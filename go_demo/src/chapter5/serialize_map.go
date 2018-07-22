package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	MA := []byte(`{"name":"Massachusetts", "area":27336,"water":25.7,
        "senators":["John kerry","Scott Brown"]}`)
	var object interface{}
	if err := json.Unmarshal(MA, &object); err != nil {
		fmt.Println(err)
	} else {
		jsonObject := object.(map[string]interface{})
		fmt.Println(jsonObjectAsString(jsonObject))
	}
}

func jsonObjectAsString(jsonObject map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	comma := ""
	for key, value := range jsonObject {
		buffer.WriteString(comma)
		switch value := value.(type) { //影子变量
		case nil:
			fmt.Fprintf(&buffer, "%q: null", key)
		case bool:
			fmt.Fprintf(&buffer, "%q: %t", key, value)
		case float64:
			fmt.Fprintf(&buffer, "%q: %f", key, value)
		case string:
			fmt.Fprintf(&buffer, "%q: %q", key, value)
		case []interface{}:
			fmt.Fprintf(&buffer, "%q: [", key)
			innerComma := ""
			for _, s := range value {
				if s, ok := s.(string); ok { //影子变量
					fmt.Fprintf(&buffer, "%s%q", innerComma, s)
					innerComma = ","
				}
			}
			buffer.WriteString("]")
		}
		comma = ","
	}

	buffer.WriteString("}")
	return buffer.String()
}
