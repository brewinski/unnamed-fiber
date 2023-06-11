package velocity_api

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type VelocityClassificationDetail struct {
	Code                 string `validate:"required"`
	Description          string `validate:"required"`
	MaximumPointsAllowed int    `validate:"gte=0"`
}

type VelocityActivityClassifications struct {
	Pending   VelocityClassificationDetail `validate:"required"`
	TierOne   VelocityClassificationDetail `validate:"required"`
	TierTwo   VelocityClassificationDetail `validate:"required"`
	TierThree VelocityClassificationDetail `validate:"required"`
}

func getVelocityActivityClassifications() VelocityActivityClassifications {
	classifications := VelocityActivityClassifications{
		Pending: VelocityClassificationDetail{
			Code:                 "CANHLSIGN",
			Description:          "Canstar Home Loan Base Points",
			MaximumPointsAllowed: 0,
		},
		TierOne: VelocityClassificationDetail{
			Code:                 "CANHLJOIN1",
			Description:          "Canstar Home Loan Base Points",
			MaximumPointsAllowed: 300000,
		},
		TierTwo: VelocityClassificationDetail{
			Code:                 "CANHLJOIN",
			Description:          "Velocity - Canstar Home Loan Join",
			MaximumPointsAllowed: 500000,
		},
		TierThree: VelocityClassificationDetail{
			Code:                 "CANHLJOIN3",
			Description:          "Canstar Base Home Loan Points",
			MaximumPointsAllowed: 1000000,
		},
	}
	validate := validator.New()
	err := validate.Struct(classifications)
	if err != nil {
		fmt.Println(err)
	}

	return classifications
}

func getActivityClassificationForPoints(pointsToAllocate int) (VelocityClassificationDetail, error) {

	var result VelocityClassificationDetail

	velocityActivityClassifications := getVelocityActivityClassifications()
	pendingMax := velocityActivityClassifications.Pending.MaximumPointsAllowed
	tierOneMax := velocityActivityClassifications.TierOne.MaximumPointsAllowed
	tierTwoMax := velocityActivityClassifications.TierTwo.MaximumPointsAllowed

	switch {
	case pointsToAllocate == pendingMax:
		result = velocityActivityClassifications.Pending

	case pointsToAllocate > pendingMax && pointsToAllocate <= tierOneMax:
		result = velocityActivityClassifications.TierOne

	case pointsToAllocate > tierOneMax && pointsToAllocate < tierTwoMax:
		result = velocityActivityClassifications.TierTwo

	// currently we don't allocate tier 3 points
	default:
		err := errors.New("invalid number of points to allocate, points to allocate must be between 0 and 500")
		return result, err
	}
	return result, nil

}
