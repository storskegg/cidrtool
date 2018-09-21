package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"os"
	"time"
)

func pushToDB(cidrList []string) error {
	fmt.Println("Pushing to DB...")
	table := "deleteme"
	pgURI := os.Getenv("RA_PG_URI")

	db, err := sql.Open("postgres", pgURI)
	if err != nil {
		return err
	}

	txn, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := txn.Prepare(pq.CopyIn(table, "ip", "origin", "created_at"))
	if err != nil {
		return err
	}

	origin := "liam"
	for _, ip := range cidrList {
		now := time.Now().Format("2006-01-02 15:04:05.999999")
		_, err2 := stmt.Exec(ip, origin, now)
		if err2 != nil {
			return err2
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	err = db.Close()
	return err
}
