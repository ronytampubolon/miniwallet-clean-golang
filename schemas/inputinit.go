package schemas

type InitInput struct {
	CustomerID string `form:"customer_xid" json:"customer_xid" validate:"required"`
}
