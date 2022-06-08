package users

type Service interface {
	GetAll() ([]User, error)
	SaveUser(nome string, sobrenome string, email string, idade string, altura float64) (User, error)
	Update(id int, nome string, sobrenome string, email string, idade string, altura float64) (User, error)
	UpdateLastNameAndAge(id int, sobrenome string, idade string) (User, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]User, error) {
	listUser, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return listUser, nil
}

func (s *service) SaveUser(nome string, sobrenome string, email string, idade string, altura float64) (User, error) {

	lastID, err := s.repository.LastID()
	if err != nil {
		return User{}, err
	}

	lastID++

	user, err := s.repository.SaveUser(lastID, nome, sobrenome, email, idade, altura)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *service) Update(id int, nome string, sobrenome string, email string, idade string, altura float64) (User, error) {
	return s.repository.Update(id, nome, sobrenome, email, idade, altura)
}

func (s *service) UpdateLastNameAndAge(id int, sobrenome string, idade string) (User, error) {
	return s.repository.UpdateLastNameAndAge(id, sobrenome, idade)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
