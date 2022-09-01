<script>
    export let data;

    function getColSpan() {
        let event_group = data.ResultsSection.AdverseEventsModule.EventGroupList.EventGroup;
        if (event_group == null) {
            return 1
        }
        return event_group.length*2;
    }
</script>
<div class="tr-indent1 tr-squishScroll" style="margin-top:2em">
    <div class="tr-indent2" style="margin-top:1em;">

        <table class="de-lightBorder" style="width:98% ; background-color:white;">
            <tbody>
            <tr>
                <td class="de-labelCellAdEv" style="border-top-color:black; width:20.0%;"><a id="overview"></a><span style="width:25px; text-align:right">Time Frame</span></td>
                <td class="de-dataCellAdEv" style="border-top-color:black;" colspan="{getColSpan()}">{data.ResultsSection.AdverseEventsModule.EventsTimeFrame}</td>
            </tr>
            <tr>
                <td class="de-labelCellAdEv" style="border-bottom-color:black;">Adverse Event Reporting Description</td>
                <td class="de-dataCellAdEv" style="border-bottom-color:black;" colspan="{getColSpan()}">{data.ResultsSection.AdverseEventsModule.EventsDescription}</td>
            </tr>
            <tr><td class="de-heavyBorderTopBottom" colspan="{getColSpan()+1}" style="padding-top:1em">&nbsp;</td></tr>

            <tr>
                <td class="de-labelCellAdEv" style="border-top-color:black ; width:20.0%">
                    <span style="width:20px; text-align:right">Arm/Group Title</span>
                </td>
                {#each data.ResultsSection.AdverseEventsModule.EventGroupList.EventGroup || [] as evtGroup}
                <td class="de-labelCellAdEv" colspan="2" style="border-top-color: black; width:20.0% ; text-align:center;background-color:white ">{evtGroup.EventGroupTitle}</td>
                {/each}
            </tr>
            <tr class="EXPAND" id="EXPAND-armGroupDescriptionRow">
                <td class="de-labelCellAdEv" style="border-bottom-color: black"><span> Arm/Group Description</span></td>
                {#each data.ResultsSection.AdverseEventsModule.EventGroupList.EventGroup || [] as evtGroup}
                <td class="de-dataCellAdEv" style="border-bottom-color: black;" colspan="2">
                    {evtGroup.EventGroupDescription}
                </td>
                {/each}
            </tr>

            <tr>
                <td colspan="{getColSpan()+1}"><b>Serious Events</b></td>
            </tr>

            {#each data.ResultsSection.AdverseEventsModule.SeriousEventList.SeriousEvent || [] as seriousEvent}
                <tr>
                    <td colspan="{getColSpan()+1}">{seriousEvent.SeriousEventTerm}</td>
                </tr>
                <tr>
                    <td class="de-labelCellAdEv" style="border-top-color:black;"></td>
                    {#each data.ResultsSection.AdverseEventsModule.EventGroupList.EventGroup || [] as evtGroup}
                        <td class="de-labelCellAdEv" colspan="2" style="border-top-color: black; width:20.0% ; text-align:center;background-color:white ">{evtGroup.EventGroupTitle}</td>
                    {/each}
                </tr>
                <tr>
                    <td></td>
                    {#each seriousEvent.SeriousEventStatsList.SeriousEventStats || [] as stats}
                        <td>Affected / at Risk (%)</td>
                        <td>{stats.SeriousEventStatsNumAffected}/{stats.SeriousEventStatsNumAtRisk} ({Math.round((stats.SeriousEventStatsNumAffected/stats.SeriousEventStatsNumAtRisk)*100 * 100) / 100}%)</td>
                    {/each}
                </tr>
            {/each}

            <tr>
                <td colspan="{getColSpan()+1}"><b>Other Events</b></td>
            </tr>

            {#each data.ResultsSection.AdverseEventsModule.OtherEventList.OtherEvent || [] as otherEvent}
                <tr>
                    <td colspan="{getColSpan()+1}">{otherEvent.OtherEventTerm}</td>
                </tr>
                <tr>
                    <td class="de-labelCellAdEv" style="border-top-color:black;"></td>
                    {#each data.ResultsSection.AdverseEventsModule.EventGroupList.EventGroup || [] as evtGroup}
                        <td class="de-labelCellAdEv" colspan="2" style="border-top-color: black; width:20.0% ; text-align:center;background-color:white ">{evtGroup.EventGroupTitle}</td>
                    {/each}
                </tr>
                <tr>
                    <td></td>
                    {#each otherEvent.OtherEventStatsList.OtherEventStats || [] as stats}
                        <td>Affected / at Risk (%)</td>
                        <td>{stats.OtherEventStatsNumAffected}/{stats.OtherEventStatsNumAtRisk} ({Math.round((stats.OtherEventStatsNumAffected/stats.OtherEventStatsNumAtRisk)*100 * 100) / 100}%)</td>
                    {/each}
                </tr>
            {/each}
            </tbody>
        </table>
    </div>
</div>
