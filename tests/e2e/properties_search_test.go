package e2e

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/suite"
)

type e2eTestSuiteSearch struct {
	suite.Suite
}

func TestSuiteSearch(t *testing.T) {
	suite.Run(t, &e2eTestSuiteSearch{})
}

func (s *e2eTestSuiteSearch) SetupSuite() {
	_, err := checkApiAvailable(s.T(), apiURL)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}
}

func (s *e2eTestSuiteSearch) TearDownSuite() {
}

func (s *e2eTestSuiteSearch) TestSearchByPolygon() {
	searchRequest := PropertySearchRequest{
		Polygon: &Polygon{
			TopLat:     "47.14668",
			TopLong:    "17.135",
			BottomLat:  "48.1434",
			BottomLong: "17.14422",
		},
		Size: 10,
	}

	body, _, err := checkPostRequest("/api/v1/properties/search", 200, searchRequest)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	var respProperties []PropertyResponse
	err = unmarshalTo(body, &respProperties)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	if len(respProperties) == 0 {
		s.Fail("expected at least 1 property in search result")
		return
	}
}

func (s *e2eTestSuiteSearch) TestSearchByOther() {
	searchRequest := PropertySearchRequest{
		CountryCode:   "sk",
		City:          "Bratislava",
		Sort:          1,
		PriceFrom:     1,
		PriceTo:       1000000,
		HomeSizeFrom:  50,
		HomeSizeTo:    1000,
		LotSizeFrom:   10,
		LotSizeTo:     10000,
		YearBuildFrom: 1911,
		YearBuildTo:   2009,
		Bathroom:      1,
		BathroomExact: lo.ToPtr(true),
		Bedroom:       1,
		BedroomExact:  lo.ToPtr(true),
		Condition:     []int{1, 2, 3},
		HomeType:      []int{1, 2, 3, 4, 5, 6, 7, 8},
	}

	body, _, err := checkPostRequest("/api/v1/properties/search", 200, searchRequest)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	var respProperties []PropertyResponse
	err = unmarshalTo(body, &respProperties)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	if len(respProperties) == 0 {
		s.Fail("expected at least 1 property in search result")
		return
	}
}

func (s *e2eTestSuiteSearch) TestSearchByAddress() {
	searchRequest := PropertySearchRequest{
		CountryCode:   "sk",
		City:          "Bratislava",
		Size:          10,
		YearBuildFrom: 2008,
		YearBuildTo:   2009,
	}

	body, _, err := checkPostRequest("/api/v1/properties/search", 200, searchRequest)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	var respProperties []PropertyResponse
	err = unmarshalTo(body, &respProperties)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	if len(respProperties) == 0 {
		s.Fail("expected at least 1 property in search result")
		return
	}
}

func (s *e2eTestSuiteSearch) TestSearchByIDs() {
	searchRequest := PropertySearchRequest{
		CountryCode:   "sk",
		City:          "Bratislava",
		Size:          10,
		YearBuildFrom: 2008,
		YearBuildTo:   2009,
	}

	body, _, err := checkPostRequest("/api/v1/properties/search", 200, searchRequest)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	var respProperties []PropertyResponse
	err = unmarshalTo(body, &respProperties)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	if len(respProperties) == 0 {
		s.Fail("expected at least 1 property in search result")
		return
	}

	var ids []int
	for _, p := range respProperties {
		ids = append(ids, p.ID)
	}

	searchRequest2 := PropertyListRequest{
		IDs: ids,
	}

	body, _, err = checkPostRequest("/api/v1/properties/list", 200, searchRequest2)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	err = unmarshalTo(body, &respProperties)
	if !s.Nil(err, err) {
		s.T().Error(err)
		return
	}

	if len(respProperties) == 0 {
		s.Fail("expected at least 1 property in search result")
		return
	}

	if len(respProperties) < 1 {
		s.Failf(
			"failed to check props list", "expected %d properties in search result but got %d",
			len(searchRequest2.IDs)-1, len(respProperties),
		)
		return
	}
}
