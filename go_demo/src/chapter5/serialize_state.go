package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type State struct {
	Name     string
	Senators []string
	Water    float64
	Area     int
}

func main() {
	MA := []byte(`{"name":"Massachusetts", "area":27336,"water":25.7,
        "senators":["John kerry","Scott Brown"]}`)
	var state State
	if err := json.Unmarshal(MA, &state); err != nil {
		fmt.Println(err)
	}
	fmt.Println(state)
}

func (state State) String() string {
	var senators []string
	for _, senator := range state.Senators {
		senators = append(senators, fmt.Sprintf("%q", senator))
	}
	return fmt.Sprintf(`{"name":%q, "area":%d, "water":%f, "senators":[%s]}`,
		state.Name, state.Area, state.Water, strings.Join(senators, ","))
}
