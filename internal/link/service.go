package link

type LinkService struct {
	repo *LinkRepo
}

func NewService(repo *LinkRepo) *LinkService {
	return &LinkService{repo: repo}
}
