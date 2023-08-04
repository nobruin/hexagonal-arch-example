package controllers

import "encoding/json"

func stringToJson(msg string) []byte {
	message := struct {
		Message string `json:"message"`
	}{
		msg,
	}
	r, err := json.Marshal(message)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
