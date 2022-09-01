// Code generated (@generated) by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdverseEventsModuleColumns holds the columns for the "adverse_events_module" table.
	AdverseEventsModuleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "events_frequency_threshold", Type: field.TypeString},
		{Name: "events_time_frame", Type: field.TypeString},
		{Name: "events_description", Type: field.TypeString},
		{Name: "results_definition_adverse_events_module", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// AdverseEventsModuleTable holds the schema information for the "adverse_events_module" table.
	AdverseEventsModuleTable = &schema.Table{
		Name:       "adverse_events_module",
		Columns:    AdverseEventsModuleColumns,
		PrimaryKey: []*schema.Column{AdverseEventsModuleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "adverse_events_module_results_definition_adverse_events_module",
				Columns:    []*schema.Column{AdverseEventsModuleColumns[4]},
				RefColumns: []*schema.Column{ResultsDefinitionColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BaselineCategoryColumns holds the columns for the "baseline_category" table.
	BaselineCategoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "baseline_category_title", Type: field.TypeString},
		{Name: "baseline_class_baseline_category_list", Type: field.TypeInt, Nullable: true},
	}
	// BaselineCategoryTable holds the schema information for the "baseline_category" table.
	BaselineCategoryTable = &schema.Table{
		Name:       "baseline_category",
		Columns:    BaselineCategoryColumns,
		PrimaryKey: []*schema.Column{BaselineCategoryColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "baseline_category_baseline_class_baseline_category_list",
				Columns:    []*schema.Column{BaselineCategoryColumns[2]},
				RefColumns: []*schema.Column{BaselineClassColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BaselineCharacteristicsModuleColumns holds the columns for the "baseline_characteristics_module" table.
	BaselineCharacteristicsModuleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "baseline_population_description", Type: field.TypeString},
		{Name: "baseline_type_units_analyzed", Type: field.TypeString},
		{Name: "results_definition_baseline_characteristics_module", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// BaselineCharacteristicsModuleTable holds the schema information for the "baseline_characteristics_module" table.
	BaselineCharacteristicsModuleTable = &schema.Table{
		Name:       "baseline_characteristics_module",
		Columns:    BaselineCharacteristicsModuleColumns,
		PrimaryKey: []*schema.Column{BaselineCharacteristicsModuleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "baseline_characteristics_module_results_definition_baseline_characteristics_module",
				Columns:    []*schema.Column{BaselineCharacteristicsModuleColumns[3]},
				RefColumns: []*schema.Column{ResultsDefinitionColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BaselineClassColumns holds the columns for the "baseline_class" table.
	BaselineClassColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "baseline_class_title", Type: field.TypeString},
		{Name: "baseline_measure_baseline_class_list", Type: field.TypeInt, Nullable: true},
	}
	// BaselineClassTable holds the schema information for the "baseline_class" table.
	BaselineClassTable = &schema.Table{
		Name:       "baseline_class",
		Columns:    BaselineClassColumns,
		PrimaryKey: []*schema.Column{BaselineClassColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "baseline_class_baseline_measure_baseline_class_list",
				Columns:    []*schema.Column{BaselineClassColumns[2]},
				RefColumns: []*schema.Column{BaselineMeasureColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BaselineClassDenomColumns holds the columns for the "baseline_class_denom" table.
	BaselineClassDenomColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "baseline_class_denom_units", Type: field.TypeString},
		{Name: "baseline_class_baseline_class_denom_list", Type: field.TypeInt, Nullable: true},
	}
	// BaselineClassDenomTable holds the schema information for the "baseline_class_denom" table.
	BaselineClassDenomTable = &schema.Table{
		Name:       "baseline_class_denom",
		Columns:    BaselineClassDenomColumns,
		PrimaryKey: []*schema.Column{BaselineClassDenomColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "baseline_class_denom_baseline_class_baseline_class_denom_list",
				Columns:    []*schema.Column{BaselineClassDenomColumns[2]},
				RefColumns: []*schema.Column{BaselineClassColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BaselineClassDenomCountColumns holds the columns for the "baseline_class_denom_count" table.
	BaselineClassDenomCountColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "baseline_class_denom_count_group_id", Type: field.TypeString},
		{Name: "baseline_class_denom_count_value", Type: field.TypeString},
		{Name: "baseline_class_denom_baseline_class_denom_count_list", Type: field.TypeInt, Nullable: true},
	}
	// BaselineClassDenomCountTable holds the schema information for the "baseline_class_denom_count" table.
	BaselineClassDenomCountTable = &schema.Table{
		Name:       "baseline_class_denom_count",
		Columns:    BaselineClassDenomCountColumns,
		PrimaryKey: []*schema.Column{BaselineClassDenomCountColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "baseline_class_denom_count_baseline_class_denom_baseline_class_denom_count_list",
				Columns:    []*schema.Column{BaselineClassDenomCountColumns[3]},
				RefColumns: []*schema.Column{BaselineClassDenomColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BaselineDenomColumns holds the columns for the "baseline_denom" table.
	BaselineDenomColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "baseline_denom_units", Type: field.TypeString},
		{Name: "baseline_characteristics_module_baseline_denom_list", Type: field.TypeInt, Nullable: true},
	}
	// BaselineDenomTable holds the schema information for the "baseline_denom" table.
	BaselineDenomTable = &schema.Table{
		Name:       "baseline_denom",
		Columns:    BaselineDenomColumns,
		PrimaryKey: []*schema.Column{BaselineDenomColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "baseline_denom_baseline_characteristics_module_baseline_denom_list",
				Columns:    []*schema.Column{BaselineDenomColumns[2]},
				RefColumns: []*schema.Column{BaselineCharacteristicsModuleColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BaselineDenomCountColumns holds the columns for the "baseline_denom_count" table.
	BaselineDenomCountColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "baseline_denom_count_group_id", Type: field.TypeString},
		{Name: "baseline_denom_count_value", Type: field.TypeString},
		{Name: "baseline_denom_baseline_denom_count_list", Type: field.TypeInt, Nullable: true},
	}
	// BaselineDenomCountTable holds the schema information for the "baseline_denom_count" table.
	BaselineDenomCountTable = &schema.Table{
		Name:       "baseline_denom_count",
		Columns:    BaselineDenomCountColumns,
		PrimaryKey: []*schema.Column{BaselineDenomCountColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "baseline_denom_count_baseline_denom_baseline_denom_count_list",
				Columns:    []*schema.Column{BaselineDenomCountColumns[3]},
				RefColumns: []*schema.Column{BaselineDenomColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BaselineGroupColumns holds the columns for the "baseline_group" table.
	BaselineGroupColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "baseline_group_id", Type: field.TypeString},
		{Name: "baseline_group_title", Type: field.TypeString},
		{Name: "baseline_group_description", Type: field.TypeString},
		{Name: "baseline_characteristics_module_baseline_group_list", Type: field.TypeInt, Nullable: true},
	}
	// BaselineGroupTable holds the schema information for the "baseline_group" table.
	BaselineGroupTable = &schema.Table{
		Name:       "baseline_group",
		Columns:    BaselineGroupColumns,
		PrimaryKey: []*schema.Column{BaselineGroupColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "baseline_group_baseline_characteristics_module_baseline_group_list",
				Columns:    []*schema.Column{BaselineGroupColumns[4]},
				RefColumns: []*schema.Column{BaselineCharacteristicsModuleColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BaselineMeasureColumns holds the columns for the "baseline_measure" table.
	BaselineMeasureColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "baseline_measure_title", Type: field.TypeString},
		{Name: "baseline_measure_description", Type: field.TypeString},
		{Name: "baseline_measure_population_description", Type: field.TypeString},
		{Name: "baseline_measure_param_type", Type: field.TypeString},
		{Name: "baseline_measure_dispersion_type", Type: field.TypeString},
		{Name: "baseline_measure_unit_of_measure", Type: field.TypeString},
		{Name: "baseline_measure_calculate_pct", Type: field.TypeString},
		{Name: "baseline_measure_denom_units_selected", Type: field.TypeString},
		{Name: "baseline_characteristics_module_baseline_measure_list", Type: field.TypeInt, Nullable: true},
	}
	// BaselineMeasureTable holds the schema information for the "baseline_measure" table.
	BaselineMeasureTable = &schema.Table{
		Name:       "baseline_measure",
		Columns:    BaselineMeasureColumns,
		PrimaryKey: []*schema.Column{BaselineMeasureColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "baseline_measure_baseline_characteristics_module_baseline_measure_list",
				Columns:    []*schema.Column{BaselineMeasureColumns[9]},
				RefColumns: []*schema.Column{BaselineCharacteristicsModuleColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BaselineMeasureDenomColumns holds the columns for the "baseline_measure_denom" table.
	BaselineMeasureDenomColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "baseline_measure_denom_units", Type: field.TypeString},
		{Name: "baseline_measure_baseline_measure_denom_list", Type: field.TypeInt, Nullable: true},
	}
	// BaselineMeasureDenomTable holds the schema information for the "baseline_measure_denom" table.
	BaselineMeasureDenomTable = &schema.Table{
		Name:       "baseline_measure_denom",
		Columns:    BaselineMeasureDenomColumns,
		PrimaryKey: []*schema.Column{BaselineMeasureDenomColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "baseline_measure_denom_baseline_measure_baseline_measure_denom_list",
				Columns:    []*schema.Column{BaselineMeasureDenomColumns[2]},
				RefColumns: []*schema.Column{BaselineMeasureColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BaselineMeasureDenomCountColumns holds the columns for the "baseline_measure_denom_count" table.
	BaselineMeasureDenomCountColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "baseline_measure_denom_count_group_id", Type: field.TypeString},
		{Name: "baseline_measure_denom_count_value", Type: field.TypeString},
		{Name: "baseline_measure_denom_baseline_measure_denom_count_list", Type: field.TypeInt, Nullable: true},
	}
	// BaselineMeasureDenomCountTable holds the schema information for the "baseline_measure_denom_count" table.
	BaselineMeasureDenomCountTable = &schema.Table{
		Name:       "baseline_measure_denom_count",
		Columns:    BaselineMeasureDenomCountColumns,
		PrimaryKey: []*schema.Column{BaselineMeasureDenomCountColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "baseline_measure_denom_count_baseline_measure_denom_baseline_measure_denom_count_list",
				Columns:    []*schema.Column{BaselineMeasureDenomCountColumns[3]},
				RefColumns: []*schema.Column{BaselineMeasureDenomColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BaselineMeasurementColumns holds the columns for the "baseline_measurement" table.
	BaselineMeasurementColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "baseline_measurement_group_id", Type: field.TypeString},
		{Name: "baseline_measurement_value", Type: field.TypeString},
		{Name: "baseline_measurement_spread", Type: field.TypeString},
		{Name: "baseline_measurement_lower_limit", Type: field.TypeString},
		{Name: "baseline_measurement_upper_limit", Type: field.TypeString},
		{Name: "baseline_measurement_comment", Type: field.TypeString},
		{Name: "baseline_category_baseline_measurement_list", Type: field.TypeInt, Nullable: true},
	}
	// BaselineMeasurementTable holds the schema information for the "baseline_measurement" table.
	BaselineMeasurementTable = &schema.Table{
		Name:       "baseline_measurement",
		Columns:    BaselineMeasurementColumns,
		PrimaryKey: []*schema.Column{BaselineMeasurementColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "baseline_measurement_baseline_category_baseline_measurement_list",
				Columns:    []*schema.Column{BaselineMeasurementColumns[7]},
				RefColumns: []*schema.Column{BaselineCategoryColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CertainAgreementColumns holds the columns for the "certain_agreement" table.
	CertainAgreementColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "agreement_pi_sponsor_employee", Type: field.TypeString},
		{Name: "agreement_restriction_type", Type: field.TypeString},
		{Name: "agreement_restrictive_agreement", Type: field.TypeString},
		{Name: "agreement_other_details", Type: field.TypeString},
		{Name: "certain_agreement_parent", Type: field.TypeInt},
	}
	// CertainAgreementTable holds the schema information for the "certain_agreement" table.
	CertainAgreementTable = &schema.Table{
		Name:       "certain_agreement",
		Columns:    CertainAgreementColumns,
		PrimaryKey: []*schema.Column{CertainAgreementColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "certain_agreement_more_info_module_parent",
				Columns:    []*schema.Column{CertainAgreementColumns[5]},
				RefColumns: []*schema.Column{MoreInfoModuleColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ClinicalTrialColumns holds the columns for the "clinical_trial" table.
	ClinicalTrialColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "study_name", Type: field.TypeString},
		{Name: "sponsor", Type: field.TypeString},
		{Name: "responsible_party", Type: field.TypeString},
		{Name: "study_type", Type: field.TypeString},
		{Name: "phase", Type: field.TypeString},
		{Name: "actual_enrollment", Type: field.TypeString},
		{Name: "allocation", Type: field.TypeString},
		{Name: "intervention_model", Type: field.TypeString},
		{Name: "masking", Type: field.TypeString},
		{Name: "primary_purpose", Type: field.TypeString},
		{Name: "official_title", Type: field.TypeString},
		{Name: "actual_study_start_date", Type: field.TypeTime},
		{Name: "actual_primary_completion_date", Type: field.TypeTime},
		{Name: "actual_study_completion_date", Type: field.TypeTime},
		{Name: "study_id", Type: field.TypeString},
	}
	// ClinicalTrialTable holds the schema information for the "clinical_trial" table.
	ClinicalTrialTable = &schema.Table{
		Name:       "clinical_trial",
		Columns:    ClinicalTrialColumns,
		PrimaryKey: []*schema.Column{ClinicalTrialColumns[0]},
	}
	// ClinicalTrialMetadataColumns holds the columns for the "clinical_trial_metadata" table.
	ClinicalTrialMetadataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tag_name", Type: field.TypeString},
		{Name: "tag_value", Type: field.TypeString},
		{Name: "use_case_code", Type: field.TypeString, Default: "default"},
		{Name: "clinical_trial_metadata", Type: field.TypeInt, Nullable: true},
	}
	// ClinicalTrialMetadataTable holds the schema information for the "clinical_trial_metadata" table.
	ClinicalTrialMetadataTable = &schema.Table{
		Name:       "clinical_trial_metadata",
		Columns:    ClinicalTrialMetadataColumns,
		PrimaryKey: []*schema.Column{ClinicalTrialMetadataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "clinical_trial_metadata_clinical_trial_metadata",
				Columns:    []*schema.Column{ClinicalTrialMetadataColumns[4]},
				RefColumns: []*schema.Column{ClinicalTrialColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DoseDescriptionColumns holds the columns for the "dose_description" table.
	DoseDescriptionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// DoseDescriptionTable holds the schema information for the "dose_description" table.
	DoseDescriptionTable = &schema.Table{
		Name:       "dose_description",
		Columns:    DoseDescriptionColumns,
		PrimaryKey: []*schema.Column{DoseDescriptionColumns[0]},
	}
	// EventGroupColumns holds the columns for the "event_group" table.
	EventGroupColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "event_group_id", Type: field.TypeString},
		{Name: "event_group_title", Type: field.TypeString},
		{Name: "event_group_description", Type: field.TypeString},
		{Name: "event_group_deaths_num_affected", Type: field.TypeString},
		{Name: "event_group_deaths_num_at_risk", Type: field.TypeString},
		{Name: "event_group_serious_num_affected", Type: field.TypeString},
		{Name: "event_group_serious_num_at_risk", Type: field.TypeString},
		{Name: "event_group_other_num_affected", Type: field.TypeString},
		{Name: "event_group_other_num_at_risk", Type: field.TypeString},
		{Name: "adverse_events_module_event_group_list", Type: field.TypeInt, Nullable: true},
	}
	// EventGroupTable holds the schema information for the "event_group" table.
	EventGroupTable = &schema.Table{
		Name:       "event_group",
		Columns:    EventGroupColumns,
		PrimaryKey: []*schema.Column{EventGroupColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "event_group_adverse_events_module_event_group_list",
				Columns:    []*schema.Column{EventGroupColumns[10]},
				RefColumns: []*schema.Column{AdverseEventsModuleColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FlowAchievementColumns holds the columns for the "flow_achievement" table.
	FlowAchievementColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "flow_achievement_group_id", Type: field.TypeString},
		{Name: "flow_achievement_comment", Type: field.TypeString},
		{Name: "flow_achievement_num_subjects", Type: field.TypeString},
		{Name: "flow_achievement_num_units", Type: field.TypeString},
		{Name: "flow_milestone_flow_achievement_list", Type: field.TypeInt, Nullable: true},
	}
	// FlowAchievementTable holds the schema information for the "flow_achievement" table.
	FlowAchievementTable = &schema.Table{
		Name:       "flow_achievement",
		Columns:    FlowAchievementColumns,
		PrimaryKey: []*schema.Column{FlowAchievementColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "flow_achievement_flow_milestone_flow_achievement_list",
				Columns:    []*schema.Column{FlowAchievementColumns[5]},
				RefColumns: []*schema.Column{FlowMilestoneColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FlowDropWithdrawColumns holds the columns for the "flow_drop_withdraw" table.
	FlowDropWithdrawColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "flow_drop_withdraw_type", Type: field.TypeString},
		{Name: "flow_drop_withdraw_comment", Type: field.TypeString},
		{Name: "flow_period_flow_drop_withdraw_list", Type: field.TypeInt, Nullable: true},
	}
	// FlowDropWithdrawTable holds the schema information for the "flow_drop_withdraw" table.
	FlowDropWithdrawTable = &schema.Table{
		Name:       "flow_drop_withdraw",
		Columns:    FlowDropWithdrawColumns,
		PrimaryKey: []*schema.Column{FlowDropWithdrawColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "flow_drop_withdraw_flow_period_flow_drop_withdraw_list",
				Columns:    []*schema.Column{FlowDropWithdrawColumns[3]},
				RefColumns: []*schema.Column{FlowPeriodColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FlowGroupColumns holds the columns for the "flow_group" table.
	FlowGroupColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "flow_group_id", Type: field.TypeString},
		{Name: "flow_group_title", Type: field.TypeString},
		{Name: "flow_group_description", Type: field.TypeString},
		{Name: "participant_flow_module_flow_group_list", Type: field.TypeInt, Nullable: true},
	}
	// FlowGroupTable holds the schema information for the "flow_group" table.
	FlowGroupTable = &schema.Table{
		Name:       "flow_group",
		Columns:    FlowGroupColumns,
		PrimaryKey: []*schema.Column{FlowGroupColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "flow_group_participant_flow_module_flow_group_list",
				Columns:    []*schema.Column{FlowGroupColumns[4]},
				RefColumns: []*schema.Column{ParticipantFlowModuleColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FlowMilestoneColumns holds the columns for the "flow_milestone" table.
	FlowMilestoneColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "flow_milestone_type", Type: field.TypeString},
		{Name: "flow_milestone_comment", Type: field.TypeString},
		{Name: "flow_period_flow_milestone_list", Type: field.TypeInt, Nullable: true},
	}
	// FlowMilestoneTable holds the schema information for the "flow_milestone" table.
	FlowMilestoneTable = &schema.Table{
		Name:       "flow_milestone",
		Columns:    FlowMilestoneColumns,
		PrimaryKey: []*schema.Column{FlowMilestoneColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "flow_milestone_flow_period_flow_milestone_list",
				Columns:    []*schema.Column{FlowMilestoneColumns[3]},
				RefColumns: []*schema.Column{FlowPeriodColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FlowPeriodColumns holds the columns for the "flow_period" table.
	FlowPeriodColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "flow_period_title", Type: field.TypeString},
		{Name: "participant_flow_module_flow_period_list", Type: field.TypeInt, Nullable: true},
	}
	// FlowPeriodTable holds the schema information for the "flow_period" table.
	FlowPeriodTable = &schema.Table{
		Name:       "flow_period",
		Columns:    FlowPeriodColumns,
		PrimaryKey: []*schema.Column{FlowPeriodColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "flow_period_participant_flow_module_flow_period_list",
				Columns:    []*schema.Column{FlowPeriodColumns[2]},
				RefColumns: []*schema.Column{ParticipantFlowModuleColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FlowReasonColumns holds the columns for the "flow_reason" table.
	FlowReasonColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "flow_reason_group_id", Type: field.TypeString},
		{Name: "flow_reason_comment", Type: field.TypeString},
		{Name: "flow_reason_num_subjects", Type: field.TypeString},
		{Name: "flow_reason_num_units", Type: field.TypeString},
		{Name: "flow_drop_withdraw_flow_reason_list", Type: field.TypeInt, Nullable: true},
	}
	// FlowReasonTable holds the schema information for the "flow_reason" table.
	FlowReasonTable = &schema.Table{
		Name:       "flow_reason",
		Columns:    FlowReasonColumns,
		PrimaryKey: []*schema.Column{FlowReasonColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "flow_reason_flow_drop_withdraw_flow_reason_list",
				Columns:    []*schema.Column{FlowReasonColumns[5]},
				RefColumns: []*schema.Column{FlowDropWithdrawColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ImmunocompromisedGroupsColumns holds the columns for the "immunocompromised_groups" table.
	ImmunocompromisedGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "group_name", Type: field.TypeString, Unique: true},
	}
	// ImmunocompromisedGroupsTable holds the schema information for the "immunocompromised_groups" table.
	ImmunocompromisedGroupsTable = &schema.Table{
		Name:       "immunocompromised_groups",
		Columns:    ImmunocompromisedGroupsColumns,
		PrimaryKey: []*schema.Column{ImmunocompromisedGroupsColumns[0]},
	}
	// ManufacturerColumns holds the columns for the "manufacturer" table.
	ManufacturerColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "manufacturer_name", Type: field.TypeString, Unique: true},
	}
	// ManufacturerTable holds the schema information for the "manufacturer" table.
	ManufacturerTable = &schema.Table{
		Name:       "manufacturer",
		Columns:    ManufacturerColumns,
		PrimaryKey: []*schema.Column{ManufacturerColumns[0]},
	}
	// MoreInfoModuleColumns holds the columns for the "more_info_module" table.
	MoreInfoModuleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "limitations_and_caveats_description", Type: field.TypeString},
		{Name: "more_info_module_certain_agreement", Type: field.TypeInt, Nullable: true},
		{Name: "more_info_module_point_of_contact", Type: field.TypeInt, Nullable: true},
		{Name: "results_definition_more_info_module", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// MoreInfoModuleTable holds the schema information for the "more_info_module" table.
	MoreInfoModuleTable = &schema.Table{
		Name:       "more_info_module",
		Columns:    MoreInfoModuleColumns,
		PrimaryKey: []*schema.Column{MoreInfoModuleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "more_info_module_certain_agreement_certain_agreement",
				Columns:    []*schema.Column{MoreInfoModuleColumns[2]},
				RefColumns: []*schema.Column{CertainAgreementColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "more_info_module_point_of_contact_point_of_contact",
				Columns:    []*schema.Column{MoreInfoModuleColumns[3]},
				RefColumns: []*schema.Column{PointOfContactColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "more_info_module_results_definition_more_info_module",
				Columns:    []*schema.Column{MoreInfoModuleColumns[4]},
				RefColumns: []*schema.Column{ResultsDefinitionColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OtherEventColumns holds the columns for the "other_event" table.
	OtherEventColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "other_event_term", Type: field.TypeString},
		{Name: "other_event_organ_system", Type: field.TypeString},
		{Name: "other_event_source_vocabulary", Type: field.TypeString},
		{Name: "other_event_assessment_type", Type: field.TypeString},
		{Name: "other_event_notes", Type: field.TypeString},
		{Name: "adverse_events_module_other_event_list", Type: field.TypeInt, Nullable: true},
	}
	// OtherEventTable holds the schema information for the "other_event" table.
	OtherEventTable = &schema.Table{
		Name:       "other_event",
		Columns:    OtherEventColumns,
		PrimaryKey: []*schema.Column{OtherEventColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "other_event_adverse_events_module_other_event_list",
				Columns:    []*schema.Column{OtherEventColumns[6]},
				RefColumns: []*schema.Column{AdverseEventsModuleColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OtherEventStatsColumns holds the columns for the "other_event_stats" table.
	OtherEventStatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "other_event_stats_group_id", Type: field.TypeString},
		{Name: "other_event_stats_num_events", Type: field.TypeString},
		{Name: "other_event_stats_num_affected", Type: field.TypeString},
		{Name: "other_event_stats_num_at_risk", Type: field.TypeString},
		{Name: "other_event_other_event_stats_list", Type: field.TypeInt, Nullable: true},
	}
	// OtherEventStatsTable holds the schema information for the "other_event_stats" table.
	OtherEventStatsTable = &schema.Table{
		Name:       "other_event_stats",
		Columns:    OtherEventStatsColumns,
		PrimaryKey: []*schema.Column{OtherEventStatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "other_event_stats_other_event_other_event_stats_list",
				Columns:    []*schema.Column{OtherEventStatsColumns[5]},
				RefColumns: []*schema.Column{OtherEventColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutcomeAnalysisColumns holds the columns for the "outcome_analysis" table.
	OutcomeAnalysisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "outcome_analysis_group_description", Type: field.TypeString},
		{Name: "outcome_analysis_tested_non_inferiority", Type: field.TypeString},
		{Name: "outcome_analysis_non_inferiority_type", Type: field.TypeString},
		{Name: "outcome_analysis_non_inferiority_comment", Type: field.TypeString},
		{Name: "outcome_analysis_p_value", Type: field.TypeString},
		{Name: "outcome_analysis_p_value_comment", Type: field.TypeString},
		{Name: "outcome_analysis_statistical_method", Type: field.TypeString},
		{Name: "outcome_analysis_statistical_comment", Type: field.TypeString},
		{Name: "outcome_analysis_param_type", Type: field.TypeString},
		{Name: "outcome_analysis_param_value", Type: field.TypeString},
		{Name: "outcome_analysis_ci_pct_value", Type: field.TypeString},
		{Name: "outcome_analysis_ci_num_sides", Type: field.TypeString},
		{Name: "outcome_analysis_ci_lower_limit", Type: field.TypeString},
		{Name: "outcome_analysis_ci_upper_limit", Type: field.TypeString},
		{Name: "outcome_analysis_ci_lower_limit_comment", Type: field.TypeString},
		{Name: "outcome_analysis_ci_upper_limit_comment", Type: field.TypeString},
		{Name: "outcome_analysis_dispersion_type", Type: field.TypeString},
		{Name: "outcome_analysis_dispersion_value", Type: field.TypeString},
		{Name: "outcome_analysis_estimate_comment", Type: field.TypeString},
		{Name: "outcome_analysis_other_analysis_description", Type: field.TypeString},
		{Name: "outcome_measure_outcome_analysis_list", Type: field.TypeInt, Nullable: true},
	}
	// OutcomeAnalysisTable holds the schema information for the "outcome_analysis" table.
	OutcomeAnalysisTable = &schema.Table{
		Name:       "outcome_analysis",
		Columns:    OutcomeAnalysisColumns,
		PrimaryKey: []*schema.Column{OutcomeAnalysisColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "outcome_analysis_outcome_measure_outcome_analysis_list",
				Columns:    []*schema.Column{OutcomeAnalysisColumns[21]},
				RefColumns: []*schema.Column{OutcomeMeasureColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutcomeAnalysisGroupIDColumns holds the columns for the "outcome_analysis_group_id" table.
	OutcomeAnalysisGroupIDColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "outcome_analysis_group_id", Type: field.TypeString},
		{Name: "outcome_analysis_outcome_analysis_group_id_list", Type: field.TypeInt, Nullable: true},
	}
	// OutcomeAnalysisGroupIDTable holds the schema information for the "outcome_analysis_group_id" table.
	OutcomeAnalysisGroupIDTable = &schema.Table{
		Name:       "outcome_analysis_group_id",
		Columns:    OutcomeAnalysisGroupIDColumns,
		PrimaryKey: []*schema.Column{OutcomeAnalysisGroupIDColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "outcome_analysis_group_id_outcome_analysis_outcome_analysis_group_id_list",
				Columns:    []*schema.Column{OutcomeAnalysisGroupIDColumns[2]},
				RefColumns: []*schema.Column{OutcomeAnalysisColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutcomeCategoryColumns holds the columns for the "outcome_category" table.
	OutcomeCategoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "outcome_category_title", Type: field.TypeString},
		{Name: "outcome_class_outcome_category_list", Type: field.TypeInt, Nullable: true},
	}
	// OutcomeCategoryTable holds the schema information for the "outcome_category" table.
	OutcomeCategoryTable = &schema.Table{
		Name:       "outcome_category",
		Columns:    OutcomeCategoryColumns,
		PrimaryKey: []*schema.Column{OutcomeCategoryColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "outcome_category_outcome_class_outcome_category_list",
				Columns:    []*schema.Column{OutcomeCategoryColumns[2]},
				RefColumns: []*schema.Column{OutcomeClassColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutcomeClassColumns holds the columns for the "outcome_class" table.
	OutcomeClassColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "outcome_class_title", Type: field.TypeString},
		{Name: "outcome_measure_outcome_class_list", Type: field.TypeInt, Nullable: true},
	}
	// OutcomeClassTable holds the schema information for the "outcome_class" table.
	OutcomeClassTable = &schema.Table{
		Name:       "outcome_class",
		Columns:    OutcomeClassColumns,
		PrimaryKey: []*schema.Column{OutcomeClassColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "outcome_class_outcome_measure_outcome_class_list",
				Columns:    []*schema.Column{OutcomeClassColumns[2]},
				RefColumns: []*schema.Column{OutcomeMeasureColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutcomeClassDenomColumns holds the columns for the "outcome_class_denom" table.
	OutcomeClassDenomColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "outcome_class_denom_units", Type: field.TypeString},
		{Name: "outcome_class_outcome_class_denom_list", Type: field.TypeInt, Nullable: true},
	}
	// OutcomeClassDenomTable holds the schema information for the "outcome_class_denom" table.
	OutcomeClassDenomTable = &schema.Table{
		Name:       "outcome_class_denom",
		Columns:    OutcomeClassDenomColumns,
		PrimaryKey: []*schema.Column{OutcomeClassDenomColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "outcome_class_denom_outcome_class_outcome_class_denom_list",
				Columns:    []*schema.Column{OutcomeClassDenomColumns[2]},
				RefColumns: []*schema.Column{OutcomeClassColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutcomeClassDenomCountColumns holds the columns for the "outcome_class_denom_count" table.
	OutcomeClassDenomCountColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "outcome_class_denom_count_group_id", Type: field.TypeString},
		{Name: "outcome_class_denom_count_value", Type: field.TypeString},
		{Name: "outcome_class_denom_outcome_class_denom_count_list", Type: field.TypeInt, Nullable: true},
	}
	// OutcomeClassDenomCountTable holds the schema information for the "outcome_class_denom_count" table.
	OutcomeClassDenomCountTable = &schema.Table{
		Name:       "outcome_class_denom_count",
		Columns:    OutcomeClassDenomCountColumns,
		PrimaryKey: []*schema.Column{OutcomeClassDenomCountColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "outcome_class_denom_count_outcome_class_denom_outcome_class_denom_count_list",
				Columns:    []*schema.Column{OutcomeClassDenomCountColumns[3]},
				RefColumns: []*schema.Column{OutcomeClassDenomColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutcomeDenomColumns holds the columns for the "outcome_denom" table.
	OutcomeDenomColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "outcome_denom_units", Type: field.TypeString},
		{Name: "outcome_measure_outcome_denom_list", Type: field.TypeInt, Nullable: true},
	}
	// OutcomeDenomTable holds the schema information for the "outcome_denom" table.
	OutcomeDenomTable = &schema.Table{
		Name:       "outcome_denom",
		Columns:    OutcomeDenomColumns,
		PrimaryKey: []*schema.Column{OutcomeDenomColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "outcome_denom_outcome_measure_outcome_denom_list",
				Columns:    []*schema.Column{OutcomeDenomColumns[2]},
				RefColumns: []*schema.Column{OutcomeMeasureColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutcomeDenomCountColumns holds the columns for the "outcome_denom_count" table.
	OutcomeDenomCountColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "outcome_denom_count_group_id", Type: field.TypeString},
		{Name: "outcome_denom_count_value", Type: field.TypeString},
		{Name: "outcome_denom_outcome_denom_count_list", Type: field.TypeInt, Nullable: true},
	}
	// OutcomeDenomCountTable holds the schema information for the "outcome_denom_count" table.
	OutcomeDenomCountTable = &schema.Table{
		Name:       "outcome_denom_count",
		Columns:    OutcomeDenomCountColumns,
		PrimaryKey: []*schema.Column{OutcomeDenomCountColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "outcome_denom_count_outcome_denom_outcome_denom_count_list",
				Columns:    []*schema.Column{OutcomeDenomCountColumns[3]},
				RefColumns: []*schema.Column{OutcomeDenomColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutcomeGroupColumns holds the columns for the "outcome_group" table.
	OutcomeGroupColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "outcome_group_id", Type: field.TypeString},
		{Name: "outcome_group_title", Type: field.TypeString},
		{Name: "outcome_group_description", Type: field.TypeString},
		{Name: "outcome_measure_outcome_group_list", Type: field.TypeInt, Nullable: true},
	}
	// OutcomeGroupTable holds the schema information for the "outcome_group" table.
	OutcomeGroupTable = &schema.Table{
		Name:       "outcome_group",
		Columns:    OutcomeGroupColumns,
		PrimaryKey: []*schema.Column{OutcomeGroupColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "outcome_group_outcome_measure_outcome_group_list",
				Columns:    []*schema.Column{OutcomeGroupColumns[4]},
				RefColumns: []*schema.Column{OutcomeMeasureColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutcomeMeasureColumns holds the columns for the "outcome_measure" table.
	OutcomeMeasureColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "outcome_measure_type", Type: field.TypeString},
		{Name: "outcome_measure_title", Type: field.TypeString},
		{Name: "outcome_measure_description", Type: field.TypeString},
		{Name: "outcome_measure_population_description", Type: field.TypeString},
		{Name: "outcome_measure_reporting_status", Type: field.TypeString},
		{Name: "outcome_measure_anticipated_posting_date", Type: field.TypeString},
		{Name: "outcome_measure_param_type", Type: field.TypeString},
		{Name: "outcome_measure_dispersion_type", Type: field.TypeString},
		{Name: "outcome_measure_unit_of_measure", Type: field.TypeString},
		{Name: "outcome_measure_calculate_pct", Type: field.TypeString},
		{Name: "outcome_measure_time_frame", Type: field.TypeString},
		{Name: "outcome_measure_type_units_analyzed", Type: field.TypeString},
		{Name: "outcome_measure_denom_units_selected", Type: field.TypeString},
		{Name: "outcome_measures_module_outcome_measure_list", Type: field.TypeInt, Nullable: true},
	}
	// OutcomeMeasureTable holds the schema information for the "outcome_measure" table.
	OutcomeMeasureTable = &schema.Table{
		Name:       "outcome_measure",
		Columns:    OutcomeMeasureColumns,
		PrimaryKey: []*schema.Column{OutcomeMeasureColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "outcome_measure_outcome_measures_module_outcome_measure_list",
				Columns:    []*schema.Column{OutcomeMeasureColumns[14]},
				RefColumns: []*schema.Column{OutcomeMeasuresModuleColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutcomeMeasurementColumns holds the columns for the "outcome_measurement" table.
	OutcomeMeasurementColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "outcome_measurement_group_id", Type: field.TypeString},
		{Name: "outcome_measurement_value", Type: field.TypeString},
		{Name: "outcome_measurement_spread", Type: field.TypeString},
		{Name: "outcome_measurement_lower_limit", Type: field.TypeString},
		{Name: "outcome_measurement_upper_limit", Type: field.TypeString},
		{Name: "outcome_measurement_comment", Type: field.TypeString},
		{Name: "outcome_category_outcome_measurement_list", Type: field.TypeInt, Nullable: true},
	}
	// OutcomeMeasurementTable holds the schema information for the "outcome_measurement" table.
	OutcomeMeasurementTable = &schema.Table{
		Name:       "outcome_measurement",
		Columns:    OutcomeMeasurementColumns,
		PrimaryKey: []*schema.Column{OutcomeMeasurementColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "outcome_measurement_outcome_category_outcome_measurement_list",
				Columns:    []*schema.Column{OutcomeMeasurementColumns[7]},
				RefColumns: []*schema.Column{OutcomeCategoryColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutcomeMeasuresModuleColumns holds the columns for the "outcome_measures_module" table.
	OutcomeMeasuresModuleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "results_definition_outcome_measures_module", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// OutcomeMeasuresModuleTable holds the schema information for the "outcome_measures_module" table.
	OutcomeMeasuresModuleTable = &schema.Table{
		Name:       "outcome_measures_module",
		Columns:    OutcomeMeasuresModuleColumns,
		PrimaryKey: []*schema.Column{OutcomeMeasuresModuleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "outcome_measures_module_results_definition_outcome_measures_module",
				Columns:    []*schema.Column{OutcomeMeasuresModuleColumns[1]},
				RefColumns: []*schema.Column{ResultsDefinitionColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutcomeOverviewColumns holds the columns for the "outcome_overview" table.
	OutcomeOverviewColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "outcome_overview_id", Type: field.TypeString},
		{Name: "outcome_overview_title", Type: field.TypeString},
		{Name: "outcome_overview_description", Type: field.TypeString},
		{Name: "outcome_overview_participants", Type: field.TypeString},
		{Name: "outcome_overview_time_frame", Type: field.TypeString},
		{Name: "outcome_overview_serotype", Type: field.TypeString},
		{Name: "outcome_overview_assay", Type: field.TypeString},
		{Name: "outcome_overview_dose_number", Type: field.TypeInt64, Default: 0},
		{Name: "outcome_overview_value", Type: field.TypeFloat64},
		{Name: "outcome_overview_upper_limit", Type: field.TypeString},
		{Name: "outcome_overview_lower_limit", Type: field.TypeString},
		{Name: "outcome_overview_group_id", Type: field.TypeString, Default: ""},
		{Name: "outcome_overview_ratio", Type: field.TypeString, Default: ""},
		{Name: "outcome_overview_measure_title", Type: field.TypeString, Default: ""},
		{Name: "outcome_overview_vaccine", Type: field.TypeString, Default: ""},
		{Name: "outcome_overview_immunocompromised_population", Type: field.TypeString, Default: "N/A"},
		{Name: "outcome_overview_manufacturer", Type: field.TypeString, Default: ""},
		{Name: "outcome_overview_confidence_interval", Type: field.TypeString, Default: ""},
		{Name: "outcome_overview_percent_responders", Type: field.TypeFloat64, Default: 0},
		{Name: "outcome_overview_time_frame_weeks", Type: field.TypeInt64, Default: 0},
		{Name: "outcome_overview_dose_description", Type: field.TypeString, Default: ""},
		{Name: "outcome_overview_schedule", Type: field.TypeString, Default: ""},
		{Name: "outcome_overview_use_case_code", Type: field.TypeString, Default: "default"},
		{Name: "outcome_measure_outcome_overview_list", Type: field.TypeInt, Nullable: true},
	}
	// OutcomeOverviewTable holds the schema information for the "outcome_overview" table.
	OutcomeOverviewTable = &schema.Table{
		Name:       "outcome_overview",
		Columns:    OutcomeOverviewColumns,
		PrimaryKey: []*schema.Column{OutcomeOverviewColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "outcome_overview_outcome_measure_outcome_overview_list",
				Columns:    []*schema.Column{OutcomeOverviewColumns[24]},
				RefColumns: []*schema.Column{OutcomeMeasureColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ParticipantFlowModuleColumns holds the columns for the "participant_flow_module" table.
	ParticipantFlowModuleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "flow_pre_assignment_details", Type: field.TypeString},
		{Name: "flow_recruitment_details", Type: field.TypeString},
		{Name: "flow_type_units_analyzed", Type: field.TypeString},
		{Name: "results_definition_participant_flow_module", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// ParticipantFlowModuleTable holds the schema information for the "participant_flow_module" table.
	ParticipantFlowModuleTable = &schema.Table{
		Name:       "participant_flow_module",
		Columns:    ParticipantFlowModuleColumns,
		PrimaryKey: []*schema.Column{ParticipantFlowModuleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "participant_flow_module_results_definition_participant_flow_module",
				Columns:    []*schema.Column{ParticipantFlowModuleColumns[4]},
				RefColumns: []*schema.Column{ResultsDefinitionColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PointOfContactColumns holds the columns for the "point_of_contact" table.
	PointOfContactColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "point_of_contact_title", Type: field.TypeString},
		{Name: "point_of_contact_organization", Type: field.TypeString},
		{Name: "point_of_contact_email", Type: field.TypeString},
		{Name: "point_of_contact_phone", Type: field.TypeString},
		{Name: "point_of_contact_phone_ext", Type: field.TypeString},
		{Name: "point_of_contact_parent", Type: field.TypeInt},
	}
	// PointOfContactTable holds the schema information for the "point_of_contact" table.
	PointOfContactTable = &schema.Table{
		Name:       "point_of_contact",
		Columns:    PointOfContactColumns,
		PrimaryKey: []*schema.Column{PointOfContactColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "point_of_contact_more_info_module_parent",
				Columns:    []*schema.Column{PointOfContactColumns[6]},
				RefColumns: []*schema.Column{MoreInfoModuleColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ResultsDefinitionColumns holds the columns for the "results_definition" table.
	ResultsDefinitionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "clinical_trial_results_definition", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// ResultsDefinitionTable holds the schema information for the "results_definition" table.
	ResultsDefinitionTable = &schema.Table{
		Name:       "results_definition",
		Columns:    ResultsDefinitionColumns,
		PrimaryKey: []*schema.Column{ResultsDefinitionColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "results_definition_clinical_trial_results_definition",
				Columns:    []*schema.Column{ResultsDefinitionColumns[1]},
				RefColumns: []*schema.Column{ClinicalTrialColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ScheduleColumns holds the columns for the "schedule" table.
	ScheduleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// ScheduleTable holds the schema information for the "schedule" table.
	ScheduleTable = &schema.Table{
		Name:       "schedule",
		Columns:    ScheduleColumns,
		PrimaryKey: []*schema.Column{ScheduleColumns[0]},
	}
	// SeriousEventColumns holds the columns for the "serious_event" table.
	SeriousEventColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "serious_event_term", Type: field.TypeString},
		{Name: "serious_event_organ_system", Type: field.TypeString},
		{Name: "serious_event_source_vocabulary", Type: field.TypeString},
		{Name: "serious_event_assessment_type", Type: field.TypeString},
		{Name: "serious_event_notes", Type: field.TypeString},
		{Name: "adverse_events_module_serious_event_list", Type: field.TypeInt, Nullable: true},
	}
	// SeriousEventTable holds the schema information for the "serious_event" table.
	SeriousEventTable = &schema.Table{
		Name:       "serious_event",
		Columns:    SeriousEventColumns,
		PrimaryKey: []*schema.Column{SeriousEventColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "serious_event_adverse_events_module_serious_event_list",
				Columns:    []*schema.Column{SeriousEventColumns[6]},
				RefColumns: []*schema.Column{AdverseEventsModuleColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SeriousEventStatsColumns holds the columns for the "serious_event_stats" table.
	SeriousEventStatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "serious_event_stats_group_id", Type: field.TypeString},
		{Name: "serious_event_stats_num_events", Type: field.TypeString},
		{Name: "serious_event_stats_num_affected", Type: field.TypeString},
		{Name: "serious_event_stats_num_at_risk", Type: field.TypeString},
		{Name: "serious_event_serious_event_stats_list", Type: field.TypeInt, Nullable: true},
	}
	// SeriousEventStatsTable holds the schema information for the "serious_event_stats" table.
	SeriousEventStatsTable = &schema.Table{
		Name:       "serious_event_stats",
		Columns:    SeriousEventStatsColumns,
		PrimaryKey: []*schema.Column{SeriousEventStatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "serious_event_stats_serious_event_serious_event_stats_list",
				Columns:    []*schema.Column{SeriousEventStatsColumns[5]},
				RefColumns: []*schema.Column{SeriousEventColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StudyEligibilityColumns holds the columns for the "study_eligibility" table.
	StudyEligibilityColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "eligibility_criteria", Type: field.TypeString},
		{Name: "healthy_volunteers", Type: field.TypeString},
		{Name: "gender", Type: field.TypeString},
		{Name: "minimum_age", Type: field.TypeString},
		{Name: "maximum_age", Type: field.TypeString},
		{Name: "std_age_list", Type: field.TypeString},
		{Name: "ethnicity", Type: field.TypeString},
		{Name: "clinical_trial_study_eligibility", Type: field.TypeInt, Nullable: true},
	}
	// StudyEligibilityTable holds the schema information for the "study_eligibility" table.
	StudyEligibilityTable = &schema.Table{
		Name:       "study_eligibility",
		Columns:    StudyEligibilityColumns,
		PrimaryKey: []*schema.Column{StudyEligibilityColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "study_eligibility_clinical_trial_study_eligibility",
				Columns:    []*schema.Column{StudyEligibilityColumns[8]},
				RefColumns: []*schema.Column{ClinicalTrialColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StudyLocationColumns holds the columns for the "study_location" table.
	StudyLocationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "location_facility", Type: field.TypeString},
		{Name: "location_city", Type: field.TypeString},
		{Name: "location_country", Type: field.TypeString},
		{Name: "location_country_code", Type: field.TypeString},
		{Name: "location_continent_name", Type: field.TypeString},
		{Name: "location_continent_code", Type: field.TypeString},
		{Name: "clinical_trial_study_locations", Type: field.TypeInt, Nullable: true},
	}
	// StudyLocationTable holds the schema information for the "study_location" table.
	StudyLocationTable = &schema.Table{
		Name:       "study_location",
		Columns:    StudyLocationColumns,
		PrimaryKey: []*schema.Column{StudyLocationColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "study_location_clinical_trial_study_locations",
				Columns:    []*schema.Column{StudyLocationColumns[7]},
				RefColumns: []*schema.Column{ClinicalTrialColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "owner", Type: field.TypeString},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"todo", "inprogress", "done"}, Default: "todo"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
	}
	// UseCaseColumns holds the columns for the "use_case" table.
	UseCaseColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "use_case_name", Type: field.TypeString},
		{Name: "use_case_description", Type: field.TypeString},
		{Name: "use_case_code", Type: field.TypeString},
	}
	// UseCaseTable holds the schema information for the "use_case" table.
	UseCaseTable = &schema.Table{
		Name:       "use_case",
		Columns:    UseCaseColumns,
		PrimaryKey: []*schema.Column{UseCaseColumns[0]},
	}
	// VaccineColumns holds the columns for the "vaccine" table.
	VaccineColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "vaccine_name", Type: field.TypeString, Unique: true},
	}
	// VaccineTable holds the schema information for the "vaccine" table.
	VaccineTable = &schema.Table{
		Name:       "vaccine",
		Columns:    VaccineColumns,
		PrimaryKey: []*schema.Column{VaccineColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdverseEventsModuleTable,
		BaselineCategoryTable,
		BaselineCharacteristicsModuleTable,
		BaselineClassTable,
		BaselineClassDenomTable,
		BaselineClassDenomCountTable,
		BaselineDenomTable,
		BaselineDenomCountTable,
		BaselineGroupTable,
		BaselineMeasureTable,
		BaselineMeasureDenomTable,
		BaselineMeasureDenomCountTable,
		BaselineMeasurementTable,
		CertainAgreementTable,
		ClinicalTrialTable,
		ClinicalTrialMetadataTable,
		DoseDescriptionTable,
		EventGroupTable,
		FlowAchievementTable,
		FlowDropWithdrawTable,
		FlowGroupTable,
		FlowMilestoneTable,
		FlowPeriodTable,
		FlowReasonTable,
		ImmunocompromisedGroupsTable,
		ManufacturerTable,
		MoreInfoModuleTable,
		OtherEventTable,
		OtherEventStatsTable,
		OutcomeAnalysisTable,
		OutcomeAnalysisGroupIDTable,
		OutcomeCategoryTable,
		OutcomeClassTable,
		OutcomeClassDenomTable,
		OutcomeClassDenomCountTable,
		OutcomeDenomTable,
		OutcomeDenomCountTable,
		OutcomeGroupTable,
		OutcomeMeasureTable,
		OutcomeMeasurementTable,
		OutcomeMeasuresModuleTable,
		OutcomeOverviewTable,
		ParticipantFlowModuleTable,
		PointOfContactTable,
		ResultsDefinitionTable,
		ScheduleTable,
		SeriousEventTable,
		SeriousEventStatsTable,
		StudyEligibilityTable,
		StudyLocationTable,
		TasksTable,
		UseCaseTable,
		VaccineTable,
	}
)

func init() {
	AdverseEventsModuleTable.ForeignKeys[0].RefTable = ResultsDefinitionTable
	AdverseEventsModuleTable.Annotation = &entsql.Annotation{
		Table: "adverse_events_module",
	}
	BaselineCategoryTable.ForeignKeys[0].RefTable = BaselineClassTable
	BaselineCategoryTable.Annotation = &entsql.Annotation{
		Table: "baseline_category",
	}
	BaselineCharacteristicsModuleTable.ForeignKeys[0].RefTable = ResultsDefinitionTable
	BaselineCharacteristicsModuleTable.Annotation = &entsql.Annotation{
		Table: "baseline_characteristics_module",
	}
	BaselineClassTable.ForeignKeys[0].RefTable = BaselineMeasureTable
	BaselineClassTable.Annotation = &entsql.Annotation{
		Table: "baseline_class",
	}
	BaselineClassDenomTable.ForeignKeys[0].RefTable = BaselineClassTable
	BaselineClassDenomTable.Annotation = &entsql.Annotation{
		Table: "baseline_class_denom",
	}
	BaselineClassDenomCountTable.ForeignKeys[0].RefTable = BaselineClassDenomTable
	BaselineClassDenomCountTable.Annotation = &entsql.Annotation{
		Table: "baseline_class_denom_count",
	}
	BaselineDenomTable.ForeignKeys[0].RefTable = BaselineCharacteristicsModuleTable
	BaselineDenomTable.Annotation = &entsql.Annotation{
		Table: "baseline_denom",
	}
	BaselineDenomCountTable.ForeignKeys[0].RefTable = BaselineDenomTable
	BaselineDenomCountTable.Annotation = &entsql.Annotation{
		Table: "baseline_denom_count",
	}
	BaselineGroupTable.ForeignKeys[0].RefTable = BaselineCharacteristicsModuleTable
	BaselineGroupTable.Annotation = &entsql.Annotation{
		Table: "baseline_group",
	}
	BaselineMeasureTable.ForeignKeys[0].RefTable = BaselineCharacteristicsModuleTable
	BaselineMeasureTable.Annotation = &entsql.Annotation{
		Table: "baseline_measure",
	}
	BaselineMeasureDenomTable.ForeignKeys[0].RefTable = BaselineMeasureTable
	BaselineMeasureDenomTable.Annotation = &entsql.Annotation{
		Table: "baseline_measure_denom",
	}
	BaselineMeasureDenomCountTable.ForeignKeys[0].RefTable = BaselineMeasureDenomTable
	BaselineMeasureDenomCountTable.Annotation = &entsql.Annotation{
		Table: "baseline_measure_denom_count",
	}
	BaselineMeasurementTable.ForeignKeys[0].RefTable = BaselineCategoryTable
	BaselineMeasurementTable.Annotation = &entsql.Annotation{
		Table: "baseline_measurement",
	}
	CertainAgreementTable.ForeignKeys[0].RefTable = MoreInfoModuleTable
	CertainAgreementTable.Annotation = &entsql.Annotation{
		Table: "certain_agreement",
	}
	ClinicalTrialTable.Annotation = &entsql.Annotation{
		Table: "clinical_trial",
	}
	ClinicalTrialMetadataTable.ForeignKeys[0].RefTable = ClinicalTrialTable
	ClinicalTrialMetadataTable.Annotation = &entsql.Annotation{
		Table: "clinical_trial_metadata",
	}
	DoseDescriptionTable.Annotation = &entsql.Annotation{
		Table: "dose_description",
	}
	EventGroupTable.ForeignKeys[0].RefTable = AdverseEventsModuleTable
	EventGroupTable.Annotation = &entsql.Annotation{
		Table: "event_group",
	}
	FlowAchievementTable.ForeignKeys[0].RefTable = FlowMilestoneTable
	FlowAchievementTable.Annotation = &entsql.Annotation{
		Table: "flow_achievement",
	}
	FlowDropWithdrawTable.ForeignKeys[0].RefTable = FlowPeriodTable
	FlowDropWithdrawTable.Annotation = &entsql.Annotation{
		Table: "flow_drop_withdraw",
	}
	FlowGroupTable.ForeignKeys[0].RefTable = ParticipantFlowModuleTable
	FlowGroupTable.Annotation = &entsql.Annotation{
		Table: "flow_group",
	}
	FlowMilestoneTable.ForeignKeys[0].RefTable = FlowPeriodTable
	FlowMilestoneTable.Annotation = &entsql.Annotation{
		Table: "flow_milestone",
	}
	FlowPeriodTable.ForeignKeys[0].RefTable = ParticipantFlowModuleTable
	FlowPeriodTable.Annotation = &entsql.Annotation{
		Table: "flow_period",
	}
	FlowReasonTable.ForeignKeys[0].RefTable = FlowDropWithdrawTable
	FlowReasonTable.Annotation = &entsql.Annotation{
		Table: "flow_reason",
	}
	ImmunocompromisedGroupsTable.Annotation = &entsql.Annotation{
		Table: "immunocompromised_groups",
	}
	ManufacturerTable.Annotation = &entsql.Annotation{
		Table: "manufacturer",
	}
	MoreInfoModuleTable.ForeignKeys[0].RefTable = CertainAgreementTable
	MoreInfoModuleTable.ForeignKeys[1].RefTable = PointOfContactTable
	MoreInfoModuleTable.ForeignKeys[2].RefTable = ResultsDefinitionTable
	MoreInfoModuleTable.Annotation = &entsql.Annotation{
		Table: "more_info_module",
	}
	OtherEventTable.ForeignKeys[0].RefTable = AdverseEventsModuleTable
	OtherEventTable.Annotation = &entsql.Annotation{
		Table: "other_event",
	}
	OtherEventStatsTable.ForeignKeys[0].RefTable = OtherEventTable
	OtherEventStatsTable.Annotation = &entsql.Annotation{
		Table: "other_event_stats",
	}
	OutcomeAnalysisTable.ForeignKeys[0].RefTable = OutcomeMeasureTable
	OutcomeAnalysisTable.Annotation = &entsql.Annotation{
		Table: "outcome_analysis",
	}
	OutcomeAnalysisGroupIDTable.ForeignKeys[0].RefTable = OutcomeAnalysisTable
	OutcomeAnalysisGroupIDTable.Annotation = &entsql.Annotation{
		Table: "outcome_analysis_group_id",
	}
	OutcomeCategoryTable.ForeignKeys[0].RefTable = OutcomeClassTable
	OutcomeCategoryTable.Annotation = &entsql.Annotation{
		Table: "outcome_category",
	}
	OutcomeClassTable.ForeignKeys[0].RefTable = OutcomeMeasureTable
	OutcomeClassTable.Annotation = &entsql.Annotation{
		Table: "outcome_class",
	}
	OutcomeClassDenomTable.ForeignKeys[0].RefTable = OutcomeClassTable
	OutcomeClassDenomTable.Annotation = &entsql.Annotation{
		Table: "outcome_class_denom",
	}
	OutcomeClassDenomCountTable.ForeignKeys[0].RefTable = OutcomeClassDenomTable
	OutcomeClassDenomCountTable.Annotation = &entsql.Annotation{
		Table: "outcome_class_denom_count",
	}
	OutcomeDenomTable.ForeignKeys[0].RefTable = OutcomeMeasureTable
	OutcomeDenomTable.Annotation = &entsql.Annotation{
		Table: "outcome_denom",
	}
	OutcomeDenomCountTable.ForeignKeys[0].RefTable = OutcomeDenomTable
	OutcomeDenomCountTable.Annotation = &entsql.Annotation{
		Table: "outcome_denom_count",
	}
	OutcomeGroupTable.ForeignKeys[0].RefTable = OutcomeMeasureTable
	OutcomeGroupTable.Annotation = &entsql.Annotation{
		Table: "outcome_group",
	}
	OutcomeMeasureTable.ForeignKeys[0].RefTable = OutcomeMeasuresModuleTable
	OutcomeMeasureTable.Annotation = &entsql.Annotation{
		Table: "outcome_measure",
	}
	OutcomeMeasurementTable.ForeignKeys[0].RefTable = OutcomeCategoryTable
	OutcomeMeasurementTable.Annotation = &entsql.Annotation{
		Table: "outcome_measurement",
	}
	OutcomeMeasuresModuleTable.ForeignKeys[0].RefTable = ResultsDefinitionTable
	OutcomeMeasuresModuleTable.Annotation = &entsql.Annotation{
		Table: "outcome_measures_module",
	}
	OutcomeOverviewTable.ForeignKeys[0].RefTable = OutcomeMeasureTable
	OutcomeOverviewTable.Annotation = &entsql.Annotation{
		Table: "outcome_overview",
	}
	ParticipantFlowModuleTable.ForeignKeys[0].RefTable = ResultsDefinitionTable
	ParticipantFlowModuleTable.Annotation = &entsql.Annotation{
		Table: "participant_flow_module",
	}
	PointOfContactTable.ForeignKeys[0].RefTable = MoreInfoModuleTable
	PointOfContactTable.Annotation = &entsql.Annotation{
		Table: "point_of_contact",
	}
	ResultsDefinitionTable.ForeignKeys[0].RefTable = ClinicalTrialTable
	ResultsDefinitionTable.Annotation = &entsql.Annotation{
		Table: "results_definition",
	}
	ScheduleTable.Annotation = &entsql.Annotation{
		Table: "schedule",
	}
	SeriousEventTable.ForeignKeys[0].RefTable = AdverseEventsModuleTable
	SeriousEventTable.Annotation = &entsql.Annotation{
		Table: "serious_event",
	}
	SeriousEventStatsTable.ForeignKeys[0].RefTable = SeriousEventTable
	SeriousEventStatsTable.Annotation = &entsql.Annotation{
		Table: "serious_event_stats",
	}
	StudyEligibilityTable.ForeignKeys[0].RefTable = ClinicalTrialTable
	StudyEligibilityTable.Annotation = &entsql.Annotation{
		Table: "study_eligibility",
	}
	StudyLocationTable.ForeignKeys[0].RefTable = ClinicalTrialTable
	StudyLocationTable.Annotation = &entsql.Annotation{
		Table: "study_location",
	}
	TasksTable.Annotation = &entsql.Annotation{
		Table: "tasks",
	}
	UseCaseTable.Annotation = &entsql.Annotation{
		Table: "use_case",
	}
	VaccineTable.Annotation = &entsql.Annotation{
		Table: "vaccine",
	}
}
