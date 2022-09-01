package manufacturer

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

type Manufacturer struct {
	ID int
	Name string
}

func (t *DatabaseContext) Get() []Manufacturer {
	results, err := t.client.Manufacturer.Query().All(t.ctx)
	if err != nil {
		return nil
	}

	output := []Manufacturer{}

	for _, result := range results {
		output = append(output, Manufacturer{
			ID:   result.ID,
			Name: result.ManufacturerName,
		})
	}

	return output
}

func (t *DatabaseContext) Edit(v Manufacturer) error {
	return t.client.Manufacturer.UpdateOneID(v.ID).SetManufacturerName(v.Name).Exec(t.ctx)
}

func (t *DatabaseContext) Delete(ID int) error {
	return t.client.Manufacturer.DeleteOneID(ID).Exec(t.ctx)
}

func (t *DatabaseContext) Create(n string) error {
	_, err := t.client.Manufacturer.Create().SetManufacturerName(n).Save(t.ctx)
	return err
}
