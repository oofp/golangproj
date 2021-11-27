package graph

import (
	"fmt"
	"sync"
	"testing"
)

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

func TestGetNextId(t *testing.T) {
	id1 := GetNextId()
	id2 := GetNextId()
	if id1 == id2 {
		t.Fatal("ids are identical")
	}
}

func TestGetNextIdAsync(t *testing.T) {
	var ids sync.Map
	flFailed := false

	for i := 0; i < 100; i++ {
		go func() {
			id := GetNextId()
			_, flLoaded := ids.LoadOrStore(id, "")
			if flLoaded {
				flFailed = true
			}
		}()
	}
	if flFailed {
		t.Fatal("duplicated ids found")
	}

	latestId := GetNextId()
	fmt.Println("LatestID:", latestId)
}
