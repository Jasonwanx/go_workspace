package main

import (
	"bufio"
	"fmt"
	"github.com/patrickmn/go-cache"
<<<<<<< HEAD
	"time"
)

func callBack(key string, value interface{}){
	fmt.Println("callBack:", key)
}
=======
	"log"
	"os"
	"time"
)

>>>>>>> 7b17158ca836f509e1dbb158e8c94c3ca3352510

func main(){
	c := cache.New(5*time.Minute, 10*time.Minute)

	c.OnEvicted(callBack)

	c.Set("foo", "bar22222222", cache.DefaultExpiration)
<<<<<<< HEAD
	c.Set("ip", "200.10.26.200", cache.DefaultExpiration)

	c.SaveFile("data.txt")
=======
	c.Set("price", float32(12.3), cache.DefaultExpiration)

	if value, err := c.IncrementFloat32("price", 0.7); err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(value)
	}

	writer := bufio.NewWriter(os.Stdout)
	c.Save(writer)
>>>>>>> 7b17158ca836f509e1dbb158e8c94c3ca3352510

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
