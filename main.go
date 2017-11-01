package main

import (
	"net/http"
	"fmt"
	"io/ioutil"

	"bytes"
	"encoding/json"
)
var token = "eyJ1c2VyX2F1dGhlbnRpY2F0aW9uX2lkIjoyNjAxODd9:1e9fUV:y2RMjwvBSDn6eKu42sDony85Wyc"

type Project struct {
	Description string `json:"description"`
	Name string `json:"name"`
}

func main() {

	body := &Project{
		Name: "Golang Test",
		Description: "Testing",
	}

	client := &http.Client{}

	bodySend, err := json.Marshal(body)
	if err != nil{
		fmt.Printf("Error Marshelling Project JSON")
	}

	req, err := http.NewRequest("POST", "https://api.taiga.io/api/v1/projects", bytes.NewBuffer(bodySend))
	if err != nil{
		fmt.Printf("Error Creating new http request")
	}


	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil{
		fmt.Printf("Error Creating a project")
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Printf("Error reading response body")
	}


	fmt.Printf("Response Body: %v", string(respBody))

}