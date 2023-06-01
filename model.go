package shorturl

type NewShort struct {
	Ori string `json:"ori" form:"ori" validate:"required"`
	Dec string `json:"dec" form:"dec" validate:"required"`
}
