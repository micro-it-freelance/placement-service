package placement

type Service interface {
}

type PlacementService struct {
	repo Repository
}

func NewPlacementService(r Repository) Service {
	return &PlacementService{
		repo: r,
	}
}
