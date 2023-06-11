package velocity_api

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/brewinski/unnamed-fiber/api/custom_http_client"
	"github.com/brewinski/unnamed-fiber/pkg/config"
)

type velocityTokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	Scope            string `json:"scope"`
}

func getEncodedCredentials() string {
	clientID := config.Config("VELOCITY_CLIENT_ID")
	clientSecret := config.Config("VELOCITY_CLIENT_SECRET")
	credentials := clientID + ":" + clientSecret
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))
	return encodedCredentials
}

func GetToken() (velocityTokenResponse, error) {
	tokenEndpoint := config.Config("VELOCITY_AUTH_ENDPOINT")
	encodedCredentials := getEncodedCredentials()
	headers := http.Header{
		"Authorization": {"Basic " + encodedCredentials},
		"Content-Type":  {"application/x-www-form-urlencoded"},
	}

	requestBody := url.Values{
		"grant_type": {"client_credentials"},
	}

	urlencodedBody := strings.NewReader(requestBody.Encode())
	client := &custom_http_client.CustomHttpClient{Client: http.DefaultClient}
	response, err := client.MakeRequest(http.MethodPost, tokenEndpoint, urlencodedBody, headers)
	if err != nil {
		return velocityTokenResponse{}, err
	}

	var decodedResponse = &velocityTokenResponse{}
	err = json.Unmarshal(response, decodedResponse)
	if err != nil {
		return velocityTokenResponse{}, err
	}
	return *decodedResponse, nil
}
