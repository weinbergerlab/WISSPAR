package clinicaltrials

import (
	"fmt"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

type OutcomeMeasureCustom struct {
	OutcomeMeasureID                     int    `json:"OutcomeMeasureID,omitempty"`
	OutcomeMeasureType                   string `json:"OutcomeMeasureType"`
	OutcomeMeasureTitle                  string `json:"OutcomeMeasureTitle"`
	OutcomeMeasureDescription            string `json:"OutcomeMeasureDescription"`
	OutcomeMeasurePopulationDescription  string `json:"OutcomeMeasurePopulationDescription"`
	OutcomeMeasureReportingStatus        string `json:"OutcomeMeasureReportingStatus"`
	OutcomeMeasureParamType              string `json:"OutcomeMeasureParamType"`
	OutcomeMeasureDispersionType         string `json:"OutcomeMeasureDispersionType,omitempty"`
	OutcomeMeasureUnitOfMeasure          string `json:"OutcomeMeasureUnitOfMeasure"`
	OutcomeMeasureTimeFrame              string `json:"OutcomeMeasureTimeFrame"`
	OutcomeMeasureAnticipatedPostingDate string `json:"OutcomeMeasureAnticipatedPostingDate"`
	OutcomeMeasureCalculatePct           string `json:"OutcomeMeasureCalculatePct"`
	OutcomeMeasureDenomUnitsSelected     string `json:"OutcomeMeasureDenomUnitsSelected"`
	OutcomeMeasureTypeUnitsAnalyzed      string `json:"OutcomeMeasureTypeUnitsAnalyzed"`
}

func (t *ResultsContext) GetOutcomeMeasures(ID string) ([]OutcomeMeasureCustom, error) {
	outcomeMeasures, err := t.client.OutcomeMeasure.Query().Where(outcomemeasure.HasParentWith(outcomemeasuresmodule.HasParentWith(resultsdefinition.HasParentWith(clinicaltrial.StudyIDEQ(ID))))).All(t.ctx)
	if err != nil {
		return nil, err
	}

	var retValue []OutcomeMeasureCustom

	for _, om := range outcomeMeasures {
		retValue = append(retValue, OutcomeMeasureCustom{
			OutcomeMeasureID:                     om.ID,
			OutcomeMeasureType:                   om.OutcomeMeasureType,
			OutcomeMeasureTitle:                  om.OutcomeMeasureTitle,
			OutcomeMeasureDescription:            om.OutcomeMeasureDescription,
			OutcomeMeasurePopulationDescription:  om.OutcomeMeasurePopulationDescription,
			OutcomeMeasureReportingStatus:        om.OutcomeMeasureReportingStatus,
			OutcomeMeasureParamType:              om.OutcomeMeasureParamType,
			OutcomeMeasureDispersionType:         om.OutcomeMeasureDispersionType,
			OutcomeMeasureUnitOfMeasure:          om.OutcomeMeasureUnitOfMeasure,
			OutcomeMeasureTimeFrame:              om.OutcomeMeasureTimeFrame,
			OutcomeMeasureAnticipatedPostingDate: om.OutcomeMeasureAnticipatedPostingDate,
			OutcomeMeasureCalculatePct:           om.OutcomeMeasureCalculatePct,
			OutcomeMeasureDenomUnitsSelected:     om.OutcomeMeasureDenomUnitsSelected,
			OutcomeMeasureTypeUnitsAnalyzed:      om.OutcomeMeasureTypeUnitsAnalyzed,
		})
	}

	return retValue, nil
}

func (t *ResultsContext) InsertCustomOutcomeMeasures(input []OutcomeMeasureCustom, ID string) error {
	rd, err := t.client.ResultsDefinition.Query().Where(resultsdefinition.HasParentWith(clinicaltrial.StudyIDEQ(ID))).First(t.ctx)
	if err != nil {
		return err
	}

	if rd == nil {
		ct, err := t.client.ClinicalTrial.Query().Where(clinicaltrial.StudyIDEQ(ID)).First(t.ctx)
		if err != nil {
			return err
		}
		rd, err = t.client.ResultsDefinition.Create().SetParent(ct).Save(t.ctx)
		if err != nil {
			return err
		}
	}

	m, err := t.client.OutcomeMeasuresModule.Create().SetParent(rd).Save(t.ctx)
	if err != nil {
		return err
	}

	for _, item := range input {
		if _, err := t.client.OutcomeMeasure.Create().
			SetOutcomeMeasureType(item.OutcomeMeasureType).
			SetOutcomeMeasureTitle(item.OutcomeMeasureType).
			SetOutcomeMeasureDescription(item.OutcomeMeasureType).
			SetOutcomeMeasurePopulationDescription(item.OutcomeMeasureType).
			SetOutcomeMeasureReportingStatus(item.OutcomeMeasureType).
			SetOutcomeMeasureParamType(item.OutcomeMeasureType).
			SetOutcomeMeasureDispersionType(item.OutcomeMeasureType).
			SetOutcomeMeasureUnitOfMeasure(item.OutcomeMeasureType).
			SetOutcomeMeasureTimeFrame(item.OutcomeMeasureType).
			SetOutcomeMeasureAnticipatedPostingDate(item.OutcomeMeasureType).
			SetOutcomeMeasureCalculatePct(item.OutcomeMeasureType).
			SetOutcomeMeasureDenomUnitsSelected(item.OutcomeMeasureType).
			SetOutcomeMeasureTypeUnitsAnalyzed(item.OutcomeMeasureType).
			SetParent(m).Save(t.ctx); err != nil {
			return err
		}
	}
	return nil
}

func (t *ResultsContext) RemoveOutcomeMeasure(ID int) error {
	count, err := t.client.OutcomeOverview.Query().Where(outcomeoverview.HasParentWith(outcomemeasure.IDEQ(ID))).Count(t.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("can't remove outcome measure that has outcome overviews")
	}

	return t.client.OutcomeMeasure.DeleteOneID(ID).Exec(t.ctx)
}

func (t *ResultsContext) InsertCustomOutcomeMeasure(item OutcomeMeasureCustom, ID string) error {
	rd, err := t.client.ResultsDefinition.Query().Where(resultsdefinition.HasParentWith(clinicaltrial.StudyIDEQ(ID))).First(t.ctx)
	if err != nil {
		return err
	}

	if rd == nil {
		ct, err := t.client.ClinicalTrial.Query().Where(clinicaltrial.StudyIDEQ(ID)).First(t.ctx)
		if err != nil {
			return err
		}
		rd, err = t.client.ResultsDefinition.Create().SetParent(ct).Save(t.ctx)
		if err != nil {
			return err
		}
	}

	m, err := t.client.OutcomeMeasuresModule.Create().SetParent(rd).Save(t.ctx)
	if err != nil {
		return err
	}

	_, err = t.client.OutcomeMeasure.Create().
			SetOutcomeMeasureType(item.OutcomeMeasureType).
			SetOutcomeMeasureTitle(item.OutcomeMeasureType).
			SetOutcomeMeasureDescription(item.OutcomeMeasureType).
			SetOutcomeMeasurePopulationDescription(item.OutcomeMeasureType).
			SetOutcomeMeasureReportingStatus(item.OutcomeMeasureType).
			SetOutcomeMeasureParamType(item.OutcomeMeasureType).
			SetOutcomeMeasureDispersionType(item.OutcomeMeasureType).
			SetOutcomeMeasureUnitOfMeasure(item.OutcomeMeasureType).
			SetOutcomeMeasureTimeFrame(item.OutcomeMeasureType).
			SetOutcomeMeasureAnticipatedPostingDate(item.OutcomeMeasureType).
			SetOutcomeMeasureCalculatePct(item.OutcomeMeasureType).
			SetOutcomeMeasureDenomUnitsSelected(item.OutcomeMeasureType).
			SetOutcomeMeasureTypeUnitsAnalyzed(item.OutcomeMeasureType).
			SetParent(m).Save(t.ctx)
	return err
}
