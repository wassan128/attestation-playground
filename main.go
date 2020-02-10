package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/fxamacker/cbor"
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

type ParsedAttestationObject struct {
	Fmt      string        `cbor:"fmt"`
	AttStmt  ParsedAttStmt `cbor:"attStmt"`
	AuthData []byte        `cbor:"authData"`
}

type ParsedAttStmt struct {
	Sig string `json:"sig"`
	X5c string `json:"x5c"`
}

func (a AttestationStatement) ParseAttestationObject() {
	atstObjStr := strings.ReplaceAll(a.Response.AttestationObject, "-", "+")
	atstObjStr = strings.ReplaceAll(atstObjStr, "_", "/") + "=="

	atstObjBuf, err := base64.StdEncoding.DecodeString(atstObjStr)
	if err != nil {
		fmt.Println("Failed base64 decode:", err)
		return
	}
	dec := cbor.NewDecoder(bytes.NewReader(atstObjBuf))

	var parsedAttestationObj ParsedAttestationObject
	if err := dec.Decode(&parsedAttestationObj); err != nil {
		fmt.Printf("error: %+v\n", err)
		return
	}
	fmt.Printf("%+v\n", parsedAttestationObj)
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

	fmt.Printf("%+v\n", attStmt)

	attStmt.ParseAttestationObject()
}
