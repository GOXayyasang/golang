package utils

import (
	"reflect"
	"time"
)

func MapInterfaceString(key string, model map[string]interface{}) string {
	if value, ok := model[key].(string); ok {
		return value
	} else {
		return ""
	}
}

func MapInterfaceInteger(key string, model map[string]interface{}) *int {
	zero := 0
	if value, ok := model[key].(int); ok {
		return &value
	} else {
		return &zero
	}
}

func CustomDecoderHook(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if t == reflect.TypeOf(time.Time{}) {
		// Specify the layout for the timestamp format
		layout := "2006-01-02T15:04:05.999Z"
		if str, ok := data.(string); ok {
			parsedTime, err := time.Parse(layout, str)
			if err != nil {
				return time.Time{}, err
			}
			return parsedTime, nil
		}
	}
	return data, nil
}
