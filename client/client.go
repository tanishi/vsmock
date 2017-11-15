package client

import (
	"context"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/soap"
)

func NewClient(ctx context.Context, u *url.URL) (*govmomi.Client, error) {
	s, err := soap.ParseURL(u.String())

	if err != nil {
		return nil, err
	}

	s.User = u.User
	return govmomi.NewClient(ctx, s, true)
}
