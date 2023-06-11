package velocity_api

const (
	VELOCITY_POINT_TYPE   = "BASE"
	VELOCITY_PARTNER_CODE = "canstar"
)

type VelocityPoint struct {
	Amount string `json:"amount" validate:"required"`
	Type   string `json:"type" validate:"required"`
}

type VelocityMember struct {
	MembershipID string     `json:"membershipId" validate:"required"`
	Name         MemberName `json:"name,omitempty"`
}

type MemberName struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

type VelocityPartner struct {
	Code      string `json:"code" validate:"required"`
	Reference string `json:"reference" validate:"required"`
}

type VelocityActivity struct {
	ReferenceDate  string `json:"referenceDate" validate:"required"`
	Classification string `json:"classification" validate:"required"`
	Description    string `json:"description" validate:"required"`
}

type VelocityEarnPointPayload struct {
	Data VelocityEarPointData `json:"data" validate:"required"`
}

type VelocityEarPointData struct {
	Member   VelocityMember   `json:"member" validate:"required"`
	Points   []VelocityPoint  `json:"points" validate:"required"`
	Partner  VelocityPartner  `json:"partner" validate:"required"`
	Activity VelocityActivity `json:"activity" validate:"required"`
}

type VelocityEarPointsParams struct {
	MembershipID         string `validate:"required" `
	PointsToAllocate     int    `validate:"required" `
	LeadUID              string `validate:"required" `
	PointsAllocationDate string `validate:"required" `
}
