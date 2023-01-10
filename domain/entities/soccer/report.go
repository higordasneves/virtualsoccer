package soccer

import (
	"strings"
)

func ParseMatchResult(mr MatchResult) (report MatchResultReport) {
	report.Title = mr.Title
	report.Date = mr.Date

	var teamGolsReported bool
	for _, mm := range mr.MatchMarkets {
		splitedMM := strings.Split(mm, "|=")
		switch splitedMM[0] {
		case "Fulltime Result":
			report.FinalResult = extractUniqueWon(splitedMM[1])
		case "Correct Score":
			report.ExactScore = extractUniqueWon(splitedMM[1])
		case "Total de Gols Exatos":
			report.ExactGoalNumber = extractUniqueWon(splitedMM[1])
		case "Total de Gols - Mais de/Menos de 0.5":
			fieldValue := extractUniqueWon(splitedMM[1])
			report.OverHalfAGoal = strings.Contains(fieldValue, "Mais")
			report.UnderHalfAGoal = !report.OverHalfAGoal
		case "Total de Gols - Mais de/Menos de 1.5":
			fieldValue := extractUniqueWon(splitedMM[1])
			report.OverOneAndAHalfAGoal = strings.Contains(fieldValue, "Mais")
			report.UnderOneAndAHalfAGoal = !report.OverOneAndAHalfAGoal
		case "Total de Gols - Mais de/Menos de 2.5":
			fieldValue := extractUniqueWon(splitedMM[1])
			report.OverTwoAndAHalfAGoal = strings.Contains(fieldValue, "Mais")
			report.UnderTwoAndAHalfAGoal = !report.OverTwoAndAHalfAGoal
		case "Total Goals Over/Under 3.5":
			fieldValue := extractUniqueWon(splitedMM[1])
			report.OverThreeAndAHalfAGoal = strings.Contains(fieldValue, "Mais")
			report.UnderThreeAndAHalfAGoal = !report.OverThreeAndAHalfAGoal
		case "Time -  Gols":
			if teamGolsReported {
				report.InvitedTeamGoals = extractUniqueWon(splitedMM[1])
			} else {
				report.HomeTeamGoals = extractUniqueWon(splitedMM[1])
				teamGolsReported = true
			}
		case "Both Teams to Score":
			fieldValue := extractUniqueWon(splitedMM[1])
			report.BothTeamsToScore = strings.Contains(fieldValue, "Sim")
		case "Home Team To Score":
			fieldValue := extractUniqueWon(splitedMM[1])
			report.HomeTeamToScore = strings.Contains(fieldValue, "Sim")
		case "Away Team To Score":
			fieldValue := extractUniqueWon(splitedMM[1])
			report.InvitedTeamToScore = strings.Contains(fieldValue, "Sim")
		}
	}

	return report
}

func extractUniqueWon(s string) string {
	values := strings.Split(s, "|")
	for _, v := range values {
		if strings.Contains(v, "~Won~") {
			return strings.ReplaceAll(v, "~Won~", "")
		}
	}

	return ""
}
