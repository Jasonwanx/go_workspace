package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main(){
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set("foo", "bar22222222", cache.DefaultExpiration)

	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}
}
