<script>
    export let data;
</script>
<div class="tr-indent3" style="text-align:center;">
    <div class="tr-indent2">
        <table class="ct-data_table tr-data_table tr-tableStyle tr-studyInfo" style="width:98%;border-spacing:0;">
            <tbody>
            <tr>
                <th class="ct-header3 tr-pale_banner_color ct-nowrap">Study Type</th>
                <td class="ct-body3" style="text-align:left;">{data.ProtocolSection.DesignModule.StudyType}</td>
            </tr>
            <tr>
                <th class="ct-header3 tr-pale_banner_color ct-nowrap">Study Design</th>
                <td class="ct-body3" style="text-align:left;">Allocation:&nbsp;{data.ProtocolSection.DesignModule.DesignInfo.DesignAllocation}; &nbsp; Intervention&nbsp;Model:&nbsp;{data.ProtocolSection.DesignModule.DesignInfo.DesignInterventionModel} &nbsp; Masking:&nbsp;{data.ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignMasking} ({#each data.ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignWhoMaskedList.DesignWhoMasked || [] as mask, i}{#if i !== 0}, {/if}{mask}{/each}); &nbsp; Primary&nbsp;Purpose:&nbsp;{data.ProtocolSection.DesignModule.DesignInfo.DesignPrimaryPurpose}</td>
            </tr>
            <tr>
                <th class="ct-header3 tr-pale_banner_color ct-nowrap">Condition</th>
                <td class="ct-body3 ct-nowrap" style="text-align:left;">
                    {#each data.ProtocolSection.ConditionsModule.ConditionList.Condition || [] as condition}
                        {condition}<br />
                    {/each}
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-pale_banner_color ct-nowrap">Interventions</th>
                <td class="ct-body3 ct-nowrap" style="text-align:left;">
                    {#each data.ProtocolSection.ArmsInterventionsModule.InterventionList.Intervention || [] as intervention}
                        {intervention.InterventionType}: {intervention.InterventionName}<br />
                    {/each}
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-pale_banner_color ct-nowrap">Enrollment</th>
                <td class="ct-body3 ct-nowrap" style="text-align:left;">
                    {data.ProtocolSection.DesignModule.EnrollmentInfo.EnrollmentCount}
                </td>
            </tr>
            </tbody></table>
    </div>
</div>
