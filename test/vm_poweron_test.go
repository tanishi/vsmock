package test

import (
	"context"
	"log"
	"testing"

	"github.com/tanishi/vsmock/client"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/view"
)

func TestVMPowerOn(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := client.NewClient(ctx, URL, USER, PASS)

	if err != nil {
		log.Println(err)
	}

	defer c.Logout(ctx)

	m := view.NewManager(c.Client)

	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)

	if err != nil {
		log.Println(err)
	}

	defer v.Destroy(ctx)

	f := find.NewFinder(c.Client, true)

	dc, err := f.DefaultDatacenter(ctx)

	if err != nil {
		log.Println(err)
	}

	f.SetDatacenter(dc)

	vm, err := f.VirtualMachine(ctx, "/DC0/vm/DC0_H0_VM0")

	if err != nil {
		log.Println(err)
	}

	vm.PowerOn(ctx)

	if s, _ := vm.PowerState(ctx); s != "poweredOn" {
		t.Errorf("%s\n", s)
	}
}
