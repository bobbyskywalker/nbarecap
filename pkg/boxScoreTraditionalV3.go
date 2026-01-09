package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	endpointV3     = "https://stats.nba.com/stats/boxscoretraditionalv3?GameID=%s&StartPeriod=0&EndPeriod=0&StartRange=0&EndRange=0&RangeType=0"
	defaultTimeout = 20 * time.Second
)

func setReqHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Referer", "https://www.nba.com/")
	req.Header.Set("Origin", "https://www.nba.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("x-nba-stats-origin", "stats")
	req.Header.Set("x-nba-stats-token", "true")
}

func FetchBoxScoreTraditionalV3JSON(gameID string) (json.RawMessage, error) {
	url := fmt.Sprintf(
		endpointV3,
		gameID,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	setReqHeaders(req)

	client := &http.Client{
		Timeout: defaultTimeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("nba stats HTTP %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}
