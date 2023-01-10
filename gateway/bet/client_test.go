package bet

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/higordasneves/virtualsoccer/domain/entities/soccer"
)

func TestClient_FindMatches(t *testing.T) {
	cookie := "aps03=ct=28&lng=33; state=413; qBvrxRyB=A8zHE3-FAQAALCPORJIv8qodoJGX5603G66Pl1bD2lZiMrvDg-omgEEtwYVxAbF5-e6ucgBSwH8AADQwAAAAAA|1|1|9facc92e1a971171a0f923d6bd670d801ca70e9c; pstk=EAC31D55E875459C92DA1110AE843BF3000004; aaat=di=34fb534d-c308-4991-92be-b6475b74ac44&am=0&at=cc843841-bc29-4bdb-b287-cf797afe201d&un=frenesi1&ts=04-01-2023 23:18:08&v=2; bet365SportsExtra=settings=0,0,0,0,0,16,0,,1,1; swt=Adn7oKPg+h/9Dv/7t51L0XwFls5SdproK4VfJTZqyrPDfMZvyWY0iC/Pi5MIXvhypdtkB6ABIXM+ig3UJLBR906QA0nkfWgNOQ59TJ+4apQyAQUOxJi7; _ga=GA1.1.1293370969.1672874290; cc=1; _ga_45M1DQFW2B=GS1.1.1672874289.1.1.1672877365.0.0.0; __cf_bm=sQc3FhG37CY7kY0KIaWXJZNd2DJudtMbt_31J8PPuhY-1672887376-0-Ab7y2fU5NFSg8SlmuN1y6c925HBwfn+GAd6sZ20WAwq4qrCeGxUEWhoCgzsO5ykB9iQbkeK51dD0YL24qrP+6m8="
	agent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"

	httpClient := Client{
		URL:       "https://extra.bet365.com",
		Cli:       &http.Client{},
		Cookie:    cookie,
		UserAgent: agent,
	}

	refDate := time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local)
	_, err := httpClient.FindMatches(context.Background(), refDate)
	assert.NoError(t, err)
}

func TestClient_GetMatchResult(t *testing.T) {
	cookie := "aps03=ct=28&lng=33; state=413; qBvrxRyB=A8zHE3-FAQAALCPORJIv8qodoJGX5603G66Pl1bD2lZiMrvDg-omgEEtwYVxAbF5-e6ucgBSwH8AADQwAAAAAA|1|1|9facc92e1a971171a0f923d6bd670d801ca70e9c; pstk=EAC31D55E875459C92DA1110AE843BF3000004; aaat=di=34fb534d-c308-4991-92be-b6475b74ac44&am=0&at=cc843841-bc29-4bdb-b287-cf797afe201d&un=frenesi1&ts=04-01-2023 23:18:08&v=2; bet365SportsExtra=settings=0,0,0,0,0,16,0,,1,1; swt=Adn7oKPg+h/9Dv/7t51L0XwFls5SdproK4VfJTZqyrPDfMZvyWY0iC/Pi5MIXvhypdtkB6ABIXM+ig3UJLBR906QA0nkfWgNOQ59TJ+4apQyAQUOxJi7; _ga=GA1.1.1293370969.1672874290; cc=1; _ga_45M1DQFW2B=GS1.1.1672874289.1.1.1672877365.0.0.0; __cf_bm=6kE_rw2_Un3MABFm0cvynqcQCpEA_heGaQ1qGEYk2XY-1672894530-0-AbiQslDaR76+iJJxPHX8QB8LiDuK4XrzbjxlBn7XNwhlQaV1uHZFlof9CGkaEg3SeSk4xNuGHgg5lwQnD+WtiOA="
	agent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"

	httpClient := Client{
		URL:       "https://extra.bet365.com",
		Cli:       &http.Client{},
		Cookie:    cookie,
		UserAgent: agent,
	}

	ft := soccer.Fixture{
		DateDesc:      "03 jan 2023",
		Desc:          "3.02  Macedônia do Norte vs Áustria",
		FixtureID:     130735294,
		ChallengeID:   84517367,
		CompetitionID: 20700663,
		Title:         nil,
		Time:          "2023-01-03T03:02:00",
		UserTime:      "2023-01-03T00:02:00",
	}

	_, err := httpClient.GetMatchResult(context.Background(), ft)
	assert.NoError(t, err)
}
