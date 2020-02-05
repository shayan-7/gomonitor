package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"time"
)

const GAP = 2 // Seconds

func Worker() {
	db, err := gorm.Open(
		"postgres",
		"host=localhost user=postgres dbname=gomono password=postgres",
	)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic("Failed to connect to postgres DB")
	}
	defer db.Close()

	urls := []URL{}
	db.Find(&urls)
	for {
		urls := []URL{}
		db.Where("available = true").Find(&urls)
		for _, url := range urls {
			fmt.Println(url.Address)
			resp, err := http.Get(url.Address)
			if err != nil {
				url.Available = false
				db.Save(&url)
				log.Printf("Unable to connect: %+v\n", err)
				fmt.Printf("\033[96mSHITTTT\n")
				continue
			}
			fmt.Println(resp.Status)
		}
		time.Sleep(GAP * time.Second)
	}
}

//func main() {
//	Worker()
//}
