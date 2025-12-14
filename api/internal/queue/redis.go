package queue

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// RedisQueue handles submission queue
type RedisQueue struct {
	Client *redis.Client
}

// NewRedisQueue creates Redis client
func NewRedisQueue(addr, password string) *RedisQueue {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return &RedisQueue{Client: rdb}
}

const SubmissionQueueKey = "oj:queue:submissions"

// EnqueueSubmission pushes submission ID to Redis
func (q *RedisQueue) EnqueueSubmission(ctx context.Context, id int64) error {
	return q.Client.LPush(ctx, SubmissionQueueKey, id).Err()
}

// DequeueSubmission blocks until a submission ID is available
func (q *RedisQueue) DequeueSubmission(ctx context.Context) (int64, error) {
	res, err := q.Client.BRPop(ctx, 0, SubmissionQueueKey).Result()
	if err != nil {
		return 0, err
	}

	var id int64
	_, err = fmt.Sscan(res[1], &id)
	return id, err
}
