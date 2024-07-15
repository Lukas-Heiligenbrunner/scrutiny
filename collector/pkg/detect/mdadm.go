package detect

import (
	"fmt"
	"github.com/analogj/scrutiny/collector/pkg/models"
	"github.com/prometheus/procfs"
)

func (d *Detect) MdadmScan() ([]models.MdRaid, error) {
	fs, err := procfs.NewFS("/proc")
	if err != nil {
		panic(err)
	}

	mdStats, err := fs.MDStat()
	if err != nil {
		panic(err)
	}

	mdRaids := make([]models.MdRaid, 0)
	for _, md := range mdStats {
		mdRaid := models.MdRaid{DeviceName: md.Name, Members: md.Devices, State: md.ActivityState}

		// todo invoke mdadm --detail /dev/mdx command to get remaining raid infos

		fmt.Printf("MD device: %s, active: %v\n", md.Name, md.ActivityState)
		fmt.Printf("MD device: %v, active: %v\n", md.DisksActive, md.DisksDown)

		mdRaids = append(mdRaids, mdRaid)
	}
	return mdRaids, nil
}
