package e2e

import (
	"encoding/json"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/stretchr/testify/suite"
)

type e2eTestSuiteSignup struct {
	suite.Suite
}

func TestSuiteSigup(t *testing.T) {
	suite.Run(t, &e2eTestSuiteSignup{})
}

func (s *e2eTestSuiteSignup) SetupSuite() {
	_, err := checkApiAvailable(s.T(), apiURL)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}
}

func (s *e2eTestSuiteSignup) TearDownSuite() {
}

func (s *e2eTestSuiteSignup) TestSignup() {
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

	if len(respToken.AccessToken) == 0 {
		s.T().Error(err)
		s.Fail("expected token in response")
		return
	}

}

func (s *e2eTestSuiteSignup) TestDuplicate() {
	signupRequest := UserSignUpRequest{
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, false, false, 10),
		Name:     gofakeit.Name(),
		Surname:  gofakeit.LastName(),
	}

	_, _, err := checkPostRequest("/api/v1/sign-up", 200, signupRequest)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	_, _, err = checkPostRequest("/api/v1/sign-up", 422, signupRequest)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}
}

func (s *e2eTestSuiteSignup) TestSignin() {
	signupRequest := UserSignUpRequest{
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, false, false, 10),
		Name:     gofakeit.Name(),
		Surname:  gofakeit.LastName(),
	}

	_, _, err := checkPostRequest("/api/v1/sign-up", 200, signupRequest)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	signinRequest := UserSignInRequest{
		Email:    signupRequest.Email,
		Password: signupRequest.Password,
	}

	body, _, err := checkPostRequest("/api/v1/sign-in", 200, signinRequest)
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

	if len(respToken.AccessToken) == 0 {
		s.T().Error(err)
		s.Fail("expected token in response")
		return
	}
}

func (s *e2eTestSuiteSignup) TestWrongPwd() {
	signupRequest := UserSignUpRequest{
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, false, false, 10),
		Name:     gofakeit.Name(),
		Surname:  gofakeit.LastName(),
	}

	_, _, err := checkPostRequest("/api/v1/sign-up", 200, signupRequest)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	signinRequest := UserSignInRequest{
		Email:    signupRequest.Email,
		Password: "wrong password",
	}

	_, _, err = checkPostRequest("/api/v1/sign-in", 422, signinRequest)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}
}

func (s *e2eTestSuiteSignup) TestNonExistingUser() {
	signinRequest := UserSignInRequest{
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, false, false, 10),
	}

	_, _, err := checkPostRequest("/api/v1/sign-in", 422, signinRequest)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}
}
