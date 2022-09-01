<script>
    export let data;

    function getInterventionList() {
        let interventionList = [];

        const armGroupList = data.ProtocolSection.ArmsInterventionsModule.ArmGroupList.ArmGroup;

        if (armGroupList == null) {
            return interventionList;
        }

        for (let i = 0; i < armGroupList.length; i++) {
            const interventionListArr = armGroupList[i].ArmGroupInterventionList.ArmGroupInterventionName;
            if (interventionListArr == null) {
                continue
            }
            for (let j = 0; j < interventionListArr.length; j++) {
                let insert = true
                for (let k = 0; k < interventionList.length; k++) {
                    if (interventionListArr[j] === interventionList[k]) {
                        insert = false
                        break
                    }
                }

                if (insert) {
                    interventionList.push(interventionListArr[j])
                }
            }
        }

        return interventionList
    }
</script>

<div id="tab-body">
    <div class="tr-indent2" style="background-color:white;">

        <!-- purpose_section -->
        <div class="tr-indent1" style="margin-top:1ex">

            <div class="tr-subsection">
                <div class="ct-header2"><h3>Study Description</h3></div>
                <div class="tr-underline"></div>
            </div>

            <div class="tr-indent2" style="margin-top:2ex">

                <div class="ct-body3">Brief Summary:</div>

                <div class="ct-body3 tr-indent2">{data.ProtocolSection.DescriptionModule.BriefSummary}</div><br>

                <!-- condition, intervention, phase summary table -->
                <table class="ct-data_table tr-data_table" border="1" style="margin:auto;width:80%;">
                    <tbody><tr style="text-align:left;">
                        <th class="ct-header3 tr-pale_banner_color">
                            <span style="display:inline;" class="term" data-term="condition/disease" title="Click to define" tabindex="0">Condition or disease <i class="fa fa-info-circle term" aria-hidden="true" data-term="condition/disease" style="border-bottom-style:none;" title="Click to define" tabindex="0"></i></span>       </th>
                        <th class="ct-header3 tr-pale_banner_color">
                            <span style="display:inline;" class="term" data-term="intervention/treatment" title="Click to define" tabindex="0">Intervention/treatment <i class="fa fa-info-circle term" aria-hidden="true" data-term="intervention/treatment" style="border-bottom-style:none;" title="Click to define" tabindex="0"></i></span>       </th>
                        <th class="ct-header3 tr-pale_banner_color">
                            <span style="display:inline;" class="term" data-term="phase" title="Click to define" tabindex="0">Phase <i class="fa fa-info-circle term" aria-hidden="true" data-term="phase" style="border-bottom-style:none;" title="Click to define" tabindex="0"></i></span>       </th>
                    </tr>
                    <tr style="text-align:left;vertical-align:top;">
                        <td class="ct-body3">
                            {#each data.ProtocolSection.ConditionsModule.ConditionList.Condition || [] as condition}
                                <span style="display:block;margin-bottom:1ex;">{condition}</span>
                            {/each}
                        </td>
                        <td class="ct-body3">
                            {#each getInterventionList() || [] as intervention}
                                <span style="display:block;margin-bottom:1ex;">{intervention}</span>
                            {/each}
                        </td>
                        <td class="ct-body3" style="white-space:nowrap;">
                            {#each data.ProtocolSection.DesignModule.PhaseList.Phase || [] as phase}
                                <span style="display:block;margin-bottom:1ex;">{phase}</span>
                            {/each}
                        </td>
                    </tr>
                    </tbody></table>
                <br>
                <!-- detailed description -->
                <div id="COLLAPSE-DetailedDesc" class="COLLAPSE" style="margin-bottom:2ex;">
                    <div class="ct-body3">Detailed Description:</div>
                    <div class="ct-body3 tr-indent2">{data.ProtocolSection.DescriptionModule.DetailedDescription}</div>
                </div>
            </div>
            <div class="tr-subsection">
                <div class="ct-header2"><h3>Study Design</h3></div>
                <div class="tr-underline"></div>
            </div>

            <p>
            </p><table class="ct-layout_table tr-tableStyle tr-studyInfo" style="margin-bottom:2ex;">
            <caption>Layout table for study information</caption>
            <thead>
            <tr>
                <th id="studyInfoColTitle"></th>
                <th id="studyInfoColData"></th>
            </tr>
            </thead>
            <tbody>
            <tr>
                <td headers="studyInfoColTitle">
                    <span style="display:inline;" class="term" data-term="study type" title="Click to define" tabindex="0">Study Type <i class="fa fa-info-circle term" aria-hidden="true" data-term="study type" style="border-bottom-style:none;" title="Click to define" tabindex="0"></i></span> :</td>
                <td headers="studyInfoColData" style="padding-left:1em">
                    {data.ProtocolSection.DesignModule.StudyType}
                    &nbsp;(Clinical Trial)
                </td>
            </tr>
            <tr>
                <td headers="studyInfoColTitle">{data.ProtocolSection.DesignModule.EnrollmentInfo.EnrollmentType}
                    <span style="display:inline;" class="term" data-term="enrollment" title="Click to define" tabindex="0">Enrollment <i class="fa fa-info-circle term" aria-hidden="true" data-term="enrollment" style="border-bottom-style:none;" title="Click to define" tabindex="0"></i></span> :</td>
                <td headers="studyInfoColData" style="padding-left:1em">{data.ProtocolSection.DesignModule.EnrollmentInfo.EnrollmentCount} participants</td>
            </tr>
            <tr>
                <td headers="studyInfoColTitle">Allocation:</td>
                <td headers="studyInfoColData" style="padding-left:1em">{data.ProtocolSection.DesignModule.DesignInfo.DesignAllocation}</td>
            </tr>
            <tr>
                <td headers="studyInfoColTitle">Intervention Model:</td>
                <td headers="studyInfoColData" style="padding-left:1em">{data.ProtocolSection.DesignModule.DesignInfo.DesignInterventionModel}</td>
            </tr>
            <tr>
                <td headers="studyInfoColTitle">Masking:</td>
                <td headers="studyInfoColData" style="padding-left:1em"> {data.ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignMasking} ({#each data.ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignWhoMaskedList.DesignWhoMasked || [] as mask, i}{#if i !== 0}, {/if}{mask}{/each})</td>
            </tr>
            <tr>
                <td headers="studyInfoColTitle">Primary Purpose:</td>
                <td headers="studyInfoColData" style="padding-left:1em"> {data.ProtocolSection.DesignModule.DesignInfo.DesignPrimaryPurpose}</td>
            </tr>
            <tr>
                <!--
                      <td headers="studyInfoColTitle" style="padding-top:2ex;">Official Title:</td>
                -->
                <td headers="studyInfoColTitle">Official Title:</td>
                <td headers="studyInfoColData" style="padding-left:1em;">{data.ProtocolSection.IdentificationModule.OfficialTitle}</td>
            </tr>

            <tr>
                <td headers="studyInfoColTitle">  {data.ProtocolSection.StatusModule.StartDateStruct.StartDateType} <span style="display:inline;" class="term" data-term="study start date" title="Click to define" tabindex="0">Study Start Date <i class="fa fa-info-circle term" aria-hidden="true" data-term="study start date" style="border-bottom-style:none;" title="Click to define" tabindex="0"></i></span> :</td>
                <td headers="studyInfoColData" style="padding-left:1em">{data.ProtocolSection.StatusModule.StartDateStruct.StartDate}</td>
            </tr>
            <tr>
                <td headers="studyInfoColTitle">  {data.ProtocolSection.StatusModule.PrimaryCompletionDateStruct.PrimaryCompletionDateType} <span style="display:inline;" class="term" data-term="primary completion date" title="Click to define" tabindex="0">Primary Completion Date <i class="fa fa-info-circle term" aria-hidden="true" data-term="primary completion date" style="border-bottom-style:none;" title="Click to define" tabindex="0"></i></span> :</td>
                <td headers="studyInfoColData" style="padding-left:1em">{data.ProtocolSection.StatusModule.PrimaryCompletionDateStruct.PrimaryCompletionDate}</td>
            </tr>
            <tr>
                <td headers="studyInfoColTitle">  {data.ProtocolSection.StatusModule.CompletionDateStruct.CompletionDateType} <span style="display:inline;" class="term" data-term="study completion date" title="Click to define" tabindex="0">Study Completion Date <i class="fa fa-info-circle term" aria-hidden="true" data-term="study completion date" style="border-bottom-style:none;" title="Click to define" tabindex="0"></i></span> :</td>
                <td headers="studyInfoColData" style="padding-left:1em">{data.ProtocolSection.StatusModule.CompletionDateStruct.CompletionDate}</td>
            </tr>
            </tbody>
        </table>

        </div>

        <!-- arms and groups table -->
        <div class="tr-subsection">
            <div class="ct-header2"><h3>Arms and Interventions</h3></div>
            <div class="tr-underline"></div>

        </div>

        <div style="padding-left:.8em; padding-right:1.2em;text-align:left;">
            <p>
            </p><table class="ct-data_table tr-data_table" style="padding:5px;width:100%;border-width:1px;">
            <thead>
            <tr style="text-align:left;">
                <th class="ct-header3 tr-pale_banner_color" style="width:50%;">
                    <span style="display:inline;" class="term" data-term="arm" title="Click to define" tabindex="0">Arm <i class="fa fa-info-circle term" aria-hidden="true" data-term="arm" style="border-bottom-style:none;" title="Click to define" tabindex="0"></i></span> 			        </th>
                <th class="ct-header3 tr-pale_banner_color">
                    <span style="display:inline;" class="term" data-term="intervention/treatment" title="Click to define" tabindex="0">Intervention/treatment <i class="fa fa-info-circle term" aria-hidden="true" data-term="intervention/treatment" style="border-bottom-style:none;" title="Click to define" tabindex="0"></i></span>         </th>
            </tr>
            </thead>
            <tbody>
            {#each data.ProtocolSection.ArmsInterventionsModule.ArmGroupList.ArmGroup || [] as armGroup}
            <tr style="text-align:left;vertical-align:top;">
                <td class="ct-body3">
                    {armGroup.ArmGroupType}: {armGroup.ArmGroupLabel}
                    <div class="tr-indent2" style="margin-top:0.5ex">{armGroup.ArmGroupDescription}</div>
                </td>
                <td class="ct-body3">
                    {#each armGroup.ArmGroupInterventionList.ArmGroupInterventionName || [] as name}
                        {name}
                        {#each data.ProtocolSection.ArmsInterventionsModule.InterventionList || [] as list}
                            {#if name === list.InterventionType + ": " + list.InterventionName}
                                <div class="tr-indent2" style="margin-top:0.5ex">{list.InterventionDescription}</div><br>
                            {/if}
                        {/each}
                    {/each}
                </td>
            </tr>
            {/each}
            </tbody>
        </table>
        </div>
        <br>


        <!-- more details -->
        <br>
        <div class="tr-subsection">
            <div class="ct-header2"><h3>Outcome Measures</h3></div>
            <div class="tr-underline"></div>

        </div>
        <br>
        <div class="tr-indent3" style="margin-right:1em;">

            <!-- primary outcomes -->
            <div class="ct-body3">
                <span style="display:inline;" class="term" data-term="primary outcome measure" title="Click to define" tabindex="0">Primary Outcome Measures <i class="fa fa-info-circle term" aria-hidden="true" data-term="primary outcome measure" style="border-bottom-style:none;" title="Click to define" tabindex="0"></i></span> :
                <ol style="margin-top:1ex; margin-bottom:1ex;">
                    {#each data.ProtocolSection.OutcomesModule.PrimaryOutcomeList.PrimaryOutcome || [] as primary_outcome}
                        <li style="margin-top:0.7ex;">{primary_outcome.PrimaryOutcomeMeasure} [&nbsp;Time&nbsp;Frame:&nbsp;{primary_outcome.PrimaryOutcomeTimeFrame}&nbsp;]<div class="tr-indent2" style="margin-top:1ex;">{primary_outcome.PrimaryOutcomeDescription}</div><br></li>
                    {/each}
                </ol>
            </div>
            <br>

            <!-- secondary outcomes -->
            <div class="ct-body3" style="margin-right:1em;">
                <span style="display:inline;" class="term" data-term="secondary outcome measure" title="Click to define" tabindex="0">Secondary Outcome Measures <i class="fa fa-info-circle term" aria-hidden="true" data-term="secondary outcome measure" style="border-bottom-style:none;" title="Click to define" tabindex="0"></i></span> :
                <ol style="margin-top:1ex; margin-bottom:1ex;">
                    {#each data.ProtocolSection.OutcomesModule.SecondaryOutcomeList.SecondaryOutcome || [] as secondary_outcome}
                        <li style="margin-top:0.7ex;">{secondary_outcome.SecondaryOutcomeMeasure} [&nbsp;Time&nbsp;Frame:&nbsp;{secondary_outcome.SecondaryOutcomeTimeFrame}&nbsp;]<div class="tr-indent2" style="margin-top:1ex;">{secondary_outcome.SecondaryOutcomeDescription}</div><br></li>
                    {/each}
                </ol>
            </div>
            <br>

            <!-- other outcomes -->
            <div class="ct-body3" style="margin-right:1em;"><b>Other Outcome Measures:</b>
                <ol style="margin-top:1ex; margin-bottom:1ex;">
                    {#each data.ProtocolSection.OutcomesModule.OtherOutcomeList.OtherOutcome || [] as other_outcome}
                        <li style="margin-top:0.7ex;">{other_outcome.OtherOutcomeMeasure} [&nbsp;Time&nbsp;Frame:&nbsp;{other_outcome.OtherOutcomeTimeFrame}&nbsp;]<div class="tr-indent2" style="margin-top:1ex;">{other_outcome.OtherOutcomeDescription}</div><br></li>
                    {/each}
                </ol>
            </div>
            <br>


        </div>


        <!-- eligibility_section -->
        <div class="tr-indent1" style="margin-top:3ex; border:1px solid white">
            <div class="tr-subsection">
                <div class="ct-header2"><h3>Eligibility Criteria</h3></div>
                <div class="tr-underline"></div>
            </div>
            <br></div>


            <br>
            <div class="tr-indent2">
                <table class="ct-layout_table tr-tableStyle tr-studyInfo">
                    <caption>Layout table for eligibility information</caption>
                    <thead>
                    <tr>
                        <th id="elgType"></th>
                        <th id="elgData"></th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <td headers="elgType" class="ct-nowrap">Ages Eligible for Study: &nbsp; </td>
                        <td headers="elgData" style="padding-left:1em">{data.ProtocolSection.EligibilityModule.MinimumAge} to {data.ProtocolSection.EligibilityModule.MaximumAge} &nbsp; ({#each data.ProtocolSection.EligibilityModule.StdAgeList.StdAge || [] as age, i}{#if i !== 0}, {/if}{age}{/each})</td>
                    </tr>
                    <tr>
                        <td headers="elgType" class="ct-nowrap">Sexes Eligible for Study: &nbsp; </td>
                        <td headers="elgData" style="padding-left:1em">{data.ProtocolSection.EligibilityModule.Gender}</td>
                    </tr>
                    <tr>
                        <td headers="elgType" class="ct-nowrap">Accepts Healthy Volunteers: &nbsp; </td>
                        <td headers="elgData" style="padding-left:1em">{#if data.ProtocolSection.EligibilityModule.HealthyVolunteers === "Accepts Healthy Volunteers"}Yes{:else}No{/if}</td>
                    </tr>
                    </tbody>
                </table>


                <div class="ct-header3" style="margin-top:2ex">Criteria</div>
                <div class="tr-indent2"><p style="margin-top:1ex; margin-bottom:1ex;">{data.ProtocolSection.EligibilityModule.EligibilityCriteria}</p></div>

            </div>

        </div>


        <!-- location_section -->
        <a id="contacts"></a>
        <div class="tr-indent1" style="margin-top:3ex">
            <div class="tr-subsection">
                <div class="ct-header2"><h3>Contacts and Locations</h3></div>
                <div class="tr-underline"></div>

            </div>
            <br>


            <div class="tr-indent2" style="margin-top:2ex">
                <!-- contacts -->

                <!-- locations -->
                <div class="ct-header3" style="margin-top:2ex;margin-bottom:2ex;">Locations</div>


                <div id="COLLAPSE-Locations" class="COLLAPSE">
                    <div class="tr-table_cover">
                        <table class="ct-layout_table tr-indent2">
                            <caption>Layout table for location information</caption>
                            <thead>
                            <tr>
                                <th id="locName">Facility</th>
                                <th id="locCity">City</th>
                                <th id="locCountry">Country</th>
                            </tr>
                            </thead>
                            <tbody>
                            {#each data.ProtocolSection.ContactsLocationsModule.LocationList.Location || [] as location}
                                <tr>
                                    <td headers="locName" style="padding-left:2em;">{location.LocationFacility}</td>
                                    <td headers="locCity">{location.LocationCity}</td>
                                    <td headers="locCountry">{location.LocationCountry}</td>
                                </tr>
                            {/each}
                            </tbody>
                        </table>
                    </div>
                </div>

                <!-- sponsors -->
                <div class="ct-header3" style="margin-top:2ex">Sponsors and Collaborators</div>
                <div class="tr-indent2" style="margin-top:1ex">PATH</div>

                <!-- overall officials -->
                <div class="ct-header3" style="margin-top:2ex">Investigators</div>
                <div class="tr-table_cover">
                    <table class="ct-layout_table tr-indent2">
                        <caption>Layout table for investigator information</caption>
                        <tbody>
                        {#each data.ProtocolSection.ContactsLocationsModule.OverallOfficialList.OverallOfficial || [] as official}
                        <tr>
                            <td headers="role">{official.OverallOfficialRole}:</td>
                            <td headers="name">{official.OverallOfficialName}</td>
                            <td headers="affiliation">{official.OverallOfficialAffiliation}</td>
                        </tr>
                        {/each}
                        </tbody></table>
                </div>


            </div>
        </div>

        <!-- More info section -->
        <div class="tr-indent1" style="margin-top:3ex; border:1px solid white">
            <div class="tr-subsection">
                <div class="ct-header2"><h3>More Information</h3></div>
                <div class="tr-underline"></div>
            </div>

            <div class="tr-indent2" style="margin-top:1em;">

                <table class="ct-layout_table tr-tableStyle tr-moreInfo" style="margin-top:1ex;">
                    <caption>Layout table for additonal information</caption>
                    <thead>
                    <tr>
                        <th id="colTitle1"></th>
                        <th id="colData1"></th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <td headers="colTitle1" class="ct-nowrap">Responsible Party:</td>
                        <td headers="colData1" style="padding-left:1em">{data.ProtocolSection.SponsorCollaboratorsModule.LeadSponsor.LeadSponsorName}</td>
                    </tr>
                    <tr>
                        <td headers="colTitle1" class="ct-nowrap">Other Study ID Numbers:</td>
                        <td headers="colData1" style="padding-left:1em">
                            {data.ProtocolSection.IdentificationModule.OrgStudyIdInfo.OrgStudyId} <br>
                        </td>
                    </tr>
                    <tr>
                        <td headers="colTitle1" class="ct-nowrap">First Posted:</td>
                        <td headers="colData1" style="padding-left:1em">{data.ProtocolSection.StatusModule.StudyFirstPostDateStruct.StudyFirstPostDate}</td>
                    </tr>
                    <tr>
                        <td headers="colTitle1" class="ct-nowrap">Results First Posted:</td>
                        <td headers="colData1" style="padding-left:1em">{data.ProtocolSection.StatusModule.ResultsFirstPostDateStruct.ResultsFirstPostDate}</td>
                    </tr>
                    <tr>
                        <td headers="colTitle1" class="ct-nowrap">Last Update Posted:</td>
                        <td headers="colData1" style="padding-left:1em">{data.ProtocolSection.StatusModule.LastUpdatePostDateStruct.LastUpdatePostDate}</td>
                    </tr>
                    <tr>
                        <td headers="colTitle1" class="ct-nowrap">Last Verified:</td>
                        <td headers="colData1" style="padding-left:1em">{data.ProtocolSection.StatusModule.StatusVerifiedDate}</td>
                    </tr>

                    <!-- Plan to Share Data -->
                    </tbody>
                </table>


                <!-- keywords -->

                <!-- mesh terms -->
                <div class="tr-tableLabel">Additional relevant MeSH terms:</div>
                <div class="tr-indent3">
                    <table class="ct-layout_table tr-meshTerms" style="width:100%;">
                        <caption>Layout table for MeSH terms</caption>
                        <thead>
                        <tr>
                            <th id="meshTermCol1" style="width:50%;"></th>
                        </tr>
                        </thead>
                        <tbody>
                        <tr>
                            <td headers="meshTermCol1" style="padding-left:1em;">
                            {#each data.DerivedSection.ConditionBrowseModule.ConditionAncestorList.ConditionAncestor || [] as list, i}
                                {list.ConditionAncestorTerm}<br />
                            {/each}
                            </td>
                        </tr>
                        </tbody></table>
                </div>
            </div>
        </div>
    </div>
