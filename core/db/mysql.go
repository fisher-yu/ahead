package db

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Policy int

const (
	RandomPolicy Policy = iota
	WeightRandomPolicy
	RoundRobinPolicy
	WeightRoundRobinPolicy
	LeastConnPolicy
)

const (
	DefaultMaxIdleConn     = 5
	DefaultMaxOpenConn     = 10
	DefaultConnMaxLifeTime = 5 * time.Minute
)

type Engine struct {
	*xorm.EngineGroup
}

type mysqlConfig struct {
	dataSource []string
	user       string
	password   string
	database   string

	*mysqlOptionConfig
}

type mysqlOptionConfig struct {
	maxIdleConn     int
	maxOpenConn     int
	maxConnLifetime time.Duration
	policy          xorm.GroupPolicy
	showLog         bool
	pingInterval    time.Duration
}

func parseConfig(cfg map[string]interface{}) (*mysqlConfig, error) {
	if cfg["host"] == nil {
		return nil, errors.New("host is not found")
	}
	if cfg["user"] == nil {
		return nil, errors.New("user is not found")
	}
	if cfg["password"] == nil {
		return nil, errors.New("password is not found")
	}
	if cfg["database"] == nil {
		return nil, errors.New("database is not found")
	}
	option := &mysqlOptionConfig{
		maxIdleConn:     DefaultMaxIdleConn,
		maxOpenConn:     DefaultMaxOpenConn,
		maxConnLifetime: DefaultConnMaxLifeTime,
		pingInterval:    0,
		policy:          xorm.RandomPolicy(),
		showLog:         false,
	}
	if cfg["max_idle_conn"] != nil {
		option.maxIdleConn = cfg["max_idle_conn"].(int)
	}
	if cfg["max_open_conn"] != nil {
		option.maxOpenConn = cfg["max_open_conn"].(int)
	}
	if cfg["max_conn_lifetime"] != nil {
		option.maxConnLifetime = time.Duration(cfg["max_conn_lifetime"].(int64)) * time.Minute
	}
	if cfg["policy"] != nil {
		switch cfg["policy"] {
		case RandomPolicy:
			option.policy = xorm.RandomPolicy()
		case WeightRandomPolicy:
			if cfg["policy_weight"] == nil {
				return nil, errors.New("policy weight is not found")
			}
			weights, e := parsePolicyWeightConf(cfg["policy_weight"].(string))
			if e != nil {
				return nil, errors.New("parse policy weight error")
			}
			option.policy = xorm.WeightRandomPolicy(weights)
		case RoundRobinPolicy:
			option.policy = xorm.RoundRobinPolicy()
		case WeightRoundRobinPolicy:
			if cfg["policy_weight"] == nil {
				return nil, errors.New("policy weight is not found")
			}
			weights, e := parsePolicyWeightConf(cfg["policy_weight"].(string))
			if e != nil {
				return nil, errors.New("parse policy weight error")
			}
			option.policy = xorm.WeightRoundRobinPolicy(weights)
		case LeastConnPolicy:
			option.policy = xorm.LeastConnPolicy()
		default:
			return nil, errors.New("policy error")
		}

	}

	if cfg["show_log"] != nil && cfg["show_log"] == true {
		option.showLog = true
	}

	if cfg["ping_interval"] != nil {
		interval := cfg["ping_interval"].(int64)
		option.pingInterval = time.Duration(interval) * time.Second
	}

	return &mysqlConfig{
		mysqlOptionConfig: option,
		dataSource:        parseDataSource(cfg),
	}, nil
}

func parsePolicyWeightConf(s string) ([]int, error) {
	var weights []int
	ws := strings.Split(s, ",")
	for _, w := range ws {
		i, e := strconv.Atoi(w)
		if e != nil {
			return nil, errors.New("policy_weight parse error")
		}
		weights = append(weights, i)
	}
	return weights, nil
}

// dataSource: user:password@tcp(host:port)/database?charset=utf8&parseTime=True&loc=Local
func parseDataSource(cfg map[string]interface{}) []string {
	var ds []string
	hosts := strings.Split(cfg["host"].(string), ",")
	for _, h := range hosts {
		ds = append(ds, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg["user"], cfg["password"], h, cfg["database"]))
	}
	return ds
}

func New(config map[string]interface{}) (*Engine, error) {
	cfg, err := parseConfig(config)
	if err != nil {
		return nil, err
	}

	eg, err := xorm.NewEngineGroup("mysql", cfg.dataSource, cfg.policy)
	if err != nil {
		return nil, err
	}
	err = eg.Ping()
	if err != nil {
		return nil, err
	}

	// 设置空闲连接池中的最大连接数
	eg.SetMaxIdleConns(cfg.maxIdleConn)
	// 设置数据库连接最大打开数
	eg.SetMaxOpenConns(cfg.maxOpenConn)
	// 设置可重用连接的最长时间，一定要小于mysql服务端的保持超时时间，否则可能会被服务端关闭
	eg.SetConnMaxLifetime(cfg.maxConnLifetime)
	//设置负载策略
	eg.SetPolicy(cfg.policy)

	if cfg.showLog {
		eg.ShowSQL(true)
	}

	// 如果设置了ping
	if cfg.pingInterval > 0 {
		go keepConnAlive(eg, cfg.pingInterval)
	}

	return &Engine{
		eg,
	}, nil
}

func keepConnAlive(eg *xorm.EngineGroup, interval time.Duration) {
	t := time.Tick(interval)
	var err error
	for {
		<-t
		err = eg.Ping()
		if err != nil {
			fmt.Println("database ping err :", err.Error())
		}
	}
}
