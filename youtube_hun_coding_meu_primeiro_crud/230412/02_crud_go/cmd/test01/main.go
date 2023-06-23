package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type payloadType struct {
	Name     string
	Email    string
	Password string
	Age      int
}

func main() {

	url := "http://localhost:8080/createUser"
	payload := payloadType{
		Name:     "Bob Marley",
		Email:    "bob@test.go",
		Password: "1234!abc",
		Age:      10,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
}
