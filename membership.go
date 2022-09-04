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

type MembershipList struct {
	Values []Membership
}

func (list *MembershipList) reader() (*bytes.Reader, error) {
	jsonBytes, err := ToJson(list)
	return bytes.NewReader(jsonBytes), err
}

// the Role resource
type Membership struct {
	Key       string `json:"key"`
	User      string `json:"userKey"`
	Role      string `json:"roleKey"`
	Version   int64  `json:"version"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
	ChangedBy string `json:"changedBy"`
}

// Get the Role in the http Response
func (member *Membership) decode(response *http.Response) (*Membership, error) {
	result := new(Membership)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// Get the FQN for the item resource
func (member *Membership) uri(baseUrl string) (string, error) {
	if len(member.Key) == 0 {
		return "", fmt.Errorf("the membership does not have a key: cannot construct Membership resource URI")
	}
	return fmt.Sprintf("%s/membership/%s", baseUrl, member.Key), nil
}

// Get a JSON bytes reader for the Serializable
func (member *Membership) reader() (*bytes.Reader, error) {
	jsonBytes, err := member.bytes()
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(*jsonBytes), err
}

// Get a []byte representing the Serializable
func (member *Membership) bytes() (*[]byte, error) {
	b, err := ToJson(member)
	return &b, err
}

func (member *Membership) valid() error {
	if len(member.Key) == 0 {
		return fmt.Errorf("membership key is missing")
	}
	if len(member.Role) == 0 {
		return fmt.Errorf("role key is missing")
	}
	if len(member.User) == 0 {
		return fmt.Errorf("user key is missing")
	}
	return nil
}
