package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"myapp/internal/models"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/uptrace/bun"
)

func NewRepository(Bun *bun.DB) *Repository {
	return &Repository{
		Bun: Bun,
	}
}

type Repository struct {
	Bun *bun.DB
}

func (db *Repository) GetAllCoinsName(ctx context.Context) ([]string, error) {
	data := []models.CoinsDB{}

	err := db.Bun.NewSelect().Model(&data).Where(`deleted_at IS NULL`).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return []string{}, nil
		}
		return nil, fmt.Errorf("[GetAllCoins] Select error: %w", err)
	}
	res := make([]string, len(data))
	for i := range data {
		res[i] = data[i].CoinName
	}

	return res, nil
}

func (db *Repository) AddCoin(ctx context.Context, coinName string) error {
	dataIns := models.CoinsDB{
		CoinName: coinName,
	}

	_, err := db.Bun.NewInsert().On("CONFLICT (coin_name) DO NOTHING").Model(&dataIns).Exec(ctx)
	if err != nil {
		return fmt.Errorf("Repository - AddCoin - db.Bun.NewInsert: %w", err)
	}
	return nil
}

func (db *Repository) GetOneProject(ctx context.Context, id uuid.UUID) (*models.Project, error) {
	project := models.Project{}
	var operatorsId []string

	err := db.Bun.NewSelect().
		Table("projects").
		Column("uuid", "project_name", "project_types.project_type", "operators").
		Join("join project_types").
		JoinOn("projects.project_type = project_types.id").
		Where("uuid = ?", id.String()).
		Scan(ctx, &project.Id, &project.ProjectName, &project.ProjectType, (pq.Array)(&operatorsId))

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Project - GetOneProject - db.Bun.NewSelect: %s", fmt.Sprintf("проекта с id = %s не существует", id.String()))
		}
		log.Println(err)
		return nil, fmt.Errorf("Project - GetOneProject - db.Bun.NewSelect: %w", err)
	}

	return &project, nil
}

func (db *Repository) UpdateProject(ctx context.Context, id uuid.UUID, project models.Project) (*models.Project, error) {

	var operatorsId []string

	err := db.Bun.NewUpdate().
		Model(&project).
		Column("project_name", "project_type").
		Where(`uuid = ?`, id.String()).
		Returning("uuid, operators").
		Scan(ctx, &project.Id, pq.Array(&operatorsId))

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Project - UpdateProject - db.Bun.NewUpdate: %s", fmt.Sprintf("проекта с id = %s не существует", id.String()))
		}
		log.Println(err)
		return nil, fmt.Errorf("Project - UpdateProject - db.Bun.NewUpdate: %w", err)
	}
	return &project, nil
}

func (db *Repository) DeleteProject(ctx context.Context, id uuid.UUID) error {

	project := &models.Project{}
	err := db.Bun.NewDelete().
		Model(project).
		Where(`uuid = ?`, id.String()).
		Returning("uuid").
		Scan(ctx, &id)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("Project - DeleteProject - db.Bun.NewDelete: %s", fmt.Sprintf("проекта с id = %s не существует", id.String()))
		}
		log.Println(err)
		return fmt.Errorf("Project - DeleteProject - db.Bun.NewDelete: %w", err)
	}
	return nil
}

func (db *Repository) AddOperatorToProject(ctx context.Context, projectId uuid.UUID, operatorId uuid.UUID) (*models.Project, error) {

	var operatorsId []string

	project, err := db.GetOneProject(ctx, projectId)
	if err != nil {
		return nil, fmt.Errorf("Project - AddOperatorToProject - db.Bun.GetOneProject: %w", err)
	}

	operatorsId = append(operatorsId, operatorId.String())

	_, err = db.Bun.NewUpdate().
		Model(project).
		Set("operators = ?", pq.Array(operatorsId)).
		Where(`uuid = ?`, projectId.String()).
		Exec(ctx)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Project - AddOperatorToProject - db.Bun.NewUpdate: %w", err)
	}

	return project, nil
}

func (db *Repository) DeleteOperatorFromProject(ctx context.Context, projectId uuid.UUID, operatorId uuid.UUID) (*models.Project, error) {

	var operatorsId []string

	project, err := db.GetOneProject(ctx, projectId)
	if err != nil {
		return nil, fmt.Errorf("Project - DeleteOperatorFromProject - db.Bun.GetOneProject: %w", err)
	}

	exist := false

	if !exist {
		return nil, fmt.Errorf("Project - DeleteOperatorFromProject - db.Bun.NewUpdate: %s", fmt.Sprintf("Данный оператор %s не участвует в проекте %s", operatorId.String(), projectId.String()))
	}

	_, err = db.Bun.NewUpdate().
		Model(project).
		Set("operators = ?", pq.Array(operatorsId)).
		Where(`uuid = ?`, projectId.String()).
		Exec(ctx)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Project - DeleteOperatorFromProject - db.Bun.NewUpdate: %w", err)
	}

	return project, nil
}
