package posts

import "github.com/Project-Fritata/fritata-backend/internal"

type GetReq struct {
	Offset int `query:"offset"`
	Limit  int `query:"limit"`
}

type GetRes struct {
	Post internal.Post `json:"post"`
	User internal.User `json:"user"`
}

type PostReq struct {
	Content string `json:"content"`
	Media   string `json:"media"`
}
