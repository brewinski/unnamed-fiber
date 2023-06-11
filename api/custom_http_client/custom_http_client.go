package custom_http_client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CustomHttpClient struct {
	*http.Client
}

// Makes an HTTP request with the specified method, URL, body, and headers.
// It returns the response body as a byte array and an HTTPError if an error occurs.
func (c *CustomHttpClient) MakeRequest(method, url string, body io.Reader, headers http.Header) ([]byte, error) {
	// Create a new HTTP request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Set the request headers
	req.Header = headers

	// Send the request
	response, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Check if the response is an error
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		errorCode := response.StatusCode
		errorMessage := fmt.Sprintf(`{"errorMessage": "%s", "errorDetail": %s}`, response.Status, responseBytes)
		fmt.Println("Request failed while calling: " + url)
		fmt.Println(errorMessage)
		return nil, fiber.NewError(errorCode, errorMessage)
	}

	// Return the response body
	return responseBytes, nil
}
