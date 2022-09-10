package schemas

type InputTransaction struct {
	Amount      float64 `form:"amount" json:"amount" validate:"required"`
	ReferenceID string  `form:"reference_id" json:"reference_id" validate:"required"`
}
