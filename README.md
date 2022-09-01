# WISSPAR - Project

##About Project
The Worldwide Index of Serotype Specific Pneumococcal Antibody Responses (WISSPAR) is a centralized platform housing data on immunogenicity from clinical trials of pneumococcal vaccines.  The data on WISSPAR are primarily curated from outcomes tables from clinicaltrials.gov and are made available in a searchable format that can be readily used for downstream analyses. Fields included in the database include product, manufacturers, dosing schedule, age group, and geographic region. You can use the embedded tools to generate customizable visualizations, export data for further analyses, or just browse through clinical trial results.

## Technical Documentation
This is a golang/svelte MPA project that uses the Ent framework for the custom schema.

## /app
This folder houses the golang code which acts as the backend http server and rest api.

### /app/gen/models
This folder is auto generated from `/app/schema` using the `go generate ./app/generator` command. DO NOT EDIT THIS FOLDER.

## /assets
This folder houses the svelte assets for the project to be compiled into the public folder.

# Live Data
This link will generate the most up to date CSV: https://wisspar.com/export-options/data-export/?use_case=pcv_antibodies&default=true

## Search Options
Here are a list of search options you can use:
 - <b>use_case</b> (example value=pcv_anitbodies) single value
 - <b>income_groups</b> (Values = High income, Upper middle income, Lower middle income, Low income, Not classified) multi value
 - <b>vaccine</b> (Values are set in admin) <i>multi value</i>
 - <b>phase</b> (Values = Phase 1, Phase 2, Phase 3, Phase 4) <i>multi value</i> 
 - <b>std_age_list</b> (Values = Child, Adult, Older Adult)  <i>multi value</i>
 - <b>country_codes</b> (ISO standard 2 letter country code)  <i>multi value</i>
 - <b>continent_codes</b> (ISO standard 2 letter continent code)  <i>multi value</i>
 - <b>responsible_party</b> (Values come from clinicaltrials.gov)  <i>multi value</i>
 - <b>ethnicity_race</b> (Values come from clinicaltrials.gov)  <i>multi value</i>
 - <b>immunocompromised_populations</b> (Values are set in admin) <i>multi value</i>


These two are used to combine field arrays in whatever way you need (you won't need to set it unless you have a specific reason)
 - <b>parent_array_separator</b> = (Default set to ',') single value
 - <b>child_array_separator</b> = (Default set to ' ') single value
 
<i>How to use multi value in URL request: .../?vaccine=Vaccine 1&vaccine=Vaccine  2</i>

## Field Options
Setting the default option to true will give you the following fields:

 - <b>location_continent</b>
 - <b>study_eligibility_standard_age_list</b>
 - <b>clinical_trial_study_id</b>
 - <b>clinical_trial_phase</b>
 - <b>outcome_overview_time_frame</b>
 - <b>outcome_overview_assay</b>
 - <b>outcome_overview_dose_number</b>
 - <b>outcome_overview_serotype</b>
 - <b>outcome_overview_value</b>
 - <b>outcome_overview_vaccine</b>
 - <b>outcome_overview_dose_description</b>
 - <b>outcome_overview_schedule</b>
 - <b>outcome_overview_time_frame_weeks</b>

However the whole list of fields you can set are below just use in URL request like this: "field_name=true" :
 - <b>location_facility</b>
 - <b>location_city</b>
 - <b>location_country</b>
 - <b>location_country_code</b>
 - <b>location_continent</b>
 - <b>location_continent_code</b>
 
 
 - <b>study_eligibility_criteria</b>
 - <b>study_eligibility_healthy_volunteers</b>
 - <b>study_eligibility_gender</b>
 - <b>study_eligibility_minimum_age</b>
 - <b>study_eligibility_maximum_age</b>
 - <b>study_eligibility_standard_age_list</b>
 - <b>study_eligibility_ethnicity</b>
 
 
 - <b>clinical_trial_study_name</b>
 - <b>clinical_trial_study_id</b>
 - <b>clinical_trial_sponsor</b>
 - <b>clinical_trial_responsible_party</b>
 - <b>clinical_trial_study_type</b>
 - <b>clinical_trial_phase</b>
 - <b>clinical_trial_actual_enrollment</b>
 - <b>clinical_trial_allocation</b>
 - <b>clinical_trial_intervention_model</b>
 - <b>clinical_trial_masking</b>
 - <b>clinical_trial_primary_purpose</b>
 - <b>clinical_trial_official_title</b>
 - <b>clinical_trial_actual_study_start_date</b>
 - <b>clinical_trial_actual_primary_completion_date</b>
 - <b>clinical_trial_actual_study_completion_date</b>
 
 
 - <b>outcome_measure_type</b>
 - <b>outcome_measure_title</b>
 - <b>outcome_measure_description</b>
 - <b>outcome_measure_population_description</b>
 - <b>outcome_measure_reporting_status</b>
 - <b>outcome_measure_anticipated_posting_date</b>
 - <b>outcome_measure_param_type</b>
 - <b>outcome_measure_dispersion_type</b>
 - <b>outcome_measure_unit_of_measure</b>
 - <b>outcome_measure_calculate_pct</b>
 - <b>outcome_measure_time_frame</b>
 - <b>outcome_measure_type_units_analyzed</b>
 - <b>outcome_measure_denom_units_selected</b>
 
 
 - <b>outcome_overview_title</b>
 - <b>outcome_overview_id</b>
 - <b>outcome_overview_description</b>
 - <b>outcome_overview_time_frame</b>
 - <b>outcome_overview_assay</b>
 - <b>outcome_overview_dose_number</b>
 - <b>outcome_overview_participants</b>
 - <b>outcome_overview_serotype</b>
 - <b>outcome_overview_value</b>
 - <b>outcome_overview_upper_limit</b>
 - <b>outcome_overview_lower_limit</b>
 - <b>outcome_overview_ratio</b>
 - <b>outcome_overview_vaccine</b>
 - <b>outcome_overview_immunocompromised_population</b>
 - <b>outcome_overview_manufacturer</b>
 - <b>outcome_overview_dose_description</b>
 - <b>outcome_overview_schedule</b>
 - <b>outcome_overview_time_frame_weeks</b>
 - <b>outcome_overview_confidence_interval</b>
 - <b>outcome_overview_percent_responders</b>

## Local Setup

### Prerequisites:
 - Docker
 - NPM
 - Yarn
 - Golang

In one terminal window run this command:
```
cd assets && yarn watch 
```
Then run this one in another terminal window:
```
docker-compose up app
```

You can access the website locally through: http://localhost:8023
