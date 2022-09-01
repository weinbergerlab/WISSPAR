<script>
    export let data;
    export let outcome;

    function getGroupName(groupID) {
        let outcomeGroup = outcome.OutcomeGroupList.OutcomeGroup;
        if (outcomeGroup == null) {
            return ""
        }
        for (let i = 0; i < outcomeGroup.length; i++) {
            if (outcomeGroup[i].OutcomeGroupId === groupID) {
                return outcomeGroup[i].OutcomeGroupTitle
            }
        }
        return ""
    }
</script>
<div style="margin:1em 0 0 0; vertical-align:top">
    <table class="de-lightBorder">
        <tbody>
        <tr>
            <th rowspan="4" class="de-outcomeLabelCell" style="width:1px">Statistical Analysis Overview</th>
            <th rowspan="1" class="de-outcomeLabelCell" style="width:1px; min-width:11em">Comparison Group Selection</th>
            <td class="de-outcomeDataCell" style="width:99%">
                {#each data.OutcomeAnalysisGroupIdList.OutcomeAnalysisGroupId || [] as groupID, k}
                    {#if k > 0}, {/if}{getGroupName(groupID)}
                {/each}
            </td>
        </tr>
        <tr>
            <th class="de-outcomeLabelCell">Comments</th>
            <td class="de-outcomeDataCell" style="padding-left:1em">{data.OutcomeAnalysisGroupDescription}</td>
        </tr>

        <tr>
            <th rowspan="1" class="de-outcomeLabelCell">
                <span style="white-space:nowrap">Type of</span>
                <span style="white-space:nowrap">Statistical Test</span>
            </th>
            <td class="de-outcomeDataCell">
                {data.OutcomeAnalysisNonInferiorityType}
            </td>
        </tr>
        <tr>
            <th class="de-outcomeLabelCell">Comments</th>
            <td class="de-outcomeDataCell" style="padding-left:1em">{data.OutcomeAnalysisNonInferiorityComment}</td>
        </tr>
        <tr>
            <th rowspan="4" colspan="1" class="de-outcomeLabelCell">
                Method of Estimation
            </th>
            <th class="de-outcomeLabelCell" style="white-space:nowrap">
                Estimation Parameter
            </th>
            <td class="de-outcomeDataCell">
                {data.OutcomeAnalysisParamType}
            </td>
        </tr>
        <tr>
            <th class="de-outcomeLabelCell">Estimated Value</th>
            <td class="de-outcomeDataCell">
                {data.OutcomeAnalysisParamValue}
            </td>
        </tr>
        <tr>
            <th class="de-outcomeLabelCell">
                Confidence Interval
            </th>

            <td class="de-outcomeDataCell">
                ({data.OutcomeAnalysisCINumSides}) {data.OutcomeAnalysisCIPctValue}%
                <div>{data.OutcomeAnalysisCILowerLimit} to {data.OutcomeAnalysisCIUpperLimit}</div>
            </td>
        </tr>
        </tbody>
    </table>
</div>
