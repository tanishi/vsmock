package main

import (
	"context"
	"testing"

	"github.com/tanishi/vsmock/helper"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/soap"
)

const (
	GOVC_URL = "GOVC_URL"
	VC_SIM   = "https://user:pass@127.0.0.1:8989/sdk"
)

const (
	VIRTUAL_MACHINE = "VirtualMachine"
	URL             = VC_SIM
)

func TestVMPowerOff(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	u, err := soap.ParseURL(URL)

	helper.LogFatal(err)

	insecure := true

	c, err := govmomi.NewClient(ctx, u, insecure)

	helper.LogFatal(err)

	defer c.Logout(ctx)

	m := view.NewManager(c.Client)

	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{VIRTUAL_MACHINE}, true)

	helper.LogFatal(err)

	defer v.Destroy(ctx)

	f := find.NewFinder(c.Client, true)

	dc, err := f.DefaultDatacenter(ctx)

	helper.LogFatal(err)

	f.SetDatacenter(dc)

	vm, err := f.VirtualMachine(ctx, "/DC0/vm/DC0_H0_VM0")

	helper.LogFatal(err)

	vm.PowerOff(ctx)
}
