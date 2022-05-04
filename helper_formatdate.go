package handlebars

import (
	"reflect"
	"time"
)

func helperFormatdate(date interface{}, format string) string {
	reflectType := reflect.TypeOf(date)
	if reflectType.Kind() == reflect.Int {
		date = int64(date.(int))
		reflectType = reflect.TypeOf(date)
	}
	if reflectType.Kind() == reflect.Int64 {
		stamp := date.(int64)
		return time.Unix(stamp, 0).Format(format)
	}

	if reflectType.AssignableTo(reflect.TypeOf(time.Time{})) {
		stamp := date.(time.Time)
		return stamp.Format(format)
	}
	return "Invalid date"
}
