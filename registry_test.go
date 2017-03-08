package hcr

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/pborman/uuid"
)

func TestGetMoteIDNoAuth(t *testing.T) {
	cl, err := GetHCRClient()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	did := uuid.NewRandom().String()[:10]
	rv, err := cl.GetMoteID(context.Background(), &GetMoteIDParams{
		DeviceId: did,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rv.Status.Okay {
		t.Fatalf("expected bad status, got %v %v", rv, err)
	}
}

func TestGetMoteID(t *testing.T) {
	cl, err := GetHCRClient()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	did := uuid.NewRandom().String()[:10]
	rv, err := cl.GetMoteID(context.Background(), &GetMoteIDParams{
		DeviceId: did,
		Auth: &Auth{
			UserSecret: "764efa883dda1e11db47671c4a3bbd9e",
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !rv.Status.Okay {
		t.Fatalf("expected okay status, got %v,%v", rv, err)
	}

	rv2, err := cl.GetMoteID(context.Background(), &GetMoteIDParams{
		DeviceId: did,
		Auth: &Auth{
			UserSecret: "764efa883dda1e11db47671c4a3bbd9e",
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !rv2.Status.Okay {
		t.Fatalf("expected okay status, got %v,%v", rv2, err)
	}
	if rv2.Mac != rv.Mac || rv2.Id != rv.Id {
		t.Fatalf("expected same ID")
	}
}

func TestCreateDeployment(t *testing.T) {
	cl, err := GetHCRClient()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	name := uuid.NewRandom().String()[:10]
	rv, err := cl.CreateDeployment(context.Background(), &CreateDeploymentParams{
		Auth: &Auth{
			UserSecret: "764efa883dda1e11db47671c4a3bbd9e",
		},
		Name: name,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !rv.Status.Okay {
		t.Fatalf("expected okay status, got %v,%v", rv, err)
	}
}
func TestCreateDuplicateDeployment(t *testing.T) {
	cl, err := GetHCRClient()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	name := uuid.NewRandom().String()[:10]
	rv, err := cl.CreateDeployment(context.Background(), &CreateDeploymentParams{
		Auth: &Auth{
			UserSecret: "764efa883dda1e11db47671c4a3bbd9e",
		},
		Name: name,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !rv.Status.Okay {
		t.Fatalf("expected okay status, got %v,%v", rv, err)
	}
	rv, err = cl.CreateDeployment(context.Background(), &CreateDeploymentParams{
		Auth: &Auth{
			UserSecret: "764efa883dda1e11db47671c4a3bbd9e",
		},
		Name: name,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rv.Status.Okay {
		t.Fatalf("expected error status, got %v,%v", rv, err)
	}
}
func TestBindBadId(t *testing.T) {
	cl, err := GetHCRClient()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	name := uuid.NewRandom().String()[:10]
	rv, err := cl.CreateDeployment(context.Background(), &CreateDeploymentParams{
		Auth: &Auth{
			UserSecret: "764efa883dda1e11db47671c4a3bbd9e",
		},
		Name: name,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !rv.Status.Okay {
		t.Fatalf("expected okay status, got %v,%v", rv, err)
	}
	rv2, err := cl.BindMote(context.Background(), &BindMoteParams{
		Auth: &Auth{
			UserSecret: "764efa883dda1e11db47671c4a3bbd9e",
		},
		Deployment: name,
		Moteid:     0x1234,
	})
	fmt.Printf("got %v %v\n", rv2, err)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rv2.Status.Okay {
		t.Fatalf("expected error status, got %v,%v", rv2, err)
	}
}

func TestBindBadDid(t *testing.T) {
	cl, err := GetHCRClient()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	name := uuid.NewRandom().String()[:10]
	rv, err := cl.CreateDeployment(context.Background(), &CreateDeploymentParams{
		Auth: &Auth{
			UserSecret: "764efa883dda1e11db47671c4a3bbd9e",
		},
		Name: name,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !rv.Status.Okay {
		t.Fatalf("expected okay status, got %v,%v", rv, err)
	}
	rv2, err := cl.BindMote(context.Background(), &BindMoteParams{
		Auth: &Auth{
			UserSecret: "764efa883dda1e11db47671c4a3bbd9e",
		},
		Deployment: "not a real deployment",
		Moteid:     0x100,
	})
	fmt.Printf("got %v %v\n", rv2, err)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rv2.Status.Okay {
		t.Fatalf("expected error status, got %v,%v", rv2, err)
	}
}

func TestMoteInfo(t *testing.T) {
	cl, err := GetHCRClient()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	name := uuid.NewRandom().String()[:10]
	rv, err := cl.CreateDeployment(context.Background(), &CreateDeploymentParams{
		Auth: &Auth{
			UserSecret: "764efa883dda1e11db47671c4a3bbd9e",
		},
		Name: name,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !rv.Status.Okay {
		t.Fatalf("expected okay status, got %v,%v", rv, err)
	}
	did := uuid.NewRandom().String()[:10]
	rvm, err := cl.GetMoteID(context.Background(), &GetMoteIDParams{
		DeviceId: did,
		Auth: &Auth{
			UserSecret: "764efa883dda1e11db47671c4a3bbd9e",
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !rvm.Status.Okay {
		t.Fatalf("expected okay status, got %v,%v", rvm, err)
	}
	rvb, err := cl.BindMote(context.Background(), &BindMoteParams{
		Auth: &Auth{
			UserSecret: "764efa883dda1e11db47671c4a3bbd9e",
		},
		Deployment: name,
		Moteid:     rvm.Id,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !rvb.Status.Okay {
		t.Fatalf("expected okay status, got %v,%v", rvb, err)
	}

	cir, err := cl.CreateInstance(context.Background(), &CreateInstanceParams{
		Auth: &Auth{
			DeploymentSecret: rv.WriteKey,
		},
		DeviceId:   did,
		Repository: "foo",
		Commit:     "bar",
		Motetype:   0x3c,
		Extradata:  "foo",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cir.Status.Okay {
		t.Fatalf("expected okay status, got %v,%v", cir, err)
	}

	rvi, err := cl.MoteInfo(context.Background(), &MoteInfoParams{
		Auth: &Auth{
			DeploymentSecret: rv.ReadKey,
		},
		Moteid: rvm.Id})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !rvi.Status.Okay {
		t.Fatalf("expected okay status, got %v,%v", rvi, err)
	}
	if !bytes.Equal(rvi.AESK, cir.AESK) {
		t.Fatalf("AES key mismatch got\n%064x expected\n%064x", rvi.AESK, cir.AESK)
	}
}
