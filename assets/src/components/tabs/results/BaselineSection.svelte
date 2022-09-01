<script>
    import {onMount} from "svelte";

    export let data;

    let width = 0;

    onMount(async () => {
        let baseline_group = data.ResultsSection.BaselineCharacteristicsModule.BaselineGroupList.BaselineGroup;
        if (baseline_group == null) {
            return
        }
        width = 100/baseline_group.length+1;
    })
</script>
<div class="tr-indent1 tr-squishScroll" style="padding-top:1em; margin-top:1em">
    <div class="tr-indent2">
        <table class="de-lightBorder" style="width:98%; background-color:white; margin-top:1em">
            <tbody>
            <tr>
                <td class="de-baselineLabelCell" style="border-top-color:black; width:1px;" colspan="2">
                    Arm/Group Title
                </td>
                {#each data.ResultsSection.BaselineCharacteristicsModule.BaselineGroupList.BaselineGroup || [] as baselineGroup}
                    <th class="de-baselineLabelCell" style="border-top-color:black ;text-align:center; width:{width}%;">{baselineGroup.BaselineGroupTitle}</th>
                {/each}
            </tr>

            <tr class="EXPAND" id="EXPAND-armGroupDescriptionRow-baseline">
                <td class="de-baselineLabelCell" style="border-bottom-color:black;" colspan="2"><span>Arm/Group Description</span></td>
                {#each data.ResultsSection.BaselineCharacteristicsModule.BaselineGroupList.BaselineGroup || [] as baselineGroup}
                    <td class="de-baselineLabelCell" style="border-bottom-color:black; text-align:left">{baselineGroup.BaselineGroupDescription}</td>
                {/each}
            </tr>

            {#each data.ResultsSection.BaselineCharacteristicsModule.BaselineDenomList.BaselineDenom || [] as baselineDenom}
            <tr>
                <th class="de-baselineLabelCell" style="border-top-color:black;" colspan="2">Overall Number of Baseline {baselineDenom.BaselineDenomUnits}</th>
                {#each baselineDenom.BaselineDenomCountList.BaselineDenomCount || [] as baselineDenomCount}
                    <td class="de-numValue_baselineDataCell" style="border-top-color:black;">{baselineDenomCount.BaselineDenomCountValue}</td>
                {/each}
            </tr>
            {/each}

            {#each data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure || [] as baselineMeasure}

                <tr>
                    <th class="de-baselineLabelCell" style="border-top-color:black ; text-align:center;" colspan="{data.ResultsSection.BaselineCharacteristicsModule.BaselineGroupList.BaselineGroup !== null ? data.ResultsSection.BaselineCharacteristicsModule.BaselineGroupList.BaselineGroup.length+2:1}">
                        {baselineMeasure.BaselineMeasureTitle}
                        <div>
                            <div class="labelSubtle">{baselineMeasure.BaselineMeasureParamType} {#if baselineMeasure.BaselineMeasureDispersionType !== undefined}({baselineMeasure.BaselineMeasureDispersionType}){/if}</div>
                            <div class="labelSubtle">Unit of measure: {baselineMeasure.BaselineMeasureUnitOfMeasure}</div>
                        </div>
                    </th>
                </tr>

                {#each baselineMeasure.BaselineClassList.BaselineClass || [] as baselineClass}
                    {#each baselineClass.BaselineCategoryList.BaselineCategory || [] as baselineCategory}
                        <tr><td class="de-baselineLabelCell" style=""></td><td class="de-baselineLabelCell" style=""></td>
                        {#each baselineCategory.BaselineMeasurementList.BaselineMeasurement || [] as baselineMeasurement}
                            <td class="de-numValue_baselineDataCell" style="vertical-align:top;">
                            <span style="white-space:nowrap">{baselineMeasurement.BaselineMeasurementValue} {#if baselineMeasurement.BaselineMeasurementSpread !== ""} ({baselineMeasurement.BaselineMeasurementSpread}) {/if}</span>
                            </td>
                        {/each}
                        </tr>
                    {/each}
                {/each}

                {#if baselineMeasure.BaselineMeasureDescription !== undefined}
                <tr style="vertical-align:top;">
                <td class="de-baselineLabelCell" style="border-bottom-color:black;" colspan="2"></td>
                <td class="de-lightBorder" style="border-bottom-color:black" colspan="{data.ResultsSection.BaselineCharacteristicsModule.BaselineGroupList.BaselineGroup !== null ? data.ResultsSection.BaselineCharacteristicsModule.BaselineGroupList.BaselineGroup.length:1}">
                    <div style="white-space:nowrap;width:100%">
                        <div class="de-commentIndex tr-note_baseline_color" style="display:inline-block;width:2em;margin-right:1em; vertical-align:top; text-align:right">[1]</div>
                        <div style="display:inline-block;max-width:80%;min-width:20em;white-space:normal; text-align:left">Measure Description: {baselineMeasure.BaselineMeasureDescription}</div>
                    </div>
                </td>
                </tr>
                {/if}
            {/each}
            </tbody>
        </table>
    </div>
</div>
