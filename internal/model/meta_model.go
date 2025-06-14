package model

import "time"

type Meta struct {
	CreatedAt int64 `bson:"created_at" json:"created_at"`
	UpdatedAt int64 `bson:"updated_at" json:"updated_at"`
	Archived  bool  `bson:"archived" json:"archived"`
}

func NewMeta() *Meta {
	now := time.Now().Unix()
	return &Meta{
		CreatedAt: now,
		UpdatedAt: now,
		Archived:  false,
	}
}
