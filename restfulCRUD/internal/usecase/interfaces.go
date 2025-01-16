package usecase

import (
	"context"
	"github.com/google/uuid"
	"myapp/internal/models"
)

type ProjectStrore interface {
	CreateProject(ctx context.Context, p models.Project) error
	GetAllProjects(ctx context.Context) ([]models.Project, error)
	DeleteProject(ctx context.Context, id uuid.UUID) error
	UpdateProject(ctx context.Context, id uuid.UUID, p models.Project) (*models.Project, error)
	GetOneProject(ctx context.Context, id uuid.UUID) (*models.Project, error)

	AddOperatorToProject(ctx context.Context, project_id uuid.UUID, operator_id uuid.UUID) (*models.Project, error)
	DeleteOperatorFromProject(ctx context.Context, project_id uuid.UUID, operator_id uuid.UUID) (*models.Project, error)
}
