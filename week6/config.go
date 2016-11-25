package main

import (
	"github.com/hashicorp/consul/api"
	"log"
)

type Config struct {
	Port string
}

var (
	cli api.Client
	Cfg Config
)

func UpdateConfig() {
	cli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err.Error())
	}
	kv := cli.KV()
	value, _, err := kv.Get("TrackServerConfig/port", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	port := string(value.Value)
	Cfg.Port = port
}
