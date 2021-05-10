package core

import (
	"ahead/core/db"
	"ahead/util/config"
)

type Model struct {
	DB *db.Engine
}

func NewModel() *Model {
	var de *db.Engine
	err := config.InitConfig()
	if err != nil {
		panic(err.Error())
	}
	if config.IsSet("mysql") {
		var err error
		cfg := config.GetStringMap("mysql")
		de, err = db.New(cfg)
		if err != nil {
			panic(err)
		}
	}

	return &Model{DB: de}
}
