<script>
    // Data for insert
    let data = {
        "ProtocolSection":  {
            "IdentificationModule":  {
                "NCTID": "",
                "OrgStudyIDInfo":  {
                    "OrgStudyID": ""
                },
                "Organization":  {
                    "OrgFullName": "",
                    "OrgClass":    ""
                },
                "BriefTitle": "",
                "OfficialTitle": ""
            },
            "StatusModule":  {
                "StatusVerifiedDate": "",
                "OverallStatus": "",
                "StartDate":  {
                    "StartDate": ""
                },
                "PrimaryCompletionDate":  {
                    "PrimaryCompletionDate": ""
                },
                "CompletionDate":  {
                    "CompletionDate": ""
                },
                "StudyFirstSubmitDate": "",
                "StudyFirstPostDate":  {
                    "StudyFirstPostDate": ""
                },
                "ResultsFirstSubmitDate": "",
                "LastUpdateSubmitDate": ""
            },
            "SponsorCollaboratorsModule":  {
                "ResponsibleParty":  {
                    "ResponsiblePartyType": ""
                },
                "LeadSponsor":  {
                    "LeadSponsorName": "",
                    "LeadSponsorClass": ""
                }
            },
            "OversightModule":  {
                "OversightHasDMC": "",
                "IsFDARegulatedDrug": ""
            },
            "DescriptionModule":  {
                "BriefSummary": "",
                "DetailedDescription": ""
            },
            "ConditionsModule":  {
                "ConditionList":  {
                    "Condition": [""]
                }
            },
            "DesignModule":  {
                "StudyType": "",
                "PhaseList":  {
                    "Phase": [""]
                },
                "DesignInfo":  {
                    "DesignAllocation":        "",
                    "DesignInterventionModel": "",
                    "DesignPrimaryPurpose":    "",
                    "DesignMaskingInfo":        {
                        "DesignMasking": "",
                        "DesignWhoMaskedList":  {
                            "DesignWhoMasked": [""]
                        }
                    }
                },
                "EnrollmentInfo":  {
                    "EnrollmentCount": "",
                    "EnrollmentType": ""
                }
            },
            "ArmsInterventionsModule":  {
                "ArmGroupList":  {
                    "ArmGroup": [
                        // {
                        //     "ArmGroupLabel": "",
                        //     "ArmGroupType": "",
                        //     "ArmGroupDescription": "",
                        //     "ArmGroupInterventionList":  {
                        //         "ArmGroupInterventionName": [""]
                        //     }
                        // }
                    ]
                },
                "InterventionList":  {
                    "Intervention": [
                        // {
                        //     "InterventionType": "",
                        //     "InterventionName": "",
                        //     "InterventionDescription": "",
                        //     "InterventionArmGroupLabelList":  {
                        //         "InterventionArmGroupLabel": [""]
                        //     }
                        // }
                    ]
                }
            },
            "OutcomesModule":  {
                "PrimaryOutcomeList":  {
                    "PrimaryOutcome": [
                        // {
                        //     "PrimaryOutcomeMeasure": "",
                        //     "PrimaryOutcomeDescription": "",
                        //     "PrimaryOutcomeTimeFrame": ""
                        // }
                    ]
                },
                "SecondaryOutcomeList":  {
                    "SecondaryOutcome": [
                        // {
                        //     "SecondaryOutcomeMeasure": "",
                        //     "SecondaryOutcomeDescription": "",
                        //     "SecondaryOutcomeTimeFrame": ""
                        // }
                    ]
                },
                "OtherOutcomeList":  {
                    "OtherOutcome": [
                        // {
                        //     "OtherOutcomeMeasure": "",
                        //     "OtherOutcomeDescription": "",
                        //     "OtherOutcomeTimeFrame": ""
                        // }
                    ]
                }
            },
            "EligibilityModule":  {
                "EligibilityCriteria": "",
                "HealthyVolunteers":   "",
                "Gender":              "",
                "MinimumAge":          "",
                "MaximumAge":          "",
                "StdAgeList":           {
                    "StdAge": [""]
                }
            },
            "ContactsLocationsModule":  {
                "OverallOfficialList":  {
                    "OverallOfficial": [
                        // {
                        //     "OverallOfficialName":        "",
                        //     "OverallOfficialAffiliation": "",
                        //     "OverallOfficialRole":        ""
                        // }
                    ]
                },
                "LocationList":  {
                    "Location": [
                        // {
                        //     "LocationFacility": "",
                        //     "LocationCity":     "",
                        //     "LocationCountry":  ""
                        // }
                    ]
                }
            },
            "ReferencesModule":  {
                "ReferenceList":  {
                    "Reference": [
                        // {
                        //     "ReferencePMID":     "",
                        //     "ReferenceType":     "",
                        //     "ReferenceCitation": ""
                        // }
                    ]
                }
            }
        },
        "ResultsSection":  {
            "BaselineCharacteristicsModule":  {
                "BaselinePopulationDescription": "",
                "BaselineTypeUnitsAnalyzed":     "",
                "BaselineGroupList":              {
                    "BaselineGroup": [
                        // {
                        //     "BaselineGroupID":         "",
                        //     "BaselineGroupTitle":      "",
                        //     "BaselineGroupDescription":""
                        // }
                    ]
                },
                "BaselineDenomList":  {
                    "BaselineDenom": [
                        // {
                        //     "BaselineDenomUnits": "",
                        //     "BaselineDenomCountList":  {
                        //         "BaselineDenomCount": [{
                        //             "BaselineDenomCountGroupID": "",
                        //             "BaselineDenomCountValue":   ""
                        //         }]
                        //     }
                        // }
                    ]
                },
                "BaselineMeasureList":  {
                    "BaselineMeasure": [
                        // {
                        //     "BaselineMeasureTitle":                 "",
                        //     "BaselineMeasureDescription":           "",
                        //     "BaselineMeasureParamType":             "",
                        //     "BaselineMeasureDispersionType":        "",
                        //     "BaselineMeasureUnitOfMeasure":         "",
                        //     "BaselineMeasureCalculatePct":          "",
                        //     "BaselineMeasurePopulationDescription": "",
                        //     "BaselineMeasureDenomUnitsSelected":    "",
                        //     "BaselineClassList":                     {
                        //         "BaselineClass": [{
                        //             "BaselineClassTitle":   "",
                        //             "BaselineCategoryList":  {
                        //                 "BaselineCategory": [{
                        //                     "BaselineCategoryTitle": "",
                        //                     "BaselineMeasurementList":  {
                        //                         "BaselineMeasurement": [{
                        //                             "BaselineMeasurementGroupID": "",
                        //                             "BaselineMeasurementValue": "",
                        //                             "BaselineMeasurementSpread": "",
                        //                             "BaselineMeasurementComment": "",
                        //                             "BaselineMeasurementUpperLimit": "",
                        //                             "BaselineMeasurementLowerLimit": ""
                        //                         }]
                        //                     }
                        //                 }]
                        //             }
                        //         }]
                        //     },
                        //     "BaselineMeasureDenomList":  {
                        //         "BaselineMeasureDenom": [{
                        //             "BaselineMeasureDenomUnits": "",
                        //             "BaselineMeasureDenomCountList":  {
                        //                 "BaselineMeasureDenomCount": [{
                        //                     "BaselineMeasureDenomCountGroupID": "",
                        //                     "BaselineMeasureDenomCountValue":   ""
                        //                 }]
                        //             }
                        //         }]
                        //     }
                        // }
                    ]
                }
            },
            "OutcomeMeasuresModule":  {
                "OutcomeMeasureList":  {
                    "OutcomeMeasure": [
                        // {
                        //     "OutcomeMeasureType": "",
                        //     "OutcomeMeasureTitle": "",
                        //     "OutcomeMeasureDescription": "",
                        //     "OutcomeMeasurePopulationDescription": "",
                        //     "OutcomeMeasureReportingStatus": "",
                        //     "OutcomeMeasureParamType": "",
                        //     "OutcomeMeasureDispersionType": "",
                        //     "OutcomeMeasureUnitOfMeasure": "",
                        //     "OutcomeMeasureTimeFrame": "",
                        //     "OutcomeMeasureAnticipatedPostingDate": "",
                        //     "OutcomeMeasureCalculatePct": "",
                        //     "OutcomeMeasureDenomUnitsSelected": "",
                        //     "OutcomeMeasureTypeUnitsAnalyzed": "",
                        //     "OutcomeGroupList": {
                        //         "OutcomeGroup": [{
                        //             "OutcomeGroupID": "",
                        //             "OutcomeGroupTitle": "",
                        //             "OutcomeGroupDescription": ""
                        //         }]
                        //     },
                        //     "OutcomeDenomList":  {
                        //         "OutcomeDenom": [{
                        //             "OutcomeDenomUnits": "",
                        //             "OutcomeDenomCountList": {
                        //                 "OutcomeDenomCount": [{
                        //                     "OutcomeDenomCountGroupID": "",
                        //                     "OutcomeDenomCountValue  ": ""
                        //                 }]
                        //             }
                        //         }]
                        //     },
                        //     "OutcomeClassList":  {
                        //         "OutcomeClass": [{
                        //             "OutcomeClassTitle": "",
                        //             "OutcomeClassDenomList":  {
                        //                 "OutcomeClassDenom": [{
                        //                     "OutcomeClassDenomUnits": "",
                        //                     "OutcomeClassDenomCountList":  {
                        //                         "OutcomeClassDenomCount": [{
                        //                             "OutcomeClassDenomCountGroupID": "",
                        //                             "OutcomeClassDenomCountValue": ""
                        //                         }]
                        //                     }
                        //                 }]
                        //             },
                        //             "OutcomeCategoryList":  {
                        //                 "OutcomeCategory": [{
                        //                     "OutcomeCategoryTitle": "",
                        //                     "OutcomeMeasurementList":  {
                        //                         "OutcomeMeasurement": [{
                        //                             "OutcomeMeasurementGroupID": "",
                        //                             "OutcomeMeasurementValue": "",
                        //                             "OutcomeMeasurementLowerLimit": "",
                        //                             "OutcomeMeasurementUpperLimit": "",
                        //                             "OutcomeMeasurementSpread": "",
                        //                             "OutcomeMeasurementComment": ""
                        //                         }]
                        //                     }
                        //                 }]
                        //             }
                        //         }]
                        //     },
                        //     "OutcomeAnalysisList":  {
                        //         "OutcomeAnalysis": [{
                        //             "OutcomeAnalysisGroupIDList":  {
                        //                 "OutcomeAnalysisGroupID": [""]
                        //             },
                        //             "OutcomeAnalysisGroupDescription": "",
                        //             "OutcomeAnalysisNonInferiorityType": "",
                        //             "OutcomeAnalysisNonInferiorityComment": "",
                        //             "OutcomeAnalysisParamType": "",
                        //             "OutcomeAnalysisParamValue": "",
                        //             "OutcomeAnalysisCIPctValue": "",
                        //             "OutcomeAnalysisCINumSides": "",
                        //             "OutcomeAnalysisCILowerLimit": "",
                        //             "OutcomeAnalysisCILowerLimitComment": "",
                        //             "OutcomeAnalysisCIUpperLimit": "",
                        //             "OutcomeAnalysisCIUpperLimitComment": "",
                        //             "OutcomeAnalysisDispersionType": "",
                        //             "OutcomeAnalysisDispersionValue": "",
                        //             "OutcomeAnalysisDescription": "",
                        //             "OutcomeAnalysisPValue": "",
                        //             "OutcomeAnalysisStatisticalComment": "",
                        //             "OutcomeAnalysisStatisticalMethod": "",
                        //             "OutcomeAnalysisTestedNonInferiority": ""
                        //         }]
                        //     }
                        // }
                    ]
                }
            },
            "AdverseEventsModule":  {
                "EventsFrequencyThreshold": "",
                "EventsTimeFrame": "",
                "EventsDescription": "",
                "EventGroupList":            {
                    "EventGroup": [
                        // {
                        //     "EventGroupID": "",
                        //     "EventGroupTitle": "",
                        //     "EventGroupDescription": "",
                        //     "EventGroupDeathsNumAffected": "",
                        //     "EventGroupDeathsNumAtRisk": "",
                        //     "EventGroupSeriousNumAffected": "",
                        //     "EventGroupSeriousNumAtRisk": "",
                        //     "EventGroupOtherNumAffected": "",
                        //     "EventGroupOtherNumAtRisk": ""
                        // }
                    ]
                },
                "SeriousEventList":  {
                    "SeriousEvent": [
                        // {
                        //     "SeriousEventTerm": "",
                        //     "SeriousEventOrganSystem": "",
                        //     "SeriousEventSourceVocabulary": "",
                        //     "SeriousEventAssessmentType": "",
                        //     "SeriousEventNotes": "",
                        //     "SeriousEventStatsList":         {
                        //         "SeriousEventStats": [{
                        //             "SeriousEventStatsGroupID": "",
                        //             "SeriousEventStatsNumEvents": "",
                        //             "SeriousEventStatsNumAffected": "",
                        //             "SeriousEventStatsNumAtRisk": ""
                        //         }]
                        //     }
                        // }
                    ]
                },
                "OtherEventList":  {
                    "OtherEvent": [
                        // {
                        //     "OtherEventTerm": "",
                        //     "OtherEventOrganSystem": "",
                        //     "OtherEventSourceVocabulary": "",
                        //     "OtherEventAssessmentType": "",
                        //     "OtherEventNotes": "",
                        //     "OtherEventStatsList":         {
                        //         "OtherEventStats": [{
                        //             "OtherEventStatsGroupID": "",
                        //             "OtherEventStatsNumEvents": "",
                        //             "OtherEventStatsNumAffected": "",
                        //             "OtherEventStatsNumAtRisk": ""
                        //         }]
                        //     }
                        // }
                    ]
                }
            },
            "MoreInfoModule":  {
                "LimitationsAndCaveatsDescription": "",
                "PointOfContact":  {
                    "PointOfContactTitle": "",
                    "PointOfContactOrganization": "",
                    "PointOfContactEmail":  "",
                    "PointOfContactPhone": "",
                    "PointOfContactPhoneExt":  ""
                }
            }
        },
        "DerivedSection":  {
            "MiscInfoModule":  {
                "VersionHolder": ""
            },
            "ConditionBrowseModule":  {
                "ConditionMeshList":  {
                    "ConditionMesh": [
                        // {
                        //     "ConditionMeshID": "",
                        //     "ConditionMeshTerm": ""
                        // }
                    ]
                },
                "ConditionAncestorList":  {
                    "ConditionAncestor": [
                        // {
                        //     "ConditionAncestorID": "",
                        //     "ConditionAncestorTerm": ""
                        // }
                    ]
                }
            }
        }
    }

    function addCondition() {
        if(selected_condition_add === ""){return;}
        data.ProtocolSection.ConditionsModule.ConditionList.Condition.push(selected_condition_add);
        data.ProtocolSection.ConditionsModule.ConditionList.Condition.sort();
        selected_condition_add = "";
    } let selected_condition_add = "";

    function removeCondition() {
        data.ProtocolSection.ConditionsModule.ConditionList.Condition = data.ProtocolSection.ConditionsModule.ConditionList.Condition.filter(condition => condition !== selected_condition_remove);
        selected_condition_remove = "";
    } let selected_condition_remove = "";

    function addPhase() {
        if(selected_phase_add === ""){return;}
        data.ProtocolSection.DesignModule.PhaseList.Phase.push(selected_phase_add);
        data.ProtocolSection.DesignModule.PhaseList.Phase.sort();
        selected_phase_add = "";
    } let selected_phase_add = "";

    function removePhase() {
        data.ProtocolSection.DesignModule.PhaseList.Phase = data.ProtocolSection.DesignModule.PhaseList.Phase.filter(condition => condition !== selected_phase_remove);
        selected_phase_remove = "";
    } let selected_phase_remove = "";

    function addDesignWhoMasked() {
        if(selected_design_who_masked_add === ""){return;}
        data.ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignWhoMaskedList.DesignWhoMasked.push(selected_design_who_masked_add);
        data.ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignWhoMaskedList.DesignWhoMasked.sort();
        selected_design_who_masked_add = "";
    } let selected_design_who_masked_add = "";

    function removeDesignWhoMasked() {
        data.ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignWhoMaskedList.DesignWhoMasked = data.ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignWhoMaskedList.DesignWhoMasked.filter(condition => condition !== selected_design_who_masked_remove);
        selected_design_who_masked_remove = "";
    } let selected_design_who_masked_remove = "";

    let default_arm_group = {
        "ArmGroupLabel": "",
        "ArmGroupType": "",
        "ArmGroupDescription": "",
        "ArmGroupInterventionList":  {
            "ArmGroupInterventionName": []
        }
    };
    function addArmGroup() {
        if(selected_arm_group_add === default_arm_group){return;}
        data.ProtocolSection.ArmsInterventionsModule.ArmGroupList.ArmGroup.push(selected_arm_group_add);
        data.ProtocolSection.ArmsInterventionsModule.ArmGroupList.ArmGroup.sort();
        selected_arm_group_add = default_arm_group;
    } let selected_arm_group_add = default_arm_group;

    function removeArmGroup(i) {
        data.ProtocolSection.ArmsInterventionsModule.ArmGroupList.ArmGroup.splice(i, 1);
    }

    function addArmGroupInterventionName(i) {
        if(selected_arm_group_intervention_name_add === ""){return;}
        data.ProtocolSection.ArmsInterventionsModule.ArmGroupList.ArmGroup[i].ArmGroupInterventionList.ArmGroupInterventionName.push(selected_arm_group_intervention_name_add);
        data.ProtocolSection.ArmsInterventionsModule.ArmGroupList.ArmGroup[i].ArmGroupInterventionList.ArmGroupInterventionName.sort();
        selected_arm_group_intervention_name_add = "";
    } let selected_arm_group_intervention_name_add = "";

    function removeArmGroupInterventionName(i, j) {
        data.ProtocolSection.ArmsInterventionsModule.ArmGroupList.ArmGroup[i].ArmGroupInterventionList.ArmGroupInterventionName.splice(j, 1);
    }

    const default_intervention = {
        "InterventionType": "",
        "InterventionName": "",
        "InterventionDescription": "",
        "InterventionArmGroupLabelList":  {
            "InterventionArmGroupLabel": []
        }
    };
    function addIntervention() {
        if(selected_intervention_add === default_intervention){return;}
        data.ProtocolSection.ArmsInterventionsModule.InterventionList.Intervention.push(selected_intervention_add);
        data.ProtocolSection.ArmsInterventionsModule.InterventionList.Intervention.sort();
        selected_intervention_add = default_intervention;
    } let selected_intervention_add = default_intervention;

    function removeIntervention(i) {
        data.ProtocolSection.ArmsInterventionsModule.InterventionList.Intervention.splice(i, 1);
    }

    function addInterventionArmGroupLabel(i) {
        if(selected_intervention_arm_group_label_add === ""){return;}
        data.ProtocolSection.ArmsInterventionsModule.InterventionList.Intervention[i].InterventionArmGroupLabelList.InterventionArmGroupLabel.push(selected_intervention_arm_group_label_add);
        data.ProtocolSection.ArmsInterventionsModule.InterventionList.Intervention[i].InterventionArmGroupLabelList.InterventionArmGroupLabel.sort();
        selected_intervention_arm_group_label_add = "";
    } let selected_intervention_arm_group_label_add = "";

    function removeInterventionArmGroupLabel(i, j) {
        data.ProtocolSection.ArmsInterventionsModule.InterventionList.Intervention[i].InterventionArmGroupLabelList.InterventionArmGroupLabel.splice(j, 1);
    }

    const default_primary_outcome = {
        "PrimaryOutcomeMeasure": "",
        "PrimaryOutcomeDescription": "",
        "PrimaryOutcomeTimeFrame": ""
    }
    function addPrimaryOutcome() {
        if(selected_primary_outcome === default_primary_outcome) { return; }
        data.ProtocolSection.OutcomesModule.PrimaryOutcomeList.PrimaryOutcome.push(selected_primary_outcome);
        data.ProtocolSection.OutcomesModule.PrimaryOutcomeList.PrimaryOutcome.sort();
        selected_primary_outcome = default_primary_outcome;
    } let selected_primary_outcome = default_primary_outcome;

    function removePrimaryOutcome(i) {
        data.ProtocolSection.OutcomesModule.PrimaryOutcomeList.PrimaryOutcome.splice(i, 1);
    }

    const default_secondary_outcome = {
        "SecondaryOutcomeMeasure": "",
        "SecondaryOutcomeDescription": "",
        "SecondaryOutcomeTimeFrame": ""
    }
    function addSecondaryOutcome() {
        if(selected_secondary_outcome === default_secondary_outcome) { return; }
        data.ProtocolSection.OutcomesModule.SecondaryOutcomeList.SecondaryOutcome.push(selected_secondary_outcome);
        data.ProtocolSection.OutcomesModule.SecondaryOutcomeList.SecondaryOutcome.sort();
        selected_secondary_outcome = default_secondary_outcome;
    } let selected_secondary_outcome = default_secondary_outcome;

    function removeSecondaryOutcome(i) {
        data.ProtocolSection.OutcomesModule.SecondaryOutcomeList.SecondaryOutcome.splice(i, 1);
    }

    const default_other_outcome = {
        "OtherOutcomeMeasure": "",
        "OtherOutcomeDescription": "",
        "OtherOutcomeTimeFrame": ""
    }
    function addOtherOutcome() {
        if(selected_other_outcome === default_other_outcome) { return; }
        data.ProtocolSection.OutcomesModule.OtherOutcomeList.OtherOutcome.push(selected_other_outcome);
        data.ProtocolSection.OutcomesModule.OtherOutcomeList.OtherOutcome.sort();
        selected_other_outcome = default_other_outcome;
    } let selected_other_outcome = default_other_outcome;

    function removeOtherOutcome(i) {
        data.ProtocolSection.OutcomesModule.OtherOutcomeList.OtherOutcome.splice(i, 1);
    }

    let std_age_list = ["Child","Adult","Older Adult"];

    function addStdAge() {
        if(selected_std_age_add === "") { return; }
        data.ProtocolSection.EligibilityModule.StdAgeList.StdAge.push(selected_std_age_add);
        std_age_list = std_age_list.filter(age => age !== selected_std_age_add);
        selected_std_age_add = "";
    } let selected_std_age_add = "";

    function removeStdAge() {
        data.ProtocolSection.EligibilityModule.StdAgeList.StdAge = data.ProtocolSection.EligibilityModule.StdAgeList.StdAge.filter(age => age !== selected_std_age_remove);
        std_age_list.push(selected_std_age_remove);
        selected_std_age_remove = "";
    } let selected_std_age_remove = "";

    const default_overall_official = {
        "OverallOfficialName":        "",
        "OverallOfficialAffiliation": "",
        "OverallOfficialRole":        ""
    }
    function addOverallOfficial() {
        if(selected_overall_official === default_overall_official) { return; }
        data.ProtocolSection.ContactsLocationsModule.OverallOfficialList.OverallOfficial.push(selected_overall_official);
        data.ProtocolSection.ContactsLocationsModule.OverallOfficialList.OverallOfficial.sort();
        selected_overall_official = default_overall_official;
    } let selected_overall_official = default_overall_official;

    function removeOverallOfficial(i) {
        data.ProtocolSection.ContactsLocationsModule.OverallOfficialList.OverallOfficial.splice(i, 1);
    }

    const default_location = {
        "LocationFacility": "",
        "LocationCity":     "",
        "LocationCountry":  ""
    }
    function addLocation() {
        if(selected_location === default_location) { return; }
        data.ProtocolSection.ContactsLocationsModule.LocationList.Location.push(selected_location);
        data.ProtocolSection.ContactsLocationsModule.LocationList.Location.sort();
        selected_location = default_location;
    } let selected_location = default_location;

    function removeLocation(i) {
        data.ProtocolSection.ContactsLocationsModule.LocationList.Location.splice(i, 1);
    }

    const default_reference = {
        "ReferencePMID":     "",
        "ReferenceType":     "",
        "ReferenceCitation": ""
    }
    function addReference() {
        if(selected_reference === default_reference) { return; }
        data.ProtocolSection.ReferencesModule.ReferenceList.Reference.push(selected_location);
        data.ProtocolSection.ReferencesModule.ReferenceList.Reference.sort();
        selected_reference = default_reference;
    } let selected_reference = default_reference;

    function removeReference(i) {
        data.ProtocolSection.ReferencesModule.ReferenceList.Reference.splice(i, 1);
    }

    const default_baseline_group = {
        "BaselineGroupID":         "",
        "BaselineGroupTitle":      "",
        "BaselineGroupDescription":""
    }
    function addBaselineGroup() {
        if(selected_baseline_group === default_baseline_group) { return; }
        data.ResultsSection.BaselineCharacteristicsModule.BaselineGroupList.BaselineGroup.push(selected_baseline_group);
        data.ResultsSection.BaselineCharacteristicsModule.BaselineGroupList.BaselineGroup.sort();
        selected_baseline_group = default_baseline_group;
    } let selected_baseline_group = default_baseline_group;

    function removeBaselineGroup(i) {
        data.ResultsSection.BaselineCharacteristicsModule.BaselineGroupList.BaselineGroup.splice(i, 1);
    }

    const default_baseline_denom = {
        "BaselineDenomUnits": "",
        "BaselineDenomCountList":  {
            "BaselineDenomCount": []
        }
    }
    function addBaselineDenom() {
        if(selected_baseline_denom === default_baseline_denom) { return; }
        data.ResultsSection.BaselineCharacteristicsModule.BaselineDenomList.BaselineDenom.push(selected_baseline_denom);
        data.ResultsSection.BaselineCharacteristicsModule.BaselineDenomList.BaselineDenom.sort();
        selected_baseline_denom = default_baseline_denom;
    } let selected_baseline_denom = default_baseline_denom;

    function removeBaselineDenom(i) {
        data.ResultsSection.BaselineCharacteristicsModule.BaselineDenomList.BaselineDenom.splice(i, 1);
    }

    const default_baseline_denom_count = {
        "BaselineDenomCountGroupID": "",
        "BaselineDenomCountValue":   ""
    }
    function addBaselineDenomCount(i) {
        if(selected_baseline_denom_count === default_baseline_denom_count) { return; }
        data.ResultsSection.BaselineCharacteristicsModule.BaselineDenomList.BaselineDenom[i].BaselineDenomCountList.BaselineDenomCount.push(selected_baseline_denom);
        data.ResultsSection.BaselineCharacteristicsModule.BaselineDenomList.BaselineDenom[i].BaselineDenomCountList.BaselineDenomCount.sort();
        selected_baseline_denom_count = default_baseline_denom_count;
    } let selected_baseline_denom_count = default_baseline_denom_count;

    function removeBaselineDenomCount(i, j) {
        data.ResultsSection.BaselineCharacteristicsModule.BaselineDenomList.BaselineDenom[i].BaselineDenomCountList.BaselineDenomCount.splice(j, 1);
    }

    const default_baseline_measure = {
        "BaselineMeasureTitle":                 "",
        "BaselineMeasureDescription":           "",
        "BaselineMeasureParamType":             "",
        "BaselineMeasureDispersionType":        "",
        "BaselineMeasureUnitOfMeasure":         "",
        "BaselineMeasureCalculatePct":          "",
        "BaselineMeasurePopulationDescription": "",
        "BaselineMeasureDenomUnitsSelected":    "",
        "BaselineClassList":                     {
            "BaselineClass": []
        },
        "BaselineMeasureDenomList":  {
            "BaselineMeasureDenom": []
        }
    }
    function addBaselineMeasure() {
        if(selected_baseline_measure === default_baseline_measure) { return; }
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure.push(selected_baseline_measure);
        data.ResultsSection.BaselineCharacteristicsModule.BaselineDenomList.BaselineDenom.sort();
        selected_baseline_measure = default_baseline_measure;
    } let selected_baseline_measure = default_baseline_measure;

    function removeBaselineMeasure(i) {
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure.splice(i, 1);
    }

    const default_baseline_class = {
        "BaselineClassTitle":   "",
        "BaselineCategoryList":  {
            "BaselineCategory": []
        }
    }
    function addBaselineClass(i) {
        if(selected_baseline_class === default_baseline_class) { return; }
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineClassList.BaselineClass.push(selected_baseline_class);
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineClassList.BaselineClass.sort();
        selected_baseline_class = default_baseline_class;
    } let selected_baseline_class = default_baseline_class;

    function removeBaselineClass(i, j) {
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineClassList.BaselineClass.splice(j, 1);
    }

    const default_baseline_category = {
        "BaselineCategoryTitle": "",
        "BaselineMeasurementList":  {
            "BaselineMeasurement": []
        }
    }
    function addBaselineCategory(i, j) {
        if(selected_baseline_category === default_baseline_category) { return; }
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineClassList.BaselineClass[j].BaselineCategoryList.BaselineCategory.push(selected_baseline_category);
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineClassList.BaselineClass[j].BaselineCategoryList.BaselineCategory.sort();
        selected_baseline_category = default_baseline_category;
    } let selected_baseline_category = default_baseline_category;

    function removeBaselineCategory(i, j, k) {
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineClassList.BaselineClass[j].BaselineCategoryList.BaselineCategory.splice(k, 1);
    }

    const default_baseline_measurement = {
        "BaselineMeasurementGroupID": "",
        "BaselineMeasurementValue": "",
        "BaselineMeasurementSpread": "",
        "BaselineMeasurementComment": "",
        "BaselineMeasurementUpperLimit": "",
        "BaselineMeasurementLowerLimit": ""
    }
    function addBaselineMeasurement(i, j, k) {
        if(selected_baseline_measurement === default_baseline_measurement) { return; }
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineClassList.BaselineClass[j].BaselineCategoryList.BaselineCategory[k].BaselineMeasurementList.BaselineMeasurement.push(selected_baseline_category);
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineClassList.BaselineClass[j].BaselineCategoryList.BaselineCategory[k].BaselineMeasurementList.BaselineMeasurement.sort();
        selected_baseline_measurement = default_baseline_measurement;
    } let selected_baseline_measurement = default_baseline_measurement;

    function removeBaselineMeasurement(i, j, k, l) {
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineClassList.BaselineClass[j].BaselineCategoryList.BaselineCategory[k].BaselineMeasurementList.BaselineMeasurement.splice(l, 1);
    }

    const default_baseline_measure_denom = {
        "BaselineMeasureDenomUnits": "",
        "BaselineMeasureDenomCountList":  {
            "BaselineMeasureDenomCount": []
        }
    }
    function addBaselineMeasureDenom(i) {
        if(selected_baseline_measure_denom === default_baseline_measure_denom) { return; }
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineMeasureDenomList.BaselineMeasureDenom.push(selected_baseline_measure_denom);
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineMeasureDenomList.BaselineMeasureDenom.sort();
        selected_baseline_measure_denom = default_baseline_measure_denom;
    } let selected_baseline_measure_denom = default_baseline_measure_denom;

    function removeBaselineMeasureDenom(i, j) {
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineMeasureDenomList.BaselineMeasureDenom.splice(j, 1);
    }

    const default_baseline_measure_denom_count = {
        "BaselineMeasureDenomCountGroupID": "",
        "BaselineMeasureDenomCountValue":   ""
    }
    function addBaselineMeasureDenomCount(i, j) {
        if(selected_baseline_measure_denom_count === default_baseline_measure_denom_count) { return; }
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineMeasureDenomList.BaselineMeasureDenom[j].BaselineMeasureDenomCountList.BaselineMeasureDenomCount.push(selected_baseline_measure_denom_count);
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineMeasureDenomList.BaselineMeasureDenom[j].BaselineMeasureDenomCountList.BaselineMeasureDenomCount.sort();
        selected_baseline_measure_denom_count = default_baseline_measure_denom_count;
    } let selected_baseline_measure_denom_count = default_baseline_measure_denom_count;

    function removeBaselineMeasureDenomCount(i, j, k) {
        data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure[i].BaselineMeasureDenomList.BaselineMeasureDenom[j].BaselineMeasureDenomCountList.BaselineMeasureDenomCount.splice(k, 1);
    }

    const default_outcome_measure = {
        "OutcomeMeasureType": "",
        "OutcomeMeasureTitle": "",
        "OutcomeMeasureDescription": "",
        "OutcomeMeasurePopulationDescription": "",
        "OutcomeMeasureReportingStatus": "",
        "OutcomeMeasureParamType": "",
        "OutcomeMeasureDispersionType": "",
        "OutcomeMeasureUnitOfMeasure": "",
        "OutcomeMeasureTimeFrame": "",
        "OutcomeMeasureAnticipatedPostingDate": "",
        "OutcomeMeasureCalculatePct": "",
        "OutcomeMeasureDenomUnitsSelected": "",
        "OutcomeMeasureTypeUnitsAnalyzed": "",
        "OutcomeGroupList": {
            "OutcomeGroup": []
        },
        "OutcomeDenomList":  {
            "OutcomeDenom": []
        },
        "OutcomeClassList":  {
            "OutcomeClass": []
        },
        "OutcomeAnalysisList":  {
            "OutcomeAnalysis": []
        }
    }
    function add_outcome_measure() {
        if(selected_outcome_measure === default_outcome_measure) { return; }
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure.push(selected_outcome_measure);
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure.sort();
        selected_outcome_measure = default_outcome_measure;
    } let selected_outcome_measure = default_outcome_measure;

    function remove_outcome_measure(i) {
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure.splice(i, 1);
    }

    const default_outcome_group = {
        "OutcomeGroupID": "",
        "OutcomeGroupTitle": "",
        "OutcomeGroupDescription": ""
    }
    function add_outcome_group(i) {
        if(selected_outcome_group === default_outcome_group) { return; }
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeGroupList.OutcomeGroup.push(selected_outcome_group);
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeGroupList.OutcomeGroup.sort();
        selected_outcome_group = default_outcome_group;
    } let selected_outcome_group = default_outcome_group;

    function remove_outcome_group(i, j) {
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeGroupList.OutcomeGroup.splice(j, 1);
    }

    const default_outcome_denom = {
        "OutcomeDenomUnits": "",
        "OutcomeDenomCountList": {
            "OutcomeDenomCount": []
        }
    }
    function add_outcome_denom(i) {
        if(selected_outcome_denom === default_outcome_denom) { return; }
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeDenomList.OutcomeDenom.push(selected_outcome_denom);
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeDenomList.OutcomeDenom.sort();
        selected_outcome_denom = default_outcome_denom;
    } let selected_outcome_denom = default_outcome_denom;

    function remove_outcome_denom(i, j) {
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeDenomList.OutcomeDenom.splice(j, 1);
    }

    const default_outcome_denom_count = {
        "OutcomeDenomCountGroupID": "",
        "OutcomeDenomCountValue  ": ""
    }
    function add_outcome_denom_count(i, j) {
        if(selected_outcome_denom_count === default_outcome_denom_count) { return; }
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeDenomList.OutcomeDenom[j].OutcomeDenomCountList.OutcomeDenomCount.push(selected_outcome_denom_count);
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeDenomList.OutcomeDenom[j].OutcomeDenomCountList.OutcomeDenomCount.sort();
        selected_outcome_denom_count = default_outcome_denom_count;
    } let selected_outcome_denom_count = default_outcome_denom_count;

    function remove_outcome_denom_count(i, j, k) {
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeDenomList.OutcomeDenom[j].OutcomeDenomCountList.OutcomeDenomCount.splice(k, 1);
    }

    const default_outcome_class = {
        "OutcomeClassTitle": "",
        "OutcomeClassDenomList":  {
            "OutcomeClassDenom": []
        },
        "OutcomeCategoryList":  {
            "OutcomeCategory": []
        },
    }
    function add_outcome_class(i) {
        if(selected_outcome_class === default_outcome_class) { return; }
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass.push(selected_outcome_class);
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass.sort();
        selected_outcome_class = default_outcome_class;
    } let selected_outcome_class = default_outcome_class;

    function remove_outcome_class(i, j) {
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass.splice(j, 1);
    }

    const default_outcome_class_denom = {
        "OutcomeClassDenomUnits": "",
        "OutcomeClassDenomCountList":  {
            "OutcomeClassDenomCount": []
        }
    }
    function add_outcome_class_denom(i, j) {
        if(selected_outcome_class_denom === default_outcome_class_denom) { return; }
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass[j].OutcomeClassDenomList.OutcomeClassDenom.push(selected_outcome_class_denom);
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass[j].OutcomeClassDenomList.OutcomeClassDenom.sort();
        selected_outcome_class_denom = default_outcome_class_denom;
    } let selected_outcome_class_denom = default_outcome_class_denom;

    function remove_outcome_class_denom(i, j, k) {
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass[j].OutcomeClassDenomList.OutcomeClassDenom.splice(k, 1);
    }

    const default_outcome_class_denom_count = {
        "OutcomeClassDenomCountGroupID": "",
        "OutcomeClassDenomCountValue": ""
    }
    function add_outcome_class_denom_count(i, j, k) {
        if(selected_outcome_class_denom_count === default_outcome_class_denom_count) { return; }
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass[j].OutcomeClassDenomList.OutcomeClassDenom[k].OutcomeClassDenomCountList.OutcomeClassDenomCount.push(selected_outcome_class_denom_count);
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass[j].OutcomeClassDenomList.OutcomeClassDenom[k].OutcomeClassDenomCountList.OutcomeClassDenomCount.sort();
        selected_outcome_class_denom_count = default_outcome_class_denom_count;
    } let selected_outcome_class_denom_count = default_outcome_class_denom_count;

    function remove_outcome_class_denom_count(i, j, k, l) {
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass[j].OutcomeClassDenomList.OutcomeClassDenom[k].OutcomeClassDenomCountList.OutcomeClassDenomCount.splice(l, 1);
    }

    const default_outcome_category = {
        "OutcomeCategoryTitle": "",
        "OutcomeMeasurementList":  {
            "OutcomeMeasurement": []
        }
    }
    function add_outcome_category(i, j) {
        if(selected_outcome_category === default_outcome_category) { return; }
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass[j].OutcomeCategoryList.OutcomeCategory.push(selected_outcome_category);
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass[j].OutcomeCategoryList.OutcomeCategory.sort();
        selected_outcome_category = default_outcome_category;
    } let selected_outcome_category = default_outcome_category;

    function remove_outcome_category(i, j, k) {
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass[j].OutcomeCategoryList.OutcomeCategory.splice(k, 1);
    }

    const default_outcome_measurement = {
        "OutcomeMeasurementGroupID": "",
        "OutcomeMeasurementValue": "",
        "OutcomeMeasurementLowerLimit": "",
        "OutcomeMeasurementUpperLimit": "",
        "OutcomeMeasurementSpread": "",
        "OutcomeMeasurementComment": ""
    }
    function add_outcome_measurement(i, j, k) {
        if(selected_outcome_measurement === default_outcome_measurement) { return; }
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass[j].OutcomeCategoryList.OutcomeCategory[k].OutcomeMeasurementList.OutcomeMeasurement.push(selected_outcome_category);
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass[j].OutcomeCategoryList.OutcomeCategory[k].OutcomeMeasurementList.OutcomeMeasurement.sort();
        selected_outcome_measurement = default_outcome_measurement;
    } let selected_outcome_measurement = default_outcome_measurement;

    function remove_outcome_measurement(i, j, k, l) {
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeClassList.OutcomeClass[j].OutcomeCategoryList.OutcomeCategory[k].OutcomeMeasurementList.OutcomeMeasurement.splice(l, 1);
    }

    const default_outcome_analysis = {
        "OutcomeAnalysisGroupIDList":  {
            "OutcomeAnalysisGroupID": []
        },
        "OutcomeAnalysisGroupDescription": "",
        "OutcomeAnalysisNonInferiorityType": "",
        "OutcomeAnalysisNonInferiorityComment": "",
        "OutcomeAnalysisParamType": "",
        "OutcomeAnalysisParamValue": "",
        "OutcomeAnalysisCIPctValue": "",
        "OutcomeAnalysisCINumSides": "",
        "OutcomeAnalysisCILowerLimit": "",
        "OutcomeAnalysisCILowerLimitComment": "",
        "OutcomeAnalysisCIUpperLimit": "",
        "OutcomeAnalysisCIUpperLimitComment": "",
        "OutcomeAnalysisDispersionType": "",
        "OutcomeAnalysisDispersionValue": "",
        "OutcomeAnalysisDescription": "",
        "OutcomeAnalysisPValue": "",
        "OutcomeAnalysisStatisticalComment": "",
        "OutcomeAnalysisStatisticalMethod": "",
        "OutcomeAnalysisTestedNonInferiority": ""
    }
    function add_outcome_analysis(i) {
        if(selected_outcome_analysis === default_outcome_analysis) { return; }
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeAnalysisList.OutcomeAnalysis.push(selected_outcome_analysis);
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeAnalysisList.OutcomeAnalysis.sort();
        selected_outcome_analysis = default_outcome_analysis;
    } let selected_outcome_analysis = default_outcome_analysis;

    function remove_outcome_analysis(i, j) {
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeAnalysisList.OutcomeAnalysis.splice(j, 1);
    }

    function add_outcome_group_id(i, j) {
        if(selected_outcome_group_id === "") { return; }
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeAnalysisList.OutcomeAnalysis[j].OutcomeAnalysisGroupIDList.OutcomeAnalysisGroupID.push(selected_outcome_analysis);
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeAnalysisList.OutcomeAnalysis[j].OutcomeAnalysisGroupIDList.OutcomeAnalysisGroupID.sort();
        selected_outcome_group_id = "";
    } let selected_outcome_group_id = "";

    function remove_outcome_group_id(i, j, k) {
        data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure[i].OutcomeAnalysisList.OutcomeAnalysis[j].OutcomeAnalysisGroupIDList.OutcomeAnalysisGroupID.splice(k, 1);
    }

</script>

<main class="container is-fluid">
    <h1 style="text-align: center" class="title">New Trial - (WIP)</h1>
    <h2 class="title">Protocol Section</h2>
    <fieldset>
        <legend><h3 class="subtitle">Identification Module</h3></legend>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="nct_id">Study ID: (012345 becomes OT012345)</label></div>
            <div class="column is-three-quarters"><input id="nct_id" bind:value={data.ProtocolSection.IdentificationModule.NCTID} class="input" type="text" placeholder="Study ID"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="brief_title">Brief Title: </label></div>
            <div class="column is-three-quarters"><input id="brief_title" bind:value={data.ProtocolSection.IdentificationModule.BriefTitle} class="input" type="text" placeholder="Brief Title"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="official_title">Official Title: </label></div>
            <div class="column is-three-quarters"><input id="official_title" bind:value={data.ProtocolSection.IdentificationModule.OfficialTitle} class="input" type="text" placeholder="Official Title"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="organization_full_name">Organization Full Name: </label></div>
            <div class="column is-three-quarters"><input id="organization_full_name" bind:value={data.ProtocolSection.IdentificationModule.Organization.OrgFullName} class="input" type="text" placeholder="Organization Full Name"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="organization_class">Organization Class: </label></div>
            <div class="column is-three-quarters"><input id="organization_class" bind:value={data.ProtocolSection.IdentificationModule.Organization.OrgClass} class="input" type="text" placeholder="Organization Class"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="organization_study_id">Organization Study ID: </label></div>
            <div class="column is-three-quarters"><input id="organization_study_id" bind:value={data.ProtocolSection.IdentificationModule.OrgStudyIDInfo.OrgStudyID} class="input" type="text" placeholder="Organization Study ID"></div>
        </div>
    </fieldset>
    <fieldset>
        <legend><h3 class="subtitle">Status Module</h3></legend>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="status_verified_date">Status Verified Date:</label></div>
            <div class="column is-three-quarters"><input id="status_verified_date" bind:value={data.ProtocolSection.StatusModule.StatusVerifiedDate} class="input" type="text" placeholder="Status Verified Date"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="overall_status">Overall Status: </label></div>
            <div class="column is-three-quarters"><input id="overall_status" bind:value={data.ProtocolSection.StatusModule.OverallStatus} class="input" type="text" placeholder="Overall Status"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="start_date">Start Date: </label></div>
            <div class="column is-three-quarters"><input id="start_date" bind:value={data.ProtocolSection.StatusModule.StartDate.StartDate} class="input" type="text" placeholder="Start Date"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="primary_completion_date">Primary Completion Date: </label></div>
            <div class="column is-three-quarters"><input id="primary_completion_date" bind:value={data.ProtocolSection.StatusModule.PrimaryCompletionDate.PrimaryCompletionDate} class="input" type="text" placeholder="Primary Completion Date"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="completion_date">Completion Date: </label></div>
            <div class="column is-three-quarters"><input id="completion_date" bind:value={data.ProtocolSection.StatusModule.CompletionDate.CompletionDate} class="input" type="text" placeholder="Completion Date"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="study_first_submit_date">Study First Submit Date: </label></div>
            <div class="column is-three-quarters"><input id="study_first_submit_date" bind:value={data.ProtocolSection.StatusModule.StudyFirstSubmitDate.StudyFirstSubmitDate} class="input" type="text" placeholder="Study First Submit Date"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="study_first_post_date">Study First Post Date: </label></div>
            <div class="column is-three-quarters"><input id="study_first_post_date" bind:value={data.ProtocolSection.StatusModule.StudyFirstPostDate.StudyFirstPostDate} class="input" type="text" placeholder="Study First Post Date"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="results_first_submit_date">Results First Submit Date: </label></div>
            <div class="column is-three-quarters"><input id="results_first_submit_date" bind:value={data.ProtocolSection.StatusModule.ResultsFirstSubmitDate.ResultsFirstSubmitDate} class="input" type="text" placeholder="Results First Post Date"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="last_update_submit_date">Last Update Submit Date: </label></div>
            <div class="column is-three-quarters"><input id="last_update_submit_date" bind:value={data.ProtocolSection.StatusModule.LastUpdateSubmitDate.LastUpdateSubmitDate} class="input" type="text" placeholder="Last Update Submit Date"></div>
        </div>
    </fieldset>
    <fieldset>
        <legend><h3 class="subtitle">Sponsor Collaborators Module</h3></legend>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="responsible_party_type">Responsible Party Type:</label></div>
            <div class="column is-three-quarters"><input id="responsible_party_type" bind:value={data.ProtocolSection.SponsorCollaboratorsModule.ResponsibleParty.ResponsiblePartyType} class="input" type="text" placeholder="Responsible Party Type"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="lead_sponsor_name">Lead Sponsor Name: </label></div>
            <div class="column is-three-quarters"><input id="lead_sponsor_name" bind:value={data.ProtocolSection.SponsorCollaboratorsModule.LeadSponsor.LeadSponsorName} class="input" type="text" placeholder="Lead Sponsor Name"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="lead_sponsor_class">Lead Sponsor Class: </label></div>
            <div class="column is-three-quarters"><input id="lead_sponsor_class" bind:value={data.ProtocolSection.SponsorCollaboratorsModule.LeadSponsor.LeadSponsorClass} class="input" type="text" placeholder="Lead Sponsor Class"></div>
        </div>
    </fieldset>
    <fieldset>
        <legend><h3 class="subtitle">Oversight Module</h3></legend>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="oversight_has_dmc">Oversight Has DMC:</label></div>
            <div class="column is-three-quarters"><input id="oversight_has_dmc" bind:value={data.ProtocolSection.OversightModule.OversightHasDMC} class="input" type="text" placeholder="Oversight Has DMC"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="is_fda_regulated_drug">Is FDA Regulated Drug: </label></div>
            <div class="column is-three-quarters"><input id="is_fda_regulated_drug" bind:value={data.ProtocolSection.OversightModule.IsFDARegulatedDrug} class="input" type="text" placeholder="Is FDA Regulated Drug"></div>
        </div>
    </fieldset>
    <fieldset>
        <legend><h3 class="subtitle">Description Module</h3></legend>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="brief_summary">Brief Summary:</label></div>
            <div class="column is-three-quarters"><input id="brief_summary" bind:value={data.ProtocolSection.DescriptionModule.BriefSummary} class="input" type="text" placeholder="Brief Summary"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="detailed_description">Detailed Description: </label></div>
            <div class="column is-three-quarters"><textarea id="detailed_description" bind:value={data.ProtocolSection.DescriptionModule.DetailedDescription} class="input textarea" ></textarea></div>
        </div>
    </fieldset>
    <fieldset>
        <legend><h3 class="subtitle">Conditions Module</h3></legend>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Conditions:</div>
            <div class="column is-three-quarters">
                <div class="columns"><div class="column is-one-quarter"><label for="add_condition">Add Condition:</label></div><div class="column is-three-quarters"><input type="text" id="add_condition" name="add_condition" class="input" bind:value={selected_condition_add} /><button class="button" on:click={() => addCondition()}>Add</button></div></div>
                <div class="columns"><div class="column is-one-quarter"><label for="remove_condition">Remove Condition:</label></div><div class="column is-three-quarters"><select id="remove_condition" name="remove_condition" class="select input" bind:value={selected_condition_remove}><option selected="selected" disabled="disabled">Not Selected</option>{#each data.ProtocolSection.ConditionsModule.ConditionList.Condition || [] as condition}<option value="{condition}">{condition}</option>{/each}</select><button class="button" on:click={() => removeCondition()}>Remove</button></div></div>
            </div>
        </div>
    </fieldset>
    <fieldset>
        <legend><h3 class="subtitle">Design Module</h3></legend>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="study_type">Study Type:</label></div>
            <div class="column is-three-quarters"><input id="study_type" bind:value={data.ProtocolSection.DesignModule.StudyType} class="input" type="text" placeholder="Study Type"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Phase List:</div>
            <div class="column is-three-quarters">
                <div class="columns"><div class="column is-one-quarter"><label for="add_phase">Add Phase:</label></div><div class="column is-three-quarters"><input type="text" id="add_phase" name="add_phase" class="input" bind:value={selected_phase_add} /><button class="button" on:click={() => addPhase()}>Add</button></div></div>
                <div class="columns"><div class="column is-one-quarter"><label for="remove_phase">Remove Phase:</label></div><div class="column is-three-quarters"><select id="remove_phase" name="remove_phase" class="select input" bind:value={selected_phase_remove}><option selected="selected" disabled="disabled">Not Selected</option>{#each data.ProtocolSection.ConditionsModule.ConditionList.Condition || [] as phase}<option value="{phase}">{phase}</option>{/each}</select><button class="button" on:click={() => removePhase()}>Remove</button></div></div>
            </div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="design_allocation">Design Allocation:</label></div>
            <div class="column is-three-quarters"><input id="design_allocation" bind:value={data.ProtocolSection.DesignModule.DesignInfo.DesignAllocation} class="input" type="text" placeholder="Design Allocation"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="design_intervention_model">Design Intervention Model:</label></div>
            <div class="column is-three-quarters"><input id="design_intervention_model" bind:value={data.ProtocolSection.DesignModule.DesignInfo.DesignInterventionModel} class="input" type="text" placeholder="Design Intervention Model"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="design_primary_purpose">Design Primary Purpose:</label></div>
            <div class="column is-three-quarters"><input id="design_primary_purpose" bind:value={data.ProtocolSection.DesignModule.DesignInfo.DesignPrimaryPurpose} class="input" type="text" placeholder="Design Primary Purpose"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="design_masking">Design Masking:</label></div>
            <div class="column is-three-quarters"><input id="design_masking" bind:value={data.ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignMasking} class="input" type="text" placeholder="Design Masking"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Designed Who Masked List:</div>
            <div class="column is-three-quarters">
                <div class="columns"><div class="column is-one-quarter"><label for="add_design_who_masked">Add Design Who Masked:</label></div><div class="column is-three-quarters"><input type="text" id="add_design_who_masked" name="add_phase" class="input" bind:value={selected_design_who_masked_add} /><button class="button" on:click={() => addDesignWhoMasked()}>Add</button></div></div>
                <div class="columns"><div class="column is-one-quarter"><label for="remove_design_who_masked">Remove Design Who Masked:</label></div><div class="column is-three-quarters"><select id="remove_design_who_masked" name="remove_phase" class="select input" bind:value={selected_design_who_masked_remove}><option selected="selected" disabled="disabled">Not Selected</option>{#each data.ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignWhoMaskedList.DesignWhoMasked || [] as design_who_masked}<option value="{design_who_masked}">{design_who_masked}</option>{/each}</select><button class="button" on:click={() => removeDesignWhoMasked()}>Remove</button></div></div>
            </div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="enrollment_count">Enrollment Count:</label></div>
            <div class="column is-three-quarters"><input id="enrollment_count" bind:value={data.ProtocolSection.DesignModule.EnrollmentInfo.EnrollmentCount} class="input" type="text" placeholder="Enrollment Count"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="enrollment_type">Enrollment Type:</label></div>
            <div class="column is-three-quarters"><input id="enrollment_type" bind:value={data.ProtocolSection.DesignModule.EnrollmentInfo.EnrollmentType} class="input" type="text" placeholder="Enrollment Type"></div>
        </div>
    </fieldset>
    <fieldset>
        <legend><h3 class="subtitle">Arms Interventions Module</h3></legend>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Arm Group List:</div>
            <div class="column is-three-quarters">
                <div class="columns">
                    <div class="column is-one-quarter">Add Arm Group:</div>
                    <div class="column is-three-quarters">
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_arm_group_label">Label:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_arm_group_label" name="add_arm_group_label" class="input" bind:value={selected_arm_group_add.ArmGroupLabel} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_arm_group_type">Type:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_arm_group_type" name="add_arm_group_type" class="input" bind:value={selected_arm_group_add.ArmGroupLabel} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_arm_group_description">Description:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_arm_group_description" name="add_arm_group_description" class="input" bind:value={selected_arm_group_add.ArmGroupDescription} /></div>
                        </div>
                        <button class="button" on:click={() => addArmGroup()}>Add</button>
                    </div>
                </div>
                <div class="columns">
                    <div class="column is-one-quarter">Arm Groups:</div>
                    <div class="column is-three-quarters">
                        {#each data.ProtocolSection.ArmsInterventionsModule.ArmGroupList.ArmGroup || [] as armGroup, i }
                            <div class="columns"><div class="column is-one-quarter"><label for="arm_group_label_{i}">Label: </label></div><div class="column is-three-quarters"><input type="text" id="arm_group_label_{i}" name="arm_group_label_{i}" bind:value={armGroup.ArmGroupLabel} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="arm_group_type_{i}">Type: </label></div><div class="column is-three-quarters"><input type="text" id="arm_group_type_{i}" name="arm_group_type_{i}" bind:value={armGroup.ArmGroupType} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="arm_group_description_{i}">Description: </label></div><div class="column is-three-quarters"><input type="text" id="arm_group_description_{i}" name="arm_group_description_{i}" bind:value={armGroup.ArmGroupDescription} /></div></div>
                            <div class="columns">
                                <div class="column is-one-quarter">Intervention List:</div>
                                <div class="column is-three-quarters">
                                    <div class="columns"><div class="column is-one-quarter"><label for="add_arm_group_intervention_name">Add Name:</label></div><div class="column is-three-quarters"><input type="text" id="add_arm_group_intervention_name" name="add_phase" class="input" bind:value={selected_arm_group_intervention_name_add} /><button class="button" on:click={() => addArmGroupInterventionName(i)}>Add</button></div></div>
                                    <div class="columns">
                                        <div class="column is-one-quarter"><label for="remove_design_who_masked">Names:</label></div>
                                        <div class="column is-three-quarters">
                                            <ul>{#each armGroup.ArmGroupInterventionList.ArmGroupInterventionName || [] as interventionName, j}
                                                <li>{interventionName} - <button class="button is-danger" on:click={() => removeArmGroupInterventionName(i, j)}>-</button></li>
                                            {/each}</ul>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <button class="button is-danger" on:click={() => removeArmGroup(i)}>Remove</button>
                        {/each}
                    </div>
                </div>
            </div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Intervention List:</div>
            <div class="column is-three-quarters">
                <div class="columns">
                    <div class="column is-one-quarter">Add Intervention:</div>
                    <div class="column is-three-quarters">
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_intervention_type">Type:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_intervention_type" name="add_arm_group_type" class="input" bind:value={selected_intervention_add.InterventionType} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_intervention_name">Name:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_intervention_name" name="add_arm_group_label" class="input" bind:value={selected_intervention_add.InterventionName} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_intervention_description">Description:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_intervention_description" name="add_intervention_description" class="input" bind:value={selected_intervention_add.InterventionDescription} /></div>
                        </div>
                        <button class="button" on:click={() => addIntervention()}>Add</button>
                    </div>
                </div>
                <div class="columns">
                    <div class="column is-one-quarter">Interventions:</div>
                    <div class="column is-three-quarters">
                        {#each data.ProtocolSection.ArmsInterventionsModule.InterventionList.Intervention || [] as intervention, i }
                            <div class="columns"><div class="column is-one-quarter"><label for="intervention_type_{i}">Type: </label></div><div class="column is-three-quarters"><input type="text" id="intervention_type_{i}" name="intervention_type_{i}" bind:value={intervention.InterventionType} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="intervention_name_{i}">Name: </label></div><div class="column is-three-quarters"><input type="text" id="intervention_name_{i}" name="intervention_name_{i}" bind:value={intervention.InterventionName} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="intervention_description_{i}">Description: </label></div><div class="column is-three-quarters"><input type="text" id="intervention_description_{i}" name="intervention_description_{i}" bind:value={intervention.InterventionDescription} /></div></div>
                            <div class="columns">
                                <div class="column is-one-quarter">Intervention List:</div>
                                <div class="column is-three-quarters">
                                    <div class="columns"><div class="column is-one-quarter"><label for="add_intervention_arm_group_label">Add Group Label:</label></div><div class="column is-three-quarters"><input type="text" id="add_intervention_arm_group_label" name="add_phase" class="input" bind:value={selected_intervention_arm_group_label_add} /><button class="button" on:click={() => addInterventionArmGroupLabel(i)}>Add</button></div></div>
                                    <div class="columns">
                                        <div class="column is-one-quarter"><label for="remove_design_who_masked">Group Label:</label></div>
                                        <div class="column is-three-quarters">
                                            <ul>{#each intervention.InterventionArmGroupLabelList.InterventionArmGroupLabel || [] as groupLabel, j}
                                                <li>{groupLabel} - <button class="button is-danger" on:click={() => removeInterventionArmGroupLabel(i, j)}>-</button></li>
                                            {/each}</ul>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <button class="button is-danger" on:click={() => removeIntervention(i)}>Remove</button>
                        {/each}
                    </div>
                </div>
            </div>
        </div>
    </fieldset>
    <fieldset>
        <legend><h3 class="subtitle">Outcomes Module</h3></legend>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Primary Outcome List:</div>
            <div class="column is-three-quarters">
                <div class="columns">
                    <div class="column is-one-quarter">Add Primary Outcome:</div>
                    <div class="column is-three-quarters">
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_primary_outcome_measure">Measure:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_primary_outcome_measure" name="add_primary_outcome_measure" class="input" bind:value={selected_primary_outcome.PrimaryOutcomeMeasure} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_primary_outcome_description">Description:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_primary_outcome_description" name="add_primary_outcome_description" class="input" bind:value={selected_primary_outcome.PrimaryOutcomeDescription} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_primary_outcome_time_frame">Time Frame:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_primary_outcome_time_frame" name="add_primary_outcome_time_frame" class="input" bind:value={selected_primary_outcome.PrimaryOutcomeTimeFrame} /></div>
                        </div>
                        <button class="button" on:click={() => addPrimaryOutcome()}>Add</button>
                    </div>
                </div>
                <div class="columns">
                    <div class="column is-one-quarter">Primary Outcomes:</div>
                    <div class="column is-three-quarters">
                        {#each data.ProtocolSection.OutcomesModule.PrimaryOutcomeList.PrimaryOutcome || [] as primaryOutcome, i }
                            <div class="columns"><div class="column is-one-quarter"><label for="primary_outcome_measure_{i}">Measure: </label></div><div class="column is-three-quarters"><input type="text" id="primary_outcome_measure_{i}" name="primary_outcome_measure_{i}" bind:value={primaryOutcome.PrimaryOutcomeMeasure} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="primary_outcome_description_{i}">Description: </label></div><div class="column is-three-quarters"><input type="text" id="primary_outcome_description_{i}" name="primary_outcome_description_{i}" bind:value={primaryOutcome.PrimaryOutcomeDescription} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="primary_outcome_time_frame_{i}">Time Frame: </label></div><div class="column is-three-quarters"><input type="text" id="primary_outcome_time_frame_{i}" name="primary_outcome_time_frame_{i}" bind:value={primaryOutcome.PrimaryOutcomeTimeFrame} /></div></div>
                            <button class="button is-danger" on:click={() => removePrimaryOutcome(i)}>Remove</button>
                        {/each}
                    </div>
                </div>
            </div>
        </div>
        <hr />
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Secondary Outcome List:</div>
            <div class="column is-three-quarters">
                <div class="columns">
                    <div class="column is-one-quarter">Add Secondary Outcome:</div>
                    <div class="column is-three-quarters">
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_secondary_outcome_measure">Measure:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_secondary_outcome_measure" name="add_secondary_outcome_measure" class="input" bind:value={selected_secondary_outcome.SecondaryOutcomeMeasure} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_secondary_outcome_description">Description:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_secondary_outcome_description" name="add_secondary_outcome_description" class="input" bind:value={selected_secondary_outcome.SecondaryOutcomeDescription} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_secondary_outcome_time_frame">Time Frame:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_secondary_outcome_time_frame" name="add_secondary_outcome_time_frame" class="input" bind:value={selected_secondary_outcome.SecondaryOutcomeTimeFrame} /></div>
                        </div>
                        <button class="button" on:click={() => addSecondaryOutcome()}>Add</button>
                    </div>
                </div>
                <div class="columns">
                    <div class="column is-one-quarter">Secondary Outcomes:</div>
                    <div class="column is-three-quarters">
                        {#each data.ProtocolSection.OutcomesModule.SecondaryOutcomeList.SecondaryOutcome || [] as secondaryOutcome, i }
                            <div class="columns"><div class="column is-one-quarter"><label for="secondary_outcome_measure_{i}">Measure: </label></div><div class="column is-three-quarters"><input type="text" id="secondary_outcome_measure_{i}" name="secondary_outcome_measure_{i}" bind:value={secondaryOutcome.SecondaryOutcomeMeasure} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="secondary_outcome_description_{i}">Description: </label></div><div class="column is-three-quarters"><input type="text" id="secondary_outcome_description_{i}" name="secondary_outcome_description_{i}" bind:value={secondaryOutcome.SecondaryOutcomeDescription} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="secondary_outcome_time_frame_{i}">Time Frame: </label></div><div class="column is-three-quarters"><input type="text" id="secondary_outcome_time_frame_{i}" name="secondary_outcome_time_frame_{i}" bind:value={secondaryOutcome.SecondaryOutcomeTimeFrame} /></div></div>
                            <button class="button is-danger" on:click={() => removeSecondaryOutcome(i)}>Remove</button>
                        {/each}
                    </div>
                </div>
            </div>
        </div>
        <hr />
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Other Outcome List:</div>
            <div class="column is-three-quarters">
                <div class="columns">
                    <div class="column is-one-quarter">Add Other Outcome:</div>
                    <div class="column is-three-quarters">
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_other_outcome_measure">Measure:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_other_outcome_measure" name="add_other_outcome_measure" class="input" bind:value={selected_other_outcome.OtherOutcomeMeasure} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_other_outcome_description">Description:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_other_outcome_description" name="add_other_outcome_description" class="input" bind:value={selected_other_outcome.OtherOutcomeDescription} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_other_outcome_time_frame">Time Frame:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_other_outcome_time_frame" name="add_other_outcome_time_frame" class="input" bind:value={selected_other_outcome.OtherOutcomeTimeFrame} /></div>
                        </div>
                        <button class="button" on:click={() => addOtherOutcome()}>Add</button>
                    </div>
                </div>
                <div class="columns">
                    <div class="column is-one-quarter">Other Outcomes:</div>
                    <div class="column is-three-quarters">
                        {#each data.ProtocolSection.OutcomesModule.OtherOutcomeList.OtherOutcome || [] as otherOutcome, i }
                            <div class="columns"><div class="column is-one-quarter"><label for="other_outcome_measure_{i}">Measure: </label></div><div class="column is-three-quarters"><input type="text" id="other_outcome_measure_{i}" name="other_outcome_measure_{i}" bind:value={otherOutcome.OtherOutcomeMeasure} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="other_outcome_description_{i}">Description: </label></div><div class="column is-three-quarters"><input type="text" id="other_outcome_description_{i}" name="other_outcome_description_{i}" bind:value={otherOutcome.OtherOutcomeDescription} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="other_outcome_time_frame_{i}">Time Frame: </label></div><div class="column is-three-quarters"><input type="text" id="other_outcome_time_frame_{i}" name="other_outcome_time_frame_{i}" bind:value={otherOutcome.OtherOutcomeTimeFrame} /></div></div>
                            <button class="button is-danger" on:click={() => removeOtherOutcome(i)}>Remove</button>
                        {/each}
                    </div>
                </div>
            </div>
        </div>
    </fieldset>
    <fieldset>
        <legend><h3 class="subtitle">Eligibility Module</h3></legend>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="eligibility_criteria">Eligibility Criteria:</label></div>
            <div class="column is-three-quarters"><input id="eligibility_criteria" bind:value={data.ProtocolSection.EligibilityModule.EligibilityCriteria} class="input" type="text" placeholder="Eligibility Criteria"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="healthy_volunteers">Healthy Volunteers:</label></div>
            <div class="column is-three-quarters"><input id="healthy_volunteers" bind:value={data.ProtocolSection.EligibilityModule.HealthyVolunteers} class="input" type="text" placeholder="Healthy Volunteers"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="gender">Gender:</label></div>
            <div class="column is-three-quarters"><input id="gender" bind:value={data.ProtocolSection.EligibilityModule.Gender} class="input" type="text" placeholder="Gender"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="minimum_age">Minimum Age:</label></div>
            <div class="column is-three-quarters"><input id="minimum_age" bind:value={data.ProtocolSection.EligibilityModule.MinimumAge} class="input" type="text" placeholder="Minimum Age"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="maximum_age">Maximum Age:</label></div>
            <div class="column is-three-quarters"><input id="maximum_age" bind:value={data.ProtocolSection.EligibilityModule.MaximumAge} class="input" type="text" placeholder="Maximum Age"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Standard Age List:</div>
            <div class="column is-three-quarters">
                <div class="columns"><div class="column is-one-quarter"><label for="add_std_age">Add Age:</label></div><div class="column is-three-quarters"><select type="text" id="add_std_age" name="add_std_age" class="input" bind:value={selected_std_age_add}><option selected="selected" disabled="disabled">Not Selected</option>{#each std_age_list || [] as age}<option value="{age}">{age}</option>{/each}</select><button class="button" on:click={() => addStdAge()}>Add</button></div></div>
                <div class="columns"><div class="column is-one-quarter"><label for="remove_std_age">Remove Age:</label></div><div class="column is-three-quarters"><select id="remove_std_age" name="remove_std_age" class="select input" bind:value={selected_std_age_remove}><option selected="selected" disabled="disabled">Not Selected</option>{#each data.ProtocolSection.EligibilityModule.StdAgeList.StdAge || [] as age}<option value="{age}">{age}</option>{/each}</select><button class="button" on:click={() => removeStdAge()}>Remove</button></div></div>
            </div>
        </div>
    </fieldset>
    <fieldset>
        <legend><h3 class="subtitle">Contacts Locations Module</h3></legend>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Overall Official List:</div>
            <div class="column is-three-quarters">
                <div class="columns">
                    <div class="column is-one-quarter">Add Overall Official:</div>
                    <div class="column is-three-quarters">
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_overall_official_name">Name:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_overall_official_name" name="add_overall_official_name" class="input" bind:value={selected_overall_official.OverallOfficialName} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_overall_official_affiliation">Affiliation:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_overall_official_affiliation" name="add_overall_official_affiliation" class="input" bind:value={selected_overall_official.OverallOfficialAffiliation} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_overall_official_role">Role:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_overall_official_role" name="add_overall_official_role" class="input" bind:value={selected_overall_official.OverallOfficialRole} /></div>
                        </div>
                        <button class="button" on:click={() => addOverallOfficial()}>Add</button>
                    </div>
                </div>
                <div class="columns">
                    <div class="column is-one-quarter">Overall Officials:</div>
                    <div class="column is-three-quarters">
                        {#each data.ProtocolSection.ContactsLocationsModule.OverallOfficialList.OverallOfficial || [] as overallOfficial, i }
                            <div class="columns"><div class="column is-one-quarter"><label for="overall_official_name_{i}">Measure: </label></div><div class="column is-three-quarters"><input type="text" id="overall_official_name_{i}" name="overall_official_name_{i}" bind:value={overallOfficial.OverallOfficialName} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="overall_official_affiliation_{i}">Description: </label></div><div class="column is-three-quarters"><input type="text" id="overall_official_affiliation_{i}" name="overall_official_affiliation_{i}" bind:value={overallOfficial.OverallOfficialAffiliation} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="overall_official_role_{i}">Time Frame: </label></div><div class="column is-three-quarters"><input type="text" id="overall_official_role_{i}" name="overall_official_role_{i}" bind:value={overallOfficial.OverallOfficialRole} /></div></div>
                            <button class="button is-danger" on:click={() => removeOverallOfficial(i)}>Remove</button>
                        {/each}
                    </div>
                </div>
            </div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Location List:</div>
            <div class="column is-three-quarters">
                <div class="columns">
                    <div class="column is-one-quarter">Add Location:</div>
                    <div class="column is-three-quarters">
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_location_facility">Facility:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_location_facility" name="add_location_facility" class="input" bind:value={selected_location.LocationFacility} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_location_city">City:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_location_city" name="add_location_city" class="input" bind:value={selected_location.LocationCity} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_location_country">Country:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_location_country" name="add_location_country" class="input" bind:value={selected_location.LocationCountry} /></div>
                        </div>
                        <button class="button" on:click={() => addLocation()}>Add</button>
                    </div>
                </div>
                <div class="columns">
                    <div class="column is-one-quarter">Overall Officials:</div>
                    <div class="column is-three-quarters">
                        {#each data.ProtocolSection.ContactsLocationsModule.LocationList.Location || [] as location, i }
                            <div class="columns"><div class="column is-one-quarter"><label for="location_facility_{i}">Facility: </label></div><div class="column is-three-quarters"><input type="text" id="location_facility_{i}" name="location_facility_{i}" bind:value={location.LocationFacility} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="location_city_{i}">City: </label></div><div class="column is-three-quarters"><input type="text" id="location_city_{i}" name="location_city_{i}" bind:value={location.LocationCity} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="location_country_{i}">Country: </label></div><div class="column is-three-quarters"><input type="text" id="location_country_{i}" name="location_country_{i}" bind:value={location.LocationCountry} /></div></div>
                            <button class="button is-danger" on:click={() => removeLocation(i)}>Remove</button>
                        {/each}
                    </div>
                </div>
            </div>
        </div>
    </fieldset>
    <fieldset>
        <legend><h3 class="subtitle">References Module</h3></legend>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Reference List:</div>
            <div class="column is-three-quarters">
                <div class="columns">
                    <div class="column is-one-quarter">Add Reference:</div>
                    <div class="column is-three-quarters">
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_reference_pmid">Pub Med ID:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_reference_pmid" name="add_reference_pmid" class="input" bind:value={selected_reference.ReferencePMID} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_reference_type">Type:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_reference_type" name="add_reference_type" class="input" bind:value={selected_reference.ReferenceType} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_reference_citation">Citation:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_reference_citation" name="add_reference_citation" class="input" bind:value={selected_reference.ReferenceCitation} /></div>
                        </div>
                        <button class="button" on:click={() => addReference()}>Add</button>
                    </div>
                </div>
                <div class="columns">
                    <div class="column is-one-quarter">References:</div>
                    <div class="column is-three-quarters">
                        {#each data.ProtocolSection.ReferencesModule.ReferenceList.Reference || [] as reference, i }
                            <div class="columns"><div class="column is-one-quarter"><label for="reference_pmid_{i}">Facility: </label></div><div class="column is-three-quarters"><input type="text" id="reference_pmid_{i}" name="reference_pmid_{i}" bind:value={reference.ReferencePMID} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="reference_type_{i}">City: </label></div><div class="column is-three-quarters"><input type="text" id="reference_type_{i}" name="reference_type_{i}" bind:value={reference.ReferenceType} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="reference_citation_{i}">Country: </label></div><div class="column is-three-quarters"><input type="text" id="reference_citation_{i}" name="reference_citation_{i}" bind:value={reference.ReferenceCitation} /></div></div>
                            <button class="button is-danger" on:click={() => removeReference(i)}>Remove</button>
                        {/each}
                    </div>
                </div>
            </div>
        </div>
    </fieldset>
    <hr />

    <h2 class="title">Results Section</h2>
    <fieldset>
        <legend><h3 class="subtitle">Baseline Characteristics Module</h3></legend>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="baseline_population_description">Baseline Population Description</label></div>
            <div class="column is-three-quarters"><input id="baseline_population_description" bind:value={data.ResultsSection.BaselineCharacteristicsModule.BaselinePopulationDescription} class="input" type="text" placeholder="Baseline Population Description"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter"><label for="baseline_type_units_analyzed">Baseline Type Units Analyzed</label></div>
            <div class="column is-three-quarters"><input id="baseline_type_units_analyzed" bind:value={data.ResultsSection.BaselineCharacteristicsModule.BaselineTypeUnitsAnalyzed} class="input" type="text" placeholder="Baseline Type Units Analyzed"></div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Baseline Group List:</div>
            <div class="column is-three-quarters">
                <div class="columns">
                    <div class="column is-one-quarter">Add Baseline Group:</div>
                    <div class="column is-three-quarters">
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_baseline_group_id">Group ID:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_baseline_group_id" name="add_baseline_group_id" class="input" bind:value={selected_baseline_group.BaselineGroupID} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_baseline_group_title">Title:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_baseline_group_title" name="add_baseline_group_title" class="input" bind:value={selected_baseline_group.BaselineGroupTitle} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_baseline_group_description">Description:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_baseline_group_description" name="add_baseline_group_description" class="input" bind:value={selected_baseline_group.BaselineGroupDescription} /></div>
                        </div>
                        <button class="button" on:click={() => addBaselineGroup()}>Add</button>
                    </div>
                </div>
                <div class="columns">
                    <div class="column is-one-quarter">Baseline Groups:</div>
                    <div class="column is-three-quarters">
                        {#each data.ResultsSection.BaselineCharacteristicsModule.BaselineGroupList.BaselineGroup || [] as baselineGroup, i }
                            <div class="columns"><div class="column is-one-quarter"><label for="baseline_group_id_{i}">Group ID: </label></div><div class="column is-three-quarters"><input type="text" id="baseline_group_id_{i}" name="baseline_group_id_{i}" bind:value={baselineGroup.BaselineGroupID} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="baseline_group_name_{i}">Name: </label></div><div class="column is-three-quarters"><input type="text" id="baseline_group_name_{i}" name="baseline_group_name_{i}" bind:value={baselineGroup.BaselineGroupTitle} /></div></div>
                            <div class="columns"><div class="column is-one-quarter"><label for="baseline_group_description_{i}">Description: </label></div><div class="column is-three-quarters"><input type="text" id="baseline_group_description_{i}" name="baseline_group_description_{i}" bind:value={baselineGroup.BaselineGroupDescription} /></div></div>
                            <button class="button is-danger" on:click={() => removeBaselineGroup(i)}>Remove</button>
                        {/each}
                    </div>
                </div>
            </div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Baseline Denom List:</div>
            <div class="column is-three-quarters">
                <div class="columns">
                    <div class="column is-one-quarter">Add Baseline Denom:</div>
                    <div class="column is-three-quarters">
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_baseline_denom_units">Units:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_baseline_denom_units" name="add_baseline_denom_units" class="input" bind:value={selected_baseline_denom.BaselineDenomUnits} /></div>
                        </div>
                        <button class="button" on:click={() => addBaselineDenom()}>Add</button>
                    </div>
                </div>
                <div class="columns">
                    <div class="column is-one-quarter">Baseline Groups:</div>
                    <div class="column is-three-quarters">
                        {#each data.ResultsSection.BaselineCharacteristicsModule.BaselineDenomList.BaselineDenom || [] as baselineDenom, i }
                            <div class="columns"><div class="column is-one-quarter"><label for="baseline_denom_units_{i}">Units: </label></div><div class="column is-three-quarters"><input type="text" id="baseline_denom_units_{i}" name="baseline_denom_units_{i}" bind:value={baselineDenom.BaselineDenomUnits} /></div></div>
                            <div class="columns">
                                <div class="column is-one-quarter">Baseline Denom Count List:</div>
                                <div class="column is-three-quarters">
                                    <div class="column is-one-quarter">Add Baseline Denom Count:</div>
                                    <div class="column is-three-quarters">
                                        <div class="columns"><div class="column is-one-quarter"><label for="add_baseline_denom_count_group_id">Group ID:</label></div><div class="column is-three-quarters"><input type="text" id="add_baseline_denom_count_group_id" name="add_baseline_denom_count_group_id" class="input" bind:value={selected_baseline_denom_count.BaselineDenomCountGroupID} /></div></div>
                                        <div class="columns"><div class="column is-one-quarter"><label for="add_baseline_denom_count_value">Value:</label></div><div class="column is-three-quarters"><input type="text" id="add_baseline_denom_count_value" name="add_baseline_denom_count_value" class="input" bind:value={selected_baseline_denom_count.BaselineDenomCountValue} /></div></div>
                                        <button class="button" on:click={() => addBaselineDenomCount(i)}>Add</button>
                                    </div>
                                    <div class="columns">
                                        <div class="column is-one-quarter"><label for="remove_design_who_masked">Group Label:</label></div>
                                        <div class="column is-three-quarters">
                                            <ul>{#each baselineDenom.BaselineDenomCountList.BaselineDenomCount || [] as baselineDenomCount, j}
                                                <li>{baselineDenomCount} - <button class="button is-danger" on:click={() => removeBaselineDenomCount(i, j)}>-</button></li>
                                            {/each}</ul>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <button class="button is-danger" on:click={() => removeBaselineDenom(i)}>Remove</button>
                        {/each}
                    </div>
                </div>
            </div>
        </div>
        <div class="columns is-centered is-mobile" >
            <div class="column is-one-quarter">Baseline Measure List:</div>
            <div class="column is-three-quarters">
                <div class="columns">
                    <div class="column is-one-quarter">Add Baseline Measure:</div>
                    <div class="column is-three-quarters">
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_baseline_measure_title">Title:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_title" name="add_baseline_measure_title" class="input" bind:value={selected_baseline_measure.BaselineMeasureTitle} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_baseline_measure_description">Description:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_description" name="add_baseline_measure_description" class="input" bind:value={selected_baseline_measure.BaselineMeasureDescription} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_baseline_measure_param_type">Param Type:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_param_type" name="add_baseline_measure_param_type" class="input" bind:value={selected_baseline_measure.BaselineMeasureParamType} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_baseline_measure_dispersion_type">Dispersion Type:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_dispersion_type" name="add_baseline_measure_dispersion_type" class="input" bind:value={selected_baseline_measure.BaselineMeasureDispersionType} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_baseline_measure_unit_of_measure">Unit Of Measure:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_unit_of_measure" name="add_baseline_measure_unit_of_measure" class="input" bind:value={selected_baseline_measure.BaselineMeasureUnitOfMeasure} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_baseline_measure_calculate_pct">Calculate PCT:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_calculate_pct" name="add_baseline_measure_calculate_pct" class="input" bind:value={selected_baseline_measure.BaselineMeasureCalculatePct} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_baseline_measure_population_description">Population Description:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_population_description" name="add_baseline_measure_population_description" class="input" bind:value={selected_baseline_measure.BaselineMeasurePopulationDescription} /></div>
                        </div>
                        <div class="columns">
                            <div class="column is-one-quarter"><label for="add_baseline_measure_denom_units_selected">Denom Units Selected:</label></div>
                            <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_denom_units_selected" name="add_baseline_measure_denom_units_selected" class="input" bind:value={selected_baseline_measure.BaselineMeasureDenomUnitsSelected} /></div>
                        </div>
                        <button class="button" on:click={() => addBaselineMeasure()}>Add</button>
                    </div>
                </div>
                <div class="columns">
                    <div class="column is-one-quarter">Baseline Measures:</div>
                    <div class="column is-three-quarters">
                        {#each data.ResultsSection.BaselineCharacteristicsModule.BaselineMeasureList.BaselineMeasure || [] as baselineMeasure, i }
                            <div class="columns">
                                <div class="column is-one-quarter"><label for="add_baseline_measure_title_{i}">Title:</label></div>
                                <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_title_{i}" name="add_baseline_measure_title_{i}" class="input" bind:value={baselineMeasure.BaselineMeasureTitle} /></div>
                            </div>
                            <div class="columns">
                                <div class="column is-one-quarter"><label for="add_baseline_measure_description_{i}">Description:</label></div>
                                <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_description_{i}" name="add_baseline_measure_description_{i}" class="input" bind:value={baselineMeasure.BaselineMeasureDescription} /></div>
                            </div>
                            <div class="columns">
                                <div class="column is-one-quarter"><label for="add_baseline_measure_param_type_{i}">Param Type:</label></div>
                                <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_param_type_{i}" name="add_baseline_measure_param_type_{i}" class="input" bind:value={baselineMeasure.BaselineMeasureParamType} /></div>
                            </div>
                            <div class="columns">
                                <div class="column is-one-quarter"><label for="add_baseline_measure_dispersion_type_{i}">Dispersion Type:</label></div>
                                <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_dispersion_type_{i}" name="add_baseline_measure_dispersion_type_{i}" class="input" bind:value={baselineMeasure.BaselineMeasureDispersionType} /></div>
                            </div>
                            <div class="columns">
                                <div class="column is-one-quarter"><label for="add_baseline_measure_unit_of_measure_{i}">Unit Of Measure:</label></div>
                                <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_unit_of_measure_{i}" name="add_baseline_measure_unit_of_measure_{i}" class="input" bind:value={baselineMeasure.BaselineMeasureUnitOfMeasure} /></div>
                            </div>
                            <div class="columns">
                                <div class="column is-one-quarter"><label for="add_baseline_measure_calculate_pct_{i}">Calculate PCT:</label></div>
                                <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_calculate_pct_{i}" name="add_baseline_measure_calculate_pct_{i}" class="input" bind:value={baselineMeasure.BaselineMeasureCalculatePct} /></div>
                            </div>
                            <div class="columns">
                                <div class="column is-one-quarter"><label for="add_baseline_measure_population_description_{i}">Population Description:</label></div>
                                <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_population_description_{i}" name="add_baseline_measure_population_description_{i}" class="input" bind:value={baselineMeasure.BaselineMeasurePopulationDescription} /></div>
                            </div>
                            <div class="columns">
                                <div class="column is-one-quarter"><label for="add_baseline_measure_denom_units_selected_{i}">Denom Units Selected:</label></div>
                                <div class="column is-three-quarters"><input type="text" id="add_baseline_measure_denom_units_selected_{i}" name="add_baseline_measure_denom_units_selected_{i}" class="input" bind:value={baselineMeasure.BaselineMeasureDenomUnitsSelected} /></div>
                            </div>

                            <div class="columns">
                                <div class="column is-one-quarter">Baseline Class List:</div>
                                <div class="column is-three-quarters">
                                    <div class="column is-one-quarter">Add Baseline Class:</div>
                                    <div class="column is-three-quarters">
                                        <div class="columns"><div class="column is-one-quarter"><label for="add_baseline_class_title">Title:</label></div><div class="column is-three-quarters"><input type="text" id="add_baseline_class_title" name="add_baseline_class_title" class="input" bind:value={selected_baseline_class.BaselineClassTitle} /></div></div>
                                        <button class="button" on:click={() => addBaselineClass(i)}>Add</button>
                                    </div>
                                    <div class="columns">
                                        <div class="column is-one-quarter">Baseline Class List:</div>
                                        <div class="column is-three-quarters">
                                            {#each baselineMeasure.BaselineClassList.BaselineClass || [] as baselineClass, j}
                                                <fieldset>
                                                    <legend>Class Title: <input type="text" bind:value={baselineClass.BaselineClassTitle} /> - <button class="button is-danger" on:click={() => removeBaselineClass(i, j)}>-</button></legend>
                                                    <h3>Add New Baseline Category</h3>
                                                    <div class="columns"><div class="column is-one-quarter"><label for="add_baseline_category_title">Category Title:</label></div><div class="column is-three-quarters"><input type="text" id="add_baseline_category_title" name="add_baseline_category_title" class="input" bind:value={selected_baseline_category.BaselineCategoryTitle} /></div></div>
                                                    <button class="button" on:click={() => addBaselineCategory(i, j)}>Add</button>
                                                    <h3>Baseline Category List</h3>
                                                    {#each baselineClass.BaselineCategoryList.BaselineCategory || [] as baselineCategory, k}
                                                        <fieldset>
                                                            <legend>Category Title: <input type="text" bind:value={baselineCategory.BaselineCategoryTitle} /> - <button class="button is-danger" on:click={() => removeBaselineCategory(i, j, k)}>-</button></legend>
                                                            <h3>Add New Baseline Measurement</h3>
                                                            <div class="columns"><div class="column is-one-quarter"><label for="add_baseline_measurement_group_id">Group ID:</label></div><div class="column is-three-quarters"><input type="text" id="add_baseline_measurement_group_id" name="add_baseline_measurement_group_id" class="input" bind:value={selected_baseline_measurement.BaselineMeasurementGroupID} /></div></div>
                                                            <div class="columns"><div class="column is-one-quarter"><label for="add_baseline_measurement_value">Value:</label></div><div class="column is-three-quarters"><input type="text" id="add_baseline_measurement_value" name="add_baseline_measurement_value" class="input" bind:value={selected_baseline_measurement.BaselineMeasurementValue} /></div></div>
                                                            <div class="columns"><div class="column is-one-quarter"><label for="add_baseline_measurement_spread">Spread:</label></div><div class="column is-three-quarters"><input type="text" id="add_baseline_measurement_spread" name="add_baseline_measurement_spread" class="input" bind:value={selected_baseline_measurement.BaselineMeasurementSpread} /></div></div>
                                                            <div class="columns"><div class="column is-one-quarter"><label for="add_baseline_measurement_comment">Comment:</label></div><div class="column is-three-quarters"><input type="text" id="add_baseline_measurement_comment" name="add_baseline_measurement_comment" class="input" bind:value={selected_baseline_measurement.BaselineMeasurementComment} /></div></div>
                                                            <div class="columns"><div class="column is-one-quarter"><label for="add_baseline_measurement_upper_limit">Upper Limit:</label></div><div class="column is-three-quarters"><input type="text" id="add_baseline_measurement_upper_limit" name="add_baseline_measurement_upper_limit" class="input" bind:value={selected_baseline_measurement.BaselineMeasurementUpperLimit} /></div></div>
                                                            <div class="columns"><div class="column is-one-quarter"><label for="add_baseline_measurement_lower_limit">Lower Limit:</label></div><div class="column is-three-quarters"><input type="text" id="add_baseline_measurement_lower_limit" name="add_baseline_measurement_lower_limit" class="input" bind:value={selected_baseline_measurement.BaselineMeasurementLowerLimit} /></div></div>
                                                            <button class="button" on:click={() => addBaselineMeasurement(i, j, k)}>Add</button>
                                                            <h3>Baseline Measurements</h3>
                                                            <table>
                                                                <tr>
                                                                    <th>Group ID</th>
                                                                    <th>Value</th>
                                                                    <th>Spread</th>
                                                                    <th>Comment</th>
                                                                    <th>Upper Limit</th>
                                                                    <th>Lower Limit</th>
                                                                    <th>Action</th>
                                                                </tr>
                                                            {#each baselineCategory.BaselineMeasurementList.BaselineMeasurement || [] as baselineMeasurement, l}
                                                                <tr>
                                                                    <td><input type="text" bind:value={baselineMeasurement.BaselineMeasurementGroupID} /></td>
                                                                    <td><input type="text" bind:value={baselineMeasurement.BaselineMeasurementValue} /></td>
                                                                    <td><input type="text" bind:value={baselineMeasurement.BaselineMeasurementSpread} /></td>
                                                                    <td><input type="text" bind:value={baselineMeasurement.BaselineMeasurementComment} /></td>
                                                                    <td><input type="text" bind:value={baselineMeasurement.BaselineMeasurementUpperLimit} /></td>
                                                                    <td><input type="text" bind:value={baselineMeasurement.BaselineMeasurementLowerLimit} /></td>
                                                                    <td><button on:click={() => removeBaselineMeasurement(i, j, k, l)} class="button is-danger">-</button></td>
                                                                </tr>
                                                            {/each}
                                                            </table>
                                                        </fieldset>
                                                    {/each}
                                                </fieldset>
                                            {/each}
                                        </div>
                                    </div>
                                </div>
                            </div>

                            <hr />

                            <div class="columns">
                                <div class="column is-one-quarter">Baseline Measure Denom List:</div>
                                <div class="column is-three-quarters">
                                    <div class="column is-one-quarter">Add Baseline Measure Denom:</div>
                                    <div class="column is-three-quarters">
                                        <div class="columns"><div class="column is-one-quarter"><label for="add_baseline_measure_denom_units">Units:</label></div><div class="column is-three-quarters"><input type="text" id="add_baseline_measure_denom_units" name="add_baseline_measure_denom_units" class="input" bind:value={selected_baseline_measure_denom.BaselineMeasureDenomUnits} /></div></div>
                                        <button class="button" on:click={() => addBaselineMeasureDenom(i)}>Add</button>
                                    </div>
                                    <div class="columns">
                                        <div class="column is-one-quarter"><label for="remove_design_who_masked">Baseline Measure Denom List:</label></div>
                                        <div class="column is-three-quarters">
                                            <ul>{#each baselineMeasure.BaselineMeasureDenomList.BaselineMeasureDenom || [] as baselineMeasureDenom, j}
                                                <li>
                                                    <input type="text" bind:value={baselineMeasureDenom.BaselineMeasureDenomUnits} /> - <button class="button is-danger" on:click={() => removeBaselineMeasureDenom(i, j)}>-</button>
                                                    <table>
                                                        <tr>
                                                            <th>Group ID</th>
                                                            <th>Value</th>
                                                            <th>Action</th>
                                                        </tr>
                                                        <tr>
                                                            <td><input type="text" bind:value={selected_baseline_measure_denom_count.BaselineMeasureDenomCountGroupID} /></td>
                                                            <td><input type="text" bind:value={selected_baseline_measure_denom_count.BaselineMeasureDenomCountValue} /></td>
                                                            <td><button on:click={addBaselineMeasureDenomCount(i, j)} class="button">Add</button></td>
                                                        </tr>
                                                        {#each baselineMeasureDenom.BaselineMeasureDenomCountList.BaselineMeasureDenomCount || [] as baselineMeasureDenomCount, k}
                                                            <tr>
                                                                <td><input type="text" bind:value={baselineMeasureDenomCount.BaselineMeasureDenomCountGroupID} /></td>
                                                                <td><input type="text" bind:value={baselineMeasureDenomCount.BaselineMeasureDenomCountValue} /></td>
                                                                <td><button on:click={() => removeBaselineMeasureDenomCount(i, j, k)} class="button is-danger">-</button></td>
                                                            </tr>
                                                        {/each}
                                                    </table>
                                                </li>
                                            {/each}</ul>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <button class="button is-danger" on:click={() => removeBaselineMeasure(i)}>Remove</button>
                        {/each}
                    </div>
                </div>
            </div>
        </div>
    </fieldset>

    <fieldset>
        <legend><h3 class="subtitle">Outcome Measures Module</h3></legend>
        <h4>Outcome Measures</h4>
        <table style="overflow-x: scroll;">
            <tr>
                <th>Type</th>
                <th>Title</th>
                <th>Description</th>
                <th>Population Description</th>
                <th>Reporting Status</th>
                <th>Param Type</th>
                <th>Dispersion Type</th>
                <th>Unit Of Measure</th>
                <th>Time Frame</th>
                <th>Anticipated Posting Date</th>
                <th>Calculate Pct</th>
                <th>Denom Units Selected</th>
                <th>Type Units Analyzed</th>
                <th>Action</th>
            </tr>
            <tr>
                <td><input type="text" bind:value={selected_outcome_measure.OutcomeMeasureType} /></td>
                <td><input type="text" bind:value={selected_outcome_measure.OutcomeMeasureTitle} /></td>
                <td><input type="text" bind:value={selected_outcome_measure.OutcomeMeasureDescription} /></td>
                <td><input type="text" bind:value={selected_outcome_measure.OutcomeMeasurePopulationDescription} /></td>
                <td><input type="text" bind:value={selected_outcome_measure.OutcomeMeasureReportingStatus} /></td>
                <td><input type="text" bind:value={selected_outcome_measure.OutcomeMeasureParamType} /></td>
                <td><input type="text" bind:value={selected_outcome_measure.OutcomeMeasureDispersionType} /></td>
                <td><input type="text" bind:value={selected_outcome_measure.OutcomeMeasureUnitOfMeasure} /></td>
                <td><input type="text" bind:value={selected_outcome_measure.OutcomeMeasureTimeFrame} /></td>
                <td><input type="text" bind:value={selected_outcome_measure.OutcomeMeasureAnticipatedPostingDate} /></td>
                <td><input type="text" bind:value={selected_outcome_measure.OutcomeMeasureCalculatePct} /></td>
                <td><input type="text" bind:value={selected_outcome_measure.OutcomeMeasureDenomUnitsSelected} /></td>
                <td><input type="text" bind:value={selected_outcome_measure.OutcomeMeasureTypeUnitsAnalyzed} /></td>
                <td><button on:click={() => add_outcome_measure()} class="button">Add</button></td>
            </tr>
            {#each data.ResultsSection.OutcomeMeasuresModule.OutcomeMeasureList.OutcomeMeasure || [] as outcomeMeasure, i}
                <tr>
                    <td><input type="text" bind:value={outcomeMeasure.OutcomeMeasureType} /></td>
                    <td><input type="text" bind:value={outcomeMeasure.OutcomeMeasureTitle} /></td>
                    <td><input type="text" bind:value={outcomeMeasure.OutcomeMeasureDescription} /></td>
                    <td><input type="text" bind:value={outcomeMeasure.OutcomeMeasurePopulationDescription} /></td>
                    <td><input type="text" bind:value={outcomeMeasure.OutcomeMeasureReportingStatus} /></td>
                    <td><input type="text" bind:value={outcomeMeasure.OutcomeMeasureParamType} /></td>
                    <td><input type="text" bind:value={outcomeMeasure.OutcomeMeasureDispersionType} /></td>
                    <td><input type="text" bind:value={outcomeMeasure.OutcomeMeasureUnitOfMeasure} /></td>
                    <td><input type="text" bind:value={outcomeMeasure.OutcomeMeasureTimeFrame} /></td>
                    <td><input type="text" bind:value={outcomeMeasure.OutcomeMeasureAnticipatedPostingDate} /></td>
                    <td><input type="text" bind:value={outcomeMeasure.OutcomeMeasureCalculatePct} /></td>
                    <td><input type="text" bind:value={outcomeMeasure.OutcomeMeasureDenomUnitsSelected} /></td>
                    <td><input type="text" bind:value={outcomeMeasure.OutcomeMeasureTypeUnitsAnalyzed} /></td>
                    <td><button on:click={() => remove_outcome_measure(i)} class="button">Remove</button></td>
                </tr>
            {/each}
        </table>
    </fieldset>
</main>
