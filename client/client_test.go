package client

import (
	"context"
	"net/url"
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

	u, err := url.Parse(URL)

	if err != nil {
		t.Errorf("%v\n", err)
	}

	u.User = url.UserPassword(USER, PASS)

	c, err := NewClient(ctx, u)

	if err != nil {
		t.Errorf("%v\n", err)
	}

	if v := c.Client.Client.Version; v != VERSION {
		t.Errorf("expected: %v, but: %v\n", VERSION, v)
	}
}
