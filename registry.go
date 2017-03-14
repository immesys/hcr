package hcr

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
)

//go:generate protoc hreg.proto --go_out=plugins=grpc:.

const registryAddress = "hcr.steelcode.com:5410"

func GetHCRClient() (HamiltonRegistryClient, error) {
	conn, err := grpc.Dial(registryAddress, grpc.WithInsecure(), grpc.WithTimeout(5*time.Second), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	client := NewHamiltonRegistryClient(conn)
	return client, nil
}

var cache map[int]cacheEntry
var cachemu sync.Mutex

type cacheEntry struct {
	mir    *MoteInfoResponse
	expiry time.Time
}

func (c *cacheEntry) Expired() bool {
	return time.Now().After(c.expiry)
}

func init() {
	cache = make(map[int]cacheEntry)
}

const expiryTime = 1 * time.Hour

func GetMoteInfo(ctx context.Context, moteid int, deploymentSecret string) (*MoteInfoResponse, error) {
	cachemu.Lock()
	ce, ok := cache[moteid]
	cachemu.Unlock()
	if ok && !ce.Expired() {
		return ce.mir, nil
	}
	conn, err := grpc.Dial(registryAddress, grpc.WithInsecure(), grpc.WithTimeout(5*time.Second), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cl := NewHamiltonRegistryClient(conn)
	rvi, err := cl.MoteInfo(ctx, &MoteInfoParams{
		Auth: &Auth{
			DeploymentSecret: deploymentSecret,
		},
		Moteid: uint32(moteid),
	})
	if err != nil {
		return nil, err
	}
	if !rvi.Status.Okay {
		return nil, fmt.Errorf("could not obtain mote info: %v", rvi.Status.Message)
	}
	cachemu.Lock()
	cache[moteid] = cacheEntry{rvi, time.Now().Add(expiryTime)}
	cachemu.Unlock()
	return rvi, nil
}
