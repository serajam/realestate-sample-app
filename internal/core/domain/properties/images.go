/* Copyright (C) Fedir Petryk */

package properties

import (
	"bytes"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Image struct {
	bun.BaseModel `bun:"table:property_images"`

	ID          uuid.UUID `bun:"type:uuid"`
	PropertyID  int       `bun:"type:integer,notnull"`
	UserID      int       `bun:"type:integer,notnull"`
	Description string    `bun:"type:text"`
	Position    int       `bun:"type:smallint,notnull"`
	Contents    []Content `bun:"-"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp,type:TIMESTAMP"`
}

func ImageObjectName(id, imgType string) string {
	return fmt.Sprintf("%s.%s", id, imgType)
}

type Dimensions struct {
	Width     float64
	Height    float64
	ImageType string
}

type Content struct {
	Content   *bytes.Buffer
	Size      int64
	ImageType string
}
