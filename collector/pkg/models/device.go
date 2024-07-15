package models

type Device struct {
	WWN string `json:"wwn"`

	DeviceName     string `json:"device_name"`
	DeviceUUID     string `json:"device_uuid"`
	DeviceSerialID string `json:"device_serial_id"`
	DeviceLabel    string `json:"device_label"`

	Manufacturer   string `json:"manufacturer"`
	ModelName      string `json:"model_name"`
	InterfaceType  string `json:"interface_type"`
	InterfaceSpeed string `json:"interface_speed"`
	SerialNumber   string `json:"serial_number"`
	Firmware       string `json:"firmware"`
	RotationSpeed  int    `json:"rotational_speed"`
	Capacity       int64  `json:"capacity"`
	FormFactor     string `json:"form_factor"`
	SmartSupport   bool   `json:"smart_support"`
	DeviceProtocol string `json:"device_protocol"` //protocol determines which smart attribute types are available (ATA, NVMe, SCSI)
	DeviceType     string `json:"device_type"`     //device type is used for querying with -d/t flag, should only be used by collector.

	// User provided metadata
	Label  string `json:"label"`
	HostId string `json:"host_id"`
}

type MdRaid struct {
	WWN string `json:"wwn"`

	DeviceName   string `json:"device_name"`
	RaidLevel    string `json:"raid_level"`
	RaidDeviceNr uint8  `json:"raid_device_nr"`
	State        string `json:"state"`
	Size         uint64 `json:"size"`
	UsedDevSize  uint64 `json:"used_dev_size"`

	Name    string   `json:"name"`
	UUID    string   `json:"uuid"`
	Members []string `json:"members"`
}

type DeviceWrapper struct {
	Success bool     `json:"success,omitempty"`
	Errors  []error  `json:"errors,omitempty"`
	Devices []Device `json:"devices"`
	MdRaids []MdRaid `json:"mdraids"`
}
