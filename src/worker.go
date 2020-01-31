package main

import (
	"time"
)

const GAP = 30 // Seconds

func Worker() {
	db, err := gorm.Open(
		"postgres",
		"host=localhost user=postgres dbname=postgres password=postgres",
	)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic("Failed to connect to postgres DB")
	}
	defer db.Close()

	for {
		urls := []Url{}
		db.Find(&urls)
		for _, url := range urls {
			resp, err := http.Get(url.Address)

		}

	}
}
