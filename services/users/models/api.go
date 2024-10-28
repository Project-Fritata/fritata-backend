package models

import "github.com/google/uuid"

type GetIdReq struct {
	Id uuid.UUID `json:"id"`
}
type GetUsernameReq struct {
	Username string `json:"username"`
}
type GetRes struct {
	Id          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Pfp         string    `json:"pfp"`
	Description string    `json:"description"`
}

type CreateReq struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type UpdateReq struct {
	Username    string `json:"username"`
	Pfp         string `json:"pfp"`
	Description string `json:"description"`
}
