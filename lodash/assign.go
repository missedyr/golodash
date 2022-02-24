package lodash

import (
	"encoding/json"
	"reflect"
)

func Assign(mapObj interface{}, sources ...interface{}) map[string]interface{} {
	var mapNew map[string]interface{}
	if reflect.ValueOf(mapObj).Kind() != reflect.Map {
		return mapNew
	}

	j, _ := json.Marshal(mapObj)
	_ = json.Unmarshal(j, &mapNew)

	for _, obj := range sources {
		if reflect.ValueOf(obj).Kind() != reflect.Map {
			continue
		}
		objNew := map[string]interface{}{}
		objM, _ := json.Marshal(obj)
		_ = json.Unmarshal(objM, &objNew)

		for k, v := range objNew {
			mapNew[k] = v
		}
	}
	return mapNew
}
