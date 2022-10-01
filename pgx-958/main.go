package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`create temporary table defer_test (
		id text primary key,
		n int not null, unique (n),
		unique (n) deferrable initially deferred )`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`drop function if exists test_trigger cascade`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`create function test_trigger() returns trigger language plpgsql as $$
	begin
	if new.n = 4 then
		raise exception 'n cant be 4!';
	end if;
	return new;
end$$`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`create constraint trigger test
		after insert or update on defer_test
		deferrable initially deferred
		for each row
		execute function test_trigger()`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`insert into defer_test (id, n) values ('a', 1), ('b', 2), ('c', 3)`)
	if err != nil {
		log.Fatal(err)
	}

	var id string
	err = db.QueryRow(`insert into defer_test (id, n) values ('e', 4) returning id`).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("No error reported.")

	// assert.Error(t, row.Scan(&id))
	// assert.Empty(t, id)

	// var pgErr *pgconn.PgError
	// require.True(t, errors.As(err, &pgErr))
	// assert.Equal(t, "n cant be 4", pgErr.Message)
}
