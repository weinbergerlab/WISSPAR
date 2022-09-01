<script>
    import Tabs from "./Tabs.svelte";
    import StudyDetails from "./tabs/StudyDetails.svelte";
    import Results from "./tabs/Results.svelte";
    import TabularView from "./tabs/TabularView.svelte";
    import {onMount} from "svelte";
    import { Confirm } from 'svelte-confirm';
    import AccordionAlt from "./tabs/accordian/AccordianAlt.svelte"

    let showStudy = false
    let data = {};
    let outcomes = new Map();
    let outcomeMeasures = [];
    let metadata = [];
    let outcome_overview_list = [];
    const trialsAPI = '/api/trials/';
    let id = "NCT03197376";
    let selected_use_case = "";

    let vaccines = [];
    let manufacturers = [];
    let doseDescriptions = [];
    let schedules = [];
    let immunocompromisedGroups = [];
    let use_cases = [];
    let use_cases_map = new Map();
    const vaccinesAPI = '/api/vaccine';
    const manufacturersAPI = '/api/manufacturer';
    const doseDescriptionsAPI = '/api/dose-description';
    const schedulesAPI = '/api/schedule';
    const icgAPI = '/api/icg';
    const useCaseAPI = '/api/usecase';

    let outcome_measure_title = "";
    let outcome_measure_description = "";
    let outcome_measure_type = "";
    let outcome_measure_population_description = "";
    let outcome_measure_reporting_status = "";
    let outcome_measure_param_type = "";
    let outcome_measure_dispersion_type = "";
    let outcome_measure_unit_of_measure = "";
    let outcome_measure_time_frame = "";
    let outcome_measure_anticipated_posting_date = "";
    let outcome_measure_calculate_pct = "";
    let outcome_measure_denom_units_selected = "";
    let outcome_measure_type_units_analyzed = "";

    let items = [
        {
            label: "Study Details",
            value: 1,
            component: StudyDetails
        },
        {
            label: "Tabular View",
            value: 2,
            component: TabularView
        },
        {
            label: "Results",
            value: 3,
            component: Results
        },
    ];

    let currentMeasurementTitle = "";
    let currentOverviewTitle = "";
    let hasClosed = true;
    let outcome_serotype = "";
    let outcome_value = 0;
    let outcome_lower = 0;
    let outcome_upper = 0;
    let outcome_ratio = "";
    let outcome_dose_number = 0;
    let outcome_time_frame_weeks = 0;
    let outcome_vaccine = "";
    let outcome_manufacturer = "";
    let outcome_dose_description = "";
    let outcome_schedule = "";
    let outcome_immunocompromised_population = "";
    let outcome_confidence_interval = "";
    let outcome_percent_responders = "";
    let outcome_assay = "";
    let reload = false;
    let stage = 1;

    let metadata_tag_name = "";
    let metadata_tag_value = "";

    let serotypecol = true;
    let valuecol = true;
    let llcol = false;
    let ulcol = false;
    let ratiocol = true;
    let dosecol = true;
    let vaccinecol = true;
    let tfwcol = false;
    let manufacturercol = false;
    let dosedescriptioncol = false;
    let schedulecol = false;
    let icgcol = false;
    let cicol = false;
    let prcol = false;
    let assaycol = false;

    onMount(async () => {
        const resVaccine = await fetch(vaccinesAPI);
        vaccines = await resVaccine.json();

        const resManufacturer = await fetch(manufacturersAPI);
        manufacturers = await resManufacturer.json();

        const resDoseDescriptions = await fetch(doseDescriptionsAPI);
        doseDescriptions = await resDoseDescriptions.json();

        const resSchedules = await fetch(schedulesAPI);
        schedules = await resSchedules.json();

        const resICG = await fetch(icgAPI);
        immunocompromisedGroups = await resICG.json();

        const resUseCase = await fetch(useCaseAPI);
        use_cases = await resUseCase.json();
        console.log(use_cases)

        for(let i = 0; i < use_cases.length; i++) {
            use_cases_map.set(use_cases[i].Code, use_cases[i])
        }
    })

    async function resetAll() {
        const studyURL = trialsAPI + id
        const outcomesURL = studyURL + "/outcome-overview-list/?use_case=" + selected_use_case
        const outcomesAPI = await fetch(outcomesURL, {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
        })
        reload = true
        actions = [];

        const outcomesJSON = await outcomesAPI.json()
        outcomes = new Map(Object.entries(outcomesJSON));

        for (const [key, value] of outcomes.entries()) {
            outcomes.set(key, new Map(Object.entries(value)))
        }
        outcomes = outcomes
        setupTrialData()
        reload = false
        await colInit()
    }

    function addToOutcomes(measure_title, overview_title, i) {
        const current_outcome = outcomes.get(measure_title).get(overview_title)[0]

        outcomes.get(measure_title).get(overview_title).splice(i, 0, {
            outcome_overview_title: current_outcome.outcome_overview_title,
            outcome_overview_description: current_outcome.outcome_overview_description,
            outcome_overview_time_frame: current_outcome.outcome_overview_time_frame,
            outcome_overview_assay: outcome_assay,
            outcome_overview_dose_number: outcome_dose_number,
            outcome_measure_title: current_outcome.outcome_measure_title,
            outcome_overview_id: current_outcome.outcome_overview_id,
            outcome_overview_participants: current_outcome.outcome_overview_participants,
            outcome_overview_serotype: outcome_serotype,
            outcome_overview_value: outcome_value,
            outcome_overview_upper_limit: outcome_upper,
            outcome_overview_lower_limit: outcome_lower,
            outcome_overview_group_id: current_outcome.outcome_overview_group_id,
            outcome_overview_ratio: outcome_ratio,
            outcome_overview_vaccine: outcome_vaccine,
            outcome_overview_time_frame_weeks: outcome_time_frame_weeks,
            outcome_overview_manufacturer: outcome_manufacturer,
            outcome_overview_dose_description: outcome_dose_description,
            outcome_overview_schedule: outcome_schedule,
            outcome_overview_immunocompromised_population: outcome_immunocompromised_population,
            outcome_overview_confidence_interval: outcome_confidence_interval,
            outcome_overview_percent_responders: outcome_percent_responders,
            outcome_overview_use_case: current_outcome.outcome_overview_use_case,
            enabled: false,
        });
        reload = true;
        outcomes = outcomes;
        itemToAction("single", measure_title, overview_title, i, true);
        colInit()
        reload = false;
        outcome_serotype = "";
        outcome_value = 0;
        outcome_lower = 0;
        outcome_upper = 0;
        outcome_ratio = "";
        outcome_dose_number = 0;
        outcome_vaccine = "";
        outcome_time_frame_weeks = 0;
        outcome_manufacturer = "";
        outcome_dose_description = "";
        outcome_schedule = "";
        outcome_immunocompromised_population = "";
        outcome_confidence_interval = "";
        outcome_percent_responders = "";
        outcome_assay = "";
    }

    function addToMetadata() {
        metadata.push({
            tag_name: metadata_tag_name,
            tag_value: metadata_tag_value,
            use_case: selected_use_case,
        });
        reload = true;
        metadata = metadata;
        reload = false;
        metadata_tag_name = "";
        metadata_tag_value = "";
    }

    function removeOutcome(measure_title, overview_title, i) {
        itemToAction("single", measure_title, overview_title, i, false);
        outcomes.get(measure_title).get(overview_title)[i].enabled = true;
        reload = true;
        outcomes = outcomes;
        reload = false;
    }

    function removeMetadata(i) {
        metadata.splice(i,1);
        reload = true;
        metadata = metadata;
        reload = false;
    }

    function removeAll(i, measure_title, overview_title) {
        itemToAction("group", measure_title, overview_title, i, false);
        for(let i = 0; i < outcomes.get(measure_title).get(overview_title).length; i++) {
            outcomes.get(measure_title).get(overview_title)[i].enabled = true
        }
        reload = true;
        outcomes = outcomes;
        selectItems.get(measure_title).set(overview_title, true);
        selectItems = selectItems;
        reload = false;
    }

    let actions = [];

    function saveAction(action) {
        actions.push(action)
    }

    function getAction() {
        return actions.pop()
    }

    function itemToAction(type, measure_title, overview_title, i, remove) {
        actions.push({
                type: type,
                i: i,
                measure_title: measure_title,
                outcome_title: overview_title,
                remove: remove
        })
    }

    function undo() {
        if (actions.length === 0) {return}
        let action = getAction();

        switch(action.type) {
            case "single":
                outcomes.get(action.measure_title).get(action.outcome_title)[action.i].enabled = action.remove
                break;
            case "group":
                selectItems.get(action.measure_title).set(action.outcome_title, action.remove)

                for(let i = 0; i < outcomes.get(action.measure_title).get(action.outcome_title).length; i++) {
                    outcomes.get(action.measure_title).get(action.outcome_title)[i].enabled = action.remove
                }
                break;
        }

        reload = true;
        selectItems = selectItems;
        outcomes = outcomes;
        colInit()
        reload = false;
    }

    const getTrial = async () => {
        if(!selected_use_case) {
            return
        }

        stage = 1.5
        const studyURL = trialsAPI + id

        const studyAPI = await fetch(studyURL, {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
        })

        const outcomesURL = studyURL + "/outcome-overview-list/?use_case=" + selected_use_case
        const outcomesAPI = await fetch(outcomesURL, {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
        })

        const metadataURL = studyURL + "/metadata/?use_case=" + selected_use_case
        const metadataAPI = await fetch(metadataURL, {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
        })

        const outcomeMeasuresURL = studyURL + "/outcome-measures/"
        const outcomeMeasuresAPI = await fetch(outcomeMeasuresURL, {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
        })

        if (studyAPI.status !== 200 && outcomesAPI.status !== 200) {
            stage = 1;
            return;
        }

        if(outcomeMeasuresAPI.status !== 200) {
            console.log(studyAPI.error())
        }
        outcomeMeasures = await outcomeMeasuresAPI.json();
        if (outcomeMeasures === null || outcomeMeasures === undefined) {
            outcomeMeasures = [];
        }

        console.log(outcomeMeasures);

        data = await studyAPI.json()
        const outcomesJSON = await outcomesAPI.json()
        outcomes = new Map(Object.entries(outcomesJSON));

        for (const [key, value] of outcomes.entries()) {
            outcomes.set(key, new Map(Object.entries(value)))
        }
        outcomes = outcomes

        metadata = await metadataAPI.json()
        showStudy = true
        stage = 2

        setupTrialData()
        await colInit()
    }

    async function colInit() {
        await new Promise(resolve => setTimeout(resolve, 500));
        if(!serotypecol       ) { hideCol("serotype-col"        ) }
        if(!valuecol          ) { hideCol("value-col"           ) }
        if(!llcol             ) { hideCol("ll-col"              ) }
        if(!ulcol             ) { hideCol("ul-col"              ) }
        if(!ratiocol          ) { hideCol("ratio-col"           ) }
        if(!dosecol           ) { hideCol("dose-col"            ) }
        if(!vaccinecol        ) { hideCol("vaccine-col"         ) }
        if(!tfwcol            ) { hideCol("tfw-col"             ) }
        if(!manufacturercol   ) { hideCol("manufacturer-col"    ) }
        if(!dosedescriptioncol) { hideCol("dose-description-col") }
        if(!schedulecol       ) { hideCol("schedule-col"        ) }
        if(!icgcol            ) { hideCol("icg-col"             ) }
        if(!cicol             ) { hideCol( "ci-col"             ) }
        if(!prcol             ) { hideCol( "pr-col"             ) }
        if(!assaycol          ) { hideCol( "assay-col"          ) }
        await new Promise(resolve => setTimeout(resolve, 500));
        if(!serotypecol       ) { hideCol("serotype-col"        ) }
        if(!valuecol          ) { hideCol("value-col"           ) }
        if(!llcol             ) { hideCol("ll-col"              ) }
        if(!ulcol             ) { hideCol("ul-col"              ) }
        if(!ratiocol          ) { hideCol("ratio-col"           ) }
        if(!dosecol           ) { hideCol("dose-col"            ) }
        if(!vaccinecol        ) { hideCol("vaccine-col"         ) }
        if(!tfwcol            ) { hideCol("tfw-col"             ) }
        if(!manufacturercol   ) { hideCol("manufacturer-col"    ) }
        if(!dosedescriptioncol) { hideCol("dose-description-col") }
        if(!schedulecol       ) { hideCol("schedule-col"        ) }
        if(!icgcol            ) { hideCol("icg-col"             ) }
        if(!cicol             ) { hideCol( "ci-col"             ) }
        if(!prcol             ) { hideCol( "pr-col"             ) }
        if(!assaycol          ) { hideCol( "assay-col"          ) }
    }

    const removeTrial = async () => {
        stage = 1;
        data = {};
        outcomes = new Map();
        metadata = [];
        showStudy = false;
    }

    let insertTrialResp = {}
    let insertTrialRespCode = 0;

    const insertTrial = async () => {
        stage = 3;

        let insertedOverviews = [];

        for (const [_, nested_overviews] of outcomes.entries()) {
            for (const [_, nested_overview_values] of nested_overviews.entries()) {
                for(let i = 0; i < nested_overview_values.length; i++) {
                    if(!nested_overview_values[i].enabled) {
                        nested_overview_values[i].outcome_overview_upper_limit = nested_overview_values[i].outcome_overview_upper_limit.toString();
                        nested_overview_values[i].outcome_overview_lower_limit = nested_overview_values[i].outcome_overview_lower_limit.toString();
                        nested_overview_values[i].outcome_overview_percent_responders  = parseFloat(nested_overview_values[i].outcome_overview_percent_responders);
                        if(Array.isArray(nested_overview_values[i].outcome_overview_immunocompromised_population)) {
                            nested_overview_values[i].outcome_overview_immunocompromised_population = nested_overview_values[i].outcome_overview_immunocompromised_population.toString()
                        }
                        insertedOverviews.push(nested_overview_values[i])
                    }
                }
            }
        }

        const trialAPI = await fetch(trialsAPI, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                data: data,
                outcome_measures: outcomeMeasures,
                outcomes: insertedOverviews,
                metadata: metadata,
                use_case: selected_use_case,
            })
        })

        insertTrialResp = await trialAPI.json()
        insertTrialRespCode = trialAPI.status;


        if (insertTrialResp.status === undefined || insertTrialResp.status === null) {
            insertTrialResp.status = "NOT OK"
        } else {
            data = {};
            outcomes = [];
            metadata = [];
            showStudy = false;
        }
    }

    function colHide(col) {
        switch (col) {
            case "serotype-col":         if(serotypecol    ) { showCol(col); return; } hideCol(col); break;
            case "value-col":            if(valuecol       ) { showCol(col); return; } hideCol(col); break;
            case "ll-col":               if(llcol          ) { showCol(col); return; } hideCol(col); break;
            case "ul-col":               if(ulcol          ) { showCol(col); return; } hideCol(col); break;
            case "ratio-col":            if(ratiocol       ) { showCol(col); return; } hideCol(col); break;
            case "dose-col":             if(dosecol        ) { showCol(col); return; } hideCol(col); break;
            case "vaccine-col":          if(vaccinecol     ) { showCol(col); return; } hideCol(col); break;
            case "tfw-col":              if(tfwcol         ) { showCol(col); return; } hideCol(col); break;
            case "manufacturer-col":     if(manufacturercol) { showCol(col); return; } hideCol(col); break;
            case "dose-description-col": if(dosedescriptioncol) { showCol(col); return; } hideCol(col); break;
            case "schedule-col":         if(schedulecol) { showCol(col); return; } hideCol(col); break;
            case "icg-col":              if(icgcol         ) { showCol(col); return; } hideCol(col); break;
            case "ci-col":               if(cicol          ) { showCol(col); return; } hideCol(col); break;
            case "pr-col":               if(prcol          ) { showCol(col); return; } hideCol(col); break;
            case "assay-col":            if(assaycol       ) { showCol(col); return; } hideCol(col); break;
        }
    }

    function showCol(col) {
        Array.from(document.querySelectorAll('.'+col)).forEach(function(colElement) {
            colElement.style.display = "";
        });
    }

    function hideCol(col) {
        Array.from(document.querySelectorAll('.'+col)).forEach(function(colElement) {
            colElement.style.display = "none";
        });
    }

    let filterShow = false;

    function hideShowFilter() {
        filterShow = !filterShow;
    }

    let trialDataShow = true;

    function hideShowTrialData() {
        trialDataShow = !trialDataShow;
    }

    let selectItems = new Map()
    let selectMeasureTitle = ""
    let selectOverviewTitle = ""

    let selectedOverview = ""

    function addToStage() {
        if(!selectedOverview) {
            return
        }

        let overview_ids = selectedOverview.split("-_-")

        let measure_title = overview_ids[0]
        let overview_title = overview_ids[1]

        selectItems.get(measure_title).set(overview_title, false)

        for(let i = 0; i < outcomes.get(measure_title).get(overview_title).length; i++) {
            if(outcomes.get(measure_title).get(overview_title)[i].custom) {
                continue;
            }
            outcomes.get(measure_title).get(overview_title)[i].enabled = false
        }

        selectedOverview = ""

        reload = true
        outcomes = outcomes
        selectItems = selectItems
        itemToAction("group", measure_title, overview_title, 0, true);
        colInit()
        reload = false
    }

    function addOutcomeMeasure() {
        outcomeMeasures.push({
            OutcomeMeasureTitle:outcome_measure_title,
            OutcomeMeasureDescription:outcome_measure_description,
            OutcomeMeasureType:outcome_measure_type,
            OutcomeMeasurePopulationDescription:outcome_measure_population_description,
            OutcomeMeasureReportingStatus:outcome_measure_reporting_status,
            OutcomeMeasureParamType:outcome_measure_param_type,
            OutcomeMeasureDispersionType:outcome_measure_dispersion_type,
            OutcomeMeasureUnitOfMeasure:outcome_measure_unit_of_measure,
            OutcomeMeasureTimeFrame:outcome_measure_time_frame,
            OutcomeMeasureAnticipatedPostingDate:outcome_measure_anticipated_posting_date,
            OutcomeMeasureCalculatePct:outcome_measure_calculate_pct,
            OutcomeMeasureDenomUnitsSelected:outcome_measure_denom_units_selected,
            OutcomeMeasureTypeUnitsAnalyzed:outcome_measure_type_units_analyzed,
        });

        outcome_measure_title = "";
        outcome_measure_description = "";
        outcome_measure_type = "";
        outcome_measure_population_description = "";
        outcome_measure_reporting_status = "";
        outcome_measure_param_type = "";
        outcome_measure_dispersion_type = "";
        outcome_measure_unit_of_measure = "";
        outcome_measure_time_frame = "";
        outcome_measure_anticipated_posting_date = "";
        outcome_measure_calculate_pct = "";
        outcome_measure_denom_units_selected = "";
        outcome_measure_type_units_analyzed = "";

        outcomeMeasures = outcomeMeasures;
    }

    function removeOutcomeMeasure(index, measure_title) {
        outcomeMeasures.splice(index, 1);
        selectItems.delete(measure_title);
        outcomeMeasures = outcomeMeasures;
        selectItems = selectItems;
    }
    let selected_outcome_measure = "";
    let selected_outcome_overview = "";
    let selected_outcome_overview_id = "";
    let selected_outcome_overview_group_id = "";
    let selected_outcome_overview_participants = "";

    function addOutcomeOverview() {
        let om = null;

        for(let i = 0; i < outcomeMeasures.length; i++) {
            if(outcomeMeasures[i].OutcomeMeasureTitle === selected_outcome_measure) {
                om = outcomeMeasures[i];
                break;
            }
        }

        if(om === null) {
            console.log("couldn't find outcome measure");
            console.log(selected_outcome_measure);
            console.log(outcomeMeasures);
            return;
        }

        let overview_example = {
            outcome_overview_title: selected_outcome_overview,
            outcome_overview_description: om.OutcomeMeasureDescription,
            outcome_overview_time_frame: om.OutcomeMeasureTimeFrame,
            outcome_measure_title: selected_outcome_measure,
            outcome_overview_id: selected_outcome_overview_id,
            outcome_overview_participants: selected_outcome_overview_participants,
            outcome_overview_group_id: selected_outcome_overview_group_id,
            outcome_overview_use_case: selected_use_case,
            enabled:true,
            custom: true,
        };

        if (selectItems.has(selected_outcome_measure)) {
            if (!selectItems.get(selected_outcome_measure).has(selected_outcome_overview)) {
                selectItems.get(selected_outcome_measure).set(selected_outcome_overview, true);
                outcomes.get(selected_outcome_measure).set(selected_outcome_overview, [overview_example]);
            }
        } else {
            let nested_overview_result = new Map();
            nested_overview_result.set(selected_outcome_overview, [overview_example]);
            outcomes.set(selected_outcome_measure, nested_overview_result);

            let nested_overviews_result = new Map();
            nested_overviews_result.set(selected_outcome_overview, true);
            selectItems.set(selected_outcome_measure, nested_overviews_result);
        }

        outcomes = outcomes;
        selectItems = selectItems;
        selected_outcome_measure = "";
        selected_outcome_overview = "";
        selected_outcome_overview_id = "";
        selected_outcome_overview_group_id = "";
        selected_outcome_overview_participants = "";
    }

    function setupTrialData() {
        for (const [measure_title, nested_overviews] of outcomes.entries()) {
            let nested_overviews_results = new Map()

            for (const [overview_title, nested_overview_values] of nested_overviews.entries()) {
                nested_overviews_results.set(overview_title, nested_overview_values[0].enabled)
            }

            selectItems.set(measure_title, nested_overviews_results)
        }
    }

    function updateAll(measure_title, overview_title, name) {
        let el = document.getElementById("all_"+name+"_"+idFilter(measure_title, overview_title))
        const val = el.value
        switch(name) {
            case "serotype":                    for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_serotype = val} break;
            case "value":                       for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_value = val} break;
            case "lower":                       for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_lower_limit = val} break;
            case "upper":                       for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_upper_limit = val} break;
            case "ratio":                       for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_ratio = val} break;
            case "dose_number":                 for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_dose_number = parseInt(val)} break;
            case "vaccine":                     for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_vaccine = val} break;
            case "time_frame_weeks":            for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_time_frame_weeks = parseInt(val)} break;
            case "manufacturer":                for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_manufacturer = val} break;
            case "dose_description":            for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_dose_description = val} break;
            case "schedule":                    for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_schedule = val} break;
            case "immunocompromised_population":for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_immunocompromised_population = getSelectValues(el)} break;
            case "confidence_interval":         for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_confidence_interval = val} break;
            case "percent_responders":          for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_percent_responders = parseFloat(val)} break;
            case "assay":                       for(let k = 0; k < outcomes.get(measure_title).get(overview_title).length; k++) {outcomes.get(measure_title).get(overview_title)[k].outcome_overview_assay = val} break;
        }

        outcomes = outcomes
    }

    function getSelectValues(select) {
        var result = [];
        var options = select && select.options;
        var opt;

        for (var i=0, iLen=options.length; i<iLen; i++) {
            opt = options[i];

            if (opt.selected) {
                result.push(opt.value || opt.text);
            }
        }
        return result;
    }

    function idFilter(measure_title, overview_title) {
        return measure_title.replace(/\s+/g, '_').toLowerCase() + "_" + overview_title.replace(/\s+/g, '_').toLowerCase()
    }

    function getShorterName(name) {
        const words = name.split(' ')
        const wordCount = words.length

        if(wordCount < 7) {
            return name;
        }

        return words.slice(0, 3).join(' ') + '...' + words.slice(-3);
    }

</script>

<main class="is-fluid">
    <h1 class="has-text-centered title">Clinical Trial - Insert Form</h1>
    <nav></nav>

    {#if stage === 1}
    <div id="stage-1">
        <div class="block has-text-centered">
            <p>Please enter the ID of the clinical trial you wish to insert (eg. <i>NCT03197376</i>)</p>
        </div>
        <form class="field has-addons" style="justify-content: center" on:submit|preventDefault={getTrial}>
            <div class="field">
                <div class="control">
                    <label><b>Clinical Trial ID</b></label>
                    <input bind:value={id} class="input" type="text" placeholder="Clinical Trial ID">
                </div>
            </div>
            <div class="field">
                <div class="control">
                    <label><b>Use Case</b></label>
                    <select bind:value={selected_use_case} class="input">
                        <option disabled="disabled" selected="selected">Not selected</option>
                        {#each use_cases || [] as use_case, i}
                            <option value="{use_case.Code}">{use_case.Name}</option>
                        {/each}
                    </select>
                </div>
            </div>

            <div class="control">
                <br />
                <button class="button is-primary">
                    <span class="icon is-small">
                      <i class="fas fa-eye"></i>
                    </span>
                </button>
            </div>
        </form>
        <br /><br />
        <div class="columns is-center is-centered">
            <div class="column is-half is-text is-centered is-center content" style="text-align: left">
                <h3 style="text-align: center" class="title is-3">How it works</h3>
                <h4 class="title is-4">1. Enter the Trial ID</h4>
                <p>Find the trial that you are looking to insert on clinicaltrials.gov or check the trials dashboard to see if it's already imported.</p>
                <h4 class="title is-4">2. Check the Information</h4>
                <p>Once you have pulled the trial up you will be presented with two panels</p>
                <h5 class="title is-5">2.1 Reviewing the trial information</h5>
                <p>The first panel on the left is all the trial information available to scan through. In order to find
                    specific information regarding the outcome measures go to Results -> Outcome Measures. You will then
                    see an outcome measure table that should correspond to the outcome measures on the right.
                </p>
                <img src="/static/images/left-panel.PNG" style="border:2px solid #333"/>
                <span style="text-align: center; width: 100%; display: block"><i>Left panel - Trial information</i></span>

                <h5 class="title is-5">2.2 Review/Enter Outcomes</h5>
                <p>As previously mentioned the right panel is for custom data input. The system has already preselected
                    the relevant outcome measures based on whether they are related to serotypes. Please check the values
                    against the corresponding data on the left. The dose is approximated based on the data so make sure
                    that especially is correct. You can delete sections and undo/reset at anytime. No data is committed
                    until you click insert all.
                </p>
                <img style="border:2px solid #333" src="/static/images/left-panel-information-edited.jpg" width="100%" />
                <span style="text-align: center; width: 100%; display: block"><i>Left panel - Outcome information</i></span>
                <br /><br />
                <img style="border:2px solid #333" src="/static/images/right-panel-information-edited.jpg" width="100%" />
                <span style="text-align: center; width: 100%; display: block"><i>Right panel - Outcome information</i></span>
                <h5 class="title is-5">2.3 Buttons</h5>
                <ul>
                    <li><b>Back</b> - Takes you back to this screen so you can insert a different trial.</li>
                    <li><b>Undo</b> - Undo previous actions including inserting an item, removing a single item or removing an entire group.</li>
                    <li><b>Reset</b> - Removes everything and resets all the data. This action cannot be undone.</li>
                    <li><b>Insert</b> - Will take you to the next stage where the data will be stored in the database.</li>
                </ul>
                <h5 class="title is-5">2.4 Add Metadata</h5>
                <p>Metadata can be anything you can think of that would need to be used to identify the trial. This can
                    be anything from vaccine, serotype or any other trial data. Add as many key value pairs as you need to.
                </p>
                <img style="border:2px solid #333" src="/static/images/metadata-panel.png" width="100%" />
                <span style="text-align: center; width: 100%; display: block"><i>Right panel - Metadata Section (At the bottom)</i></span>
                <h4 class="title is-4">3. Insert the Trial</h4>
                <p>Once you have clicked insert all please wait for confirmation that the trial has been inserted into
                    the database.
                </p>
            </div>
        </div>
    </div>
    {/if}

    {#if stage === 1.5}
        <div id="stage-1-5" class="content has-text-centered">
            <h2 class="title has-text-centered is-4">Loading</h2>
            <p class="has-text-centered">Getting Trial Information... Please wait.</p>
            <div class="loader">Loading...</div>
        </div>
    {/if}

    {#if stage === 2}
        <div id="stage-2">
            <div class="columns is-centered">
                {#if trialDataShow}<div class="column"><h2 class="subtitle has-text-centered" style="font-size: 27px; font-weight: bold;">Trial Data</h2></div>{/if}
                <div class="column">
                    <h2 class="subtitle has-text-centered" style="font-size: 27px; font-weight: bold;">Use Case Data: {use_cases_map.get(selected_use_case).Name}</h2>
                    <h4 class="subtitle has-text-centered" style="font-size: 18px; font-weight: normal;">Use Case Description: {use_cases_map.get(selected_use_case).Description}</h4>
                </div>
            </div>
            <div class="columns is-centered">
                {#if trialDataShow}
                    <div id="trial-data" class="column">
                        <div id="study">
                            {#if showStudy}
                                <Tabs {items} {data} />
                            {/if}
                        </div>
                    </div>
                {/if}
                <div class="column">
        <!--            <form>-->
                    <div class="block" >
                        <div class="columns">
                            <div class="column is-half" style="position:relative;">
                                <button class="button" on:click={() => hideShowFilter()}>Filters</button>
                                <button class="button" on:click={() => hideShowTrialData()}>{#if trialDataShow}Hide{:else }Show{/if} Trial Data</button>
                                {#if filterShow}
                                    <div class="filter-cols">
                                        Serotype <input type=checkbox bind:checked={serotypecol} on:change={() => colHide("serotype-col")}> <br />
                                        Value <input type=checkbox bind:checked={valuecol} on:change={() => colHide("value-col")}> <br />
                                        Lower Limit <input type=checkbox bind:checked={llcol} on:change={() => colHide("ll-col")}> <br />
                                        Upper Limit <input type=checkbox bind:checked={ulcol} on:change={() => colHide("ul-col")}> <br />
                                        Ratio <input type=checkbox bind:checked={ratiocol} on:change={() => colHide("ratio-col")}> <br />
                                        Dose <input type=checkbox bind:checked={dosecol} on:change={() => colHide("dose-col")}> <br />
                                        Vaccine <input type=checkbox bind:checked={vaccinecol} on:change={() => colHide("vaccine-col")}> <br />
                                        Time Frame Weeks <input type=checkbox bind:checked={tfwcol} on:change={() => colHide("tfw-col")}> <br />
                                        Manufacturer <input type=checkbox bind:checked={manufacturercol} on:change={() => colHide("manufacturer-col")}> <br />
                                        Dose Description <input type=checkbox bind:checked={dosedescriptioncol} on:change={() => colHide("dose-description-col")}> <br />
                                        Schedule <input type=checkbox bind:checked={schedulecol} on:change={() => colHide("schedule-col")}> <br />
                                        Immunocompromised Groups <input type=checkbox bind:checked={icgcol} on:change={() => colHide("icg-col")}> <br />
                                        Confidence Interval <input type=checkbox bind:checked={cicol} on:change={() => colHide("ci-col")}> <br />
                                        &percnt; Responders <input type=checkbox bind:checked={prcol} on:change={() => colHide("pr-col")}> <br />
                                        Assay <input type=checkbox bind:checked={assaycol} on:change={() => colHide("assay-col")}> <br />
                                    </div>
                                {/if}
                            </div>
                            <div id="buttonmenu" class="column is-half" style="text-align: right">
                                <button on:click={removeTrial} class="button">Back</button>
                                <button on:click={() => undo()} class="button">Undo</button>
                                <Confirm
                                        confirmTitle="Reset"
                                        cancelTitle="Cancel"
                                        let:confirm="{confirmThis}"
                                        style="text-align: center"
                                >
                                    <button on:click={() => confirmThis(resetAll)} class="button">Reset</button>
                                    <span slot="title" class="slot-title-reset">Reset the Outcomes</span>
                                    <span slot="description" class="slot-title-reset">You won't be able to revert this!</span>
                                </Confirm>
                                <button on:click={insertTrial} class="button">Insert</button>
                            </div>
                        </div>
                    </div>

                    <div class="block" style="border-bottom: 3px solid #777; margin-bottom: 0; padding-bottom: 14px;">
                        <div class="columns">
                            <div class="column is-four-fifths">
                                <select bind:value={selectedOverview} class="input">
                                    <option disabled="disabled" selected="selected">Not Selected</option>
                                    {#each [...selectItems] as [measure_title, overview_items]}
                                        <option disabled="disabled">{measure_title}</option>
                                        {#each [...overview_items] as [overview_title, enabled]}
                                            <option value="{measure_title}-_-{overview_title}" disabled="{!enabled}" >- {overview_title}</option>
                                        {/each}
                                    {/each}
                                </select>
                            </div>
                            <div class="column is-one-fifth"><button on:click={() => addToStage()} class="button" style="margin: 0 auto; display: block">Add Outcome</button></div>
                        </div>
                    </div>

                    <div class="set-scroll-box">
                        <h3 style="text-align: center; margin-bottom: 10px">Outcome Measures</h3>
                            {#if !reload}
                                {#each [...outcomes] as [measure_title, outcomeList] }
                                    <div class="outcome-measure-list">
                                    <table class="unwrapped" style="width: 100%">
                                        {#each [...outcomeList] as [overview_title, nested_outcome_list] }
                                            {#if !selectItems.get(measure_title).get(overview_title)}
                                            {#each nested_outcome_list || [] as outcome, i}
                                                {#if i === 0 }
                                                    <tr class="td-header">
                                                        <td colspan="14">
                                                            <table class="unwrapped" style="width: 100%">
                                                                <tr class="td-header">
                                                                    <td colspan="13">
                                                                        <h3 class="subtitle" style="font-size: 20px;">{outcome.outcome_measure_title}</h3>
                                                                        <h3 class="subtitle" style="font-size: 18px;">{outcome.outcome_overview_title}</h3>
                                                                        <p class="is-medium">{outcome.outcome_overview_description}</p>
                                                                    </td>
                                                                    <td colspan="1" style="text-align: center; vertical-align: middle;"><button class="button is-center" on:click={() => removeAll(i, outcome.outcome_measure_title, outcome.outcome_overview_title)}>Remove All</button></td>
                                                                </tr>
                                                            </table>
                                                        </td>
                                                    </tr>
                                                {/if}
                                                {#if i === 0 }
                                                    <tr>
                                                        <td colspan="1" class="serotype-col">Serotype</td>
                                                        <td colspan="1" class="value-col">Value</td>
                                                        <td colspan="1" class="ll-col">Lower Limit</td>
                                                        <td colspan="1" class="ul-col">Upper Limit</td>
                                                        <td colspan="1" class="ratio-col">Ratio</td>
                                                        <td colspan="1" class="dose-col">Dose</td>
                                                        <td colspan="1" class="vaccine-col">Vaccine</td>
                                                        <td colspan="1" class="tfw-col">Time Frame Weeks</td>
                                                        <td colspan="1" class="manufacturer-col">Manufacturer</td>
                                                        <td colspan="1" class="dose-description-col">Dose Description</td>
                                                        <td colspan="1" class="schedule-col">Schedule</td>
                                                        <td colspan="1" class="icg-col">Immunocompromised groups</td>
                                                        <td colspan="1" class="ci-col">Confidence Interval</td>
                                                        <td colspan="1" class="pr-col">&percnt; Responders</td>
                                                        <td colspan="1" class="assay-col">Assay</td>
                                                        <td colspan="1">Action</td>
                                                    </tr>
                                                    <tr style="background-color: #ddd">
                                                        <td class="serotype-col" colspan="1"><input type="text" class="input" bind:value={outcome_serotype} /></td>
                                                        <td class="value-col"    colspan="1"><input type="number" step="any" class="input" bind:value={outcome_value} /></td>
                                                        <td class="ll-col"       colspan="1"><input type="text" class="input" bind:value={outcome_lower} /></td>
                                                        <td class="ul-col"       colspan="1"><input type="text" class="input" bind:value={outcome_upper} /></td>
                                                        <td class="ratio-col"    colspan="1"><input type="text" class="input" bind:value={outcome_ratio} /></td>
                                                        <td class="dose-col"     colspan="1"><input type="number" class="input" bind:value={outcome_dose_number} /></td>
                                                        <td class="vaccine-col"  colspan="1">
                                                            <select class="input" bind:value={outcome_vaccine}>
                                                                <option selected="{(outcome_vaccine === '') ? 'selected' : ''}">Not selected</option>
                                                                {#each vaccines as vaccine}<option value={vaccine.Name}>{vaccine.Name}</option>{/each}
                                                            </select>
                                                        </td>
                                                        <td class="tfw-col" colspan="1"><input type="number" class="input" bind:value={outcome_time_frame_weeks} /></td>
                                                        <td class="manufacturer-col" colspan="1">
                                                            <select class="input" bind:value={outcome_manufacturer}>
                                                                <option selected="{(outcome_manufacturer === '') ? 'selected': ''}">Not selected</option>
                                                                {#each manufacturers as manufacturer}<option value={manufacturer.Name}>{manufacturer.Name}</option>{/each}
                                                            </select>
                                                        </td>
                                                        <td class="dose-description-col"  colspan="1">
                                                            <select class="input" bind:value={outcome_dose_description}>
                                                                <option selected="{(outcome_dose_description === '') ? 'selected' : ''}">Not selected</option>
                                                                {#each doseDescriptions as doseDescription}<option value={doseDescription.Name}>{doseDescription.Name}</option>{/each}
                                                            </select>
                                                        </td>
                                                        <td class="schedule-col"  colspan="1">
                                                            <select class="input" bind:value={outcome_schedule}>
                                                                <option selected="{(outcome_schedule === '') ? 'selected' : ''}">Not selected</option>
                                                                {#each schedules as schedule}<option value={schedule.Name}>{schedule.Name}</option>{/each}
                                                            </select>
                                                        </td>
                                                        <td class="icg-col" colspan="1">
                                                            <select multiple class="input" bind:value={outcome_immunocompromised_population}>
                                                                <option selected="{(outcome_immunocompromised_population === '') ? 'selected': ''}">Not selected</option>
                                                                {#each immunocompromisedGroups as icg}<option value={icg.Name}>{icg.Name}</option>{/each}
                                                            </select>
                                                        </td>
                                                        <td class="ci-col" colspan="1"><input type="text" class="input" bind:value={outcome_confidence_interval} /></td>
                                                        <td class="pr-col" colspan="1"><input type="number" class="input" bind:value={outcome_percent_responders} /></td>
                                                        <td class="assay-col" colspan="1">
                                                            <select class="input" bind:value={outcome_assay}>
                                                                <option selected="{(outcome_assay === '') ? 'selected': ''}">Not selected</option>
                                                                <option value="GMC">GMC</option>
                                                                <option value="OPA">OPA</option>
                                                            </select>
                                                        </td>
                                                        <td colspan="1"><button class="button is-center" on:click={() => addToOutcomes(measure_title, overview_title, i)}>+</button></td>
                                                    </tr>
                                                {/if}
                                                {#if !outcome.enabled}
                                                    <tr>
                                                    <td class="serotype-col" colspan="1"><b><input type="text"
                                                                  id="overview_serotype_{outcome.outcome_overview_id}_{i}"
                                                                  name="overview_serotype_{outcome.outcome_overview_id}_{i}"
                                                                  class="input"
                                                                  bind:value={outcome.outcome_overview_serotype} /></b></td>
                                                    <td class="value-col" colspan="1"><input type="number" step="any"
                                                                  id="overview_value_{outcome.outcome_overview_id}_{i}"
                                                                  name="overview_value_{outcome.outcome_overview_id}_{i}"
                                                                  class="input"
                                                                  bind:value={outcome.outcome_overview_value}  /></td>
                                                    <td class="ll-col" colspan="1"><input type="text"
                                                                  id="overview_lower_{outcome.outcome_overview_id}_{i}"
                                                                  name="overview_lower_{outcome.outcome_overview_id}_{i}"
                                                                  class="input"
                                                                  bind:value={outcome.outcome_overview_lower_limit}  />
                                                    </td>
                                                    <td class="ul-col" colspan="1"><input type="text"
                                                                    id="overview_upper_{outcome.outcome_overview_id}_{i}"
                                                                    name="overview_upper_{outcome.outcome_overview_id}_{i}"
                                                                    class="input"
                                                                    bind:value={outcome.outcome_overview_upper_limit}  />
                                                    </td>
                                                    <td class="ratio-col" colspan="1"><input type="text"
                                                                id="overview_ratio_{outcome.outcome_overview_id}_{i}"
                                                                name="overview_ratio_{outcome.outcome_overview_id}_{i}"
                                                                class="input"
                                                                bind:value={outcome.outcome_overview_ratio}  />
                                                    </td>
                                                    <td class="dose-col" colspan="1"><input type="number"
                                                                id="overview_dose_{outcome.outcome_overview_id}_{i}"
                                                                name="overview_dose_{outcome.outcome_overview_id}_{i}"
                                                                class="input"
                                                                bind:value={outcome.outcome_overview_dose_number}  />
                                                    </td>
                                                    <td class="vaccine-col" colspan="1">
                                                        <select
                                                                id="overview_vaccine_{outcome.outcome_overview_id}_{i}"
                                                                name="overview_vaccine_{outcome.outcome_overview_id}_{i}"
                                                                class="input"
                                                                bind:value={outcome.outcome_overview_vaccine}>
                                                            <option selected="{(outcome.outcome_overview_vaccine === '') ? 'selected' : ''}">Not selected</option>
                                                            {#each vaccines as vaccine}
                                                                <option value={vaccine.Name}>
                                                                    {vaccine.Name}
                                                                </option>
                                                            {/each}
                                                        </select>
                                                    </td>
                                                    <td class="tfw-col" colspan="1"><input type="number"
                                                                   id="overview_time_frame_weeks_{outcome.outcome_overview_id}_{i}"
                                                                   name="overview_dose_{outcome.outcome_overview_id}_{i}"
                                                                   class="input"
                                                                   bind:value={outcome.outcome_overview_time_frame_weeks}  />
                                                    </td>
                                                    <td class="manufacturer-col" colspan="1">
                                                        <select
                                                                id="overview_manufacturer_{outcome.outcome_overview_id}_{i}"
                                                                name="overview_manufacturer_{outcome.outcome_overview_id}_{i}"
                                                                class="input"
                                                                bind:value={outcome.outcome_overview_manufacturer}>
                                                            <option selected="{(outcome.outcome_overview_manufacturer === '') ? 'selected': ''}">Not selected</option>
                                                            {#each manufacturers as manufacturer}
                                                                <option value={manufacturer.Name}>
                                                                    {manufacturer.Name}
                                                                </option>
                                                            {/each}
                                                        </select>
                                                    </td>
                                                    <td class="dose-description-col" colspan="1">
                                                        <select
                                                                id="overview_dose_description_{outcome.outcome_overview_id}_{i}"
                                                                name="overview_dose_description_{outcome.outcome_overview_id}_{i}"
                                                                class="input"
                                                                bind:value={outcome.outcome_overview_dose_description}>
                                                            <option selected="{(outcome.outcome_overview_dose_description === '') ? 'selected': ''}">Not selected</option>
                                                            {#each doseDescriptions as doseDescription}
                                                                <option value={doseDescription.Name}>
                                                                    {doseDescription.Name}
                                                                </option>
                                                            {/each}
                                                        </select>
                                                    </td>
                                                    <td class="schedule-col" colspan="1">
                                                        <select
                                                                id="overview_schedule_{outcome.outcome_overview_id}_{i}"
                                                                name="overview_schedule_{outcome.outcome_overview_id}_{i}"
                                                                class="input"
                                                                bind:value={outcome.outcome_overview_schedule}>
                                                            <option selected="{(outcome.outcome_overview_schedule === '') ? 'selected': ''}">Not selected</option>
                                                            {#each schedules as schedule}
                                                                <option value={schedule.Name}>
                                                                    {schedule.Name}
                                                                </option>
                                                            {/each}
                                                        </select>
                                                    </td>
                                                    <td class="icg-col" colspan="1">
                                                        <select multiple
                                                                id="overview_immunocompromised_populations_{outcome.outcome_overview_id}_{i}"
                                                                name="overview_immunocompromised_populations_{outcome.outcome_overview_id}_{i}"
                                                                class="input" bind:value={outcome.outcome_overview_immunocompromised_population}>
                                                            <option selected="{(outcome.outcome_overview_immunocompromised_population === '') ? 'selected': ''}">Not selected</option>
                                                            {#each immunocompromisedGroups as icg}<option value={icg.Name}>{icg.Name}</option>{/each}
                                                        </select>
                                                    </td>
                                                    <td class="ci-col" colspan="1"><input type="text"
                                                               id="overview_confidence_interval_{outcome.outcome_overview_id}_{i}"
                                                               name="overview_confidence_interval_{outcome.outcome_overview_id}_{i}"
                                                               class="input"
                                                               bind:value={outcome.outcome_overview_confidence_interval}  />
                                                    </td>
                                                    <td class="pr-col" colspan="1"><input type="number"
                                                               id="overview_percent_responders_{outcome.outcome_overview_id}_{i}"
                                                               name="overview_percent_responders_{outcome.outcome_overview_id}_{i}"
                                                               class="input"
                                                               bind:value={outcome.outcome_overview_percent_responders}  />
                                                    </td>
                                                    <td class="assay-col" colspan="1">
                                                        <select
                                                                id="overview_assay_{outcome.outcome_overview_id}_{i}"
                                                                name="overview_assay_{outcome.outcome_overview_id}_{i}"
                                                                class="input" bind:value={outcome.outcome_overview_assay}>
                                                            <option selected="{(outcome.outcome_overview_assay === '') ? 'selected': ''}">Not selected</option>
                                                            <option value="GMC">GMC</option>
                                                            <option value="OPA">OPA</option>
                                                        </select>
                                                    </td>
                                                    <td colspan="1">
                                                        <button class="button is-center" on:click={() => removeOutcome(measure_title, overview_title, i)} >X</button>
                                                    </td>
                                </tr>
                                                {/if}
                                                {#if i === nested_outcome_list.length-1}
                                                    <tr style="background-color: #ccc">
                                                        <td class="serotype-col" colspan="1">
                                                            <div class="control"><input type="text" id="all_serotype_{idFilter(measure_title,overview_title)}" name="all_serotype_{idFilter(measure_title,overview_title)}" class="input" /></div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title, overview_title, "serotype")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td class="value-col"    colspan="1">
                                                            <div class="control"><input type="number" step="any" id="all_value_{idFilter(measure_title,overview_title)}" name="all_value_{idFilter(measure_title,overview_title)}" class="input" /></div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title, overview_title, "value")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td class="ll-col"       colspan="1">
                                                            <div class="control"><input type="text" id="all_lower_{idFilter(measure_title,overview_title)}" name="all_lower_{idFilter(measure_title,overview_title)}" class="input" /></div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title, overview_title, "lower")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td class="ul-col"       colspan="1">
                                                            <div class="control"><input type="text" id="all_upper_{idFilter(measure_title,overview_title)}" name="all_upper_{idFilter(measure_title,overview_title)}" class="input" /></div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title, overview_title, "upper")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td class="ratio-col"    colspan="1">
                                                            <div class="control"><input type="text" id="all_ratio_{idFilter(measure_title,overview_title)}" name="all_ratio_{idFilter(measure_title,overview_title)}" class="input" /></div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title, overview_title, "ratio")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td class="dose-col"     colspan="1">
                                                            <div class="control"><input type="number" id="all_dose_number_{idFilter(measure_title,overview_title)}" name="all_dose_number_{idFilter(measure_title,overview_title)}" class="input" /></div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title, overview_title, "dose_number")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td class="vaccine-col"  colspan="1">
                                                            <div class="control">
                                                                <select class="input" id="all_vaccine_{idFilter(measure_title,overview_title)}" name="all_vaccine_{idFilter(measure_title,overview_title)}" >
                                                                    <option selected="selected">Not selected</option>
                                                                    {#each vaccines as vaccine}<option value={vaccine.Name}>{vaccine.Name}</option>{/each}
                                                                </select>
                                                            </div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title,overview_title, "vaccine")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td class="tfw-col" colspan="1">
                                                            <div class="control"><input type="number" id="all_time_frame_weeks_{idFilter(measure_title,overview_title)}" name="all_time_frame_weeks_{idFilter(measure_title,overview_title)}" class="input" /></div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title, overview_title, "time_frame_weeks")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td class="manufacturer-col" colspan="1">
                                                            <div class="control"><select class="input" id="all_manufacturer_{idFilter(measure_title,overview_title)}" name="all_manufacturer_{idFilter(measure_title,overview_title)}">
                                                                <option selected="selected">Not selected</option>
                                                                {#each manufacturers as manufacturer}<option value={manufacturer.Name}>{manufacturer.Name}</option>{/each}
                                                            </select>
                                                            </div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title, overview_title, "manufacturer")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td class="dose-description-col" colspan="1">
                                                            <div class="control"><select class="input" id="all_dose_description_{idFilter(measure_title,overview_title)}" name="all_dose_description_{idFilter(measure_title,overview_title)}">
                                                                <option selected="selected">Not selected</option>
                                                                {#each doseDescriptions as doseDescription}<option value={doseDescription.Name}>{doseDescription.Name}</option>{/each}
                                                            </select>
                                                            </div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title, overview_title, "dose_description")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td class="schedule-col" colspan="1">
                                                            <div class="control"><select class="input" id="all_schedule_{idFilter(measure_title,overview_title)}" name="all_schedule_{idFilter(measure_title,overview_title)}">
                                                                <option selected="selected">Not selected</option>
                                                                {#each schedules as schedule}<option value={schedule.Name}>{schedule.Name}</option>{/each}
                                                            </select>
                                                            </div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title, overview_title, "schedule")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td class="icg-col" colspan="1">
                                                            <div class="control"><select multiple class="input" id="all_immunocompromised_population_{idFilter(measure_title,overview_title)}" name="all_immunocompromised_population_{idFilter(measure_title,overview_title)}">
                                                                <option selected="selected">Not selected</option>
                                                                {#each immunocompromisedGroups as icg}<option value={icg.Name}>{icg.Name}</option>{/each}
                                                            </select>
                                                            </div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title, overview_title, "immunocompromised_population")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td class="ci-col" colspan="1">
                                                            <div class="control"><input type="text" id="all_confidence_interval_{idFilter(measure_title,overview_title)}" name="all_confidence_interval_{idFilter(measure_title,overview_title)}" class="input" /></div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title, overview_title, "confidence_interval")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td class="pr-col" colspan="1">
                                                            <div class="control"><input type="number" id="all_percent_responders_{idFilter(measure_title,overview_title)}" name="all_percent_responders_{idFilter(measure_title,overview_title)}" class="input" /></div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title, overview_title, "percent_responders")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td class="assay-col" colspan="1">
                                                            <div class="control"><select class="input" id="all_assay_{idFilter(measure_title,overview_title)}" name="all_assay_{idFilter(measure_title,overview_title)}">
                                                                <option selected="selected">Not selected</option>
                                                                <option value="GMC">GMC</option>
                                                                <option value="OPA">OPA</option>
                                                            </select></div>
                                                            <div class="control"><button class="button is-primary full-button" on:click={() => updateAll(measure_title, overview_title, "assay")}><span class="icon is-small"><i class="fas fa-solid fa-globe"></i></span></button></div>
                                                        </td>
                                                        <td colspan="1"></td>
                                                    </tr>
                                                {/if}
                                            {/each}
                                            {/if}
                                        {/each}
                                    </table>
                                    </div>
                                {/each}
                            {/if}

                        <hr class="lbbig" style="background: #ccc; height: 5px;" />

                        <div class="columns">
                            <div class="column block is-full">
                                <h3 style="text-align: center; margin-bottom: 10px">Metadata</h3>

                                <div class="metaform" style="background-color: #eee">
                                <div class="columns">
                                    <div class="column is-one-fifth"></div>
                                    <div class="column is-three-fifths"><input type="text" class="input" placeholder="Tag Name" bind:value={metadata_tag_name} /></div>
                                    <div class="column is-one-fifths"></div>
                                </div>
                                <div class="columns">
                                    <div class="column is-one-fifth"></div>
                                    <div class="column is-three-fifths"><input type="text" class="input" placeholder="Tag Value" bind:value={metadata_tag_value} /></div>
                                    <div class="column is-one-fifth"></div>
                                </div>
                                <div class="columns">
                                    <div class="column is-one-fifth"></div>
                                    <div class="column is-three-fifths" style="text-align: center"><button class="input button is-center is-dark" on:click={() => addToMetadata()}>Add</button></div>
                                    <div class="column is-one-fifth"></div>
                                </div>
                                </div>

                                <table width="100%">
                                <tr style="border-bottom: 3px solid #ddd">
                                    <td>Tag Name</td>
                                    <td>Tag Value</td>
                                    <td>Action</td>
                                </tr>
                                {#each metadata || [] as metadatum, i}
                                    <tr style="border-bottom: 1px solid #ddd">
                                        <td style="vertical-align: middle">{metadatum.tag_name}</td>
                                        <td style="vertical-align: middle">{metadatum.tag_value}</td>
                                        <td style="vertical-align: middle"><button class="button is-center" on:click={() => removeMetadata(i)} >X</button></td>
                                    </tr>
                                {/each}
                            </table>
                            </div>
                        </div>

                        <div class="columns" style="max-width: 100%; overflow: hidden">
                            <div class="column block is-full">
                                <h3 style="text-align: center; margin-bottom: 10px">Outcome Measures</h3>
                                <hr />
                                {#each outcomeMeasures || [] as outcome_measure, i}
                                    <AccordionAlt>
                                    <span slot="head">{outcome_measure.OutcomeMeasureTitle}</span>
                                    <div slot="details">
                                        <div class="columns"><div class="column is-one-quarter">Title</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure.OutcomeMeasureTitle}" type="text" /></div></div>
                                        <div class="columns"><div class="column is-one-quarter">Description</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure.OutcomeMeasureDescription}" type="text" /></div></div>
                                        <div class="columns"><div class="column is-one-quarter">Type</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure.OutcomeMeasureType}" type="text" /></div></div>
                                        <div class="columns"><div class="column is-one-quarter">Population Description</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure.OutcomeMeasurePopulationDescription}" type="text" /></div></div>
                                        <div class="columns"><div class="column is-one-quarter">Reporting Status</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure.OutcomeMeasureReportingStatus}" type="text" /></div></div>
                                        <div class="columns"><div class="column is-one-quarter">Param Type</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure.OutcomeMeasureParamType}" type="text" /></div></div>
                                        <div class="columns"><div class="column is-one-quarter">Dispersion Type</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure.OutcomeMeasureDispersionType}" type="text" /></div></div>
                                        <div class="columns"><div class="column is-one-quarter">Unit Of Measure</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure.OutcomeMeasureUnitOfMeasure}" type="text" /></div></div>
                                        <div class="columns"><div class="column is-one-quarter">Time Frame</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure.OutcomeMeasureTimeFrame}" type="text" /></div></div>
                                        <div class="columns"><div class="column is-one-quarter">Anticipated Posting Date</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure.OutcomeMeasureAnticipatedPostingDate}" type="text" /></div></div>
                                        <div class="columns"><div class="column is-one-quarter">Calculate PCT</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure.OutcomeMeasureCalculatePct}" type="text" /></div></div>
                                        <div class="columns"><div class="column is-one-quarter">Denom Units Selected</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure.OutcomeMeasureDenomUnitsSelected}" type="text" /></div></div>
                                        <div class="columns"><div class="column is-one-quarter">Type Units Analyzed</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure.OutcomeMeasureTypeUnitsAnalyzed}" type="text" /></div></div>
                                        <div class="columns"><div class="column is-full"><button class="button" on:click={removeOutcomeMeasure(i, outcome_measure.OutcomeMeasureTitle)}>Remove</button></div></div>
                                    </div>
                                    </AccordionAlt>
                                {/each}
                                <br />
                                <hr />
                                <br />
                                <h4><b>Add New Outcome Measure</b></h4>
                                <br />
                                <p><i>Don't add an outcome measure with the same name as another one as it could cause one to be overwritten.</i></p>
                                <div class="columns"><div class="column is-one-quarter">Title</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure_title}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Description</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure_description}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Type</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure_type}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Population Description</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure_population_description}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Reporting Status</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure_reporting_status}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Param Type</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure_param_type}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Dispersion Type</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure_dispersion_type}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Unit Of Measure</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure_unit_of_measure}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Time Frame</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure_time_frame}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Anticipated Posting Date</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure_anticipated_posting_date}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Calculate PCT</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure_calculate_pct}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Denom Units Selected</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure_denom_units_selected}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Type Units Analyzed</div><div class="column is-three-quarters"><input class="input" bind:value="{outcome_measure_type_units_analyzed}" type="text" /></div></div>
                                <div class="columns"><div class="column is-full"><button class="button" on:click={() => addOutcomeMeasure()}>Add</button></div></div>

                                <hr />
                                <h4><b>Add New Outcome Overview Title</b></h4>
                                <br />
                                <div class="columns">
                                    <div class="column is-one-quarter">Outcome Measure Parent</div>
                                    <div class="column is-three-quarters">
                                        <select class="select" style="max-width: 100%;" bind:value="{selected_outcome_measure}" on:change="{() => selected_outcome_overview = ''}">
                                            <option value="">None Selected</option>
                                            {#each outcomeMeasures || [] as outcome_measure, i}
                                                <option value="{outcome_measure.OutcomeMeasureTitle}">{getShorterName(outcome_measure.OutcomeMeasureTitle)}</option>
                                            {/each}
                                        </select>
                                    </div>
                                </div>
                                <div class="columns"><div class="column is-one-quarter">Title</div><div class="column is-three-quarters"><input class="input" bind:value="{selected_outcome_overview}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Overview ID</div><div class="column is-three-quarters"><input class="input" bind:value="{selected_outcome_overview_id}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Overview Group ID</div><div class="column is-three-quarters"><input class="input" bind:value="{selected_outcome_overview_group_id}" type="text" /></div></div>
                                <div class="columns"><div class="column is-one-quarter">Overview Participants</div><div class="column is-three-quarters"><input class="input" bind:value="{selected_outcome_overview_participants}" type="text" /></div></div>
                                <button class="button" on:click={() => addOutcomeOverview()}>Add Overview</button>
                            </div>
                        </div>
        <!--            </form>-->
                    </div>
                </div>
            </div>
        </div>
    {/if}
    <!--[if lte IE 9]>
    <style>
        .path {stroke-dasharray: 0 !important;}
    </style>
    <![endif]-->
    {#if stage === 3}
        <div id="stage-3" class="content has-text-centered">
            {#if insertTrialResp.status === "Ok"}
                <h2 class="title has-text-centered is-4">Success Inserting Trial Insertion</h2>
                <p class="has-text-centered">Trial data has been inserted correctly, the reporting dashboard will be updated soon.</p>
                <svg version="1.1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 130.2 130.2">
                    <circle class="path circle" fill="none" stroke="#73AF55" stroke-width="6" stroke-miterlimit="10" cx="65.1" cy="65.1" r="62.1"/>
                    <polyline class="path check" fill="none" stroke="#73AF55" stroke-width="6" stroke-linecap="round" stroke-miterlimit="10" points="100.2,40.2 51.5,88.8 29.8,67.5 "/>
                </svg>
                <br /><br />
                <button on:click={removeTrial} class="button">Back To Start</button>
            {/if}
            {#if insertTrialRespCode === 0}
                <h2 class="title has-text-centered is-4">Inserting Trial Information</h2>
                <p class="has-text-centered">Please wait for confirmation of trial insertion.</p>
                <div class="loader">Loading...</div>
            {/if}
            {#if insertTrialResp.status !== "Ok" && insertTrialRespCode !== 0}
                <h2 class="title has-text-centered is-4">Error Inserting Trial Information</h2>
                <p class="has-text-centered">{JSON.stringify(insertTrialResp)}</p>
                <svg version="1.1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 130.2 130.2">
                    <circle class="path circle" fill="none" stroke="#D06079" stroke-width="6" stroke-miterlimit="10" cx="65.1" cy="65.1" r="62.1"/>
                    <line class="path line" fill="none" stroke="#D06079" stroke-width="6" stroke-linecap="round" stroke-miterlimit="10" x1="34.4" y1="37.9" x2="95.8" y2="92.3"/>
                    <line class="path line" fill="none" stroke="#D06079" stroke-width="6" stroke-linecap="round" stroke-miterlimit="10" x1="95.8" y1="38" x2="34.4" y2="92.2"/>
                </svg>
                <br /><br />
                <button on:click={removeTrial} class="button">Back To Start</button>
            {/if}
        </div>
    {/if}
</main>
