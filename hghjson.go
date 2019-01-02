package hghstring

import "encoding/json"

// Jsonencode 对json编码的一个封装
func Jsonencode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Jsondecode 对一个[]byte进行json解析
func Jsondecode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
