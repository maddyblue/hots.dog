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
				CREATE TABLE cache (
					id STRING PRIMARY KEY,
					until TIMESTAMP,
					data BYTES,
					gzip BYTES,
					last_hit TIMESTAMP
				);

				CREATE TABLE config (
					key STRING PRIMARY KEY,
					i INT NULL,
					s STRING NULL
				);
			`,
		},
		{
			ID: "2",
			Up: `
				CREATE TABLE playerskills (
					region INT,
					blizzid INT,
					build INT,
					mode INT,
					skill INT,
					PRIMARY KEY (region, blizzid, build, mode)
				);

				CREATE TABLE skillstats (
					build INT,
					mode INT,
					data BYTES,
					PRIMARY KEY (build, mode)
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
	// Clear the cache because implementations may have changed. This assumes
	// the cron job is running the correct image, which may not be true.
	mustExec(db, `UPDATE cache SET until = NULL`)
	log.Printf("applied %d migrations", n)
}
