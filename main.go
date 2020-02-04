package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

type AttestationStatement struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	RawId    string `json:"rawId"`
	Response struct {
		AttestationObject string `json:"attestationObject"`
		ClientDataJSON    string `json:"clientDataJSON"`
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Print("Usage: go run main.go <attestation json>")
		return
	}

	var attStmt AttestationStatement
	err := json.Unmarshal([]byte(args[0]), &attStmt)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", attStmt)
}