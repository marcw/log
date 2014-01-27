package gogol

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type encodeState struct {
	buffer     *bytes.Buffer
	writeBegan bool
}

// Marshal will encode any structure or map in a k=v format
func Marshal(data interface{}) ([]byte, error) {
	e := &encodeState{}
	err := e.marshal("", data)
	return e.buffer.Bytes(), err
}

// MarshalString is a proxy method for Marshal that drops any error and always return a string
func MarshalString(data interface{}) string {
	e := &encodeState{}
	e.marshal("", data)
	return e.buffer.String()
}

func (e *encodeState) marshal(prefix string, data interface{}) error {
	if e.writeBegan {
		e.writeBegan = false
	}
	if e.buffer == nil {
		e.buffer = bytes.NewBuffer(nil)
	}

	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		return e.marshalPtr(prefix, value.Interface())
	}

	if value.Kind() == reflect.Struct {
		return e.marshalStruct(prefix, value.Interface())
	}

	if value.Kind() == reflect.Map {
		return e.marshalMap(prefix, value.Interface())
	}

	return fmt.Errorf("don't know how to marshal this type")
}

func (e *encodeState) writeTab() error {
	if !e.writeBegan {
		return nil
	}

	_, err := e.buffer.WriteRune('\t')
	return err
}

func (e *encodeState) marshalPtr(prefix string, data interface{}) error {
	value := reflect.ValueOf(data)
	value = value.Elem()

	if value.Kind() == reflect.Struct {
		return e.marshalStruct(prefix, value.Interface())
	}

	return fmt.Errorf("don't know how to marshal this pointer")
}

func (e *encodeState) marshalStruct(prefix string, data interface{}) error {
	ref := reflect.TypeOf(data)
	value := reflect.ValueOf(data)

	for i := 0; i < value.NumField(); i++ {
		fieldType := ref.Field(i)
		// Ignore unexported fields
		if fieldType.PkgPath != "" {
			continue
		}
		// Get the key for this field
		var key string
		if key = fieldType.Tag.Get("kv"); key == "-" {
			continue
		} else if key == "" {
			key = fieldType.Name
		}
		e.writeTab()

		field := value.Field(i)
		if field.Kind() == reflect.Ptr || field.Kind() == reflect.Struct || field.Kind() == reflect.Map {
			if len(key) > 0 {
				key = key + "."
			}
			err := e.marshal(key, field.Interface())
			if err != nil {
				return err
			}
			continue
		}

		value := stringify(field.Interface())

		if _, err := e.buffer.WriteString(fmt.Sprintf("%s%s=%s", prefix, key, value)); err != nil {
			return err
		}
		e.writeBegan = true
	}

	return nil
}

func (e *encodeState) marshalMap(prefix string, data interface{}) error {
	value := reflect.ValueOf(data)
	keys := value.MapKeys()
	for _, v := range keys {
		e.writeTab()
		mapKey, _ := strconv.Unquote(stringify(v.Interface()))
		mapValue := stringify(value.MapIndex(v).Interface())
		if _, err := e.buffer.WriteString(fmt.Sprintf("%s%s=%s", prefix, mapKey, mapValue)); err != nil {
			return err
		}
		e.writeBegan = true
	}

	return nil
}

func stringify(data interface{}) string {
	value := ""
	switch t := data.(type) {
	case bool:
		value = strconv.FormatBool(t)
	case rune:
		value = strconv.QuoteRune(t)
	case time.Time:
		value = t.UTC().Format(time.RFC3339)
	case int:
		value = strconv.Itoa(int(t))
	case fmt.Stringer:
		value = t.String()
	case string:
		value = strconv.Quote(t)
	case []string:
		value = strconv.Quote(strings.Join(t, ";"))
	default:
		value = "n/a"
	}
	return value
}
