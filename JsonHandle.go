package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	// Json unmarshal to struct
	var s Serverslice
	str := `{"servers":[{"serverName": "Beijing_VPN", "serverIP": "127.0.0.1"},{"serverName": "Shanghai_VPN", "serverIP": "127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println("Json after unmarshal", s.Servers[0].ServerIP)
	fmt.Println("Json after unmarshal", s.Servers[1].ServerName)


	// Json unmarshal to interface{}
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez","Morticial"]}`)
	var f interface{}

	json.Unmarshal(b, &f)
	// handleError("Error when unmarshal json to interface{}", err)

	m := f.(map[string]interface{})
	for k,v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, " is string ", vv)
		case int: // age do not judged to int ?????
			fmt.Println(k, " is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, " is of a type I don`t know how to handle")
		}
	}


}
