package dbManager

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"reflect"
	"time"
)

//数据对象公共结构
type CommonItem struct {
	UpdatedAt time.Time `ckv:"-"` //数据更新时间
}

func (s *CommonItem) GetUpdateTime() int64 {
	return s.UpdatedAt.Unix()
}

func (s *CommonItem) SetUpdateTime(t int64) {
	s.UpdatedAt = time.Unix(t, 0)
}

func (s *CommonItem) SetValue(dest, src interface{}, skipfield ...string) (bool, map[string]interface{}) {
	return SetValue(dest, src, skipfield...)
}

func SetValue(dest, src interface{}, skipfield ...string) (bool, map[string]interface{}) {
	if dest == src {
		return true, nil
	}
	valDest := reflect.ValueOf(dest)
	val := reflect.ValueOf(src)
	if val.Kind() != reflect.Ptr {
		log.Println(fmt.Errorf("Err <ParseValue> cannot use non-ptr model struct `%s`", val.Type().Name()))
		return false, nil
	}
	if valDest.Kind() != reflect.Ptr {
		log.Println(fmt.Errorf("Err <ParseValue> cannot use non-ptr model struct `%s`", valDest.Type().Name()))
		return false, nil
	}
	valDest = reflect.Indirect(valDest)
	val = reflect.Indirect(val)
	changedDa := make(map[string]interface{})
	isChange := setIndVal(valDest, val, "", false, changedDa, skipfield...)
	if isChange {
		_, ok := changedDa["UpdatedAt"]
		if len(changedDa) == 1 && ok {
			isChange = false
		}
		zap.L().Info("数据更新 设置值", zap.String("name", valDest.Type().Name()),
			zap.String("dest", fmt.Sprintf("%+v", dest)),
		)
		zap.L().Info("数据更新 设置值",
			zap.String("name", valDest.Type().Name()),
			zap.String("changed data", fmt.Sprintf("%+v", changedDa)),
		)
	}
	return isChange, changedDa
}

func setIndVal(destInd reflect.Value, src reflect.Value, destFieldName string, isChange bool,
	changedata map[string]interface{}, skipfield ...string) bool {
	/*如果只有时间更新 就跳过此次修改*/
	tpname := destInd.Type().Name()
	if destInd.Type().Kind() == reflect.Struct && tpname != "Time" {
		for i := destInd.NumField() - 1; i >= 0; i-- { //这里用倒序的方式是 让updatetime最后更新
			destFieldName := destInd.Type().Field(i).Name
			if setIndVal(destInd.Field(i), src, destFieldName, isChange, changedata) {
				isChange = true
			}
		}
	} else {
		if !(len(destFieldName) > 0 && destFieldName[0] >= 'A' && destFieldName[0] <= 'Z') {
			//fmt.Println("field 非public成员", destFieldName)
			return false
		}
		if len(skipfield) > 0 {
			for _, v := range skipfield {
				if destFieldName == v {
					return false
				}
			}
		}
		for i := src.Type().NumField() - 1; i >= 0; i-- {
			destIndFieldType := destInd.Type().Name()
			srcFieldName := src.Type().Field(i).Name
			srcFieldType := src.Field(i).Type().Name()
			if len(srcFieldName) > 0 && srcFieldName[0] >= 'a' && srcFieldName[0] <= 'z' {
				continue
			}
			if src.Field(i).Type().Kind() == reflect.Struct && srcFieldType != "Time" {
				//fmt.Println("Struct", src.Type().Field(i).Name, src.Field(i).Type().Name())
				if setIndVal(destInd, src.Field(i), destFieldName, isChange, changedata) {
					isChange = true
				}
			} else {
				//fmt.Println("destFieldName", destFieldName)
				//fmt.Println("====", destFieldName, src.Type().Field(i).Name)
				if srcFieldName == destFieldName {
					if srcFieldType == destIndFieldType && srcFieldType != "" {
						//fmt.Println("field 相等", destFieldName, src.Type().Field(i).Name, src.Field(i).Type().Name(), destInd.Type().Name())
						destIndVal := destInd.Interface()
						field := src.Field(i)
						srcVal := field.Interface()
						if !reflect.DeepEqual(destIndVal, srcVal) {
							destInd.Set(src.Field(i))
							changedata[destFieldName] = srcVal
							if destFieldName == "UpdatedAt" {
								return false
							} else {
								return true
							}
						}
					}
				}
			}
		}
	}
	return isChange
}
