package redirecter

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

const topicClick = "click"

type Pubsub struct {
	redis *redis.Client
}

func NewPubsub(redisUrl string) *Pubsub {
	return &Pubsub{
		redis: redis.NewClient(&redis.Options{Addr: redisUrl}),
	}
}

func (p *Pubsub) Pub(topic string, msg interface{}) error {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	subscribers, err := p.redis.SMembers(fmt.Sprintf("mq_subscribers:%s", topic)).Result()
	if err != nil {
		return err
	}

	pipe := p.redis.Pipeline()
	for _, sub := range subscribers {
		pipe.LPush(sub, msgBytes)
	}

	if _, err := pipe.Exec(); err != nil {
		return err
	}

	return nil
}
