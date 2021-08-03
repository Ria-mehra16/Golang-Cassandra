package deleting

type Repository interface {
	DeleteCandy(Candy) (string, error)
}

type Service interface {
	DeleteCandy(Candy) (string, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) DeleteCandy(c Candy) (string, error) {
	id, err := s.r.DeleteCandy(c)
	if err != nil {
		return "", err
	}
	return id, nil
}
