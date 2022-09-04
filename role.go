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

type RoleList struct {
	Values []Role
}

func (list *RoleList) reader() (*bytes.Reader, error) {
	jsonBytes, err := ToJson(list)
	return bytes.NewReader(jsonBytes), err
}

// the Role resource
type Role struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	Level       int    `json:"level"`
	Version     int64  `json:"version"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	ChangedBy   string `json:"changedBy"`
}

// Get the Role in the http Response
func (role *Role) decode(response *http.Response) (*Role, error) {
	result := new(Role)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// Get the FQN for the item resource
func (role *Role) uri(baseUrl string) (string, error) {
	if len(role.Key) == 0 {
		return "", fmt.Errorf("the role does not have a key: cannot construct Role resource URI")
	}
	return fmt.Sprintf("%s/role/%s", baseUrl, role.Key), nil
}

// Get a JSON bytes reader for the Serializable
func (role *Role) reader() (*bytes.Reader, error) {
	jsonBytes, err := role.bytes()
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(*jsonBytes), err
}

// Get a []byte representing the Serializable
func (role *Role) bytes() (*[]byte, error) {
	b, err := ToJson(role)
	return &b, err
}

func (role *Role) valid() error {
	if len(role.Key) == 0 {
		return fmt.Errorf("role key is missing")
	}
	if len(role.Name) == 0 {
		return fmt.Errorf("role name is missing")
	}
	return nil
}
