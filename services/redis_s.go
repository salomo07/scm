package services

import (
	"context"
	"log"
	"scm/config"

	"github.com/redis/go-redis/v9"
)

type ChannelObj struct {
	NamaChannel string
	Channel     *redis.PubSub
}

var Channels ChannelObj

func SaveValueRedis(key string, value string) {
	var ctx = context.Background()
	opt, _ := redis.ParseURL(config.GetCredRedis())
	client := redis.NewClient(opt)

	client.Set(ctx, key, value, 0)
	print(key + " is saved")
}

func GetValueRedis(key string) (val string, err string) {
	var ctx = context.Background()

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

	Channels.NamaChannel = channelname
	Channels.Channel = pubsub
}
func Publish(channelName string, data any) {
	opt, _ := redis.ParseURL(config.GetCredRedis())
	client := redis.NewClient(opt)
	err := client.Publish(context.Background(), channelName, StructToJson(data)).Err()
	if err != nil {
		log.Fatal(err)
	}
}
func Unsubscribe(channelName string) {
	opt, _ := redis.ParseURL(config.GetCredRedis())
	client := redis.NewClient(opt)
	pubsub := client.Subscribe(context.Background(), channelName)
	err := pubsub.Unsubscribe(context.Background(), channelName)
	if err != nil {
		log.Fatal(err)
	}
	errClose := pubsub.Close()
	if errClose != nil {
		log.Fatal(errClose)
	}
}
