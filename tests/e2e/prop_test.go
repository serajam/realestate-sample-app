package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"
)

type e2eTestSuiteProp struct {
	suite.Suite
}

func TestSuiteProp(t *testing.T) {
	suite.Run(t, &e2eTestSuiteProp{})
}

func (s *e2eTestSuiteProp) SetupSuite() {
	_, err := checkApiAvailable(s.T(), apiURL)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}
}

func (s *e2eTestSuiteProp) TearDownSuite() {
}

func (s *e2eTestSuiteProp) TestCreate() {
	signupRequest := UserSignUpRequest{
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, false, false, 10),
		Name:     gofakeit.Name(),
		Surname:  gofakeit.LastName(),
	}

	body, _, err := checkPostRequest("/api/v1/sign-up", 200, signupRequest)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	var respToken authToken
	err = json.Unmarshal(body, &respToken)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	property := NewTestPropertyRequest()
	propBody, err := json.Marshal(property)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, apiURL+"/api/v1/user/properties", bytes.NewBuffer(propBody))
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+respToken.AccessToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		s.T().Errorf("expected status code 200, got %d", res.StatusCode)
		respErr := ErrorResponse{}
		err := json.NewDecoder(res.Body).Decode(&respErr)
		if err != nil {
			s.T().Error(err)
			return
		}
		s.T().Error(respErr.Message)
		s.T().Error(respErr.FieldErrors)
		s.T().Error(respErr.Error)
		return
	}

	data, err := io.ReadAll(res.Body)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	var respProperty PropertyResponse
	err = unmarshalTo(data, &respProperty)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	property = NewTestPropertyRequest()
	property.ID = respProperty.ID
	propBody, err = json.Marshal(property)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	req, err = http.NewRequest(
		http.MethodPut, apiURL+fmt.Sprintf("/api/v1/user/properties/%d", respProperty.ID), bytes.NewBuffer(propBody),
	)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+respToken.AccessToken)
	req.Header.Add("Content-Type", "application/json")

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		s.T().Errorf("expected status code 200, got %d", res.StatusCode)
		respErr := ErrorResponse{}
		err := json.NewDecoder(res.Body).Decode(&respErr)
		if err != nil {
			s.T().Error(err)
			return
		}
		s.T().Error(respErr.Message)
		s.T().Error(respErr.Error)
		return
	}

	data, err = io.ReadAll(res.Body)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	err = unmarshalTo(data, &respProperty)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}
}
