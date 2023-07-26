package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/auth"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/handlers"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
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

type TestLoginRespDataToken struct {
	Token string `json:"token"`
}

type TestLoginResp struct {
	Status     int                    `json:"status"`
	Data       TestLoginRespDataToken `json:"data"`
	Error      string                 `json:"error"`
	Errors     map[string][]string    `json:"errors"`
	FormErrors map[string][]string    `json:"formErrors"`
}

func initApp() *httptest.Server {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Can not load .env file")
	}
	store.InitDB()
	db := store.GetDB()
	if err := db.Exec("delete from jwt_info").Error; err != nil {
		fmt.Println("Error deleting entities:", err)
	}
	if err := db.Exec("delete from user").Error; err != nil {
		fmt.Println("Error deleting entities:", err)
	}
	handler := handlers.MakeHandler()
	ts := httptest.NewServer(handler)
	return ts
}

func registerUser(t *testing.T, ts *httptest.Server) (*TestRegisterResp, string, string) {
	email := "vladimir.teplov@gmail.com"
	password := "12345678"
	body := dto.SignUp{
		Nickname:    "test_" + helpers.GenerateRandomString(5),
		Email:       email,
		Password:    password,
		BillingPlan: constants.FreeBillingPlan,
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
		return nil, "", ""
	}

	registerResp := TestRegisterResp{}
	err = json.Unmarshal(responseBody, &registerResp)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, "", ""
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}

	assert.Equal(t, registerResp.Status, http.StatusOK)
	assert.NotEmpty(t, registerResp.Data.Token)
	assert.Empty(t, registerResp.Error)
	assert.Empty(t, registerResp.Error)
	assert.Empty(t, registerResp.Errors)
	assert.Empty(t, registerResp.FormErrors)

	return &registerResp, email, password
}

func loginUser(t *testing.T, ts *httptest.Server, email, password string) string {

	registerResp, email, password := registerUser(t, ts)

	body := dto.Login{
		Email:    email,
		Password: password,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodGet, ts.URL+"/api/v1/security/login", bytes.NewReader(bodyBytes))
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
		log.Fatal(err)
	}

	loginResp := TestLoginResp{}
	err = json.Unmarshal(responseBody, &loginResp)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}

	assert.Equal(t, registerResp.Status, http.StatusOK)
	assert.NotEmpty(t, registerResp.Data.Token)
	assert.Empty(t, registerResp.Error)
	assert.Empty(t, registerResp.Error)
	assert.Empty(t, registerResp.Errors)
	assert.Empty(t, registerResp.FormErrors)

	return registerResp.Data.Token
}

func checkToken(t *testing.T, db *gorm.DB, token string) *models.User {
	isValid, err := auth.VerifyJWT(db, []byte(token))
	if err != nil || token == "" || !isValid {
		log.Fatalln(err)
	}

	jwtPayload, err := auth.ParseJWTPayload([]byte(token))
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

	return userByJWTInfo
}
