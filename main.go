package main

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/consul/api"
)

type browserJSON struct {
	URL    string
	Action string
	Expire int
}

func getConsul(kv *api.KV) []byte {
	pair, _, err := kv.Get("foo/json", nil)
	if err != nil {
		fmt.Print("Error getting json value from Consul", err)
	}
	return pair.Value
}

func setConsul(kv *api.KV, jsonData []byte) {
	p := &api.KVPair{Key: "foo/json", Value: []byte(jsonData)}
	_, err := kv.Put(p, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "localhost:8500"
	consulConfig.Scheme = "http"
	consulConfig.Datacenter = "dc1"

	client, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Print("Error creating consul client", err)
	}

	kv := client.KV()

	var consulValue browserJSON

	// =================

	urlToOpen := "https://osu.edu"
	urlAction := "open"
	urlExpire := 0

	browser := browserJSON{
		urlToOpen,
		urlAction,
		urlExpire,
	}

	b, err := json.Marshal(browser)
	if err != nil {
		fmt.Print(err)
	}

	setConsul(kv, b)

	foo := getConsul(kv)
	err = json.Unmarshal(foo, &consulValue)

	if consulValue.URL != "" {
		fmt.Println("URL: ", consulValue.URL)
	} else {
		fmt.Println("consulValue.URL is empty")
	}

	if consulValue.Action != "" {
		fmt.Println("Action: ", consulValue.Action)
	} else {
		fmt.Println("consulValue.Action is empty")
	}

	if consulValue.Expire > 0 {
		fmt.Println("Expire: ", consulValue.Expire)
	} else {
		fmt.Println("consulValue.Expire is zero")
	}

	fmt.Printf("\n---------------\n")

	// =================

	urlToOpen = ""
	urlAction = "reload"
	urlExpire = 0

	browser2 := browserJSON{
		urlToOpen,
		urlAction,
		urlExpire,
	}

	c, err := json.Marshal(browser2)
	if err != nil {
		fmt.Print(err)
	}

	setConsul(kv, c)

	bar := getConsul(kv)
	err = json.Unmarshal(bar, &consulValue)

	if consulValue.URL != "" {
		fmt.Println("URL: ", consulValue.URL)
	} else {
		fmt.Println("consulValue.URL is empty")
	}

	if consulValue.Action != "" {
		fmt.Println("Action: ", consulValue.Action)
	} else {
		fmt.Println("consulValue.Action is empty")
	}

	if consulValue.Expire > 0 {
		fmt.Println("Expire: ", consulValue.Expire)
	} else {
		fmt.Println("consulValue.Expire is zero")
	}

	fmt.Printf("\n---------------\n")

	// =================

	urlToOpen = "https://github.com"
	urlAction = "open"
	urlExpire = 1000

	browser3 := browserJSON{
		urlToOpen,
		urlAction,
		urlExpire,
	}

	d, err := json.Marshal(browser3)
	if err != nil {
		fmt.Print(err)
	}

	setConsul(kv, d)

	baz := getConsul(kv)
	err = json.Unmarshal(baz, &consulValue)

	if consulValue.URL != "" {
		fmt.Println("URL: ", consulValue.URL)
	} else {
		fmt.Println("consulValue.URL is empty")
	}

	if consulValue.Action != "" {
		fmt.Println("Action: ", consulValue.Action)
	} else {
		fmt.Println("consulValue.Action is empty")
	}

	if consulValue.Expire > 0 {
		fmt.Println("Expire: ", consulValue.Expire)
	} else {
		fmt.Println("consulValue.Expire is zero")
	}

}
