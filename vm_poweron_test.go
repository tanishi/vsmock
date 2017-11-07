package main

import (
	"context"
	"testing"

	"github.com/tanishi/vsmock/constant"
	"github.com/tanishi/vsmock/helper"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/soap"
)

func TestVMPowerOn(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	u, err := soap.ParseURL(constant.URL)

	helper.LogFatal(err)

	insecure := true

	c, err := govmomi.NewClient(ctx, u, insecure)

	helper.LogFatal(err)

	defer c.Logout(ctx)

	m := view.NewManager(c.Client)

	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{constant.VIRTUAL_MACHINE}, true)

	helper.LogFatal(err)

	defer v.Destroy(ctx)

	f := find.NewFinder(c.Client, true)

	dc, err := f.DefaultDatacenter(ctx)

	helper.LogFatal(err)

	f.SetDatacenter(dc)

	vm, err := f.VirtualMachine(ctx, "/DC0/vm/DC0_H0_VM0")

	helper.LogFatal(err)

	vm.PowerOn(ctx)
}
