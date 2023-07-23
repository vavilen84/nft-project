package handlers

import (
	"encoding/json"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/helpers"
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
	resp = dto.Response{
		Error:  err.Error(),
		Status: status,
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
