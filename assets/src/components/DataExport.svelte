<script>
    import {onMount} from "svelte";

    // Field select

    let location_facility = false;
    let location_city = false;
    let location_country = false;
    let location_country_code = false;
    let location_continent = true;
    let location_continent_code = false;
    let study_eligibility_criteria = false;
    let study_eligibility_healthy_volunteers = false;
    let study_eligibility_gender = false;
    let study_eligibility_minimum_age = false;
    let study_eligibility_maximum_age = false;
    let study_eligibility_standard_age_list = true;
    let study_eligibility_ethnicity = false;
    let clinical_trial_study_name = false;
    let clinical_trial_study_id = true;
    let clinical_trial_sponsor = true;
    let clinical_trial_responsible_party = false;
    let clinical_trial_study_type = false;
    let clinical_trial_phase = true;
    let clinical_trial_actual_enrollment = false;
    let clinical_trial_allocation = false;
    let clinical_trial_intervention_model = false;
    let clinical_trial_masking = false;
    let clinical_trial_primary_purpose = false;
    let clinical_trial_official_title = false;
    let clinical_trial_actual_study_start_date = false;
    let clinical_trial_actual_primary_completion_date = false;
    let clinical_trial_actual_study_completion_date = false;
    let outcome_measure_title = false;
    let outcome_measure_type = false;
    let outcome_measure_description = false;
    let outcome_measure_population_description = false;
    let outcome_measure_reporting_status = false;
    let outcome_measure_anticipated_posting_date = false;
    let outcome_measure_param_type = false;
    let outcome_measure_dispersion_type = false;
    let outcome_measure_unit_of_measure = false;
    let outcome_measure_calculate_pct = false;
    let outcome_measure_time_frame = false;
    let outcome_measure_type_units_analyzed = false;
    let outcome_measure_denom_units_selected = false;
    let outcome_overview_title = false;
    let outcome_overview_id = false;
    let outcome_overview_description = false;
    let outcome_overview_time_frame = true;
    let outcome_overview_assay = true;
    let outcome_overview_dose_number = true;
    let outcome_overview_participants = false;
    let outcome_overview_serotype = true;
    let outcome_overview_value = true;
    let outcome_overview_upper_limit = true;
    let outcome_overview_lower_limit = true;
    let outcome_overview_ratio = false;
    let outcome_overview_vaccine = true;
    let outcome_overview_immunocompromised_population = false;
    let outcome_overview_manufacturer = false;
    let outcome_overview_dose_description = true;
    let outcome_overview_schedule = true;
    let outcome_overview_time_frame_weeks = true;
    let outcome_overview_confidence_interval = false;
    let outcome_overview_percent_responders = false;



    $: downloadDataCheck = !location_facility
        && !location_city
        && !location_country
        && !location_country_code
        && !location_continent
        && !location_continent_code
        && !study_eligibility_criteria
        && !study_eligibility_healthy_volunteers
        && !study_eligibility_gender
        && !study_eligibility_minimum_age
        && !study_eligibility_maximum_age
        && !study_eligibility_standard_age_list
        && !study_eligibility_ethnicity
        && !clinical_trial_study_name
        && !clinical_trial_study_id
        && !clinical_trial_sponsor
        && !clinical_trial_responsible_party
        && !clinical_trial_study_type
        && !clinical_trial_phase
        && !clinical_trial_actual_enrollment
        && !clinical_trial_allocation
        && !clinical_trial_intervention_model
        && !clinical_trial_masking
        && !clinical_trial_primary_purpose
        && !clinical_trial_official_title
        && !clinical_trial_actual_study_start_date
        && !clinical_trial_actual_primary_completion_date
        && !clinical_trial_actual_study_completion_date
        && !outcome_measure_title
        && !outcome_measure_type
        && !outcome_measure_description
        && !outcome_measure_population_description
        && !outcome_measure_reporting_status
        && !outcome_measure_anticipated_posting_date
        && !outcome_measure_param_type
        && !outcome_measure_dispersion_type
        && !outcome_measure_unit_of_measure
        && !outcome_measure_calculate_pct
        && !outcome_measure_time_frame
        && !outcome_measure_type_units_analyzed
        && !outcome_measure_denom_units_selected
        && !outcome_overview_title
        && !outcome_overview_id
        && !outcome_overview_description
        && !outcome_overview_time_frame
        && !outcome_overview_assay
        && !outcome_overview_dose_number
        && !outcome_overview_participants
        && !outcome_overview_serotype
        && !outcome_overview_value
        && !outcome_overview_upper_limit
        && !outcome_overview_lower_limit
        && !outcome_overview_ratio
        && !outcome_overview_vaccine
        && !outcome_overview_immunocompromised_population
        && !outcome_overview_manufacturer
        && !outcome_overview_dose_description
        && !outcome_overview_schedule
        && !outcome_overview_time_frame_weeks
        && !outcome_overview_confidence_interval
        && !outcome_overview_percent_responders;

    // Search filters

    let income_groups = [];
    let parent_array_separator = ",";
    let child_array_separator = " ";
    let vaccine = [];
    let phase = [];
    let std_age_list = [];
    let country_codes = [];
    let continent_codes = [];
    let responsible_party = [];
    let ethnicity_race = [];
    let immunocompromised_populations = [];

    const apiBase = "/export-options/"

    // Static content
    let responsible_parties = [];
    const responsiblePartiesAPI = apiBase + 'responsible-parties/';

    let immunocompromised_groups = [];
    const immunocompromisedGroupsAPI = apiBase + 'icg';

    let vaccines = [];
    const vaccinesAPI = apiBase + 'vaccine';

    let use_cases = [];
    let selected_use_case = "";
    const useCaseAPI = apiBase + 'usecase';

    const dataExportAPI = apiBase + 'data-export/';

    onMount(async () => {
        const res = await fetch(vaccinesAPI);
        vaccines = await res.json();

        const resICG = await fetch(immunocompromisedGroupsAPI);
        immunocompromised_groups = await resICG.json();

        const resUseCase = await fetch(useCaseAPI);
        use_cases = await resUseCase.json();

        const resResponsibleParties = await fetch(responsiblePartiesAPI);
        responsible_parties = await resResponsibleParties.json();
    });

    async function downloadData() {
        const res = await fetch(dataExportAPI, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                search_options: {
                    use_case: selected_use_case,
                    income_groups: income_groups,
                    parent_array_separator: parent_array_separator,
                    child_array_separator: child_array_separator,
                    vaccine: vaccine,
                    phase: phase,
                    std_age_list: std_age_list,
                    country_codes: country_codes,
                    continent_codes: continent_codes,
                    responsible_party: responsible_party,
                    ethnicity_race: ethnicity_race,
                    immunocompromised_populations: immunocompromised_populations,
                },
                field_options: {
                    location_facility: (String(location_facility) === 'true'),
                    location_city: (String(location_city) === 'true'),
                    location_country: (String(location_country) === 'true'),
                    location_country_code: (String(location_country_code) === 'true'),
                    location_continent: (String(location_continent) === 'true'),
                    location_continent_code: (String(location_continent_code) === 'true'),
                    study_eligibility_criteria: (String(study_eligibility_criteria) === 'true'),
                    study_eligibility_healthy_volunteers: (String(study_eligibility_healthy_volunteers) === 'true'),
                    study_eligibility_gender: (String(study_eligibility_gender) === 'true'),
                    study_eligibility_minimum_age: (String(study_eligibility_minimum_age) === 'true'),
                    study_eligibility_maximum_age: (String(study_eligibility_maximum_age) === 'true'),
                    study_eligibility_standard_age_list: (String(study_eligibility_standard_age_list) === 'true'),
                    study_eligibility_ethnicity: (String(study_eligibility_ethnicity) === 'true'),
                    clinical_trial_study_name: (String(clinical_trial_study_name) === 'true'),
                    clinical_trial_study_id: (String(clinical_trial_study_id) === 'true'),
                    clinical_trial_sponsor: (String(clinical_trial_sponsor) === 'true'),
                    clinical_trial_responsible_party: (String(clinical_trial_responsible_party) === 'true'),
                    clinical_trial_study_type: (String(clinical_trial_study_type) === 'true'),
                    clinical_trial_phase: (String(clinical_trial_phase) === 'true'),
                    clinical_trial_actual_enrollment: (String(clinical_trial_actual_enrollment) === 'true'),
                    clinical_trial_allocation: (String(clinical_trial_allocation) === 'true'),
                    clinical_trial_intervention_model: (String(clinical_trial_intervention_model) === 'true'),
                    clinical_trial_masking: (String(clinical_trial_masking) === 'true'),
                    clinical_trial_primary_purpose: (String(clinical_trial_primary_purpose) === 'true'),
                    clinical_trial_official_title: (String(clinical_trial_official_title) === 'true'),
                    clinical_trial_actual_study_start_date: (String(clinical_trial_actual_study_start_date) === 'true'),
                    clinical_trial_actual_primary_completion_date: (String(clinical_trial_actual_primary_completion_date) === 'true'),
                    clinical_trial_actual_study_completion_date: (String(clinical_trial_actual_study_completion_date) === 'true'),
                    outcome_measure_type: (String(outcome_measure_type) === 'true'),
                    outcome_measure_title: (String(outcome_measure_title) === 'true'),
                    outcome_measure_description: (String(outcome_measure_description) === 'true'),
                    outcome_measure_population_description: (String(outcome_measure_population_description) === 'true'),
                    outcome_measure_reporting_status: (String(outcome_measure_reporting_status) === 'true'),
                    outcome_measure_anticipated_posting_date: (String(outcome_measure_anticipated_posting_date) === 'true'),
                    outcome_measure_param_type: (String(outcome_measure_param_type) === 'true'),
                    outcome_measure_dispersion_type: (String(outcome_measure_dispersion_type) === 'true'),
                    outcome_measure_unit_of_measure: (String(outcome_measure_unit_of_measure) === 'true'),
                    outcome_measure_calculate_pct: (String(outcome_measure_calculate_pct) === 'true'),
                    outcome_measure_time_frame: (String(outcome_measure_time_frame) === 'true'),
                    outcome_measure_type_units_analyzed: (String(outcome_measure_type_units_analyzed) === 'true'),
                    outcome_measure_denom_units_selected: (String(outcome_measure_denom_units_selected) === 'true'),
                    outcome_overview_title: (String(outcome_overview_title) === 'true'),
                    outcome_overview_id: (String(outcome_overview_id) === 'true'),
                    outcome_overview_description: (String(outcome_overview_description) === 'true'),
                    outcome_overview_time_frame: (String(outcome_overview_time_frame) === 'true'),
                    outcome_overview_assay: (String(outcome_overview_assay) === 'true'),
                    outcome_overview_dose_number: (String(outcome_overview_dose_number) === 'true'),
                    outcome_overview_participants: (String(outcome_overview_participants) === 'true'),
                    outcome_overview_serotype: (String(outcome_overview_serotype) === 'true'),
                    outcome_overview_value: (String(outcome_overview_value) === 'true'),
                    outcome_overview_upper_limit: (String(outcome_overview_upper_limit) === 'true'),
                    outcome_overview_lower_limit: (String(outcome_overview_lower_limit) === 'true'),
                    outcome_overview_ratio: (String(outcome_overview_ratio) === 'true'),
                    outcome_overview_vaccine: (String(outcome_overview_vaccine) === 'true'),
                    outcome_overview_immunocompromised_population: (String(outcome_overview_immunocompromised_population) === 'true'),
                    outcome_overview_manufacturer: (String(outcome_overview_manufacturer) === 'true'),
                    outcome_overview_dose_description: (String(outcome_overview_dose_description) === 'true'),
                    outcome_overview_schedule: (String(outcome_overview_schedule) === 'true'),
                    outcome_overview_time_frame_weeks: (String(outcome_overview_time_frame_weeks) === 'true'),
                    outcome_overview_confidence_interval: (String(outcome_overview_confidence_interval) === 'true'),
                    outcome_overview_percent_responders: (String(outcome_overview_percent_responders) === 'true'),
                },
            })
        });
        const data = await res.text();
        console.log(data)

        const hiddenElement = document.createElement('a');
              hiddenElement.href = 'data:text/csv;charset=utf-8,' + encodeURI(data);
              hiddenElement.target = '_blank';
              hiddenElement.download = 'wisspar_export.csv';
              hiddenElement.click();
    }

    const next_stage = () => {
        if (stage === 3) {
            return
        }
        progressBar.handleProgress(+1)
        stage++;
    }
    const previous_stage = () => {
        if (stage === 1) {
            return
        }
        progressBar.handleProgress(-1)
        stage--;
    }

    let stage = 1;
    import ProgressBar from './progress/Bar.svelte';
    let steps = ['', '', ''], currentActive = 1, progressBar;

</script>

<main class="is-fluid">
    <h1 class="has-text-centered title">Immunogenicity Data Export</h1>
    <h2 class="has-text-centered subtitle">Data Exporting Process</h2>
    <div class="columns">
        <div class="column" style="text-align: center" class:selected-stage="{stage === 1}" ><b>Step 1) Choose Use Case <i>(*required)</i></b>
            <p>Use cases have been curated on WISSPAR for different purposes. Please select the one that fits your needs.</p></div>
        <div class="column" style="text-align: center" class:selected-stage="{stage === 2}">
            <b>Step 2) Apply Search Filters <i>(optional)</i></b>
            <p>Applying search filters to your data will search for it in the outcome data, however, it will not filter for only that data field.
                As an example if you are trying to search for a particular country, other studies that took place in multiple locations including
                the one you selected will be included.</p>
        </div>
        <div class="column" style="text-align: center" class:selected-stage="{stage === 3}"><b>Step 3) Select Data Fields <i>(*required)</i></b>
            <p>Tick any field you wish for your analysis, just keep in mind the more fields you add may increase the time to create the data.</p></div>
    </div>
    <ProgressBar style="text-align: center" {steps} bind:currentActive bind:this={progressBar}/>
    {#if stage === 1}
    <div class="columns">
        <div class="column" style="text-align: center">
            <h3>Choose Use Case</h3>
            <label for="use-case-select">Use cases are sets of outcome overviews that are curated for a specific purpose.</label><br /><br />
            <select id="use-case-select" name="use-case-select" bind:value={selected_use_case}>
                <option selected="selected" disabled="disabled">Select use case</option>
                {#each use_cases as use_case}
                    <option value="{use_case.Code}">{use_case.Name} - {use_case.Description}</option>
                {/each}
            </select>
        </div>
    </div>
    <div class="columns">
        <div class="column">
            <button on:click={next_stage} disabled="{selected_use_case === ''}" class="button is-pulled-right">Next</button>
        </div>
    </div>
    {/if}
    {#if stage === 2}
    <div class="columns">
        <div class="column" style="text-align: center">
            <h3>Apply Search Filters</h3>
        </div>
    </div>
    <div class="columns">
        <div class="column">
            <b>Income group filter</b><br />
            <label><input type=checkbox bind:group={income_groups} name="income_groups" value="High income"> High income</label><br />
            <label><input type=checkbox bind:group={income_groups} name="income_groups" value="Upper middle income"> Upper middle income</label><br />
            <label><input type=checkbox bind:group={income_groups} name="income_groups" value="Lower middle income"> Lower middle income</label><br />
            <label><input type=checkbox bind:group={income_groups} name="income_groups" value="Low income"> Low income</label><br />
            <label><input type=checkbox bind:group={income_groups} name="income_groups" value="Not classified"> Not classified</label>
        </div>
        <div class="column">
            <b>Vaccine filter</b><br />
            {#each vaccines as v (v.ID)}
                <label><input type=checkbox bind:group={vaccine} name="vaccine" value={v.Name}> {v.Name}</label><br />
            {/each}
        </div>
        <div class="column">
            <b>Immunocompromised groups filter</b><br />
            {#each immunocompromised_groups as ig (ig.ID)}
                <label><input type=checkbox bind:group={immunocompromised_populations} name="vaccine" value={ig.Name}> {ig.Name}</label> <br />
            {/each}
            <hr />
            <label><input type=checkbox bind:group={immunocompromised_populations} name="vaccine" value=""> Healthy</label> <br />
        </div>
    </div>
    <div class="columns">
        <div class="column">
            <b>Trial phase</b><br />
            <label><input type=checkbox bind:group={phase} name="phase_list" value="Phase 1"> Phase 1</label><br />
            <label><input type=checkbox bind:group={phase} name="phase_list" value="Phase 2"> Phase 2</label><br />
            <label><input type=checkbox bind:group={phase} name="phase_list" value="Phase 3"> Phase 3</label><br />
            <label><input type=checkbox bind:group={phase} name="phase_list" value="Phase 4"> Phase 4</label>
        </div>
        <div class="column">
            <b>Age list</b><br />
            <label><input type=checkbox bind:group={std_age_list} name="std_age_list" value="Child"> Child</label> <br />
            <label><input type=checkbox bind:group={std_age_list} name="std_age_list" value="Adult"> Adult</label> <br />
            <label><input type=checkbox bind:group={std_age_list} name="std_age_list" value="Older Adult"> Older Adult</label>
        </div>
        <div class="column">
            <b>Responsible Parties</b><br />
            {#each responsible_parties as rp_item}
                <label><input type=checkbox bind:group={responsible_party} name="responsible_party" value={rp_item}> {rp_item}</label> <br />
            {/each}
        </div>
    </div>
        <div class="columns">
            <div class="column">
                <button on:click={previous_stage} class="button is-pulled-left">Previous</button>
                <button on:click={next_stage} class="button is-pulled-right">Next</button>
            </div>
        </div>
    {/if}
    {#if stage === 3}
    <div class="columns">
        <div class="column" style="text-align: center">
            <h3>Select Data Fields</h3>
        </div>
    </div>
    <div class="columns">
        <div class="column">
            <b>Clinical Trial - <i>Default data</i></b><br />
            <label><input type=checkbox bind:checked={clinical_trial_study_name} name="clinical_trial_study_name"> Study Name</label><br />
            <label><input type=checkbox bind:checked={clinical_trial_study_id} name="clinical_trial_study_name"> Study ID (NCT Trial Number)</label><br />
            <label><input type=checkbox bind:checked={clinical_trial_sponsor} name="clinical_trial_sponsor"> Sponsor</label><br />
            <label><input type=checkbox bind:checked={clinical_trial_responsible_party} name="clinical_trial_responsible_party"> Responsible Party</label><br />
            <label><input type=checkbox bind:checked={clinical_trial_study_type} name="clinical_trial_study_type"> Study Type</label><br />
            <label><input type=checkbox bind:checked={clinical_trial_phase} name="clinical_trial_phase"> Phase</label><br />
            <label><input type=checkbox bind:checked={clinical_trial_actual_enrollment} name="clinical_trial_actual_enrollment"> Actual Enrollment</label><br />
            <label><input type=checkbox bind:checked={clinical_trial_allocation} name="clinical_trial_allocation"> Allocation</label><br />
            <label><input type=checkbox bind:checked={clinical_trial_intervention_model} name="clinical_trial_intervention_model"> Intervention Model</label><br />
            <label><input type=checkbox bind:checked={clinical_trial_masking} name="clinical_trial_masking"> Masking</label><br />
            <label><input type=checkbox bind:checked={clinical_trial_primary_purpose} name="clinical_trial_primary_purpose"> Primary Purpose</label><br />
            <label><input type=checkbox bind:checked={clinical_trial_official_title} name="clinical_trial_official_title"> Official Title</label><br />
            <label><input type=checkbox bind:checked={clinical_trial_actual_study_start_date} name="clinical_trial_actual_study_start_date"> Actual Study Start Date</label><br />
            <label><input type=checkbox bind:checked={clinical_trial_actual_primary_completion_date} name="clinical_trial_actual_primary_completion_date"> Actual Primary Completion Date</label><br />
            <label><input type=checkbox bind:checked={clinical_trial_actual_study_completion_date} name="clinical_trial_actual_study_completion_date"> Actual Study Completion Date</label><br /><br />
            <b>Study Eligibility - <i>Default data</i></b><br />
            <label><input type=checkbox bind:checked={study_eligibility_criteria} name="study_eligibility_criteria"> Eligibility Criteria</label><br />
            <label><input type=checkbox bind:checked={study_eligibility_healthy_volunteers} name="study_eligibility_healthy_volunteers"> Healthy Volunteers</label><br />
            <label><input type=checkbox bind:checked={study_eligibility_gender} name="study_eligibility_gender"> Gender</label><br />
            <label><input type=checkbox bind:checked={study_eligibility_minimum_age} name="study_eligibility_minimum_age"> Minimum Age</label><br />
            <label><input type=checkbox bind:checked={study_eligibility_maximum_age} name="study_eligibility_maximum_age"> Maximum Age</label><br />
            <label><input type=checkbox bind:checked={study_eligibility_standard_age_list} name="study_eligibility_standard_age_list"> Standard Age List</label><br />
            <label><input type=checkbox bind:checked={study_eligibility_ethnicity} name="study_eligibility_ethnicity"> Ethnicity</label><br /><br />
        </div>
        <div class="column">
            <b>Study Location - <i>Curated data</i></b><br />
            <label><input type=checkbox bind:checked={location_facility} name="location_facility"> Facilities</label> <br />
            <label><input type=checkbox bind:checked={location_city} name="location_city"> Cities</label> <br />
            <label><input type=checkbox bind:checked={location_country} name="location_country"> Countries</label> <br />
            <label><input type=checkbox bind:checked={location_country_code} name="location_country_code"> Country Codes</label> <br />
            <label><input type=checkbox bind:checked={location_continent} name="location_continent"> Continents</label> <br />
            <label><input type=checkbox bind:checked={location_continent_code} name="location_continent_code"> Continent Codes</label><br /><br />
            <b>Outcome Measure - <i>Default data</i></b><br />
            <label><input type=checkbox bind:checked={outcome_measure_title} name="outcome_measure_title"> Title</label> <br />
            <label><input type=checkbox bind:checked={outcome_measure_type} name="outcome_measure_type"> Type</label> <br />
            <label><input type=checkbox bind:checked={outcome_measure_description} name="outcome_measure_description"> Description</label> <br />
            <label><input type=checkbox bind:checked={outcome_measure_population_description} name="outcome_measure_population_description"> Population Description</label> <br />
            <label><input type=checkbox bind:checked={outcome_measure_reporting_status} name="outcome_measure_reporting_status"> Reporting Status</label> <br />
            <label><input type=checkbox bind:checked={outcome_measure_anticipated_posting_date} name="outcome_measure_anticipated_posting_date"> Anticipated Posting Date</label> <br />
            <label><input type=checkbox bind:checked={outcome_measure_param_type} name="outcome_measure_param_type"> Param Type</label> <br />
            <label><input type=checkbox bind:checked={outcome_measure_dispersion_type} name="outcome_measure_dispersion_type"> Dispersion Type</label> <br />
            <label><input type=checkbox bind:checked={outcome_measure_unit_of_measure} name="outcome_measure_unit_of_measure"> Unit Of Measure</label> <br />
            <label><input type=checkbox bind:checked={outcome_measure_calculate_pct} name="outcome_measure_calculate_pct"> Calculate PCT</label> <br />
            <label><input type=checkbox bind:checked={outcome_measure_time_frame} name="outcome_measure_time_frame"> Time Frame</label> <br />
            <label><input type=checkbox bind:checked={outcome_measure_type_units_analyzed} name="outcome_measure_type_units_analyzed"> Type Units Analyzed</label> <br />
            <label><input type=checkbox bind:checked={outcome_measure_denom_units_selected} name="outcome_measure_denom_units_selected"> Denom Units Selected</label>
        </div>
        <div class="column">
            <b>Outcome Overview - <i>Curated data</i></b><br />
            <label><input type=checkbox bind:checked={outcome_overview_title} name="outcome_overview_title"> Title</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_id} name="outcome_overview_id"> ID</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_description} name="outcome_overview_description"> Description</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_time_frame} name="outcome_overview_time_frame"> Time Frame</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_time_frame_weeks} name="outcome_overview_time_frame_weeks"> Time Frame Weeks</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_assay} name="outcome_overview_assay"> Assay</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_dose_number} name="outcome_overview_dose_number"> Dose Number</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_participants} name="outcome_overview_participants"> Participants</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_serotype} name="outcome_overview_serotype"> Serotype</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_value} name="outcome_overview_value"> Value</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_upper_limit} name="outcome_overview_upper_limit"> Upper Limit</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_lower_limit} name="outcome_overview_lower_limit"> Lower Limit</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_ratio} name="outcome_overview_ratio"> Ratio</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_vaccine} name="outcome_overview_vaccine"> Vaccine</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_immunocompromised_population} name="outcome_overview_immunocompromised_population"> Immunocompromised Populations</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_manufacturer} name="outcome_overview_manufacturer"> Manufacturer</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_dose_description} name="outcome_overview_manufacturer"> Dose Description</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_schedule} name="outcome_overview_manufacturer"> Schedule</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_confidence_interval} name="outcome_overview_confidence_interval"> Confidence Interval</label><br />
            <label><input type=checkbox bind:checked={outcome_overview_percent_responders} name="outcome_overview_percent_responders"> Percent Responders</label>
        </div>
    </div>
    <div class="columns">
        <div class="column">
            <button class="button is-pulled-left" on:click={previous_stage}>Previous</button>
            <button class="button is-pulled-right" disabled="{downloadDataCheck}" on:click={() => downloadData()}>Download Data</button>
        </div>
    </div>
    {/if}
</main>
