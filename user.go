package main

type Game struct {
	Build int64
	Map   string
	ID    int64
	Patch string
}

type Player struct {
	Build  int64
	Map    string
	Game   int64
	Hero   string
	Name   string
	Winner bool
	Team   int64
}

type Talent struct {
	Build  int64
	Map    string
	Game   int64
	Hero   string
	Tier   int
	Name   string
	Winner bool
}
