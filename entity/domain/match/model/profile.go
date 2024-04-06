package model

// Gender is an enum that represents gender.
type Gender int

const (
	GenderUnspecified Gender = iota
	GenderMale
	GenderFemale
	GenderOther
)

var genderMap = map[Gender]string{
	GenderUnspecified: "unspecified",
	GenderMale:        "male",
	GenderFemale:      "female",
	GenderOther:       "other",
}

func (x *Gender) String() string {
	return genderMap[*x]
}

// Profile is a value object that represents a user profile.
type Profile struct {
	Name   string `json:"name,omitempty"`
	Age    uint   `json:"age,omitempty"`
	Gender Gender `json:"gender,omitempty"`
	Height uint   `json:"height,omitempty"`
}
