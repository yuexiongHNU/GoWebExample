package main

import (
	"encoding/xml"
	"os"
	"io/ioutil"
	"fmt"
	"bufio"
)

type Recurlyservers struct {
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
	Description string `xml:",innerxml"`
}

type server struct {
	XMLName	xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}

type servers2 struct {
	XMLName xml.Name `xml:"servers2"`
	Version string `xml:"version,attr"`
	Comment string `xml:",comment"`
	CharacterData string `xml:",chardata"`
	Region1 string `xml:"region>region1"`
	Region2 string `xml:"region>region_2"`
	Svs []server2 `xml:"server2"`
}

type server2 struct {
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}

func main() {
	file, err := os.Open("servers.xml")
	handleError("Error when open xml: ", err)
	defer file.Close()

	reader := bufio.NewReader(file)
	data, err := ioutil.ReadAll(reader)
	handleError("Error when read file:", err)

	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	handleError("Error when unmarshal:", err)

	fmt.Println("###### V #######", v)

	s := servers2{}
	s.Version = "1"
	s.Comment = "This is a comment"
	s.Region1 = "Beijing"
	s.Region2 = "Shanghai"
	s.CharacterData = "gb2312"
	s.Svs = append(s.Svs, server2{ "Shanghai_VPN", "127.0.0.1"})
	s.Svs = append(s.Svs, server2{"Beijing_VPN", "127.0.0.2"})

	output, err := xml.MarshalIndent(s, " ", "  ")
	handleError("Error when marshal to xml:", err)
	// Add header info
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}

func handleError(customInfo string, err error) {
	if err != nil {
		fmt.Println(customInfo, err)
		os.Exit(1)
	}
}
