package handlers

import (
	"encoding/json"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/validation"
	"net/http"
)

type BaseController struct{}

func (*BaseController) WriteSuccessResponse(w http.ResponseWriter, data interface{}, status int) {
	resp := dto.Response{
		Status: status,
		Data:   data,
	}
	writeResponse(w, resp, status)
}

func (*BaseController) WriteErrorResponse(w http.ResponseWriter, err error, status int) {
	resp := dto.Response{}
	helpers.LogError(err)
	errs, ok := err.(validation.Errors)
	if ok {
		formErrors := make(map[string][]string)
		for field, fieldErrors := range errs {
			fieldErrMsgs := make([]string, 0)
			for _, v := range fieldErrors {
				fieldErrMsgs = append(fieldErrMsgs, v.Message)
			}
			formErrors[field] = fieldErrMsgs
		}
		resp = dto.Response{
			FormErrors: formErrors,
			Status:     status,
		}
	} else {
		resp = dto.Response{
			Error:  err.Error(),
			Status: status,
		}
	}
	writeResponse(w, resp, status)
}

func writeResponse(w http.ResponseWriter, resp dto.Response, status int) {
	b, e := json.Marshal(resp)
	if e != nil {
		helpers.LogError(e)
		return
	}
	_, e = w.Write(b)
	if e != nil {
		helpers.LogError(e)
		return
	}
	w.WriteHeader(status)
}
