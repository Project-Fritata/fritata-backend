package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Project-Fritata/fritata-backend/internal/apierrors"
	"github.com/Project-Fritata/fritata-backend/internal/env"
	"github.com/Project-Fritata/fritata-backend/services/posts/models"
	"github.com/gofiber/fiber/v3/log"
)

func CheckModerationStatus(post models.Post) (bool, error) {

	getModerationReq := models.GetModerationReq{
		Comment: models.ReqComment{
			Text: post.Content,
		},
		Languages:           []string{"en"},
		RequestedAttributes: models.ReqRequestedAttributes{},
	}

	// Get user data for post
	client := &http.Client{}
	jsonBody, err := json.Marshal(getModerationReq)
	if err != nil {
		log.Errorf("Error marshalling get moderation request: %+v\n%w", getModerationReq, err)
		return false, apierrors.DefaultError()
	}
	req, err := http.NewRequest(
		"POST",
		"https://commentanalyzer.googleapis.com/v1alpha1/comments:analyze?key="+env.GetEnvVar("API_MODERATION_KEY"),
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		log.Errorf("Error getting moderation for request: %+v\n%w", getModerationReq, err)
		return false, apierrors.DefaultError()
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error getting moderation for request: %+v\n%w", getModerationReq, err)
		return false, apierrors.DefaultError()
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Errorf("Error getting moderation for request: %+v\n%w", getModerationReq, err)
		return false, apierrors.DefaultError()
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error reading response body of moderation response: %+v\n%w", getModerationReq, err)
		return false, apierrors.DefaultError()
	}

	var moderationRes models.GetModerationRes
	err = json.Unmarshal(body, &moderationRes)
	if err != nil {
		log.Errorf("Error unmarshalling moderation response body: %s\n%w", string(body), err)
		return false, apierrors.DefaultError()
	}

	toxicityValue := moderationRes.AttributeScores.Toxicity.SummaryScore.Value
	if toxicityValue > 0.5 {
		log.Errorf("Flagged post: %s by user: %s with confidence: %f", post.Content, post.Id_User.String(), toxicityValue)
		return false, fmt.Errorf("moderation system flagged post as innapropriate")
	}

	return true, nil
}
