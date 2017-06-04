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
				CREATE TABLE games (
					id SERIAL PRIMARY KEY,
					build INT,
					map STRING,
					patch STRING
				);
				CREATE TABLE players (
					id SERIAL PRIMARY KEY,
					game INT REFERENCES games,
					build INT,
					map STRING,
					hero STRING,
					winner BOOL,
					name STRING,
					team INT,
					INDEX (map, hero) STORING (winner)
				);
				CREATE TABLE talents (
					id SERIAL PRIMARY KEY,
					game INT REFERENCES games,
					player INT REFERENCES players,
					build INT,
					map STRING,
					hero STRING,
					winner BOOL,
					tier INT,
					name STRING
				);
				CREATE TABLE maps (
					name STRING PRIMARY KEY
				);
				INSERT INTO maps VALUES
					('Battlefield of Eternity'),
					('Blackheart''s Bay'),
					('Braxis Holdout'),
					('Cursed Hollow'),
					('Dragon Shire'),
					('Garden of Terror'),
					('Hanamura'),
					('Haunted Mines'),
					('Infernal Shrines'),
					('Sky Temple'),
					('Tomb of the Spider Queen'),
					('Towers of Doom'),
					('Warhead Junction')
				;
				CREATE TABLE heroes (
					name STRING PRIMARY KEY
				);
				INSERT INTO heroes VALUES
					('Abathur'),
					('Alarak'),
					('Anub''arak'),
					('Artanis'),
					('Arthas'),
					('Auriel'),
					('Azmodan'),
					('Brightwing'),
					('Cassia'),
					('Chen'),
					('Cho'),
					('Chromie'),
					('Dehaka'),
					('Diablo'),
					('D.Va'),
					('E.T.C.'),
					('Falstad'),
					('Gall'),
					('Gazlowe'),
					('Genji'),
					('Greymane'),
					('Gul''dan'),
					('Illidan'),
					('Jaina'),
					('Johanna'),
					('Kael''thas'),
					('Kerrigan'),
					('Kharazim'),
					('Leoric'),
					('Li Li'),
					('Li-Ming'),
					('Lt. Morales'),
					('Lúcio'),
					('Lunara'),
					('Malfurion'),
					('Malthael'),
					('Medivh'),
					('Muradin'),
					('Murky'),
					('Nazeebo'),
					('Nova'),
					('Probius'),
					('Ragnaros'),
					('Raynor'),
					('Rehgar'),
					('Rexxar'),
					('Samuro'),
					('Sgt. Hammer'),
					('Sonya'),
					('Stitches'),
					('Sylvanas'),
					('Tassadar'),
					('The Butcher'),
					('The Lost Vikings'),
					('Thrall'),
					('Tracer'),
					('Tychus'),
					('Tyrael'),
					('Tyrande'),
					('Uther'),
					('Valeera'),
					('Valla'),
					('Varian'),
					('Xul'),
					('Zagara'),
					('Zarya'),
					('Zeratul'),
					('Zul''jin')
				;
				CREATE TABLE translations (
					orig STRING PRIMARY KEY,
					type STRING,
					translation STRING NULL,
					known_bad BOOL,
					INDEX (translation)
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
