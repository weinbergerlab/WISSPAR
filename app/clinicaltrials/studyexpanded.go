package clinicaltrials

type StudyExpanded struct {
	ProtocolSection ProtocolSection `json:"ProtocolSection" html:"l=Protocol Section,c=top-level-section"`
	ResultsSection ResultsSection `json:"ResultsSection" html:"l=Results Section,c=top-level-section"`
	DocumentSection DocumentSection `json:"DocumentSection" html:"l=Document Section,c=top-level-section"`
	DerivedSection DerivedSection `json:"DerivedSection" html:"l=Derived Section,c=top-level-section"`
}

// Protocol Section Start
type ProtocolSection struct {
	IdentificationModule IdentificationModule `json:"IdentificationModule" html:"l=Identification Module,c=top-level-module"`
	StatusModule StatusModule `json:"StatusModule" html:"l=Status Module,c=top-level-module"`
	SponsorCollaboratorsModule SponsorCollaboratorsModule `json:"SponsorCollaboratorsModule" html:"l=Sponsor Collaborators Module,c=top-level-module"`
	OversightModule OversightModule `json:"OversightModule" html:"l=Oversight Module,c=top-level-module"`
	DescriptionModule DescriptionModule `json:"DescriptionModule" html:"l=Description Module,c=top-level-module"`
	ConditionsModule ConditionsModule `json:"ConditionsModule" html:"l=Condition Module,c=top-level-module"`
	DesignModule DesignModule `json:"DesignModule" html:"l=Design Module,c=top-level-module"`
	ArmsInterventionsModule ArmsInterventionsModule `json:"ArmsInterventionsModule" html:"l=Arms Interventions Module,c=top-level-module"`
	OutcomesModule OutcomesModule `json:"OutcomesModule" html:"l=Outcomes Module,c=top-level-module"`
	EligibilityModule EligibilityModule `json:"EligibilityModule" html:"l=Eligibility Module,c=top-level-module"`
	ContactsLocationsModule ContactsLocationsModule `json:"ContactsLocationsModule" html:"l=Contacts Locations Module,c=top-level-module"`
	ReferencesModule ReferencesModule `json:"ReferencesModule" html:"l=References Module,c=top-level-module"`
}

type IdentificationModule struct {
	NCTID          string `json:"NCTId" html:"l=Study ID,e=span"`
	OrgStudyIDInfo OrgStudyIDInfo `json:"OrgStudyIdInfo"`
	Organization Organization `json:"Organization" html:"l=Organisation Details"`
	BriefTitle    string `json:"BriefTitle" html:"l=Brief Title,e=span"`
	OfficialTitle string `json:"OfficialTitle" html:"l=Official Title,e=span"`
}

type OrgStudyIDInfo struct {
	OrgStudyID string `json:"OrgStudyId" html:"l=Organisation Study ID,e=span"`
}

type Organization struct {
	OrgFullName string `json:"OrgFullName" html:"l=Full Name,e=span"`
	OrgClass    string `json:"OrgClass" html:"l=Class,e=span"`
}

type StatusModule struct {
	StatusVerifiedDate string `json:"StatusVerifiedDate" html:"l=Status Verified Date,e=span"`
	OverallStatus      string `json:"OverallStatus" html:"l=Overall Status,e=span"`
	ExpandedAccessInfo ExpandedAccessInfo `json:"ExpandedAccessInfo"`
	StartDateStruct StartDateStruct `json:"StartDateStruct" html:"l=Start Date,c=date"`
	PrimaryCompletionDateStruct PrimaryCompletionDateStruct `json:"PrimaryCompletionDateStruct" html:"l=Primary Completion Date,c=date"`
	CompletionDateStruct CompletionDateStruct `json:"CompletionDateStruct" html:"l=Completion Date,c=date"`
	StudyFirstSubmitDate     string `json:"StudyFirstSubmitDate" html:"l=Study First Submit Date,e=span"`
	StudyFirstSubmitQCDate   string `json:"StudyFirstSubmitQCDate" html:"l=Study First Submit QC Date,e=span"`
	StudyFirstPostDateStruct StudyFirstPostDateStruct `json:"StudyFirstPostDateStruct" html:"l=Study First Post Date,c=date"`
	ResultsFirstSubmitDate     string `json:"ResultsFirstSubmitDate,e=span"`
	ResultsFirstSubmitQCDate   string `json:"ResultsFirstSubmitQCDate,e=span"`
	ResultsFirstPostDateStruct ResultsFirstPostDateStruct `json:"ResultsFirstPostDateStruct" html:"l=Results First Post Date,c=date"`
	LastUpdateSubmitDate     string `json:"LastUpdateSubmitDate,e=span"`
	LastUpdatePostDateStruct LastUpdatePostDateStruct `json:"LastUpdatePostDateStruct" html:"l=Results First Post Date,c=date"`
}

type ExpandedAccessInfo struct {
	HasExpandedAccess string `json:"HasExpandedAccess" html:"l=Has Expanded Access,e=span"`
}

type StartDateStruct struct {
	StartDate     string `json:"StartDate,e=span"`
	StartDateType string `json:"StartDateType,e=span"`
}

type PrimaryCompletionDateStruct struct {
	PrimaryCompletionDate     string `json:"PrimaryCompletionDate,e=span"`
	PrimaryCompletionDateType string `json:"PrimaryCompletionDateType,e=span"`
}

type CompletionDateStruct struct {
	CompletionDate     string `json:"CompletionDate,e=span"`
	CompletionDateType string `json:"CompletionDateType,e=span"`
}

type StudyFirstPostDateStruct struct {
	StudyFirstPostDate     string `json:"StudyFirstPostDate,e=span"`
	StudyFirstPostDateType string `json:"StudyFirstPostDateType,e=span"`
}

type ResultsFirstPostDateStruct struct {
	ResultsFirstPostDate     string `json:"ResultsFirstPostDate,e=span"`
	ResultsFirstPostDateType string `json:"ResultsFirstPostDateType,e=span"`
}

type LastUpdatePostDateStruct struct {
	LastUpdatePostDate     string `json:"LastUpdatePostDate,e=span"`
	LastUpdatePostDateType string `json:"LastUpdatePostDateType,e=span"`
}

type SponsorCollaboratorsModule struct {
	ResponsibleParty ResponsibleParty `json:"ResponsibleParty"`
	LeadSponsor LeadSponsor `json:"LeadSponsor" html:"l=Lead Sponsor"`
}

type ResponsibleParty struct {
	ResponsiblePartyType string `json:"ResponsiblePartyType" html:"l=Responsible Party Type,e=span"`
}

type LeadSponsor struct {
	LeadSponsorName  string `json:"LeadSponsorName" html:"l=Name,e=span"`
	LeadSponsorClass string `json:"LeadSponsorClass" html:"l=Class,e=span"`
}

type OversightModule struct {
	OversightHasDMC      string `json:"OversightHasDMC,e=span"`
	IsFDARegulatedDrug   string `json:"IsFDARegulatedDrug,e=span"`
	IsFDARegulatedDevice string `json:"IsFDARegulatedDevice,e=span"`
}

type DescriptionModule struct {
	BriefSummary        string `json:"BriefSummary" html:"l=Brief Summary,e=span"`
	DetailedDescription string `json:"DetailedDescription" html:"l=Detailed Description,e=span"`
}

type ConditionsModule struct {
	ConditionList ConditionList `json:"ConditionList" html:"l=Condition List,e=ul"`
}

type ConditionList struct {
	Condition []string `json:"Condition" html:"e=li"`
}

type DesignModule struct {
	StudyType string `json:"StudyType,e=span"`
	PhaseList PhaseList `json:"PhaseList" html:"l=Phases,e=ul"`
	DesignInfo DesignInfo `json:"DesignInfo" html:"l=Design Information"`
	EnrollmentInfo EnrollmentInfo `json:"EnrollmentInfo" html:"l=Enrollment Information"`
}

type PhaseList struct {
	Phase []string `json:"Phase" html:"e=li"`
}

type DesignInfo struct {
	DesignAllocation        string `json:"DesignAllocation" html:"l=Allocation,e=span"`
	DesignInterventionModel string `json:"DesignInterventionModel" html:"l=Intervention Model,e=span"`
	DesignPrimaryPurpose    string `json:"DesignPrimaryPurpose" html:"l=Primary Purpose,e=span"`
	DesignMaskingInfo       DesignMaskingInfo `json:"DesignMaskingInfo" html:"l=Design Masking Information"`
}

type DesignMaskingInfo  struct {
	DesignMasking       string `json:"DesignMasking" html:"l=Design Masking,e=span"`
	DesignWhoMaskedList DesignWhoMaskedList `json:"DesignWhoMaskedList" html:"l=Who Masked List,e=ul"`
}

type DesignWhoMaskedList struct {
	DesignWhoMasked []string `json:"DesignWhoMasked" html:"e=li"`
}

type EnrollmentInfo struct {
	EnrollmentCount string `json:"EnrollmentCount" html:"l=Count,e=span"`
	EnrollmentType  string `json:"EnrollmentType" html:"l=Type,e=span"`
}

type ArmsInterventionsModule struct {
	ArmGroupList ArmGroupList `json:"ArmGroupList" html:"l=Arm Group List"`
	InterventionList InterventionList `json:"InterventionList" html:"l=Intervention List,e=ul"`
}

type ArmGroupList struct {
	ArmGroup []ArmGroup `json:"ArmGroup" html:"row"`
}

type ArmGroup struct {
	ArmGroupLabel            string `json:"ArmGroupLabel" html:"l=Label,e=span"`
	ArmGroupType             string `json:"ArmGroupType" html:"l=Type,e=span"`
	ArmGroupDescription      string `json:"ArmGroupDescription" html:"l=Description,e=span"`
	ArmGroupInterventionList ArmGroupInterventionList `json:"ArmGroupInterventionList" html:"l=Intervention List,e=ul"`
}

type ArmGroupInterventionList struct {
	ArmGroupInterventionName []string `json:"ArmGroupInterventionName" html:"e=li"`
}

type InterventionList struct {
	Intervention []Intervention `json:"Intervention" html:"e=li"`
}

type Intervention struct {
	InterventionType              string `json:"InterventionType" html:"l=Type,e=span"`
	InterventionName              string `json:"InterventionName" html:"l=Name,e=span"`
	InterventionDescription       string `json:"InterventionDescription" html:"l=Description,e=span"`
	InterventionArmGroupLabelList InterventionArmGroupLabelList `json:"InterventionArmGroupLabelList" html:"l=Arm Group Label List,e=ul"`
}

type InterventionArmGroupLabelList struct {
	InterventionArmGroupLabel []string `json:"InterventionArmGroupLabel" html:"e=li"`
}

type OutcomesModule struct {
	PrimaryOutcomeList PrimaryOutcomeList `json:"PrimaryOutcomeList" html:"l=Primary Outcome List,e=ul"`
	SecondaryOutcomeList SecondaryOutcomeList `json:"SecondaryOutcomeList" html:"l=Secondary Outcome List,e=ul"`
	OtherOutcomeList OtherOutcomeList `json:"OtherOutcomeList" html:"l=Other Outcome List,e=ul"`
}

type PrimaryOutcomeList struct {
	PrimaryOutcome []PrimaryOutcome `json:"PrimaryOutcome" html:"e=li"`
}

type PrimaryOutcome struct {
	PrimaryOutcomeMeasure     string `json:"PrimaryOutcomeMeasure" html:"l=Outcome Measure,e=span"`
	PrimaryOutcomeDescription string `json:"PrimaryOutcomeDescription" html:"l=Outcome Description,e=span"`
	PrimaryOutcomeTimeFrame   string `json:"PrimaryOutcomeTimeFrame" html:"l=Outcome Time Frame,e=span"`
}

type SecondaryOutcomeList struct {
	SecondaryOutcome []SecondaryOutcome `json:"SecondaryOutcome" html:"e=li"`
}

type SecondaryOutcome struct {
	SecondaryOutcomeMeasure     string `json:"SecondaryOutcomeMeasure" html:"l=Outcome Measure,e=span"`
	SecondaryOutcomeDescription string `json:"SecondaryOutcomeDescription" html:"l=Outcome Description,e=span"`
	SecondaryOutcomeTimeFrame   string `json:"SecondaryOutcomeTimeFrame" html:"l=Outcome Time Frame,e=span"`
}

type OtherOutcomeList struct {
	OtherOutcome OtherOutcome `json:"OtherOutcome" html:"e=li"`
}

type OtherOutcome struct {
	OtherOutcomeMeasure     string `json:"OtherOutcomeMeasure" html:"l=Outcome Measure,e=span"`
	OtherOutcomeDescription string `json:"OtherOutcomeDescription" html:"l=Outcome Description,e=span"`
	OtherOutcomeTimeFrame   string `json:"OtherOutcomeTimeFrame" html:"l=Outcome Time Frame,e=span"`
}

type EligibilityModule struct {
	EligibilityCriteria string `json:"EligibilityCriteria,e=span"`
	HealthyVolunteers   string `json:"HealthyVolunteers,e=span"`
	Gender              string `json:"Gender,e=span"`
	MinimumAge          string `json:"MinimumAge,e=span"`
	MaximumAge          string `json:"MaximumAge,e=span"`
	StdAgeList          StdAgeList `json:"StdAgeList" html:"l=Standard Age List,e=ul"`
}

type StdAgeList struct {
	StdAge []string `json:"StdAge" html:"e=li"`
}

type ContactsLocationsModule struct {
	OverallOfficialList OverallOfficialList `json:"OverallOfficialList" html:"l=Overall Official List,e=ul"`
	LocationList LocationList `json:"LocationList" html:"l=Location List"`
}

type OverallOfficialList struct {
	OverallOfficial OverallOfficial `json:"OverallOfficial" html:"e=li"`
}

type OverallOfficial struct {
	OverallOfficialName        string `json:"OverallOfficialName" html:"l=Name,e=span"`
	OverallOfficialAffiliation string `json:"OverallOfficialAffiliation" html:"l=Affiliation,e=span"`
	OverallOfficialRole        string `json:"OverallOfficialRole" html:"l=Role,e=span"`
}

type LocationList struct {
	Location []Location `json:"Location" html:"row"`
}

type Location struct {
	LocationFacility string `json:"LocationFacility" html:"l=Facility,e=span"`
	LocationCity     string `json:"LocationCity" html:"l=City,e=span"`
	LocationCountry  string `json:"LocationCountry" html:"l=Country,e=span"`
}

type ReferencesModule struct {
	ReferenceList ReferenceList `json:"ReferenceList" html:"l=Reference List,e=ul"`
}

type ReferenceList struct {
	Reference []Reference `json:"Reference" html:"e=li"`
}

type Reference struct {
	ReferencePMID     string `json:"ReferencePMID" html:"l=PM ID,e=span"`
	ReferenceType     string `json:"ReferenceType" html:"l=Type,e=span"`
	ReferenceCitation string `json:"ReferenceCitation" html:"l=Citation,e=span"`
}

// Results Section Start
type ResultsSection struct {
	ParticipantFlowModule ParticipantFlowModule `json:"ParticipantFlowModule" html:"l=Participant Flow Module,c=top-level-module"`
	BaselineCharacteristicsModule BaselineCharacteristicsModule `json:"BaselineCharacteristicsModule" html:"l=Baseline Characteristics Module,c=top-level-module"`
	OutcomeMeasuresModule OutcomeMeasuresModule `json:"OutcomeMeasuresModule" html:"l=Outcome Measures Module,c=top-level-module"`
	AdverseEventsModule AdverseEventsModule `json:"AdverseEventsModule" html:"l=AdverseEventsModule,c=top-level-module"`
	MoreInfoModule MoreInfoModule `json:"MoreInfoModule" html:"l=More Info Module,c=top-level-module"`
}

type ParticipantFlowModule struct {
	FlowPreAssignmentDetails string `json:"FlowPreAssignmentDetails" html:"l=Flow Pre Assignment Details,e=span"`
	FlowRecruitmentDetails   string `json:"FlowRecruitmentDetails" html:"l=Flow Recruitment Details,e=span"`
	FlowTypeUnitsAnalyzed    string `json:"FlowTypeUnitsAnalyzed" html:"l=Flow Type Units Analyzed,e=span"`
	FlowGroupList            FlowGroupList `json:"FlowGroupList" html:"l=Flow Group List"`
	FlowPeriodList FlowPeriodList `json:"FlowPeriodList" html:"l=Flow Period List,e=ul"`
}

type FlowGroupList struct {
	FlowGroup []FlowGroup `json:"FlowGroup" html:"row"`
}

type FlowGroup struct {
	FlowGroupID          string `json:"FlowGroupId" html:"l=ID,e=span"`
	FlowGroupTitle       string `json:"FlowGroupTitle" html:"l=Title,e=span"`
	FlowGroupDescription string `json:"FlowGroupDescription" html:"l=Description,e=span"`
}

type FlowPeriodList struct {
	FlowPeriod []FlowPeriod `json:"FlowPeriod" html:"e=li"`
}

type FlowPeriod struct {
	FlowPeriodTitle   string `json:"FlowPeriodTitle" html:"l=Title,e=span"`
	FlowMilestoneList FlowMilestoneList `json:"FlowMilestoneList" html:"l=Flow Milestone List,e=ul"`
	FlowDropWithdrawList FlowDropWithdrawList `json:"FlowDropWithdrawList" html:"l=Flow Drop Withdraw List,e=ul"`
}

type FlowMilestoneList struct {
	FlowMilestone []FlowMilestone `json:"FlowMilestone" html:"e=li"`
}

type FlowMilestone struct {
	FlowMilestoneType   string `json:"FlowMilestoneType" html:"l=Type,e=span"`
	FlowAchievementList FlowAchievementList `json:"FlowAchievementList" html:"l=Flow Achievement List"`
	FlowMilestoneComment string `json:"FlowMilestoneComment,omitempty" html:"l=Comment,e=span,omitempty"`
}

type FlowAchievementList struct {
	FlowAchievement []FlowAchievement `json:"FlowAchievement" html:"row"`
}

type FlowAchievement struct {
	FlowAchievementGroupID     string `json:"FlowAchievementGroupId" html:"l=Group ID,e=span"`
	FlowAchievementNumSubjects string `json:"FlowAchievementNumSubjects" html:"l=Num Subjects,e=span"`
	FlowAchievementComment     string `json:"FlowAchievementComment" html:"l=Comment,e=span"`
	FlowAchievementNumUnits    string `json:"FlowAchievementNumUnits" html:"l=Num Units,e=span"`
}

type FlowDropWithdrawList struct {
	FlowDropWithdraw []FlowDropWithdraw `json:"FlowDropWithdraw" html:"e=li"`
}

type FlowDropWithdraw struct {
	FlowDropWithdrawType    string `json:"FlowDropWithdrawType" html:"l=Type,e=span"`
	FlowDropWithdrawComment string `json:"FlowDropWithdrawComment" html:"l=Comment,e=span,omitempty"`
	FlowReasonList          FlowReasonList `json:"FlowReasonList" html:"l=Flow Reason List"`
}

type FlowReasonList struct {
	FlowReason []FlowReason `json:"FlowReason" html:"row"`
}

type FlowReason struct {
	FlowReasonGroupID     string `json:"FlowReasonGroupId" html:"l=Group ID,e=span"`
	FlowReasonNumSubjects string `json:"FlowReasonNumSubjects" html:"l=Num Subjects,e=span"`
	FlowReasonNumUnits    string `json:"FlowReasonNumUnits" html:"l=Num Units,e=span"`
	FlowReasonComment     string `json:"FlowReasonComment" html:"l=Comment,e=span"`
}

type BaselineCharacteristicsModule struct {
	BaselinePopulationDescription string `json:"BaselinePopulationDescription" html:"l=Population Description,e=span,omitempty"`
	BaselineTypeUnitsAnalyzed     string `json:"BaselineTypeUnitsAnalyzed" html:"l=Type Units Analyzed,e=span"`
	BaselineGroupList             BaselineGroupList `json:"BaselineGroupList" html:"l=Baseline Group List"`
	BaselineDenomList BaselineDenomList `json:"BaselineDenomList" html:"l=Baseline Denom List,e=ul"`
	BaselineMeasureList BaselineMeasureList `json:"BaselineMeasureList" html:"l=Baseline Measure List,e=ul"`
}

type BaselineGroupList struct {
	BaselineGroup []BaselineGroup `json:"BaselineGroup" html:"row"`
}

type BaselineGroup struct {
	BaselineGroupID          string `json:"BaselineGroupId" html:"l=Group ID,e=span"`
	BaselineGroupTitle       string `json:"BaselineGroupTitle" html:"l=Title,e=span"`
	BaselineGroupDescription string `json:"BaselineGroupDescription" html:"l=Description,e=span"`
}

type BaselineDenomList struct {
	BaselineDenom []BaselineDenom `json:"BaselineDenom" html:"e=li"`
}

type BaselineDenom struct {
	BaselineDenomUnits     string `json:"BaselineDenomUnits" html:"l=Denom Units,e=span"`
	BaselineDenomCountList BaselineDenomCountList `json:"BaselineDenomCountList" html:"l=Baseline Denom Count List"`
}

type BaselineDenomCountList struct {
	BaselineDenomCount []BaselineDenomCount `json:"BaselineDenomCount" html:"row"`
}

type BaselineDenomCount struct {
	BaselineDenomCountGroupID string `json:"BaselineDenomCountGroupId" html:"l=Group ID,e=span"`
	BaselineDenomCountValue   string `json:"BaselineDenomCountValue" html:"l=Value,e=span"`
}

type BaselineMeasureList struct {
	BaselineMeasure []BaselineMeasure `json:"BaselineMeasure" html:"e=li"`
}

type BaselineMeasure struct {
	BaselineMeasureTitle                 string `json:"BaselineMeasureTitle" html:"l=Title,e=span"`
	BaselineMeasureDescription           string `json:"BaselineMeasureDescription,omitempty" html:"l=Description,e=span"`
	BaselineMeasureParamType             string `json:"BaselineMeasureParamType" html:"l=Param Type,e=span"`
	BaselineMeasureDispersionType        string `json:"BaselineMeasureDispersionType,omitempty" html:"l=Dispersion Type,e=span"`
	BaselineMeasureUnitOfMeasure         string `json:"BaselineMeasureUnitOfMeasure" html:"l=Unit Of Measure,e=span"`
	BaselineMeasureCalculatePct          string `json:"BaselineMeasureCalculatePct" html:"l=Calculate PCT,e=span,omitempty"`
	BaselineMeasurePopulationDescription string `json:"BaselineMeasurePopulationDescription" html:"l=Population Description,e=span,omitempty"`
	BaselineMeasureDenomUnitsSelected    string `json:"BaselineMeasureDenomUnitsSelected" html:"l=Denom Units Selected,e=span"`
	BaselineClassList                    BaselineClassList `json:"BaselineClassList" html:"l=Baseline Class List,e=ul"`
	BaselineMeasureDenomList BaselineMeasureDenomList `json:"BaselineMeasureDenomList" html:"l=Baseline Measure Denom List,e=ul"`
}

type BaselineClassList struct {
	BaselineClass []BaselineClass `json:"BaselineClass" html:"e=li"`
}

type BaselineClass struct {
	BaselineClassTitle   string `json:"BaselineClassTitle" html:"l=Title,e=span"`
	BaselineCategoryList BaselineCategoryList `json:"BaselineCategoryList,e=ul"`
}

type BaselineCategoryList struct {
	BaselineCategory []BaselineCategory `json:"BaselineCategory" html:"e=li"`
}

type BaselineCategory struct {
	BaselineCategoryTitle   string `json:"BaselineCategoryTitle" html:"l=Title,e=span"`
	BaselineMeasurementList BaselineMeasurementList `json:"BaselineMeasurementList" html:"l=Baseline Measurement List"`
}

type BaselineMeasurementList struct {
	BaselineMeasurement []BaselineMeasurement `json:"BaselineMeasurement" html:"row"`
}

type BaselineMeasurement struct {
	BaselineMeasurementGroupID    string `json:"BaselineMeasurementGroupId" html:"l=Group ID,e=span"`
	BaselineMeasurementValue      string `json:"BaselineMeasurementValue" html:"l=Value,e=span"`
	BaselineMeasurementSpread     string `json:"BaselineMeasurementSpread" html:"l=Spread,e=span"`
	BaselineMeasurementComment    string `json:"BaselineMeasurementComment" html:"l=Comment,e=span"`
	BaselineMeasurementUpperLimit string `json:"BaselineMeasurementUpperLimit" html:"l=Upper Limit,e=span"`
	BaselineMeasurementLowerLimit string `json:"BaselineMeasurementLowerLimit" html:"l=Lower Limit,e=span"`
}

type BaselineMeasureDenomList struct {
	BaselineMeasureDenom []BaselineMeasureDenom `json:"BaselineMeasureDenom" html:"e=li"`
}

type BaselineMeasureDenom struct {
	BaselineMeasureDenomUnits     string `json:"BaselineMeasureDenomUnits" html:"l=Denom Units,e=span"`
	BaselineMeasureDenomCountList BaselineMeasureDenomCountList `json:"BaselineMeasureDenomCountList" html:"l=Baseline Measure Denom Count List"`
}

type BaselineMeasureDenomCountList struct {
	BaselineMeasureDenomCount []BaselineMeasureDenomCount `json:"BaselineMeasureDenomCount" html:"row"`
}

type BaselineMeasureDenomCount struct {
	BaselineMeasureDenomCountGroupID string `json:"BaselineMeasureDenomCountGroupID" html:"l=Group ID,e=span"`
	BaselineMeasureDenomCountValue   string `json:"BaselineMeasureDenomCountValue" html:"l=Value,e=span"`
}

type OutcomeMeasuresModule struct {
	OutcomeMeasureList OutcomeMeasureList `json:"OutcomeMeasureList" html:"l=Outcome Measure List,e=ul"`
}

type OutcomeMeasureList struct {
	OutcomeMeasure []OutcomeMeasure `json:"OutcomeMeasure" html:"e=li"`
}

type OutcomeMeasure struct {
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
	OutcomeGroupList                     OutcomeGroupList `json:"OutcomeGroupList" html:"l=Outcome Group List"`
	OutcomeDenomList OutcomeDenomList `json:"OutcomeDenomList" html:"l=Outcome Denom List,e=ul"`
	OutcomeClassList OutcomeClassList `json:"OutcomeClassList" html:"l=Outcome Class List,e=ul"`
	OutcomeAnalysisList OutcomeAnalysisList `json:"OutcomeAnalysisList,omitempty" html:"l=Outcome Analysis List,e=ul"`
}

type OutcomeGroupList struct {
	OutcomeGroup []OutcomeGroup `json:"OutcomeGroup" html:"row"`
}

type OutcomeGroup struct {
	OutcomeGroupID          string `json:"OutcomeGroupId" html:"l=Group ID,e=span"`
	OutcomeGroupTitle       string `json:"OutcomeGroupTitle" html:"l=Group Title,e=span"`
	OutcomeGroupDescription string `json:"OutcomeGroupDescription" html:"l=Description,e=span"`
}

type OutcomeDenomList struct {
	OutcomeDenom []OutcomeDenom `json:"OutcomeDenom" html:"e=li"`
}

type OutcomeDenom struct {
	OutcomeDenomUnits     string `json:"OutcomeDenomUnits" html:"l=Denom Units,e=span"`
	OutcomeDenomCountList OutcomeDenomCountList `json:"OutcomeDenomCountList" html:"l=Outcome Denom Count List"`
}

type OutcomeDenomCountList struct {
	OutcomeDenomCount []OutcomeDenomCount `json:"OutcomeDenomCount" html:"row"`
}

type OutcomeDenomCount struct {
	OutcomeDenomCountGroupID string `json:"OutcomeDenomCountGroupId" html:"l=Group ID,e=span"`
	OutcomeDenomCountValue   string `json:"OutcomeDenomCountValue" html:"l=Value,e=span"`
}

type OutcomeClassList struct {
	OutcomeClass []OutcomeClass `json:"OutcomeClass" html:"e=li"`
}

type OutcomeClass struct {
	OutcomeClassTitle     string `json:"OutcomeClassTitle" html:"l=Title,e=span"`
	OutcomeClassDenomList OutcomeClassDenomList `json:"OutcomeClassDenomList" html:"l=Outcome Class Denom List,e=ul"`
	OutcomeCategoryList OutcomeCategoryList `json:"OutcomeCategoryList" html:"l=Outcome Category List,e=ul"`
}

type OutcomeClassDenomList struct {
	OutcomeClassDenom []OutcomeClassDenom `json:"OutcomeClassDenom" html:"e=li"`
}

type OutcomeClassDenom struct {
	OutcomeClassDenomUnits     string `json:"OutcomeClassDenomUnits" html:"l=Denom Units,e=span"`
	OutcomeClassDenomCountList OutcomeClassDenomCountList `json:"OutcomeClassDenomCountList" html:"l=Outcome Class Denom Count List"`
}

type OutcomeClassDenomCountList struct {
	OutcomeClassDenomCount []OutcomeClassDenomCount `json:"OutcomeClassDenomCount" html:"row"`
}

type OutcomeClassDenomCount struct {
	OutcomeClassDenomCountGroupID string `json:"OutcomeClassDenomCountGroupId" html:"l=Group ID,e=span"`
	OutcomeClassDenomCountValue   string `json:"OutcomeClassDenomCountValue" html:"l=Value,e=span"`
}

type OutcomeCategoryList struct {
	OutcomeCategory []OutcomeCategory `json:"OutcomeCategory" html:"e=li"`
}

type OutcomeCategory struct {
	OutcomeCategoryTitle   string `json:"OutcomeCategoryTitle" html:"l=Title,e=span"`
	OutcomeMeasurementList OutcomeMeasurementList `json:"OutcomeMeasurementList" html:"l=Outcome Measurement List"`
}

type OutcomeMeasurementList struct {
	OutcomeMeasurement []OutcomeMeasurement `json:"OutcomeMeasurement" html:"row"`
}

type OutcomeMeasurement struct {
	OutcomeMeasurementGroupID    string `json:"OutcomeMeasurementGroupId" html:"l=Group ID,e=span"`
	OutcomeMeasurementValue      string `json:"OutcomeMeasurementValue" html:"l=Value,e=span"`
	OutcomeMeasurementLowerLimit string `json:"OutcomeMeasurementLowerLimit" html:"l=Lower Limit,e=span"`
	OutcomeMeasurementUpperLimit string `json:"OutcomeMeasurementUpperLimit" html:"l=Upper Limit,e=span"`
	OutcomeMeasurementSpread     string `json:"OutcomeMeasurementSpread" html:"l=Spread,e=span"`
	OutcomeMeasurementComment    string `json:"OutcomeMeasurementComment" html:"l=Comment,e=span"`
}

type OutcomeAnalysisList struct {
	OutcomeAnalysis []OutcomeAnalysis `json:"OutcomeAnalysis" html:"e=li"`
}

type OutcomeAnalysis struct {
	OutcomeAnalysisGroupIDList OutcomeAnalysisGroupIDList `json:"OutcomeAnalysisGroupIdList" html:"l=Outcome Analysis Group ID List,e=ul"`
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
}

type OutcomeAnalysisGroupIDList struct {
	OutcomeAnalysisGroupID []string `json:"OutcomeAnalysisGroupId" html:"e=li"`
}

type AdverseEventsModule struct {
	EventsFrequencyThreshold string `json:"EventsFrequencyThreshold" html:"l=Frequency Threshold,e=span"`
	EventsTimeFrame          string `json:"EventsTimeFrame" html:"l=Time Frame,e=span"`
	EventsDescription        string `json:"EventsDescription" html:"l=Description,e=span"`
	EventGroupList           EventGroupList `json:"EventGroupList" html:"l=Event Group List,e=ul"`
	SeriousEventList SeriousEventList `json:"SeriousEventList" html:"l=Serious Event List,e=ul"`
	OtherEventList OtherEventList `json:"OtherEventList" html:"l=Other Event List,e=ul"`
}

type EventGroupList struct {
	EventGroup []EventGroup `json:"EventGroup" html:"e=li"`
}

type EventGroup struct {
	EventGroupID                 string `json:"EventGroupId" html:"l=Id,e=span"`
	EventGroupTitle              string `json:"EventGroupTitle" html:"l=Title,e=span"`
	EventGroupDescription        string `json:"EventGroupDescription" html:"l=Description,e=span"`
	EventGroupDeathsNumAffected  string `json:"EventGroupDeathsNumAffected" html:"l=Deaths Num Affected,e=span"`
	EventGroupDeathsNumAtRisk    string `json:"EventGroupDeathsNumAtRisk" html:"l=Deaths Num At Risk,e=span"`
	EventGroupSeriousNumAffected string `json:"EventGroupSeriousNumAffected" html:"l=Serious Num Affected,e=span"`
	EventGroupSeriousNumAtRisk   string `json:"EventGroupSeriousNumAtRisk" html:"l=Serious Num At Risk,e=span"`
	EventGroupOtherNumAffected   string `json:"EventGroupOtherNumAffected" html:"l=Other Num Affected,e=span"`
	EventGroupOtherNumAtRisk     string `json:"EventGroupOtherNumAtRisk" html:"l=Other Num At Risk,e=span"`
}

type SeriousEventList struct {
	SeriousEvent []SeriousEvent `json:"SeriousEvent" html:"e=li"`
}

type SeriousEvent struct {
	SeriousEventTerm             string `json:"SeriousEventTerm" html:"l=Term,e=span"`
	SeriousEventOrganSystem      string `json:"SeriousEventOrganSystem" html:"l=Organ System,e=span"`
	SeriousEventSourceVocabulary string `json:"SeriousEventSourceVocabulary" html:"l=Source Vocabulary,e=span"`
	SeriousEventAssessmentType   string `json:"SeriousEventAssessmentType" html:"l=Assessment Type,e=span"`
	SeriousEventNotes            string `json:"SeriousEventNotes" html:"l=Notes,e=span"`
	SeriousEventStatsList        SeriousEventStatsList `json:"SeriousEventStatsList" html:"l=Serious Event Stats List"`
}

type SeriousEventStatsList struct {
	SeriousEventStats []SeriousEventStats `json:"SeriousEventStats" html:"row"`
}

type SeriousEventStats struct {
	SeriousEventStatsGroupID     string `json:"SeriousEventStatsGroupId" html:"l=Group ID,e=span"`
	SeriousEventStatsNumEvents   string `json:"SeriousEventStatsNumEvents" html:"l=Num Events,e=span"`
	SeriousEventStatsNumAffected string `json:"SeriousEventStatsNumAffected" html:"l=Num Affected,e=span"`
	SeriousEventStatsNumAtRisk   string `json:"SeriousEventStatsNumAtRisk" html:"l=Num At Risk,e=span"`
}

type OtherEventList struct {
	OtherEvent []OtherEvent `json:"OtherEvent" html:"e=li"`
}

type OtherEvent struct {
	OtherEventTerm             string `json:"OtherEventTerm" html:"l=Term,e=span"`
	OtherEventOrganSystem      string `json:"OtherEventOrganSystem" html:"l=Organ System,e=span"`
	OtherEventSourceVocabulary string `json:"OtherEventSourceVocabulary" html:"l=Source Vocabulary,e=span"`
	OtherEventAssessmentType   string `json:"OtherEventAssessmentType" html:"l=Assessment Type,e=span"`
	OtherEventNotes            string `json:"OtherEventNotes" html:"l=Notes,e=span"`
	OtherEventStatsList        OtherEventStatsList `json:"OtherEventStatsList" html:"l=Other Event Stats List"`
}

type OtherEventStatsList struct {
	OtherEventStats []OtherEventStats `json:"OtherEventStats" html:"row"`
}

type OtherEventStats struct {
	OtherEventStatsGroupID     string `json:"OtherEventStatsGroupId" html:"l=Group ID,e=span"`
	OtherEventStatsNumEvents   string `json:"OtherEventStatsNumEvents" html:"l=Num Events,e=span"`
	OtherEventStatsNumAffected string `json:"OtherEventStatsNumAffected" html:"l=Num Affected,e=span"`
	OtherEventStatsNumAtRisk   string `json:"OtherEventStatsNumAtRisk" html:"l=Num At Risk,e=span"`
}

type MoreInfoModule struct {
	LimitationsAndCaveatsDescription string `json:"LimitationsAndCaveatsDescription" html:"l=Limitations and Caveats Description,e=span"`
	CertainAgreement                 CertainAgreement `json:"CertainAgreement" html:"l=CertainAgreement"`
	PointOfContact PointOfContact `json:"PointOfContact" html:"l=PointOfContact"`
}

type CertainAgreement struct {
	AgreementPISponsorEmployee    string `json:"AgreementPISponsorEmployee" html:"l=PI Sponsor Employee,e=span"`
	AgreementRestrictiveAgreement string `json:"AgreementRestrictiveAgreement" html:"l=Restrictive Agreement,e=span"`
	AgreementRestrictionType      string `json:"AgreementRestrictionType" html:"l=Restriction Type,e=span"`
	AgreementOtherDetails         string `json:"AgreementOtherDetails" html:"l=OtherDetails,e=span"`
}

type PointOfContact struct {
	PointOfContactTitle        string `json:"PointOfContactTitle" html:"l=Title,e=span"`
	PointOfContactOrganization string `json:"PointOfContactOrganization" html:"l=Organization,e=span"`
	PointOfContactEmail        string `json:"PointOfContactEMail" html:"l=EMail,e=span"`
	PointOfContactPhone        string `json:"PointOfContactPhone" html:"l=Phone,e=span"`
	PointOfContactPhoneExt     string `json:"PointOfContactPhoneExt" html:"l=Phone Extension,e=span"`
}

// Document Section Start
type DocumentSection struct {
	LargeDocumentModule LargeDocumentModule `json:"LargeDocumentModule" html:"l=Large Document Module,c=top-level-module"`
}

type LargeDocumentModule struct {
	LargeDocList LargeDocList `json:"LargeDocList" html:"l=Large Doc List,e=ul"`
}

type LargeDocList struct {
	LargeDoc []LargeDoc `json:"LargeDoc" html:"e=li"`
}

type LargeDoc struct {
	LargeDocTypeAbbrev  string `json:"LargeDocTypeAbbrev" html:"l=Type Abbrev,e=span"`
	LargeDocHasProtocol string `json:"LargeDocHasProtocol" html:"l=Has Protocol,e=span"`
	LargeDocHasSAP      string `json:"LargeDocHasSAP" html:"l=Has SAP,e=span"`
	LargeDocHasICF      string `json:"LargeDocHasICF" html:"l=Has ICF,e=span"`
	LargeDocLabel       string `json:"LargeDocLabel" html:"l=Label,e=span"`
	LargeDocDate        string `json:"LargeDocDate" html:"l=Date,e=span"`
	LargeDocUploadDate  string `json:"LargeDocUploadDate" html:"l=Upload Date,e=span"`
	LargeDocFilename    string `json:"LargeDocFilename" html:"l=Filename,e=span"`
}

// Derived Section Start
type DerivedSection struct {
	MiscInfoModule MiscInfoModule `json:"MiscInfoModule" html:"l=MiscInfoModule,c=top-level-module"`
	ConditionBrowseModule ConditionBrowseModule `json:"ConditionBrowseModule" html:"l=ConditionBrowseModule,c=top-level-module"`
	InterventionBrowseModule InterventionBrowseModule `json:"InterventionBrowseModule" html:"l=Intervention Browse Module,c=top-level-module"`
}

type MiscInfoModule struct {
	VersionHolder string `json:"VersionHolder" html:"l=Version Holder,e=span"`
}

type ConditionBrowseModule struct {
	ConditionMeshList ConditionMeshList `json:"ConditionMeshList" html:"l=Condition Mesh List"`
	ConditionAncestorList ConditionAncestorList `json:"ConditionAncestorList" html:"l=Condition Ancestor List"`
	ConditionBrowseLeafList ConditionBrowseLeafList `json:"ConditionBrowseLeafList" html:"l=Condition Browse Leaf List"`
	ConditionBrowseBranchList ConditionBrowseBranchList `json:"ConditionBrowseBranchList" html:"l=Condition Browse Branch List"`
}

type ConditionMeshList struct {
	ConditionMesh []ConditionMesh `json:"ConditionMesh" html:"row"`
}

type ConditionMesh struct {
	ConditionMeshID   string `json:"ConditionMeshId" html:"l=Id,e=span"`
	ConditionMeshTerm string `json:"ConditionMeshTerm" html:"l=Term,e=span"`
}

type ConditionAncestorList struct {
	ConditionAncestor []ConditionAncestor `json:"ConditionAncestor" html:"row"`
}

type ConditionAncestor struct {
	ConditionAncestorID   string `json:"ConditionAncestorId" html:"l=Id,e=span"`
	ConditionAncestorTerm string `json:"ConditionAncestorTerm" html:"l=Term,e=span"`
}

type ConditionBrowseLeafList struct {
	ConditionBrowseLeaf []ConditionBrowseLeaf `json:"ConditionBrowseLeaf" html:"row"`
}

type ConditionBrowseLeaf struct {
	ConditionBrowseLeafID        string `json:"ConditionBrowseLeafId" html:"l=Id,e=span"`
	ConditionBrowseLeafName      string `json:"ConditionBrowseLeafName" html:"l=Name,e=span"`
	ConditionBrowseLeafRelevance string `json:"ConditionBrowseLeafRelevance" html:"l=Relevance,e=span"`
	ConditionBrowseLeafAsFound   string `json:"ConditionBrowseLeafAsFound,omitempty" html:"l=As Found,e=span"`
}

type ConditionBrowseBranchList struct {
	ConditionBrowseBranch []ConditionBrowseBranch `json:"ConditionBrowseBranch" html:"row"`
}

type ConditionBrowseBranch struct {
	ConditionBrowseBranchAbbrev string `json:"ConditionBrowseBranchAbbrev" html:"l=Abbrev,e=span"`
	ConditionBrowseBranchName   string `json:"ConditionBrowseBranchName" html:"l=Name,e=span"`
}

type InterventionBrowseModule struct {
	InterventionBrowseLeafList InterventionBrowseLeafList `json:"InterventionBrowseLeafList" html:"l=Intervention Browse Leaf List"`
	InterventionBrowseBranchList InterventionBrowseBranchList `json:"InterventionBrowseBranchList" html:"l=Intervention Browse Branch List"`
}

type InterventionBrowseLeafList struct {
	InterventionBrowseLeaf []InterventionBrowseLeaf `json:"InterventionBrowseLeaf" html:"row"`
}

type InterventionBrowseLeaf struct {
	InterventionBrowseLeafID        string `json:"InterventionBrowseLeafId" html:"l=Id,e=span"`
	InterventionBrowseLeafName      string `json:"InterventionBrowseLeafName" html:"l=Name,e=span"`
	InterventionBrowseLeafRelevance string `json:"InterventionBrowseLeafRelevance" html:"l=Relevance,e=span"`
}

type InterventionBrowseBranchList struct {
	InterventionBrowseBranch []InterventionBrowseBranch `json:"InterventionBrowseBranch" html:"row"`
}

type InterventionBrowseBranch struct {
	InterventionBrowseBranchAbbrev string `json:"InterventionBrowseBranchAbbrev" html:"l=Abbrev,e=span"`
	InterventionBrowseBranchName   string `json:"InterventionBrowseBranchName" html:"l=Name,e=span"`
}
