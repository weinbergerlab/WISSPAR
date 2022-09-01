<script>
    import {onMount} from "svelte";
    import Switch from './support/Switch.svelte';
    import { slide } from 'svelte/transition';
    import FusionCharts from "fusioncharts"; //Import the Fusioncharts library
    import Maps from "fusioncharts/fusioncharts.maps"; //Import the FusionMaps
    import World from "fusioncharts/maps/fusioncharts.world"; //Import World Map
    import FusionTheme from "fusioncharts/themes/fusioncharts.theme.fusion";
    import Charts from "fusioncharts/fusioncharts.charts";

    //Import the Svelte component
    import SvelteFC, { fcRoot } from "svelte-fusioncharts";

    // Always set FusionCharts as the first parameter
    fcRoot(FusionCharts, Charts, Maps, World, FusionTheme);

    let items = [];
    let map_items = []

    //STEP 2 : Preparing the Map Data
    let mapData = [
        {
            id: "NA",
            value: "0",
            showLabel: "1"
        },
        {
            id: "SA",
            value: "0",
            showLabel: "1"
        },
        {
            id: "AS",
            value: "0",
            showLabel: "1"
        },
        {
            id: "EU",
            value: "0",
            showLabel: "1"
        },
        {
            id: "AF",
            value: "0",
            showLabel: "1"
        },
        {
            id: "AU",
            value: "0",
            showLabel: "1"
        }
    ];

    let mapConfigs = {
        type: 'world', // Map type
        width: '100%', // Width of the chart
        height: '450', // Height of the chart
        dataFormat: 'json', // Data Type
        renderAt:'chart-container',  //container where the chart will render
        dataSource: {
            // Map Configuration
            "chart": {
                "caption": "World Map",
                "labelsepchar": ": ",
                "entityFillHoverColor": "#FFF9C4",
                "theme": "fusion",
                "showLegend": false,
            },
            // Source data as JSON --> id represents continents of the world.
            "data": mapData
        }
    };

    let is_or = false
    let age_list = [];
    let vaccines = [];
    let schedules = [];
    let manufacturers = [];
    let continents = [];
    let immunocompromised_groups = [];


    let sortBy = {col: "id", ascending: true};

    $: sort = (column) => {

        if (sortBy.col == column) {
            sortBy.ascending = !sortBy.ascending
        } else {
            sortBy.col = column
            sortBy.ascending = true
        }

        // Modifier to sorting function for ascending or descending
        let sortModifier = (sortBy.ascending) ? 1 : -1;

        let sort = (a, b) =>
            (a[column] < b[column])
                ? -1 * sortModifier
                : (a[column] > b[column])
                ? 1 * sortModifier
                : 0;

        items = items.sort(sort);
    }

    const trialsAPI = '/dashboard/clinical-trial/';
    const manufacturerAPI = '/dashboard/manufacturer/';
    const scheduleAPI = '/dashboard/schedule/';
    const apiBase = "/export-options/"
    const vaccinesAPI = apiBase + 'vaccine';
    const immunocompromisedGroupsAPI = apiBase + 'icg';

    let default_vaccines = [];
    let default_manufacturers = [];
    let default_schedules = [];
    let default_immunocompromised_groups = [];
    const default_age_list = ["Child","Adult","Older Adult"];
    const default_continent_list = ["North America","South America","Europe","Africa","Asia","Australia"];

    let use_cases = [];
    let use_case_data = [];
    const useCaseAPI = apiBase + 'usecase';

    let reload = false;
    onMount(async () => {
        const resVaccine = await fetch(vaccinesAPI);
        default_vaccines = await resVaccine.json();

        const resManufacturer = await fetch(manufacturerAPI);
        default_manufacturers = await resManufacturer.json();

        const resSchedule = await fetch(scheduleAPI);
        default_schedules = await resSchedule.json();

        const resImmunocompromisedGroups = await fetch(immunocompromisedGroupsAPI);
        default_immunocompromised_groups = await resImmunocompromisedGroups.json();

        const resUseCase = await fetch(useCaseAPI);
        use_cases = await resUseCase.json();
        for(let i = 0; i < use_cases.length; i++) {
            use_case_data.push({
                label:use_cases[i].Name,
                value:0
            })
        }

        await getTrial();
    })

    let chartConfigs = {
        type: "column2d", //Select the chart type
        width: '100%', //Set the width of the chart
        height: 450, //Set the height of the chart
        dataFormat: "json", //Set the input dataFormat to json
        dataSource: {
            chart: {
                caption: "Amount of Outcomes Curated",
                subCaption: "Split by Use Case",
                xAxisName: "Use Cases", //Assign a relevant name to your x-axis
                yAxisName: "Outcomes", //Assign a relevant name to your y-axis
                numberSuffix: "",
                theme: "fusion" //Apply a theme to your chart
            },
            //Include chartData from STEP 2
            data: use_case_data
        }
    };

    let showMap = false

    const getTrial = async () => {
        items = [];
        await fetch(trialsAPI, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                is_or: is_or,
                age_list: age_list,
                vaccines: vaccines,
                schedules: schedules,
                manufacturers: manufacturers,
                continents: continents,
                immunocompromised_groups: immunocompromised_groups,
            })
        }).then(r => r.json())
          .then(data => {
              for(let i = 0; i < data.length; i++) {
                  if(data[i].age_list !== undefined) {
                      data[i].age_list = data[i].age_list.replace('[', '').replace(']','').replaceAll(',',', ')
                  }
                  if(data[i].ethnicity !== undefined) {
                      data[i].ethnicity = data[i].ethnicity.replace('[', '').replace(']','').replaceAll(',',', ')
                  }
              }
              items = data
        });

        mapData[0].value = 0;
        mapData[1].value = 0;
        mapData[2].value = 0;
        mapData[3].value = 0;
        mapData[4].value = 0;
        mapData[5].value = 0;

        for(let i = 0; i < use_case_data.length; i++) {
            use_case_data[i].value = 0
        }

        if (items.length === 0) {
            showMap = false;

            mapConfigs = {
                type: 'world', // Map type
                width: '100%', // Width of the chart
                height: '450', // Height of the chart
                dataFormat: 'json', // Data Type
                renderAt:'chart-container',  //container where the chart will render
                dataSource: {
                    // Map Configuration
                    "chart": {
                        "caption": "World Map",
                        "labelsepchar": ": ",
                        "entityFillHoverColor": "#FFF9C4",
                        "theme": "fusion",
                        "showLegend": false,
                    },
                    // Source data as JSON --> id represents continents of the world.
                    "data": mapData
                }
            };

            chartConfigs = {
                type: "column2d", //Select the chart type
                width: '100%', //Set the width of the chart
                height: 450, //Set the height of the chart
                dataFormat: "json", //Set the input dataFormat to json
                dataSource: {
                    chart: {
                        caption: "Number of Outcomes Curated",
                        subCaption: "Split by Use Case",
                        xAxisName: "Use Cases", //Assign a relevant name to your x-axis
                        yAxisName: "Outcomes", //Assign a relevant name to your y-axis
                        numberSuffix: "",
                        theme: "fusion" //Apply a theme to your chart
                    },
                    //Include chartData from STEP 2
                    data: use_case_data
                }
            };

            showMap = true
            return
        }

        for(let i = 0; i < items.length; i++) {
            // Loop through entries for use cases
            if(items[i].outcomes_curated_per_use_case !== null && items[i].outcomes_curated_per_use_case !== undefined) {
                for (const key in items[i].outcomes_curated_per_use_case) {
                    for(let j = 0; j < use_case_data.length; j++) {
                        for(let k = 0; k < use_cases.length; k++) {
                            if(use_case_data[j].label === use_cases[k].Name && key === use_cases[k].Code) {
                                use_case_data[j].value = use_case_data[j].value+items[i].outcomes_curated_per_use_case[key]
                            }
                        }
                    }
                }
            }

            for(let j = 0; j < items[i].location_continents.length; j++) {
                if(items[i].location_continents[j].includes("North America")) { mapData[0].value = mapData[0].value+1 }
                if(items[i].location_continents[j].includes("South America")) { mapData[1].value = mapData[1].value+1 }
                if(items[i].location_continents[j].includes("Asia")) { mapData[2].value = mapData[2].value+1 }
                if(items[i].location_continents[j].includes("Europe")) { mapData[3].value = mapData[3].value+1 }
                if(items[i].location_continents[j].includes("Africa")) { mapData[4].value = mapData[4].value+1 }
                if(items[i].location_continents[j].includes("Australia")) { mapData[5].value = mapData[5].value+1 }
            }
        }

        showMap = false;

        mapConfigs = {
            type: 'world', // Map type
            width: '100%', // Width of the chart
            height: '450', // Height of the chart
            dataFormat: 'json', // Data Type
            renderAt:'chart-container',  //container where the chart will render
            dataSource: {
                // Map Configuration
                "chart": {
                    "caption": "World Map",
                    "labelsepchar": ": ",
                    "entityFillHoverColor": "#FFF9C4",
                    "theme": "fusion",
                    "showLegend": false,
                },
                // Source data as JSON --> id represents continents of the world.
                "data": mapData
            }
        };

        chartConfigs = {
            type: "column2d", //Select the chart type
            width: '100%', //Set the width of the chart
            height: 450, //Set the height of the chart
            dataFormat: "json", //Set the input dataFormat to json
            dataSource: {
                chart: {
                    caption: "Number of Outcomes Curated",
                    subCaption: "Split by Use Case",
                    xAxisName: "Use Cases", //Assign a relevant name to your x-axis
                    yAxisName: "Outcomes", //Assign a relevant name to your y-axis
                    numberSuffix: "",
                    theme: "fusion" //Apply a theme to your chart
                },
                //Include chartData from STEP 2
                data: use_case_data
            }
        };

        showMap = true

    }

    const removeTrial = async () => { items = [] }

    const reset = async () => { is_or = false; age_list = []; vaccines = []; schedules = []; manufacturers = []; continents = []; await getTrial(); }

    const checkbox_list = ["vaccine_checkboxes", "schedule_checkboxes", "manufacturer_checkboxes", "age_list_checkboxes", "continent_checkboxes"]

    function showCheckboxes(id) {
        for(let i = 0; i < checkbox_list.length; i++) {
          if(checkbox_list[i] !== id) {
              document.getElementById(checkbox_list[i]).style.display = "none";
          }
        }
        const checkboxes = document.getElementById(id);
        if (checkboxes.style.display === "none") { checkboxes.style.display = "block"; return }
        checkboxes.style.display = "none";
    }
</script>

<main class="is-fluid">
    <h1 class="has-text-centered title">Clinical Trials</h1>
    <div class="columns">
        <div class="column">
            <h3>Vaccine Comparison Dashboard</h3>
            <p>Please use the options on the right to search for the trials you'd like to find based on your search criteria.</p>
            <p>You can also sort the results by column by clicking on the header by ascending and then again for descending.</p>
        </div>
        <div class="column is-pulled-right" style="text-align: right;">
            <div class="vaccine-comparison">
                <b>Vaccine Comparison</b>
                <div class="multiselect" style="display: inline-block;">
                    <div class="selectBox" on:click={() => showCheckboxes('vaccine_checkboxes')}>
                        <select class="input">
                            <option>Select Vaccines</option>
                        </select>
                        <div class="overSelect"></div>
                    </div>
                    <div id="vaccine_checkboxes" class="checkboxes" style="display: none">
                        {#each default_vaccines as default_vaccine}
                            <label for="default_vaccine_list_{default_vaccine.Name.toLowerCase().replace(/\s/g, '')}">
                                <input type="checkbox" bind:group={vaccines} id="default_vaccine_list_{default_vaccine.Name.toLowerCase().replace(/\s/g, '')}" value="{default_vaccine.Name}" />{default_vaccine.Name}</label>
                        {/each}
                    </div>
                </div>
                <span style="display: inline-block;"><span style="vertical-align: sub">AND</span> <Switch bind:checked={is_or}></Switch> <span style="vertical-align: sub">OR</span></span>
            </div>
            <div class="multiselect" style="display: inline-block;">
                <div class="selectBox" on:click={() => showCheckboxes('schedule_checkboxes')}>
                    <select class="input">
                        <option>Select Schedules</option>
                    </select>
                    <div class="overSelect"></div>
                </div>
                <div id="schedule_checkboxes" class="checkboxes" style="display: none">
                    {#each default_schedules as default_schedule}
                        <label for="default_schedule_{default_schedule.Name.toLowerCase().replace(/\s/g, '')}">
                            <input type="checkbox" bind:group={schedules} id="default_schedule_{default_schedule.Name.toLowerCase().replace(/\s/g, '')}" value="{default_schedule.Name}" />{default_schedule.Name}</label>
                    {/each}
                </div>
            </div>
            <div class="multiselect" style="display: inline-block;">
                <div class="selectBox" on:click={() => showCheckboxes('manufacturer_checkboxes')}>
                    <select class="input">
                        <option>Select Manufacturers</option>
                    </select>
                    <div class="overSelect"></div>
                </div>
                <div id="manufacturer_checkboxes" class="checkboxes" style="display: none">
                    {#each default_manufacturers as default_manufacturer}
                        <label for="default_manufacturers_{default_manufacturer.toLowerCase().replace(/\s/g, '')}">
                            <input type="checkbox" bind:group={manufacturers} id="default_manufacturers_{default_manufacturer.toLowerCase().replace(/\s/g, '')}" value="{default_manufacturer}" />{default_manufacturer}</label>
                    {/each}
                </div>
            </div>
            <div class="multiselect" style="display: inline-block;">
                <div class="selectBox" on:click={() => showCheckboxes('age_list_checkboxes')}>
                    <select class="input">
                        <option>Select Ages</option>
                    </select>
                    <div class="overSelect"></div>
                </div>
                <div id="age_list_checkboxes" class="checkboxes" style="display: none">
                    {#each default_age_list as age}
                        <label for="default_age_list_{age.toLowerCase().replace(/\s/g, '')}">
                            <input type="checkbox" bind:group={age_list} id="default_age_list_{age.toLowerCase().replace(/\s/g, '')}" value="{age}" />{age}</label>
                    {/each}
                </div>
            </div>
            <div class="multiselect" style="display: inline-block;">
                <div class="selectBox" on:click={() => showCheckboxes('continent_checkboxes')}>
                    <select class="input">
                        <option>Select Continents</option>
                    </select>
                    <div class="overSelect"></div>
                </div>
                <div id="continent_checkboxes" class="checkboxes" style="display: none">
                    {#each default_continent_list as continent}
                        <label for="default_continent_{continent.toLowerCase().replace(/\s/g, '')}">
                            <input type="checkbox" bind:group={continents} id="default_continent_{continent.toLowerCase().replace(/\s/g, '')}" value="{continent}" />{continent}</label>
                    {/each}
                </div>
            </div>
            <div class="multiselect" style="display: inline-block;">
                <div class="selectBox" on:click={() => showCheckboxes('immunocompromised_groups_checkboxes')}>
                    <select class="input">
                        <option>Select Immunocompromised Groups</option>
                    </select>
                    <div class="overSelect"></div>
                </div>
                <div id="immunocompromised_groups_checkboxes" class="checkboxes" style="display: none">
                    {#each default_immunocompromised_groups as ig}
                        <label for="default_immunocompromised_groups_{ig.Name.toLowerCase().replace(/\s/g, '')}">
                            <input type="checkbox" bind:group={immunocompromised_groups} id="default_immunocompromised_groups_{ig.Name.toLowerCase().replace(/\s/g, '')}" value="{ig.Name}" />{ig.Name}</label>
                    {/each}
                    <hr />
                    <label for="default_immunocompromised_groups_healthy">
                        <input type="checkbox" bind:group={immunocompromised_groups} id="default_immunocompromised_groups_healthy" value="" />Healthy</label>
                </div>
            </div>
            <br /><br />
            <button class="button" style="display: inline-block;" on:click={() => reset()}>Reset</button>
            <button class="button" style="display: inline-block;" on:click={() => getTrial()}>Search</button>
        </div>
    </div>

    <table class="de-lightBorder" style="width: 100%; overflow-x: auto; display: block;">
        <thead>
        <tr>
            <th style="min-width: 90px; color: #fff !important" on:click={sort("nct_id")}>NCT ID <i class="fas fa-sort"></i></th>
            <th style="min-width: 90px; color: #fff !important" on:click={sort("study_name")}>Study Name <i class="fas fa-sort"></i></th>
            <th style="min-width: 90px; color: #fff !important" on:click={sort("phase")}>Phase <i class="fas fa-sort"></i></th>
            <th style="min-width: 90px; color: #fff !important" on:click={sort("vaccine")}>Vaccines <i class="fas fa-sort"></i></th>
            <th style="min-width: 116px; color: #fff !important" on:click={sort("schedule")}>Schedules <i class="fas fa-sort"></i></th>
            <th style="min-width: 150px; color: #fff !important" on:click={sort("manufacturer")}>Manufacturers <i class="fas fa-sort"></i></th>
            <th style="min-width: 93px; color: #fff !important" on:click={sort("age_list")}>Age List <i class="fas fa-sort"></i></th>
            <th style="min-width: 90px; color: #fff !important" on:click={sort("ethnicity")}>Ethnicity <i class="fas fa-sort"></i></th>
            <th style="min-width: 120px; color: #fff !important" on:click={sort("location_continents")}>Continents <i class="fas fa-sort"></i></th>
            <th style="min-width: 120px; color: #fff !important" on:click={sort("immunocompromised_groups")}>Immunocompromised Groups <i class="fas fa-sort"></i></th>
            <th style="min-width: 110px; color: #fff !important" on:click={sort("gender")}>Gender <i class="fas fa-sort"></i></th>
        </tr>
        </thead>
        <tbody transition:slide>
        {#each items as item}
        <tr>
            <td>{item.nct_id}</td>
            <td>{item.study_name}</td>
            <td>{item.phase}</td>
            <td><ul style="list-style: disc;">{#each item.vaccine as v}<li>{v}</li>{/each}</ul></td>
            <td><ul style="list-style: disc;">{#each item.schedule as s}<li>{s}</li>{/each}</ul></td>
            <td><ul style="list-style: disc;">{#each item.manufacturer as m}<li>{m}</li>{/each}</ul></td>
            <td>{item.age_list}</td>
            <td>{item.ethnicity}</td>
            <td><ul style="list-style: disc;">{#each item.location_continents as l}<li>{l}</li>{/each}</ul></td>
            <td><ul style="list-style: disc;">{#each item.immunocompromised_groups as ig}<li>{ig}</li>{/each}</ul></td>
            <td>{item.gender}</td>
        </tr>
        {/each}
        </tbody>
    </table>

    <div class="columns">
        <div class="column">
            {#if showMap}<SvelteFC {...mapConfigs} />{/if}
        </div>
        <div class="column">
            {#if showMap}<SvelteFC {...chartConfigs} />{/if}
        </div>
    </div>
</main>
