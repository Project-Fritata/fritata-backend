package models

type GetModerationReq struct {
	Comment             ReqComment             `json:"comment"`
	Languages           []string               `json:"languages"`
	RequestedAttributes ReqRequestedAttributes `json:"requestedAttributes"`
}

type ReqComment struct {
	Text string `json:"text"`
}

type ReqRequestedAttributes struct {
	TOXICITY ReqToxicity `json:"TOXICITY"`
}

type ReqToxicity struct {
}

type GetModerationRes struct {
	AttributeScores struct {
		Toxicity struct {
			SpanScores []struct {
				Begin int `json:"begin"`
				End   int `json:"end"`
				Score struct {
					Value float64 `json:"value"`
					Type  string  `json:"type"`
				} `json:"score"`
			} `json:"spanScores"`
			SummaryScore struct {
				Value float64 `json:"value"`
				Type  string  `json:"type"`
			} `json:"summaryScore"`
		} `json:"TOXICITY"`
	} `json:"attributeScores"`
	Languages         []string `json:"languages"`
	DetectedLanguages []string `json:"detectedLanguages"`
}
