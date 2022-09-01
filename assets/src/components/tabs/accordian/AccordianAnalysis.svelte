<script>
    import {onMount} from "svelte";
    import OutcomeAnalysis from "./OutcomeAnalysis.svelte";

    let items;
    export let outcome = {};

    let open = [];

    onMount(async () => {
        if (outcome !== {} && outcome !== null && outcome !== undefined ) {
            console.log(outcome)
            if (outcome.OutcomeAnalysisList.OutcomeAnalysis !== null) {
                outcome.OutcomeAnalysisList.OutcomeAnalysis.forEach(function () {
                    open.push(false)
                })
                items = outcome.OutcomeAnalysisList.OutcomeAnalysis
            }
        }
    })

    function toggle(i) {
        open[i] = !open[i];
    }
</script>

{#each items || [] as item, i}
    <button on:click={() => toggle(i)} aria-expanded={open[i]} class="accordianbutton2"><svg width="20" height="20" fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" stroke="currentColor"><path d="M9 5l7 7-7 7"></path></svg> Statistical Analysis {i+1} </button>
    {#if open[i]}<OutcomeAnalysis data={item} outcome={outcome} />{/if}
{:else}
    <p>No Statistical Analysis for this outcome measure</p>
{/each}
