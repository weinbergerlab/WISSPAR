package vaccine

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

type Vaccine struct {
	ID int
	Name string
}

func (t *DatabaseContext) Get() []Vaccine {
	results, err := t.client.Vaccine.Query().All(t.ctx)
	if err != nil {
		return nil
	}

	output := []Vaccine{}

	for _, result := range results {
		output = append(output, Vaccine{
			ID:   result.ID,
			Name: result.VaccineName,
		})
	}

	return output
}

func (t *DatabaseContext) Edit(v Vaccine) error {
	return t.client.Vaccine.UpdateOneID(v.ID).SetVaccineName(v.Name).Exec(t.ctx)
}

func (t *DatabaseContext) Delete(ID int) error {
	return t.client.Vaccine.DeleteOneID(ID).Exec(t.ctx)
}

func (t *DatabaseContext) Create(n string) error {
	_, err := t.client.Vaccine.Create().SetVaccineName(n).Save(t.ctx)
	return err
}
