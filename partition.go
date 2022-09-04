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

type PartitionList struct {
	Values []Partition
}

func (list *PartitionList) reader() (*bytes.Reader, error) {
	jsonBytes, err := ToJson(list)
	return bytes.NewReader(jsonBytes), err
}

// the Partition resource
type Partition struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	Version     int64  `json:"version"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	ChangedBy   string `json:"changedBy"`
}

// Get the Partition in the http Response
func (partition *Partition) decode(response *http.Response) (*Partition, error) {
	result := new(Partition)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// Get the FQN for the item resource
func (partition *Partition) uri(baseUrl string) (string, error) {
	if len(partition.Key) == 0 {
		return "", fmt.Errorf("the partition does not have a key: cannot construct Partition resource URI")
	}
	return fmt.Sprintf("%s/partition/%s", baseUrl, partition.Key), nil
}

// Get a JSON bytes reader for the Serializable
func (partition *Partition) reader() (*bytes.Reader, error) {
	jsonBytes, err := partition.bytes()
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(*jsonBytes), err
}

// Get a []byte representing the Serializable
func (partition *Partition) bytes() (*[]byte, error) {
	b, err := ToJson(partition)
	return &b, err
}

func (partition *Partition) valid() error {
	if len(partition.Key) == 0 {
		return fmt.Errorf("partition key is missing")
	}
	if len(partition.Name) == 0 {
		return fmt.Errorf("partition name is missing")
	}
	return nil
}
