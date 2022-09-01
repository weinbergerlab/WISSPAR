<script>
    export let data = {};
</script>
<div id="tab-body">
    <!-- field list -->
    <div class="ct-indent1">

        <table class="ct-data_table tr-data_table tr-tableStyle">
            <tbody><tr>
                <th class="ct-header2 tr-banner_color tr-sectionBanner" colspan="2">Tracking Information</th>
            </tr>

            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">First Submitted Date&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.StatusModule.StudyFirstSubmitDate}</td>
            </tr>

            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">First Posted Date&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.StatusModule.StudyFirstPostDateStruct.StudyFirstPostDate}</td>
            </tr>

            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Results First Submitted Date&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.StatusModule.ResultsFirstSubmitDate}
                </td>
            </tr>

            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Results First Posted Date&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.StatusModule.ResultsFirstPostDateStruct.ResultsFirstPostDate}
                </td>
            </tr>

            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Last Update Posted Date </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.StatusModule.LastUpdatePostDateStruct.LastUpdatePostDate}
                </td>
            </tr>

            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Actual Study Start Date&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.StatusModule.StartDateStruct.StartDate}
                </td>
            </tr>

            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Actual Primary Completion Date </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.StatusModule.PrimaryCompletionDateStruct.PrimaryCompletionDate} &nbsp; (Final data collection date for primary outcome measure)
                </td>
            </tr>

            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Current Primary Outcome Measures&nbsp;<sup style="color:blue"> ICMJE </sup></th>
                <td class="ct-body3 tr-rowBody"><ul style="margin:0ex 1em; padding:0ex 0em">
                    {#each data.ProtocolSection.OutcomesModule.PrimaryOutcomeList.PrimaryOutcome || [] as primary_outcome}
                        <li style="margin-top:0.7ex;">{primary_outcome.PrimaryOutcomeMeasure} [&nbsp;Time&nbsp;Frame:&nbsp;{primary_outcome.PrimaryOutcomeTimeFrame}&nbsp;]<div class="tr-indent2" style="margin-top:1ex;">{primary_outcome.PrimaryOutcomeDescription}</div><br></li>
                    {/each}
                </ul>

                </td>
            </tr>

            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Current Secondary Outcome Measures&nbsp;<sup style="color:blue"> ICMJE </sup></th>
                <td class="ct-body3 tr-rowBody"><ul style="margin:0ex 1em; padding:0ex 0em">
                    {#each data.ProtocolSection.OutcomesModule.SecondaryOutcomeList.SecondaryOutcome || [] as secondary_outcome}
                        <li style="margin-top:0.7ex;">{secondary_outcome.SecondaryOutcomeMeasure} [&nbsp;Time&nbsp;Frame:&nbsp;{secondary_outcome.SecondaryOutcomeTimeFrame}&nbsp;]<div class="tr-indent2" style="margin-top:1ex;">{secondary_outcome.SecondaryOutcomeDescription}</div><br></li>
                    {/each}
                </ul>

                </td>
            </tr>

            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Current Other Pre-specified Outcome Measures</th>
                <td class="ct-body3 tr-rowBody"><ul style="margin:0ex 1em; padding:0ex 0em">
                    {#each data.ProtocolSection.OutcomesModule.OtherOutcomeList.OtherOutcome || [] as other_outcome}
                        <li style="margin-top:0.7ex;">{other_outcome.OtherOutcomeMeasure} [&nbsp;Time&nbsp;Frame:&nbsp;{other_outcome.OtherOutcomeTimeFrame}&nbsp;]<div class="tr-indent2" style="margin-top:1ex;">{other_outcome.OtherOutcomeDescription}</div><br></li>
                    {/each}
                </ul>

                </td>
            </tr>

            <tr><td colspan="2">&nbsp;</td></tr>

            <tr>
                <th class="ct-header2 tr-banner_color tr-sectionBanner" colspan="2">Descriptive Information</th>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Brief Title&nbsp;<sup style="color:blue"> ICMJE </sup>
                </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.IdentificationModule.BriefTitle}</td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Official Title&nbsp;<sup style="color:blue"> ICMJE </sup>
                </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.IdentificationModule.OfficialTitle}</td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Brief Summary
                </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.DescriptionModule.BriefSummary}</td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Detailed Description</th>
                <td class="ct-body3 tr-rowBody"><p style="margin-top:1ex; margin-bottom:1ex;">{data.ProtocolSection.DescriptionModule.DetailedDescription}</p></td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Study Type&nbsp;<sup style="color:blue"> ICMJE </sup></th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.DesignModule.StudyType}</td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Study Phase&nbsp;<sup style="color:blue"> ICMJE </sup>
                </th>
                <td class="ct-body3 tr-rowBody">{#each data.ProtocolSection.DesignModule.PhaseList.Phase || [] as phase}<span style="display:block;margin-bottom:1ex;">{phase}</span>{/each}</td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Study Design&nbsp;<sup style="color:blue"> ICMJE </sup>
                </th>
                <td class="ct-body3 tr-rowBody">Allocation: {data.ProtocolSection.DesignModule.DesignInfo.DesignAllocation}<br>Intervention Model: {data.ProtocolSection.DesignModule.DesignInfo.DesignInterventionModel}<br>Masking: {data.ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignMasking} ({#each data.ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignWhoMaskedList.DesignWhoMasked || [] as mask, i}{#if i !== 0}, {/if}{mask}{/each})<br>Primary Purpose: {data.ProtocolSection.DesignModule.DesignInfo.DesignPrimaryPurpose}
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Condition&nbsp;<sup style="color:blue"> ICMJE </sup>
                </th>
                <td class="ct-body3 tr-rowBody">{#each data.ProtocolSection.ConditionsModule.ConditionList.Condition || [] as condition}
                    <span style="display:block;margin-bottom:1ex;">{condition}</span>
                {/each}</td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Intervention&nbsp;<sup style="color:blue"> ICMJE </sup></th>
                <td class="ct-body3 tr-rowBody"><ul style="margin:0ex 1em; padding:0ex 0em">
                    {#each data.ProtocolSection.ArmsInterventionsModule.InterventionList.Intervention || [] as intervention}
                    <li style="margin:1ex 0.5ex">{intervention.InterventionType}: {intervention.InterventionName}
                        <div class="tr-indent2" style="margin-top:0.5ex">{intervention.InterventionDescription}</div></li>
                    {/each}
                </ul>
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Study Arms&nbsp;<sup style="color:blue"> ICMJE </sup>
                </th>
                <td class="ct-body3 tr-rowBody"><ul style="margin:0ex 1em; padding:0ex 0em">
                    {#each data.ProtocolSection.ArmsInterventionsModule.ArmGroupList.ArmGroup || [] as armgroup}
                    <li style="margin:1ex 0.5ex">{armgroup.ArmGroupType}: {armgroup.ArmGroupLabel}
                        <div class="tr-indent2" style="margin-top:0.5ex">{armgroup.ArmGroupDescription}</div>
                        <div class="tr-indent2" style="margin-top:0.5ex">Intervention:
                            {#each armgroup.ArmGroupInterventionList.ArmGroupInterventionName || [] as intervention}
                                <span style="display: block">{intervention}</span>
                            {/each}
                        </div>
                    </li>
                    {/each}
                </ul>
                </td>
            </tr>

            <tr><td colspan="2">&nbsp;</td></tr>

            <tr>
                <th class="ct-header2 tr-banner_color tr-sectionBanner" colspan="2">Recruitment Information</th>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Recruitment Status&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.StatusModule.OverallStatus}
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">{data.ProtocolSection.DesignModule.EnrollmentInfo.EnrollmentType} Enrollment&nbsp;<sup style="color:blue"> ICMJE </sup> <br></th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.DesignModule.EnrollmentInfo.EnrollmentCount}
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Actual Study Completion Date&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.StatusModule.CompletionDateStruct.CompletionDate}
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Actual Primary Completion Date </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.StatusModule.PrimaryCompletionDateStruct.PrimaryCompletionDate} &nbsp; (Final data collection date for primary outcome measure)
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Eligibility Criteria&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody">
                    {data.ProtocolSection.EligibilityModule.EligibilityCriteria}
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Sex/Gender&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody"><table class="tr-layout_table">
                    <tbody><tr>
                        <td style="vertical-align:top; padding:0px; margin:0ex; border:none; white-space:nowrap;">Sexes Eligible for Study:</td>
                        <td style="vertical-align:top; padding:0px 0px 0px 1em; margin:0ex; border:none;">{data.ProtocolSection.EligibilityModule.Gender}</td>
                    </tr>
                    </tbody></table>

                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Ages&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.EligibilityModule.MinimumAge} to {data.ProtocolSection.EligibilityModule.MaximumAge} &nbsp; ({#each data.ProtocolSection.EligibilityModule.StdAgeList.StdAge || [] as age, i}{#if i !== 0}, {/if}{age}{/each})
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Accepts Healthy Volunteers&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody">{#if data.ProtocolSection.EligibilityModule.HealthyVolunteers === "Accepts Healthy Volunteers"}Yes{:else}No{/if}
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Listed Location Countries&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody">{#each data.ProtocolSection.ContactsLocationsModule.LocationList.Location || [] as location}
                    <span style="display: block">{location.LocationCountry}</span>
                {/each}
                </td>
            </tr>

            <tr><td colspan="2">&nbsp;</td></tr>

            <tr>
                <th class="ct-header2 tr-banner_color tr-sectionBanner" colspan="2">Administrative Information</th>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">NCT Number&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.IdentificationModule.NCTId}
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Other Study ID Numbers&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.IdentificationModule.OrgStudyIdInfo.OrgStudyId}
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Responsible Party </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.IdentificationModule.Organization.OrgFullName}
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Study Sponsor&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.SponsorCollaboratorsModule.LeadSponsor.LeadSponsorName}
                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Investigators&nbsp;<sup style="color:blue"> ICMJE </sup> </th>
                <td class="ct-body3 tr-rowBody"><table class="tr-layout_table">
                    <tbody>
                    {#each data.ProtocolSection.ContactsLocationsModule.OverallOfficialList.OverallOfficial || [] as official}
                        <tr>
                            <td style="vertical-align:top; padding:0px; margin:0ex; border:none; white-space:nowrap;">{official.OverallOfficialRole}:</td>
                            <td style="vertical-align:top; padding:0px 0px 0px 1em; margin:0ex; border:none;">{official.OverallOfficialName}</td>
                            <td style="vertical-align:top; padding:0px 0px 0px 1em; margin:0ex; border:none;">{official.OverallOfficialAffiliation}</td>
                        </tr>
                    {/each}
                    </tbody></table>

                </td>
            </tr>
            <tr>
                <th class="ct-header3 tr-banner_color tr-rowHeader">Verification Date </th>
                <td class="ct-body3 tr-rowBody">{data.ProtocolSection.StatusModule.StatusVerifiedDate}
                </td>
            </tr>
            </tbody>
        </table>
    </div>
</div>
