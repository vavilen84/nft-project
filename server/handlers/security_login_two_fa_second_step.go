package handlers

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/auth"
	"github.com/vavilen84/nft-project/aws"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"gorm.io/gorm"
	"net/http"
)

func (c *SecurityController) TwoFaLoginStepTwo(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dtoModel := dto.TwoFaLoginStepTwo{}
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
	u, err := models.FindUserByTwoFAToken(db, dtoModel.EmailTwoFaCode)
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

	u.EmailTwoFaCode = ""
	err = models.SetEmailTwoFaCode(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}

	token := helpers.GenerateRandomString(6)
	u.EmailTwoFaCode = token

	err = models.SetEmailTwoFaCode(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}

	err = aws.SendLoginTwoFaCode(u.Email, token)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	tok, err := auth.CreateJWT(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	resp := make(dto.ResponseData)
	resp["token"] = string(tok)

	c.WriteSuccessResponse(w, resp, http.StatusOK)
}

func (c *SecurityController) TwoFaLoginStepOne(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dtoModel := dto.TwoFaLoginStepTwo{}
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
	u, err := models.FindUserByTwoFAToken(db, dtoModel.EmailTwoFaCode)
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

	u.EmailTwoFaCode = ""
	err = models.SetEmailTwoFaCode(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}

	token := helpers.GenerateRandomString(6)
	u.EmailTwoFaCode = token

	err = models.SetEmailTwoFaCode(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}

	err = aws.SendLoginTwoFaCode(u.Email, token)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	tok, err := auth.CreateJWT(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	resp := make(dto.ResponseData)
	resp["token"] = string(tok)

	c.WriteSuccessResponse(w, resp, http.StatusOK)
}
