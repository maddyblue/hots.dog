package skills

type Team struct {
	PlayerRatings
}

func NewTeam() Team {
	return Team{make(PlayerRatings)}
}

func (t Team) AddPlayer(p interface{}, r Rating) {
	t.PlayerRatings[p] = r
}

func (t Team) PlayerCount() int {
	return len(t.PlayerRatings)
}

func (t Team) Players() []interface{} {
	ps := []interface{}{}
	for p := range t.PlayerRatings {
		ps = append(ps, p)
	}
	return ps
}

func (t Team) PlayerRating(p interface{}) Rating {
	return t.PlayerRatings[p]
}
