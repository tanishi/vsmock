package client

import (
	"context"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/soap"
)

type Client struct {
	URL           *url.URL
	GovmomiClient *govmomi.Client
	user          string
	pass          string
}

func NewClient(ctx context.Context, u, user, pass string) (*Client, error) {
	s, err := soap.ParseURL(u)

	if err != nil {
		return nil, err
	}

	s.User = url.UserPassword(user, pass)

	gc, err := govmomi.NewClient(ctx, s, true)

	if err != nil {
		return nil, err
	}

	c := &Client{
		URL:           s,
		user:          user,
		pass:          pass,
		GovmomiClient: gc,
	}

	return c, nil
}
