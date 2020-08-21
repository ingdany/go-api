package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8000/user/delete/3"

	var jsonStr= []byte(`{ "username":"danilitro", "first_name":"Daniel", "last_name":"Perez" }`)
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))

	fmt.Println("URL:>", url)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body,_ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}