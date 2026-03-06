package identification

type Service interface {
	Status() (Identification, error)
}
