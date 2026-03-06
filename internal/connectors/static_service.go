package connectors

type StaticService struct{}

func (staticService StaticService) List() ([]Connector, error) {
	return []Connector{
		{Name: "Socket A", Type: CCS1, Status: Available},
		{Name: "Socket B", Type: CCS2, Status: Charging},
	}, nil
}
