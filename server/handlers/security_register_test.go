package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/auth"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"io"
	"log"
	"net/http"
	"testing"
)

type TestRegisterRespDataToken struct {
	Token string `json:"token"`
}

type TestRegisterResp struct {
	Status     int                       `json:"status"`
	Data       TestRegisterRespDataToken `json:"data"`
	Error      string                    `json:"error"`
	Errors     map[string][]string       `json:"errors"`
	FormErrors map[string][]string       `json:"formErrors"`
}

func TestRegister_OK(t *testing.T) {

	ts := initApp()
	defer ts.Close()
	db := store.GetDB()
	registerResp, email, _ := registerUser(t, ts)

	isValid, err := auth.VerifyJWT(db, []byte(registerResp.Data.Token))
	if err != nil || registerResp.Data.Token == "" || !isValid {
		log.Fatalln(err)
	}

	jwtPayload, err := auth.ParseJWTPayload([]byte(registerResp.Data.Token))
	if err != nil {
		log.Fatalln(err)
	}
	assert.NotEmpty(t, jwtPayload.JWTInfoId)

	jwtInfo, err := models.FindJWTInfoById(db, jwtPayload.JWTInfoId)
	if err != nil {
		log.Fatalln(err)
	}

	userByJWTInfo, err := models.FindUserById(db, jwtInfo.UserId)
	if err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, userByJWTInfo.Email, email)
}

func TestRegister_NotOK(t *testing.T) {

	ts := initApp()
	defer ts.Close()

	body := dto.SignUp{
		Nickname:    "",
		Email:       "vladimir.teplovgmail.com",
		Password:    "1234567",
		BillingPlan: 10,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/security/register", bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	registerResp := TestRegisterResp{}
	err = json.Unmarshal(responseBody, &registerResp)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res.StatusCode)
	}

	assert.Equal(t, registerResp.Status, http.StatusBadRequest)
	assert.Empty(t, registerResp.Data.Token)
	assert.Empty(t, registerResp.Error)
	assert.Empty(t, registerResp.Error)
	assert.Empty(t, registerResp.Errors)
	assert.NotEmpty(t, registerResp.FormErrors)
	assert.NotEmpty(t, registerResp.FormErrors["BillingPlan"][0])
	assert.NotEmpty(t, registerResp.FormErrors["Email"][0])
	assert.NotEmpty(t, registerResp.FormErrors["Nickname"][0])
	assert.NotEmpty(t, registerResp.FormErrors["Password"][0])
}
