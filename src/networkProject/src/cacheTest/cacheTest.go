package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

func main(){
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set("foo", "bar22222222", cache.DefaultExpiration)

	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	cfg,err := ini.Load("TsiDemoARConfig.ini")
	if err != nil{
		log.Fatal(err)
	}

	OrderFileDir := cfg.Section("SearchDir").Key("OrderFileDir")
	fmt.Println(OrderFileDir)
}
