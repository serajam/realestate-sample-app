package data

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"io"
	"io/fs"
	"log"
)

type TestProperty struct {
	Address struct {
		HouseNumber   string `json:"house_number"`
		Road          string `json:"road"`
		Neighbourhood string `json:"neighbourhood"`
		CityDistrict  string `json:"city_district"`
		Town          string `json:"town"`
		County        string `json:"county"`
		State         string `json:"state"`
		ISO31662Lvl4  string `json:"ISO3166-2-lvl4"`
		Postcode      string `json:"postcode"`
		Country       string `json:"country"`
		CountryCode   string `json:"country_code"`
	} `json:"address"`
	City        string `json:"city"`
	FullAddress string `json:"fullAddress"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
}

func LoadProps(file fs.File) ([]TestProperty, error) {
	ff, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(ff)

	zipListing, err := zip.NewReader(reader, int64(len(ff)))
	if err != nil {
		log.Fatal(err)
	}

	properties := make([]TestProperty, 0, 20000)

	for _, file := range zipListing.File {
		f, err := file.Open()
		if err != nil {
			return nil, err
		}
		err = json.NewDecoder(f).Decode(&properties)
	}

	if err != nil && err != io.EOF {
		return nil, err
	}

	return properties, nil
}
