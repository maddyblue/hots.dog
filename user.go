package main

import "time"

type Game struct {
	ID int64

	Patch string
	Seed  int64
	Time  time.Time

	Build  int64
	Length int64
	Map    string
	Mode   Mode
}

type Player struct {
	ID   int64
	Game int64

	Hero   string
	Name   string
	Team   int64
	Winner bool

	Build  int64
	Length int64
	Map    string
	Mode   Mode
}

type Talent struct {
	ID     int64
	Game   int64
	Player int64

	Hero   string
	Name   string
	Tier   int
	Winner bool

	Build  int64
	Length int64
	Map    string
	Mode   Mode
}

type Mode int

const (
	ModeQuickMatch Mode = iota + 1
	ModeBrawl
	ModeUnrankedDraft
	ModeHeroLeague
	ModeTeamLeague
)

var modeMap = map[int64]Mode{
	50001: ModeQuickMatch,
	50031: ModeBrawl,
	50051: ModeUnrankedDraft,
	50061: ModeHeroLeague,
	50071: ModeTeamLeague,
}

var modeNames = map[Mode]string{
	ModeQuickMatch:    "Quick Match",
	ModeBrawl:         "Brawl",
	ModeUnrankedDraft: "Unranked Draft",
	ModeHeroLeague:    "Hero League",
	ModeTeamLeague:    "Team League",
}
