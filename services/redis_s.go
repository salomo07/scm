package services

import (
	"context"
	"log"
	"scm/config"
	"time"

	"github.com/redis/go-redis/v9"
)

type ChannelObj struct {
	NamaChannel string
	PubSub      *redis.PubSub
}

var channelsMap = make(map[string]*redis.PubSub)

// key string, value string, expired time.Duration
func SaveValueRedis(data ...string) {
	var ctx = context.Background()
	opt, _ := redis.ParseURL(config.GetCredRedis())
	client := redis.NewClient(opt)
	duration := time.Hour * 1
	log.Println(data)
	if len(data) == 3 {
		if data[2] != "" {
			dur, _ := time.ParseDuration(data[2])
			duration = dur
		}
		client.Set(ctx, string(data[0]), data[1], duration)
	} else {
		client.Set(ctx, string(data[0]), data[1], duration)
	}

	print(data[0] + " is saved")
}

func GetValueRedis(key string) (val string, err string) {
	var ctx = context.Background()
	print("\n" + config.GetCredRedis() + "\n")
	opt, error := redis.ParseURL(config.GetCredRedis())
	if error != nil {
		return "", error.Error()
	}
	client := redis.NewClient(opt)
	var res = client.Get(ctx, key).Val()
	return res, ""
}

func SubscribeRedis(channelname string) {
	opt, _ := redis.ParseURL(config.GetCredRedis())
	client := redis.NewClient(opt)
	pubsub := client.Subscribe(context.Background(), channelname)

	channel := pubsub.Channel()
	go func() {
		for msg := range channel {
			log.Println("Terima Pesan: " + msg.Payload)
			println("\nMunkin bisa send to client pakai websocket")
		}
	}()
	println("Redis channel " + channelname + " is active")
	channelsMap[channelname] = pubsub
}
func Publish(channelName string, data any) {
	opt, _ := redis.ParseURL(config.GetCredRedis())
	client := redis.NewClient(opt)
	err := client.Publish(context.Background(), channelName, StructToJson(data)).Err()
	if err != nil {
		log.Fatal(err)
	}
	Unsubscribe(channelName)
}
func Unsubscribe(channelName string) {
	pubsub := channelsMap[channelName]
	err := pubsub.Unsubscribe(context.Background(), channelName)
	if err != nil {
		log.Fatal(err)
	}
	errClose := pubsub.Close()
	if errClose != nil {
		log.Fatal(errClose)
	}
	println("\nUnsubscribe channel " + channelName + "\n")
}
