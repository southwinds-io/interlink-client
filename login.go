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
	"fmt"
)

// Login information for users authenticating with client devices such as web browsers
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Get a JSON bytes reader for the Serializable
func (login *Login) reader() (*bytes.Reader, error) {
	jsonBytes, err := login.bytes()
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(*jsonBytes), err
}

// Get a []byte representing the Serializable
func (login *Login) bytes() (*[]byte, error) {
	b, err := ToJson(login)
	return &b, err
}

// Get the FQN for the item resource
func (login *Login) uri(baseUrl string) (string, error) {
	return fmt.Sprintf("%s/login", baseUrl), nil
}

func (login *Login) valid() error {
	if len(login.Username) == 0 {
		return fmt.Errorf("username is missing")
	}
	if len(login.Password) == 0 {
		return fmt.Errorf("user password is missing")
	}
	return nil
}
