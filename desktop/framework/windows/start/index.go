/**
 ******************************************************************************
 * @file    index.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package StartWindows

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Api struct {
	ctx context.Context
}

type ReturnResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Init() *Api {
	return &Api{}
}

func (start *Api) Startup(ctx context.Context) {
	start.ctx = ctx
}

func (start *Api) Shutdown(ctx context.Context) {

}

func (start *Api) DeviceRequest(host string, path string, method string, parameter any) map[string]interface{} {

	response := map[string]interface{}{"code": 0}

	requestBody, err := json.Marshal(parameter)
	if err != nil {
		response["code"] = 10000
		return response
	}

	urlWithParams := "http://" + host + path
	if method == "GET" {
		queryParams := url.Values{}
		parameters := map[string]string{}
		json.Unmarshal(requestBody, &parameters)
		for key, value := range parameters {
			queryParams.Add(key, value)
		}
		urlWithParams += "?" + queryParams.Encode()
	}

	request, err := http.NewRequest(method, urlWithParams, bytes.NewReader(requestBody))
	if err != nil {
		response["code"] = 10000
		return response
	}

	responseBody, err := start.onRequest(request, "application/json", "")
	if err != nil {
		response["code"] = 10000
		return response
	}
	defer responseBody.Close()

	if err := json.NewDecoder(responseBody).Decode(&response); err != nil {
		response["code"] = 10000
		return response
	}
	response["code"] = 0
	return response
}

func (start *Api) ServiceRequest(path string, method string, parameter any, token string) ReturnResponse {

	response := ReturnResponse{}

	requestBody, err := json.Marshal(parameter)
	if err != nil {
		response.Code = 10000
		return response
	}

	urlWithParams := "https://gateway.geekros.com" + path
	if method == "GET" {
		queryParams := url.Values{}
		parameters := map[string]string{}
		json.Unmarshal(requestBody, &parameters)
		for key, value := range parameters {
			queryParams.Add(key, value)
		}
		urlWithParams += "?" + queryParams.Encode()
	}

	request, err := http.NewRequest(method, urlWithParams, bytes.NewReader(requestBody))
	if err != nil {
		response.Code = 10000
		return response
	}

	responseBody, err := start.onRequest(request, "application/json", token)
	if err != nil {
		response.Code = 10000
		return response
	}
	defer responseBody.Close()

	if err := json.NewDecoder(responseBody).Decode(&response); err != nil {
		response.Code = 10000
		return response
	}

	return response
}

func (start *Api) onRequest(request *http.Request, contentType string, token string) (io.ReadCloser, error) {

	request.Header.Set("Content-Type", contentType)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 GEEKCNC/1.0.0")
	request.Header.Set("Account-Token", token)

	HttpClient := &http.Client{
		Timeout: 15 * time.Second,
	}

	response, err := HttpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error request: %w", err)
	}

	if response.StatusCode < 200 || response.StatusCode >= 400 {
		respBody, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("error request: %s", respBody)
	}

	return response.Body, nil
}
