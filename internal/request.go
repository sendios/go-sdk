package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Client *http.Client
	Auth   *Auth
}

func (r *Request) Post(url, route string, data interface{}) ([]byte, error) {
	postBody, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling data: %s", err)
	}

	body := bytes.NewBuffer(postBody)

	req, err := http.NewRequest(http.MethodPost, url+route, body)
	if err != nil {
		return nil, fmt.Errorf("error while creating request: %s", err)
	}

	key := Sha1Encoder(r.Auth.AuthKey)
	req.SetBasicAuth(r.Auth.ClientID, key)

	response, err := r.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %s", err)
	}

	defer response.Body.Close()

	var result interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("error while decoding response data: %s", err)
	}

	encoded, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling response data: %s", err)
	}

	return encoded, nil
}

func (r *Request) Get(url, route string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", url, route), nil)
	if err != nil {
		return nil, fmt.Errorf("error while creating request: %s", err)
	}

	key := Sha1Encoder(r.Auth.AuthKey)
	req.SetBasicAuth(r.Auth.ClientID, key)

	response, err := r.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %s", err)
	}

	defer response.Body.Close()

	var result interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("error while decoding response data: %s", err)
	}

	encoded, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling response data: %s", err)
	}

	return encoded, nil
}

func (r *Request) Delete(url, route string, data interface{}) ([]byte, error) {
	postBody, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling data: %s", err)
	}

	body := bytes.NewBuffer(postBody)

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", url, route), body)
	if err != nil {
		return nil, fmt.Errorf("error while creating request: %s", err)
	}

	key := Sha1Encoder(r.Auth.AuthKey)
	req.SetBasicAuth(r.Auth.ClientID, key)

	response, err := r.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %s", err)
	}

	defer response.Body.Close()

	var result interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("error while decoding response data: %s", err)
	}

	encoded, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling response data: %s", err)
	}

	return encoded, nil
}

func (r *Request) Put(url, route string, data interface{}) ([]byte, error) {
	putBody, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling data: %s", err)
	}

	body := bytes.NewBuffer(putBody)

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", url, route), body)
	if err != nil {
		return nil, fmt.Errorf("error while creating request: %s", err)
	}

	key := Sha1Encoder(r.Auth.AuthKey)
	req.SetBasicAuth(r.Auth.ClientID, key)

	response, err := r.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %s", err)
	}

	defer response.Body.Close()

	var result interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("error while decoding response data: %s", err)
	}

	encoded, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling response data: %s", err)
	}

	return encoded, nil
}
