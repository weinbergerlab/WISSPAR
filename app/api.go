package app

import (
	"fmt"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/clinicaltrials"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/dashboards"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/dosedescription"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/task"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/immunocompromisedgroups"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/manufacturer"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/metadata"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/schedule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/usecase"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/vaccine"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"

	"github.com/lithammer/shortuuid/v3"

	"github.com/go-chi/render"

	"github.com/adnaan/authn"
)


// Get clinical trial and put it into a Study struct
func Get(t Context)  http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		rc := clinicaltrials.NewResultsContext(r.Context() , t.db)
		study, err := rc.GetTrial(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		format := r.URL.Query().Get("format")
		if format == "" {
			format = "json"
		}

		section := r.URL.Query().Get("section")

		if format == "html" {

			studyHTML := ""

			var htmlErr error

			switch section {
			case "document":
				studyHTML, htmlErr = clinicaltrials.GetDocumentHTML(study)
			case "protocol":
				studyHTML, htmlErr = clinicaltrials.GetProtocolHTML(study)
			case "results":
				studyHTML, htmlErr = clinicaltrials.GetResultsHTML(study)
			case "derived":
				studyHTML, htmlErr = clinicaltrials.GetDerivedHTML(study)
			default:
				studyHTML, htmlErr = clinicaltrials.GetTrialHTML(study)
			}

			if htmlErr != nil {
				render.Render(w, r, ErrInternal(err))
				return
			}

			render.HTML(w, r, studyHTML)
			return
		}

		render.JSON(w, r, study)
	}
}

type overviewList struct {
	SelectedOverviews []clinicaltrials.OutcomeOverview `json:"selected_overviews"`
	TrialOverviews []clinicaltrials.OutcomeOverview `json:"trial_overviews"`
}

// GetOutcomeOverviewList for manual input
func PostCSVExport(t Context) http.HandlerFunc {
	type req struct {
		SearchFilters clinicaltrials.SearchOptions `json:"search_options"`
		FieldFilters clinicaltrials.FieldOptions `json:"field_options"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := clinicaltrials.NewResultsContext(r.Context() , t.db)

		csvData, err := rc.GetExport(req.SearchFilters, req.FieldFilters)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment;filename=wisspar_data_export.csv")
		w.Write(csvData.Bytes())
	}
}

// GetOutcomeOverviewList for manual input
func GetCSVExport(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()

		sf := clinicaltrials.SearchOptions{
			UseCase: r.URL.Query().Get("use_case"),
			IncomeGroups: r.Form["income_groups"],
			ParentArraySeparator: r.URL.Query().Get("parent_array_separator"),
			ChildArraySeparator: r.URL.Query().Get("child_array_separator"),
			Vaccine: r.Form["vaccine"],
			Phase: r.Form["phase"],
			STDAgeList: r.Form["std_age_list"],
			CountryCodes: r.Form["country_codes"],
			ContinentCodes: r.Form["continent_codes"],
			ResponsibleParty: r.Form["responsible_party"],
			EthnicityRace: r.Form["ethnicity_race"],
			ImmunocompromisedPopulations: r.Form["immunocompromised_populations"],
		}

		ff := clinicaltrials.FieldOptions{
			LocationFacility:                           isFieldTrue(r, "location_facility"),
			LocationCity:                               isFieldTrue(r, "location_city"),
			LocationCountry:                            isFieldTrue(r, "location_country"),
			LocationCountryCode:                        isFieldTrue(r, "location_country_code"),
			LocationContinentName:                      isFieldTrue(r, "location_continent"),
			LocationContinentCode:                      isFieldTrue(r, "location_continent_code"),
			EligibilityCriteria:                        isFieldTrue(r, "study_eligibility_criteria"),
			HealthyVolunteers:                          isFieldTrue(r, "study_eligibility_healthy_volunteers"),
			Gender:                                     isFieldTrue(r, "study_eligibility_gender"),
			MinimumAge:                                 isFieldTrue(r, "study_eligibility_minimum_age"),
			MaximumAge:                                 isFieldTrue(r, "study_eligibility_maximum_age"),
			StandardAgeList:                            isFieldTrue(r, "study_eligibility_standard_age_list"),
			Ethnicity:                                  isFieldTrue(r, "study_eligibility_ethnicity"),
			ClinicalTrialStudyName:                     isFieldTrue(r, "clinical_trial_study_name"),
			ClinicalTrialStudyID:                       isFieldTrue(r, "clinical_trial_study_id"),
			Sponsor:                                    isFieldTrue(r, "clinical_trial_sponsor"),
			ResponsibleParty:                           isFieldTrue(r, "clinical_trial_responsible_party"),
			StudyType:                                  isFieldTrue(r, "clinical_trial_study_type"),
			Phase:                                      isFieldTrue(r, "clinical_trial_phase"),
			ActualEnrollment:                           isFieldTrue(r, "clinical_trial_actual_enrollment"),
			Allocation:                                 isFieldTrue(r, "clinical_trial_allocation"),
			InterventionModel:                          isFieldTrue(r, "clinical_trial_intervention_model"),
			Masking:                                    isFieldTrue(r, "clinical_trial_masking"),
			PrimaryPurpose:                             isFieldTrue(r, "clinical_trial_primary_purpose"),
			OfficialTitle:                              isFieldTrue(r, "clinical_trial_official_title"),
			ActualStudyStartDate:                       isFieldTrue(r, "clinical_trial_actual_study_start_date"),
			ActualPrimaryCompletionDate:                isFieldTrue(r, "clinical_trial_actual_primary_completion_date"),
			ActualStudyCompletionDate:                  isFieldTrue(r, "clinical_trial_actual_study_completion_date"),
			OutcomeMeasureType:                         isFieldTrue(r, "outcome_measure_type"),
			OutcomeMeasureTitle:                        isFieldTrue(r, "outcome_measure_title"),
			OutcomeMeasureDescription:                  isFieldTrue(r, "outcome_measure_description"),
			OutcomeMeasurePopulationDescription:        isFieldTrue(r, "outcome_measure_population_description"),
			OutcomeMeasureReportingStatus:              isFieldTrue(r, "outcome_measure_reporting_status"),
			OutcomeMeasureAnticipatedPostingDate:       isFieldTrue(r, "outcome_measure_anticipated_posting_date"),
			OutcomeMeasureParamType:                    isFieldTrue(r, "outcome_measure_param_type"),
			OutcomeMeasureDispersionType:               isFieldTrue(r, "outcome_measure_dispersion_type"),
			OutcomeMeasureUnitOfMeasure:                isFieldTrue(r, "outcome_measure_unit_of_measure"),
			OutcomeMeasureCalculatePCT:                 isFieldTrue(r, "outcome_measure_calculate_pct"),
			OutcomeMeasureTimeFrame:                    isFieldTrue(r, "outcome_measure_time_frame"),
			OutcomeMeasureTypeUnitsAnalyzed:            isFieldTrue(r, "outcome_measure_type_units_analyzed"),
			OutcomeMeasureDenomUnitsSelected:           isFieldTrue(r, "outcome_measure_denom_units_selected"),
			OutcomeOverviewTitle:                       isFieldTrue(r, "outcome_overview_title"),
			OutcomeOverviewID:                          isFieldTrue(r, "outcome_overview_id"),
			OutcomeOverviewDescription:                 isFieldTrue(r, "outcome_overview_description"),
			OutcomeOverviewTimeFrame:                   isFieldTrue(r, "outcome_overview_time_frame"),
			OutcomeOverviewAssay:                       isFieldTrue(r, "outcome_overview_assay"),
			OutcomeOverviewDoseNumber:                  isFieldTrue(r, "outcome_overview_dose_number"),
			OutcomeOverviewParticipants:                isFieldTrue(r, "outcome_overview_participants"),
			OutcomeOverviewSerotype:                    isFieldTrue(r, "outcome_overview_serotype"),
			OutcomeOverviewValue:                       isFieldTrue(r, "outcome_overview_value"),
			OutcomeOverviewUpperLimit:                  isFieldTrue(r, "outcome_overview_upper_limit"),
			OutcomeOverviewLowerLimit:                  isFieldTrue(r, "outcome_overview_lower_limit"),
			OutcomeOverviewRatio:                       isFieldTrue(r, "outcome_overview_ratio"),
			OutcomeOverviewVaccine:                     isFieldTrue(r, "outcome_overview_vaccine"),
			OutcomeOverviewImmunocompromisedPopulation: isFieldTrue(r, "outcome_overview_immunocompromised_population"),
			OutcomeOverviewManufacturer:                isFieldTrue(r, "outcome_overview_manufacturer"),
			OutcomeOverviewDoseDescription:             isFieldTrue(r, "outcome_overview_dose_description"),
			OutcomeOverviewSchedule:                    isFieldTrue(r, "outcome_overview_schedule"),
			OutcomeOverviewTimeFrameWeeks:              isFieldTrue(r, "outcome_overview_time_frame_weeks"),
			OutcomeOverviewConfidenceInterval:          isFieldTrue(r, "outcome_overview_confidence_interval"),
			OutcomeOverviewPercentResponders:           isFieldTrue(r, "outcome_overview_percent_responders"),
		}


		rc := clinicaltrials.NewResultsContext(r.Context() , t.db)

		csvData, err := rc.GetExport(sf, ff)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment;filename=wisspar_data_export.csv")
		w.Write(csvData.Bytes())
	}
}

func isFieldTrue(r *http.Request, str string) bool {
	if strings.ToLower(r.URL.Query().Get("default")) == "true" {
		switch str {
		case "location_continent",
		   	 "study_eligibility_standard_age_list",
		   	 "clinical_trial_study_id",
		   	 "clinical_trial_phase",
		   	 "outcome_overview_time_frame",
		   	 "outcome_overview_assay",
		   	 "outcome_overview_dose_number",
		   	 "outcome_overview_serotype",
		   	 "outcome_overview_value",
		   	 "outcome_overview_vaccine",
		   	 "outcome_overview_dose_description",
		   	 "outcome_overview_schedule",
		   	 "outcome_overview_time_frame_weeks":
			return true
		}
	}
	return strings.ToLower(r.URL.Query().Get(str)) == "true"
}

// GetOutcomeOverviewList for manual input
func GetOutcomeOverviewList(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		rc := clinicaltrials.NewResultsContext(r.Context() , t.db)

		study, err := rc.GetTrial(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		useCase := r.URL.Query().Get("use_case")
		if useCase == "" {
			render.Render(w, r, ErrInvalidRequest(fmt.Errorf("use case not set")))
			return
		}

		render.JSON(w, r, rc.GetOutcomeOverviewList(study, useCase))
	}
}

// insertStudy for manual input
func insertStudy(t Context) http.HandlerFunc {
	type req struct {
		Data clinicaltrials.Study `json:"data"`
		OutcomeMeasures []clinicaltrials.OutcomeMeasureCustom `json:"outcome_measures"`
		OutcomeOverviews []clinicaltrials.OutcomeOverview `json:"outcomes"`
		Metadata []clinicaltrials.Metadata `json:"metadata"`
		UseCase string `json:"use_case"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := clinicaltrials.NewResultsContext(r.Context() , t.db)

		if err := rc.InsertAll(req.Data, req.OutcomeMeasures, req.OutcomeOverviews, req.Metadata, req.UseCase); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		type resp struct {
			Status string `json:"status"`
		}

		render.JSON(w, r, resp{Status: "Ok"})
	}
}

// InsertStudy for manual input
func InsertCustomOutcomeMeasure(t Context) http.HandlerFunc {
	type req struct {
		OutcomeOverviews []clinicaltrials.OutcomeMeasureCustom `json:"outcome_measures"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := clinicaltrials.NewResultsContext(r.Context() , t.db)

		if err := rc.InsertCustomOutcomeMeasures(req.OutcomeOverviews, id); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		type resp struct {
			Status string `json:"status"`
		}

		render.JSON(w, r, resp{Status: "Ok"})
	}
}

func GetOutcomeMeasures(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		rc := clinicaltrials.NewResultsContext(r.Context() , t.db)

		items, err := rc.GetOutcomeMeasures(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.JSON(w, r, items)
	}
}

func DeleteOutcomeMeasure(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "omid")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := clinicaltrials.NewResultsContext(r.Context() , t.db)
		err = rc.RemoveOutcomeMeasure(idInt)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func list(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := authn.AccountIDFromContext(r)
		tasks, err := t.db.Task.Query().Where(task.Owner(userID)).All(r.Context())
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.JSON(w, r, tasks)
	}
}

func create(t Context) http.HandlerFunc {
	type req struct {
		Text string `json:"text"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(req)
		userID := authn.AccountIDFromContext(r)
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		newTask, err := t.db.Task.Create().
			SetID(shortuuid.New()).
			SetStatus(task.StatusInprogress).
			SetOwner(userID).
			SetText(req.Text).
			Save(r.Context())
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}
		render.JSON(w, r, newTask)
	}
}

func updateStatus(t Context) http.HandlerFunc {
	type req struct {
		Status string `json:"status"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(req)
		id := chi.URLParam(r, "id")

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		updatedTask, err := t.db.Task.
			UpdateOneID(id).
			SetUpdatedAt(time.Now()).
			SetStatus(task.Status(req.Status)).
			Save(r.Context())
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}
		render.JSON(w, r, updatedTask)
	}
}

func updateText(t Context) http.HandlerFunc {
	type req struct {
		Text string `json:"text"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(req)
		id := chi.URLParam(r, "id")

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		updatedTask, err := t.db.Task.
			UpdateOneID(id).
			SetUpdatedAt(time.Now()).
			SetText(req.Text).
			Save(r.Context())
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}
		render.JSON(w, r, updatedTask)
	}
}

func delete(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		err := t.db.Task.DeleteOneID(id).Exec(r.Context())
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func GetVaccine(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rc := vaccine.NewDatabaseContext(r.Context() , t.db)
		render.JSON(w, r, rc.Get())
	}
}

func UpdateVaccine(t Context) http.HandlerFunc {
	type req struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := vaccine.NewDatabaseContext(r.Context() , t.db)
		err = rc.Edit(vaccine.Vaccine{
			ID: idInt,
			Name: req.Name,
		})

		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func CreateVaccine(t Context) http.HandlerFunc {
	type req struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := vaccine.NewDatabaseContext(r.Context() , t.db)
		if err := rc.Create(req.Name); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func DeleteVaccine(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := vaccine.NewDatabaseContext(r.Context() , t.db)
		err = rc.Delete(idInt)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func GetManufacturer(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rc := manufacturer.NewDatabaseContext(r.Context() , t.db)
		render.JSON(w, r, rc.Get())
	}
}

func UpdateManufacturer(t Context) http.HandlerFunc {
	type req struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := manufacturer.NewDatabaseContext(r.Context() , t.db)
		err = rc.Edit(manufacturer.Manufacturer{
			ID: idInt,
			Name: req.Name,
		})

		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func CreateManufacturer(t Context) http.HandlerFunc {
	type req struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := manufacturer.NewDatabaseContext(r.Context() , t.db)
		if err := rc.Create(req.Name); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func DeleteManufacturer(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := manufacturer.NewDatabaseContext(r.Context() , t.db)
		err = rc.Delete(idInt)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func GetMetadata(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		useCase := r.URL.Query().Get("use_case")
		if useCase == "" {
			render.Render(w, r, ErrInvalidRequest(fmt.Errorf("use case not set")))
			return
		}

		rc := metadata.NewDatabaseContext(r.Context() , t.db)
		render.JSON(w, r, rc.Get(id, useCase))
	}
}

func GetICG(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rc := immunocompromisedgroups.NewDatabaseContext(r.Context() , t.db)
		render.JSON(w, r, rc.Get())
	}
}

func UpdateICG(t Context) http.HandlerFunc {
	type req struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := immunocompromisedgroups.NewDatabaseContext(r.Context() , t.db)
		err = rc.Edit(immunocompromisedgroups.ImmunocompromisedGroups{
			ID: idInt,
			Name: req.Name,
		})

		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func CreateICG(t Context) http.HandlerFunc {
	type req struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := immunocompromisedgroups.NewDatabaseContext(r.Context() , t.db)
		if err := rc.Create(req.Name); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func DeleteICG(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := immunocompromisedgroups.NewDatabaseContext(r.Context() , t.db)
		err = rc.Delete(idInt)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func GetUseCase(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rc := usecase.NewDatabaseContext(r.Context() , t.db)
		render.JSON(w, r, rc.Get())
	}
}

func UpdateUseCase(t Context) http.HandlerFunc {
	type req struct {
		Name string `json:"name"`
		Description string `json:"description"`
		Code string `json:"code"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := usecase.NewDatabaseContext(r.Context() , t.db)
		err = rc.Edit(usecase.UseCase{
			ID: idInt,
			Name: req.Name,
		})

		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func CreateUseCase(t Context) http.HandlerFunc {
	type req struct {
		Name string `json:"name"`
		Description string `json:"description"`
		Code string `json:"code"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := usecase.NewDatabaseContext(r.Context() , t.db)
		if err := rc.Create(req.Name, req.Description, req.Code); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func DeleteUseCase(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := usecase.NewDatabaseContext(r.Context() , t.db)
		err = rc.Delete(idInt)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func GetResponsibleParties(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trial := clinicaltrials.NewResultsContext(r.Context(), t.db)

		rp, err := trial.GetResponsibleParties()
		if err != nil {
			render.Render(w, r, ErrInternal(err))
		}

		render.JSON(w, r, rp)
	}
}

func GetDistinctManufacturers(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		results := []string{}
		res, err := t.db.Manufacturer.Query().All(r.Context())
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		for _, man := range res { results = append(results, man.ManufacturerName) }

		render.Status(r, http.StatusOK)
		render.JSON(w, r, results)
	}
}

func GetVsTrials(t Context) http.HandlerFunc {
	type req struct {
		IsOr          bool     `json:"is_or"`
		AgeList       []string `json:"age_list"`
		Vaccines      []string `json:"vaccines"`
		Schedules      []string `json:"schedules"`
		Manufacturers []string `json:"manufacturers"`
		Continents    []string `json:"continents"`
		ImmunocompromisedGroups    []string `json:"immunocompromised_groups"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := dashboards.NewResultsContext(r.Context(), t.db)
		results, err := rc.GetClinicalTrial(req.IsOr, req.AgeList, req.Vaccines, req.Schedules, req.Manufacturers, req.Continents, req.ImmunocompromisedGroups)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, results)
	}
}


func GetDoseDescription(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rc := dosedescription.NewDatabaseContext(r.Context() , t.db)
		render.JSON(w, r, rc.Get())
	}
}

func UpdateDoseDescription(t Context) http.HandlerFunc {
	type req struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := dosedescription.NewDatabaseContext(r.Context() , t.db)
		err = rc.Edit(dosedescription.DoseDescription{
			ID: idInt,
			Name: req.Name,
		})

		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func CreateDoseDescription(t Context) http.HandlerFunc {
	type req struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := dosedescription.NewDatabaseContext(r.Context() , t.db)
		if err := rc.Create(req.Name); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func DeleteDoseDescription(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := dosedescription.NewDatabaseContext(r.Context() , t.db)
		err = rc.Delete(idInt)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func GetSchedule(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rc := schedule.NewDatabaseContext(r.Context() , t.db)
		render.JSON(w, r, rc.Get())
	}
}

func UpdateSchedule(t Context) http.HandlerFunc {
	type req struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := schedule.NewDatabaseContext(r.Context() , t.db)
		err = rc.Edit(schedule.Schedule{
			ID: idInt,
			Name: req.Name,
		})

		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func CreateSchedule(t Context) http.HandlerFunc {
	type req struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(req)
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := schedule.NewDatabaseContext(r.Context() , t.db)
		if err := rc.Create(req.Name); err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}

func DeleteSchedule(t Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		rc := schedule.NewDatabaseContext(r.Context() , t.db)
		err = rc.Delete(idInt)
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	}
}
