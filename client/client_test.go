package client

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/autsu/rpcz/loadbalance"
	"github.com/autsu/rpcz/registry"
)

func TestGetServerAddr(t *testing.T) {
	reg, err := registry.NewEtcdClient([]string{"127.0.0.1:2379"})
	if err != nil {
		t.Fatal(err)
	}

	prefix := "/register-servier"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	for i := 0; i < 1000; i++ {
		addr, err := GetServerAddr(ctx, reg, &loadbalance.RoundRobin{}, prefix+"/service1")
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("get server addr: %v\n", addr)
	}
}
