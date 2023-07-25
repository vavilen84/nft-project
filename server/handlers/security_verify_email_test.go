package handlers_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/auth"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"log"
	"net/http"
	"testing"
)

func TestVerifyEmail_OK(t *testing.T) {

	ts := initApp()
	defer ts.Close()

	registerResp, _, _ := registerUser(t, ts)
	jwtPayload, err := auth.ParseJWTPayload([]byte(registerResp.Data.Token))
	if err != nil {
		log.Fatal(err)
	}
	db := store.GetDB()
	jwtInfo, err := models.FindJWTInfoById(db, jwtPayload.JWTInfoId)
	if err != nil {
		log.Fatal(err)
	}
	u, err := models.FindUserById(db, jwtInfo.UserId)
	if err != nil {
		log.Fatal(err)
	}
	assert.NotEmpty(t, u.EmailTwoFaCode)
	assert.False(t, u.IsEmailVerified)

	foundUser, err := models.FindUserByTwoFAToken(db, u.EmailTwoFaCode)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, foundUser.Id, u.Id)

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(ts.URL+"/api/v1/security/verify-email?token=%s", u.EmailTwoFaCode),
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	u, err = models.FindUserById(db, jwtInfo.UserId)
	if err != nil {
		log.Fatal(err)
	}
	assert.Empty(t, u.EmailTwoFaCode)
	assert.True(t, u.IsEmailVerified)
}

func TestVerifyEmail_NotOK(t *testing.T) {

	ts := initApp()
	defer ts.Close()

	registerResp, _, _ := registerUser(t, ts)
	jwtPayload, err := auth.ParseJWTPayload([]byte(registerResp.Data.Token))
	if err != nil {
		log.Fatal(err)
	}
	db := store.GetDB()
	jwtInfo, err := models.FindJWTInfoById(db, jwtPayload.JWTInfoId)
	if err != nil {
		log.Fatal(err)
	}
	u, err := models.FindUserById(db, jwtInfo.UserId)
	if err != nil {
		log.Fatal(err)
	}
	assert.NotEmpty(t, u.EmailTwoFaCode)
	assert.False(t, u.IsEmailVerified)

	foundUser, err := models.FindUserByTwoFAToken(db, u.EmailTwoFaCode)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, foundUser.Id, u.Id)

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(ts.URL+"/api/v1/security/verify-email?token=%s1", u.EmailTwoFaCode),
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	u, err = models.FindUserById(db, jwtInfo.UserId)
	assert.NotEmpty(t, u.EmailTwoFaCode)
	assert.False(t, u.IsEmailVerified)
}
