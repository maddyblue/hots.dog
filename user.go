package main

import "time"

type Build struct {
	ID            string
	Start, Finish time.Time
}

type Mode int

const (
	ModeQuickMatch Mode = iota + 1
	ModeBrawl
	ModeUnrankedDraft
	ModeHeroLeague
	ModeTeamLeague
)

var gameModes = map[string]Mode{
	"QuickMatch":    ModeQuickMatch,
	"UnrankedDraft": ModeUnrankedDraft,
	"HeroLeague":    ModeHeroLeague,
	"TeamLeague":    ModeTeamLeague,
}

var modeNames = map[Mode]string{
	ModeQuickMatch:    "Quick Match",
	ModeBrawl:         "Brawl",
	ModeUnrankedDraft: "Unranked Draft",
	ModeHeroLeague:    "Hero League",
	ModeTeamLeague:    "Team League",
}
