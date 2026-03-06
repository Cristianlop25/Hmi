package identification

import "hmi-sonic/internal/connectors"

type StaticService struct{}

func (staticService StaticService) Status() (Identification, error) {
	return Identification{
		Connector: connectors.Connector{
			Name:   "Socket A",
			Type:   connectors.CCS1,
			Status: connectors.Available,
		},
		Rfid:         true,
		Qr:           false,
		QrText:       "",
		ReservedDate: "",
	}, nil
}
