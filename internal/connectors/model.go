package connectors

type Connector struct {
	Name   string
	Type   ConnectorType
	Status ConnectorStatus
}

type ConnectorType string

const (
	CCS2    ConnectorType = "CCS2"
	CCS1    ConnectorType = "CCS1"
	Schuko  ConnectorType = "Schuko"
	CHAdeMO ConnectorType = "CHAdeMO"
	Type1   ConnectorType = "Type1"
	Type2   ConnectorType = "Type2"
)

const connectorsDir = "/assets/images/connectors/"

var connectorIcons = map[ConnectorType]string{
	CCS2:    "ccs2.svg",
	CCS1:    "ccs1.svg",
	Schuko:  "schuko.svg",
	Type1:   "type1.svg",
	Type2:   "type2.svg",
	CHAdeMO: "chademo.svg",
}

func (connectorType ConnectorType) IconPath() string {
	if icon, ok := connectorIcons[connectorType]; ok {
		return connectorsDir + icon
	}
	return connectorsDir + "default.svg"
}

type ConnectorStatus string

const (
	Available ConnectorStatus = "Available"
	Charging  ConnectorStatus = "Charging"
	Paused    ConnectorStatus = "Paused"
	Disabled  ConnectorStatus = "Disabled"
	Error     ConnectorStatus = "Error"
	Reserved  ConnectorStatus = "Reserved"
	Finished  ConnectorStatus = "Finished"
)

var statusColors = map[ConnectorStatus]string{
	Available: "bg-primary-green",
	Finished:  "bg-primary-green",
	Charging:  "bg-primary-blue",
	Paused:    "bg-primary-grey",
	Disabled:  "bg-primary-grey",
	Error:     "bg-primary-red",
	Reserved:  "bg-primary-cyan",
}

func (status ConnectorStatus) ColorClass() string {
	if color, ok := statusColors[status]; ok {
		return color
	}
	return "bg-primary-green"
}
