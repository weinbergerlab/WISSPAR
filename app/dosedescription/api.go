package dosedescription

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

type DoseDescription struct {
	ID int
	Name string
}

func (t *DatabaseContext) Get() []DoseDescription {
	results, err := t.client.DoseDescription.Query().All(t.ctx)
	if err != nil {
		return nil
	}

	output := []DoseDescription{}

	for _, result := range results {
		output = append(output, DoseDescription{
			ID:   result.ID,
			Name: result.Name,
		})
	}

	return output
}

func (t *DatabaseContext) Edit(v DoseDescription) error {
	return t.client.DoseDescription.UpdateOneID(v.ID).SetName(v.Name).Exec(t.ctx)
}

func (t *DatabaseContext) Delete(ID int) error {
	return t.client.DoseDescription.DeleteOneID(ID).Exec(t.ctx)
}

func (t *DatabaseContext) Create(n string) error {
	_, err := t.client.DoseDescription.Create().SetName(n).Save(t.ctx)
	return err
}
