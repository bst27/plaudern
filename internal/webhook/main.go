package webhook

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Receiver interface {
	Receive(url string, payload map[string]interface{}) error
}

func New() Receiver {
	return &receiver{}
}

type receiver struct{}

func (receiver) Receive(url string, payload map[string]interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(data))
	return err
}
