package velocity_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brewinski/unnamed-fiber/api/custom_http_client"
	"github.com/brewinski/unnamed-fiber/pkg/config"
)

type PointsEarnResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Partner  PartnerData  `json:"partner"`
	Velocity VelocityData `json:"velocity"`
}

type PartnerData struct {
	Reference string `json:"reference"`
}

type VelocityData struct {
	Reference string `json:"reference"`
}

func AllocatePoint(requestParams VelocityEarPointsParams) (PointsEarnResponse, error) {
	var PointsEarnResponse PointsEarnResponse
	apiURL := config.Config("VELOCITY_POINTS_EARN_ENDPOINT")

	header, err := generateRequestHeader()
	if err != nil {
		return PointsEarnResponse, err
	}
	requestPayload, err := generateRequestPayload(requestParams)
	if err != nil {
		return PointsEarnResponse, err
	}

	client := &custom_http_client.CustomHttpClient{Client: http.DefaultClient}
	response, err := client.MakeRequest(http.MethodPost, apiURL, bytes.NewReader(requestPayload), header)
	if err != nil {
		return PointsEarnResponse, err
	}

	err = json.Unmarshal(response, &PointsEarnResponse)
	if err != nil {
		return PointsEarnResponse, err
	}

	return PointsEarnResponse, nil
}

// Generates HTTP header containing the necessary information for making Points Earn API requests.
// It includes the X-APIKey, Content-Type, and Authorization headers.
// The Authorization header includes a bearer token obtained from the GetToken function.
// If the token retrieval fails, an empty header and the corresponding error are returned.
func generateRequestHeader() (http.Header, error) {
	tokenResponse, err := GetToken()

	if err != nil {
		return http.Header{}, err
	}

	header := http.Header{
		"X-APIKey":      {config.Config("VELOCITY_API_KEY")},
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer" + " " + tokenResponse.AccessToken},
	}
	return header, nil
}

// generates the request body payload for  points points api based on the provided parameters.
// It retrieves the activity classification for the given points allocation and constructs the payload accordingly.
// If the activity classification retrieval fails, an empty payload and the corresponding error are returned.
func generateRequestPayload(params VelocityEarPointsParams) ([]byte, error) {

	activityData, err := getActivityClassificationForPoints(params.PointsToAllocate)
	if err != nil {
		return nil, err
	}

	var payload = VelocityEarnPointPayload{
		Data: VelocityEarPointData{
			Member: VelocityMember{
				MembershipID: params.MembershipID,
				// uncomment if we are validating the firstName and lastName as well, will also need to include them in the params
				// Name: MemberName{
				// 	FirstName: "test",
				// 	LastName:  "test",
				// },
			},
			Points: []VelocityPoint{
				{
					Amount: fmt.Sprint(params.PointsToAllocate),
					Type:   VELOCITY_POINT_TYPE,
				}},
			Partner: VelocityPartner{
				Code:      VELOCITY_PARTNER_CODE,
				Reference: params.LeadUID,
			},
			Activity: VelocityActivity{
				ReferenceDate:  params.PointsAllocationDate,
				Classification: activityData.Code,
				Description:    activityData.Description,
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return jsonPayload, nil
}
