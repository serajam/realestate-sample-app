/* Copyright (C) Fedir Petryk */

package properties

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

type Properties []Property

type Property struct {
	bun.BaseModel `bun:"properties,alias:p"`

	ID            int     `bun:"type:id,pk,autoincrement"`
	UserID        int     `bun:"type:integer,notnull"`
	Location      Point   `bun:"type:geometry(point,4326),notnull"`
	Price         float32 `bun:"type:numeric,notnull"`
	PriceCurrency string  `bun:"price_currency"`

	Address      string `bun:"type:text,notnull"`
	Country      string `bun:"country"`
	City         string `bun:"city"`
	State        string `bun:"state"`
	Street       string `bun:"street"`
	ZipCode      string `bun:"zip_code"`
	HouseNumber  string `bun:"house_number"`
	Neighborhood string `bun:"neighbourhood"`

	HomeSize     float32 `bun:"type:numeric(15,2)"`
	LotSize      float32 `bun:"type:numeric(15,2)"`
	LivingSize   float32 `bun:"type:numeric(15,2)"`
	YearBuild    uint16  `bun:"type:smallint"`
	Bedroom      uint8   `bun:"type:smallint"`
	Bathroom     uint8   `bun:"type:smallint"`
	Floor        uint8   `bun:"floor"`
	TotalFloors  uint8   `bun:"total_floors"`
	PropertyType uint8   `bun:"type:smallint"`
	HomeType     uint8   `bun:"type:smallint"`
	Condition    uint8   `bun:"type:smallint"`

	BrokerName  string `bun:"broker_name"`
	Description string `bun:"type:text"`

	Active       *bool `bun:"active"`
	HasImages    *bool `bun:"has_images"`
	HasGarage    *bool `bun:"has_garage"`
	HasVideo     *bool `bun:"has_video"`
	Has3DTour    *bool `bun:"has_3d_tour"`
	TotalParking uint8 `bun:"total_parking"`
	HasAC        *bool `bun:"has_ac"`

	PetsAllowed *bool `bun:"pets_allowed"`
	Appliance   *bool `bun:"appliance"`
	Heating     uint8 `bun:"heating"`

	Images []*Image `bun:"rel:has-many,join:id=property_id"`
	// `bun:"rel:has-many,join:id=account_id"`

	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp,type:TIMESTAMP"`
	UpdatedAt time.Time `bun:",default:current_timestamp"`
}

func (p Property) ActualDays() int {
	return int(time.Since(p.CreatedAt).Hours() / 24)
}

// PolygonPoint represents an x,y coordinate in EPSG:4326 for PostGIS.
type Point [2]float64

func (p Point) String() string {
	return fmt.Sprintf("SRID=4326;POINT(%v %v)", p[0], p[1])
}

// Scan implements the sql.Scanner interface.
func (p *Point) Scan(val interface{}) error {
	b, err := hex.DecodeString(string(val.([]uint8)))
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	var wkbByteOrder uint8
	if err := binary.Read(r, binary.LittleEndian, &wkbByteOrder); err != nil {
		return err
	}

	var byteOrder binary.ByteOrder
	switch wkbByteOrder {
	case 0:
		byteOrder = binary.BigEndian
	case 1:
		byteOrder = binary.LittleEndian
	default:
		return fmt.Errorf("Invalid byte order %d", wkbByteOrder)
	}

	var wkbGeometryType uint64
	if err := binary.Read(r, byteOrder, &wkbGeometryType); err != nil {
		return err
	}

	if err := binary.Read(r, byteOrder, p); err != nil {
		return err
	}

	return nil
}

// Value impl.
func (p Point) Value() (driver.Value, error) {
	return p.String(), nil
}
