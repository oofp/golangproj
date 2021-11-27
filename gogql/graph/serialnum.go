package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"strconv"
	"sync/atomic"
)

var last_id int64

func GetNextId() string {
	atomic.AddInt64(&last_id, 1)
	new_id := strconv.FormatInt(last_id, 10)
	return new_id
}
