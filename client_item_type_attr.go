/*
   Interlink Configuration Management Database - HTTP Client
   © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   Contributors to this project, hereby assign copyright in this code to the project,
   to be licensed under the same terms as the rest of the code.
*/

package ilink

// issue a Put http request with the Item Type Attribute data as payload to the resource URI
func (c *Client) PutItemTypeAttr(typeAttr *ItemTypeAttribute) (*Result, error) {
	if err := typeAttr.valid(); err != nil {
		return nil, err
	}
	uri, err := typeAttr.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}
	resp, err := c.Put(uri, typeAttr, c.addHttpHeaders)
	if resp != nil {
		if err == nil {
			return newResult(resp)
		}
	}
	return nil, err
}

// issue a Delete http request to the resource URI
func (c *Client) DeleteItemTypeAttr(typeAttr *ItemTypeAttribute) (*Result, error) {
	uri, err := typeAttr.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}
	resp, err := c.Delete(uri, c.addHttpHeaders)
	return result(resp, err)
}

// issue a Get http request to the resource URI
func (c *Client) GetItemTypeAttr(typeAttr *ItemTypeAttribute) (*ItemTypeAttribute, error) {
	uri, err := typeAttr.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}
	result, err := c.Get(uri, c.addHttpHeaders)
	if err != nil {
		return nil, err
	}
	return typeAttr.decode(result)
}
