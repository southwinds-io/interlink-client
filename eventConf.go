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
	"crypto/tls"
	"errors"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"os"
)

// configuration for the event manager (mqtt broker)
type EventConfig struct {
	// the MQTT Server url
	Server string
	// the item type for which to get notification changes (itemInstance must be empty)
	ItemType string
	// the item instance for which to get notification changes (ItemType must be empty)
	ItemInstance string
	// the quality of service for message delivery - 0: at most once, 1: at least once, 2: exactly once
	Qos int
	// authentication Username
	Username string
	// authentication Password
	Password string
	// skip tls certificate verification
	InsecureSkipVerify bool
	// the policy the Server will follow for TLS Client Authentication
	ClientAuthType tls.ClientAuthType
	// a function to process received messages
	OnMsgReceived MQTT.MessageHandler
}

func (c *EventConfig) hasCredentials() bool {
	return len(c.Username) > 0 && len(c.Password) > 0
}

func (c *EventConfig) topic() string {
	if len(c.ItemInstance) > 0 {
		return fmt.Sprintf("II_%s", c.ItemInstance)
	}
	return fmt.Sprintf("IT_%s", c.ItemType)
}

// unique identifier for the client
func (c *EventConfig) clientId() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-host"
	}
	return fmt.Sprintf("%s-%s-%s", c.topic(), hostname, uuid.New())
}

// check the configuration is valid
func (c *EventConfig) isValid() (bool, error) {
	if len(c.Server) == 0 {
		return false, errors.New("server property not provided")
	}
	if len(c.ItemInstance) > 0 && len(c.ItemType) > 0 {
		return false, errors.New("itemType and itemInstance both have values, only one is allowed")
	}
	if len(c.ItemInstance) == 0 && len(c.ItemType) == 0 {
		return false, errors.New("itemType and itemInstance do not have values, one is required")
	}
	if len(c.Username) > 0 && len(c.Password) == 0 {
		return false, errors.New("username with no password, provide password")
	}
	if c.OnMsgReceived == nil {
		return false, errors.New("a handler for received messages must be provided")
	}
	return true, nil
}
