package main

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func v1() {
	var rows []row
	fmt.Println("inserting rows")
	for i := 0; i < numberOfRows; i++ {
		rows = append(rows, row{
			key:       uuid.New().String(),
			val:       rand.Int(),
			ttl:       time.Second * time.Duration(rand.Intn(10)),
			createdAt: time.Now(),
		})
	}
	fmt.Println("inserting finished")

	fmt.Println("deleting rows")
	for len(rows) > 1 {
		count := 0
		for i, r := range rows {
			j := i - count
			expiration := r.createdAt.Add(r.ttl)
			if time.Now().After(expiration) {
				rows = append(rows[:j], rows[j+1:]...)
				count++
			}
		}
	}
	fmt.Println("process finished")
	fmt.Printf("rows: %v", rows)
}
