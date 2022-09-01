package schedule

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

type Schedule struct {
	ID int
	Name string
}

func (t *DatabaseContext) Get() []Schedule {
	results, err := t.client.Schedule.Query().All(t.ctx)
	if err != nil {
		return nil
	}

	output := []Schedule{}

	for _, result := range results {
		output = append(output, Schedule{
			ID:   result.ID,
			Name: result.Name,
		})
	}

	return output
}

func (t *DatabaseContext) Edit(v Schedule) error {
	return t.client.Schedule.UpdateOneID(v.ID).SetName(v.Name).Exec(t.ctx)
}

func (t *DatabaseContext) Delete(ID int) error {
	return t.client.Schedule.DeleteOneID(ID).Exec(t.ctx)
}

func (t *DatabaseContext) Create(n string) error {
	_, err := t.client.Schedule.Create().SetName(n).Save(t.ctx)
	return err
}
