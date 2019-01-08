package realmdefense

import "encoding/json"

func ToJson(i interface{}) []byte {
	b, _ := json.Marshal(i)
	return b
}
