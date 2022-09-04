/*
   Interlink Configuration Management Database - HTTP Client
   © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   Contributors to this project, hereby assign copyright in this code to the project,
   to be licensed under the same terms as the rest of the code.
*/

package ilink

// issue a Put http request with the Membership data as payload to the resource URI
func (c *Client) PutMembership(member *Membership) (*Result, error) {
	if err := member.valid(); err != nil {
		return nil, err
	}
	uri, err := member.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}
	resp, err := c.Put(uri, member, c.addHttpHeaders)
	if resp != nil {
		return newResult(resp)
	}
	return nil, err
}

// issue a Delete http request to the resource URI
func (c *Client) DeleteMembership(member *Membership) (*Result, error) {
	uri, err := member.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}
	resp, err := c.Delete(uri, c.addHttpHeaders)
	return result(resp, err)
}

// issue a Get http request to the resource URI
func (c *Client) GetMembership(member *Membership) (*Membership, error) {
	uri, err := member.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}
	result, err := c.Get(uri, c.addHttpHeaders)
	if err != nil {
		return nil, err
	}
	i, err := member.decode(result)
	defer func() {
		if ferr := result.Body.Close(); ferr != nil {
			err = ferr
		}
	}()
	return i, err
}
