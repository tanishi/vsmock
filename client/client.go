package client

import (
	"context"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/soap"
)

func NewClient(ctx context.Context, u, user, pass string) (*govmomi.Client, error) {
	s, err := soap.ParseURL(u)

	if err != nil {
		return nil, err
	}

	s.User = url.UserPassword(user, pass)

	return govmomi.NewClient(ctx, s, true)
}
