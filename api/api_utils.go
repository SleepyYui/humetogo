package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h *HumeApi) printDebug(message string) {
	if h.Debug {
		fmt.Println("[DEBUG] " + message)
	}
}

func (h *HumeApi) printError(message string) {
	fmt.Println("[ERROR] " + message)
}

func (h *HumeApi) get(url string, target interface{}) error {
	h.printDebug("GET " + url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		h.printError(err.Error())
		return err
	}
	req.Header.Add("X-Hume-Api-Key", h.ApiKey)
	req.Header.Add("X-Hume-Api-Secret", h.ApiSecret)

	resp, err := RestClient.Do(req)
	if err != nil {
		h.printError(err.Error())
		return err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			h.printError(err.Error())
		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		h.printError(err.Error())
		return err
	}

	return nil
}

func (h *HumeApi) post(url string, body interface{}, target interface{}) error {
	h.printDebug("POST " + url)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		h.printError(err.Error())
		return err
	}
	req.Header.Add("X-Hume-Api-Key", h.ApiKey)
	req.Header.Add("X-Hume-Api-Secret", h.ApiSecret)

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		h.printError(err.Error())
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	resp, err := RestClient.Do(req)
	if err != nil {
		h.printError(err.Error())
		return err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			h.printError(err.Error())
		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		h.printError(err.Error())
		return err
	}

	return nil
}
