/* Copyright (C) Fedir Petryk */

package properties

import (
	"fmt"
	"strings"
)

const TwoPointPolygonFormat = `SRID=4326;POLYGON((%s %s, %s %s, %s %s, %s %s, %s %s))`
const PolygonFormat = `SRID=4326;POLYGON((%s))`

type TwoPointPolygon struct {
	topLat  string
	topLong string

	bottomLat  string
	bottomLong string
}

func NewPolygon(topLat string, topLong string, bottomLat string, bottomLong string) TwoPointPolygon {
	return TwoPointPolygon{
		topLat:     topLat,
		topLong:    topLong,
		bottomLat:  bottomLat,
		bottomLong: bottomLong,
	}
}

func (p TwoPointPolygon) String() string {
	return fmt.Sprintf(
		TwoPointPolygonFormat,
		p.topLat, p.topLong,
		p.topLat, p.bottomLong,
		p.bottomLat, p.bottomLong,
		p.bottomLat, p.topLong,
		p.topLat, p.topLong,
	)
}

type PolygonPoint struct {
	Coordinates [2]float64
}

func (p PolygonPoint) String() string {
	return fmt.Sprintf("%v %v", p.Coordinates[0], p.Coordinates[1])
}

type Polygon struct {
	Points []PolygonPoint
}

func (p Polygon) String() string {
	polygonBuilder := strings.Builder{}

	for _, point := range p.Points {
		polygonBuilder.WriteString(point.String() + ", ")
	}

	return fmt.Sprintf(PolygonFormat, polygonBuilder.String()[:polygonBuilder.Len()-2])
}

type GeoObject struct {
	Polygon Polygon
}
