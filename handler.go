package main

import (
	"encoding/json"
	"net/http"
)

func ping(writer http.ResponseWriter, request *http.Request) {
	message := Message{
		Body: "pong",
	}
	data, _ := json.Marshal(message)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(data)
}

func newMessage(writer http.ResponseWriter, request *http.Request) {
	var message Message

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&message)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(err.Error())
		_, _ = writer.Write(data)
	} else {
		data, _ := json.Marshal(message)
		// append new message to our data container
		messages = append(messages, message)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		_, _ = writer.Write(data)
	}
}

func allMessages(writer http.ResponseWriter, request *http.Request) {
	data, _ := json.Marshal(messages)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(data)
}
