package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/vavilen84/nft-project/auth"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"github.com/vavilen84/nft-project/validation"
	"gorm.io/gorm"
	"net/http"
)

type SecurityController struct {
	BaseController
}

func (c *SecurityController) Register(w http.ResponseWriter, r *http.Request) {
	db := store.GetDB()
	dec := json.NewDecoder(r.Body)
	dtoModel := dto.SignUp{}
	err := dec.Decode(&dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	err = validation.ValidateByScenario(constants.ScenarioSignIn, &dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusNotFound)
		return
	}
	u, err := models.FindUserByEmail(db, dtoModel.Email)
	if err != nil {
		helpers.LogError(err)
		if err != gorm.ErrRecordNotFound {
			http.Error(w, "Bad Request", http.StatusInternalServerError)
			return
		}
	} else {
		err := errors.New(fmt.Sprintf("user with email %s already exists", dtoModel.Email))
		helpers.LogError(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u = &models.User{
		Email:    dtoModel.Email,
		Password: dtoModel.Password,
		Nickname: dtoModel.Nickname,
	}
	err = models.InsertUser(db, u)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	tok, err := auth.CreateJWT(db, u)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	resp := make(dto.ResponseData)
	resp["token"] = string(tok)
	c.WriteSuccessResponse(w, resp, http.StatusOK)
}

func (c *SecurityController) Login(w http.ResponseWriter, r *http.Request) {
	db := store.GetDB()
	dec := json.NewDecoder(r.Body)
	model := dto.Login{}
	err := dec.Decode(&model)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	err = validation.ValidateByScenario(constants.ScenarioSignIn, &model)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	u, err := models.FindUserByEmail(db, model.Email)
	if err != nil {
		helpers.LogError(err)
		if err == gorm.ErrRecordNotFound {
			err = errors.New(fmt.Sprintf("user with email %s not found", model.Email))
			c.WriteErrorResponse(w, err, http.StatusNotFound)
		} else {
			c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		}
		return
	}
	passwordIsValid := password.Verify(model.Password, u.PasswordSalt, u.Password, nil)
	if !passwordIsValid {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.UnauthorizedError, http.StatusUnauthorized)
		return
	}
	tok, err := auth.CreateJWT(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}
	resp := make(dto.ResponseData)
	resp["token"] = string(tok)
	c.WriteSuccessResponse(w, resp, http.StatusOK)
}
