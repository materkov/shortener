package analytics

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Pubsub struct {
	redis *redis.Client
}

func NewPubsub(addr string) *Pubsub {
	return &Pubsub{
		redis: redis.NewClient(&redis.Options{Addr: addr}),
	}
}

func (p *Pubsub) Listen(topics []string, queue string, listener func(topic string, req []byte)) error {
	prefixLen := len(fmt.Sprintf("queue:%s:", queue))

	listenKeys := make([]string, len(topics))
	for idx, topic := range topics {
		listenKeys[idx] = fmt.Sprintf("queue:%s:%s", queue, topic)
	}

	_, err := p.redis.Pipelined(func(p redis.Pipeliner) error {
		for i, topic := range topics {
			p.SAdd(fmt.Sprintf("mq_subscribers:%s", topic), listenKeys[i])
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error subscribing: %s", err)
	}

	for {
		res, err := p.redis.BLPop(time.Second*5, listenKeys...).Result()
		if err == redis.Nil {
			continue
		} else if err != nil {
			return fmt.Errorf("error gettig from queue: %s", err)
		} else if len(res) != 2 {
			return fmt.Errorf("bad redis response")
		}

		listener(res[0][prefixLen:], []byte(res[1]))
	}
}
