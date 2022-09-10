package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/ronytampubolon/miniwallet/models"
	"github.com/ronytampubolon/miniwallet/schemas"
	services "github.com/ronytampubolon/miniwallet/services/remote"

	"github.com/ronytampubolon/miniwallet/utils"
)

type handler struct {
	service services.WalletService
}

var validate *validator.Validate

func NewWalletHandler(service services.WalletService) *handler {
	return &handler{service: service}
}

func (h *handler) InitWalletHandler(ctx *gin.Context) {
	var input schemas.InitInput
	ctx.Bind(&input)
	validate = validator.New()
	err := validate.Struct(input)
	if err != nil {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, utils.ParseError(err))
		return
	}
	account, _ := h.service.InitWallet(&input)
	accessTokenData := map[string]interface{}{"id": account.OwnedBy}
	accessToken, _ := utils.Sign(accessTokenData, "JWT_SECRET", 24*60*1)

	utils.APIResponse(ctx, http.StatusOK, "success", map[string]interface{}{"token": accessToken})
}

// Handler Get Balance
func (h *handler) GetBalanceHandler(ctx *gin.Context) {
	token, _ := utils.VerifyTokenHeader(ctx, "JWT_SECRET")
	claims := utils.DecodeToken(token)
	account, err := h.service.BalanceWallet(claims.Claims.ID)
	switch err.Code {
	case 400:
		utils.ValidatorErrorResponse(ctx, err.Code, err.Type)
		return
	default:
		utils.APIResponse(ctx, http.StatusOK, "success", map[string]interface{}{
			"wallet": models.TransformEnableResponse(*account),
		})
	}

}

// Handler Enable Wallet
func (h *handler) EnableWalletHandler(ctx *gin.Context) {
	token, _ := utils.VerifyTokenHeader(ctx, "JWT_SECRET")
	claims := utils.DecodeToken(token)
	account, err := h.service.EnableWallet(claims.Claims.ID)
	switch err.Code {
	case 400:
		utils.ValidatorErrorResponse(ctx, err.Code, err.Type)
		return
	default:
		utils.APIResponse(ctx, http.StatusOK, "success", map[string]interface{}{
			"wallet": models.TransformEnableResponse(*account),
		})
	}
}

// Handler Disabled Wallet
func (h *handler) DisableWalletHandler(ctx *gin.Context) {
	var input schemas.InputDisabled
	ctx.Bind(&input)
	validate = validator.New()
	validationErr := validate.Struct(input)
	if validationErr != nil {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, utils.ParseError(validationErr))
		return
	}
	if input.IsDisabled == false {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, map[string]interface{}{
			"is_disabled": "should be true",
		})
		return
	}

	token, _ := utils.VerifyTokenHeader(ctx, "JWT_SECRET")
	claims := utils.DecodeToken(token)
	account, err := h.service.DisableWallet(claims.Claims.ID)
	switch err.Code {
	case 400:
		utils.ValidatorErrorResponse(ctx, err.Code, err.Type)
		return
	default:
		utils.APIResponse(ctx, http.StatusOK, "success", map[string]interface{}{
			"wallet": models.TransformDisabledResponse(*account),
		})
	}
}

// Handler Deposit Wallter
func (h *handler) DepositHandler(ctx *gin.Context) {
	var input schemas.InputTransaction
	ctx.Bind(&input)
	validate = validator.New()
	validationErr := validate.Struct(input)
	if validationErr != nil {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, utils.ParseError(validationErr))
		return
	}

	token, _ := utils.VerifyTokenHeader(ctx, "JWT_SECRET")
	claims := utils.DecodeToken(token)
	trx, err := h.service.Deposit(claims.Claims.ID, &input)
	switch err.Code {
	case 400:
		utils.ValidatorErrorResponse(ctx, err.Code, err.Type)
		return
	default:
		utils.APIResponse(ctx, http.StatusOK, "success", map[string]interface{}{
			"deposit": trx,
		})
	}
}

// Handler Withdrawn Wallet
func (h *handler) WithdrawnHandler(ctx *gin.Context) {
	var input schemas.InputTransaction
	ctx.Bind(&input)
	validate = validator.New()
	validationErr := validate.Struct(input)
	if validationErr != nil {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, utils.ParseError(validationErr))
		return
	}

	token, _ := utils.VerifyTokenHeader(ctx, "JWT_SECRET")
	claims := utils.DecodeToken(token)
	trx, err := h.service.Withdrawn(claims.Claims.ID, &input)
	switch err.Code {
	case 400:
		utils.ValidatorErrorResponse(ctx, err.Code, err.Type)
		return
	default:
		utils.APIResponse(ctx, http.StatusOK, "success", map[string]interface{}{
			"withdrawal": trx,
		})
	}
}
