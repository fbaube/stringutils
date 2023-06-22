package stringutils

import (
	"database/sql"
	"fmt"
	// _ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func DBSimple() error {

	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		// log.Fatal(err)
		return err
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		// log.Fatal(err)
		return err
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		// log.Fatal(err)
		return err
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちは世界%03d", i))
		if err != nil {
			// log.Fatal(err)
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		// log.Fatal(err)
		return err
	}

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		// log.Fatal(err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			// log.Fatal(err)
			return err
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		// log.Fatal(err)
		return err
	}

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		// log.Fatal(err)
		return err
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		// log.Fatal(err)
		return err
	}
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		// log.Fatal(err)
		return err
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		// log.Fatal(err)
		return err
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		// log.Fatal(err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			// log.Fatal(err)
			return err
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		// log.Fatal(err)
		return err
	}
	return nil
}
