package hghstring

import (
	"fmt"
	"testing"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func TestJsonencode(t *testing.T) {
	m := Message{"Alice", "Hello", 120001}

	b, err := Jsonencode(m)
	if err != nil {
		t.Errorf("Json Encode Failed!\n")
		return
	}
	fmt.Println("Encode: ", string(b))
}

func TestJsondecode(t *testing.T) {
	b := []byte(`{"Name": "Bob", "Body": "Hello", "Time": 123456}`)
	var m Message

	err := Jsondecode(b, &m)
	if err != nil {
		t.Errorf("Decode Failed!\n")
		return
	}
	fmt.Println("Message:", m)

	var i interface{}

	err1 := Jsondecode(b, &i)
	if err1 != nil {
		t.Errorf("Decode Failed!\n")
		return
	}
	m1 := i.(map[string]interface{})

	for k, v := range m1 {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string ", vv)
		case float64:
			fmt.Println(k, "is number ", vv)
		case []interface{}:
			fmt.Println(k, "is an object")
		default:
			fmt.Println(k, "unknown Type")
		}
	}
}
