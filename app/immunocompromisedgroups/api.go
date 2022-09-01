package immunocompromisedgroups

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

type ImmunocompromisedGroups struct {
	ID int
	Name string
}

func (t *DatabaseContext) Get() []ImmunocompromisedGroups {
	results, err := t.client.ImmunocompromisedGroups.Query().All(t.ctx)
	if err != nil {
		return nil
	}

	output := []ImmunocompromisedGroups{}

	for _, result := range results {
		output = append(output, ImmunocompromisedGroups{
			ID:   result.ID,
			Name: result.GroupName,
		})
	}

	return output
}

func (t *DatabaseContext) Edit(v ImmunocompromisedGroups) error {
	return t.client.ImmunocompromisedGroups.UpdateOneID(v.ID).SetGroupName(v.Name).Exec(t.ctx)
}

func (t *DatabaseContext) Delete(ID int) error {
	return t.client.ImmunocompromisedGroups.DeleteOneID(ID).Exec(t.ctx)
}

func (t *DatabaseContext) Create(n string) error {
	_, err := t.client.ImmunocompromisedGroups.Create().SetGroupName(n).Save(t.ctx)
	return err
}
