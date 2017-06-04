package main

type Game struct {
	ID    int64
	Build int64
	Map   string
	Patch string
}

type Player struct {
	ID     int64
	Game   int64
	Build  int64
	Map    string
	Hero   string
	Name   string
	Winner bool
	Team   int64
}

type Talent struct {
	ID     int64
	Game   int64
	Player int64
	Build  int64
	Map    string
	Hero   string
	Tier   int
	Name   string
	Winner bool
}
