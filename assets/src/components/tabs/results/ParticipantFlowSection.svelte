<script>
    import {onMount} from "svelte";

    export let data;

    let width = 0;

    onMount(async () => {
        let flow_group = data.ResultsSection.ParticipantFlowModule.FlowGroupList.FlowGroup;
        if (flow_group == null) {
            return
        }
        width = 100/flow_group.length+1;
    })
</script>
<div class="tr-indent1 tr-squishScroll" style="padding-top:1em">
    <div class="tr-indent2">

        <table class="de-lightBorder" style="width:98% ; background-color: white;">
            <tbody>
            <tr>
                <td class="de-popFlowLabelCell" style="text-align:right">Recruitment Details</td>
                <td class="de-popFlowCell">{data.ResultsSection.ParticipantFlowModule.FlowRecruitmentDetails}</td>
            </tr>
            <tr>
                <td class="de-popFlowLabelCell" style="text-align:right; white-space:nowrap; width:15%">Pre-assignment Details</td>
                <td class="de-popFlowCell" style="width:83%">{data.ResultsSection.ParticipantFlowModule.FlowPreAssignmentDetails}</td>
            </tr>
            </tbody>
        </table>

        <table class="de-lightBorder" style="width:98% ; background-color: white; margin-top:1em">
            <thead>
            <tr style="height:0; border:0 none; background: #EBFFDD">
                {#each Array(data.ResultsSection.ParticipantFlowModule.FlowGroupList.FlowGroup) || [] as _, _}
                    <th style="width:{width}%"></th>
                {/each}
            </tr>
            </thead>
            <tbody>
            <tr>
                <td class="de-popFlowLabelCell" style="border-top:none; vertical-align:top; text-align:right">Arm/Group Title</td>
                {#each data.ResultsSection.ParticipantFlowModule.FlowGroupList.FlowGroup || [] as flowgroup}
                    <th class="de-popFlowLabelCell" colspan="1" style="border-top:none; vertical-align:top; text-align:center">{flowgroup.FlowGroupTitle}</th>
                {/each}
            </tr>
            <tr class="EXPAND" id="EXPAND-armGroupDescriptionRow-Result-ParticipantFlow">
                <td class="de-popFlowLabelCell" style="vertical-align:top ; text-align:right"><span style="padding-bottom:3px"> Arm/Group Description</span></td>
                {#each data.ResultsSection.ParticipantFlowModule.FlowGroupList.FlowGroup || [] as flowgroup}
                    <td class="de-popFlowLabelCell" colspan="1" style="text-align:left">{flowgroup.FlowGroupDescription}</td>
                {/each}
            </tr>

            {#each data.ResultsSection.ParticipantFlowModule.FlowPeriodList.FlowPeriod || [] as flowPeriod}
                <tr>
                    <td class="de-popFlowCell" colspan="5" style="padding-top:1em">
                        <span>Period Title: </span>
                        <span style="font-weight:bold">{flowPeriod.FlowPeriodTitle}</span>
                    </td>
                </tr>
                {#each flowPeriod.FlowMilestoneList.FlowMilestone || [] as flowMilestone}
                <tr>
                    <td class="de-popFlowLabelCell" style="padding-right:2em">{flowMilestone.FlowMilestoneType}</td>
                    {#each flowMilestone.FlowAchievementList.FlowAchievement || [] as flowAchievment}
                        <td class="de-numValue_popFlowCell" style="white-space:nowrap">{flowAchievment.FlowAchievementNumSubjects}</td>
                    {/each}
                </tr>
                {/each}

                {#each flowPeriod.FlowDropWithdrawList.FlowDropWithdraw || [] as flowDropWithdraw}
                <tr>
                    <td class="de-popFlowLabelCell">{flowDropWithdraw.FlowDropWithdrawType}</td>
                    {#each flowDropWithdraw.FlowReasonList.FlowReason || [] as flowReason}
                        <td class="de-numValue_popFlowCell" style="white-space:nowrap"><span style="padding-left:2em">{flowReason.FlowReasonNumSubjects}</span></td>
                    {/each}
                </tr>
                {/each}
            {/each}
            </tbody>
        </table>
    </div>
</div>
