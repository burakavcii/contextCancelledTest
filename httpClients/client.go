package httpClients

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const URL = "http://google.com"

type HTTPContextClient struct {
	client http.Client
	url    *url.URL
}

func NewHTTPContextClient() HTTPContextClient {
	parsedUrl, _ := url.Parse(URL)
	return HTTPContextClient{client: http.Client{}, url: parsedUrl}
}
func NewHTTPContextClient3(URL string) HTTPContextClient {
	parsedUrl, _ := url.Parse(URL)
	return HTTPContextClient{client: http.Client{}, url: parsedUrl}
}

func (c *HTTPContextClient) CallCancelBefore() error {
	httpRequest := http.Request{URL: c.url, Method: "GET"}
	resp, err := c.doRequestWithTimeout(&httpRequest)
	if err != nil {
		log.Printf("error while doing request: %s\n", err)
		return err
	}
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			log.Printf("error while closing body: %s\n", closeErr)
		}
	}()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error while reading body: %s\n", err)
		return err
	}
	return nil
}

func (c *HTTPContextClient) CallCancelBefore2() error {
	httpRequest := http.Request{URL: c.url, Method: "GET"}
	resp, err := c.doRequestWithTimeout(&httpRequest)
	if err != nil {
		log.Printf("error while doing request: %s\n", err)
		return err
	}
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			log.Printf("error while closing body: %s\n", closeErr)
		}
	}()
	time.Sleep(time.Millisecond * 1)
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error while reading body: %s\n", err)
		return err
	}
	return nil
}

func (c *HTTPContextClient) CallCancelBefore3() error {
	httpRequest := http.Request{URL: c.url, Method: "GET"}
	resp, cancelFunc, err := c.doRequestWithTimeout3(&httpRequest)
	defer cancelFunc()
	if err != nil {
		log.Printf("error while doing request: %s\n", err)
		return err
	}
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			log.Printf("error while closing body: %s\n", closeErr)
		}
	}()
	time.Sleep(time.Millisecond * 1)
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error while reading body: %s\n", err)
		return err
	}
	return nil
}

func (c *HTTPContextClient) doRequestWithTimeout(request *http.Request) (*http.Response, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*3000)
	*request = *request.WithContext(ctx)
	defer cancelFunc()

	resp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *HTTPContextClient) doRequestWithTimeout3(request *http.Request) (*http.Response, context.CancelFunc, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*3000)
	*request = *request.WithContext(ctx)

	resp, err := c.client.Do(request)
	if err != nil {
		return nil, cancelFunc, err
	}

	return resp, cancelFunc, nil
}

func (c *HTTPContextClient) CallCancelAfter() error {
	httpRequest := http.Request{URL: c.url, Method: "GET"}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*3000)
	httpRequest = *httpRequest.WithContext(ctx)
	defer cancelFunc()

	resp, err := c.client.Do(&httpRequest)
	if err != nil {
		log.Printf("error while doing request: %s\n", err)
		return err
	}

	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			log.Printf("error while closing body: %s\n", closeErr)
		}
	}()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error while reading body: %s\n", err)
		return err
	}
	return nil
}
