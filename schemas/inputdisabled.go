package schemas

type InputDisabled struct {
	IsDisabled bool `form:"is_disabled" json:"is_disabled" validate:"required"`
}
