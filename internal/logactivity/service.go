package logactivity

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, req CreateLogActivityRequest) (*LogActivity, error) {
	log := &LogActivity{
		ActionType: req.ActionType,
		IPAddress:  req.IPAddress,
		DurationMS: req.DurationMS,
		Metadata:   req.Metadata,
		UserID:     req.UserID,
	}

	err := s.repo.Create(ctx, log)
	if err != nil {
		return nil, err
	}

	return log, nil
}

func (s *Service) FindAll(ctx context.Context) ([]LogActivity, error) {
	return s.repo.FindAll(ctx)
}
