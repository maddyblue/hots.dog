package main

import (
	"database/sql"
	"log"
)

type migration struct {
	ID string
	Up string
}

// mustMigrate panics if it fails.
func mustMigrate(db *sql.DB) {
	migrations := []migration{
		{
			ID: "1",
			Up: `
				CREATE TABLE config (
					key STRING PRIMARY KEY,
					i INT,
					s STRING
				);
				CREATE TABLE builds (
					id STRING PRIMARY KEY,
					start TIMESTAMP,
					finish TIMESTAMP
				);
				CREATE TABLE battletags (
					id INT PRIMARY KEY,
					name STRING,
					INDEX (name)
				);
				CREATE TABLE maps (
					name STRING PRIMARY KEY
				);
				CREATE TABLE heroes (
					slug STRING PRIMARY KEY,
					name STRING,
					roles STRING,
					icon STRING,
					INDEX (name)
				);
				CREATE TABLE games (
					id INT PRIMARY KEY,
					url STRING,
					time TIMESTAMP,

					build STRING,
					length INT,
					map STRING,
					mode INT,
					region INT
				);
				CREATE TABLE players (
					id SERIAL PRIMARY KEY,
					game INT REFERENCES games,
					blizzid INT REFERENCES battletags,

					hero STRING,
					hero_level INT,
					team INT,
					winner BOOL,

					build STRING,
					length INT,
					map STRING,
					mode INT,
					region INT,

					INDEX (blizzid),
					INDEX (build, hero_level, map, mode) STORING (hero, winner),
					INDEX (build, hero_level, mode) STORING (hero, winner)
				);
			`,
		},
		{
			ID: "2",
			Up: `
					CREATE TABLE cache (
						id string PRIMARY KEY,
						until TIMESTAMP,
						data BYTES,
						gzip BYTES
					);
				`,
		},
	}

	const migrateTable = "migrations"

	mustExec(db, `CREATE TABLE IF NOT EXISTS `+migrateTable+` (
		id string PRIMARY KEY,
		created timestamp DEFAULT NOW()
	)`)

	var n int
	seen := make(map[string]bool)
	for _, migration := range migrations {
		// Sanity checks.
		if migration.ID == "" {
			panic("empty migration ID")
		}
		if seen[migration.ID] {
			panic("duplicate ID")
		}
		seen[migration.ID] = true

		// Check if migration has been run already.
		var i int
		if err := db.QueryRow("SELECT count(*) from "+migrateTable+" WHERE id = $1", migration.ID).Scan(&i); err != nil {
			panic(err)
		}
		if i != 0 {
			continue
		}

		// Migrate.
		mustExec(db, migration.Up)
		n++

		mustExec(db, "INSERT INTO "+migrateTable+" (id) VALUES ($1)", migration.ID)
	}
	log.Printf("applied %d migrations", n)
}

func mustExec(db *sql.DB, stmt string, args ...interface{}) {
	if _, err := db.Exec(stmt, args...); err != nil {
		panic(err)
	}
}
