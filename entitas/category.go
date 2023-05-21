package entitas

import "time"

type Category struct {
	Id                 uint
	Name               string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	CreatedAtFormatted string
}
