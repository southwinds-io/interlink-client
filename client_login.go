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
	"encoding/json"
	"fmt"
	"io/ioutil"
	h "southwinds.dev/http"
	"strings"
	"time"
)

// Login check that the user is authenticated using the CMDB as user store
// and returns a list of access controls for the user
func (c *Client) Login(credentials *Login) (*h.UserPrincipal, error) {
	// validates user
	if err := credentials.valid(); err != nil {
		return nil, err
	}
	uri, err := credentials.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}
	resp, err := c.Post(uri, credentials, c.addHttpHeaders)
	// if there is a technical error
	if err != nil {
		return nil, fmt.Errorf("login failed for user '%s' due to error: '%s'\n", credentials.Username, err)
	}
	// if the response was unauthorised, login failed
	if resp.StatusCode == 401 {
		return nil, fmt.Errorf("authentication failed for user '%s'\n", credentials.Username)
	}
	// otherwise, get the list of controls from the user information
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("user was authenticated but failed to read response body for user '%s'; cannot return list of access controls; error: '%s'\n", credentials.Username, err)
	}
	var user User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, fmt.Errorf("user was authenticated but failed to unmarhsal response body for user '%s'; cannot return list of access controls; error: '%s'\n", credentials.Username, err)
	}
	controls, err := newControls(user.ACL)
	if err != nil {
		return nil, fmt.Errorf("user was authenticated but failed to parse access controls for user '%s': '%s'\n", credentials.Username, err)
	}
	// constructs a principal and returns
	return &h.UserPrincipal{
		Username: credentials.Username,
		Rights:   controls,
		Created:  time.Now(),
	}, nil
}

func newControls(acl string) (h.Controls, error) {
	var ctls h.Controls
	// if acl is empty then return an empty list of controls
	if len(strings.Trim(acl, " ")) == 0 {
		return h.Controls{}, nil
	}
	parts := strings.Split(acl, ",")
	for _, part := range parts {
		control, err := newControl(part)
		if err != nil {
			return nil, err
		}
		ctls = append(ctls, control)
	}
	return ctls, nil
}

func newControl(ac string) (h.Control, error) {
	parts := strings.Split(ac, ":")
	if len(parts) != 3 {
		return h.Control{}, fmt.Errorf("Invalid control format '%s', it should be realm:uri:method\n", ac)
	}
	return h.Control{
		Realm:  parts[0],
		URI:    parts[1],
		Method: strings.Split(parts[2], "|"),
	}, nil
}
