package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/xuri/excelize/v2"

	"github.com/higordasneves/virtualsoccer/domain/entities/soccer"
	"github.com/higordasneves/virtualsoccer/gateway/bet"
)

var mu sync.Mutex
var contagem int

func main() {
	cookie := "aps03=ct=28&lng=33; state=413; qBvrxRyB=A8zHE3-FAQAALCPORJIv8qodoJGX5603G66Pl1bD2lZiMrvDg-omgEEtwYVxAbF5-e6ucgBSwH8AADQwAAAAAA|1|1|9facc92e1a971171a0f923d6bd670d801ca70e9c; bet365SportsExtra=settings=0,0,0,0,0,16,0,,1,1; swt=Adn7oKPg+h/9Dv/7t51L0XwFls5SdproK4VfJTZqyrPDfMZvyWY0iC/Pi5MIXvhypdtkB6ABIXM+ig3UJLBR906QA0nkfWgNOQ59TJ+4apQyAQUOxJi7; _ga=GA1.1.1293370969.1672874290; cc=1; Affiliates=Code=365_01012170/162022043193&prd=Sports; pstk=09BACF11D531E0DDADF0866C36E4DC73000004; aaat=di=34fb534d-c308-4991-92be-b6475b74ac44&am=0&at=34eaf09b-b9c8-4c30-b601-262f46473054&un=frenesi1&ts=05-01-2023 19:32:52&v=2; _ga_45M1DQFW2B=GS1.1.1672949722.8.0.1672949722.0.0.0; __cf_bm=xNOlZlCqB5N1gQGmGToCo97lcCaIKp3gWpjOQhrBRGA-1672960368-0-AW+HbmcxwV7h9LL0mz7HUizpdJI8pIh1yDS5k+IlK2c6vFDJdYyE7c1WSrkj3zhGvEr+TkrcgJ4IiY/xKQpWsh4="
	agent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
	ctx := context.Background()

	client := bet.Client{
		URL:       "https://extra.bet365.com",
		Cli:       &http.Client{},
		Cookie:    cookie,
		UserAgent: agent,
	}

	log.Print("initializing")

	from := time.Date(2022, 12, 1, 0, 0, 0, 0, time.Local)
	to := time.Date(2023, 1, 5, 0, 0, 0, 0, time.Local)

	filename := "matches"
	for d := from; d.Before(to); d = d.AddDate(0, 0, 1) {
		txtReport := make([][]string, 0)
		log.Printf("initializing day %s", d)
		txtReport = append(txtReport, generateDailyReport(ctx, d, client)...)
		generateFile(filename, txtReport)
		log.Printf("finish for %s", d)
	}

	log.Print("finish", contagem)
}

func generateDailyReport(ctx context.Context, date time.Time, client bet.Client) [][]string {
	matches, err := client.FindMatches(ctx, date)
	if err != nil {
		log.Print(err)
		return [][]string{}
	}
	contagem += len(matches.Fixtures)

	txtReport := make([][]string, 0)
	c := make(chan soccer.Fixture)
	go func() {
		for _, ft := range matches.Fixtures {
			c <- ft
		}
		close(c)
	}()

	results := make([]soccer.MatchResult, 0, len(matches.Fixtures))
	wg := sync.WaitGroup{}
	for ft := range c {
		v := ft
		wg.Add(1)
		go func() {
			result, err := client.GetMatchResult(ctx, v)
			if err != nil {
				log.Print(err)
			} else {
				mu.Lock()
				results = append(results, result)
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	for _, result := range results {
		report := soccer.ParseMatchResult(result)
		txtReport = append(txtReport, []string{
			report.Date,
			report.Title,
			report.FinalResult,
			report.ExactScore,
			report.ExactGoalNumber,
			strconv.FormatBool(report.OverHalfAGoal),
			strconv.FormatBool(report.UnderHalfAGoal),
			strconv.FormatBool(report.OverOneAndAHalfAGoal),
			strconv.FormatBool(report.UnderOneAndAHalfAGoal),
			strconv.FormatBool(report.OverTwoAndAHalfAGoal),
			strconv.FormatBool(report.UnderTwoAndAHalfAGoal),
			strconv.FormatBool(report.OverThreeAndAHalfAGoal),
			strconv.FormatBool(report.UnderThreeAndAHalfAGoal),
			report.HomeTeamGoals,
			report.InvitedTeamGoals,
			strconv.FormatBool(report.BothTeamsToScore),
			strconv.FormatBool(report.HomeTeamToScore),
			strconv.FormatBool(report.InvitedTeamToScore),
		})
	}

	return txtReport
}

func generateFile(filename string, report [][]string) {
	sheetName := "results"
	rowNumber := 1

	xlsx, err := excelize.OpenFile(filename + ".xlsx")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			xlsx = excelize.NewFile()
			oldSheetName := xlsx.GetSheetName(0)
			xlsx.SetSheetName(oldSheetName, sheetName)
			xlsx.SaveAs(filename + ".xlsx")

			var header = []string{
				"date", "title", "final_result", "exact_score", "exact_goal_number",
				"over_half_a_goal", "under_half_a_goal", "over_one_and_a_half_a_goal", "under_one_and_a_half_a_goal",
				"over_two_and_a_half_a_goal", "under_two_and_a_half_a_goal", "over_three_and_a_half_a_goal", "under_three_and_a_half_a_goal",
				"home_team_goals", "invited_team_goals", "both_teams_to_score", "home_team_to_score", "invited_team_to_score",
			}
			xlsx.SetSheetRow(sheetName, "A"+strconv.Itoa(rowNumber), &header)
		}
	}

	rows, _ := xlsx.Rows(sheetName)
	for rows.Next() {
		rowNumber++
	}

	for _, r := range report {
		xlsx.SetSheetRow(sheetName, "A"+strconv.Itoa(rowNumber), &r)
		rowNumber++
	}

	xlsx.SaveAs(filename + ".xlsx")
	xlsx.Close()
}
