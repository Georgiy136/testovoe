package usecase

import (
	"context"
	"fmt"
	"myapp/internal/models"

	"github.com/google/uuid"
)

type UseCases struct {
	store CoinStrore
}

func NewUsecases(st CoinStrore) *UseCases {
	return &UseCases{
		store: st,
	}
}

func (us *UseCases) AddProject(ctx context.Context, p models.Project) (*models.Project, error) {
	p.Id = uuid.New()
	err := us.store.CreateProject(ctx, p)
	if err != nil {
		return nil, fmt.Errorf("ProjectUseCases - AddProject - us.store.CreateProject: %w", err)
	}
	return &p, nil
}

func (us *UseCases) ListCoins(ctx context.Context) ([]models.CoinsDB, error) {
	p, err := us.store.GetAllCoins(ctx)
	if err != nil {
		return nil, fmt.Errorf("UseCases - ListCoins - us.store.GetAllCoins: %w", err)
	}
	return p, nil
}

func (us *UseCases) DeleteProject(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("ProjectUseCases - DeleteProject - uuid.Parse: %w", err)
	}
	err = us.store.DeleteProject(ctx, uid)
	if err != nil {
		return fmt.Errorf("ProjectUseCases - DeleteProject - us.store.DeleteProject: %w", err)
	}
	return nil
}

func (us *UseCases) UpdateProject(ctx context.Context, id string, p models.Project) (*models.Project, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("ProjectUseCases - UpdateProject - uuid.Parse: %w", err)
	}
	project, err := us.store.UpdateProject(ctx, uid, p)
	if err != nil {
		return nil, fmt.Errorf("ProjectUseCases - UpdateProject - us.store.UpdateProject: %w", err)
	}
	return project, nil
}

func (us *UseCases) GetOneProject(ctx context.Context, id string) (*models.Project, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("ProjectUseCases - GetOneProject - uuid.Parse: %w", err)
	}
	p, err := us.store.GetOneProject(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("ProjectUseCases - GetOneProject - us.store.GetOneProject: %w", err)
	}

	return p, nil
}
