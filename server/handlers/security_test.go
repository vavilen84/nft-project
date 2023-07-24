package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/helpers"
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

	body := dto.SignUp{
		Nickname:    "test_" + helpers.GenerateRandomString(5),
		Email:       "vladimir.teplov@gmail.com",
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

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}

	assert.Equal(t, registerResp.Status, http.StatusOK)
	assert.NotEmpty(t, registerResp.Data.Token)
	assert.Empty(t, registerResp.Error)
	assert.Empty(t, registerResp.Error)
	assert.Empty(t, registerResp.Errors)
	assert.Empty(t, registerResp.FormErrors)
}
