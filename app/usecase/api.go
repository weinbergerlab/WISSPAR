package usecase

import (
	"context"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models"
)

type DatabaseContext struct {
	client *models.Client
	ctx    context.Context
}

func NewDatabaseContext(ctx context.Context, client *models.Client) *DatabaseContext {
	return &DatabaseContext{
		client: client,
		ctx:    ctx,
	}
}

type UseCase struct {
	ID int
	Name string
	Description string
	Code string
}

func (t *DatabaseContext) Get() []UseCase {
	results, err := t.client.UseCase.Query().All(t.ctx)
	if err != nil {
		return nil
	}

	output := []UseCase{}

	for _, result := range results {
		output = append(output, UseCase{
			ID:   result.ID,
			Name: result.UseCaseName,
			Description: result.UseCaseDescription,
			Code: result.UseCaseCode,
		})
	}

	return output
}

func (t *DatabaseContext) Edit(u UseCase) error {
	return t.client.UseCase.
		UpdateOneID(u.ID).
		SetUseCaseName(u.Name).
		SetUseCaseDescription(u.Description).
		SetUseCaseCode(u.Code).
		Exec(t.ctx)
}

func (t *DatabaseContext) Delete(ID int) error {
	return t.client.UseCase.DeleteOneID(ID).Exec(t.ctx)
}

func (t *DatabaseContext) Create(n, d ,c string) error {
	_, err := t.client.UseCase.Create().
		SetUseCaseName(n).
		SetUseCaseDescription(d).
		SetUseCaseCode(c).
		Save(t.ctx)
	return err
}
