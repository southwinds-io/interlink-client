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
	"fmt"
	"net/http"
)

// clear all data in the database
func (c *Client) Clear() (*Result, error) {
	resp, err := c.Delete(fmt.Sprintf("%s/clear", c.conf.BaseURI), c.addHttpHeaders)
	return result(resp, err)
}

// generic function to check for errors and retrieve a result
func result(resp *http.Response, err error) (*Result, error) {
	// if there is a response
	if resp != nil {
		// extract the response result
		result, err2 := newResult(resp)
		if err2 != nil {
			return result, err2
		}
		return result, err
	}
	return nil, err
}
