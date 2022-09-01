<script>
    import Tabs from "./Tabs.svelte";
    import StudyDetails from "./tabs/StudyDetails.svelte";
    import Results from "./tabs/Results.svelte";
    import TabularView from "./tabs/TabularView.svelte";
    import {onMount} from "svelte";
    import { Confirm } from 'svelte-confirm'

    let showStudy = false
    let data = {};
    const trialsAPI = '/api/trials/';
    let id = "NCT03197376";

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

    let reload = false;
    let stage = 1;
    onMount(async () => {})

    const getTrial = async () => {
        stage = 1.5
        const studyURL = trialsAPI + id

        const studyAPI = await fetch(studyURL, {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
        })


        if (studyAPI.status !== 200) {
            stage = 1;
            return;
        }

        data = await studyAPI.json()
        showStudy = true
        stage = 2
    }


    const removeTrial = async () => {
        stage = 1;
        data = {};
        showStudy = false;
    }

</script>

<main class="is-fluid">
    <h1 class="has-text-centered title">Look Up A Trial</h1>

    {#if stage === 1}
    <div id="stage-1">
        <div class="block has-text-centered">
            <p>Please enter the NCT ID of the clinical trial you wish to view (eg. <i>NCT03197376</i>)</p>
        </div>
        <form class="field has-addons" style="justify-content: center" on:submit|preventDefault={getTrial}>
            <div class="field">
                <div class="control">
                    <label for="clinical-trial-id"><b>Clinical Trial ID</b></label>
                    <input bind:value={id} id="clinical-trial-id" name="clinical-trial-id" class="input" type="text" placeholder="Clinical Trial ID">
                </div>
            </div>

            <div class="control">
                <br />
                <button class="button is-primary" aria-label="Search">
                    <span class="icon is-small" role="none">
                      <i class="fas fa-eye" role="none"></i>
                    </span>
                </button>
            </div>
        </form>
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
                <div class="column">
                    <h2 class="subtitle has-text-centered" style="font-size: 27px; font-weight: bold;">Trial Data</h2>
                    <button on:click={removeTrial} class="button is-pulled-right">Back</button>
                </div>
            </div>
            <div class="columns is-centered">
                <div id="trial-data" class="column">
                    <div id="study">
                        {#if showStudy}<Tabs {items} {data} />{/if}
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
</main>
