package metadata

import (
	"context"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrialmetadata"
)

type Metadata struct {
	TagName string `json:"tag_name"`
	TagValue string `json:"tag_value"`
	TagUseCase string `json:"tag_use_case"`
}

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

func (t *DatabaseContext) Get(ID string, useCase string) []Metadata {
	output := []Metadata{}

	metadata, err := t.client.ClinicalTrial.Query().Where(clinicaltrial.StudyIDEQ(ID)).QueryMetadata().Where(clinicaltrialmetadata.UseCaseCodeEQ(useCase)).All(t.ctx)
	if err != nil {
		return output
	}

	if len(metadata) == 0 {
		return output
	}

	for _, result := range metadata {
		output = append(output, Metadata{
			TagName:   result.TagName,
			TagValue: result.TagValue,
			TagUseCase: result.UseCaseCode,
		})
	}

	return output
}
