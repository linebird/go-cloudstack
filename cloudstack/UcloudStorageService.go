package cloudstack

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type CreateStorageParams struct {
	p map[string]interface{}
}

func (*UcloudStorageService) GetToken(url string, username string, password string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Storage-User", username)
	req.Header.Set("X-Storage-Pass", password)
	req.Header.Set("Accept", "*/*")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	token := res.Header.Get("X-Storage-Token")

	return token, nil
}

func (*UcloudStorageService) GetStorageObject(url string, token string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Auth-Token", token)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(string(data))
	}

	return data, nil
}

func (*UcloudStorageService) PutStorageObject(url string, token string, contentType string, filePath string) (bool, error) {
	var r io.Reader
	r, _ = os.Open(filePath)
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, url, r)
	req.Header.Set("X-Auth-Token", token)
	res, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	if res.StatusCode != 201 {
		return false, errors.New("put request is fail!")
	}

	return true, nil
}

func (*UcloudStorageService) DeleteStorageObject(url string, token string) (bool, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("X-Auth-Token", token)
	res, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	if res.StatusCode != 204 {
		return false, errors.New("delete request is fail!")
	}

	log.Printf("HTTP Status Code : %d", res.StatusCode)

	return true, err
}
