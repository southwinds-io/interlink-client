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

type LinkTypeAttributeList struct {
	Values []LinkTypeAttribute
}

func (list *LinkTypeAttributeList) reader() (*bytes.Reader, error) {
	jsonBytes, err := ToJson(list)
	return bytes.NewReader(jsonBytes), err
}

type LinkTypeAttribute struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	DefValue    string `json:"defValue"`
	Required    bool   `json:"required"`
	Regex       string `json:"regex"`
	LinkTypeKey string `json:"linkTypeKey"`
	Version     int64  `json:"version"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	ChangedBy   string `json:"changedBy"`
}

// Get the Link Type Attribute in the http Response
func (typeAttr *LinkTypeAttribute) decode(response *http.Response) (*LinkTypeAttribute, error) {
	result := new(LinkTypeAttribute)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// Get the FQN for the item type attribute resource
func (typeAttr *LinkTypeAttribute) uri(baseUrl string) (string, error) {
	if len(typeAttr.LinkTypeKey) == 0 {
		return "", fmt.Errorf("the link type attribute does not have an link type key: cannot construct itemtype attr resource URI")
	}
	if len(typeAttr.Key) == 0 {
		return "", fmt.Errorf("the link type attribute does not have a key: cannot construct itemtype attr resource URI")
	}
	return fmt.Sprintf("%s/linktype/%s/attribute/%s", baseUrl, typeAttr.LinkTypeKey, typeAttr.Key), nil
}

// Get a JSON bytes reader for the Serializable
func (typeAttr *LinkTypeAttribute) reader() (*bytes.Reader, error) {
	jsonBytes, err := typeAttr.bytes()
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(*jsonBytes), err
}

// Get a []byte representing the Serializable
func (typeAttr *LinkTypeAttribute) bytes() (*[]byte, error) {
	b, err := ToJson(typeAttr)
	return &b, err
}

func (typeAttr *LinkTypeAttribute) valid() error {
	if len(typeAttr.Key) == 0 {
		return fmt.Errorf("link type attribute key is missing")
	}
	if len(typeAttr.LinkTypeKey) == 0 {
		return fmt.Errorf("link type key is missing")
	}
	return nil
}
