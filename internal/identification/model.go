package identification

import "hmi-sonic/internal/connectors"

type Identification struct {
	Connector    connectors.Connector
	Rfid         bool
	Qr           bool
	QrText       string
	ReservedDate string
}
