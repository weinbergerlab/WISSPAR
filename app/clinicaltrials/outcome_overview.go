package clinicaltrials

type OutcomeOverview struct {
	OutcomeOverviewTitle                       string  `json:"outcome_overview_title"`
	OutcomeOverviewID                          string  `json:"outcome_overview_id"`
	OutcomeOverviewDescription                 string  `json:"outcome_overview_description"`
	OutcomeOverviewTimeFrame                   string  `json:"outcome_overview_time_frame"`
	OutcomeOverviewAssay                       string  `json:"outcome_overview_assay"`
	OutcomeOverviewDoseNumber                  int64   `json:"outcome_overview_dose_number"`
	OutcomeMeasureTitle                        string  `json:"outcome_measure_title"`
	OutcomeOverviewParticipants                string  `json:"outcome_overview_participants"`
	OutcomeOverviewSerotype                    string  `json:"outcome_overview_serotype"`
	OutcomeOverviewValue                       float64 `json:"outcome_overview_value"`
	OutcomeOverviewUpperLimit                  string  `json:"outcome_overview_upper_limit"`
	OutcomeOverviewLowerLimit                  string  `json:"outcome_overview_lower_limit"`
	OutcomeOverviewGroupID                     string  `json:"outcome_overview_group_id"`
	OutcomeOverviewRatio                       string  `json:"outcome_overview_ratio"`
	OutcomeOverviewVaccine                     string  `json:"outcome_overview_vaccine"`
	OutcomeOverviewImmunocompromisedPopulation string  `json:"outcome_overview_immunocompromised_population"`
	OutcomeOverviewManufacturer                string  `json:"outcome_overview_manufacturer"`
	OutcomeOverviewTimeFrameWeeks              int64   `json:"outcome_overview_time_frame_weeks"`
	OutcomeOverviewDoseDescription             string  `json:"outcome_overview_dose_description"`
	OutcomeOverviewSchedule                    string  `json:"outcome_overview_schedule"`
	OutcomeOverviewConfidenceInterval          string  `json:"outcome_overview_confidence_interval"`
	OutcomeOverviewPercentResponders           float64 `json:"outcome_overview_percent_responders"`
	OutcomeOverviewUseCase                     string  `json:"outcome_overview_use_case"`
	Enabled                                    bool    `json:"enabled"`
}
