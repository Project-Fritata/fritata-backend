package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/services/posts/models"
)

func CheckModerationStatus(post models.Post) (bool, error) {

	getModerationReq := models.GetModerationReq{
		Token: internal.GetEnvVar("API_MODERATION_KEY"),
		Text:  post.Content,
	}

	// Get user data for post
	client := &http.Client{}
	jsonBody, err := json.Marshal(getModerationReq)
	if err != nil {
		return false, err
	}
	req, err := http.NewRequest("GET", "https://api.moderatehatespeech.com/api/v1/moderate/", bytes.NewBuffer(jsonBody))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, errors.New("failed to get moderation status in external service")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, errors.New("failed to read response body")
	}

	var moderationRes models.GetModerationRes
	err = json.Unmarshal(body, &moderationRes)
	if err != nil {
		return false, errors.New("failed to parse response body")
	}

	if moderationRes.Class == "flag" {
		fmt.Println("Flagged post: " + post.Content + " by user: " + post.Id_User.String() + " with confidence: " + moderationRes.Confidence)
		return false, nil
	}

	return true, nil
}
