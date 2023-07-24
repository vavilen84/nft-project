package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/handlers"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/store"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initApp() *httptest.Server {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Can not load .env file")
	}
	store.InitDB()
	handler := handlers.MakeHandler()
	ts := httptest.NewServer(handler)
	return ts
}

// TODO: make fixtures load/clear before each integration test
// for now skipped and only for debug purposes
func TestRegister(t *testing.T) {
	//t.Skip()

	ts := initApp()
	defer ts.Close()

	body := dto.SignUp{
		Nickname: "test_" + helpers.GenerateRandomString(5),
		Email:    "vladimir.teplov@gmail.com",
		//Email:       "user_" + helpers.GenerateRandomString(5) + "example.com",
		Password:    "12345678",
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

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}
}
