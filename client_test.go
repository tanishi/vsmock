package client

import (
	"context"
	"testing"
)

const (
	URL     = "localhost:8989"
	USER    = "user"
	PASS    = "pass"
	VERSION = "6.5"
)

func TestClient(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := NewClient(ctx, URL, USER, PASS)

	if err != nil {
		t.Errorf("%v\n", err)
	}

	if v := c.GovmomiClient.Version; v != VERSION {
		t.Errorf("expected: %v, but: %v\n", VERSION, v)
	}
}
