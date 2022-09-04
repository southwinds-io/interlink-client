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
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

// how to use the event manager
func TestReceiver(t *testing.T) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	// create a new instance of the event manager
	m, err := NewEventManager(&EventConfig{
		Server:             "tcp://127.0.0.1:1883",
		ItemInstance:       "TEST_APP_01",
		Qos:                2,
		InsecureSkipVerify: true,
		OnMsgReceived:      onMsgReceived,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	// connect and subscribe
	m.Connect()
	<-done
}

// a handler to process received messages
func onMsgReceived(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", msg.Topic(), msg.Payload())
}
