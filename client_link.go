/*
   Interlink Configuration Management Database - HTTP Client
   © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   Contributors to this project, hereby assign copyright in this code to the project,
   to be licensed under the same terms as the rest of the code.
*/

package ilink

// PutLink issue a Put http request with the Link data as payload to the resource URI
func (c *Client) PutLink(link *Link) (*Result, error) {
	if err := link.valid(); err != nil {
		return nil, err
	}
	uri, err := link.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}
	resp, err := c.Put(uri, link, c.addHttpHeaders)
	return result(resp, err)
}

// DeleteLink issue a Delete http request to the resource URI
func (c *Client) DeleteLink(link *Link) (*Result, error) {
	uri, err := link.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}
	resp, err := c.Delete(uri, c.addHttpHeaders)
	return result(resp, err)
}

// GetLink issue a Get http request to the resource URI
func (c *Client) GetLink(link *Link) (*Link, error) {
	uri, err := link.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}
	result, err := c.Get(uri, c.addHttpHeaders)
	if err != nil {
		return nil, err
	}
	return link.decode(result)
}

func (c *Client) GetLinks() (*LinkList, error) {
	uri, err := uriLinks(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}
	result, err := c.Get(uri, c.addHttpHeaders)
	if err != nil {
		return nil, err
	}

	list, err := decodeLinkList(result)

	defer func() {
		if ferr := result.Body.Close(); ferr != nil {
			err = ferr
		}
	}()

	return list, err
}
