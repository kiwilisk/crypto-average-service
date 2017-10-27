package main

import (
	"time"
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

type RestClient interface {
	GetResponse(url string) (*HttpResponse, error)
}

type HttpResponse struct {
	StatusCode int
	Body       []byte
	Size       int64
	ReceivedAt time.Time
}

type DefaultRestClient struct {
	client http.Client
}

type RetryingRestClient struct {
	client       http.Client
	retryTimeout time.Duration
}

func (c DefaultRestClient) Get(url string) (*HttpResponse, error) {
	response, err := c.client.Get(url)
	if err != nil {
		err = fmt.Errorf("failed to get response from url %s, %v", url, err)
		return nil, err
	}
	if response.StatusCode >= 300 {
		err = fmt.Errorf("failed to get successful response from url %s. Status was: %s, %v", url, response.Status, err)
		return &HttpResponse{StatusCode: response.StatusCode, ReceivedAt: time.Now()}, err
	}
	return parseBody(response)
}

func (c RetryingRestClient) Get(url string) (*HttpResponse, error) {
	deadline := time.Now().Add(c.retryTimeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		response, err := c.client.Get(url)
		if err != nil {
			log.Printf("failed to get response from url %s, %v. Retrying", url, err)
			time.Sleep(time.Second << uint(tries)) // exponential backoff
			continue
		}
		if response.StatusCode >= 300 {
			err = fmt.Errorf("failed to get successful response from url %s. Status was %d - %s, %v", url, response.StatusCode, response.Status, err)
			return nil, err
		}
		return parseBody(response)
	}
	return nil, fmt.Errorf("failed to contact server with url %s within backoff period", url)
}

func parseBody(response *http.Response) (*HttpResponse, error) {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("failed to read body, %v", err)
		return nil, err
	}
	return &HttpResponse{
		response.StatusCode,
		body,
		response.ContentLength,
		time.Now(),
	}, nil
}
