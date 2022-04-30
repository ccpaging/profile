package config

import (
	"reflect"
	"strings"
)

func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

func cloneMap(in map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	for k, v := range in {
		vm, ok := v.(map[string]interface{})
		if ok {
			out[k] = cloneMap(vm)
		} else {
			out[k] = v
		}
	}
	return out
}

func flattenKey(section, key string) string {
	if section == "" {
		return strings.ToUpper(key)
	}
	if key == "" {
		return strings.ToUpper(section)
	}
	return strings.ToUpper(section + DELIMITER + key)
}

// flattenMap a nested map so that the result is a single map with keys corresponding to the
// path through the original map. For example,
// {
//     "a": {
//         "b": 1
//     },
//     "c": "sea"
// }
// would flatten to
// {
//     "a.b": 1,
//     "c": "sea"
// }
func flattenMap(in map[string]interface{}, leadKey string) map[string]interface{} {
	out := make(map[string]interface{})

	for key, value := range in {
		if valueAsMap, ok := value.(map[string]interface{}); ok {
			sub := flattenMap(valueAsMap, leadKey)

			for subKey, subValue := range sub {
				out[flattenKey(leadKey+key, subKey)] = subValue
			}
		} else {
			out[flattenKey(leadKey+key, "")] = value
		}
	}

	return out
}
