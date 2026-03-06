package connectors

type Service interface {
	List() ([]Connector, error)
}
