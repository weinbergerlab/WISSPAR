<script>
    import TitleSection from "./results/TitleSection.svelte";
    import ParticipantFlowSection from "./results/ParticipantFlowSection.svelte";
    import BaselineSection from "./results/BaselineSection.svelte";
    import MoreInfoSection from "./results/MoreInfoSection.svelte";
    import AdverseEventsSection from "./results/AdverseEventsSection.svelte";
    import OutcomeMeasuresSection from "./results/OutcomeMeasuresSection.svelte";
    import Accordian from "./accordian/Accordian.svelte";

    export let data;

    let items = [
        {
            label: "Title Section",
            component: TitleSection,
            is_open: true,
        },
        {
            label: "Participant Flow",
            component: ParticipantFlowSection,
            is_open: false,
        },
        {
            label: "Baseline Characteristics",
            component: BaselineSection,
            is_open: false,
        },
        {
            label: "Outcome Measures",
            component: OutcomeMeasuresSection,
            is_open: false,
        },
        {
            label: "Adverse Events",
            component: AdverseEventsSection,
            is_open: false,
        },
        {
            label: "More Info",
            component: MoreInfoSection,
            is_open: false,
        }
    ];
</script>
<div id="tab-body">
    <div class="tr-indent2" style="background-color:white;">
        <Accordian {items} {data} />
    </div>
</div>
