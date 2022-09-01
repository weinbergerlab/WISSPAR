<script>
    import Outcomes from "./Outcomes.svelte";

    export let data;
</script>
<div class="tr-indent1 tr-squishScroll" style="margin-top:2em">
    <div class="tr-indent2" style="margin-top:1em;">
        <Outcomes {data} />
    </div>
</div>
