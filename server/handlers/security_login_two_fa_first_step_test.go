package handlers_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/store"
	"testing"
)

func TestLogin_OK(t *testing.T) {
	ts := initApp()
	defer ts.Close()
	db := store.GetDB()
	registerResp, email, password := registerUser(t, ts)

	registeredUser := checkToken(t, db, registerResp.Data.Token)
	assert.Equal(t, registeredUser.Email, email)

	loggedInUserToken := loginUser(t, ts, email, password)
	loggedInUser := checkToken(t, db, loggedInUserToken)
	assert.Equal(t, loggedInUser.Email, registeredUser.Email)
}
