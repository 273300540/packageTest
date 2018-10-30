package main

//import (
//	"fmt"
//	"pakF/pakLTwo"
//	"github.com/go-redis/redis"
//)
//
//func main() {
//	fmt.Print(pakLTwo.Hello())
//	client := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//	pingResult, err := client.Ping().Result();
//	if err != nil {
//		fmt.Print(err);
//		return;
//	}
//	fmt.Print(pingResult)
//}
func Hello2() string {
	return "hello2"
}
