package util

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
)

func StructAssign(dst interface{}, src interface{}) {
	if dst == nil || src == nil {
		panic("structAssign dst == nil || src == nil")
	}
	bVal := reflect.ValueOf(dst).Elem() // 获取reflect.Type类型
	vVal := reflect.ValueOf(src).Elem() // 获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {

		if vVal.Field(i).IsZero() {
			continue
		}

		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		vField := vTypeOfT.Field(i)
		name := vField.Name
		bField, ok := bVal.Type().FieldByName(name)

		if ok && bField.Type == vField.Type {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}

func SaveJSON(fileName string, data interface{}) {
	saveDir := "storage/tmp/"
	_ = os.MkdirAll(saveDir, 0777)
	file, err := os.OpenFile(saveDir+fileName+".json", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln("os.OpenFile err:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err = encoder.Encode(data); err != nil {
		log.Fatalln("encoder.Encode(data) err:", err)
	}
}
