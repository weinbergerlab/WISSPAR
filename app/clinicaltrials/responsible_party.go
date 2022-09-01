package clinicaltrials

import (
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
)

func (t *ResultsContext) GetResponsibleParties() ([]string, error) {
	return t.client.ClinicalTrial.Query().
		GroupBy(clinicaltrial.FieldResponsibleParty).
		Strings(t.ctx)
}
