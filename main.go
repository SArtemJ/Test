package main

import (
	"github.com/go-redis/redis"
	"strconv"
)

var Client *redis.Client

func main() {

	//Подключаемся к Redis
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	allD := GetAllDevicesFromDB()
	for {
		CreateMetric(allD)
	}

}

//Установить значения по ключу
func setValues(key int, value string) {
	keyToStr := strconv.Itoa(key)
	err := Client.Set(keyToStr, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

//Получить значения по ключу
func getValues(key int) string {
	keyToStr := strconv.Itoa(key)
	val, err := Client.Get(keyToStr).Result()
	if err != nil {
		//panic(err)
	}
	return val
}
