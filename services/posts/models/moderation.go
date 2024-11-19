package models

type GetModerationReq struct {
	Token string `json:"token"`
	Text  string `json:"text"`
}

type GetModerationRes struct {
	Response   string `json:"response"`
	Class      string `json:"class"`
	Confidence string `json:"confidence"`
}
