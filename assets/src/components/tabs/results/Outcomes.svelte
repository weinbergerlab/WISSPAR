<script>
    import OutcomeAnalysis from "../accordian/OutcomeAnalysis.svelte"
    import AccordianAnalysis from "../accordian/AccordianAnalysis.svelte";

    export let data;

    let currentOutcomeType = "";

    function isCurrentOutcomeType(outcomeType) {
        if (currentOutcomeType !== outcomeType) {
            currentOutcomeType = outcomeType
            return false
        }
        return true
    }

    function getGroupName(outcome, groupID) {
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

    function analysisToItems() {
        let items = [];
        outcome.OutcomeAnalysisList.OutcomeAnalysis.forEach(function (item, i) {
            items.push({
                label: "Statistical Analysis " + i,
                component: OutcomeAnalysis,
                data: item,
                outcome: outcome,
                is_open: false,
            })
        })
    }

</script>
{#each data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure || [] as outcome, i}
    {#if !isCurrentOutcomeType(outcome.OutcomeMeasureType)}
        <h3>{outcome.OutcomeMeasureType}</h3>
    {/if}
    <fieldset class="de-fieldset">
    <legend><span style="padding:0 1em">{i+1}.<span style="margin-left: 1em">{outcome.OutcomeMeasureType}</span></span></legend>
    <div style="padding:1em; overflow-x:auto;">
        <table class="de-lightBorder" style="width:100%;">
            <thead>
            <tr>
                <th style="min-width:150px; width:16%; visibility:hidden;"></th>
                <th style="visibility:hidden;"></th>
            </tr>
            </thead>
            <tbody>
            <tr>
                <th class="de-outcomeLabelCell">Title</th>
                <td class="de-outcomeDataCell">{outcome.OutcomeMeasureTitle}</td>
            </tr>
            <tr>
                <th class="de-outcomeLabelCell" style="white-space:nowrap">Description</th>
                <td class="de-outcomeDataCell">{outcome.OutcomeMeasureDescription}</td>
            </tr>
            <tr>
                <th class="de-outcomeLabelCell">Time Frame</th>
                <td class="de-outcomeDataCell">{outcome.OutcomeMeasureTimeFrame}</td>
            </tr>
            </tbody>
        </table>

        <div style="margin-top:1em;">
            <div style="width:100%">
                <div style="margin-bottom:1em;">
                    <span>Outcome Measure Data</span>
                </div>
                <div>
                    <div class="de-outcomeLabelCell" style="text-align:left ; width:100%"><span>Analysis Population Description</span></div>
                    <div class="de-outcomeDataCell" style="width:100%; border-top:none">{outcome.OutcomeMeasurePopulationDescription}</div>
                </div>

                    <table class="de-lightBorder" style="width:100%;">
                        <tbody>

                        <tr>
                            <td class="de-outcomeLabelCell" colspan="2" style="border-top:none;">
                                Arm/Group Title
                            </td>
                            {#each outcome.OutcomeGroupList.OutcomeGroup || [] as outcomeGroup}
                            <th class="de-outcomeLabelCell" style="border-top:none; text-align:center">{outcomeGroup.OutcomeGroupTitle}</th>
                            {/each}
                        </tr>
                        <tr class="EXPAND" id="EXPAND-outcome-armGroupDescriptionRow-outcome-1">
                            <td class="de-outcomeLabelCell" colspan="2">Arm/Group Description:</td>
                            {#each outcome.OutcomeGroupList.OutcomeGroup || [] as outcomeGroup}
                            <td class="de-outcomeLabelCell" style="text-align:left">
                                <div>{outcomeGroup.OutcomeGroupDescription}</div>
                            </td>
                            {/each}
                        </tr>

                        <tr>
                            <th class="de-outcomeLabelCell" colspan="2">Overall Number of Participants Analyzed</th>
                            {#each outcome.OutcomeDenomList.OutcomeDenom || [] as outcomeDenom}
                                {#each outcomeDenom.OutcomeDenomCountList.OutcomeDenomCount || [] as outcomeDenomCount}
                                    <td class="de-numValue_outcomeDataCell" style="vertical-align:top;">{outcomeDenomCount.OutcomeDenomCountValue}</td>
                                {/each}
                            {/each}
                        </tr>
                        <tr>
                            <th class="de-outcomeLabelCell" colspan="2">
                                {outcome.OutcomeMeasureParamType} {#if outcome.OutcomeMeasureDispersionType !== undefined}({outcome.OutcomeMeasureDispersionType}){/if}<br><div class="labelSubtle">Unit of Measure: {outcome.OutcomeMeasureUnitOfMeasure}</div>   </th>
                            <td class="de-outcomeDataCell"></td>
                            <td class="de-outcomeDataCell"></td>
                            <td class="de-outcomeDataCell"></td>
                        </tr>
                        {#each outcome.OutcomeClassList.OutcomeClass || [] as outcomeClass}
                            {#each outcomeClass.OutcomeClassDenomList.OutcomeClassDenom || [] as outcomeClassDenom, j}
                        <tr>
                            <th class="de-outcomeLabelCell" rowspan="2">{outcomeClass.OutcomeClassTitle}</th>
                            <th class="de-outcomeLabelCell">Number Analyzed</th>
                            {#each outcomeClassDenom.OutcomeClassDenomCountList.OutcomeClassDenomCount || [] as outcomeClassDenomCount}
                                <td class="de-numValue_outcomeDataCell">{outcomeClassDenomCount.OutcomeClassDenomCountValue} {outcomeClassDenom.OutcomeClassDenomUnits}                </td>
                            {/each}

                        </tr>
                        <tr>
                            <th class="de-outcomeLabelCell">
                            </th>

                            {#each outcomeClass.OutcomeCategoryList.OutcomeCategory[j].OutcomeMeasurementList.OutcomeMeasurement || [] as outcomeMeasurement}
                            <td class="de-numValue_outcomeDataCell">
                                <div>{outcomeMeasurement.OutcomeMeasurementValue}</div>
                                <div>({outcomeMeasurement.OutcomeMeasurementLowerLimit} to {outcomeMeasurement.OutcomeMeasurementUpperLimit})</div>
                            </td>
                            {/each}
                        </tr>
                            {/each}
                        {/each}
                        </tbody>
                    </table>
                <div>
                    <h3>Statistical Analysis</h3>
                    <AccordianAnalysis {outcome} />
            </div>
            </div>
        </div>
    </div>
</fieldset>
{/each}
