package bet

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/higordasneves/virtualsoccer/domain/entities/soccer"
)

type Client struct {
	URL       string
	Cli       *http.Client
	Cookie    string
	UserAgent string
}

const iso8601Format = "2006-01-02"

func (c Client) FindMatches(ctx context.Context, refDate time.Time) (soccer.Matches, error) {
	url := fmt.Sprintf("%s/ResultsApi/GetFixtures?sportId=146&competitionId=20700663&challengeId=0&fixtureId=0&fromDate=%s&toDate=%s&isDynamic=false&linkId=0&teamId=0&sportDescriptor=FutebolVirtual&ct=28&lng=33&st=413&tz=GMTMinus3&ot=Decimals", strings.TrimSuffix(c.URL, "/"), refDate.Format(iso8601Format), refDate.Format(iso8601Format))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return soccer.Matches{}, fmt.Errorf("creating request %w", err)
	}

	req.Header.Add("Cookie", c.Cookie)
	req.Header.Add("User-Agent", c.UserAgent)

	resp, err := c.Cli.Do(req)
	if err != nil {
		return soccer.Matches{}, fmt.Errorf("request %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return soccer.Matches{}, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	var result soccer.Matches

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return soccer.Matches{}, fmt.Errorf("failed to decode %w", err)
	}

	return result, nil
}

func (c Client) GetMatchResult(ctx context.Context, ft soccer.Fixture) (soccer.MatchResult, error) {
	url := fmt.Sprintf("%s/ResultsApi/GetResults?sportName=sport&sportId=146&fixtureId=%d&competitionId=%d&fromDate=%s&toDate=%s&challengeId=%d&marketOverride=&ct=28&lng=33&st=413&tz=GMTMinus3&ot=Decimals", strings.TrimSuffix(c.URL, "/"), ft.FixtureID, ft.CompetitionID, ft.Time, ft.Time, ft.ChallengeID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return soccer.MatchResult{}, fmt.Errorf("creating request %w", err)
	}

	req.Header.Add("Cookie", c.Cookie)
	req.Header.Add("User-Agent", c.UserAgent)

	resp, err := c.Cli.Do(req)
	if err != nil {
		return soccer.MatchResult{}, fmt.Errorf("request %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return soccer.MatchResult{}, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	var result soccer.MatchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return soccer.MatchResult{}, fmt.Errorf("failed to decode %w", err)
	}

	return result, nil
}

// https://extra.bet365.com/ResultsApi/GetFixtures?sportId=146&competitionId=20700663&challengeId=0&fixtureId=0&fromDate=2023-01-03&toDate=2023-01-03&isDynamic=false&linkId=0&teamId=0&sportDescriptor=FutebolVirtual&ct=28&lng=33&st=413&tz=GMTMinus3&ot=Decimals
// https://extra.bet365.com/ResultsApi/GetResults?sportName=sport&sportId=146&fixtureId=130735294&competitionId=20700663&fromDate=2023-01-03T03:02:00&toDate=2023-01-03T03:02:00&challengeId=84517367&marketOverride=&ct=28&lng=33&st=413&tz=GMTMinus3&ot=Decimals
