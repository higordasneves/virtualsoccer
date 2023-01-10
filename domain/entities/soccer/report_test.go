package soccer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMatchResult(t *testing.T) {
	tests := []struct {
		name       string
		arg        MatchResult
		wantReport MatchResultReport
	}{
		{
			name: "case 1",
			arg: MatchResult{
				Date:  "2023-01-03T00:02:00",
				Title: "3.02  Macedônia do Norte vs Áustria",
				MatchMarkets: []string{
					"Fulltime Result|=Macedônia do Norte~Won~|Áustria~~|Empate~~",
					"Correct Score|=Macedônia do Norte #2-1~Won~|Áustria #1-0~~|Áustria #2-0~~|Áustria #2-1~~|Áustria #3-0~~|Áustria #3-1~~|Áustria #3-2~~|Áustria #4-0~~|Empate #0-0~~|Empate #1-1~~|Empate #2-2~~|Macedônia do Norte #1-0~~|Macedônia do Norte #2-0~~|Macedônia do Norte #3-0~~|Macedônia do Norte #3-1~~|Macedônia do Norte #3-2~~|Macedônia do Norte #4-0~~|Para Áustria ganhar por qualquer outro resultado~~|Para Macedônia do Norte ganhar por qualquer outro resultado~~|Qualquer outro resultado de empate~~",
					"Total de Gols Exatos|=3 Gols~Won~|0 Golos~~|1 Gol~~|2 Gols~~|4 Gols~~|5+ Golos~~",
					"Total de Gols - Mais de/Menos de 0.5|=Mais de 0.5~Won~|Menos de 0.5~~",
					"Total de Gols - Mais de/Menos de 1.5|=Mais de 1.5~Won~|Menos de 1.5~~",
					"Total de Gols - Mais de/Menos de 2.5|=Mais de 2.5~Won~|Menos de 2.5~~",
					"Total Goals Over/Under 3.5|=Menos de 3.5~Won~|Mais de 3.5~~",
					"Time -  Gols|=Macedônia do Norte - 2 Golos~Won~|Macedônia do Norte - 0 Golos~~|Macedônia do Norte - 1 Golo~~|Macedônia do Norte - 3 Golos~~|Macedônia do Norte - 4 Golos~~|Macedônia do Norte - 5+ Golos~~",
					"Time -  Gols|=Áustria - 1 Golo~Won~|Áustria - 0 Golos~~|Áustria - 2 Golos~~|Áustria - 3 Golos~~|Áustria - 4 Golos~~|Áustria - 5+ Golos~~",
					"Both Teams to Score|=Sim~Won~|Sem Gols~~",
					"Home Team To Score|=Sim~Won~|Não~~",
					"Away Team To Score|=Sim~Won~|Não~~",
				},
			},
			wantReport: MatchResultReport{
				Date:                    "2023-01-03T00:02:00",
				Title:                   "3.02  Macedônia do Norte vs Áustria",
				FinalResult:             "Macedônia do Norte",
				ExactScore:              "Macedônia do Norte #2-1",
				ExactGoalNumber:         "3 Gols",
				OverHalfAGoal:           true,
				UnderHalfAGoal:          false,
				OverOneAndAHalfAGoal:    true,
				UnderOneAndAHalfAGoal:   false,
				OverTwoAndAHalfAGoal:    true,
				UnderTwoAndAHalfAGoal:   false,
				OverThreeAndAHalfAGoal:  false,
				UnderThreeAndAHalfAGoal: true,
				HomeTeamGoals:           "Macedônia do Norte - 2 Golos",
				InvitedTeamGoals:        "Áustria - 1 Golo",
				BothTeamsToScore:        true,
				HomeTeamToScore:         true,
				InvitedTeamToScore:      true,
			},
		},
		{
			name: "case 0x0",
			arg: MatchResult{
				MatchMarkets: []string{
					"Total de Gols - Mais de/Menos de 0.5|=Mais de 0.5~~|Menos de 0.5~Won~",
					"Total de Gols - Mais de/Menos de 1.5|=Mais de 1.5~~|Menos de 1.5~Won~",
					"Total de Gols - Mais de/Menos de 2.5|=Mais de 2.5~~|Menos de 2.5~Won~",
					"Total Goals Over/Under 3.5|=Menos de 3.5~Won~|Mais de 3.5~~",
					"Both Teams to Score|=Sim~~|Sem Gols~Won~",
					"Home Team To Score|=Sim~~|Não~Won~",
					"Away Team To Score|=Sim~~|Não~Won~",
				},
			},
			wantReport: MatchResultReport{

				OverHalfAGoal:           false,
				UnderHalfAGoal:          true,
				OverOneAndAHalfAGoal:    false,
				UnderOneAndAHalfAGoal:   true,
				OverTwoAndAHalfAGoal:    false,
				UnderTwoAndAHalfAGoal:   true,
				OverThreeAndAHalfAGoal:  false,
				UnderThreeAndAHalfAGoal: true,
				BothTeamsToScore:        false,
				HomeTeamToScore:         false,
				InvitedTeamToScore:      false,
			},
		},
		{
			name: "case 1x0",
			arg: MatchResult{
				MatchMarkets: []string{
					"Total de Gols - Mais de/Menos de 0.5|=Mais de 0.5~Won~|Menos de 0.5~~",
					"Total de Gols - Mais de/Menos de 1.5|=Mais de 1.5~~|Menos de 1.5~Won~",
					"Total de Gols - Mais de/Menos de 2.5|=Mais de 2.5~~|Menos de 2.5~Won~",
					"Total Goals Over/Under 3.5|=Menos de 3.5~Won~|Mais de 3.5~~",
					"Both Teams to Score|=Sim~~|Sem Gols~Won~",
					"Home Team To Score|=Sim~Won~|Não~~",
					"Away Team To Score|=Sim~~|Não~Won~",
				},
			},
			wantReport: MatchResultReport{

				OverHalfAGoal:           true,
				UnderHalfAGoal:          false,
				OverOneAndAHalfAGoal:    false,
				UnderOneAndAHalfAGoal:   true,
				OverTwoAndAHalfAGoal:    false,
				UnderTwoAndAHalfAGoal:   true,
				OverThreeAndAHalfAGoal:  false,
				UnderThreeAndAHalfAGoal: true,
				BothTeamsToScore:        false,
				HomeTeamToScore:         true,
				InvitedTeamToScore:      false,
			},
		},
		{
			name: "case 1x1",
			arg: MatchResult{
				MatchMarkets: []string{
					"Total de Gols - Mais de/Menos de 0.5|=Mais de 0.5~Won~|Menos de 0.5~~",
					"Total de Gols - Mais de/Menos de 1.5|=Mais de 1.5~Won~|Menos de 1.5~~",
					"Total de Gols - Mais de/Menos de 2.5|=Mais de 2.5~~|Menos de 2.5~Won~",
					"Total Goals Over/Under 3.5|=Menos de 3.5~Won~|Mais de 3.5~~",
					"Both Teams to Score|=Sim~Won~|Sem Gols~~",
					"Home Team To Score|=Sim~Won~|Não~~",
					"Away Team To Score|=Sim~Won~|Não~~",
				},
			},
			wantReport: MatchResultReport{

				OverHalfAGoal:           true,
				UnderHalfAGoal:          false,
				OverOneAndAHalfAGoal:    true,
				UnderOneAndAHalfAGoal:   false,
				OverTwoAndAHalfAGoal:    false,
				UnderTwoAndAHalfAGoal:   true,
				OverThreeAndAHalfAGoal:  false,
				UnderThreeAndAHalfAGoal: true,
				BothTeamsToScore:        true,
				HomeTeamToScore:         true,
				InvitedTeamToScore:      true,
			},
		},
		{
			name: "case 0x4",
			arg: MatchResult{
				MatchMarkets: []string{
					"Total de Gols - Mais de/Menos de 0.5|=Mais de 0.5~Won~|Menos de 0.5~~",
					"Total de Gols - Mais de/Menos de 1.5|=Mais de 1.5~Won~|Menos de 1.5~~",
					"Total de Gols - Mais de/Menos de 2.5|=Mais de 2.5~Won~|Menos de 2.5~~",
					"Total Goals Over/Under 3.5|=Menos de 3.5~~|Mais de 3.5~Won~",
					"Both Teams to Score|=Sim~~|Sem Gols~Won~",
					"Home Team To Score|=Sim~~|Não~Won~",
					"Away Team To Score|=Sim~Won~|Não~~",
				},
			},
			wantReport: MatchResultReport{

				OverHalfAGoal:           true,
				UnderHalfAGoal:          false,
				OverOneAndAHalfAGoal:    true,
				UnderOneAndAHalfAGoal:   false,
				OverTwoAndAHalfAGoal:    true,
				UnderTwoAndAHalfAGoal:   false,
				OverThreeAndAHalfAGoal:  true,
				UnderThreeAndAHalfAGoal: false,
				BothTeamsToScore:        false,
				HomeTeamToScore:         false,
				InvitedTeamToScore:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseMatchResult(tt.arg)
			assert.Equal(t, tt.wantReport, result)
		})
	}
}
