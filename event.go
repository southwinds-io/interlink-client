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
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// EventManager MQTT client for change notifications
type EventManager struct {
	done   chan bool
	cfg    *EventConfig
	client MQTT.Client
}

// NewEventManager creates a new event manager subscribed to a specific topic
// cfg: the mqtt server configuration
func NewEventManager(cfg *EventConfig) (*EventManager, error) {
	// check the configuration is valid (preconditions)
	if ok, err := cfg.isValid(); !ok {
		return nil, err
	}
	m := new(EventManager)
	// create connection configuration
	connOpts := MQTT.NewClientOptions().AddBroker(cfg.Server).SetClientID(cfg.clientId()).SetCleanSession(true)
	// add credentials if provided
	if cfg.hasCredentials() {
		connOpts.SetUsername(cfg.Username)
		connOpts.SetPassword(cfg.Password)
	}
	// setup tls configuration
	tlsConfig := &tls.Config{
		InsecureSkipVerify: cfg.InsecureSkipVerify,
		ClientAuth:         cfg.ClientAuthType,
	}
	connOpts.SetTLSConfig(tlsConfig)
	// subscribe to the topic on connection
	connOpts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(cfg.topic(), byte(cfg.Qos), cfg.OnMsgReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	// finally create the client
	client := MQTT.NewClient(connOpts)
	// set up the manager
	m.client = client
	m.cfg = cfg
	// return a new setup manager
	return m, nil
}

// Connect to the message broker
func (m *EventManager) Connect() error {
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	fmt.Printf("Connected to %s\n", m.cfg.Server)
	return nil
}

// Disconnect from the message broker
func (m *EventManager) Disconnect(timeoutMilSecs uint) {
	m.client.Disconnect(timeoutMilSecs)
}
