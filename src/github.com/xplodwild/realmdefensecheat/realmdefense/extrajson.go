package realmdefense

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ExtraJSON is a special base structure that allows us to keep any extra field that wouldn't yet be in
// our data structure, and be able to marshal them back to JSON later without loosing any data.

type _SaveData SaveData

func (t SaveData) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})

	// Take everything in Extra
	for k, v := range t.Extra {
		data[k] = v
	}

	// Take all the struct values with a json tag
	val := reflect.ValueOf(t)
	typ := reflect.TypeOf(t)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldv := val.Field(i)
		jsonTag := strings.Split(field.Tag.Get("json"), ",")[0]
		if jsonTag != "" && jsonTag != "-" {
			data[jsonTag] = fieldv.Interface()
		}
	}
	return json.Marshal(data)
}

func (t *SaveData) UnmarshalJSON(b []byte) error {
	t2 := _SaveData{}
	err := json.Unmarshal(b, &t2)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &(t2.Extra))
	if err != nil {
		return err
	}

	typ := reflect.TypeOf(t2)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := strings.Split(field.Tag.Get("json"), ",")[0]
		if jsonTag != "" && jsonTag != "-" {
			delete(t2.Extra, jsonTag)
		}
	}

	*t = SaveData(t2)

	return nil
}
