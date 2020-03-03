package main

import (
	"encoding/json"
	"fmt"
)

func Output(v interface{}) {
	j, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("failed json.Marshal() v=%+v, err=%+v\n", v, err)
	}
	fmt.Printf("%s\n", string(j))
}
