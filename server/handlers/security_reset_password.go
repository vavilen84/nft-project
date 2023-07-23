package handlers

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"gorm.io/gorm"
	"net/http"
)

func (c *SecurityController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dtoModel := dto.ResetPassword{}
	err := dec.Decode(&dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	db := store.GetDB()
	u, err := models.FindUserByResetPasswordToken(db, dtoModel.Token)
	if err != nil {
		helpers.LogError(err)
		if err == gorm.ErrRecordNotFound {
			err = errors.New("user not found")
			c.WriteErrorResponse(w, err, http.StatusNotFound)
		} else {
			c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		}
		return
	}
	u.Password = dtoModel.NewPassword
	err = models.UserResetPassword(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}
	resp := make(dto.ResponseData)
	c.WriteSuccessResponse(w, resp, http.StatusOK)
}
