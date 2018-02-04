package main

import ("fmt"
	"github.com/go-redis/redis"
	"strconv"
)


var Client *redis.Client

func main() {

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

func setValues(key int, value string) {
	keyToStr := strconv.Itoa(key)
	err := Client.Set(keyToStr, value, 0).Err()
	if err != nil {
		panic(err)
	}
}


func getValues(key int) {
	keyToStr := strconv.Itoa(key)
	val, err := Client.Get(keyToStr).Result()
	if err != nil {
		//panic(err)
	}
	fmt.Println(keyToStr, val)
}

