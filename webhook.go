package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func PostMessage(url string, message string) error {
	body, err := json.Marshal(map[string]string{"text": message})
	if err != nil {
		return err
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	err = response.Body.Close()
	if err != nil {
		return err
	}

	return nil
}
