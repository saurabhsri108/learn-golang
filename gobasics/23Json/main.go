package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 399, "LearnCodeOnline.in", "abd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "abe123", nil},
	}

	finalJson, err := json.MarshalIndent(courses, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"name": "MERN Bootcamp",
		"price": 399,
		"website": "LearnCodeOnline.in"
	}
	`)
	// var course course
	// checkValid := json.Valid(jsonDataFromWeb)
	// if checkValid {
	// 	fmt.Println("JSON was valid")
	// 	json.Unmarshal(jsonDataFromWeb, &course)
	// 	fmt.Printf("%#v\n", course)
	// } else {
	// 	fmt.Println("Json wasn't valid")
	// }

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v, value is %v and type is %T\n", k, v, v)
	}
}
