package main

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"sync"
	"time"
)

type v4rows map[*uuid.UUID]v4row

type v4row struct {
	val       int
	ttl       time.Duration
	createdAt time.Time
}

func v4() {
	rows := make(v4rows)
	var mu sync.Mutex

	fmt.Println("inserting rows")
	fmt.Println("deleting rows")
	for i := 0; i < million; i++ {
		id := uuid.New()
		ttl := time.Second * time.Duration(rand.Intn(10))
		mu.Lock()
		rows[&id] = v4row{
			val:       rand.Int(),
			ttl:       time.Second * time.Duration(rand.Intn(10)),
			createdAt: time.Now(),
		}
		mu.Unlock()

		go deleteRow(&id, rows, ttl, &mu)
	}
	fmt.Println("inserting finished")
	time.Sleep(10 * time.Second)

	fmt.Println("process finished")
	mu.Lock()
	fmt.Printf("rows: %v", rows)
	mu.Unlock()
}

func deleteRow(id *uuid.UUID, rows v4rows, ttl time.Duration, mu *sync.Mutex) {
	select {
	case <-time.After(ttl):
		mu.Lock()
		delete(rows, id)
		mu.Unlock()
		return
	}
}
