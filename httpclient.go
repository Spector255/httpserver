package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"bytes"
	"fmt"
	"os"
)

func command(filename string) []byte{
	location, err := os.Open(filename)
	if err != nil {
		fmt.Println("error", err)
	}
	byteValue, err := ioutil.ReadAll(location)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(byteValue)
	return byteValue

}

func main() {
	client := &http.Client{}
	
	var body bytes.Buffer


	req := ""
	inp := ""

	fmt.Println("Request: ")
	fmt.Scanln(&req)

	fmt.Print("(data/update/delete/read.json): ")
	fmt.Scanln(&inp)

	var text = command(inp)

	fmt.Println(text)
	
	body.Write([]byte(text))
	
	req, err := http.NewRequest("PUT", "http://localhost:8080/", &body)
	
	resp, err := client.Do(req)
	if err != nil {panic(err) }
	
	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	
	err = ioutil.WriteFile("p.html", data, 0666)
	if err != nil {panic(err) }
	
	
}
