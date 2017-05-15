package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/ftp"
)

func ServerIsoInsert(ctx Context, params *IsoInsertServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerIsoInsert is failed: %s", e)
	}

	if p.Instance.CDROM != nil && p.Instance.CDROM.ID != sacloud.EmptyID {
		fmt.Fprint(GlobalOption.Err, "ISOImage is already inserted to server\n")
		return nil
	}

	imageID := params.IsoImageId
	if imageID == sacloud.EmptyID {

		//validate
		isoParams := &CreateISOImageParam{
			Tags:        params.Tags,
			IconId:      params.IconId,
			Size:        params.Size,
			Name:        params.Name,
			Description: params.Description,
			IsoFile:     params.IsoFile,
		}
		if errs := isoParams.Validate(); len(errs) > 0 {
			return flattenErrors(errs)
		}

		// Upload iso image
		api := client.GetCDROMAPI()
		iso := api.New()

		// set params
		iso.SetTags(params.Tags)
		iso.SetIconByID(params.IconId)
		iso.SetSizeGB(params.Size)
		iso.SetName(params.Name)
		iso.SetDescription(params.Description)

		// call Create(id)
		res, ftpServer, err := api.Create(iso)
		if err != nil {
			return fmt.Errorf("ServerIsoInsert is failed: %s", err)
		}

		// upload
		ftpsClient := ftp.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
		err = ftpsClient.Upload(params.GetIsoFile())
		if err != nil {
			return fmt.Errorf("ServerIsoInsert is failed: %s", err)
		}

		// close FTP
		_, err = api.CloseFTP(res.ID)
		if err != nil {
			return fmt.Errorf("ServerIsoInsert is failed: %s", err)
		}

		imageID = res.ID
	}

	// call manipurate functions
	_, err := api.InsertCDROM(params.Id, imageID)
	if err != nil {
		return fmt.Errorf("ServerIsoInsert is failed: %s", err)
	}

	return nil
	// return ctx.GetOutput().Print(res)

}
