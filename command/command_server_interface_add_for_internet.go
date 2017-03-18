package command

import (
	"fmt"
)

func ServerInterfaceAddForInternet(ctx Context, params *InterfaceAddForInternetServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	diskAPI := client.GetDiskAPI()
	interfaceAPI := client.GetInterfaceAPI()

	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerInterfaceAddForInternet is failed: %s", e)
	}

	if len(p.GetInterfaces()) > 0 {
		return fmt.Errorf("Interface to connect to the Internet(shared-segment) must be the first interface of server")
	}

	// validate connected disk if need DiskEdit
	if !params.WithoutDiskEdit {
		disks := p.GetDisks()
		if len(disks) == 0 {
			return fmt.Errorf("Server haven't any disks. Can't call EditDisk API without server connected disks.")
		}
		// EditDisk API supported?
		res, err := diskAPI.CanEditDisk(disks[0].ID)
		if err != nil {
			return fmt.Errorf("ServerInterfaceAddForInternet is failed: %s", e)
		}
		if !res {
			return fmt.Errorf("Can't call EditDisk API with disk(%d)", disks[0].ID)

		}

	}

	if p.IsUp() {
		return fmt.Errorf("ServerInterfaceAddForInternet is failed: %s", "server is running")
	}

	// call manipurate functions
	nic, err := interfaceAPI.CreateAndConnectToServer(params.Id)
	if err != nil {
		return fmt.Errorf("ServerInterfaceAddForInternet is failed: %s", err)
	}

	_, err = interfaceAPI.ConnectToSharedSegment(nic.ID)
	if err != nil {
		return fmt.Errorf("ServerInterfaceAddForInternet is failed: %s", err)
	}

	if !params.WithoutDiskEdit {
		// disk edit
		editParam := diskAPI.NewCondig()
		_, err := diskAPI.Config(p.GetDisks()[0].ID, editParam)
		if err != nil {
			return fmt.Errorf("ServerInterfaceAddForInternet is failed: %s", err)
		}
	}

	return nil

}
