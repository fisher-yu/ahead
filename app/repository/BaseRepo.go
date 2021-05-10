package repository

import (
	"reflect"
	"time"
)

type BaseRepo struct {
}

// 将validator的struct映射到model的struct
func (r *BaseRepo) BindModel(validator interface{}, model interface{}) {
	pv := reflect.ValueOf(validator).Elem()
	mv := reflect.ValueOf(model).Elem()

	pt := pv.Type()

	for i := 0; i < pv.NumField(); i++ {
		name := pt.Field(i).Name
		if mv.FieldByName(name).IsValid() {
			//注意，这里必须类型相同，int64和int不能映射
			mv.FieldByName(name).Set(reflect.ValueOf(pv.Field(i).Interface()))
		}
	}
}

// 格式化时间
func (r *BaseRepo) MarshalTime(t time.Time) string {
	//当返回时间为空时，需特殊处理
	if time.Time(t).IsZero() {
		return ""
	}

	return time.Time(t).Format("2006-01-02 15:04:05")
}
