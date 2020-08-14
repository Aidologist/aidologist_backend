package moohttp

import (
	"encoding/json"
	"strconv"
)

type MooRequest struct {
	CtxInput	[]byte
	Data 		map[string]interface{}
}

func (m *MooRequest) RequestBodyToMap() {
	requestBody := make(map[string]interface{})
	var unmarshal interface{}
	json.Unmarshal(m.CtxInput, &unmarshal)
	if unmarshal != nil {
		var mm map[string]interface{}
		mm = unmarshal.(map[string]interface{})
		for k, v := range mm {
			requestBody[k] = v
		}
	}
	m.Data = requestBody
}

func (m *MooRequest) ReadString(s string) string {
	var requestBody = m.Data
	return requestBody[s].(string)
}

func (m *MooRequest) ReadInt(s string) int {
	var requestBody = m.Data
	i, _ := strconv.Atoi(requestBody[s].(string))
	return i
}