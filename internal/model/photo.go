package model

import (
	"encoding/json"
	"errors"
	"time"
)

type Photo struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	Caption   string    `json:"caption"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type PhotoUser struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	Caption   string    `json:"caption"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
}

type CreatePhoto struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	Caption string `json:"caption"`
}

func (p *CreatePhoto) ValidateInput() error {
	errVal := map[string]interface{}{}

	if p.Title == "" {
		errVal["title"] = "title cant empty"
	}
	if p.Caption == "" {
		errVal["caption"] = "caption cant empty"
	}
	if p.Url == "" {
		errVal["url"] = "url cant empty"
	}

	jsonErr, err := json.Marshal(errVal)
	if err != nil {
		return err
	}
	if len(errVal) > 0 {
		return errors.New(string(jsonErr))
	}

	return nil
}
