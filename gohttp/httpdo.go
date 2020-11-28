package gohttp

import "net/http"

func (c *httpClient) do(method string,
	url string,
	headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}
	request, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	fullHeaders := c.getRequestHeaders(headers)
	request.Header = fullHeaders
	defer client.CloseIdleConnections()
	return client.Do(request)
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)
	for header, value := range c.Headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	return result
}
