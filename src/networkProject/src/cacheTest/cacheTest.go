package main

import (
	"bufio"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/patrickmn/go-cache"
	"log"
	"os"
	"time"
)


func main(){
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set("foo", "bar22222222", cache.DefaultExpiration)
	c.Set("price", float32(12.3), cache.DefaultExpiration)

	if value, err := c.IncrementFloat32("price", 0.7); err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(value)
	}

	writer := bufio.NewWriter(os.Stdout)
	c.Save(writer)

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
