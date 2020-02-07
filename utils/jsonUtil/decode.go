package jsonUtil
import (
	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
	"reflect"
)

func Decode(msg []byte, obj interface{}) error {

	objR := reflect.TypeOf(obj)
	if objR.Kind() != reflect.Ptr {
		return errors.New("Decode: second argument must be a pointer to result object")
	}

	objT := objR.Elem()
	var objV = reflect.Indirect(reflect.ValueOf(obj))

	switch objT.Kind() {
	case reflect.Struct:
		if jsonUnmarshaller, ok := objV.Addr().Interface().(json.Unmarshaler); ok {
			if err := jsonUnmarshaller.UnmarshalJSON(msg); err != nil {
				return err
			}
			return nil
		}

		var jsonParserValue []byte
		var jsonParserLastError error
		var jsonParserDataType jsonparser.ValueType
		//var lastOffset int
		for i := 0; i < objT.NumField(); i++ {
			fieldValue := objV.Field(i)
			if !fieldValue.CanSet() {
				continue
			}

			jsonParserValue = nil
			f := objT.Field(i)

			var fieldNamesToFind = []string{f.Name}
			if tagName := f.Tag.Get("json"); tagName != "" {
				if tagName == "-" {
					continue
				}
				fieldNamesToFind = append([]string{tagName}, fieldNamesToFind...)
			}
			for _, fieldName := range fieldNamesToFind {
				jsonParserValue, jsonParserDataType, _, jsonParserLastError = jsonparser.Get(msg, fieldName)
				if jsonParserDataType != jsonparser.NotExist && jsonParserLastError != nil {
					return jsonParserLastError
				}
				if jsonParserDataType != jsonparser.NotExist {
					break
				}
			}

			if jsonParserValue == nil {
				continue
			}

			if decoding_err := Decode(jsonParserValue, fieldValue.Addr().Interface()); decoding_err != nil {
				return errors.Wrap(decoding_err, f.Name+" decoding error("+string(jsonParserValue)+")")
			}
		}
	case reflect.Slice:
		jsonparser.ArrayEach(msg, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			newSliceElem := reflect.New(objV.Type().Elem())
			err = Decode(value, newSliceElem.Interface())
			objV.Set(reflect.Append(objV, newSliceElem.Elem()))
		})

	case reflect.String:
		parsedString, parse_err := jsonparser.ParseString(msg)
		if parse_err != nil {
			return parse_err
		}
		objV.SetString(parsedString)
	case reflect.Bool:
		parsedBool, parse_err := jsonparser.ParseBoolean(msg)
		if parse_err != nil {
			return parse_err
		}
		objV.SetBool(parsedBool)
	case reflect.Float32, reflect.Float64:
		parsedFloat, parse_err := jsonparser.ParseFloat(msg)
		if parse_err != nil {
			return parse_err
		}
		objV.SetFloat(parsedFloat)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		parsedInt, parse_err := jsonparser.ParseInt(msg)
		if parse_err != nil {
			return parse_err
		}
		objV.SetInt(parsedInt)
	default:
		if jsonUnmarshaller, ok := objV.Addr().Interface().(json.Unmarshaler); ok {
			if err := jsonUnmarshaller.UnmarshalJSON(msg); err != nil {
				return err
			}
		}
	}
	return nil
}