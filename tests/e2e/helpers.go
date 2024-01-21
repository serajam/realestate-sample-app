package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

const apiURL = "http://localhost:8080"

func checkApiAvailable(t *testing.T, url string) (*http.Response, error) {
	healthUrl := url + "/health"
	ticker := time.NewTicker(time.Second * 10)
	maxRetries := 20
	for {
		if maxRetries == 0 {
			return nil, errors.New("Maximum attempts reached. Api offline")
		}

		resp, err := http.Get(healthUrl)
		if err != nil {
			t.Logf("Api offline: %s", err)
			t.Logf("Retrying in 10 seconds. Retries left: %d", maxRetries)
			maxRetries--

			select {
			case <-ticker.C:
			}
			continue
		}

		return resp, nil
	}
}

func checkPostRequest(url string, expectedStatus int, request interface{}) (respBody []byte, status int, err error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, 0, err
	}

	resp, err := http.Post(apiURL+url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, 0, err
	}

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	if resp.StatusCode != expectedStatus {
		respErr := ErrorResponse{}
		err := json.NewDecoder(resp.Body).Decode(&respErr)
		if err != nil && err != io.EOF {
			return nil, 0, err
		}

		return respBody, resp.StatusCode, fmt.Errorf(
			"expected status 200, got %d, error: %s, %s", resp.StatusCode, respErr.Message, respErr.Error,
		)
	}

	return respBody, resp.StatusCode, nil
}

func unmarshalTo(body []byte, to interface{}) error {
	var defeaultResp DefaultResponse

	err := json.Unmarshal(body, &defeaultResp)
	if err != nil {
		return err
	}

	err = json.Unmarshal(defeaultResp.Data, to)
	if err != nil {
		return err
	}

	return nil
}

func NewTestPropertyRequest() PropertyCreateRequest {
	maxLat := 48.20396774477816
	maxLong := 17.0443878146695
	minLat := 48.105638
	minLong := 17.203723

	return PropertyCreateRequest{
		Location:      Location{gofakeit.Float64Range(minLat, maxLat), gofakeit.Float64Range(maxLong, minLong)},
		Price:         gofakeit.Float32Range(10000, 5000000),
		PriceCurrency: "EUR",
		FullAddress:   gofakeit.Address().Address,
		Address: Address{
			Country:      gofakeit.CountryAbr(),
			City:         gofakeit.City(),
			State:        gofakeit.State(),
			Street:       gofakeit.Street(),
			ZipCode:      gofakeit.Zip(),
			HouseNumber:  gofakeit.StreetNumber(),
			Neighborhood: gofakeit.State(),
		},
		HomeSize:       gofakeit.Float32Range(50, 500),
		LotSize:        gofakeit.Float32Range(50, 500),
		LivingSize:     gofakeit.Float32Range(50, 500),
		YearBuild:      uint16(gofakeit.Year()),
		Bedroom:        uint8(gofakeit.IntRange(1, 5)),
		Bathroom:       uint8(gofakeit.IntRange(1, 5)),
		Floor:          uint8(gofakeit.IntRange(1, 100)),
		TotalFloors:    uint8(gofakeit.IntRange(1, 100)),
		PropertyType:   uint8(gofakeit.IntRange(1, 3)),
		HomeType:       uint8(gofakeit.IntRange(1, 9)),
		Condition:      uint8(gofakeit.IntRange(1, 3)),
		BrokerName:     gofakeit.Name(),
		Description:    gofakeit.Sentence(100),
		IsActiveStatus: lo.ToPtr(gofakeit.Bool()),
		HasImages:      lo.ToPtr(gofakeit.Bool()),
		HasGarage:      lo.ToPtr(gofakeit.Bool()),
		HasVideo:       lo.ToPtr(gofakeit.Bool()),
		Has3DTour:      lo.ToPtr(gofakeit.Bool()),
		TotalParking:   uint8(gofakeit.IntRange(1, 5)),
		HasAC:          lo.ToPtr(gofakeit.Bool()),
		Appliance:      lo.ToPtr(gofakeit.Bool()),
		Heating:        uint8(gofakeit.IntRange(1, 3)),
		PetsAllowed:    lo.ToPtr(gofakeit.Bool()),
	}
}
