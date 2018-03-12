package skills

import (
	"fmt"
)

type RankedTeams struct {
	teams []Team
	ranks []int
}

func NewRankedTeams(teams []Team, ranks []int) *RankedTeams {
	if len(teams) != len(ranks) {
		panic(fmt.Errorf("Number of teams [%v] does not match number of ranks [%v]", len(teams), len(ranks)))
	}
	return &RankedTeams{teams, ranks}
}

func (rt *RankedTeams) AddTeam(team Team, rank int) {
	rt.teams = append(rt.teams, team)
	rt.ranks = append(rt.ranks, rank)
}

func (rt *RankedTeams) Len() int           { return len(rt.teams) }
func (rt *RankedTeams) Less(i, j int) bool { return rt.ranks[i] < rt.ranks[j] }

func (rt *RankedTeams) Swap(i, j int) {
	rt.teams[i], rt.teams[j] = rt.teams[j], rt.teams[i]
	rt.ranks[i], rt.ranks[j] = rt.ranks[j], rt.ranks[i]
}
