package clinicaltrials

type Study struct {
	ProtocolSection struct {
		IdentificationModule struct {
			NCTID          string `json:"NCTId" html:"l=Study ID,e=span"`
			OrgStudyIDInfo struct {
				OrgStudyID string `json:"OrgStudyId" html:"l=Organisation Study ID,e=span"`
			} `json:"OrgStudyIdInfo"`
			Organization struct {
				OrgFullName string `json:"OrgFullName" html:"l=Full Name,e=span"`
				OrgClass    string `json:"OrgClass" html:"l=Class,e=span"`
			} `json:"Organization" html:"l=Organisation Details"`
			BriefTitle    string `json:"BriefTitle" html:"l=Brief Title,e=span"`
			OfficialTitle string `json:"OfficialTitle" html:"l=Official Title,e=span"`
		} `json:"IdentificationModule" html:"l=Identification Module,c=top-level-module"`
		StatusModule struct {
			StatusVerifiedDate string `json:"StatusVerifiedDate" html:"l=Status Verified Date,e=span"`
			OverallStatus      string `json:"OverallStatus" html:"l=Overall Status,e=span"`
			ExpandedAccessInfo struct {
				HasExpandedAccess string `json:"HasExpandedAccess" html:"l=Has Expanded Access,e=span"`
			} `json:"ExpandedAccessInfo"`
			StartDateStruct struct {
				StartDate     string `json:"StartDate,e=span"`
				StartDateType string `json:"StartDateType,e=span"`
			} `json:"StartDateStruct" html:"l=Start Date,c=date"`
			PrimaryCompletionDateStruct struct {
				PrimaryCompletionDate     string `json:"PrimaryCompletionDate,e=span"`
				PrimaryCompletionDateType string `json:"PrimaryCompletionDateType,e=span"`
			} `json:"PrimaryCompletionDateStruct" html:"l=Primary Completion Date,c=date"`
			CompletionDateStruct struct {
				CompletionDate     string `json:"CompletionDate,e=span"`
				CompletionDateType string `json:"CompletionDateType,e=span"`
			} `json:"CompletionDateStruct" html:"l=Completion Date,c=date"`
			StudyFirstSubmitDate     string `json:"StudyFirstSubmitDate" html:"l=Study First Submit Date,e=span"`
			StudyFirstSubmitQCDate   string `json:"StudyFirstSubmitQCDate" html:"l=Study First Submit QC Date,e=span"`
			StudyFirstPostDateStruct struct {
				StudyFirstPostDate     string `json:"StudyFirstPostDate,e=span"`
				StudyFirstPostDateType string `json:"StudyFirstPostDateType,e=span"`
			} `json:"StudyFirstPostDateStruct" html:"l=Study First Post Date,c=date"`
			ResultsFirstSubmitDate     string `json:"ResultsFirstSubmitDate,e=span"`
			ResultsFirstSubmitQCDate   string `json:"ResultsFirstSubmitQCDate,e=span"`
			ResultsFirstPostDateStruct struct {
				ResultsFirstPostDate     string `json:"ResultsFirstPostDate,e=span"`
				ResultsFirstPostDateType string `json:"ResultsFirstPostDateType,e=span"`
			} `json:"ResultsFirstPostDateStruct" html:"l=Results First Post Date,c=date"`
			LastUpdateSubmitDate     string `json:"LastUpdateSubmitDate,e=span"`
			LastUpdatePostDateStruct struct {
				LastUpdatePostDate     string `json:"LastUpdatePostDate,e=span"`
				LastUpdatePostDateType string `json:"LastUpdatePostDateType,e=span"`
			} `json:"LastUpdatePostDateStruct" html:"l=Results First Post Date,c=date"`
		} `json:"StatusModule" html:"l=Status Module,c=top-level-module"`
		SponsorCollaboratorsModule struct {
			ResponsibleParty struct {
				ResponsiblePartyType string `json:"ResponsiblePartyType" html:"l=Responsible Party Type,e=span"`
			} `json:"ResponsibleParty"`
			LeadSponsor struct {
				LeadSponsorName  string `json:"LeadSponsorName" html:"l=Name,e=span"`
				LeadSponsorClass string `json:"LeadSponsorClass" html:"l=Class,e=span"`
			} `json:"LeadSponsor" html:"l=Lead Sponsor"`
		} `json:"SponsorCollaboratorsModule" html:"l=Sponsor Collaborators Module,c=top-level-module"`
		OversightModule struct {
			OversightHasDMC      string `json:"OversightHasDMC,e=span"`
			IsFDARegulatedDrug   string `json:"IsFDARegulatedDrug,e=span"`
			IsFDARegulatedDevice string `json:"IsFDARegulatedDevice,e=span"`
		} `json:"OversightModule" html:"l=Oversight Module,c=top-level-module"`
		DescriptionModule struct {
			BriefSummary        string `json:"BriefSummary" html:"l=Brief Summary,e=span"`
			DetailedDescription string `json:"DetailedDescription" html:"l=Detailed Description,e=span"`
		} `json:"DescriptionModule" html:"l=Description Module,c=top-level-module"`
		ConditionsModule struct {
			ConditionList struct {
				Condition []string `json:"Condition" html:"e=li"`
			} `json:"ConditionList" html:"l=Condition List,e=ul"`
		} `json:"ConditionsModule" html:"l=Condition Module,c=top-level-module"`
		DesignModule struct {
			StudyType string `json:"StudyType,e=span"`
			PhaseList struct {
				Phase []string `json:"Phase" html:"e=li"`
			} `json:"PhaseList" html:"l=Phases,e=ul"`
			DesignInfo struct {
				DesignAllocation        string `json:"DesignAllocation" html:"l=Allocation,e=span"`
				DesignInterventionModel string `json:"DesignInterventionModel" html:"l=Intervention Model,e=span"`
				DesignPrimaryPurpose    string `json:"DesignPrimaryPurpose" html:"l=Primary Purpose,e=span"`
				DesignMaskingInfo       struct {
					DesignMasking       string `json:"DesignMasking" html:"l=Design Masking,e=span"`
					DesignWhoMaskedList struct {
						DesignWhoMasked []string `json:"DesignWhoMasked" html:"e=li"`
					} `json:"DesignWhoMaskedList" html:"l=Who Masked List,e=ul"`
				} `json:"DesignMaskingInfo" html:"l=Design Masking Information"`
			} `json:"DesignInfo" html:"l=Design Information"`
			EnrollmentInfo struct {
				EnrollmentCount string `json:"EnrollmentCount" html:"l=Count,e=span"`
				EnrollmentType  string `json:"EnrollmentType" html:"l=Type,e=span"`
			} `json:"EnrollmentInfo" html:"l=Enrollment Information"`
		} `json:"DesignModule" html:"l=Design Module,c=top-level-module"`
		ArmsInterventionsModule struct {
			ArmGroupList struct {
				ArmGroup []struct {
					ArmGroupLabel            string `json:"ArmGroupLabel" html:"l=Label,e=span"`
					ArmGroupType             string `json:"ArmGroupType" html:"l=Type,e=span"`
					ArmGroupDescription      string `json:"ArmGroupDescription" html:"l=Description,e=span"`
					ArmGroupInterventionList struct {
						ArmGroupInterventionName []string `json:"ArmGroupInterventionName" html:"e=li"`
					} `json:"ArmGroupInterventionList" html:"l=Intervention List,e=ul"`
				} `json:"ArmGroup" html:"row"`
			} `json:"ArmGroupList" html:"l=Arm Group List"`
			InterventionList struct {
				Intervention []struct {
					InterventionType              string `json:"InterventionType" html:"l=Type,e=span"`
					InterventionName              string `json:"InterventionName" html:"l=Name,e=span"`
					InterventionDescription       string `json:"InterventionDescription" html:"l=Description,e=span"`
					InterventionArmGroupLabelList struct {
						InterventionArmGroupLabel []string `json:"InterventionArmGroupLabel" html:"e=li"`
					} `json:"InterventionArmGroupLabelList" html:"l=Arm Group Label List,e=ul"`
				} `json:"Intervention" html:"e=li"`
			} `json:"InterventionList" html:"l=Intervention List,e=ul"`
		} `json:"ArmsInterventionsModule" html:"l=Arms Interventions Module,c=top-level-module"`
		OutcomesModule struct {
			PrimaryOutcomeList struct {
				PrimaryOutcome []struct {
					PrimaryOutcomeMeasure     string `json:"PrimaryOutcomeMeasure" html:"l=Outcome Measure,e=span"`
					PrimaryOutcomeDescription string `json:"PrimaryOutcomeDescription" html:"l=Outcome Description,e=span"`
					PrimaryOutcomeTimeFrame   string `json:"PrimaryOutcomeTimeFrame" html:"l=Outcome Time Frame,e=span"`
				} `json:"PrimaryOutcome" html:"e=li"`
			} `json:"PrimaryOutcomeList" html:"l=Primary Outcome List,e=ul"`
			SecondaryOutcomeList struct {
				SecondaryOutcome []struct {
					SecondaryOutcomeMeasure     string `json:"SecondaryOutcomeMeasure" html:"l=Outcome Measure,e=span"`
					SecondaryOutcomeDescription string `json:"SecondaryOutcomeDescription" html:"l=Outcome Description,e=span"`
					SecondaryOutcomeTimeFrame   string `json:"SecondaryOutcomeTimeFrame" html:"l=Outcome Time Frame,e=span"`
				} `json:"SecondaryOutcome" html:"e=li"`
			} `json:"SecondaryOutcomeList" html:"l=Secondary Outcome List,e=ul"`
			OtherOutcomeList struct {
				OtherOutcome []struct {
					OtherOutcomeMeasure     string `json:"OtherOutcomeMeasure" html:"l=Outcome Measure,e=span"`
					OtherOutcomeDescription string `json:"OtherOutcomeDescription" html:"l=Outcome Description,e=span"`
					OtherOutcomeTimeFrame   string `json:"OtherOutcomeTimeFrame" html:"l=Outcome Time Frame,e=span"`
				} `json:"OtherOutcome" html:"e=li"`
			} `json:"OtherOutcomeList" html:"l=Other Outcome List,e=ul"`
		} `json:"OutcomesModule" html:"l=Outcomes Module,c=top-level-module"`
		EligibilityModule struct {
			EligibilityCriteria string `json:"EligibilityCriteria,e=span"`
			HealthyVolunteers   string `json:"HealthyVolunteers,e=span"`
			Gender              string `json:"Gender,e=span"`
			MinimumAge          string `json:"MinimumAge,e=span"`
			MaximumAge          string `json:"MaximumAge,e=span"`
			StdAgeList          struct {
				StdAge []string `json:"StdAge" html:"e=li"`
			} `json:"StdAgeList" html:"l=Standard Age List,e=ul"`
		} `json:"EligibilityModule" html:"l=Eligibility Module,c=top-level-module"`
		ContactsLocationsModule struct {
			OverallOfficialList struct {
				OverallOfficial []struct {
					OverallOfficialName        string `json:"OverallOfficialName" html:"l=Name,e=span"`
					OverallOfficialAffiliation string `json:"OverallOfficialAffiliation" html:"l=Affiliation,e=span"`
					OverallOfficialRole        string `json:"OverallOfficialRole" html:"l=Role,e=span"`
				} `json:"OverallOfficial" html:"e=li"`
			} `json:"OverallOfficialList" html:"l=Overall Official List,e=ul"`
			LocationList struct {
				Location []struct {
					LocationFacility string `json:"LocationFacility" html:"l=Facility,e=span"`
					LocationCity     string `json:"LocationCity" html:"l=City,e=span"`
					LocationCountry  string `json:"LocationCountry" html:"l=Country,e=span"`
				} `json:"Location" html:"row"`
			} `json:"LocationList" html:"l=Location List"`
		} `json:"ContactsLocationsModule" html:"l=Contacts Locations Module,c=top-level-module"`
		ReferencesModule struct {
			ReferenceList struct {
				Reference []struct {
					ReferencePMID     string `json:"ReferencePMID" html:"l=PM ID,e=span"`
					ReferenceType     string `json:"ReferenceType" html:"l=Type,e=span"`
					ReferenceCitation string `json:"ReferenceCitation" html:"l=Citation,e=span"`
				} `json:"Reference" html:"e=li"`
			} `json:"ReferenceList" html:"l=Reference List,e=ul"`
		} `json:"ReferencesModule" html:"l=References Module,c=top-level-module"`
	} `json:"ProtocolSection" html:"l=Protocol Section,c=top-level-section"`
	ResultsSection struct {
		ParticipantFlowModule struct {
			FlowPreAssignmentDetails string `json:"FlowPreAssignmentDetails" html:"l=Flow Pre Assignment Details,e=span"`
			FlowRecruitmentDetails   string `json:"FlowRecruitmentDetails" html:"l=Flow Recruitment Details,e=span"`
			FlowTypeUnitsAnalyzed    string `json:"FlowTypeUnitsAnalyzed" html:"l=Flow Type Units Analyzed,e=span"`
			FlowGroupList            struct {
				FlowGroup []struct {
					FlowGroupID          string `json:"FlowGroupId" html:"l=ID,e=span"`
					FlowGroupTitle       string `json:"FlowGroupTitle" html:"l=Title,e=span"`
					FlowGroupDescription string `json:"FlowGroupDescription" html:"l=Description,e=span"`
				} `json:"FlowGroup" html:"row"`
			} `json:"FlowGroupList" html:"l=Flow Group List"`
			FlowPeriodList struct {
				FlowPeriod []struct {
					FlowPeriodTitle   string `json:"FlowPeriodTitle" html:"l=Title,e=span"`
					FlowMilestoneList struct {
						FlowMilestone []struct {
							FlowMilestoneType   string `json:"FlowMilestoneType" html:"l=Type,e=span"`
							FlowAchievementList struct {
								FlowAchievement []struct {
									FlowAchievementGroupID     string `json:"FlowAchievementGroupId" html:"l=Group ID,e=span"`
									FlowAchievementNumSubjects string `json:"FlowAchievementNumSubjects" html:"l=Num Subjects,e=span"`
									FlowAchievementComment     string `json:"FlowAchievementComment" html:"l=Comment,e=span"`
									FlowAchievementNumUnits    string `json:"FlowAchievementNumUnits" html:"l=Num Units,e=span"`
								} `json:"FlowAchievement" html:"row"`
							} `json:"FlowAchievementList" html:"l=Flow Achievement List"`
							FlowMilestoneComment string `json:"FlowMilestoneComment,omitempty" html:"l=Comment,e=span,omitempty"`
						} `json:"FlowMilestone" html:"e=li"`
					} `json:"FlowMilestoneList" html:"l=Flow Milestone List,e=ul"`
					FlowDropWithdrawList struct {
						FlowDropWithdraw []struct {
							FlowDropWithdrawType    string `json:"FlowDropWithdrawType" html:"l=Type,e=span"`
							FlowDropWithdrawComment string `json:"FlowDropWithdrawComment" html:"l=Comment,e=span,omitempty"`
							FlowReasonList          struct {
								FlowReason []struct {
									FlowReasonGroupID     string `json:"FlowReasonGroupId" html:"l=Group ID,e=span"`
									FlowReasonNumSubjects string `json:"FlowReasonNumSubjects" html:"l=Num Subjects,e=span"`
									FlowReasonNumUnits    string `json:"FlowReasonNumUnits" html:"l=Num Units,e=span"`
									FlowReasonComment     string `json:"FlowReasonComment" html:"l=Comment,e=span"`
								} `json:"FlowReason" html:"row"`
							} `json:"FlowReasonList" html:"l=Flow Reason List"`
						} `json:"FlowDropWithdraw" html:"e=li"`
					} `json:"FlowDropWithdrawList" html:"l=Flow Drop Withdraw List,e=ul"`
				} `json:"FlowPeriod" html:"e=li"`
			} `json:"FlowPeriodList" html:"l=Flow Period List,e=ul"`
		} `json:"ParticipantFlowModule" html:"l=Participant Flow Module,c=top-level-module"`
		BaselineCharacteristicsModule struct {
			BaselinePopulationDescription string `json:"BaselinePopulationDescription" html:"l=Population Description,e=span,omitempty"`
			BaselineTypeUnitsAnalyzed     string `json:"BaselineTypeUnitsAnalyzed" html:"l=Type Units Analyzed,e=span"`
			BaselineGroupList             struct {
				BaselineGroup []struct {
					BaselineGroupID          string `json:"BaselineGroupId" html:"l=Group ID,e=span"`
					BaselineGroupTitle       string `json:"BaselineGroupTitle" html:"l=Title,e=span"`
					BaselineGroupDescription string `json:"BaselineGroupDescription" html:"l=Description,e=span"`
				} `json:"BaselineGroup" html:"row"`
			} `json:"BaselineGroupList" html:"l=Baseline Group List"`
			BaselineDenomList struct {
				BaselineDenom []struct {
					BaselineDenomUnits     string `json:"BaselineDenomUnits" html:"l=Denom Units,e=span"`
					BaselineDenomCountList struct {
						BaselineDenomCount []struct {
							BaselineDenomCountGroupID string `json:"BaselineDenomCountGroupId" html:"l=Group ID,e=span"`
							BaselineDenomCountValue   string `json:"BaselineDenomCountValue" html:"l=Value,e=span"`
						} `json:"BaselineDenomCount" html:"row"`
					} `json:"BaselineDenomCountList" html:"l=Baseline Denom Count List"`
				} `json:"BaselineDenom" html:"e=li"`
			} `json:"BaselineDenomList" html:"l=Baseline Denom List,e=ul"`
			BaselineMeasureList struct {
				BaselineMeasure []struct {
					BaselineMeasureTitle                 string `json:"BaselineMeasureTitle" html:"l=Title,e=span"`
					BaselineMeasureDescription           string `json:"BaselineMeasureDescription,omitempty" html:"l=Description,e=span"`
					BaselineMeasureParamType             string `json:"BaselineMeasureParamType" html:"l=Param Type,e=span"`
					BaselineMeasureDispersionType        string `json:"BaselineMeasureDispersionType,omitempty" html:"l=Dispersion Type,e=span"`
					BaselineMeasureUnitOfMeasure         string `json:"BaselineMeasureUnitOfMeasure" html:"l=Unit Of Measure,e=span"`
					BaselineMeasureCalculatePct          string `json:"BaselineMeasureCalculatePct" html:"l=Calculate PCT,e=span,omitempty"`
					BaselineMeasurePopulationDescription string `json:"BaselineMeasurePopulationDescription" html:"l=Population Description,e=span,omitempty"`
					BaselineMeasureDenomUnitsSelected    string `json:"BaselineMeasureDenomUnitsSelected" html:"l=Denom Units Selected,e=span"`
					BaselineClassList                    struct {
						BaselineClass []struct {
							BaselineClassTitle   string `json:"BaselineClassTitle" html:"l=Title,e=span"`
							BaselineCategoryList struct {
								BaselineCategory []struct {
									BaselineCategoryTitle   string `json:"BaselineCategoryTitle" html:"l=Title,e=span"`
									BaselineMeasurementList struct {
										BaselineMeasurement []struct {
											BaselineMeasurementGroupID    string `json:"BaselineMeasurementGroupId" html:"l=Group ID,e=span"`
											BaselineMeasurementValue      string `json:"BaselineMeasurementValue" html:"l=Value,e=span"`
											BaselineMeasurementSpread     string `json:"BaselineMeasurementSpread" html:"l=Spread,e=span"`
											BaselineMeasurementComment    string `json:"BaselineMeasurementComment" html:"l=Comment,e=span"`
											BaselineMeasurementUpperLimit string `json:"BaselineMeasurementUpperLimit" html:"l=Upper Limit,e=span"`
											BaselineMeasurementLowerLimit string `json:"BaselineMeasurementLowerLimit" html:"l=Lower Limit,e=span"`
										} `json:"BaselineMeasurement" html:"row"`
									} `json:"BaselineMeasurementList" html:"l=Baseline Measurement List"`
								} `json:"BaselineCategory" html:"e=li"`
							} `json:"BaselineCategoryList,e=ul"`
						} `json:"BaselineClass" html:"e=li"`
					} `json:"BaselineClassList" html:"l=Baseline Class List,e=ul"`
					BaselineMeasureDenomList struct {
						BaselineMeasureDenom []struct {
							BaselineMeasureDenomUnits     string `json:"BaselineMeasureDenomUnits" html:"l=Denom Units,e=span"`
							BaselineMeasureDenomCountList struct {
								BaselineMeasureDenomCount []struct {
									BaselineMeasureDenomCountGroupID string `json:"BaselineMeasureDenomCountGroupID" html:"l=Group ID,e=span"`
									BaselineMeasureDenomCountValue   string `json:"BaselineMeasureDenomCountValue" html:"l=Value,e=span"`
								} `json:"BaselineMeasureDenomCount" html:"row"`
							} `json:"BaselineMeasureDenomCountList" html:"l=Baseline Measure Denom Count List"`
						} `json:"BaselineMeasureDenom" html:"e=li"`
					} `json:"BaselineMeasureDenomList" html:"l=Baseline Measure Denom List,e=ul"`
				} `json:"BaselineMeasure" html:"e=li"`
			} `json:"BaselineMeasureList" html:"l=Baseline Measure List,e=ul"`
		} `json:"BaselineCharacteristicsModule" html:"l=Baseline Characteristics Module,c=top-level-module"`
		OutcomeMeasuresModule struct {
			OutcomeMeasureList struct {
				OutcomeMeasure []struct {
					OutcomeMeasureType                   string `json:"OutcomeMeasureType" html:"l=Type,e=span"`
					OutcomeMeasureTitle                  string `json:"OutcomeMeasureTitle" html:"l=Title,e=span"`
					OutcomeMeasureDescription            string `json:"OutcomeMeasureDescription" html:"l=Description,e=span"`
					OutcomeMeasurePopulationDescription  string `json:"OutcomeMeasurePopulationDescription" html:"l=Population Description,e=span"`
					OutcomeMeasureReportingStatus        string `json:"OutcomeMeasureReportingStatus" html:"l=Reporting Status,e=span"`
					OutcomeMeasureParamType              string `json:"OutcomeMeasureParamType" html:"l=Param Type,e=span"`
					OutcomeMeasureDispersionType         string `json:"OutcomeMeasureDispersionType,omitempty" html:"l=Dispersion Type,e=span"`
					OutcomeMeasureUnitOfMeasure          string `json:"OutcomeMeasureUnitOfMeasure" html:"l=Unit Of Measure,e=span"`
					OutcomeMeasureTimeFrame              string `json:"OutcomeMeasureTimeFrame" html:"l=Time Frame,e=span"`
					OutcomeMeasureAnticipatedPostingDate string `json:"OutcomeMeasureAnticipatedPostingDate" html:"l=Anticipated Posting Date,e=span"`
					OutcomeMeasureCalculatePct           string `json:"OutcomeMeasureCalculatePct" html:"l=Calculate PCT,e=span,omitempty"`
					OutcomeMeasureDenomUnitsSelected     string `json:"OutcomeMeasureDenomUnitsSelected" html:"l=Denom Units Selected,e=span,omitempty"`
					OutcomeMeasureTypeUnitsAnalyzed      string `json:"OutcomeMeasureTypeUnitsAnalyzed" html:"l=Type Units Analyzed,e=span,omitempty"`
					OutcomeGroupList                     struct {
						OutcomeGroup []struct {
							OutcomeGroupID          string `json:"OutcomeGroupId" html:"l=Group ID,e=span"`
							OutcomeGroupTitle       string `json:"OutcomeGroupTitle" html:"l=Group Title,e=span"`
							OutcomeGroupDescription string `json:"OutcomeGroupDescription" html:"l=Description,e=span"`
						} `json:"OutcomeGroup" html:"row"`
					} `json:"OutcomeGroupList" html:"l=Outcome Group List"`
					OutcomeDenomList struct {
						OutcomeDenom []struct {
							OutcomeDenomUnits     string `json:"OutcomeDenomUnits" html:"l=Denom Units,e=span"`
							OutcomeDenomCountList struct {
								OutcomeDenomCount []struct {
									OutcomeDenomCountGroupID string `json:"OutcomeDenomCountGroupId" html:"l=Group ID,e=span"`
									OutcomeDenomCountValue   string `json:"OutcomeDenomCountValue" html:"l=Value,e=span"`
								} `json:"OutcomeDenomCount" html:"row"`
							} `json:"OutcomeDenomCountList" html:"l=Outcome Denom Count List"`
						} `json:"OutcomeDenom" html:"e=li"`
					} `json:"OutcomeDenomList" html:"l=Outcome Denom List,e=ul"`
					OutcomeClassList struct {
						OutcomeClass []struct {
							OutcomeClassTitle     string `json:"OutcomeClassTitle" html:"l=Title,e=span"`
							OutcomeClassDenomList struct {
								OutcomeClassDenom []struct {
									OutcomeClassDenomUnits     string `json:"OutcomeClassDenomUnits" html:"l=Denom Units,e=span"`
									OutcomeClassDenomCountList struct {
										OutcomeClassDenomCount []struct {
											OutcomeClassDenomCountGroupID string `json:"OutcomeClassDenomCountGroupId" html:"l=Group ID,e=span"`
											OutcomeClassDenomCountValue   string `json:"OutcomeClassDenomCountValue" html:"l=Value,e=span"`
										} `json:"OutcomeClassDenomCount" html:"row"`
									} `json:"OutcomeClassDenomCountList" html:"l=Outcome Class Denom Count List"`
								} `json:"OutcomeClassDenom" html:"e=li"`
							} `json:"OutcomeClassDenomList" html:"l=Outcome Class Denom List,e=ul"`
							OutcomeCategoryList struct {
								OutcomeCategory []struct {
									OutcomeCategoryTitle   string `json:"OutcomeCategoryTitle" html:"l=Title,e=span"`
									OutcomeMeasurementList struct {
										OutcomeMeasurement []struct {
											OutcomeMeasurementGroupID    string `json:"OutcomeMeasurementGroupId" html:"l=Group ID,e=span"`
											OutcomeMeasurementValue      string `json:"OutcomeMeasurementValue" html:"l=Value,e=span"`
											OutcomeMeasurementLowerLimit string `json:"OutcomeMeasurementLowerLimit" html:"l=Lower Limit,e=span"`
											OutcomeMeasurementUpperLimit string `json:"OutcomeMeasurementUpperLimit" html:"l=Upper Limit,e=span"`
											OutcomeMeasurementSpread     string `json:"OutcomeMeasurementSpread" html:"l=Spread,e=span"`
											OutcomeMeasurementComment    string `json:"OutcomeMeasurementComment" html:"l=Comment,e=span"`
										} `json:"OutcomeMeasurement" html:"row"`
									} `json:"OutcomeMeasurementList" html:"l=Outcome Measurement List"`
								} `json:"OutcomeCategory" html:"e=li"`
							} `json:"OutcomeCategoryList" html:"l=Outcome Category List,e=ul"`
						} `json:"OutcomeClass" html:"e=li"`
					} `json:"OutcomeClassList" html:"l=Outcome Class List,e=ul"`
					OutcomeAnalysisList struct {
						OutcomeAnalysis []struct {
							OutcomeAnalysisGroupIDList struct {
								OutcomeAnalysisGroupID []string `json:"OutcomeAnalysisGroupId" html:"e=li"`
							} `json:"OutcomeAnalysisGroupIdList" html:"l=Outcome Analysis Group ID List,e=ul"`
							OutcomeAnalysisGroupDescription      string `json:"OutcomeAnalysisGroupDescription" html:"l=Group Description,e=span,omitempty"`
							OutcomeAnalysisNonInferiorityType    string `json:"OutcomeAnalysisNonInferiorityType" html:"l=Non Inferiority Type,e=span"`
							OutcomeAnalysisNonInferiorityComment string `json:"OutcomeAnalysisNonInferiorityComment" html:"l=Non Inferiority Comment,e=span,omitempty"`
							OutcomeAnalysisParamType             string `json:"OutcomeAnalysisParamType" html:"l=Param Type,e=span"`
							OutcomeAnalysisParamValue            string `json:"OutcomeAnalysisParamValue" html:"l=Param Value,e=span"`
							OutcomeAnalysisCIPctValue            string `json:"OutcomeAnalysisCIPctValue" html:"l=CI PCT Value,e=span"`
							OutcomeAnalysisCINumSides            string `json:"OutcomeAnalysisCINumSides" html:"l=CI Num Sides,e=span"`
							OutcomeAnalysisCILowerLimit          string `json:"OutcomeAnalysisCILowerLimit" html:"l=CI Lower Limit,e=span"`
							OutcomeAnalysisCILowerLimitComment   string `json:"OutcomeAnalysisCILowerLimitComment" html:"l=CI Lower Limit Comment,e=span,omitempty"`
							OutcomeAnalysisCIUpperLimit          string `json:"OutcomeAnalysisCIUpperLimit" html:"l=CI Upper Limit,e=span"`
							OutcomeAnalysisCIUpperLimitComment   string `json:"OutcomeAnalysisCIUpperLimitComment" html:"l=CI Upper Limit Comment,e=span,omitempty"`
							OutcomeAnalysisDispersionType        string `json:"OutcomeAnalysisDispersionType" html:"l=Dispersion Type,e=span"`
							OutcomeAnalysisDispersionValue       string `json:"OutcomeAnalysisDispersionValue" html:"l=Dispersion Value,e=span"`
							OutcomeAnalysisEstimateComment       string `json:"OutcomeAnalysisEstimateComment" html:"l=Estimate Comment,e=span,omitempty"`
							OutcomeAnalysisDescription           string `json:"OutcomeAnalysisDescription" html:"l=Description,e=span"`
							OutcomeAnalysisPValue                string `json:"OutcomeAnalysisPValue" html:"l=P Value,e=span"`
							OutcomeAnalysisPValueComment         string `json:"OutcomeAnalysisPValueComment" html:"l=P Value Comment,e=span,omitempty"`
							OutcomeAnalysisStatisticalComment    string `json:"OutcomeAnalysisStatisticalComment" html:"l=Statistical Comment,e=span,omitempty"`
							OutcomeAnalysisStatisticalMethod     string `json:"OutcomeAnalysisStatisticalMethod" html:"l=Statistical Method,e=span"`
							OutcomeAnalysisTestedNonInferiority  string `json:"OutcomeAnalysisTestedNonInferiority" html:"l=Tested Non Inferiority,e=span"`
						} `json:"OutcomeAnalysis" html:"e=li"`
					} `json:"OutcomeAnalysisList,omitempty" html:"l=Outcome Analysis List,e=ul"`
				} `json:"OutcomeMeasure" html:"e=li"`
			} `json:"OutcomeMeasureList" html:"l=Outcome Measure List,e=ul"`
		} `json:"OutcomeMeasuresModule" html:"l=Outcome Measures Module,c=top-level-module"`
		AdverseEventsModule struct {
			EventsFrequencyThreshold string `json:"EventsFrequencyThreshold" html:"l=Frequency Threshold,e=span"`
			EventsTimeFrame          string `json:"EventsTimeFrame" html:"l=Time Frame,e=span"`
			EventsDescription        string `json:"EventsDescription" html:"l=Description,e=span"`
			EventGroupList           struct {
				EventGroup []struct {
					EventGroupID                 string `json:"EventGroupId" html:"l=Id,e=span"`
					EventGroupTitle              string `json:"EventGroupTitle" html:"l=Title,e=span"`
					EventGroupDescription        string `json:"EventGroupDescription" html:"l=Description,e=span"`
					EventGroupDeathsNumAffected  string `json:"EventGroupDeathsNumAffected" html:"l=Deaths Num Affected,e=span"`
					EventGroupDeathsNumAtRisk    string `json:"EventGroupDeathsNumAtRisk" html:"l=Deaths Num At Risk,e=span"`
					EventGroupSeriousNumAffected string `json:"EventGroupSeriousNumAffected" html:"l=Serious Num Affected,e=span"`
					EventGroupSeriousNumAtRisk   string `json:"EventGroupSeriousNumAtRisk" html:"l=Serious Num At Risk,e=span"`
					EventGroupOtherNumAffected   string `json:"EventGroupOtherNumAffected" html:"l=Other Num Affected,e=span"`
					EventGroupOtherNumAtRisk     string `json:"EventGroupOtherNumAtRisk" html:"l=Other Num At Risk,e=span"`
				} `json:"EventGroup" html:"e=li"`
			} `json:"EventGroupList" html:"l=Event Group List,e=ul"`
			SeriousEventList struct {
				SeriousEvent []struct {
					SeriousEventTerm             string `json:"SeriousEventTerm" html:"l=Term,e=span"`
					SeriousEventOrganSystem      string `json:"SeriousEventOrganSystem" html:"l=Organ System,e=span"`
					SeriousEventSourceVocabulary string `json:"SeriousEventSourceVocabulary" html:"l=Source Vocabulary,e=span"`
					SeriousEventAssessmentType   string `json:"SeriousEventAssessmentType" html:"l=Assessment Type,e=span"`
					SeriousEventNotes            string `json:"SeriousEventNotes" html:"l=Notes,e=span"`
					SeriousEventStatsList        struct {
						SeriousEventStats []struct {
							SeriousEventStatsGroupID     string `json:"SeriousEventStatsGroupId" html:"l=Group ID,e=span"`
							SeriousEventStatsNumEvents   string `json:"SeriousEventStatsNumEvents" html:"l=Num Events,e=span"`
							SeriousEventStatsNumAffected string `json:"SeriousEventStatsNumAffected" html:"l=Num Affected,e=span"`
							SeriousEventStatsNumAtRisk   string `json:"SeriousEventStatsNumAtRisk" html:"l=Num At Risk,e=span"`
						} `json:"SeriousEventStats" html:"row"`
					} `json:"SeriousEventStatsList" html:"l=Serious Event Stats List"`
				} `json:"SeriousEvent" html:"e=li"`
			} `json:"SeriousEventList" html:"l=Serious Event List,e=ul"`
			OtherEventList struct {
				OtherEvent []struct {
					OtherEventTerm             string `json:"OtherEventTerm" html:"l=Term,e=span"`
					OtherEventOrganSystem      string `json:"OtherEventOrganSystem" html:"l=Organ System,e=span"`
					OtherEventSourceVocabulary string `json:"OtherEventSourceVocabulary" html:"l=Source Vocabulary,e=span"`
					OtherEventAssessmentType   string `json:"OtherEventAssessmentType" html:"l=Assessment Type,e=span"`
					OtherEventNotes            string `json:"OtherEventNotes" html:"l=Notes,e=span"`
					OtherEventStatsList        struct {
						OtherEventStats []struct {
							OtherEventStatsGroupID     string `json:"OtherEventStatsGroupId" html:"l=Group ID,e=span"`
							OtherEventStatsNumEvents   string `json:"OtherEventStatsNumEvents" html:"l=Num Events,e=span"`
							OtherEventStatsNumAffected string `json:"OtherEventStatsNumAffected" html:"l=Num Affected,e=span"`
							OtherEventStatsNumAtRisk   string `json:"OtherEventStatsNumAtRisk" html:"l=Num At Risk,e=span"`
						} `json:"OtherEventStats" html:"row"`
					} `json:"OtherEventStatsList" html:"l=Other Event Stats List"`
				} `json:"OtherEvent" html:"e=li"`
			} `json:"OtherEventList" html:"l=Other Event List,e=ul"`
		} `json:"AdverseEventsModule" html:"l=AdverseEventsModule,c=top-level-module"`
		MoreInfoModule struct {
			LimitationsAndCaveatsDescription string `json:"LimitationsAndCaveatsDescription" html:"l=Limitations and Caveats Description,e=span"`
			CertainAgreement                 struct {
				AgreementPISponsorEmployee    string `json:"AgreementPISponsorEmployee" html:"l=PI Sponsor Employee,e=span"`
				AgreementRestrictiveAgreement string `json:"AgreementRestrictiveAgreement" html:"l=Restrictive Agreement,e=span"`
				AgreementRestrictionType      string `json:"AgreementRestrictionType" html:"l=Restriction Type,e=span"`
				AgreementOtherDetails         string `json:"AgreementOtherDetails" html:"l=OtherDetails,e=span"`
			} `json:"CertainAgreement" html:"l=CertainAgreement"`
			PointOfContact struct {
				PointOfContactTitle        string `json:"PointOfContactTitle" html:"l=Title,e=span"`
				PointOfContactOrganization string `json:"PointOfContactOrganization" html:"l=Organization,e=span"`
				PointOfContactEmail        string `json:"PointOfContactEMail" html:"l=EMail,e=span"`
				PointOfContactPhone        string `json:"PointOfContactPhone" html:"l=Phone,e=span"`
				PointOfContactPhoneExt     string `json:"PointOfContactPhoneExt" html:"l=Phone Extension,e=span"`
			} `json:"PointOfContact" html:"l=PointOfContact"`
		} `json:"MoreInfoModule" html:"l=More Info Module,c=top-level-module"`
	} `json:"ResultsSection,omitempty" html:"l=Results Section,c=top-level-section"`
	DocumentSection struct {
		LargeDocumentModule struct {
			LargeDocList struct {
				LargeDoc []struct {
					LargeDocTypeAbbrev  string `json:"LargeDocTypeAbbrev" html:"l=Type Abbrev,e=span"`
					LargeDocHasProtocol string `json:"LargeDocHasProtocol" html:"l=Has Protocol,e=span"`
					LargeDocHasSAP      string `json:"LargeDocHasSAP" html:"l=Has SAP,e=span"`
					LargeDocHasICF      string `json:"LargeDocHasICF" html:"l=Has ICF,e=span"`
					LargeDocLabel       string `json:"LargeDocLabel" html:"l=Label,e=span"`
					LargeDocDate        string `json:"LargeDocDate" html:"l=Date,e=span"`
					LargeDocUploadDate  string `json:"LargeDocUploadDate" html:"l=Upload Date,e=span"`
					LargeDocFilename    string `json:"LargeDocFilename" html:"l=Filename,e=span"`
				} `json:"LargeDoc" html:"e=li"`
			} `json:"LargeDocList" html:"l=Large Doc List,e=ul"`
		} `json:"LargeDocumentModule" html:"l=Large Document Module,c=top-level-module"`
	} `json:"DocumentSection" html:"l=Document Section,c=top-level-section"`
	DerivedSection struct {
		MiscInfoModule struct {
			VersionHolder string `json:"VersionHolder" html:"l=Version Holder,e=span"`
		} `json:"MiscInfoModule" html:"l=MiscInfoModule,c=top-level-module"`
		ConditionBrowseModule struct {
			ConditionMeshList struct {
				ConditionMesh []struct {
					ConditionMeshID   string `json:"ConditionMeshId" html:"l=Id,e=span"`
					ConditionMeshTerm string `json:"ConditionMeshTerm" html:"l=Term,e=span"`
				} `json:"ConditionMesh" html:"row"`
			} `json:"ConditionMeshList" html:"l=Condition Mesh List"`
			ConditionAncestorList struct {
				ConditionAncestor []struct {
					ConditionAncestorID   string `json:"ConditionAncestorId" html:"l=Id,e=span"`
					ConditionAncestorTerm string `json:"ConditionAncestorTerm" html:"l=Term,e=span"`
				} `json:"ConditionAncestor" html:"row"`
			} `json:"ConditionAncestorList" html:"l=Condition Ancestor List"`
			ConditionBrowseLeafList struct {
				ConditionBrowseLeaf []struct {
					ConditionBrowseLeafID        string `json:"ConditionBrowseLeafId" html:"l=Id,e=span"`
					ConditionBrowseLeafName      string `json:"ConditionBrowseLeafName" html:"l=Name,e=span"`
					ConditionBrowseLeafRelevance string `json:"ConditionBrowseLeafRelevance" html:"l=Relevance,e=span"`
					ConditionBrowseLeafAsFound   string `json:"ConditionBrowseLeafAsFound,omitempty" html:"l=As Found,e=span"`
				} `json:"ConditionBrowseLeaf" html:"row"`
			} `json:"ConditionBrowseLeafList" html:"l=Condition Browse Leaf List"`
			ConditionBrowseBranchList struct {
				ConditionBrowseBranch []struct {
					ConditionBrowseBranchAbbrev string `json:"ConditionBrowseBranchAbbrev" html:"l=Abbrev,e=span"`
					ConditionBrowseBranchName   string `json:"ConditionBrowseBranchName" html:"l=Name,e=span"`
				} `json:"ConditionBrowseBranch" html:"row"`
			} `json:"ConditionBrowseBranchList" html:"l=Condition Browse Branch List"`
		} `json:"ConditionBrowseModule" html:"l=ConditionBrowseModule,c=top-level-module"`
		InterventionBrowseModule struct {
			InterventionBrowseLeafList struct {
				InterventionBrowseLeaf []struct {
					InterventionBrowseLeafID        string `json:"InterventionBrowseLeafId" html:"l=Id,e=span"`
					InterventionBrowseLeafName      string `json:"InterventionBrowseLeafName" html:"l=Name,e=span"`
					InterventionBrowseLeafRelevance string `json:"InterventionBrowseLeafRelevance" html:"l=Relevance,e=span"`
				} `json:"InterventionBrowseLeaf" html:"row"`
			} `json:"InterventionBrowseLeafList" html:"l=Intervention Browse Leaf List"`
			InterventionBrowseBranchList struct {
				InterventionBrowseBranch []struct {
					InterventionBrowseBranchAbbrev string `json:"InterventionBrowseBranchAbbrev" html:"l=Abbrev,e=span"`
					InterventionBrowseBranchName   string `json:"InterventionBrowseBranchName" html:"l=Name,e=span"`
				} `json:"InterventionBrowseBranch" html:"row"`
			} `json:"InterventionBrowseBranchList" html:"l=Intervention Browse Branch List"`
		} `json:"InterventionBrowseModule" html:"l=Intervention Browse Module,c=top-level-module"`
	} `json:"DerivedSection" html:"l=Derived Section,c=top-level-section"`
}
