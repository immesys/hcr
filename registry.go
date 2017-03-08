package hcr

import (
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
