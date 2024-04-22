package main

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

const numberOfRows = 100000

func v3() {
	rows := make([]v2row, 0, numberOfRows)
	fmt.Println("inserting rows")
	for i := 0; i < numberOfRows; i++ {
		id := uuid.New()
		rows = append(rows, v2row{
			key:       &id,
			val:       rand.Int(),
			ttl:       time.Second * time.Duration(rand.Intn(10)),
			createdAt: time.Now(),
		})
	}
	fmt.Println("inserting finished")

	fmt.Println("deleting rows")
	for len(rows) >= 1 {
		count := 0
		for i, r := range rows {
			j := i - count
			expiration := r.createdAt.Add(r.ttl)
			if time.Now().After(expiration) {
				rows[j] = rows[len(rows)-1]
				rows = rows[:len(rows)-1]
				count++
			}
		}
	}
	fmt.Println("process finished")
	fmt.Printf("rows: %v", rows)
}
