<script>
    export let open = false;
    import { slide } from 'svelte/transition';
    const handleClick = () => open = !open
</script>

<div class="accordion">
    <div class="header">
        <div class="text">
            <slot name="head"></slot>
        </div>

        <button class="button" on:click={handleClick} >
            +/-
        </button>
    </div>
    {#if open}
        <div class="details" transition:slide>
            <slot name="details">
            </slot>
        </div>
    {/if}
</div>
