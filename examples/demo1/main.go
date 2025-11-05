package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func measureSaturation() float64 {
	return 95.0 + float64(5*rand.Float64())
}

func getDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "oximeter.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS saturation " +
			"(value FLOAT, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP)",
	)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func persistSaturation(db *sql.DB, value float64) {
	_, err := db.Exec("INSERT INTO saturation (value) VALUES(?)", value)
	if err != nil {
		log.Fatal(err)
	}
}

func retrieveAverageSaturation(db *sql.DB, tail time.Duration) float64 {
	row := db.QueryRow(
		"SELECT avg(value) FROM saturation WHERE timestamp >= ?",
		time.Now().UTC().Add(-tail),
	)
	var average float64
	err := row.Scan(&average)
	if err != nil {
		log.Fatal(err)
	}
	return average
}

func main() {
	db := getDatabase()
	go func() {
		for {
			persistSaturation(db, measureSaturation())
			time.Sleep(15 * time.Second)
		}
	}()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tail, err := time.ParseDuration(r.URL.Query()["tail"][0])
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(w, fmt.Sprintf("%f\n", retrieveAverageSaturation(db, tail)))
	})
	http.ListenAndServe(":8080", handler)
}
