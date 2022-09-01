// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"time"

	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrialmetadata"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/dosedescription"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/immunocompromisedgroups"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/manufacturer"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/schedule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/task"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/vaccine"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	clinicaltrialmetadataFields := schema.ClinicalTrialMetadata{}.Fields()
	_ = clinicaltrialmetadataFields
	// clinicaltrialmetadataDescUseCaseCode is the schema descriptor for use_case_code field.
	clinicaltrialmetadataDescUseCaseCode := clinicaltrialmetadataFields[2].Descriptor()
	// clinicaltrialmetadata.DefaultUseCaseCode holds the default value on creation for the use_case_code field.
	clinicaltrialmetadata.DefaultUseCaseCode = clinicaltrialmetadataDescUseCaseCode.Default.(string)
	dosedescriptionFields := schema.DoseDescription{}.Fields()
	_ = dosedescriptionFields
	// dosedescriptionDescName is the schema descriptor for name field.
	dosedescriptionDescName := dosedescriptionFields[0].Descriptor()
	// dosedescription.NameValidator is a validator for the "name" field. It is called by the builders before save.
	dosedescription.NameValidator = dosedescriptionDescName.Validators[0].(func(string) error)
	immunocompromisedgroupsFields := schema.ImmunocompromisedGroups{}.Fields()
	_ = immunocompromisedgroupsFields
	// immunocompromisedgroupsDescGroupName is the schema descriptor for group_name field.
	immunocompromisedgroupsDescGroupName := immunocompromisedgroupsFields[0].Descriptor()
	// immunocompromisedgroups.GroupNameValidator is a validator for the "group_name" field. It is called by the builders before save.
	immunocompromisedgroups.GroupNameValidator = immunocompromisedgroupsDescGroupName.Validators[0].(func(string) error)
	manufacturerFields := schema.Manufacturer{}.Fields()
	_ = manufacturerFields
	// manufacturerDescManufacturerName is the schema descriptor for manufacturer_name field.
	manufacturerDescManufacturerName := manufacturerFields[0].Descriptor()
	// manufacturer.ManufacturerNameValidator is a validator for the "manufacturer_name" field. It is called by the builders before save.
	manufacturer.ManufacturerNameValidator = manufacturerDescManufacturerName.Validators[0].(func(string) error)
	outcomeoverviewFields := schema.OutcomeOverview{}.Fields()
	_ = outcomeoverviewFields
	// outcomeoverviewDescOutcomeOverviewDoseNumber is the schema descriptor for outcome_overview_dose_number field.
	outcomeoverviewDescOutcomeOverviewDoseNumber := outcomeoverviewFields[7].Descriptor()
	// outcomeoverview.DefaultOutcomeOverviewDoseNumber holds the default value on creation for the outcome_overview_dose_number field.
	outcomeoverview.DefaultOutcomeOverviewDoseNumber = outcomeoverviewDescOutcomeOverviewDoseNumber.Default.(int64)
	// outcomeoverviewDescOutcomeOverviewGroupID is the schema descriptor for outcome_overview_group_id field.
	outcomeoverviewDescOutcomeOverviewGroupID := outcomeoverviewFields[11].Descriptor()
	// outcomeoverview.DefaultOutcomeOverviewGroupID holds the default value on creation for the outcome_overview_group_id field.
	outcomeoverview.DefaultOutcomeOverviewGroupID = outcomeoverviewDescOutcomeOverviewGroupID.Default.(string)
	// outcomeoverviewDescOutcomeOverviewRatio is the schema descriptor for outcome_overview_ratio field.
	outcomeoverviewDescOutcomeOverviewRatio := outcomeoverviewFields[12].Descriptor()
	// outcomeoverview.DefaultOutcomeOverviewRatio holds the default value on creation for the outcome_overview_ratio field.
	outcomeoverview.DefaultOutcomeOverviewRatio = outcomeoverviewDescOutcomeOverviewRatio.Default.(string)
	// outcomeoverviewDescOutcomeOverviewMeasureTitle is the schema descriptor for outcome_overview_measure_title field.
	outcomeoverviewDescOutcomeOverviewMeasureTitle := outcomeoverviewFields[13].Descriptor()
	// outcomeoverview.DefaultOutcomeOverviewMeasureTitle holds the default value on creation for the outcome_overview_measure_title field.
	outcomeoverview.DefaultOutcomeOverviewMeasureTitle = outcomeoverviewDescOutcomeOverviewMeasureTitle.Default.(string)
	// outcomeoverviewDescOutcomeOverviewVaccine is the schema descriptor for outcome_overview_vaccine field.
	outcomeoverviewDescOutcomeOverviewVaccine := outcomeoverviewFields[14].Descriptor()
	// outcomeoverview.DefaultOutcomeOverviewVaccine holds the default value on creation for the outcome_overview_vaccine field.
	outcomeoverview.DefaultOutcomeOverviewVaccine = outcomeoverviewDescOutcomeOverviewVaccine.Default.(string)
	// outcomeoverviewDescOutcomeOverviewImmunocompromisedPopulation is the schema descriptor for outcome_overview_immunocompromised_population field.
	outcomeoverviewDescOutcomeOverviewImmunocompromisedPopulation := outcomeoverviewFields[15].Descriptor()
	// outcomeoverview.DefaultOutcomeOverviewImmunocompromisedPopulation holds the default value on creation for the outcome_overview_immunocompromised_population field.
	outcomeoverview.DefaultOutcomeOverviewImmunocompromisedPopulation = outcomeoverviewDescOutcomeOverviewImmunocompromisedPopulation.Default.(string)
	// outcomeoverviewDescOutcomeOverviewManufacturer is the schema descriptor for outcome_overview_manufacturer field.
	outcomeoverviewDescOutcomeOverviewManufacturer := outcomeoverviewFields[16].Descriptor()
	// outcomeoverview.DefaultOutcomeOverviewManufacturer holds the default value on creation for the outcome_overview_manufacturer field.
	outcomeoverview.DefaultOutcomeOverviewManufacturer = outcomeoverviewDescOutcomeOverviewManufacturer.Default.(string)
	// outcomeoverviewDescOutcomeOverviewConfidenceInterval is the schema descriptor for outcome_overview_confidence_interval field.
	outcomeoverviewDescOutcomeOverviewConfidenceInterval := outcomeoverviewFields[17].Descriptor()
	// outcomeoverview.DefaultOutcomeOverviewConfidenceInterval holds the default value on creation for the outcome_overview_confidence_interval field.
	outcomeoverview.DefaultOutcomeOverviewConfidenceInterval = outcomeoverviewDescOutcomeOverviewConfidenceInterval.Default.(string)
	// outcomeoverviewDescOutcomeOverviewPercentResponders is the schema descriptor for outcome_overview_percent_responders field.
	outcomeoverviewDescOutcomeOverviewPercentResponders := outcomeoverviewFields[18].Descriptor()
	// outcomeoverview.DefaultOutcomeOverviewPercentResponders holds the default value on creation for the outcome_overview_percent_responders field.
	outcomeoverview.DefaultOutcomeOverviewPercentResponders = outcomeoverviewDescOutcomeOverviewPercentResponders.Default.(float64)
	// outcomeoverviewDescOutcomeOverviewTimeFrameWeeks is the schema descriptor for outcome_overview_time_frame_weeks field.
	outcomeoverviewDescOutcomeOverviewTimeFrameWeeks := outcomeoverviewFields[19].Descriptor()
	// outcomeoverview.DefaultOutcomeOverviewTimeFrameWeeks holds the default value on creation for the outcome_overview_time_frame_weeks field.
	outcomeoverview.DefaultOutcomeOverviewTimeFrameWeeks = outcomeoverviewDescOutcomeOverviewTimeFrameWeeks.Default.(int64)
	// outcomeoverviewDescOutcomeOverviewDoseDescription is the schema descriptor for outcome_overview_dose_description field.
	outcomeoverviewDescOutcomeOverviewDoseDescription := outcomeoverviewFields[20].Descriptor()
	// outcomeoverview.DefaultOutcomeOverviewDoseDescription holds the default value on creation for the outcome_overview_dose_description field.
	outcomeoverview.DefaultOutcomeOverviewDoseDescription = outcomeoverviewDescOutcomeOverviewDoseDescription.Default.(string)
	// outcomeoverviewDescOutcomeOverviewSchedule is the schema descriptor for outcome_overview_schedule field.
	outcomeoverviewDescOutcomeOverviewSchedule := outcomeoverviewFields[21].Descriptor()
	// outcomeoverview.DefaultOutcomeOverviewSchedule holds the default value on creation for the outcome_overview_schedule field.
	outcomeoverview.DefaultOutcomeOverviewSchedule = outcomeoverviewDescOutcomeOverviewSchedule.Default.(string)
	// outcomeoverviewDescOutcomeOverviewUseCaseCode is the schema descriptor for outcome_overview_use_case_code field.
	outcomeoverviewDescOutcomeOverviewUseCaseCode := outcomeoverviewFields[22].Descriptor()
	// outcomeoverview.DefaultOutcomeOverviewUseCaseCode holds the default value on creation for the outcome_overview_use_case_code field.
	outcomeoverview.DefaultOutcomeOverviewUseCaseCode = outcomeoverviewDescOutcomeOverviewUseCaseCode.Default.(string)
	scheduleFields := schema.Schedule{}.Fields()
	_ = scheduleFields
	// scheduleDescName is the schema descriptor for name field.
	scheduleDescName := scheduleFields[0].Descriptor()
	// schedule.NameValidator is a validator for the "name" field. It is called by the builders before save.
	schedule.NameValidator = scheduleDescName.Validators[0].(func(string) error)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskFields[4].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescUpdatedAt is the schema descriptor for updated_at field.
	taskDescUpdatedAt := taskFields[5].Descriptor()
	// task.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	task.DefaultUpdatedAt = taskDescUpdatedAt.Default.(func() time.Time)
	// task.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	task.UpdateDefaultUpdatedAt = taskDescUpdatedAt.UpdateDefault.(func() time.Time)
	vaccineFields := schema.Vaccine{}.Fields()
	_ = vaccineFields
	// vaccineDescVaccineName is the schema descriptor for vaccine_name field.
	vaccineDescVaccineName := vaccineFields[0].Descriptor()
	// vaccine.VaccineNameValidator is a validator for the "vaccine_name" field. It is called by the builders before save.
	vaccine.VaccineNameValidator = vaccineDescVaccineName.Validators[0].(func(string) error)
}
