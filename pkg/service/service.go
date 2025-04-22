package service

type VshepService interface {
	SendRequest(request []byte) (int, []byte, error)
}

type Service struct {
	VshepService
}

func NewService(vshepURL string) *Service {
	return &Service{
		NewVshepServiceImpl(vshepURL),
	}
}
