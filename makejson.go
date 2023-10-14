package main

import (
	"encoding/json"
	"fmt"
)

func MainMakeJson() {

	var name string
	var address string

	people := make(map[string]string)

	fmt.Print("Get me your name")
	fmt.Scan(&name)
	fmt.Print("Get me your address")
	fmt.Scan(&address)
	people[name] = address

	peopleJSON, _ := json.Marshal(people)

	fmt.Print(string(peopleJSON))

}
