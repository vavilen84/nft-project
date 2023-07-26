package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/store"
	"log"
	"net/http"
	"testing"
)

func Test_ForgotPassword(t *testing.T) {
	ts := initApp()
	defer ts.Close()
	db := store.GetDB()
	registerResp, email, _ := registerUser(t, ts)

	registeredUser := checkToken(t, db, registerResp.Data.Token)
	assert.Equal(t, registeredUser.Email, email)

	body := dto.ForgotPassword{
		Email: email,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/security/forgot-password", bytes.NewReader(bodyBytes))
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
