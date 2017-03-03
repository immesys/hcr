package hcr

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMoteID(t *testing.T) {
	cl, err := GetHCRClient()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	rv, err := cl.GetMoteID(context.Background(), &GetMoteIDParams{
		DeviceId: "foobar",
	})
	fmt.Printf("here %v %v\n", rv, err)
}
