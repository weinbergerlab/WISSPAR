<script>
    import {onMount} from "svelte";

    export let items = [];
    export let data = {};

    let open = [];

    onMount(async () => {
        items.forEach(function () {
            open.push(false)
        })
    })

    function toggle(i) {
        open[i] = !open[i];
    }
</script>
{#each items || [] as item, i}
    <button on:click={() => toggle(i)} aria-expanded={open[i]} class="accordianbutton"><svg style="tran"  width="20" height="20" fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" stroke="currentColor"><path d="M9 5l7 7-7 7"></path></svg> {item.label} </button>
    {#if open[i]}<svelte:component this={item.component} {data} />{/if}
{/each}
