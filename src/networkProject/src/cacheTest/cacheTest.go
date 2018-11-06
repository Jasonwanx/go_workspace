package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func callBack(key string, value interface{}){
	fmt.Println("callBack:", key)
}

func main(){
	c := cache.New(5*time.Minute, 10*time.Minute)

	c.OnEvicted(callBack)

	c.Set("foo", "bar22222222", cache.DefaultExpiration)
	c.Set("ip", "200.10.26.200", cache.DefaultExpiration)

	c.SaveFile("data.txt")

	foo, found := c.Get("foo")
	ip, found := c.Get("ip")
	if found {
		fmt.Println(foo)
		fmt.Println(ip)
	}

	c.Delete("foo")

	//cfg,err := ini.Load("TsiDemoARConfig.ini")
	//if err != nil{
	//	log.Fatal(err)
	//}
	//
	//OrderFileDir := cfg.Section("SearchDir").Key("OrderFileDir")
	//fmt.Println(OrderFileDir)
}
