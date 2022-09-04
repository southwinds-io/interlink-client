/*
   Interlink Configuration Management Database - HTTP Client
   © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   Contributors to this project, hereby assign copyright in this code to the project,
   to be licensed under the same terms as the rest of the code.
*/

package ilink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GraphData struct {
	Models             []Model             `json:"models"`
	ItemTypes          []ItemType          `json:"itemTypes"`
	ItemTypeAttributes []ItemTypeAttribute `json:"itemTypeAttributes"`
	LinkTypes          []LinkType          `json:"linkTypes"`
	LinkTypeAttribute  []LinkTypeAttribute `json:"linkTypeAttributes"`
	LinkRules          []LinkRule          `json:"linkRules"`
	Items              []Item              `json:"items"`
	Links              []Link              `json:"links"`
}

// Get the Item in the http Response
func (data *GraphData) decode(response *http.Response) (*GraphData, error) {
	result := new(GraphData)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// Get the FQN for the data resource
func (data *GraphData) uri(baseUrl string) (string, error) {
	return fmt.Sprintf("%s/data", baseUrl), nil
}

// Get a JSON bytes reader for the Serializable
func (data *GraphData) reader() (*bytes.Reader, error) {
	jsonBytes, err := data.bytes()
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(*jsonBytes), err
}

// Get a []byte representing the Serializable
func (data *GraphData) bytes() (*[]byte, error) {
	b, err := ToJson(data)
	return &b, err
}
