/*
   Interlink Configuration Management Database - HTTP Client
   © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   Contributors to this project, hereby assign copyright in this code to the project,
   to be licensed under the same terms as the rest of the code.
*/

package ilink

// issue a Put http request with the Item Type data as payload to the resource URI
func (c *Client) PutItemType(itemType *ItemType) (*Result, error) {
	// validates item type
	if err := itemType.valid(); err != nil {
		return nil, err
	}
	uri, err := itemType.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}
	resp, err := c.Put(uri, itemType, c.addHttpHeaders)
	if resp != nil {
		return newResult(resp)
	}
	return nil, err
}

// issue a Delete http request to the resource URI
func (c *Client) DeleteItemType(itemType *ItemType) (*Result, error) {
	uri, err := itemType.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}
	resp, err := c.Delete(uri, c.addHttpHeaders)
	return result(resp, err)
}

// issue a Get http request to the resource URI
// itemType: an instance of the Item Type with the key of the item to retrieve
func (c *Client) GetItemType(itemType *ItemType) (*ItemType, error) {
	uri, err := itemType.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}
	result, err := c.Get(uri, c.addHttpHeaders)
	if err != nil {
		return nil, err
	}
	it, err := itemType.decode(result)
	defer func() {
		if ferr := result.Body.Close(); ferr != nil {
			err = ferr
		}
	}()
	return it, err
}
