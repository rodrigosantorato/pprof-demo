package main

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type v2row struct {
	key       *uuid.UUID
	val       int
	ttl       time.Duration
	createdAt time.Time
}

func v2() {
	var rows []v2row
	fmt.Println("inserting rows")
	for i := 0; i < 1000000; i++ {
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
