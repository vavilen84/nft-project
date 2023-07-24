package handlers_test

//func TestVerifyEmail_OK(t *testing.T) {
//
//	ts := initApp()
//	defer ts.Close()
//
//	registerResp := registerUser(t, ts)
//
//	body := dto.Ver{
//		Nickname:    "test_" + helpers.GenerateRandomString(5),
//		Email:       "vladimir.teplov@gmail.com",
//		Password:    "12345678",
//		BillingPlan: constants.FreeBillingPlan,
//	}
//	bodyBytes, err := json.Marshal(body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	req, err := http.NewRequest(http.MethodGet, ts.URL+"/api/v1/security/register", bytes.NewReader(bodyBytes))
//	if err != nil {
//		t.Fatalf("Failed to create request: %v", err)
//	}
//
//	res, err := http.DefaultClient.Do(req)
//	if err != nil {
//		t.Fatalf("Failed to send request: %v", err)
//	}
//	defer res.Body.Close()
//
//	responseBody, err := io.ReadAll(res.Body)
//	if err != nil {
//		fmt.Println("Error reading response body:", err)
//		return
//	}
//
//	registerResp := TestRegisterResp{}
//	err = json.Unmarshal(responseBody, &registerResp)
//	if err != nil {
//		fmt.Println("Error reading response body:", err)
//		return
//	}
//
//	if res.StatusCode != http.StatusOK {
//		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
//	}
//
//	assert.Equal(t, registerResp.Status, http.StatusOK)
//	assert.NotEmpty(t, registerResp.Data.Token)
//	assert.Empty(t, registerResp.Error)
//	assert.Empty(t, registerResp.Error)
//	assert.Empty(t, registerResp.Errors)
//	assert.Empty(t, registerResp.FormErrors)
//}
