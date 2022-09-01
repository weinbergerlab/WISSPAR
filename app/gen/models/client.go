// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"log"

	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/migrate"

	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/adverseeventsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasurement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/certainagreement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrialmetadata"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/dosedescription"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/eventgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowachievement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowdropwithdraw"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowmilestone"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowreason"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/immunocompromisedgroups"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/manufacturer"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/otherevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/othereventstats"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysis"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysisgroupid"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclassdenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasurement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/pointofcontact"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/schedule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriousevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriouseventstats"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studyeligibility"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studylocation"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/task"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/usecase"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/vaccine"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AdverseEventsModule is the client for interacting with the AdverseEventsModule builders.
	AdverseEventsModule *AdverseEventsModuleClient
	// BaselineCategory is the client for interacting with the BaselineCategory builders.
	BaselineCategory *BaselineCategoryClient
	// BaselineCharacteristicsModule is the client for interacting with the BaselineCharacteristicsModule builders.
	BaselineCharacteristicsModule *BaselineCharacteristicsModuleClient
	// BaselineClass is the client for interacting with the BaselineClass builders.
	BaselineClass *BaselineClassClient
	// BaselineClassDenom is the client for interacting with the BaselineClassDenom builders.
	BaselineClassDenom *BaselineClassDenomClient
	// BaselineClassDenomCount is the client for interacting with the BaselineClassDenomCount builders.
	BaselineClassDenomCount *BaselineClassDenomCountClient
	// BaselineDenom is the client for interacting with the BaselineDenom builders.
	BaselineDenom *BaselineDenomClient
	// BaselineDenomCount is the client for interacting with the BaselineDenomCount builders.
	BaselineDenomCount *BaselineDenomCountClient
	// BaselineGroup is the client for interacting with the BaselineGroup builders.
	BaselineGroup *BaselineGroupClient
	// BaselineMeasure is the client for interacting with the BaselineMeasure builders.
	BaselineMeasure *BaselineMeasureClient
	// BaselineMeasureDenom is the client for interacting with the BaselineMeasureDenom builders.
	BaselineMeasureDenom *BaselineMeasureDenomClient
	// BaselineMeasureDenomCount is the client for interacting with the BaselineMeasureDenomCount builders.
	BaselineMeasureDenomCount *BaselineMeasureDenomCountClient
	// BaselineMeasurement is the client for interacting with the BaselineMeasurement builders.
	BaselineMeasurement *BaselineMeasurementClient
	// CertainAgreement is the client for interacting with the CertainAgreement builders.
	CertainAgreement *CertainAgreementClient
	// ClinicalTrial is the client for interacting with the ClinicalTrial builders.
	ClinicalTrial *ClinicalTrialClient
	// ClinicalTrialMetadata is the client for interacting with the ClinicalTrialMetadata builders.
	ClinicalTrialMetadata *ClinicalTrialMetadataClient
	// DoseDescription is the client for interacting with the DoseDescription builders.
	DoseDescription *DoseDescriptionClient
	// EventGroup is the client for interacting with the EventGroup builders.
	EventGroup *EventGroupClient
	// FlowAchievement is the client for interacting with the FlowAchievement builders.
	FlowAchievement *FlowAchievementClient
	// FlowDropWithdraw is the client for interacting with the FlowDropWithdraw builders.
	FlowDropWithdraw *FlowDropWithdrawClient
	// FlowGroup is the client for interacting with the FlowGroup builders.
	FlowGroup *FlowGroupClient
	// FlowMilestone is the client for interacting with the FlowMilestone builders.
	FlowMilestone *FlowMilestoneClient
	// FlowPeriod is the client for interacting with the FlowPeriod builders.
	FlowPeriod *FlowPeriodClient
	// FlowReason is the client for interacting with the FlowReason builders.
	FlowReason *FlowReasonClient
	// ImmunocompromisedGroups is the client for interacting with the ImmunocompromisedGroups builders.
	ImmunocompromisedGroups *ImmunocompromisedGroupsClient
	// Manufacturer is the client for interacting with the Manufacturer builders.
	Manufacturer *ManufacturerClient
	// MoreInfoModule is the client for interacting with the MoreInfoModule builders.
	MoreInfoModule *MoreInfoModuleClient
	// OtherEvent is the client for interacting with the OtherEvent builders.
	OtherEvent *OtherEventClient
	// OtherEventStats is the client for interacting with the OtherEventStats builders.
	OtherEventStats *OtherEventStatsClient
	// OutcomeAnalysis is the client for interacting with the OutcomeAnalysis builders.
	OutcomeAnalysis *OutcomeAnalysisClient
	// OutcomeAnalysisGroupID is the client for interacting with the OutcomeAnalysisGroupID builders.
	OutcomeAnalysisGroupID *OutcomeAnalysisGroupIDClient
	// OutcomeCategory is the client for interacting with the OutcomeCategory builders.
	OutcomeCategory *OutcomeCategoryClient
	// OutcomeClass is the client for interacting with the OutcomeClass builders.
	OutcomeClass *OutcomeClassClient
	// OutcomeClassDenom is the client for interacting with the OutcomeClassDenom builders.
	OutcomeClassDenom *OutcomeClassDenomClient
	// OutcomeClassDenomCount is the client for interacting with the OutcomeClassDenomCount builders.
	OutcomeClassDenomCount *OutcomeClassDenomCountClient
	// OutcomeDenom is the client for interacting with the OutcomeDenom builders.
	OutcomeDenom *OutcomeDenomClient
	// OutcomeDenomCount is the client for interacting with the OutcomeDenomCount builders.
	OutcomeDenomCount *OutcomeDenomCountClient
	// OutcomeGroup is the client for interacting with the OutcomeGroup builders.
	OutcomeGroup *OutcomeGroupClient
	// OutcomeMeasure is the client for interacting with the OutcomeMeasure builders.
	OutcomeMeasure *OutcomeMeasureClient
	// OutcomeMeasurement is the client for interacting with the OutcomeMeasurement builders.
	OutcomeMeasurement *OutcomeMeasurementClient
	// OutcomeMeasuresModule is the client for interacting with the OutcomeMeasuresModule builders.
	OutcomeMeasuresModule *OutcomeMeasuresModuleClient
	// OutcomeOverview is the client for interacting with the OutcomeOverview builders.
	OutcomeOverview *OutcomeOverviewClient
	// ParticipantFlowModule is the client for interacting with the ParticipantFlowModule builders.
	ParticipantFlowModule *ParticipantFlowModuleClient
	// PointOfContact is the client for interacting with the PointOfContact builders.
	PointOfContact *PointOfContactClient
	// ResultsDefinition is the client for interacting with the ResultsDefinition builders.
	ResultsDefinition *ResultsDefinitionClient
	// Schedule is the client for interacting with the Schedule builders.
	Schedule *ScheduleClient
	// SeriousEvent is the client for interacting with the SeriousEvent builders.
	SeriousEvent *SeriousEventClient
	// SeriousEventStats is the client for interacting with the SeriousEventStats builders.
	SeriousEventStats *SeriousEventStatsClient
	// StudyEligibility is the client for interacting with the StudyEligibility builders.
	StudyEligibility *StudyEligibilityClient
	// StudyLocation is the client for interacting with the StudyLocation builders.
	StudyLocation *StudyLocationClient
	// Task is the client for interacting with the Task builders.
	Task *TaskClient
	// UseCase is the client for interacting with the UseCase builders.
	UseCase *UseCaseClient
	// Vaccine is the client for interacting with the Vaccine builders.
	Vaccine *VaccineClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.AdverseEventsModule = NewAdverseEventsModuleClient(c.config)
	c.BaselineCategory = NewBaselineCategoryClient(c.config)
	c.BaselineCharacteristicsModule = NewBaselineCharacteristicsModuleClient(c.config)
	c.BaselineClass = NewBaselineClassClient(c.config)
	c.BaselineClassDenom = NewBaselineClassDenomClient(c.config)
	c.BaselineClassDenomCount = NewBaselineClassDenomCountClient(c.config)
	c.BaselineDenom = NewBaselineDenomClient(c.config)
	c.BaselineDenomCount = NewBaselineDenomCountClient(c.config)
	c.BaselineGroup = NewBaselineGroupClient(c.config)
	c.BaselineMeasure = NewBaselineMeasureClient(c.config)
	c.BaselineMeasureDenom = NewBaselineMeasureDenomClient(c.config)
	c.BaselineMeasureDenomCount = NewBaselineMeasureDenomCountClient(c.config)
	c.BaselineMeasurement = NewBaselineMeasurementClient(c.config)
	c.CertainAgreement = NewCertainAgreementClient(c.config)
	c.ClinicalTrial = NewClinicalTrialClient(c.config)
	c.ClinicalTrialMetadata = NewClinicalTrialMetadataClient(c.config)
	c.DoseDescription = NewDoseDescriptionClient(c.config)
	c.EventGroup = NewEventGroupClient(c.config)
	c.FlowAchievement = NewFlowAchievementClient(c.config)
	c.FlowDropWithdraw = NewFlowDropWithdrawClient(c.config)
	c.FlowGroup = NewFlowGroupClient(c.config)
	c.FlowMilestone = NewFlowMilestoneClient(c.config)
	c.FlowPeriod = NewFlowPeriodClient(c.config)
	c.FlowReason = NewFlowReasonClient(c.config)
	c.ImmunocompromisedGroups = NewImmunocompromisedGroupsClient(c.config)
	c.Manufacturer = NewManufacturerClient(c.config)
	c.MoreInfoModule = NewMoreInfoModuleClient(c.config)
	c.OtherEvent = NewOtherEventClient(c.config)
	c.OtherEventStats = NewOtherEventStatsClient(c.config)
	c.OutcomeAnalysis = NewOutcomeAnalysisClient(c.config)
	c.OutcomeAnalysisGroupID = NewOutcomeAnalysisGroupIDClient(c.config)
	c.OutcomeCategory = NewOutcomeCategoryClient(c.config)
	c.OutcomeClass = NewOutcomeClassClient(c.config)
	c.OutcomeClassDenom = NewOutcomeClassDenomClient(c.config)
	c.OutcomeClassDenomCount = NewOutcomeClassDenomCountClient(c.config)
	c.OutcomeDenom = NewOutcomeDenomClient(c.config)
	c.OutcomeDenomCount = NewOutcomeDenomCountClient(c.config)
	c.OutcomeGroup = NewOutcomeGroupClient(c.config)
	c.OutcomeMeasure = NewOutcomeMeasureClient(c.config)
	c.OutcomeMeasurement = NewOutcomeMeasurementClient(c.config)
	c.OutcomeMeasuresModule = NewOutcomeMeasuresModuleClient(c.config)
	c.OutcomeOverview = NewOutcomeOverviewClient(c.config)
	c.ParticipantFlowModule = NewParticipantFlowModuleClient(c.config)
	c.PointOfContact = NewPointOfContactClient(c.config)
	c.ResultsDefinition = NewResultsDefinitionClient(c.config)
	c.Schedule = NewScheduleClient(c.config)
	c.SeriousEvent = NewSeriousEventClient(c.config)
	c.SeriousEventStats = NewSeriousEventStatsClient(c.config)
	c.StudyEligibility = NewStudyEligibilityClient(c.config)
	c.StudyLocation = NewStudyLocationClient(c.config)
	c.Task = NewTaskClient(c.config)
	c.UseCase = NewUseCaseClient(c.config)
	c.Vaccine = NewVaccineClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("models: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("models: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                           ctx,
		config:                        cfg,
		AdverseEventsModule:           NewAdverseEventsModuleClient(cfg),
		BaselineCategory:              NewBaselineCategoryClient(cfg),
		BaselineCharacteristicsModule: NewBaselineCharacteristicsModuleClient(cfg),
		BaselineClass:                 NewBaselineClassClient(cfg),
		BaselineClassDenom:            NewBaselineClassDenomClient(cfg),
		BaselineClassDenomCount:       NewBaselineClassDenomCountClient(cfg),
		BaselineDenom:                 NewBaselineDenomClient(cfg),
		BaselineDenomCount:            NewBaselineDenomCountClient(cfg),
		BaselineGroup:                 NewBaselineGroupClient(cfg),
		BaselineMeasure:               NewBaselineMeasureClient(cfg),
		BaselineMeasureDenom:          NewBaselineMeasureDenomClient(cfg),
		BaselineMeasureDenomCount:     NewBaselineMeasureDenomCountClient(cfg),
		BaselineMeasurement:           NewBaselineMeasurementClient(cfg),
		CertainAgreement:              NewCertainAgreementClient(cfg),
		ClinicalTrial:                 NewClinicalTrialClient(cfg),
		ClinicalTrialMetadata:         NewClinicalTrialMetadataClient(cfg),
		DoseDescription:               NewDoseDescriptionClient(cfg),
		EventGroup:                    NewEventGroupClient(cfg),
		FlowAchievement:               NewFlowAchievementClient(cfg),
		FlowDropWithdraw:              NewFlowDropWithdrawClient(cfg),
		FlowGroup:                     NewFlowGroupClient(cfg),
		FlowMilestone:                 NewFlowMilestoneClient(cfg),
		FlowPeriod:                    NewFlowPeriodClient(cfg),
		FlowReason:                    NewFlowReasonClient(cfg),
		ImmunocompromisedGroups:       NewImmunocompromisedGroupsClient(cfg),
		Manufacturer:                  NewManufacturerClient(cfg),
		MoreInfoModule:                NewMoreInfoModuleClient(cfg),
		OtherEvent:                    NewOtherEventClient(cfg),
		OtherEventStats:               NewOtherEventStatsClient(cfg),
		OutcomeAnalysis:               NewOutcomeAnalysisClient(cfg),
		OutcomeAnalysisGroupID:        NewOutcomeAnalysisGroupIDClient(cfg),
		OutcomeCategory:               NewOutcomeCategoryClient(cfg),
		OutcomeClass:                  NewOutcomeClassClient(cfg),
		OutcomeClassDenom:             NewOutcomeClassDenomClient(cfg),
		OutcomeClassDenomCount:        NewOutcomeClassDenomCountClient(cfg),
		OutcomeDenom:                  NewOutcomeDenomClient(cfg),
		OutcomeDenomCount:             NewOutcomeDenomCountClient(cfg),
		OutcomeGroup:                  NewOutcomeGroupClient(cfg),
		OutcomeMeasure:                NewOutcomeMeasureClient(cfg),
		OutcomeMeasurement:            NewOutcomeMeasurementClient(cfg),
		OutcomeMeasuresModule:         NewOutcomeMeasuresModuleClient(cfg),
		OutcomeOverview:               NewOutcomeOverviewClient(cfg),
		ParticipantFlowModule:         NewParticipantFlowModuleClient(cfg),
		PointOfContact:                NewPointOfContactClient(cfg),
		ResultsDefinition:             NewResultsDefinitionClient(cfg),
		Schedule:                      NewScheduleClient(cfg),
		SeriousEvent:                  NewSeriousEventClient(cfg),
		SeriousEventStats:             NewSeriousEventStatsClient(cfg),
		StudyEligibility:              NewStudyEligibilityClient(cfg),
		StudyLocation:                 NewStudyLocationClient(cfg),
		Task:                          NewTaskClient(cfg),
		UseCase:                       NewUseCaseClient(cfg),
		Vaccine:                       NewVaccineClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                           ctx,
		config:                        cfg,
		AdverseEventsModule:           NewAdverseEventsModuleClient(cfg),
		BaselineCategory:              NewBaselineCategoryClient(cfg),
		BaselineCharacteristicsModule: NewBaselineCharacteristicsModuleClient(cfg),
		BaselineClass:                 NewBaselineClassClient(cfg),
		BaselineClassDenom:            NewBaselineClassDenomClient(cfg),
		BaselineClassDenomCount:       NewBaselineClassDenomCountClient(cfg),
		BaselineDenom:                 NewBaselineDenomClient(cfg),
		BaselineDenomCount:            NewBaselineDenomCountClient(cfg),
		BaselineGroup:                 NewBaselineGroupClient(cfg),
		BaselineMeasure:               NewBaselineMeasureClient(cfg),
		BaselineMeasureDenom:          NewBaselineMeasureDenomClient(cfg),
		BaselineMeasureDenomCount:     NewBaselineMeasureDenomCountClient(cfg),
		BaselineMeasurement:           NewBaselineMeasurementClient(cfg),
		CertainAgreement:              NewCertainAgreementClient(cfg),
		ClinicalTrial:                 NewClinicalTrialClient(cfg),
		ClinicalTrialMetadata:         NewClinicalTrialMetadataClient(cfg),
		DoseDescription:               NewDoseDescriptionClient(cfg),
		EventGroup:                    NewEventGroupClient(cfg),
		FlowAchievement:               NewFlowAchievementClient(cfg),
		FlowDropWithdraw:              NewFlowDropWithdrawClient(cfg),
		FlowGroup:                     NewFlowGroupClient(cfg),
		FlowMilestone:                 NewFlowMilestoneClient(cfg),
		FlowPeriod:                    NewFlowPeriodClient(cfg),
		FlowReason:                    NewFlowReasonClient(cfg),
		ImmunocompromisedGroups:       NewImmunocompromisedGroupsClient(cfg),
		Manufacturer:                  NewManufacturerClient(cfg),
		MoreInfoModule:                NewMoreInfoModuleClient(cfg),
		OtherEvent:                    NewOtherEventClient(cfg),
		OtherEventStats:               NewOtherEventStatsClient(cfg),
		OutcomeAnalysis:               NewOutcomeAnalysisClient(cfg),
		OutcomeAnalysisGroupID:        NewOutcomeAnalysisGroupIDClient(cfg),
		OutcomeCategory:               NewOutcomeCategoryClient(cfg),
		OutcomeClass:                  NewOutcomeClassClient(cfg),
		OutcomeClassDenom:             NewOutcomeClassDenomClient(cfg),
		OutcomeClassDenomCount:        NewOutcomeClassDenomCountClient(cfg),
		OutcomeDenom:                  NewOutcomeDenomClient(cfg),
		OutcomeDenomCount:             NewOutcomeDenomCountClient(cfg),
		OutcomeGroup:                  NewOutcomeGroupClient(cfg),
		OutcomeMeasure:                NewOutcomeMeasureClient(cfg),
		OutcomeMeasurement:            NewOutcomeMeasurementClient(cfg),
		OutcomeMeasuresModule:         NewOutcomeMeasuresModuleClient(cfg),
		OutcomeOverview:               NewOutcomeOverviewClient(cfg),
		ParticipantFlowModule:         NewParticipantFlowModuleClient(cfg),
		PointOfContact:                NewPointOfContactClient(cfg),
		ResultsDefinition:             NewResultsDefinitionClient(cfg),
		Schedule:                      NewScheduleClient(cfg),
		SeriousEvent:                  NewSeriousEventClient(cfg),
		SeriousEventStats:             NewSeriousEventStatsClient(cfg),
		StudyEligibility:              NewStudyEligibilityClient(cfg),
		StudyLocation:                 NewStudyLocationClient(cfg),
		Task:                          NewTaskClient(cfg),
		UseCase:                       NewUseCaseClient(cfg),
		Vaccine:                       NewVaccineClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AdverseEventsModule.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.AdverseEventsModule.Use(hooks...)
	c.BaselineCategory.Use(hooks...)
	c.BaselineCharacteristicsModule.Use(hooks...)
	c.BaselineClass.Use(hooks...)
	c.BaselineClassDenom.Use(hooks...)
	c.BaselineClassDenomCount.Use(hooks...)
	c.BaselineDenom.Use(hooks...)
	c.BaselineDenomCount.Use(hooks...)
	c.BaselineGroup.Use(hooks...)
	c.BaselineMeasure.Use(hooks...)
	c.BaselineMeasureDenom.Use(hooks...)
	c.BaselineMeasureDenomCount.Use(hooks...)
	c.BaselineMeasurement.Use(hooks...)
	c.CertainAgreement.Use(hooks...)
	c.ClinicalTrial.Use(hooks...)
	c.ClinicalTrialMetadata.Use(hooks...)
	c.DoseDescription.Use(hooks...)
	c.EventGroup.Use(hooks...)
	c.FlowAchievement.Use(hooks...)
	c.FlowDropWithdraw.Use(hooks...)
	c.FlowGroup.Use(hooks...)
	c.FlowMilestone.Use(hooks...)
	c.FlowPeriod.Use(hooks...)
	c.FlowReason.Use(hooks...)
	c.ImmunocompromisedGroups.Use(hooks...)
	c.Manufacturer.Use(hooks...)
	c.MoreInfoModule.Use(hooks...)
	c.OtherEvent.Use(hooks...)
	c.OtherEventStats.Use(hooks...)
	c.OutcomeAnalysis.Use(hooks...)
	c.OutcomeAnalysisGroupID.Use(hooks...)
	c.OutcomeCategory.Use(hooks...)
	c.OutcomeClass.Use(hooks...)
	c.OutcomeClassDenom.Use(hooks...)
	c.OutcomeClassDenomCount.Use(hooks...)
	c.OutcomeDenom.Use(hooks...)
	c.OutcomeDenomCount.Use(hooks...)
	c.OutcomeGroup.Use(hooks...)
	c.OutcomeMeasure.Use(hooks...)
	c.OutcomeMeasurement.Use(hooks...)
	c.OutcomeMeasuresModule.Use(hooks...)
	c.OutcomeOverview.Use(hooks...)
	c.ParticipantFlowModule.Use(hooks...)
	c.PointOfContact.Use(hooks...)
	c.ResultsDefinition.Use(hooks...)
	c.Schedule.Use(hooks...)
	c.SeriousEvent.Use(hooks...)
	c.SeriousEventStats.Use(hooks...)
	c.StudyEligibility.Use(hooks...)
	c.StudyLocation.Use(hooks...)
	c.Task.Use(hooks...)
	c.UseCase.Use(hooks...)
	c.Vaccine.Use(hooks...)
}

// AdverseEventsModuleClient is a client for the AdverseEventsModule schema.
type AdverseEventsModuleClient struct {
	config
}

// NewAdverseEventsModuleClient returns a client for the AdverseEventsModule from the given config.
func NewAdverseEventsModuleClient(c config) *AdverseEventsModuleClient {
	return &AdverseEventsModuleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `adverseeventsmodule.Hooks(f(g(h())))`.
func (c *AdverseEventsModuleClient) Use(hooks ...Hook) {
	c.hooks.AdverseEventsModule = append(c.hooks.AdverseEventsModule, hooks...)
}

// Create returns a create builder for AdverseEventsModule.
func (c *AdverseEventsModuleClient) Create() *AdverseEventsModuleCreate {
	mutation := newAdverseEventsModuleMutation(c.config, OpCreate)
	return &AdverseEventsModuleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AdverseEventsModule entities.
func (c *AdverseEventsModuleClient) CreateBulk(builders ...*AdverseEventsModuleCreate) *AdverseEventsModuleCreateBulk {
	return &AdverseEventsModuleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AdverseEventsModule.
func (c *AdverseEventsModuleClient) Update() *AdverseEventsModuleUpdate {
	mutation := newAdverseEventsModuleMutation(c.config, OpUpdate)
	return &AdverseEventsModuleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AdverseEventsModuleClient) UpdateOne(aem *AdverseEventsModule) *AdverseEventsModuleUpdateOne {
	mutation := newAdverseEventsModuleMutation(c.config, OpUpdateOne, withAdverseEventsModule(aem))
	return &AdverseEventsModuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AdverseEventsModuleClient) UpdateOneID(id int) *AdverseEventsModuleUpdateOne {
	mutation := newAdverseEventsModuleMutation(c.config, OpUpdateOne, withAdverseEventsModuleID(id))
	return &AdverseEventsModuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AdverseEventsModule.
func (c *AdverseEventsModuleClient) Delete() *AdverseEventsModuleDelete {
	mutation := newAdverseEventsModuleMutation(c.config, OpDelete)
	return &AdverseEventsModuleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AdverseEventsModuleClient) DeleteOne(aem *AdverseEventsModule) *AdverseEventsModuleDeleteOne {
	return c.DeleteOneID(aem.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AdverseEventsModuleClient) DeleteOneID(id int) *AdverseEventsModuleDeleteOne {
	builder := c.Delete().Where(adverseeventsmodule.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AdverseEventsModuleDeleteOne{builder}
}

// Query returns a query builder for AdverseEventsModule.
func (c *AdverseEventsModuleClient) Query() *AdverseEventsModuleQuery {
	return &AdverseEventsModuleQuery{
		config: c.config,
	}
}

// Get returns a AdverseEventsModule entity by its id.
func (c *AdverseEventsModuleClient) Get(ctx context.Context, id int) (*AdverseEventsModule, error) {
	return c.Query().Where(adverseeventsmodule.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AdverseEventsModuleClient) GetX(ctx context.Context, id int) *AdverseEventsModule {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a AdverseEventsModule.
func (c *AdverseEventsModuleClient) QueryParent(aem *AdverseEventsModule) *ResultsDefinitionQuery {
	query := &ResultsDefinitionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := aem.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(adverseeventsmodule.Table, adverseeventsmodule.FieldID, id),
			sqlgraph.To(resultsdefinition.Table, resultsdefinition.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, adverseeventsmodule.ParentTable, adverseeventsmodule.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(aem.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEventGroupList queries the event_group_list edge of a AdverseEventsModule.
func (c *AdverseEventsModuleClient) QueryEventGroupList(aem *AdverseEventsModule) *EventGroupQuery {
	query := &EventGroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := aem.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(adverseeventsmodule.Table, adverseeventsmodule.FieldID, id),
			sqlgraph.To(eventgroup.Table, eventgroup.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, adverseeventsmodule.EventGroupListTable, adverseeventsmodule.EventGroupListColumn),
		)
		fromV = sqlgraph.Neighbors(aem.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySeriousEventList queries the serious_event_list edge of a AdverseEventsModule.
func (c *AdverseEventsModuleClient) QuerySeriousEventList(aem *AdverseEventsModule) *SeriousEventQuery {
	query := &SeriousEventQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := aem.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(adverseeventsmodule.Table, adverseeventsmodule.FieldID, id),
			sqlgraph.To(seriousevent.Table, seriousevent.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, adverseeventsmodule.SeriousEventListTable, adverseeventsmodule.SeriousEventListColumn),
		)
		fromV = sqlgraph.Neighbors(aem.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOtherEventList queries the other_event_list edge of a AdverseEventsModule.
func (c *AdverseEventsModuleClient) QueryOtherEventList(aem *AdverseEventsModule) *OtherEventQuery {
	query := &OtherEventQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := aem.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(adverseeventsmodule.Table, adverseeventsmodule.FieldID, id),
			sqlgraph.To(otherevent.Table, otherevent.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, adverseeventsmodule.OtherEventListTable, adverseeventsmodule.OtherEventListColumn),
		)
		fromV = sqlgraph.Neighbors(aem.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AdverseEventsModuleClient) Hooks() []Hook {
	return c.hooks.AdverseEventsModule
}

// BaselineCategoryClient is a client for the BaselineCategory schema.
type BaselineCategoryClient struct {
	config
}

// NewBaselineCategoryClient returns a client for the BaselineCategory from the given config.
func NewBaselineCategoryClient(c config) *BaselineCategoryClient {
	return &BaselineCategoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `baselinecategory.Hooks(f(g(h())))`.
func (c *BaselineCategoryClient) Use(hooks ...Hook) {
	c.hooks.BaselineCategory = append(c.hooks.BaselineCategory, hooks...)
}

// Create returns a create builder for BaselineCategory.
func (c *BaselineCategoryClient) Create() *BaselineCategoryCreate {
	mutation := newBaselineCategoryMutation(c.config, OpCreate)
	return &BaselineCategoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaselineCategory entities.
func (c *BaselineCategoryClient) CreateBulk(builders ...*BaselineCategoryCreate) *BaselineCategoryCreateBulk {
	return &BaselineCategoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaselineCategory.
func (c *BaselineCategoryClient) Update() *BaselineCategoryUpdate {
	mutation := newBaselineCategoryMutation(c.config, OpUpdate)
	return &BaselineCategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaselineCategoryClient) UpdateOne(bc *BaselineCategory) *BaselineCategoryUpdateOne {
	mutation := newBaselineCategoryMutation(c.config, OpUpdateOne, withBaselineCategory(bc))
	return &BaselineCategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaselineCategoryClient) UpdateOneID(id int) *BaselineCategoryUpdateOne {
	mutation := newBaselineCategoryMutation(c.config, OpUpdateOne, withBaselineCategoryID(id))
	return &BaselineCategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaselineCategory.
func (c *BaselineCategoryClient) Delete() *BaselineCategoryDelete {
	mutation := newBaselineCategoryMutation(c.config, OpDelete)
	return &BaselineCategoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BaselineCategoryClient) DeleteOne(bc *BaselineCategory) *BaselineCategoryDeleteOne {
	return c.DeleteOneID(bc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BaselineCategoryClient) DeleteOneID(id int) *BaselineCategoryDeleteOne {
	builder := c.Delete().Where(baselinecategory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaselineCategoryDeleteOne{builder}
}

// Query returns a query builder for BaselineCategory.
func (c *BaselineCategoryClient) Query() *BaselineCategoryQuery {
	return &BaselineCategoryQuery{
		config: c.config,
	}
}

// Get returns a BaselineCategory entity by its id.
func (c *BaselineCategoryClient) Get(ctx context.Context, id int) (*BaselineCategory, error) {
	return c.Query().Where(baselinecategory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaselineCategoryClient) GetX(ctx context.Context, id int) *BaselineCategory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a BaselineCategory.
func (c *BaselineCategoryClient) QueryParent(bc *BaselineCategory) *BaselineClassQuery {
	query := &BaselineClassQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinecategory.Table, baselinecategory.FieldID, id),
			sqlgraph.To(baselineclass.Table, baselineclass.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinecategory.ParentTable, baselinecategory.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(bc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBaselineMeasurementList queries the baseline_measurement_list edge of a BaselineCategory.
func (c *BaselineCategoryClient) QueryBaselineMeasurementList(bc *BaselineCategory) *BaselineMeasurementQuery {
	query := &BaselineMeasurementQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinecategory.Table, baselinecategory.FieldID, id),
			sqlgraph.To(baselinemeasurement.Table, baselinemeasurement.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinecategory.BaselineMeasurementListTable, baselinecategory.BaselineMeasurementListColumn),
		)
		fromV = sqlgraph.Neighbors(bc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BaselineCategoryClient) Hooks() []Hook {
	return c.hooks.BaselineCategory
}

// BaselineCharacteristicsModuleClient is a client for the BaselineCharacteristicsModule schema.
type BaselineCharacteristicsModuleClient struct {
	config
}

// NewBaselineCharacteristicsModuleClient returns a client for the BaselineCharacteristicsModule from the given config.
func NewBaselineCharacteristicsModuleClient(c config) *BaselineCharacteristicsModuleClient {
	return &BaselineCharacteristicsModuleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `baselinecharacteristicsmodule.Hooks(f(g(h())))`.
func (c *BaselineCharacteristicsModuleClient) Use(hooks ...Hook) {
	c.hooks.BaselineCharacteristicsModule = append(c.hooks.BaselineCharacteristicsModule, hooks...)
}

// Create returns a create builder for BaselineCharacteristicsModule.
func (c *BaselineCharacteristicsModuleClient) Create() *BaselineCharacteristicsModuleCreate {
	mutation := newBaselineCharacteristicsModuleMutation(c.config, OpCreate)
	return &BaselineCharacteristicsModuleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaselineCharacteristicsModule entities.
func (c *BaselineCharacteristicsModuleClient) CreateBulk(builders ...*BaselineCharacteristicsModuleCreate) *BaselineCharacteristicsModuleCreateBulk {
	return &BaselineCharacteristicsModuleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaselineCharacteristicsModule.
func (c *BaselineCharacteristicsModuleClient) Update() *BaselineCharacteristicsModuleUpdate {
	mutation := newBaselineCharacteristicsModuleMutation(c.config, OpUpdate)
	return &BaselineCharacteristicsModuleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaselineCharacteristicsModuleClient) UpdateOne(bcm *BaselineCharacteristicsModule) *BaselineCharacteristicsModuleUpdateOne {
	mutation := newBaselineCharacteristicsModuleMutation(c.config, OpUpdateOne, withBaselineCharacteristicsModule(bcm))
	return &BaselineCharacteristicsModuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaselineCharacteristicsModuleClient) UpdateOneID(id int) *BaselineCharacteristicsModuleUpdateOne {
	mutation := newBaselineCharacteristicsModuleMutation(c.config, OpUpdateOne, withBaselineCharacteristicsModuleID(id))
	return &BaselineCharacteristicsModuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaselineCharacteristicsModule.
func (c *BaselineCharacteristicsModuleClient) Delete() *BaselineCharacteristicsModuleDelete {
	mutation := newBaselineCharacteristicsModuleMutation(c.config, OpDelete)
	return &BaselineCharacteristicsModuleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BaselineCharacteristicsModuleClient) DeleteOne(bcm *BaselineCharacteristicsModule) *BaselineCharacteristicsModuleDeleteOne {
	return c.DeleteOneID(bcm.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BaselineCharacteristicsModuleClient) DeleteOneID(id int) *BaselineCharacteristicsModuleDeleteOne {
	builder := c.Delete().Where(baselinecharacteristicsmodule.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaselineCharacteristicsModuleDeleteOne{builder}
}

// Query returns a query builder for BaselineCharacteristicsModule.
func (c *BaselineCharacteristicsModuleClient) Query() *BaselineCharacteristicsModuleQuery {
	return &BaselineCharacteristicsModuleQuery{
		config: c.config,
	}
}

// Get returns a BaselineCharacteristicsModule entity by its id.
func (c *BaselineCharacteristicsModuleClient) Get(ctx context.Context, id int) (*BaselineCharacteristicsModule, error) {
	return c.Query().Where(baselinecharacteristicsmodule.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaselineCharacteristicsModuleClient) GetX(ctx context.Context, id int) *BaselineCharacteristicsModule {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a BaselineCharacteristicsModule.
func (c *BaselineCharacteristicsModuleClient) QueryParent(bcm *BaselineCharacteristicsModule) *ResultsDefinitionQuery {
	query := &ResultsDefinitionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bcm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID, id),
			sqlgraph.To(resultsdefinition.Table, resultsdefinition.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, baselinecharacteristicsmodule.ParentTable, baselinecharacteristicsmodule.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(bcm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBaselineGroupList queries the baseline_group_list edge of a BaselineCharacteristicsModule.
func (c *BaselineCharacteristicsModuleClient) QueryBaselineGroupList(bcm *BaselineCharacteristicsModule) *BaselineGroupQuery {
	query := &BaselineGroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bcm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID, id),
			sqlgraph.To(baselinegroup.Table, baselinegroup.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinecharacteristicsmodule.BaselineGroupListTable, baselinecharacteristicsmodule.BaselineGroupListColumn),
		)
		fromV = sqlgraph.Neighbors(bcm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBaselineDenomList queries the baseline_denom_list edge of a BaselineCharacteristicsModule.
func (c *BaselineCharacteristicsModuleClient) QueryBaselineDenomList(bcm *BaselineCharacteristicsModule) *BaselineDenomQuery {
	query := &BaselineDenomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bcm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID, id),
			sqlgraph.To(baselinedenom.Table, baselinedenom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinecharacteristicsmodule.BaselineDenomListTable, baselinecharacteristicsmodule.BaselineDenomListColumn),
		)
		fromV = sqlgraph.Neighbors(bcm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBaselineMeasureList queries the baseline_measure_list edge of a BaselineCharacteristicsModule.
func (c *BaselineCharacteristicsModuleClient) QueryBaselineMeasureList(bcm *BaselineCharacteristicsModule) *BaselineMeasureQuery {
	query := &BaselineMeasureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bcm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID, id),
			sqlgraph.To(baselinemeasure.Table, baselinemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinecharacteristicsmodule.BaselineMeasureListTable, baselinecharacteristicsmodule.BaselineMeasureListColumn),
		)
		fromV = sqlgraph.Neighbors(bcm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BaselineCharacteristicsModuleClient) Hooks() []Hook {
	return c.hooks.BaselineCharacteristicsModule
}

// BaselineClassClient is a client for the BaselineClass schema.
type BaselineClassClient struct {
	config
}

// NewBaselineClassClient returns a client for the BaselineClass from the given config.
func NewBaselineClassClient(c config) *BaselineClassClient {
	return &BaselineClassClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `baselineclass.Hooks(f(g(h())))`.
func (c *BaselineClassClient) Use(hooks ...Hook) {
	c.hooks.BaselineClass = append(c.hooks.BaselineClass, hooks...)
}

// Create returns a create builder for BaselineClass.
func (c *BaselineClassClient) Create() *BaselineClassCreate {
	mutation := newBaselineClassMutation(c.config, OpCreate)
	return &BaselineClassCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaselineClass entities.
func (c *BaselineClassClient) CreateBulk(builders ...*BaselineClassCreate) *BaselineClassCreateBulk {
	return &BaselineClassCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaselineClass.
func (c *BaselineClassClient) Update() *BaselineClassUpdate {
	mutation := newBaselineClassMutation(c.config, OpUpdate)
	return &BaselineClassUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaselineClassClient) UpdateOne(bc *BaselineClass) *BaselineClassUpdateOne {
	mutation := newBaselineClassMutation(c.config, OpUpdateOne, withBaselineClass(bc))
	return &BaselineClassUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaselineClassClient) UpdateOneID(id int) *BaselineClassUpdateOne {
	mutation := newBaselineClassMutation(c.config, OpUpdateOne, withBaselineClassID(id))
	return &BaselineClassUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaselineClass.
func (c *BaselineClassClient) Delete() *BaselineClassDelete {
	mutation := newBaselineClassMutation(c.config, OpDelete)
	return &BaselineClassDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BaselineClassClient) DeleteOne(bc *BaselineClass) *BaselineClassDeleteOne {
	return c.DeleteOneID(bc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BaselineClassClient) DeleteOneID(id int) *BaselineClassDeleteOne {
	builder := c.Delete().Where(baselineclass.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaselineClassDeleteOne{builder}
}

// Query returns a query builder for BaselineClass.
func (c *BaselineClassClient) Query() *BaselineClassQuery {
	return &BaselineClassQuery{
		config: c.config,
	}
}

// Get returns a BaselineClass entity by its id.
func (c *BaselineClassClient) Get(ctx context.Context, id int) (*BaselineClass, error) {
	return c.Query().Where(baselineclass.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaselineClassClient) GetX(ctx context.Context, id int) *BaselineClass {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a BaselineClass.
func (c *BaselineClassClient) QueryParent(bc *BaselineClass) *BaselineMeasureQuery {
	query := &BaselineMeasureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselineclass.Table, baselineclass.FieldID, id),
			sqlgraph.To(baselinemeasure.Table, baselinemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselineclass.ParentTable, baselineclass.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(bc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBaselineClassDenomList queries the baseline_class_denom_list edge of a BaselineClass.
func (c *BaselineClassClient) QueryBaselineClassDenomList(bc *BaselineClass) *BaselineClassDenomQuery {
	query := &BaselineClassDenomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselineclass.Table, baselineclass.FieldID, id),
			sqlgraph.To(baselineclassdenom.Table, baselineclassdenom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselineclass.BaselineClassDenomListTable, baselineclass.BaselineClassDenomListColumn),
		)
		fromV = sqlgraph.Neighbors(bc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBaselineCategoryList queries the baseline_category_list edge of a BaselineClass.
func (c *BaselineClassClient) QueryBaselineCategoryList(bc *BaselineClass) *BaselineCategoryQuery {
	query := &BaselineCategoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselineclass.Table, baselineclass.FieldID, id),
			sqlgraph.To(baselinecategory.Table, baselinecategory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselineclass.BaselineCategoryListTable, baselineclass.BaselineCategoryListColumn),
		)
		fromV = sqlgraph.Neighbors(bc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BaselineClassClient) Hooks() []Hook {
	return c.hooks.BaselineClass
}

// BaselineClassDenomClient is a client for the BaselineClassDenom schema.
type BaselineClassDenomClient struct {
	config
}

// NewBaselineClassDenomClient returns a client for the BaselineClassDenom from the given config.
func NewBaselineClassDenomClient(c config) *BaselineClassDenomClient {
	return &BaselineClassDenomClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `baselineclassdenom.Hooks(f(g(h())))`.
func (c *BaselineClassDenomClient) Use(hooks ...Hook) {
	c.hooks.BaselineClassDenom = append(c.hooks.BaselineClassDenom, hooks...)
}

// Create returns a create builder for BaselineClassDenom.
func (c *BaselineClassDenomClient) Create() *BaselineClassDenomCreate {
	mutation := newBaselineClassDenomMutation(c.config, OpCreate)
	return &BaselineClassDenomCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaselineClassDenom entities.
func (c *BaselineClassDenomClient) CreateBulk(builders ...*BaselineClassDenomCreate) *BaselineClassDenomCreateBulk {
	return &BaselineClassDenomCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaselineClassDenom.
func (c *BaselineClassDenomClient) Update() *BaselineClassDenomUpdate {
	mutation := newBaselineClassDenomMutation(c.config, OpUpdate)
	return &BaselineClassDenomUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaselineClassDenomClient) UpdateOne(bcd *BaselineClassDenom) *BaselineClassDenomUpdateOne {
	mutation := newBaselineClassDenomMutation(c.config, OpUpdateOne, withBaselineClassDenom(bcd))
	return &BaselineClassDenomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaselineClassDenomClient) UpdateOneID(id int) *BaselineClassDenomUpdateOne {
	mutation := newBaselineClassDenomMutation(c.config, OpUpdateOne, withBaselineClassDenomID(id))
	return &BaselineClassDenomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaselineClassDenom.
func (c *BaselineClassDenomClient) Delete() *BaselineClassDenomDelete {
	mutation := newBaselineClassDenomMutation(c.config, OpDelete)
	return &BaselineClassDenomDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BaselineClassDenomClient) DeleteOne(bcd *BaselineClassDenom) *BaselineClassDenomDeleteOne {
	return c.DeleteOneID(bcd.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BaselineClassDenomClient) DeleteOneID(id int) *BaselineClassDenomDeleteOne {
	builder := c.Delete().Where(baselineclassdenom.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaselineClassDenomDeleteOne{builder}
}

// Query returns a query builder for BaselineClassDenom.
func (c *BaselineClassDenomClient) Query() *BaselineClassDenomQuery {
	return &BaselineClassDenomQuery{
		config: c.config,
	}
}

// Get returns a BaselineClassDenom entity by its id.
func (c *BaselineClassDenomClient) Get(ctx context.Context, id int) (*BaselineClassDenom, error) {
	return c.Query().Where(baselineclassdenom.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaselineClassDenomClient) GetX(ctx context.Context, id int) *BaselineClassDenom {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a BaselineClassDenom.
func (c *BaselineClassDenomClient) QueryParent(bcd *BaselineClassDenom) *BaselineClassQuery {
	query := &BaselineClassQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bcd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselineclassdenom.Table, baselineclassdenom.FieldID, id),
			sqlgraph.To(baselineclass.Table, baselineclass.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselineclassdenom.ParentTable, baselineclassdenom.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(bcd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBaselineClassDenomCountList queries the baseline_class_denom_count_list edge of a BaselineClassDenom.
func (c *BaselineClassDenomClient) QueryBaselineClassDenomCountList(bcd *BaselineClassDenom) *BaselineClassDenomCountQuery {
	query := &BaselineClassDenomCountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bcd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselineclassdenom.Table, baselineclassdenom.FieldID, id),
			sqlgraph.To(baselineclassdenomcount.Table, baselineclassdenomcount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselineclassdenom.BaselineClassDenomCountListTable, baselineclassdenom.BaselineClassDenomCountListColumn),
		)
		fromV = sqlgraph.Neighbors(bcd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BaselineClassDenomClient) Hooks() []Hook {
	return c.hooks.BaselineClassDenom
}

// BaselineClassDenomCountClient is a client for the BaselineClassDenomCount schema.
type BaselineClassDenomCountClient struct {
	config
}

// NewBaselineClassDenomCountClient returns a client for the BaselineClassDenomCount from the given config.
func NewBaselineClassDenomCountClient(c config) *BaselineClassDenomCountClient {
	return &BaselineClassDenomCountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `baselineclassdenomcount.Hooks(f(g(h())))`.
func (c *BaselineClassDenomCountClient) Use(hooks ...Hook) {
	c.hooks.BaselineClassDenomCount = append(c.hooks.BaselineClassDenomCount, hooks...)
}

// Create returns a create builder for BaselineClassDenomCount.
func (c *BaselineClassDenomCountClient) Create() *BaselineClassDenomCountCreate {
	mutation := newBaselineClassDenomCountMutation(c.config, OpCreate)
	return &BaselineClassDenomCountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaselineClassDenomCount entities.
func (c *BaselineClassDenomCountClient) CreateBulk(builders ...*BaselineClassDenomCountCreate) *BaselineClassDenomCountCreateBulk {
	return &BaselineClassDenomCountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaselineClassDenomCount.
func (c *BaselineClassDenomCountClient) Update() *BaselineClassDenomCountUpdate {
	mutation := newBaselineClassDenomCountMutation(c.config, OpUpdate)
	return &BaselineClassDenomCountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaselineClassDenomCountClient) UpdateOne(bcdc *BaselineClassDenomCount) *BaselineClassDenomCountUpdateOne {
	mutation := newBaselineClassDenomCountMutation(c.config, OpUpdateOne, withBaselineClassDenomCount(bcdc))
	return &BaselineClassDenomCountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaselineClassDenomCountClient) UpdateOneID(id int) *BaselineClassDenomCountUpdateOne {
	mutation := newBaselineClassDenomCountMutation(c.config, OpUpdateOne, withBaselineClassDenomCountID(id))
	return &BaselineClassDenomCountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaselineClassDenomCount.
func (c *BaselineClassDenomCountClient) Delete() *BaselineClassDenomCountDelete {
	mutation := newBaselineClassDenomCountMutation(c.config, OpDelete)
	return &BaselineClassDenomCountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BaselineClassDenomCountClient) DeleteOne(bcdc *BaselineClassDenomCount) *BaselineClassDenomCountDeleteOne {
	return c.DeleteOneID(bcdc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BaselineClassDenomCountClient) DeleteOneID(id int) *BaselineClassDenomCountDeleteOne {
	builder := c.Delete().Where(baselineclassdenomcount.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaselineClassDenomCountDeleteOne{builder}
}

// Query returns a query builder for BaselineClassDenomCount.
func (c *BaselineClassDenomCountClient) Query() *BaselineClassDenomCountQuery {
	return &BaselineClassDenomCountQuery{
		config: c.config,
	}
}

// Get returns a BaselineClassDenomCount entity by its id.
func (c *BaselineClassDenomCountClient) Get(ctx context.Context, id int) (*BaselineClassDenomCount, error) {
	return c.Query().Where(baselineclassdenomcount.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaselineClassDenomCountClient) GetX(ctx context.Context, id int) *BaselineClassDenomCount {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a BaselineClassDenomCount.
func (c *BaselineClassDenomCountClient) QueryParent(bcdc *BaselineClassDenomCount) *BaselineClassDenomQuery {
	query := &BaselineClassDenomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bcdc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselineclassdenomcount.Table, baselineclassdenomcount.FieldID, id),
			sqlgraph.To(baselineclassdenom.Table, baselineclassdenom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselineclassdenomcount.ParentTable, baselineclassdenomcount.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(bcdc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BaselineClassDenomCountClient) Hooks() []Hook {
	return c.hooks.BaselineClassDenomCount
}

// BaselineDenomClient is a client for the BaselineDenom schema.
type BaselineDenomClient struct {
	config
}

// NewBaselineDenomClient returns a client for the BaselineDenom from the given config.
func NewBaselineDenomClient(c config) *BaselineDenomClient {
	return &BaselineDenomClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `baselinedenom.Hooks(f(g(h())))`.
func (c *BaselineDenomClient) Use(hooks ...Hook) {
	c.hooks.BaselineDenom = append(c.hooks.BaselineDenom, hooks...)
}

// Create returns a create builder for BaselineDenom.
func (c *BaselineDenomClient) Create() *BaselineDenomCreate {
	mutation := newBaselineDenomMutation(c.config, OpCreate)
	return &BaselineDenomCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaselineDenom entities.
func (c *BaselineDenomClient) CreateBulk(builders ...*BaselineDenomCreate) *BaselineDenomCreateBulk {
	return &BaselineDenomCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaselineDenom.
func (c *BaselineDenomClient) Update() *BaselineDenomUpdate {
	mutation := newBaselineDenomMutation(c.config, OpUpdate)
	return &BaselineDenomUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaselineDenomClient) UpdateOne(bd *BaselineDenom) *BaselineDenomUpdateOne {
	mutation := newBaselineDenomMutation(c.config, OpUpdateOne, withBaselineDenom(bd))
	return &BaselineDenomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaselineDenomClient) UpdateOneID(id int) *BaselineDenomUpdateOne {
	mutation := newBaselineDenomMutation(c.config, OpUpdateOne, withBaselineDenomID(id))
	return &BaselineDenomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaselineDenom.
func (c *BaselineDenomClient) Delete() *BaselineDenomDelete {
	mutation := newBaselineDenomMutation(c.config, OpDelete)
	return &BaselineDenomDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BaselineDenomClient) DeleteOne(bd *BaselineDenom) *BaselineDenomDeleteOne {
	return c.DeleteOneID(bd.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BaselineDenomClient) DeleteOneID(id int) *BaselineDenomDeleteOne {
	builder := c.Delete().Where(baselinedenom.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaselineDenomDeleteOne{builder}
}

// Query returns a query builder for BaselineDenom.
func (c *BaselineDenomClient) Query() *BaselineDenomQuery {
	return &BaselineDenomQuery{
		config: c.config,
	}
}

// Get returns a BaselineDenom entity by its id.
func (c *BaselineDenomClient) Get(ctx context.Context, id int) (*BaselineDenom, error) {
	return c.Query().Where(baselinedenom.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaselineDenomClient) GetX(ctx context.Context, id int) *BaselineDenom {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a BaselineDenom.
func (c *BaselineDenomClient) QueryParent(bd *BaselineDenom) *BaselineCharacteristicsModuleQuery {
	query := &BaselineCharacteristicsModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinedenom.Table, baselinedenom.FieldID, id),
			sqlgraph.To(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinedenom.ParentTable, baselinedenom.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(bd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBaselineDenomCountList queries the baseline_denom_count_list edge of a BaselineDenom.
func (c *BaselineDenomClient) QueryBaselineDenomCountList(bd *BaselineDenom) *BaselineDenomCountQuery {
	query := &BaselineDenomCountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinedenom.Table, baselinedenom.FieldID, id),
			sqlgraph.To(baselinedenomcount.Table, baselinedenomcount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinedenom.BaselineDenomCountListTable, baselinedenom.BaselineDenomCountListColumn),
		)
		fromV = sqlgraph.Neighbors(bd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BaselineDenomClient) Hooks() []Hook {
	return c.hooks.BaselineDenom
}

// BaselineDenomCountClient is a client for the BaselineDenomCount schema.
type BaselineDenomCountClient struct {
	config
}

// NewBaselineDenomCountClient returns a client for the BaselineDenomCount from the given config.
func NewBaselineDenomCountClient(c config) *BaselineDenomCountClient {
	return &BaselineDenomCountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `baselinedenomcount.Hooks(f(g(h())))`.
func (c *BaselineDenomCountClient) Use(hooks ...Hook) {
	c.hooks.BaselineDenomCount = append(c.hooks.BaselineDenomCount, hooks...)
}

// Create returns a create builder for BaselineDenomCount.
func (c *BaselineDenomCountClient) Create() *BaselineDenomCountCreate {
	mutation := newBaselineDenomCountMutation(c.config, OpCreate)
	return &BaselineDenomCountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaselineDenomCount entities.
func (c *BaselineDenomCountClient) CreateBulk(builders ...*BaselineDenomCountCreate) *BaselineDenomCountCreateBulk {
	return &BaselineDenomCountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaselineDenomCount.
func (c *BaselineDenomCountClient) Update() *BaselineDenomCountUpdate {
	mutation := newBaselineDenomCountMutation(c.config, OpUpdate)
	return &BaselineDenomCountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaselineDenomCountClient) UpdateOne(bdc *BaselineDenomCount) *BaselineDenomCountUpdateOne {
	mutation := newBaselineDenomCountMutation(c.config, OpUpdateOne, withBaselineDenomCount(bdc))
	return &BaselineDenomCountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaselineDenomCountClient) UpdateOneID(id int) *BaselineDenomCountUpdateOne {
	mutation := newBaselineDenomCountMutation(c.config, OpUpdateOne, withBaselineDenomCountID(id))
	return &BaselineDenomCountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaselineDenomCount.
func (c *BaselineDenomCountClient) Delete() *BaselineDenomCountDelete {
	mutation := newBaselineDenomCountMutation(c.config, OpDelete)
	return &BaselineDenomCountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BaselineDenomCountClient) DeleteOne(bdc *BaselineDenomCount) *BaselineDenomCountDeleteOne {
	return c.DeleteOneID(bdc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BaselineDenomCountClient) DeleteOneID(id int) *BaselineDenomCountDeleteOne {
	builder := c.Delete().Where(baselinedenomcount.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaselineDenomCountDeleteOne{builder}
}

// Query returns a query builder for BaselineDenomCount.
func (c *BaselineDenomCountClient) Query() *BaselineDenomCountQuery {
	return &BaselineDenomCountQuery{
		config: c.config,
	}
}

// Get returns a BaselineDenomCount entity by its id.
func (c *BaselineDenomCountClient) Get(ctx context.Context, id int) (*BaselineDenomCount, error) {
	return c.Query().Where(baselinedenomcount.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaselineDenomCountClient) GetX(ctx context.Context, id int) *BaselineDenomCount {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a BaselineDenomCount.
func (c *BaselineDenomCountClient) QueryParent(bdc *BaselineDenomCount) *BaselineDenomQuery {
	query := &BaselineDenomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bdc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinedenomcount.Table, baselinedenomcount.FieldID, id),
			sqlgraph.To(baselinedenom.Table, baselinedenom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinedenomcount.ParentTable, baselinedenomcount.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(bdc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BaselineDenomCountClient) Hooks() []Hook {
	return c.hooks.BaselineDenomCount
}

// BaselineGroupClient is a client for the BaselineGroup schema.
type BaselineGroupClient struct {
	config
}

// NewBaselineGroupClient returns a client for the BaselineGroup from the given config.
func NewBaselineGroupClient(c config) *BaselineGroupClient {
	return &BaselineGroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `baselinegroup.Hooks(f(g(h())))`.
func (c *BaselineGroupClient) Use(hooks ...Hook) {
	c.hooks.BaselineGroup = append(c.hooks.BaselineGroup, hooks...)
}

// Create returns a create builder for BaselineGroup.
func (c *BaselineGroupClient) Create() *BaselineGroupCreate {
	mutation := newBaselineGroupMutation(c.config, OpCreate)
	return &BaselineGroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaselineGroup entities.
func (c *BaselineGroupClient) CreateBulk(builders ...*BaselineGroupCreate) *BaselineGroupCreateBulk {
	return &BaselineGroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaselineGroup.
func (c *BaselineGroupClient) Update() *BaselineGroupUpdate {
	mutation := newBaselineGroupMutation(c.config, OpUpdate)
	return &BaselineGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaselineGroupClient) UpdateOne(bg *BaselineGroup) *BaselineGroupUpdateOne {
	mutation := newBaselineGroupMutation(c.config, OpUpdateOne, withBaselineGroup(bg))
	return &BaselineGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaselineGroupClient) UpdateOneID(id int) *BaselineGroupUpdateOne {
	mutation := newBaselineGroupMutation(c.config, OpUpdateOne, withBaselineGroupID(id))
	return &BaselineGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaselineGroup.
func (c *BaselineGroupClient) Delete() *BaselineGroupDelete {
	mutation := newBaselineGroupMutation(c.config, OpDelete)
	return &BaselineGroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BaselineGroupClient) DeleteOne(bg *BaselineGroup) *BaselineGroupDeleteOne {
	return c.DeleteOneID(bg.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BaselineGroupClient) DeleteOneID(id int) *BaselineGroupDeleteOne {
	builder := c.Delete().Where(baselinegroup.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaselineGroupDeleteOne{builder}
}

// Query returns a query builder for BaselineGroup.
func (c *BaselineGroupClient) Query() *BaselineGroupQuery {
	return &BaselineGroupQuery{
		config: c.config,
	}
}

// Get returns a BaselineGroup entity by its id.
func (c *BaselineGroupClient) Get(ctx context.Context, id int) (*BaselineGroup, error) {
	return c.Query().Where(baselinegroup.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaselineGroupClient) GetX(ctx context.Context, id int) *BaselineGroup {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a BaselineGroup.
func (c *BaselineGroupClient) QueryParent(bg *BaselineGroup) *BaselineCharacteristicsModuleQuery {
	query := &BaselineCharacteristicsModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bg.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinegroup.Table, baselinegroup.FieldID, id),
			sqlgraph.To(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinegroup.ParentTable, baselinegroup.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(bg.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BaselineGroupClient) Hooks() []Hook {
	return c.hooks.BaselineGroup
}

// BaselineMeasureClient is a client for the BaselineMeasure schema.
type BaselineMeasureClient struct {
	config
}

// NewBaselineMeasureClient returns a client for the BaselineMeasure from the given config.
func NewBaselineMeasureClient(c config) *BaselineMeasureClient {
	return &BaselineMeasureClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `baselinemeasure.Hooks(f(g(h())))`.
func (c *BaselineMeasureClient) Use(hooks ...Hook) {
	c.hooks.BaselineMeasure = append(c.hooks.BaselineMeasure, hooks...)
}

// Create returns a create builder for BaselineMeasure.
func (c *BaselineMeasureClient) Create() *BaselineMeasureCreate {
	mutation := newBaselineMeasureMutation(c.config, OpCreate)
	return &BaselineMeasureCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaselineMeasure entities.
func (c *BaselineMeasureClient) CreateBulk(builders ...*BaselineMeasureCreate) *BaselineMeasureCreateBulk {
	return &BaselineMeasureCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaselineMeasure.
func (c *BaselineMeasureClient) Update() *BaselineMeasureUpdate {
	mutation := newBaselineMeasureMutation(c.config, OpUpdate)
	return &BaselineMeasureUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaselineMeasureClient) UpdateOne(bm *BaselineMeasure) *BaselineMeasureUpdateOne {
	mutation := newBaselineMeasureMutation(c.config, OpUpdateOne, withBaselineMeasure(bm))
	return &BaselineMeasureUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaselineMeasureClient) UpdateOneID(id int) *BaselineMeasureUpdateOne {
	mutation := newBaselineMeasureMutation(c.config, OpUpdateOne, withBaselineMeasureID(id))
	return &BaselineMeasureUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaselineMeasure.
func (c *BaselineMeasureClient) Delete() *BaselineMeasureDelete {
	mutation := newBaselineMeasureMutation(c.config, OpDelete)
	return &BaselineMeasureDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BaselineMeasureClient) DeleteOne(bm *BaselineMeasure) *BaselineMeasureDeleteOne {
	return c.DeleteOneID(bm.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BaselineMeasureClient) DeleteOneID(id int) *BaselineMeasureDeleteOne {
	builder := c.Delete().Where(baselinemeasure.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaselineMeasureDeleteOne{builder}
}

// Query returns a query builder for BaselineMeasure.
func (c *BaselineMeasureClient) Query() *BaselineMeasureQuery {
	return &BaselineMeasureQuery{
		config: c.config,
	}
}

// Get returns a BaselineMeasure entity by its id.
func (c *BaselineMeasureClient) Get(ctx context.Context, id int) (*BaselineMeasure, error) {
	return c.Query().Where(baselinemeasure.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaselineMeasureClient) GetX(ctx context.Context, id int) *BaselineMeasure {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a BaselineMeasure.
func (c *BaselineMeasureClient) QueryParent(bm *BaselineMeasure) *BaselineCharacteristicsModuleQuery {
	query := &BaselineCharacteristicsModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinemeasure.Table, baselinemeasure.FieldID, id),
			sqlgraph.To(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinemeasure.ParentTable, baselinemeasure.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(bm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBaselineMeasureDenomList queries the baseline_measure_denom_list edge of a BaselineMeasure.
func (c *BaselineMeasureClient) QueryBaselineMeasureDenomList(bm *BaselineMeasure) *BaselineMeasureDenomQuery {
	query := &BaselineMeasureDenomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinemeasure.Table, baselinemeasure.FieldID, id),
			sqlgraph.To(baselinemeasuredenom.Table, baselinemeasuredenom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinemeasure.BaselineMeasureDenomListTable, baselinemeasure.BaselineMeasureDenomListColumn),
		)
		fromV = sqlgraph.Neighbors(bm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBaselineClassList queries the baseline_class_list edge of a BaselineMeasure.
func (c *BaselineMeasureClient) QueryBaselineClassList(bm *BaselineMeasure) *BaselineClassQuery {
	query := &BaselineClassQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinemeasure.Table, baselinemeasure.FieldID, id),
			sqlgraph.To(baselineclass.Table, baselineclass.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinemeasure.BaselineClassListTable, baselinemeasure.BaselineClassListColumn),
		)
		fromV = sqlgraph.Neighbors(bm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BaselineMeasureClient) Hooks() []Hook {
	return c.hooks.BaselineMeasure
}

// BaselineMeasureDenomClient is a client for the BaselineMeasureDenom schema.
type BaselineMeasureDenomClient struct {
	config
}

// NewBaselineMeasureDenomClient returns a client for the BaselineMeasureDenom from the given config.
func NewBaselineMeasureDenomClient(c config) *BaselineMeasureDenomClient {
	return &BaselineMeasureDenomClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `baselinemeasuredenom.Hooks(f(g(h())))`.
func (c *BaselineMeasureDenomClient) Use(hooks ...Hook) {
	c.hooks.BaselineMeasureDenom = append(c.hooks.BaselineMeasureDenom, hooks...)
}

// Create returns a create builder for BaselineMeasureDenom.
func (c *BaselineMeasureDenomClient) Create() *BaselineMeasureDenomCreate {
	mutation := newBaselineMeasureDenomMutation(c.config, OpCreate)
	return &BaselineMeasureDenomCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaselineMeasureDenom entities.
func (c *BaselineMeasureDenomClient) CreateBulk(builders ...*BaselineMeasureDenomCreate) *BaselineMeasureDenomCreateBulk {
	return &BaselineMeasureDenomCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaselineMeasureDenom.
func (c *BaselineMeasureDenomClient) Update() *BaselineMeasureDenomUpdate {
	mutation := newBaselineMeasureDenomMutation(c.config, OpUpdate)
	return &BaselineMeasureDenomUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaselineMeasureDenomClient) UpdateOne(bmd *BaselineMeasureDenom) *BaselineMeasureDenomUpdateOne {
	mutation := newBaselineMeasureDenomMutation(c.config, OpUpdateOne, withBaselineMeasureDenom(bmd))
	return &BaselineMeasureDenomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaselineMeasureDenomClient) UpdateOneID(id int) *BaselineMeasureDenomUpdateOne {
	mutation := newBaselineMeasureDenomMutation(c.config, OpUpdateOne, withBaselineMeasureDenomID(id))
	return &BaselineMeasureDenomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaselineMeasureDenom.
func (c *BaselineMeasureDenomClient) Delete() *BaselineMeasureDenomDelete {
	mutation := newBaselineMeasureDenomMutation(c.config, OpDelete)
	return &BaselineMeasureDenomDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BaselineMeasureDenomClient) DeleteOne(bmd *BaselineMeasureDenom) *BaselineMeasureDenomDeleteOne {
	return c.DeleteOneID(bmd.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BaselineMeasureDenomClient) DeleteOneID(id int) *BaselineMeasureDenomDeleteOne {
	builder := c.Delete().Where(baselinemeasuredenom.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaselineMeasureDenomDeleteOne{builder}
}

// Query returns a query builder for BaselineMeasureDenom.
func (c *BaselineMeasureDenomClient) Query() *BaselineMeasureDenomQuery {
	return &BaselineMeasureDenomQuery{
		config: c.config,
	}
}

// Get returns a BaselineMeasureDenom entity by its id.
func (c *BaselineMeasureDenomClient) Get(ctx context.Context, id int) (*BaselineMeasureDenom, error) {
	return c.Query().Where(baselinemeasuredenom.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaselineMeasureDenomClient) GetX(ctx context.Context, id int) *BaselineMeasureDenom {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a BaselineMeasureDenom.
func (c *BaselineMeasureDenomClient) QueryParent(bmd *BaselineMeasureDenom) *BaselineMeasureQuery {
	query := &BaselineMeasureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bmd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinemeasuredenom.Table, baselinemeasuredenom.FieldID, id),
			sqlgraph.To(baselinemeasure.Table, baselinemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinemeasuredenom.ParentTable, baselinemeasuredenom.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(bmd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBaselineMeasureDenomCountList queries the baseline_measure_denom_count_list edge of a BaselineMeasureDenom.
func (c *BaselineMeasureDenomClient) QueryBaselineMeasureDenomCountList(bmd *BaselineMeasureDenom) *BaselineMeasureDenomCountQuery {
	query := &BaselineMeasureDenomCountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bmd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinemeasuredenom.Table, baselinemeasuredenom.FieldID, id),
			sqlgraph.To(baselinemeasuredenomcount.Table, baselinemeasuredenomcount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinemeasuredenom.BaselineMeasureDenomCountListTable, baselinemeasuredenom.BaselineMeasureDenomCountListColumn),
		)
		fromV = sqlgraph.Neighbors(bmd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BaselineMeasureDenomClient) Hooks() []Hook {
	return c.hooks.BaselineMeasureDenom
}

// BaselineMeasureDenomCountClient is a client for the BaselineMeasureDenomCount schema.
type BaselineMeasureDenomCountClient struct {
	config
}

// NewBaselineMeasureDenomCountClient returns a client for the BaselineMeasureDenomCount from the given config.
func NewBaselineMeasureDenomCountClient(c config) *BaselineMeasureDenomCountClient {
	return &BaselineMeasureDenomCountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `baselinemeasuredenomcount.Hooks(f(g(h())))`.
func (c *BaselineMeasureDenomCountClient) Use(hooks ...Hook) {
	c.hooks.BaselineMeasureDenomCount = append(c.hooks.BaselineMeasureDenomCount, hooks...)
}

// Create returns a create builder for BaselineMeasureDenomCount.
func (c *BaselineMeasureDenomCountClient) Create() *BaselineMeasureDenomCountCreate {
	mutation := newBaselineMeasureDenomCountMutation(c.config, OpCreate)
	return &BaselineMeasureDenomCountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaselineMeasureDenomCount entities.
func (c *BaselineMeasureDenomCountClient) CreateBulk(builders ...*BaselineMeasureDenomCountCreate) *BaselineMeasureDenomCountCreateBulk {
	return &BaselineMeasureDenomCountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaselineMeasureDenomCount.
func (c *BaselineMeasureDenomCountClient) Update() *BaselineMeasureDenomCountUpdate {
	mutation := newBaselineMeasureDenomCountMutation(c.config, OpUpdate)
	return &BaselineMeasureDenomCountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaselineMeasureDenomCountClient) UpdateOne(bmdc *BaselineMeasureDenomCount) *BaselineMeasureDenomCountUpdateOne {
	mutation := newBaselineMeasureDenomCountMutation(c.config, OpUpdateOne, withBaselineMeasureDenomCount(bmdc))
	return &BaselineMeasureDenomCountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaselineMeasureDenomCountClient) UpdateOneID(id int) *BaselineMeasureDenomCountUpdateOne {
	mutation := newBaselineMeasureDenomCountMutation(c.config, OpUpdateOne, withBaselineMeasureDenomCountID(id))
	return &BaselineMeasureDenomCountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaselineMeasureDenomCount.
func (c *BaselineMeasureDenomCountClient) Delete() *BaselineMeasureDenomCountDelete {
	mutation := newBaselineMeasureDenomCountMutation(c.config, OpDelete)
	return &BaselineMeasureDenomCountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BaselineMeasureDenomCountClient) DeleteOne(bmdc *BaselineMeasureDenomCount) *BaselineMeasureDenomCountDeleteOne {
	return c.DeleteOneID(bmdc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BaselineMeasureDenomCountClient) DeleteOneID(id int) *BaselineMeasureDenomCountDeleteOne {
	builder := c.Delete().Where(baselinemeasuredenomcount.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaselineMeasureDenomCountDeleteOne{builder}
}

// Query returns a query builder for BaselineMeasureDenomCount.
func (c *BaselineMeasureDenomCountClient) Query() *BaselineMeasureDenomCountQuery {
	return &BaselineMeasureDenomCountQuery{
		config: c.config,
	}
}

// Get returns a BaselineMeasureDenomCount entity by its id.
func (c *BaselineMeasureDenomCountClient) Get(ctx context.Context, id int) (*BaselineMeasureDenomCount, error) {
	return c.Query().Where(baselinemeasuredenomcount.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaselineMeasureDenomCountClient) GetX(ctx context.Context, id int) *BaselineMeasureDenomCount {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a BaselineMeasureDenomCount.
func (c *BaselineMeasureDenomCountClient) QueryParent(bmdc *BaselineMeasureDenomCount) *BaselineMeasureDenomQuery {
	query := &BaselineMeasureDenomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bmdc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinemeasuredenomcount.Table, baselinemeasuredenomcount.FieldID, id),
			sqlgraph.To(baselinemeasuredenom.Table, baselinemeasuredenom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinemeasuredenomcount.ParentTable, baselinemeasuredenomcount.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(bmdc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BaselineMeasureDenomCountClient) Hooks() []Hook {
	return c.hooks.BaselineMeasureDenomCount
}

// BaselineMeasurementClient is a client for the BaselineMeasurement schema.
type BaselineMeasurementClient struct {
	config
}

// NewBaselineMeasurementClient returns a client for the BaselineMeasurement from the given config.
func NewBaselineMeasurementClient(c config) *BaselineMeasurementClient {
	return &BaselineMeasurementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `baselinemeasurement.Hooks(f(g(h())))`.
func (c *BaselineMeasurementClient) Use(hooks ...Hook) {
	c.hooks.BaselineMeasurement = append(c.hooks.BaselineMeasurement, hooks...)
}

// Create returns a create builder for BaselineMeasurement.
func (c *BaselineMeasurementClient) Create() *BaselineMeasurementCreate {
	mutation := newBaselineMeasurementMutation(c.config, OpCreate)
	return &BaselineMeasurementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaselineMeasurement entities.
func (c *BaselineMeasurementClient) CreateBulk(builders ...*BaselineMeasurementCreate) *BaselineMeasurementCreateBulk {
	return &BaselineMeasurementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaselineMeasurement.
func (c *BaselineMeasurementClient) Update() *BaselineMeasurementUpdate {
	mutation := newBaselineMeasurementMutation(c.config, OpUpdate)
	return &BaselineMeasurementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaselineMeasurementClient) UpdateOne(bm *BaselineMeasurement) *BaselineMeasurementUpdateOne {
	mutation := newBaselineMeasurementMutation(c.config, OpUpdateOne, withBaselineMeasurement(bm))
	return &BaselineMeasurementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaselineMeasurementClient) UpdateOneID(id int) *BaselineMeasurementUpdateOne {
	mutation := newBaselineMeasurementMutation(c.config, OpUpdateOne, withBaselineMeasurementID(id))
	return &BaselineMeasurementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaselineMeasurement.
func (c *BaselineMeasurementClient) Delete() *BaselineMeasurementDelete {
	mutation := newBaselineMeasurementMutation(c.config, OpDelete)
	return &BaselineMeasurementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BaselineMeasurementClient) DeleteOne(bm *BaselineMeasurement) *BaselineMeasurementDeleteOne {
	return c.DeleteOneID(bm.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BaselineMeasurementClient) DeleteOneID(id int) *BaselineMeasurementDeleteOne {
	builder := c.Delete().Where(baselinemeasurement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaselineMeasurementDeleteOne{builder}
}

// Query returns a query builder for BaselineMeasurement.
func (c *BaselineMeasurementClient) Query() *BaselineMeasurementQuery {
	return &BaselineMeasurementQuery{
		config: c.config,
	}
}

// Get returns a BaselineMeasurement entity by its id.
func (c *BaselineMeasurementClient) Get(ctx context.Context, id int) (*BaselineMeasurement, error) {
	return c.Query().Where(baselinemeasurement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaselineMeasurementClient) GetX(ctx context.Context, id int) *BaselineMeasurement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a BaselineMeasurement.
func (c *BaselineMeasurementClient) QueryParent(bm *BaselineMeasurement) *BaselineCategoryQuery {
	query := &BaselineCategoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinemeasurement.Table, baselinemeasurement.FieldID, id),
			sqlgraph.To(baselinecategory.Table, baselinecategory.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinemeasurement.ParentTable, baselinemeasurement.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(bm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BaselineMeasurementClient) Hooks() []Hook {
	return c.hooks.BaselineMeasurement
}

// CertainAgreementClient is a client for the CertainAgreement schema.
type CertainAgreementClient struct {
	config
}

// NewCertainAgreementClient returns a client for the CertainAgreement from the given config.
func NewCertainAgreementClient(c config) *CertainAgreementClient {
	return &CertainAgreementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `certainagreement.Hooks(f(g(h())))`.
func (c *CertainAgreementClient) Use(hooks ...Hook) {
	c.hooks.CertainAgreement = append(c.hooks.CertainAgreement, hooks...)
}

// Create returns a create builder for CertainAgreement.
func (c *CertainAgreementClient) Create() *CertainAgreementCreate {
	mutation := newCertainAgreementMutation(c.config, OpCreate)
	return &CertainAgreementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CertainAgreement entities.
func (c *CertainAgreementClient) CreateBulk(builders ...*CertainAgreementCreate) *CertainAgreementCreateBulk {
	return &CertainAgreementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CertainAgreement.
func (c *CertainAgreementClient) Update() *CertainAgreementUpdate {
	mutation := newCertainAgreementMutation(c.config, OpUpdate)
	return &CertainAgreementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CertainAgreementClient) UpdateOne(ca *CertainAgreement) *CertainAgreementUpdateOne {
	mutation := newCertainAgreementMutation(c.config, OpUpdateOne, withCertainAgreement(ca))
	return &CertainAgreementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CertainAgreementClient) UpdateOneID(id int) *CertainAgreementUpdateOne {
	mutation := newCertainAgreementMutation(c.config, OpUpdateOne, withCertainAgreementID(id))
	return &CertainAgreementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CertainAgreement.
func (c *CertainAgreementClient) Delete() *CertainAgreementDelete {
	mutation := newCertainAgreementMutation(c.config, OpDelete)
	return &CertainAgreementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CertainAgreementClient) DeleteOne(ca *CertainAgreement) *CertainAgreementDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CertainAgreementClient) DeleteOneID(id int) *CertainAgreementDeleteOne {
	builder := c.Delete().Where(certainagreement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CertainAgreementDeleteOne{builder}
}

// Query returns a query builder for CertainAgreement.
func (c *CertainAgreementClient) Query() *CertainAgreementQuery {
	return &CertainAgreementQuery{
		config: c.config,
	}
}

// Get returns a CertainAgreement entity by its id.
func (c *CertainAgreementClient) Get(ctx context.Context, id int) (*CertainAgreement, error) {
	return c.Query().Where(certainagreement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CertainAgreementClient) GetX(ctx context.Context, id int) *CertainAgreement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a CertainAgreement.
func (c *CertainAgreementClient) QueryParent(ca *CertainAgreement) *MoreInfoModuleQuery {
	query := &MoreInfoModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(certainagreement.Table, certainagreement.FieldID, id),
			sqlgraph.To(moreinfomodule.Table, moreinfomodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, certainagreement.ParentTable, certainagreement.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CertainAgreementClient) Hooks() []Hook {
	return c.hooks.CertainAgreement
}

// ClinicalTrialClient is a client for the ClinicalTrial schema.
type ClinicalTrialClient struct {
	config
}

// NewClinicalTrialClient returns a client for the ClinicalTrial from the given config.
func NewClinicalTrialClient(c config) *ClinicalTrialClient {
	return &ClinicalTrialClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `clinicaltrial.Hooks(f(g(h())))`.
func (c *ClinicalTrialClient) Use(hooks ...Hook) {
	c.hooks.ClinicalTrial = append(c.hooks.ClinicalTrial, hooks...)
}

// Create returns a create builder for ClinicalTrial.
func (c *ClinicalTrialClient) Create() *ClinicalTrialCreate {
	mutation := newClinicalTrialMutation(c.config, OpCreate)
	return &ClinicalTrialCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ClinicalTrial entities.
func (c *ClinicalTrialClient) CreateBulk(builders ...*ClinicalTrialCreate) *ClinicalTrialCreateBulk {
	return &ClinicalTrialCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ClinicalTrial.
func (c *ClinicalTrialClient) Update() *ClinicalTrialUpdate {
	mutation := newClinicalTrialMutation(c.config, OpUpdate)
	return &ClinicalTrialUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ClinicalTrialClient) UpdateOne(ct *ClinicalTrial) *ClinicalTrialUpdateOne {
	mutation := newClinicalTrialMutation(c.config, OpUpdateOne, withClinicalTrial(ct))
	return &ClinicalTrialUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ClinicalTrialClient) UpdateOneID(id int) *ClinicalTrialUpdateOne {
	mutation := newClinicalTrialMutation(c.config, OpUpdateOne, withClinicalTrialID(id))
	return &ClinicalTrialUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ClinicalTrial.
func (c *ClinicalTrialClient) Delete() *ClinicalTrialDelete {
	mutation := newClinicalTrialMutation(c.config, OpDelete)
	return &ClinicalTrialDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ClinicalTrialClient) DeleteOne(ct *ClinicalTrial) *ClinicalTrialDeleteOne {
	return c.DeleteOneID(ct.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ClinicalTrialClient) DeleteOneID(id int) *ClinicalTrialDeleteOne {
	builder := c.Delete().Where(clinicaltrial.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ClinicalTrialDeleteOne{builder}
}

// Query returns a query builder for ClinicalTrial.
func (c *ClinicalTrialClient) Query() *ClinicalTrialQuery {
	return &ClinicalTrialQuery{
		config: c.config,
	}
}

// Get returns a ClinicalTrial entity by its id.
func (c *ClinicalTrialClient) Get(ctx context.Context, id int) (*ClinicalTrial, error) {
	return c.Query().Where(clinicaltrial.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ClinicalTrialClient) GetX(ctx context.Context, id int) *ClinicalTrial {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryResultsDefinition queries the results_definition edge of a ClinicalTrial.
func (c *ClinicalTrialClient) QueryResultsDefinition(ct *ClinicalTrial) *ResultsDefinitionQuery {
	query := &ResultsDefinitionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ct.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(clinicaltrial.Table, clinicaltrial.FieldID, id),
			sqlgraph.To(resultsdefinition.Table, resultsdefinition.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, clinicaltrial.ResultsDefinitionTable, clinicaltrial.ResultsDefinitionColumn),
		)
		fromV = sqlgraph.Neighbors(ct.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStudyLocations queries the study_locations edge of a ClinicalTrial.
func (c *ClinicalTrialClient) QueryStudyLocations(ct *ClinicalTrial) *StudyLocationQuery {
	query := &StudyLocationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ct.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(clinicaltrial.Table, clinicaltrial.FieldID, id),
			sqlgraph.To(studylocation.Table, studylocation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, clinicaltrial.StudyLocationsTable, clinicaltrial.StudyLocationsColumn),
		)
		fromV = sqlgraph.Neighbors(ct.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStudyEligibility queries the study_eligibility edge of a ClinicalTrial.
func (c *ClinicalTrialClient) QueryStudyEligibility(ct *ClinicalTrial) *StudyEligibilityQuery {
	query := &StudyEligibilityQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ct.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(clinicaltrial.Table, clinicaltrial.FieldID, id),
			sqlgraph.To(studyeligibility.Table, studyeligibility.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, clinicaltrial.StudyEligibilityTable, clinicaltrial.StudyEligibilityColumn),
		)
		fromV = sqlgraph.Neighbors(ct.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMetadata queries the metadata edge of a ClinicalTrial.
func (c *ClinicalTrialClient) QueryMetadata(ct *ClinicalTrial) *ClinicalTrialMetadataQuery {
	query := &ClinicalTrialMetadataQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ct.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(clinicaltrial.Table, clinicaltrial.FieldID, id),
			sqlgraph.To(clinicaltrialmetadata.Table, clinicaltrialmetadata.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, clinicaltrial.MetadataTable, clinicaltrial.MetadataColumn),
		)
		fromV = sqlgraph.Neighbors(ct.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ClinicalTrialClient) Hooks() []Hook {
	return c.hooks.ClinicalTrial
}

// ClinicalTrialMetadataClient is a client for the ClinicalTrialMetadata schema.
type ClinicalTrialMetadataClient struct {
	config
}

// NewClinicalTrialMetadataClient returns a client for the ClinicalTrialMetadata from the given config.
func NewClinicalTrialMetadataClient(c config) *ClinicalTrialMetadataClient {
	return &ClinicalTrialMetadataClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `clinicaltrialmetadata.Hooks(f(g(h())))`.
func (c *ClinicalTrialMetadataClient) Use(hooks ...Hook) {
	c.hooks.ClinicalTrialMetadata = append(c.hooks.ClinicalTrialMetadata, hooks...)
}

// Create returns a create builder for ClinicalTrialMetadata.
func (c *ClinicalTrialMetadataClient) Create() *ClinicalTrialMetadataCreate {
	mutation := newClinicalTrialMetadataMutation(c.config, OpCreate)
	return &ClinicalTrialMetadataCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ClinicalTrialMetadata entities.
func (c *ClinicalTrialMetadataClient) CreateBulk(builders ...*ClinicalTrialMetadataCreate) *ClinicalTrialMetadataCreateBulk {
	return &ClinicalTrialMetadataCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ClinicalTrialMetadata.
func (c *ClinicalTrialMetadataClient) Update() *ClinicalTrialMetadataUpdate {
	mutation := newClinicalTrialMetadataMutation(c.config, OpUpdate)
	return &ClinicalTrialMetadataUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ClinicalTrialMetadataClient) UpdateOne(ctm *ClinicalTrialMetadata) *ClinicalTrialMetadataUpdateOne {
	mutation := newClinicalTrialMetadataMutation(c.config, OpUpdateOne, withClinicalTrialMetadata(ctm))
	return &ClinicalTrialMetadataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ClinicalTrialMetadataClient) UpdateOneID(id int) *ClinicalTrialMetadataUpdateOne {
	mutation := newClinicalTrialMetadataMutation(c.config, OpUpdateOne, withClinicalTrialMetadataID(id))
	return &ClinicalTrialMetadataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ClinicalTrialMetadata.
func (c *ClinicalTrialMetadataClient) Delete() *ClinicalTrialMetadataDelete {
	mutation := newClinicalTrialMetadataMutation(c.config, OpDelete)
	return &ClinicalTrialMetadataDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ClinicalTrialMetadataClient) DeleteOne(ctm *ClinicalTrialMetadata) *ClinicalTrialMetadataDeleteOne {
	return c.DeleteOneID(ctm.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ClinicalTrialMetadataClient) DeleteOneID(id int) *ClinicalTrialMetadataDeleteOne {
	builder := c.Delete().Where(clinicaltrialmetadata.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ClinicalTrialMetadataDeleteOne{builder}
}

// Query returns a query builder for ClinicalTrialMetadata.
func (c *ClinicalTrialMetadataClient) Query() *ClinicalTrialMetadataQuery {
	return &ClinicalTrialMetadataQuery{
		config: c.config,
	}
}

// Get returns a ClinicalTrialMetadata entity by its id.
func (c *ClinicalTrialMetadataClient) Get(ctx context.Context, id int) (*ClinicalTrialMetadata, error) {
	return c.Query().Where(clinicaltrialmetadata.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ClinicalTrialMetadataClient) GetX(ctx context.Context, id int) *ClinicalTrialMetadata {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a ClinicalTrialMetadata.
func (c *ClinicalTrialMetadataClient) QueryParent(ctm *ClinicalTrialMetadata) *ClinicalTrialQuery {
	query := &ClinicalTrialQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ctm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(clinicaltrialmetadata.Table, clinicaltrialmetadata.FieldID, id),
			sqlgraph.To(clinicaltrial.Table, clinicaltrial.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, clinicaltrialmetadata.ParentTable, clinicaltrialmetadata.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(ctm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ClinicalTrialMetadataClient) Hooks() []Hook {
	return c.hooks.ClinicalTrialMetadata
}

// DoseDescriptionClient is a client for the DoseDescription schema.
type DoseDescriptionClient struct {
	config
}

// NewDoseDescriptionClient returns a client for the DoseDescription from the given config.
func NewDoseDescriptionClient(c config) *DoseDescriptionClient {
	return &DoseDescriptionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dosedescription.Hooks(f(g(h())))`.
func (c *DoseDescriptionClient) Use(hooks ...Hook) {
	c.hooks.DoseDescription = append(c.hooks.DoseDescription, hooks...)
}

// Create returns a create builder for DoseDescription.
func (c *DoseDescriptionClient) Create() *DoseDescriptionCreate {
	mutation := newDoseDescriptionMutation(c.config, OpCreate)
	return &DoseDescriptionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of DoseDescription entities.
func (c *DoseDescriptionClient) CreateBulk(builders ...*DoseDescriptionCreate) *DoseDescriptionCreateBulk {
	return &DoseDescriptionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DoseDescription.
func (c *DoseDescriptionClient) Update() *DoseDescriptionUpdate {
	mutation := newDoseDescriptionMutation(c.config, OpUpdate)
	return &DoseDescriptionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DoseDescriptionClient) UpdateOne(dd *DoseDescription) *DoseDescriptionUpdateOne {
	mutation := newDoseDescriptionMutation(c.config, OpUpdateOne, withDoseDescription(dd))
	return &DoseDescriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DoseDescriptionClient) UpdateOneID(id int) *DoseDescriptionUpdateOne {
	mutation := newDoseDescriptionMutation(c.config, OpUpdateOne, withDoseDescriptionID(id))
	return &DoseDescriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DoseDescription.
func (c *DoseDescriptionClient) Delete() *DoseDescriptionDelete {
	mutation := newDoseDescriptionMutation(c.config, OpDelete)
	return &DoseDescriptionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DoseDescriptionClient) DeleteOne(dd *DoseDescription) *DoseDescriptionDeleteOne {
	return c.DeleteOneID(dd.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DoseDescriptionClient) DeleteOneID(id int) *DoseDescriptionDeleteOne {
	builder := c.Delete().Where(dosedescription.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DoseDescriptionDeleteOne{builder}
}

// Query returns a query builder for DoseDescription.
func (c *DoseDescriptionClient) Query() *DoseDescriptionQuery {
	return &DoseDescriptionQuery{
		config: c.config,
	}
}

// Get returns a DoseDescription entity by its id.
func (c *DoseDescriptionClient) Get(ctx context.Context, id int) (*DoseDescription, error) {
	return c.Query().Where(dosedescription.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DoseDescriptionClient) GetX(ctx context.Context, id int) *DoseDescription {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DoseDescriptionClient) Hooks() []Hook {
	return c.hooks.DoseDescription
}

// EventGroupClient is a client for the EventGroup schema.
type EventGroupClient struct {
	config
}

// NewEventGroupClient returns a client for the EventGroup from the given config.
func NewEventGroupClient(c config) *EventGroupClient {
	return &EventGroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `eventgroup.Hooks(f(g(h())))`.
func (c *EventGroupClient) Use(hooks ...Hook) {
	c.hooks.EventGroup = append(c.hooks.EventGroup, hooks...)
}

// Create returns a create builder for EventGroup.
func (c *EventGroupClient) Create() *EventGroupCreate {
	mutation := newEventGroupMutation(c.config, OpCreate)
	return &EventGroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EventGroup entities.
func (c *EventGroupClient) CreateBulk(builders ...*EventGroupCreate) *EventGroupCreateBulk {
	return &EventGroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EventGroup.
func (c *EventGroupClient) Update() *EventGroupUpdate {
	mutation := newEventGroupMutation(c.config, OpUpdate)
	return &EventGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EventGroupClient) UpdateOne(eg *EventGroup) *EventGroupUpdateOne {
	mutation := newEventGroupMutation(c.config, OpUpdateOne, withEventGroup(eg))
	return &EventGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EventGroupClient) UpdateOneID(id int) *EventGroupUpdateOne {
	mutation := newEventGroupMutation(c.config, OpUpdateOne, withEventGroupID(id))
	return &EventGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EventGroup.
func (c *EventGroupClient) Delete() *EventGroupDelete {
	mutation := newEventGroupMutation(c.config, OpDelete)
	return &EventGroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *EventGroupClient) DeleteOne(eg *EventGroup) *EventGroupDeleteOne {
	return c.DeleteOneID(eg.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *EventGroupClient) DeleteOneID(id int) *EventGroupDeleteOne {
	builder := c.Delete().Where(eventgroup.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EventGroupDeleteOne{builder}
}

// Query returns a query builder for EventGroup.
func (c *EventGroupClient) Query() *EventGroupQuery {
	return &EventGroupQuery{
		config: c.config,
	}
}

// Get returns a EventGroup entity by its id.
func (c *EventGroupClient) Get(ctx context.Context, id int) (*EventGroup, error) {
	return c.Query().Where(eventgroup.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EventGroupClient) GetX(ctx context.Context, id int) *EventGroup {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a EventGroup.
func (c *EventGroupClient) QueryParent(eg *EventGroup) *AdverseEventsModuleQuery {
	query := &AdverseEventsModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := eg.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(eventgroup.Table, eventgroup.FieldID, id),
			sqlgraph.To(adverseeventsmodule.Table, adverseeventsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, eventgroup.ParentTable, eventgroup.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(eg.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EventGroupClient) Hooks() []Hook {
	return c.hooks.EventGroup
}

// FlowAchievementClient is a client for the FlowAchievement schema.
type FlowAchievementClient struct {
	config
}

// NewFlowAchievementClient returns a client for the FlowAchievement from the given config.
func NewFlowAchievementClient(c config) *FlowAchievementClient {
	return &FlowAchievementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `flowachievement.Hooks(f(g(h())))`.
func (c *FlowAchievementClient) Use(hooks ...Hook) {
	c.hooks.FlowAchievement = append(c.hooks.FlowAchievement, hooks...)
}

// Create returns a create builder for FlowAchievement.
func (c *FlowAchievementClient) Create() *FlowAchievementCreate {
	mutation := newFlowAchievementMutation(c.config, OpCreate)
	return &FlowAchievementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FlowAchievement entities.
func (c *FlowAchievementClient) CreateBulk(builders ...*FlowAchievementCreate) *FlowAchievementCreateBulk {
	return &FlowAchievementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FlowAchievement.
func (c *FlowAchievementClient) Update() *FlowAchievementUpdate {
	mutation := newFlowAchievementMutation(c.config, OpUpdate)
	return &FlowAchievementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FlowAchievementClient) UpdateOne(fa *FlowAchievement) *FlowAchievementUpdateOne {
	mutation := newFlowAchievementMutation(c.config, OpUpdateOne, withFlowAchievement(fa))
	return &FlowAchievementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FlowAchievementClient) UpdateOneID(id int) *FlowAchievementUpdateOne {
	mutation := newFlowAchievementMutation(c.config, OpUpdateOne, withFlowAchievementID(id))
	return &FlowAchievementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FlowAchievement.
func (c *FlowAchievementClient) Delete() *FlowAchievementDelete {
	mutation := newFlowAchievementMutation(c.config, OpDelete)
	return &FlowAchievementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FlowAchievementClient) DeleteOne(fa *FlowAchievement) *FlowAchievementDeleteOne {
	return c.DeleteOneID(fa.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FlowAchievementClient) DeleteOneID(id int) *FlowAchievementDeleteOne {
	builder := c.Delete().Where(flowachievement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FlowAchievementDeleteOne{builder}
}

// Query returns a query builder for FlowAchievement.
func (c *FlowAchievementClient) Query() *FlowAchievementQuery {
	return &FlowAchievementQuery{
		config: c.config,
	}
}

// Get returns a FlowAchievement entity by its id.
func (c *FlowAchievementClient) Get(ctx context.Context, id int) (*FlowAchievement, error) {
	return c.Query().Where(flowachievement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FlowAchievementClient) GetX(ctx context.Context, id int) *FlowAchievement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a FlowAchievement.
func (c *FlowAchievementClient) QueryParent(fa *FlowAchievement) *FlowMilestoneQuery {
	query := &FlowMilestoneQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := fa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(flowachievement.Table, flowachievement.FieldID, id),
			sqlgraph.To(flowmilestone.Table, flowmilestone.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flowachievement.ParentTable, flowachievement.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(fa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FlowAchievementClient) Hooks() []Hook {
	return c.hooks.FlowAchievement
}

// FlowDropWithdrawClient is a client for the FlowDropWithdraw schema.
type FlowDropWithdrawClient struct {
	config
}

// NewFlowDropWithdrawClient returns a client for the FlowDropWithdraw from the given config.
func NewFlowDropWithdrawClient(c config) *FlowDropWithdrawClient {
	return &FlowDropWithdrawClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `flowdropwithdraw.Hooks(f(g(h())))`.
func (c *FlowDropWithdrawClient) Use(hooks ...Hook) {
	c.hooks.FlowDropWithdraw = append(c.hooks.FlowDropWithdraw, hooks...)
}

// Create returns a create builder for FlowDropWithdraw.
func (c *FlowDropWithdrawClient) Create() *FlowDropWithdrawCreate {
	mutation := newFlowDropWithdrawMutation(c.config, OpCreate)
	return &FlowDropWithdrawCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FlowDropWithdraw entities.
func (c *FlowDropWithdrawClient) CreateBulk(builders ...*FlowDropWithdrawCreate) *FlowDropWithdrawCreateBulk {
	return &FlowDropWithdrawCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FlowDropWithdraw.
func (c *FlowDropWithdrawClient) Update() *FlowDropWithdrawUpdate {
	mutation := newFlowDropWithdrawMutation(c.config, OpUpdate)
	return &FlowDropWithdrawUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FlowDropWithdrawClient) UpdateOne(fdw *FlowDropWithdraw) *FlowDropWithdrawUpdateOne {
	mutation := newFlowDropWithdrawMutation(c.config, OpUpdateOne, withFlowDropWithdraw(fdw))
	return &FlowDropWithdrawUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FlowDropWithdrawClient) UpdateOneID(id int) *FlowDropWithdrawUpdateOne {
	mutation := newFlowDropWithdrawMutation(c.config, OpUpdateOne, withFlowDropWithdrawID(id))
	return &FlowDropWithdrawUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FlowDropWithdraw.
func (c *FlowDropWithdrawClient) Delete() *FlowDropWithdrawDelete {
	mutation := newFlowDropWithdrawMutation(c.config, OpDelete)
	return &FlowDropWithdrawDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FlowDropWithdrawClient) DeleteOne(fdw *FlowDropWithdraw) *FlowDropWithdrawDeleteOne {
	return c.DeleteOneID(fdw.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FlowDropWithdrawClient) DeleteOneID(id int) *FlowDropWithdrawDeleteOne {
	builder := c.Delete().Where(flowdropwithdraw.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FlowDropWithdrawDeleteOne{builder}
}

// Query returns a query builder for FlowDropWithdraw.
func (c *FlowDropWithdrawClient) Query() *FlowDropWithdrawQuery {
	return &FlowDropWithdrawQuery{
		config: c.config,
	}
}

// Get returns a FlowDropWithdraw entity by its id.
func (c *FlowDropWithdrawClient) Get(ctx context.Context, id int) (*FlowDropWithdraw, error) {
	return c.Query().Where(flowdropwithdraw.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FlowDropWithdrawClient) GetX(ctx context.Context, id int) *FlowDropWithdraw {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a FlowDropWithdraw.
func (c *FlowDropWithdrawClient) QueryParent(fdw *FlowDropWithdraw) *FlowPeriodQuery {
	query := &FlowPeriodQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := fdw.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(flowdropwithdraw.Table, flowdropwithdraw.FieldID, id),
			sqlgraph.To(flowperiod.Table, flowperiod.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flowdropwithdraw.ParentTable, flowdropwithdraw.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(fdw.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFlowReasonList queries the flow_reason_list edge of a FlowDropWithdraw.
func (c *FlowDropWithdrawClient) QueryFlowReasonList(fdw *FlowDropWithdraw) *FlowReasonQuery {
	query := &FlowReasonQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := fdw.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(flowdropwithdraw.Table, flowdropwithdraw.FieldID, id),
			sqlgraph.To(flowreason.Table, flowreason.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, flowdropwithdraw.FlowReasonListTable, flowdropwithdraw.FlowReasonListColumn),
		)
		fromV = sqlgraph.Neighbors(fdw.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FlowDropWithdrawClient) Hooks() []Hook {
	return c.hooks.FlowDropWithdraw
}

// FlowGroupClient is a client for the FlowGroup schema.
type FlowGroupClient struct {
	config
}

// NewFlowGroupClient returns a client for the FlowGroup from the given config.
func NewFlowGroupClient(c config) *FlowGroupClient {
	return &FlowGroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `flowgroup.Hooks(f(g(h())))`.
func (c *FlowGroupClient) Use(hooks ...Hook) {
	c.hooks.FlowGroup = append(c.hooks.FlowGroup, hooks...)
}

// Create returns a create builder for FlowGroup.
func (c *FlowGroupClient) Create() *FlowGroupCreate {
	mutation := newFlowGroupMutation(c.config, OpCreate)
	return &FlowGroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FlowGroup entities.
func (c *FlowGroupClient) CreateBulk(builders ...*FlowGroupCreate) *FlowGroupCreateBulk {
	return &FlowGroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FlowGroup.
func (c *FlowGroupClient) Update() *FlowGroupUpdate {
	mutation := newFlowGroupMutation(c.config, OpUpdate)
	return &FlowGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FlowGroupClient) UpdateOne(fg *FlowGroup) *FlowGroupUpdateOne {
	mutation := newFlowGroupMutation(c.config, OpUpdateOne, withFlowGroup(fg))
	return &FlowGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FlowGroupClient) UpdateOneID(id int) *FlowGroupUpdateOne {
	mutation := newFlowGroupMutation(c.config, OpUpdateOne, withFlowGroupID(id))
	return &FlowGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FlowGroup.
func (c *FlowGroupClient) Delete() *FlowGroupDelete {
	mutation := newFlowGroupMutation(c.config, OpDelete)
	return &FlowGroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FlowGroupClient) DeleteOne(fg *FlowGroup) *FlowGroupDeleteOne {
	return c.DeleteOneID(fg.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FlowGroupClient) DeleteOneID(id int) *FlowGroupDeleteOne {
	builder := c.Delete().Where(flowgroup.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FlowGroupDeleteOne{builder}
}

// Query returns a query builder for FlowGroup.
func (c *FlowGroupClient) Query() *FlowGroupQuery {
	return &FlowGroupQuery{
		config: c.config,
	}
}

// Get returns a FlowGroup entity by its id.
func (c *FlowGroupClient) Get(ctx context.Context, id int) (*FlowGroup, error) {
	return c.Query().Where(flowgroup.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FlowGroupClient) GetX(ctx context.Context, id int) *FlowGroup {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a FlowGroup.
func (c *FlowGroupClient) QueryParent(fg *FlowGroup) *ParticipantFlowModuleQuery {
	query := &ParticipantFlowModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := fg.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(flowgroup.Table, flowgroup.FieldID, id),
			sqlgraph.To(participantflowmodule.Table, participantflowmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flowgroup.ParentTable, flowgroup.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(fg.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FlowGroupClient) Hooks() []Hook {
	return c.hooks.FlowGroup
}

// FlowMilestoneClient is a client for the FlowMilestone schema.
type FlowMilestoneClient struct {
	config
}

// NewFlowMilestoneClient returns a client for the FlowMilestone from the given config.
func NewFlowMilestoneClient(c config) *FlowMilestoneClient {
	return &FlowMilestoneClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `flowmilestone.Hooks(f(g(h())))`.
func (c *FlowMilestoneClient) Use(hooks ...Hook) {
	c.hooks.FlowMilestone = append(c.hooks.FlowMilestone, hooks...)
}

// Create returns a create builder for FlowMilestone.
func (c *FlowMilestoneClient) Create() *FlowMilestoneCreate {
	mutation := newFlowMilestoneMutation(c.config, OpCreate)
	return &FlowMilestoneCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FlowMilestone entities.
func (c *FlowMilestoneClient) CreateBulk(builders ...*FlowMilestoneCreate) *FlowMilestoneCreateBulk {
	return &FlowMilestoneCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FlowMilestone.
func (c *FlowMilestoneClient) Update() *FlowMilestoneUpdate {
	mutation := newFlowMilestoneMutation(c.config, OpUpdate)
	return &FlowMilestoneUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FlowMilestoneClient) UpdateOne(fm *FlowMilestone) *FlowMilestoneUpdateOne {
	mutation := newFlowMilestoneMutation(c.config, OpUpdateOne, withFlowMilestone(fm))
	return &FlowMilestoneUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FlowMilestoneClient) UpdateOneID(id int) *FlowMilestoneUpdateOne {
	mutation := newFlowMilestoneMutation(c.config, OpUpdateOne, withFlowMilestoneID(id))
	return &FlowMilestoneUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FlowMilestone.
func (c *FlowMilestoneClient) Delete() *FlowMilestoneDelete {
	mutation := newFlowMilestoneMutation(c.config, OpDelete)
	return &FlowMilestoneDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FlowMilestoneClient) DeleteOne(fm *FlowMilestone) *FlowMilestoneDeleteOne {
	return c.DeleteOneID(fm.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FlowMilestoneClient) DeleteOneID(id int) *FlowMilestoneDeleteOne {
	builder := c.Delete().Where(flowmilestone.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FlowMilestoneDeleteOne{builder}
}

// Query returns a query builder for FlowMilestone.
func (c *FlowMilestoneClient) Query() *FlowMilestoneQuery {
	return &FlowMilestoneQuery{
		config: c.config,
	}
}

// Get returns a FlowMilestone entity by its id.
func (c *FlowMilestoneClient) Get(ctx context.Context, id int) (*FlowMilestone, error) {
	return c.Query().Where(flowmilestone.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FlowMilestoneClient) GetX(ctx context.Context, id int) *FlowMilestone {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a FlowMilestone.
func (c *FlowMilestoneClient) QueryParent(fm *FlowMilestone) *FlowPeriodQuery {
	query := &FlowPeriodQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := fm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(flowmilestone.Table, flowmilestone.FieldID, id),
			sqlgraph.To(flowperiod.Table, flowperiod.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flowmilestone.ParentTable, flowmilestone.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(fm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFlowAchievementList queries the flow_achievement_list edge of a FlowMilestone.
func (c *FlowMilestoneClient) QueryFlowAchievementList(fm *FlowMilestone) *FlowAchievementQuery {
	query := &FlowAchievementQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := fm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(flowmilestone.Table, flowmilestone.FieldID, id),
			sqlgraph.To(flowachievement.Table, flowachievement.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, flowmilestone.FlowAchievementListTable, flowmilestone.FlowAchievementListColumn),
		)
		fromV = sqlgraph.Neighbors(fm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FlowMilestoneClient) Hooks() []Hook {
	return c.hooks.FlowMilestone
}

// FlowPeriodClient is a client for the FlowPeriod schema.
type FlowPeriodClient struct {
	config
}

// NewFlowPeriodClient returns a client for the FlowPeriod from the given config.
func NewFlowPeriodClient(c config) *FlowPeriodClient {
	return &FlowPeriodClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `flowperiod.Hooks(f(g(h())))`.
func (c *FlowPeriodClient) Use(hooks ...Hook) {
	c.hooks.FlowPeriod = append(c.hooks.FlowPeriod, hooks...)
}

// Create returns a create builder for FlowPeriod.
func (c *FlowPeriodClient) Create() *FlowPeriodCreate {
	mutation := newFlowPeriodMutation(c.config, OpCreate)
	return &FlowPeriodCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FlowPeriod entities.
func (c *FlowPeriodClient) CreateBulk(builders ...*FlowPeriodCreate) *FlowPeriodCreateBulk {
	return &FlowPeriodCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FlowPeriod.
func (c *FlowPeriodClient) Update() *FlowPeriodUpdate {
	mutation := newFlowPeriodMutation(c.config, OpUpdate)
	return &FlowPeriodUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FlowPeriodClient) UpdateOne(fp *FlowPeriod) *FlowPeriodUpdateOne {
	mutation := newFlowPeriodMutation(c.config, OpUpdateOne, withFlowPeriod(fp))
	return &FlowPeriodUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FlowPeriodClient) UpdateOneID(id int) *FlowPeriodUpdateOne {
	mutation := newFlowPeriodMutation(c.config, OpUpdateOne, withFlowPeriodID(id))
	return &FlowPeriodUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FlowPeriod.
func (c *FlowPeriodClient) Delete() *FlowPeriodDelete {
	mutation := newFlowPeriodMutation(c.config, OpDelete)
	return &FlowPeriodDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FlowPeriodClient) DeleteOne(fp *FlowPeriod) *FlowPeriodDeleteOne {
	return c.DeleteOneID(fp.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FlowPeriodClient) DeleteOneID(id int) *FlowPeriodDeleteOne {
	builder := c.Delete().Where(flowperiod.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FlowPeriodDeleteOne{builder}
}

// Query returns a query builder for FlowPeriod.
func (c *FlowPeriodClient) Query() *FlowPeriodQuery {
	return &FlowPeriodQuery{
		config: c.config,
	}
}

// Get returns a FlowPeriod entity by its id.
func (c *FlowPeriodClient) Get(ctx context.Context, id int) (*FlowPeriod, error) {
	return c.Query().Where(flowperiod.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FlowPeriodClient) GetX(ctx context.Context, id int) *FlowPeriod {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a FlowPeriod.
func (c *FlowPeriodClient) QueryParent(fp *FlowPeriod) *ParticipantFlowModuleQuery {
	query := &ParticipantFlowModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := fp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(flowperiod.Table, flowperiod.FieldID, id),
			sqlgraph.To(participantflowmodule.Table, participantflowmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flowperiod.ParentTable, flowperiod.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(fp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFlowMilestoneList queries the flow_milestone_list edge of a FlowPeriod.
func (c *FlowPeriodClient) QueryFlowMilestoneList(fp *FlowPeriod) *FlowMilestoneQuery {
	query := &FlowMilestoneQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := fp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(flowperiod.Table, flowperiod.FieldID, id),
			sqlgraph.To(flowmilestone.Table, flowmilestone.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, flowperiod.FlowMilestoneListTable, flowperiod.FlowMilestoneListColumn),
		)
		fromV = sqlgraph.Neighbors(fp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFlowDropWithdrawList queries the flow_drop_withdraw_list edge of a FlowPeriod.
func (c *FlowPeriodClient) QueryFlowDropWithdrawList(fp *FlowPeriod) *FlowDropWithdrawQuery {
	query := &FlowDropWithdrawQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := fp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(flowperiod.Table, flowperiod.FieldID, id),
			sqlgraph.To(flowdropwithdraw.Table, flowdropwithdraw.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, flowperiod.FlowDropWithdrawListTable, flowperiod.FlowDropWithdrawListColumn),
		)
		fromV = sqlgraph.Neighbors(fp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FlowPeriodClient) Hooks() []Hook {
	return c.hooks.FlowPeriod
}

// FlowReasonClient is a client for the FlowReason schema.
type FlowReasonClient struct {
	config
}

// NewFlowReasonClient returns a client for the FlowReason from the given config.
func NewFlowReasonClient(c config) *FlowReasonClient {
	return &FlowReasonClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `flowreason.Hooks(f(g(h())))`.
func (c *FlowReasonClient) Use(hooks ...Hook) {
	c.hooks.FlowReason = append(c.hooks.FlowReason, hooks...)
}

// Create returns a create builder for FlowReason.
func (c *FlowReasonClient) Create() *FlowReasonCreate {
	mutation := newFlowReasonMutation(c.config, OpCreate)
	return &FlowReasonCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FlowReason entities.
func (c *FlowReasonClient) CreateBulk(builders ...*FlowReasonCreate) *FlowReasonCreateBulk {
	return &FlowReasonCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FlowReason.
func (c *FlowReasonClient) Update() *FlowReasonUpdate {
	mutation := newFlowReasonMutation(c.config, OpUpdate)
	return &FlowReasonUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FlowReasonClient) UpdateOne(fr *FlowReason) *FlowReasonUpdateOne {
	mutation := newFlowReasonMutation(c.config, OpUpdateOne, withFlowReason(fr))
	return &FlowReasonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FlowReasonClient) UpdateOneID(id int) *FlowReasonUpdateOne {
	mutation := newFlowReasonMutation(c.config, OpUpdateOne, withFlowReasonID(id))
	return &FlowReasonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FlowReason.
func (c *FlowReasonClient) Delete() *FlowReasonDelete {
	mutation := newFlowReasonMutation(c.config, OpDelete)
	return &FlowReasonDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FlowReasonClient) DeleteOne(fr *FlowReason) *FlowReasonDeleteOne {
	return c.DeleteOneID(fr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FlowReasonClient) DeleteOneID(id int) *FlowReasonDeleteOne {
	builder := c.Delete().Where(flowreason.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FlowReasonDeleteOne{builder}
}

// Query returns a query builder for FlowReason.
func (c *FlowReasonClient) Query() *FlowReasonQuery {
	return &FlowReasonQuery{
		config: c.config,
	}
}

// Get returns a FlowReason entity by its id.
func (c *FlowReasonClient) Get(ctx context.Context, id int) (*FlowReason, error) {
	return c.Query().Where(flowreason.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FlowReasonClient) GetX(ctx context.Context, id int) *FlowReason {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a FlowReason.
func (c *FlowReasonClient) QueryParent(fr *FlowReason) *FlowDropWithdrawQuery {
	query := &FlowDropWithdrawQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := fr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(flowreason.Table, flowreason.FieldID, id),
			sqlgraph.To(flowdropwithdraw.Table, flowdropwithdraw.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flowreason.ParentTable, flowreason.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(fr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FlowReasonClient) Hooks() []Hook {
	return c.hooks.FlowReason
}

// ImmunocompromisedGroupsClient is a client for the ImmunocompromisedGroups schema.
type ImmunocompromisedGroupsClient struct {
	config
}

// NewImmunocompromisedGroupsClient returns a client for the ImmunocompromisedGroups from the given config.
func NewImmunocompromisedGroupsClient(c config) *ImmunocompromisedGroupsClient {
	return &ImmunocompromisedGroupsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `immunocompromisedgroups.Hooks(f(g(h())))`.
func (c *ImmunocompromisedGroupsClient) Use(hooks ...Hook) {
	c.hooks.ImmunocompromisedGroups = append(c.hooks.ImmunocompromisedGroups, hooks...)
}

// Create returns a create builder for ImmunocompromisedGroups.
func (c *ImmunocompromisedGroupsClient) Create() *ImmunocompromisedGroupsCreate {
	mutation := newImmunocompromisedGroupsMutation(c.config, OpCreate)
	return &ImmunocompromisedGroupsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ImmunocompromisedGroups entities.
func (c *ImmunocompromisedGroupsClient) CreateBulk(builders ...*ImmunocompromisedGroupsCreate) *ImmunocompromisedGroupsCreateBulk {
	return &ImmunocompromisedGroupsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ImmunocompromisedGroups.
func (c *ImmunocompromisedGroupsClient) Update() *ImmunocompromisedGroupsUpdate {
	mutation := newImmunocompromisedGroupsMutation(c.config, OpUpdate)
	return &ImmunocompromisedGroupsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ImmunocompromisedGroupsClient) UpdateOne(ig *ImmunocompromisedGroups) *ImmunocompromisedGroupsUpdateOne {
	mutation := newImmunocompromisedGroupsMutation(c.config, OpUpdateOne, withImmunocompromisedGroups(ig))
	return &ImmunocompromisedGroupsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ImmunocompromisedGroupsClient) UpdateOneID(id int) *ImmunocompromisedGroupsUpdateOne {
	mutation := newImmunocompromisedGroupsMutation(c.config, OpUpdateOne, withImmunocompromisedGroupsID(id))
	return &ImmunocompromisedGroupsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ImmunocompromisedGroups.
func (c *ImmunocompromisedGroupsClient) Delete() *ImmunocompromisedGroupsDelete {
	mutation := newImmunocompromisedGroupsMutation(c.config, OpDelete)
	return &ImmunocompromisedGroupsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ImmunocompromisedGroupsClient) DeleteOne(ig *ImmunocompromisedGroups) *ImmunocompromisedGroupsDeleteOne {
	return c.DeleteOneID(ig.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ImmunocompromisedGroupsClient) DeleteOneID(id int) *ImmunocompromisedGroupsDeleteOne {
	builder := c.Delete().Where(immunocompromisedgroups.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ImmunocompromisedGroupsDeleteOne{builder}
}

// Query returns a query builder for ImmunocompromisedGroups.
func (c *ImmunocompromisedGroupsClient) Query() *ImmunocompromisedGroupsQuery {
	return &ImmunocompromisedGroupsQuery{
		config: c.config,
	}
}

// Get returns a ImmunocompromisedGroups entity by its id.
func (c *ImmunocompromisedGroupsClient) Get(ctx context.Context, id int) (*ImmunocompromisedGroups, error) {
	return c.Query().Where(immunocompromisedgroups.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ImmunocompromisedGroupsClient) GetX(ctx context.Context, id int) *ImmunocompromisedGroups {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ImmunocompromisedGroupsClient) Hooks() []Hook {
	return c.hooks.ImmunocompromisedGroups
}

// ManufacturerClient is a client for the Manufacturer schema.
type ManufacturerClient struct {
	config
}

// NewManufacturerClient returns a client for the Manufacturer from the given config.
func NewManufacturerClient(c config) *ManufacturerClient {
	return &ManufacturerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `manufacturer.Hooks(f(g(h())))`.
func (c *ManufacturerClient) Use(hooks ...Hook) {
	c.hooks.Manufacturer = append(c.hooks.Manufacturer, hooks...)
}

// Create returns a create builder for Manufacturer.
func (c *ManufacturerClient) Create() *ManufacturerCreate {
	mutation := newManufacturerMutation(c.config, OpCreate)
	return &ManufacturerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Manufacturer entities.
func (c *ManufacturerClient) CreateBulk(builders ...*ManufacturerCreate) *ManufacturerCreateBulk {
	return &ManufacturerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Manufacturer.
func (c *ManufacturerClient) Update() *ManufacturerUpdate {
	mutation := newManufacturerMutation(c.config, OpUpdate)
	return &ManufacturerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ManufacturerClient) UpdateOne(m *Manufacturer) *ManufacturerUpdateOne {
	mutation := newManufacturerMutation(c.config, OpUpdateOne, withManufacturer(m))
	return &ManufacturerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ManufacturerClient) UpdateOneID(id int) *ManufacturerUpdateOne {
	mutation := newManufacturerMutation(c.config, OpUpdateOne, withManufacturerID(id))
	return &ManufacturerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Manufacturer.
func (c *ManufacturerClient) Delete() *ManufacturerDelete {
	mutation := newManufacturerMutation(c.config, OpDelete)
	return &ManufacturerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ManufacturerClient) DeleteOne(m *Manufacturer) *ManufacturerDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ManufacturerClient) DeleteOneID(id int) *ManufacturerDeleteOne {
	builder := c.Delete().Where(manufacturer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ManufacturerDeleteOne{builder}
}

// Query returns a query builder for Manufacturer.
func (c *ManufacturerClient) Query() *ManufacturerQuery {
	return &ManufacturerQuery{
		config: c.config,
	}
}

// Get returns a Manufacturer entity by its id.
func (c *ManufacturerClient) Get(ctx context.Context, id int) (*Manufacturer, error) {
	return c.Query().Where(manufacturer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ManufacturerClient) GetX(ctx context.Context, id int) *Manufacturer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ManufacturerClient) Hooks() []Hook {
	return c.hooks.Manufacturer
}

// MoreInfoModuleClient is a client for the MoreInfoModule schema.
type MoreInfoModuleClient struct {
	config
}

// NewMoreInfoModuleClient returns a client for the MoreInfoModule from the given config.
func NewMoreInfoModuleClient(c config) *MoreInfoModuleClient {
	return &MoreInfoModuleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `moreinfomodule.Hooks(f(g(h())))`.
func (c *MoreInfoModuleClient) Use(hooks ...Hook) {
	c.hooks.MoreInfoModule = append(c.hooks.MoreInfoModule, hooks...)
}

// Create returns a create builder for MoreInfoModule.
func (c *MoreInfoModuleClient) Create() *MoreInfoModuleCreate {
	mutation := newMoreInfoModuleMutation(c.config, OpCreate)
	return &MoreInfoModuleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MoreInfoModule entities.
func (c *MoreInfoModuleClient) CreateBulk(builders ...*MoreInfoModuleCreate) *MoreInfoModuleCreateBulk {
	return &MoreInfoModuleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MoreInfoModule.
func (c *MoreInfoModuleClient) Update() *MoreInfoModuleUpdate {
	mutation := newMoreInfoModuleMutation(c.config, OpUpdate)
	return &MoreInfoModuleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MoreInfoModuleClient) UpdateOne(mim *MoreInfoModule) *MoreInfoModuleUpdateOne {
	mutation := newMoreInfoModuleMutation(c.config, OpUpdateOne, withMoreInfoModule(mim))
	return &MoreInfoModuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MoreInfoModuleClient) UpdateOneID(id int) *MoreInfoModuleUpdateOne {
	mutation := newMoreInfoModuleMutation(c.config, OpUpdateOne, withMoreInfoModuleID(id))
	return &MoreInfoModuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MoreInfoModule.
func (c *MoreInfoModuleClient) Delete() *MoreInfoModuleDelete {
	mutation := newMoreInfoModuleMutation(c.config, OpDelete)
	return &MoreInfoModuleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MoreInfoModuleClient) DeleteOne(mim *MoreInfoModule) *MoreInfoModuleDeleteOne {
	return c.DeleteOneID(mim.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MoreInfoModuleClient) DeleteOneID(id int) *MoreInfoModuleDeleteOne {
	builder := c.Delete().Where(moreinfomodule.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MoreInfoModuleDeleteOne{builder}
}

// Query returns a query builder for MoreInfoModule.
func (c *MoreInfoModuleClient) Query() *MoreInfoModuleQuery {
	return &MoreInfoModuleQuery{
		config: c.config,
	}
}

// Get returns a MoreInfoModule entity by its id.
func (c *MoreInfoModuleClient) Get(ctx context.Context, id int) (*MoreInfoModule, error) {
	return c.Query().Where(moreinfomodule.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MoreInfoModuleClient) GetX(ctx context.Context, id int) *MoreInfoModule {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a MoreInfoModule.
func (c *MoreInfoModuleClient) QueryParent(mim *MoreInfoModule) *ResultsDefinitionQuery {
	query := &ResultsDefinitionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := mim.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(moreinfomodule.Table, moreinfomodule.FieldID, id),
			sqlgraph.To(resultsdefinition.Table, resultsdefinition.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, moreinfomodule.ParentTable, moreinfomodule.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(mim.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCertainAgreement queries the certain_agreement edge of a MoreInfoModule.
func (c *MoreInfoModuleClient) QueryCertainAgreement(mim *MoreInfoModule) *CertainAgreementQuery {
	query := &CertainAgreementQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := mim.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(moreinfomodule.Table, moreinfomodule.FieldID, id),
			sqlgraph.To(certainagreement.Table, certainagreement.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, moreinfomodule.CertainAgreementTable, moreinfomodule.CertainAgreementColumn),
		)
		fromV = sqlgraph.Neighbors(mim.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPointOfContact queries the point_of_contact edge of a MoreInfoModule.
func (c *MoreInfoModuleClient) QueryPointOfContact(mim *MoreInfoModule) *PointOfContactQuery {
	query := &PointOfContactQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := mim.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(moreinfomodule.Table, moreinfomodule.FieldID, id),
			sqlgraph.To(pointofcontact.Table, pointofcontact.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, moreinfomodule.PointOfContactTable, moreinfomodule.PointOfContactColumn),
		)
		fromV = sqlgraph.Neighbors(mim.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MoreInfoModuleClient) Hooks() []Hook {
	return c.hooks.MoreInfoModule
}

// OtherEventClient is a client for the OtherEvent schema.
type OtherEventClient struct {
	config
}

// NewOtherEventClient returns a client for the OtherEvent from the given config.
func NewOtherEventClient(c config) *OtherEventClient {
	return &OtherEventClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `otherevent.Hooks(f(g(h())))`.
func (c *OtherEventClient) Use(hooks ...Hook) {
	c.hooks.OtherEvent = append(c.hooks.OtherEvent, hooks...)
}

// Create returns a create builder for OtherEvent.
func (c *OtherEventClient) Create() *OtherEventCreate {
	mutation := newOtherEventMutation(c.config, OpCreate)
	return &OtherEventCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OtherEvent entities.
func (c *OtherEventClient) CreateBulk(builders ...*OtherEventCreate) *OtherEventCreateBulk {
	return &OtherEventCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OtherEvent.
func (c *OtherEventClient) Update() *OtherEventUpdate {
	mutation := newOtherEventMutation(c.config, OpUpdate)
	return &OtherEventUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OtherEventClient) UpdateOne(oe *OtherEvent) *OtherEventUpdateOne {
	mutation := newOtherEventMutation(c.config, OpUpdateOne, withOtherEvent(oe))
	return &OtherEventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OtherEventClient) UpdateOneID(id int) *OtherEventUpdateOne {
	mutation := newOtherEventMutation(c.config, OpUpdateOne, withOtherEventID(id))
	return &OtherEventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OtherEvent.
func (c *OtherEventClient) Delete() *OtherEventDelete {
	mutation := newOtherEventMutation(c.config, OpDelete)
	return &OtherEventDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OtherEventClient) DeleteOne(oe *OtherEvent) *OtherEventDeleteOne {
	return c.DeleteOneID(oe.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OtherEventClient) DeleteOneID(id int) *OtherEventDeleteOne {
	builder := c.Delete().Where(otherevent.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OtherEventDeleteOne{builder}
}

// Query returns a query builder for OtherEvent.
func (c *OtherEventClient) Query() *OtherEventQuery {
	return &OtherEventQuery{
		config: c.config,
	}
}

// Get returns a OtherEvent entity by its id.
func (c *OtherEventClient) Get(ctx context.Context, id int) (*OtherEvent, error) {
	return c.Query().Where(otherevent.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OtherEventClient) GetX(ctx context.Context, id int) *OtherEvent {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OtherEvent.
func (c *OtherEventClient) QueryParent(oe *OtherEvent) *AdverseEventsModuleQuery {
	query := &AdverseEventsModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := oe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(otherevent.Table, otherevent.FieldID, id),
			sqlgraph.To(adverseeventsmodule.Table, adverseeventsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, otherevent.ParentTable, otherevent.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(oe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOtherEventStatsList queries the other_event_stats_list edge of a OtherEvent.
func (c *OtherEventClient) QueryOtherEventStatsList(oe *OtherEvent) *OtherEventStatsQuery {
	query := &OtherEventStatsQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := oe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(otherevent.Table, otherevent.FieldID, id),
			sqlgraph.To(othereventstats.Table, othereventstats.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, otherevent.OtherEventStatsListTable, otherevent.OtherEventStatsListColumn),
		)
		fromV = sqlgraph.Neighbors(oe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OtherEventClient) Hooks() []Hook {
	return c.hooks.OtherEvent
}

// OtherEventStatsClient is a client for the OtherEventStats schema.
type OtherEventStatsClient struct {
	config
}

// NewOtherEventStatsClient returns a client for the OtherEventStats from the given config.
func NewOtherEventStatsClient(c config) *OtherEventStatsClient {
	return &OtherEventStatsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `othereventstats.Hooks(f(g(h())))`.
func (c *OtherEventStatsClient) Use(hooks ...Hook) {
	c.hooks.OtherEventStats = append(c.hooks.OtherEventStats, hooks...)
}

// Create returns a create builder for OtherEventStats.
func (c *OtherEventStatsClient) Create() *OtherEventStatsCreate {
	mutation := newOtherEventStatsMutation(c.config, OpCreate)
	return &OtherEventStatsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OtherEventStats entities.
func (c *OtherEventStatsClient) CreateBulk(builders ...*OtherEventStatsCreate) *OtherEventStatsCreateBulk {
	return &OtherEventStatsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OtherEventStats.
func (c *OtherEventStatsClient) Update() *OtherEventStatsUpdate {
	mutation := newOtherEventStatsMutation(c.config, OpUpdate)
	return &OtherEventStatsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OtherEventStatsClient) UpdateOne(oes *OtherEventStats) *OtherEventStatsUpdateOne {
	mutation := newOtherEventStatsMutation(c.config, OpUpdateOne, withOtherEventStats(oes))
	return &OtherEventStatsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OtherEventStatsClient) UpdateOneID(id int) *OtherEventStatsUpdateOne {
	mutation := newOtherEventStatsMutation(c.config, OpUpdateOne, withOtherEventStatsID(id))
	return &OtherEventStatsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OtherEventStats.
func (c *OtherEventStatsClient) Delete() *OtherEventStatsDelete {
	mutation := newOtherEventStatsMutation(c.config, OpDelete)
	return &OtherEventStatsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OtherEventStatsClient) DeleteOne(oes *OtherEventStats) *OtherEventStatsDeleteOne {
	return c.DeleteOneID(oes.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OtherEventStatsClient) DeleteOneID(id int) *OtherEventStatsDeleteOne {
	builder := c.Delete().Where(othereventstats.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OtherEventStatsDeleteOne{builder}
}

// Query returns a query builder for OtherEventStats.
func (c *OtherEventStatsClient) Query() *OtherEventStatsQuery {
	return &OtherEventStatsQuery{
		config: c.config,
	}
}

// Get returns a OtherEventStats entity by its id.
func (c *OtherEventStatsClient) Get(ctx context.Context, id int) (*OtherEventStats, error) {
	return c.Query().Where(othereventstats.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OtherEventStatsClient) GetX(ctx context.Context, id int) *OtherEventStats {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OtherEventStats.
func (c *OtherEventStatsClient) QueryParent(oes *OtherEventStats) *OtherEventQuery {
	query := &OtherEventQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := oes.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(othereventstats.Table, othereventstats.FieldID, id),
			sqlgraph.To(otherevent.Table, otherevent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, othereventstats.ParentTable, othereventstats.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(oes.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OtherEventStatsClient) Hooks() []Hook {
	return c.hooks.OtherEventStats
}

// OutcomeAnalysisClient is a client for the OutcomeAnalysis schema.
type OutcomeAnalysisClient struct {
	config
}

// NewOutcomeAnalysisClient returns a client for the OutcomeAnalysis from the given config.
func NewOutcomeAnalysisClient(c config) *OutcomeAnalysisClient {
	return &OutcomeAnalysisClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outcomeanalysis.Hooks(f(g(h())))`.
func (c *OutcomeAnalysisClient) Use(hooks ...Hook) {
	c.hooks.OutcomeAnalysis = append(c.hooks.OutcomeAnalysis, hooks...)
}

// Create returns a create builder for OutcomeAnalysis.
func (c *OutcomeAnalysisClient) Create() *OutcomeAnalysisCreate {
	mutation := newOutcomeAnalysisMutation(c.config, OpCreate)
	return &OutcomeAnalysisCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutcomeAnalysis entities.
func (c *OutcomeAnalysisClient) CreateBulk(builders ...*OutcomeAnalysisCreate) *OutcomeAnalysisCreateBulk {
	return &OutcomeAnalysisCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutcomeAnalysis.
func (c *OutcomeAnalysisClient) Update() *OutcomeAnalysisUpdate {
	mutation := newOutcomeAnalysisMutation(c.config, OpUpdate)
	return &OutcomeAnalysisUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutcomeAnalysisClient) UpdateOne(oa *OutcomeAnalysis) *OutcomeAnalysisUpdateOne {
	mutation := newOutcomeAnalysisMutation(c.config, OpUpdateOne, withOutcomeAnalysis(oa))
	return &OutcomeAnalysisUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutcomeAnalysisClient) UpdateOneID(id int) *OutcomeAnalysisUpdateOne {
	mutation := newOutcomeAnalysisMutation(c.config, OpUpdateOne, withOutcomeAnalysisID(id))
	return &OutcomeAnalysisUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutcomeAnalysis.
func (c *OutcomeAnalysisClient) Delete() *OutcomeAnalysisDelete {
	mutation := newOutcomeAnalysisMutation(c.config, OpDelete)
	return &OutcomeAnalysisDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutcomeAnalysisClient) DeleteOne(oa *OutcomeAnalysis) *OutcomeAnalysisDeleteOne {
	return c.DeleteOneID(oa.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutcomeAnalysisClient) DeleteOneID(id int) *OutcomeAnalysisDeleteOne {
	builder := c.Delete().Where(outcomeanalysis.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutcomeAnalysisDeleteOne{builder}
}

// Query returns a query builder for OutcomeAnalysis.
func (c *OutcomeAnalysisClient) Query() *OutcomeAnalysisQuery {
	return &OutcomeAnalysisQuery{
		config: c.config,
	}
}

// Get returns a OutcomeAnalysis entity by its id.
func (c *OutcomeAnalysisClient) Get(ctx context.Context, id int) (*OutcomeAnalysis, error) {
	return c.Query().Where(outcomeanalysis.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutcomeAnalysisClient) GetX(ctx context.Context, id int) *OutcomeAnalysis {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OutcomeAnalysis.
func (c *OutcomeAnalysisClient) QueryParent(oa *OutcomeAnalysis) *OutcomeMeasureQuery {
	query := &OutcomeMeasureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := oa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeanalysis.Table, outcomeanalysis.FieldID, id),
			sqlgraph.To(outcomemeasure.Table, outcomemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomeanalysis.ParentTable, outcomeanalysis.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(oa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOutcomeAnalysisGroupIDList queries the outcome_analysis_group_id_list edge of a OutcomeAnalysis.
func (c *OutcomeAnalysisClient) QueryOutcomeAnalysisGroupIDList(oa *OutcomeAnalysis) *OutcomeAnalysisGroupIDQuery {
	query := &OutcomeAnalysisGroupIDQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := oa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeanalysis.Table, outcomeanalysis.FieldID, id),
			sqlgraph.To(outcomeanalysisgroupid.Table, outcomeanalysisgroupid.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomeanalysis.OutcomeAnalysisGroupIDListTable, outcomeanalysis.OutcomeAnalysisGroupIDListColumn),
		)
		fromV = sqlgraph.Neighbors(oa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutcomeAnalysisClient) Hooks() []Hook {
	return c.hooks.OutcomeAnalysis
}

// OutcomeAnalysisGroupIDClient is a client for the OutcomeAnalysisGroupID schema.
type OutcomeAnalysisGroupIDClient struct {
	config
}

// NewOutcomeAnalysisGroupIDClient returns a client for the OutcomeAnalysisGroupID from the given config.
func NewOutcomeAnalysisGroupIDClient(c config) *OutcomeAnalysisGroupIDClient {
	return &OutcomeAnalysisGroupIDClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outcomeanalysisgroupid.Hooks(f(g(h())))`.
func (c *OutcomeAnalysisGroupIDClient) Use(hooks ...Hook) {
	c.hooks.OutcomeAnalysisGroupID = append(c.hooks.OutcomeAnalysisGroupID, hooks...)
}

// Create returns a create builder for OutcomeAnalysisGroupID.
func (c *OutcomeAnalysisGroupIDClient) Create() *OutcomeAnalysisGroupIDCreate {
	mutation := newOutcomeAnalysisGroupIDMutation(c.config, OpCreate)
	return &OutcomeAnalysisGroupIDCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutcomeAnalysisGroupID entities.
func (c *OutcomeAnalysisGroupIDClient) CreateBulk(builders ...*OutcomeAnalysisGroupIDCreate) *OutcomeAnalysisGroupIDCreateBulk {
	return &OutcomeAnalysisGroupIDCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutcomeAnalysisGroupID.
func (c *OutcomeAnalysisGroupIDClient) Update() *OutcomeAnalysisGroupIDUpdate {
	mutation := newOutcomeAnalysisGroupIDMutation(c.config, OpUpdate)
	return &OutcomeAnalysisGroupIDUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutcomeAnalysisGroupIDClient) UpdateOne(oagi *OutcomeAnalysisGroupID) *OutcomeAnalysisGroupIDUpdateOne {
	mutation := newOutcomeAnalysisGroupIDMutation(c.config, OpUpdateOne, withOutcomeAnalysisGroupID(oagi))
	return &OutcomeAnalysisGroupIDUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutcomeAnalysisGroupIDClient) UpdateOneID(id int) *OutcomeAnalysisGroupIDUpdateOne {
	mutation := newOutcomeAnalysisGroupIDMutation(c.config, OpUpdateOne, withOutcomeAnalysisGroupIDID(id))
	return &OutcomeAnalysisGroupIDUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutcomeAnalysisGroupID.
func (c *OutcomeAnalysisGroupIDClient) Delete() *OutcomeAnalysisGroupIDDelete {
	mutation := newOutcomeAnalysisGroupIDMutation(c.config, OpDelete)
	return &OutcomeAnalysisGroupIDDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutcomeAnalysisGroupIDClient) DeleteOne(oagi *OutcomeAnalysisGroupID) *OutcomeAnalysisGroupIDDeleteOne {
	return c.DeleteOneID(oagi.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutcomeAnalysisGroupIDClient) DeleteOneID(id int) *OutcomeAnalysisGroupIDDeleteOne {
	builder := c.Delete().Where(outcomeanalysisgroupid.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutcomeAnalysisGroupIDDeleteOne{builder}
}

// Query returns a query builder for OutcomeAnalysisGroupID.
func (c *OutcomeAnalysisGroupIDClient) Query() *OutcomeAnalysisGroupIDQuery {
	return &OutcomeAnalysisGroupIDQuery{
		config: c.config,
	}
}

// Get returns a OutcomeAnalysisGroupID entity by its id.
func (c *OutcomeAnalysisGroupIDClient) Get(ctx context.Context, id int) (*OutcomeAnalysisGroupID, error) {
	return c.Query().Where(outcomeanalysisgroupid.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutcomeAnalysisGroupIDClient) GetX(ctx context.Context, id int) *OutcomeAnalysisGroupID {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OutcomeAnalysisGroupID.
func (c *OutcomeAnalysisGroupIDClient) QueryParent(oagi *OutcomeAnalysisGroupID) *OutcomeAnalysisQuery {
	query := &OutcomeAnalysisQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := oagi.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeanalysisgroupid.Table, outcomeanalysisgroupid.FieldID, id),
			sqlgraph.To(outcomeanalysis.Table, outcomeanalysis.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomeanalysisgroupid.ParentTable, outcomeanalysisgroupid.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(oagi.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutcomeAnalysisGroupIDClient) Hooks() []Hook {
	return c.hooks.OutcomeAnalysisGroupID
}

// OutcomeCategoryClient is a client for the OutcomeCategory schema.
type OutcomeCategoryClient struct {
	config
}

// NewOutcomeCategoryClient returns a client for the OutcomeCategory from the given config.
func NewOutcomeCategoryClient(c config) *OutcomeCategoryClient {
	return &OutcomeCategoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outcomecategory.Hooks(f(g(h())))`.
func (c *OutcomeCategoryClient) Use(hooks ...Hook) {
	c.hooks.OutcomeCategory = append(c.hooks.OutcomeCategory, hooks...)
}

// Create returns a create builder for OutcomeCategory.
func (c *OutcomeCategoryClient) Create() *OutcomeCategoryCreate {
	mutation := newOutcomeCategoryMutation(c.config, OpCreate)
	return &OutcomeCategoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutcomeCategory entities.
func (c *OutcomeCategoryClient) CreateBulk(builders ...*OutcomeCategoryCreate) *OutcomeCategoryCreateBulk {
	return &OutcomeCategoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutcomeCategory.
func (c *OutcomeCategoryClient) Update() *OutcomeCategoryUpdate {
	mutation := newOutcomeCategoryMutation(c.config, OpUpdate)
	return &OutcomeCategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutcomeCategoryClient) UpdateOne(oc *OutcomeCategory) *OutcomeCategoryUpdateOne {
	mutation := newOutcomeCategoryMutation(c.config, OpUpdateOne, withOutcomeCategory(oc))
	return &OutcomeCategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutcomeCategoryClient) UpdateOneID(id int) *OutcomeCategoryUpdateOne {
	mutation := newOutcomeCategoryMutation(c.config, OpUpdateOne, withOutcomeCategoryID(id))
	return &OutcomeCategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutcomeCategory.
func (c *OutcomeCategoryClient) Delete() *OutcomeCategoryDelete {
	mutation := newOutcomeCategoryMutation(c.config, OpDelete)
	return &OutcomeCategoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutcomeCategoryClient) DeleteOne(oc *OutcomeCategory) *OutcomeCategoryDeleteOne {
	return c.DeleteOneID(oc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutcomeCategoryClient) DeleteOneID(id int) *OutcomeCategoryDeleteOne {
	builder := c.Delete().Where(outcomecategory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutcomeCategoryDeleteOne{builder}
}

// Query returns a query builder for OutcomeCategory.
func (c *OutcomeCategoryClient) Query() *OutcomeCategoryQuery {
	return &OutcomeCategoryQuery{
		config: c.config,
	}
}

// Get returns a OutcomeCategory entity by its id.
func (c *OutcomeCategoryClient) Get(ctx context.Context, id int) (*OutcomeCategory, error) {
	return c.Query().Where(outcomecategory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutcomeCategoryClient) GetX(ctx context.Context, id int) *OutcomeCategory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OutcomeCategory.
func (c *OutcomeCategoryClient) QueryParent(oc *OutcomeCategory) *OutcomeClassQuery {
	query := &OutcomeClassQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := oc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomecategory.Table, outcomecategory.FieldID, id),
			sqlgraph.To(outcomeclass.Table, outcomeclass.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomecategory.ParentTable, outcomecategory.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(oc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOutcomeMeasurementList queries the outcome_measurement_list edge of a OutcomeCategory.
func (c *OutcomeCategoryClient) QueryOutcomeMeasurementList(oc *OutcomeCategory) *OutcomeMeasurementQuery {
	query := &OutcomeMeasurementQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := oc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomecategory.Table, outcomecategory.FieldID, id),
			sqlgraph.To(outcomemeasurement.Table, outcomemeasurement.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomecategory.OutcomeMeasurementListTable, outcomecategory.OutcomeMeasurementListColumn),
		)
		fromV = sqlgraph.Neighbors(oc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutcomeCategoryClient) Hooks() []Hook {
	return c.hooks.OutcomeCategory
}

// OutcomeClassClient is a client for the OutcomeClass schema.
type OutcomeClassClient struct {
	config
}

// NewOutcomeClassClient returns a client for the OutcomeClass from the given config.
func NewOutcomeClassClient(c config) *OutcomeClassClient {
	return &OutcomeClassClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outcomeclass.Hooks(f(g(h())))`.
func (c *OutcomeClassClient) Use(hooks ...Hook) {
	c.hooks.OutcomeClass = append(c.hooks.OutcomeClass, hooks...)
}

// Create returns a create builder for OutcomeClass.
func (c *OutcomeClassClient) Create() *OutcomeClassCreate {
	mutation := newOutcomeClassMutation(c.config, OpCreate)
	return &OutcomeClassCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutcomeClass entities.
func (c *OutcomeClassClient) CreateBulk(builders ...*OutcomeClassCreate) *OutcomeClassCreateBulk {
	return &OutcomeClassCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutcomeClass.
func (c *OutcomeClassClient) Update() *OutcomeClassUpdate {
	mutation := newOutcomeClassMutation(c.config, OpUpdate)
	return &OutcomeClassUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutcomeClassClient) UpdateOne(oc *OutcomeClass) *OutcomeClassUpdateOne {
	mutation := newOutcomeClassMutation(c.config, OpUpdateOne, withOutcomeClass(oc))
	return &OutcomeClassUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutcomeClassClient) UpdateOneID(id int) *OutcomeClassUpdateOne {
	mutation := newOutcomeClassMutation(c.config, OpUpdateOne, withOutcomeClassID(id))
	return &OutcomeClassUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutcomeClass.
func (c *OutcomeClassClient) Delete() *OutcomeClassDelete {
	mutation := newOutcomeClassMutation(c.config, OpDelete)
	return &OutcomeClassDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutcomeClassClient) DeleteOne(oc *OutcomeClass) *OutcomeClassDeleteOne {
	return c.DeleteOneID(oc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutcomeClassClient) DeleteOneID(id int) *OutcomeClassDeleteOne {
	builder := c.Delete().Where(outcomeclass.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutcomeClassDeleteOne{builder}
}

// Query returns a query builder for OutcomeClass.
func (c *OutcomeClassClient) Query() *OutcomeClassQuery {
	return &OutcomeClassQuery{
		config: c.config,
	}
}

// Get returns a OutcomeClass entity by its id.
func (c *OutcomeClassClient) Get(ctx context.Context, id int) (*OutcomeClass, error) {
	return c.Query().Where(outcomeclass.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutcomeClassClient) GetX(ctx context.Context, id int) *OutcomeClass {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OutcomeClass.
func (c *OutcomeClassClient) QueryParent(oc *OutcomeClass) *OutcomeMeasureQuery {
	query := &OutcomeMeasureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := oc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeclass.Table, outcomeclass.FieldID, id),
			sqlgraph.To(outcomemeasure.Table, outcomemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomeclass.ParentTable, outcomeclass.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(oc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOutcomeClassDenomList queries the outcome_class_denom_list edge of a OutcomeClass.
func (c *OutcomeClassClient) QueryOutcomeClassDenomList(oc *OutcomeClass) *OutcomeClassDenomQuery {
	query := &OutcomeClassDenomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := oc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeclass.Table, outcomeclass.FieldID, id),
			sqlgraph.To(outcomeclassdenom.Table, outcomeclassdenom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomeclass.OutcomeClassDenomListTable, outcomeclass.OutcomeClassDenomListColumn),
		)
		fromV = sqlgraph.Neighbors(oc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOutcomeCategoryList queries the outcome_category_list edge of a OutcomeClass.
func (c *OutcomeClassClient) QueryOutcomeCategoryList(oc *OutcomeClass) *OutcomeCategoryQuery {
	query := &OutcomeCategoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := oc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeclass.Table, outcomeclass.FieldID, id),
			sqlgraph.To(outcomecategory.Table, outcomecategory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomeclass.OutcomeCategoryListTable, outcomeclass.OutcomeCategoryListColumn),
		)
		fromV = sqlgraph.Neighbors(oc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutcomeClassClient) Hooks() []Hook {
	return c.hooks.OutcomeClass
}

// OutcomeClassDenomClient is a client for the OutcomeClassDenom schema.
type OutcomeClassDenomClient struct {
	config
}

// NewOutcomeClassDenomClient returns a client for the OutcomeClassDenom from the given config.
func NewOutcomeClassDenomClient(c config) *OutcomeClassDenomClient {
	return &OutcomeClassDenomClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outcomeclassdenom.Hooks(f(g(h())))`.
func (c *OutcomeClassDenomClient) Use(hooks ...Hook) {
	c.hooks.OutcomeClassDenom = append(c.hooks.OutcomeClassDenom, hooks...)
}

// Create returns a create builder for OutcomeClassDenom.
func (c *OutcomeClassDenomClient) Create() *OutcomeClassDenomCreate {
	mutation := newOutcomeClassDenomMutation(c.config, OpCreate)
	return &OutcomeClassDenomCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutcomeClassDenom entities.
func (c *OutcomeClassDenomClient) CreateBulk(builders ...*OutcomeClassDenomCreate) *OutcomeClassDenomCreateBulk {
	return &OutcomeClassDenomCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutcomeClassDenom.
func (c *OutcomeClassDenomClient) Update() *OutcomeClassDenomUpdate {
	mutation := newOutcomeClassDenomMutation(c.config, OpUpdate)
	return &OutcomeClassDenomUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutcomeClassDenomClient) UpdateOne(ocd *OutcomeClassDenom) *OutcomeClassDenomUpdateOne {
	mutation := newOutcomeClassDenomMutation(c.config, OpUpdateOne, withOutcomeClassDenom(ocd))
	return &OutcomeClassDenomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutcomeClassDenomClient) UpdateOneID(id int) *OutcomeClassDenomUpdateOne {
	mutation := newOutcomeClassDenomMutation(c.config, OpUpdateOne, withOutcomeClassDenomID(id))
	return &OutcomeClassDenomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutcomeClassDenom.
func (c *OutcomeClassDenomClient) Delete() *OutcomeClassDenomDelete {
	mutation := newOutcomeClassDenomMutation(c.config, OpDelete)
	return &OutcomeClassDenomDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutcomeClassDenomClient) DeleteOne(ocd *OutcomeClassDenom) *OutcomeClassDenomDeleteOne {
	return c.DeleteOneID(ocd.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutcomeClassDenomClient) DeleteOneID(id int) *OutcomeClassDenomDeleteOne {
	builder := c.Delete().Where(outcomeclassdenom.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutcomeClassDenomDeleteOne{builder}
}

// Query returns a query builder for OutcomeClassDenom.
func (c *OutcomeClassDenomClient) Query() *OutcomeClassDenomQuery {
	return &OutcomeClassDenomQuery{
		config: c.config,
	}
}

// Get returns a OutcomeClassDenom entity by its id.
func (c *OutcomeClassDenomClient) Get(ctx context.Context, id int) (*OutcomeClassDenom, error) {
	return c.Query().Where(outcomeclassdenom.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutcomeClassDenomClient) GetX(ctx context.Context, id int) *OutcomeClassDenom {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OutcomeClassDenom.
func (c *OutcomeClassDenomClient) QueryParent(ocd *OutcomeClassDenom) *OutcomeClassQuery {
	query := &OutcomeClassQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ocd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeclassdenom.Table, outcomeclassdenom.FieldID, id),
			sqlgraph.To(outcomeclass.Table, outcomeclass.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomeclassdenom.ParentTable, outcomeclassdenom.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(ocd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOutcomeClassDenomCountList queries the outcome_class_denom_count_list edge of a OutcomeClassDenom.
func (c *OutcomeClassDenomClient) QueryOutcomeClassDenomCountList(ocd *OutcomeClassDenom) *OutcomeClassDenomCountQuery {
	query := &OutcomeClassDenomCountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ocd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeclassdenom.Table, outcomeclassdenom.FieldID, id),
			sqlgraph.To(outcomeclassdenomcount.Table, outcomeclassdenomcount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomeclassdenom.OutcomeClassDenomCountListTable, outcomeclassdenom.OutcomeClassDenomCountListColumn),
		)
		fromV = sqlgraph.Neighbors(ocd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutcomeClassDenomClient) Hooks() []Hook {
	return c.hooks.OutcomeClassDenom
}

// OutcomeClassDenomCountClient is a client for the OutcomeClassDenomCount schema.
type OutcomeClassDenomCountClient struct {
	config
}

// NewOutcomeClassDenomCountClient returns a client for the OutcomeClassDenomCount from the given config.
func NewOutcomeClassDenomCountClient(c config) *OutcomeClassDenomCountClient {
	return &OutcomeClassDenomCountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outcomeclassdenomcount.Hooks(f(g(h())))`.
func (c *OutcomeClassDenomCountClient) Use(hooks ...Hook) {
	c.hooks.OutcomeClassDenomCount = append(c.hooks.OutcomeClassDenomCount, hooks...)
}

// Create returns a create builder for OutcomeClassDenomCount.
func (c *OutcomeClassDenomCountClient) Create() *OutcomeClassDenomCountCreate {
	mutation := newOutcomeClassDenomCountMutation(c.config, OpCreate)
	return &OutcomeClassDenomCountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutcomeClassDenomCount entities.
func (c *OutcomeClassDenomCountClient) CreateBulk(builders ...*OutcomeClassDenomCountCreate) *OutcomeClassDenomCountCreateBulk {
	return &OutcomeClassDenomCountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutcomeClassDenomCount.
func (c *OutcomeClassDenomCountClient) Update() *OutcomeClassDenomCountUpdate {
	mutation := newOutcomeClassDenomCountMutation(c.config, OpUpdate)
	return &OutcomeClassDenomCountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutcomeClassDenomCountClient) UpdateOne(ocdc *OutcomeClassDenomCount) *OutcomeClassDenomCountUpdateOne {
	mutation := newOutcomeClassDenomCountMutation(c.config, OpUpdateOne, withOutcomeClassDenomCount(ocdc))
	return &OutcomeClassDenomCountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutcomeClassDenomCountClient) UpdateOneID(id int) *OutcomeClassDenomCountUpdateOne {
	mutation := newOutcomeClassDenomCountMutation(c.config, OpUpdateOne, withOutcomeClassDenomCountID(id))
	return &OutcomeClassDenomCountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutcomeClassDenomCount.
func (c *OutcomeClassDenomCountClient) Delete() *OutcomeClassDenomCountDelete {
	mutation := newOutcomeClassDenomCountMutation(c.config, OpDelete)
	return &OutcomeClassDenomCountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutcomeClassDenomCountClient) DeleteOne(ocdc *OutcomeClassDenomCount) *OutcomeClassDenomCountDeleteOne {
	return c.DeleteOneID(ocdc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutcomeClassDenomCountClient) DeleteOneID(id int) *OutcomeClassDenomCountDeleteOne {
	builder := c.Delete().Where(outcomeclassdenomcount.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutcomeClassDenomCountDeleteOne{builder}
}

// Query returns a query builder for OutcomeClassDenomCount.
func (c *OutcomeClassDenomCountClient) Query() *OutcomeClassDenomCountQuery {
	return &OutcomeClassDenomCountQuery{
		config: c.config,
	}
}

// Get returns a OutcomeClassDenomCount entity by its id.
func (c *OutcomeClassDenomCountClient) Get(ctx context.Context, id int) (*OutcomeClassDenomCount, error) {
	return c.Query().Where(outcomeclassdenomcount.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutcomeClassDenomCountClient) GetX(ctx context.Context, id int) *OutcomeClassDenomCount {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OutcomeClassDenomCount.
func (c *OutcomeClassDenomCountClient) QueryParent(ocdc *OutcomeClassDenomCount) *OutcomeClassDenomQuery {
	query := &OutcomeClassDenomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ocdc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeclassdenomcount.Table, outcomeclassdenomcount.FieldID, id),
			sqlgraph.To(outcomeclassdenom.Table, outcomeclassdenom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomeclassdenomcount.ParentTable, outcomeclassdenomcount.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(ocdc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutcomeClassDenomCountClient) Hooks() []Hook {
	return c.hooks.OutcomeClassDenomCount
}

// OutcomeDenomClient is a client for the OutcomeDenom schema.
type OutcomeDenomClient struct {
	config
}

// NewOutcomeDenomClient returns a client for the OutcomeDenom from the given config.
func NewOutcomeDenomClient(c config) *OutcomeDenomClient {
	return &OutcomeDenomClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outcomedenom.Hooks(f(g(h())))`.
func (c *OutcomeDenomClient) Use(hooks ...Hook) {
	c.hooks.OutcomeDenom = append(c.hooks.OutcomeDenom, hooks...)
}

// Create returns a create builder for OutcomeDenom.
func (c *OutcomeDenomClient) Create() *OutcomeDenomCreate {
	mutation := newOutcomeDenomMutation(c.config, OpCreate)
	return &OutcomeDenomCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutcomeDenom entities.
func (c *OutcomeDenomClient) CreateBulk(builders ...*OutcomeDenomCreate) *OutcomeDenomCreateBulk {
	return &OutcomeDenomCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutcomeDenom.
func (c *OutcomeDenomClient) Update() *OutcomeDenomUpdate {
	mutation := newOutcomeDenomMutation(c.config, OpUpdate)
	return &OutcomeDenomUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutcomeDenomClient) UpdateOne(od *OutcomeDenom) *OutcomeDenomUpdateOne {
	mutation := newOutcomeDenomMutation(c.config, OpUpdateOne, withOutcomeDenom(od))
	return &OutcomeDenomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutcomeDenomClient) UpdateOneID(id int) *OutcomeDenomUpdateOne {
	mutation := newOutcomeDenomMutation(c.config, OpUpdateOne, withOutcomeDenomID(id))
	return &OutcomeDenomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutcomeDenom.
func (c *OutcomeDenomClient) Delete() *OutcomeDenomDelete {
	mutation := newOutcomeDenomMutation(c.config, OpDelete)
	return &OutcomeDenomDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutcomeDenomClient) DeleteOne(od *OutcomeDenom) *OutcomeDenomDeleteOne {
	return c.DeleteOneID(od.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutcomeDenomClient) DeleteOneID(id int) *OutcomeDenomDeleteOne {
	builder := c.Delete().Where(outcomedenom.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutcomeDenomDeleteOne{builder}
}

// Query returns a query builder for OutcomeDenom.
func (c *OutcomeDenomClient) Query() *OutcomeDenomQuery {
	return &OutcomeDenomQuery{
		config: c.config,
	}
}

// Get returns a OutcomeDenom entity by its id.
func (c *OutcomeDenomClient) Get(ctx context.Context, id int) (*OutcomeDenom, error) {
	return c.Query().Where(outcomedenom.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutcomeDenomClient) GetX(ctx context.Context, id int) *OutcomeDenom {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OutcomeDenom.
func (c *OutcomeDenomClient) QueryParent(od *OutcomeDenom) *OutcomeMeasureQuery {
	query := &OutcomeMeasureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := od.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomedenom.Table, outcomedenom.FieldID, id),
			sqlgraph.To(outcomemeasure.Table, outcomemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomedenom.ParentTable, outcomedenom.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(od.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOutcomeDenomCountList queries the outcome_denom_count_list edge of a OutcomeDenom.
func (c *OutcomeDenomClient) QueryOutcomeDenomCountList(od *OutcomeDenom) *OutcomeDenomCountQuery {
	query := &OutcomeDenomCountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := od.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomedenom.Table, outcomedenom.FieldID, id),
			sqlgraph.To(outcomedenomcount.Table, outcomedenomcount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomedenom.OutcomeDenomCountListTable, outcomedenom.OutcomeDenomCountListColumn),
		)
		fromV = sqlgraph.Neighbors(od.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutcomeDenomClient) Hooks() []Hook {
	return c.hooks.OutcomeDenom
}

// OutcomeDenomCountClient is a client for the OutcomeDenomCount schema.
type OutcomeDenomCountClient struct {
	config
}

// NewOutcomeDenomCountClient returns a client for the OutcomeDenomCount from the given config.
func NewOutcomeDenomCountClient(c config) *OutcomeDenomCountClient {
	return &OutcomeDenomCountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outcomedenomcount.Hooks(f(g(h())))`.
func (c *OutcomeDenomCountClient) Use(hooks ...Hook) {
	c.hooks.OutcomeDenomCount = append(c.hooks.OutcomeDenomCount, hooks...)
}

// Create returns a create builder for OutcomeDenomCount.
func (c *OutcomeDenomCountClient) Create() *OutcomeDenomCountCreate {
	mutation := newOutcomeDenomCountMutation(c.config, OpCreate)
	return &OutcomeDenomCountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutcomeDenomCount entities.
func (c *OutcomeDenomCountClient) CreateBulk(builders ...*OutcomeDenomCountCreate) *OutcomeDenomCountCreateBulk {
	return &OutcomeDenomCountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutcomeDenomCount.
func (c *OutcomeDenomCountClient) Update() *OutcomeDenomCountUpdate {
	mutation := newOutcomeDenomCountMutation(c.config, OpUpdate)
	return &OutcomeDenomCountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutcomeDenomCountClient) UpdateOne(odc *OutcomeDenomCount) *OutcomeDenomCountUpdateOne {
	mutation := newOutcomeDenomCountMutation(c.config, OpUpdateOne, withOutcomeDenomCount(odc))
	return &OutcomeDenomCountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutcomeDenomCountClient) UpdateOneID(id int) *OutcomeDenomCountUpdateOne {
	mutation := newOutcomeDenomCountMutation(c.config, OpUpdateOne, withOutcomeDenomCountID(id))
	return &OutcomeDenomCountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutcomeDenomCount.
func (c *OutcomeDenomCountClient) Delete() *OutcomeDenomCountDelete {
	mutation := newOutcomeDenomCountMutation(c.config, OpDelete)
	return &OutcomeDenomCountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutcomeDenomCountClient) DeleteOne(odc *OutcomeDenomCount) *OutcomeDenomCountDeleteOne {
	return c.DeleteOneID(odc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutcomeDenomCountClient) DeleteOneID(id int) *OutcomeDenomCountDeleteOne {
	builder := c.Delete().Where(outcomedenomcount.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutcomeDenomCountDeleteOne{builder}
}

// Query returns a query builder for OutcomeDenomCount.
func (c *OutcomeDenomCountClient) Query() *OutcomeDenomCountQuery {
	return &OutcomeDenomCountQuery{
		config: c.config,
	}
}

// Get returns a OutcomeDenomCount entity by its id.
func (c *OutcomeDenomCountClient) Get(ctx context.Context, id int) (*OutcomeDenomCount, error) {
	return c.Query().Where(outcomedenomcount.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutcomeDenomCountClient) GetX(ctx context.Context, id int) *OutcomeDenomCount {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OutcomeDenomCount.
func (c *OutcomeDenomCountClient) QueryParent(odc *OutcomeDenomCount) *OutcomeDenomQuery {
	query := &OutcomeDenomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := odc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomedenomcount.Table, outcomedenomcount.FieldID, id),
			sqlgraph.To(outcomedenom.Table, outcomedenom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomedenomcount.ParentTable, outcomedenomcount.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(odc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutcomeDenomCountClient) Hooks() []Hook {
	return c.hooks.OutcomeDenomCount
}

// OutcomeGroupClient is a client for the OutcomeGroup schema.
type OutcomeGroupClient struct {
	config
}

// NewOutcomeGroupClient returns a client for the OutcomeGroup from the given config.
func NewOutcomeGroupClient(c config) *OutcomeGroupClient {
	return &OutcomeGroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outcomegroup.Hooks(f(g(h())))`.
func (c *OutcomeGroupClient) Use(hooks ...Hook) {
	c.hooks.OutcomeGroup = append(c.hooks.OutcomeGroup, hooks...)
}

// Create returns a create builder for OutcomeGroup.
func (c *OutcomeGroupClient) Create() *OutcomeGroupCreate {
	mutation := newOutcomeGroupMutation(c.config, OpCreate)
	return &OutcomeGroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutcomeGroup entities.
func (c *OutcomeGroupClient) CreateBulk(builders ...*OutcomeGroupCreate) *OutcomeGroupCreateBulk {
	return &OutcomeGroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutcomeGroup.
func (c *OutcomeGroupClient) Update() *OutcomeGroupUpdate {
	mutation := newOutcomeGroupMutation(c.config, OpUpdate)
	return &OutcomeGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutcomeGroupClient) UpdateOne(og *OutcomeGroup) *OutcomeGroupUpdateOne {
	mutation := newOutcomeGroupMutation(c.config, OpUpdateOne, withOutcomeGroup(og))
	return &OutcomeGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutcomeGroupClient) UpdateOneID(id int) *OutcomeGroupUpdateOne {
	mutation := newOutcomeGroupMutation(c.config, OpUpdateOne, withOutcomeGroupID(id))
	return &OutcomeGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutcomeGroup.
func (c *OutcomeGroupClient) Delete() *OutcomeGroupDelete {
	mutation := newOutcomeGroupMutation(c.config, OpDelete)
	return &OutcomeGroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutcomeGroupClient) DeleteOne(og *OutcomeGroup) *OutcomeGroupDeleteOne {
	return c.DeleteOneID(og.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutcomeGroupClient) DeleteOneID(id int) *OutcomeGroupDeleteOne {
	builder := c.Delete().Where(outcomegroup.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutcomeGroupDeleteOne{builder}
}

// Query returns a query builder for OutcomeGroup.
func (c *OutcomeGroupClient) Query() *OutcomeGroupQuery {
	return &OutcomeGroupQuery{
		config: c.config,
	}
}

// Get returns a OutcomeGroup entity by its id.
func (c *OutcomeGroupClient) Get(ctx context.Context, id int) (*OutcomeGroup, error) {
	return c.Query().Where(outcomegroup.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutcomeGroupClient) GetX(ctx context.Context, id int) *OutcomeGroup {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OutcomeGroup.
func (c *OutcomeGroupClient) QueryParent(og *OutcomeGroup) *OutcomeMeasureQuery {
	query := &OutcomeMeasureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := og.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomegroup.Table, outcomegroup.FieldID, id),
			sqlgraph.To(outcomemeasure.Table, outcomemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomegroup.ParentTable, outcomegroup.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(og.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutcomeGroupClient) Hooks() []Hook {
	return c.hooks.OutcomeGroup
}

// OutcomeMeasureClient is a client for the OutcomeMeasure schema.
type OutcomeMeasureClient struct {
	config
}

// NewOutcomeMeasureClient returns a client for the OutcomeMeasure from the given config.
func NewOutcomeMeasureClient(c config) *OutcomeMeasureClient {
	return &OutcomeMeasureClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outcomemeasure.Hooks(f(g(h())))`.
func (c *OutcomeMeasureClient) Use(hooks ...Hook) {
	c.hooks.OutcomeMeasure = append(c.hooks.OutcomeMeasure, hooks...)
}

// Create returns a create builder for OutcomeMeasure.
func (c *OutcomeMeasureClient) Create() *OutcomeMeasureCreate {
	mutation := newOutcomeMeasureMutation(c.config, OpCreate)
	return &OutcomeMeasureCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutcomeMeasure entities.
func (c *OutcomeMeasureClient) CreateBulk(builders ...*OutcomeMeasureCreate) *OutcomeMeasureCreateBulk {
	return &OutcomeMeasureCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutcomeMeasure.
func (c *OutcomeMeasureClient) Update() *OutcomeMeasureUpdate {
	mutation := newOutcomeMeasureMutation(c.config, OpUpdate)
	return &OutcomeMeasureUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutcomeMeasureClient) UpdateOne(om *OutcomeMeasure) *OutcomeMeasureUpdateOne {
	mutation := newOutcomeMeasureMutation(c.config, OpUpdateOne, withOutcomeMeasure(om))
	return &OutcomeMeasureUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutcomeMeasureClient) UpdateOneID(id int) *OutcomeMeasureUpdateOne {
	mutation := newOutcomeMeasureMutation(c.config, OpUpdateOne, withOutcomeMeasureID(id))
	return &OutcomeMeasureUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutcomeMeasure.
func (c *OutcomeMeasureClient) Delete() *OutcomeMeasureDelete {
	mutation := newOutcomeMeasureMutation(c.config, OpDelete)
	return &OutcomeMeasureDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutcomeMeasureClient) DeleteOne(om *OutcomeMeasure) *OutcomeMeasureDeleteOne {
	return c.DeleteOneID(om.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutcomeMeasureClient) DeleteOneID(id int) *OutcomeMeasureDeleteOne {
	builder := c.Delete().Where(outcomemeasure.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutcomeMeasureDeleteOne{builder}
}

// Query returns a query builder for OutcomeMeasure.
func (c *OutcomeMeasureClient) Query() *OutcomeMeasureQuery {
	return &OutcomeMeasureQuery{
		config: c.config,
	}
}

// Get returns a OutcomeMeasure entity by its id.
func (c *OutcomeMeasureClient) Get(ctx context.Context, id int) (*OutcomeMeasure, error) {
	return c.Query().Where(outcomemeasure.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutcomeMeasureClient) GetX(ctx context.Context, id int) *OutcomeMeasure {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OutcomeMeasure.
func (c *OutcomeMeasureClient) QueryParent(om *OutcomeMeasure) *OutcomeMeasuresModuleQuery {
	query := &OutcomeMeasuresModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := om.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasure.Table, outcomemeasure.FieldID, id),
			sqlgraph.To(outcomemeasuresmodule.Table, outcomemeasuresmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomemeasure.ParentTable, outcomemeasure.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(om.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOutcomeGroupList queries the outcome_group_list edge of a OutcomeMeasure.
func (c *OutcomeMeasureClient) QueryOutcomeGroupList(om *OutcomeMeasure) *OutcomeGroupQuery {
	query := &OutcomeGroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := om.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasure.Table, outcomemeasure.FieldID, id),
			sqlgraph.To(outcomegroup.Table, outcomegroup.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomemeasure.OutcomeGroupListTable, outcomemeasure.OutcomeGroupListColumn),
		)
		fromV = sqlgraph.Neighbors(om.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOutcomeOverviewList queries the outcome_overview_list edge of a OutcomeMeasure.
func (c *OutcomeMeasureClient) QueryOutcomeOverviewList(om *OutcomeMeasure) *OutcomeOverviewQuery {
	query := &OutcomeOverviewQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := om.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasure.Table, outcomemeasure.FieldID, id),
			sqlgraph.To(outcomeoverview.Table, outcomeoverview.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomemeasure.OutcomeOverviewListTable, outcomemeasure.OutcomeOverviewListColumn),
		)
		fromV = sqlgraph.Neighbors(om.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOutcomeDenomList queries the outcome_denom_list edge of a OutcomeMeasure.
func (c *OutcomeMeasureClient) QueryOutcomeDenomList(om *OutcomeMeasure) *OutcomeDenomQuery {
	query := &OutcomeDenomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := om.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasure.Table, outcomemeasure.FieldID, id),
			sqlgraph.To(outcomedenom.Table, outcomedenom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomemeasure.OutcomeDenomListTable, outcomemeasure.OutcomeDenomListColumn),
		)
		fromV = sqlgraph.Neighbors(om.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOutcomeClassList queries the outcome_class_list edge of a OutcomeMeasure.
func (c *OutcomeMeasureClient) QueryOutcomeClassList(om *OutcomeMeasure) *OutcomeClassQuery {
	query := &OutcomeClassQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := om.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasure.Table, outcomemeasure.FieldID, id),
			sqlgraph.To(outcomeclass.Table, outcomeclass.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomemeasure.OutcomeClassListTable, outcomemeasure.OutcomeClassListColumn),
		)
		fromV = sqlgraph.Neighbors(om.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOutcomeAnalysisList queries the outcome_analysis_list edge of a OutcomeMeasure.
func (c *OutcomeMeasureClient) QueryOutcomeAnalysisList(om *OutcomeMeasure) *OutcomeAnalysisQuery {
	query := &OutcomeAnalysisQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := om.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasure.Table, outcomemeasure.FieldID, id),
			sqlgraph.To(outcomeanalysis.Table, outcomeanalysis.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomemeasure.OutcomeAnalysisListTable, outcomemeasure.OutcomeAnalysisListColumn),
		)
		fromV = sqlgraph.Neighbors(om.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutcomeMeasureClient) Hooks() []Hook {
	return c.hooks.OutcomeMeasure
}

// OutcomeMeasurementClient is a client for the OutcomeMeasurement schema.
type OutcomeMeasurementClient struct {
	config
}

// NewOutcomeMeasurementClient returns a client for the OutcomeMeasurement from the given config.
func NewOutcomeMeasurementClient(c config) *OutcomeMeasurementClient {
	return &OutcomeMeasurementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outcomemeasurement.Hooks(f(g(h())))`.
func (c *OutcomeMeasurementClient) Use(hooks ...Hook) {
	c.hooks.OutcomeMeasurement = append(c.hooks.OutcomeMeasurement, hooks...)
}

// Create returns a create builder for OutcomeMeasurement.
func (c *OutcomeMeasurementClient) Create() *OutcomeMeasurementCreate {
	mutation := newOutcomeMeasurementMutation(c.config, OpCreate)
	return &OutcomeMeasurementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutcomeMeasurement entities.
func (c *OutcomeMeasurementClient) CreateBulk(builders ...*OutcomeMeasurementCreate) *OutcomeMeasurementCreateBulk {
	return &OutcomeMeasurementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutcomeMeasurement.
func (c *OutcomeMeasurementClient) Update() *OutcomeMeasurementUpdate {
	mutation := newOutcomeMeasurementMutation(c.config, OpUpdate)
	return &OutcomeMeasurementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutcomeMeasurementClient) UpdateOne(om *OutcomeMeasurement) *OutcomeMeasurementUpdateOne {
	mutation := newOutcomeMeasurementMutation(c.config, OpUpdateOne, withOutcomeMeasurement(om))
	return &OutcomeMeasurementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutcomeMeasurementClient) UpdateOneID(id int) *OutcomeMeasurementUpdateOne {
	mutation := newOutcomeMeasurementMutation(c.config, OpUpdateOne, withOutcomeMeasurementID(id))
	return &OutcomeMeasurementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutcomeMeasurement.
func (c *OutcomeMeasurementClient) Delete() *OutcomeMeasurementDelete {
	mutation := newOutcomeMeasurementMutation(c.config, OpDelete)
	return &OutcomeMeasurementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutcomeMeasurementClient) DeleteOne(om *OutcomeMeasurement) *OutcomeMeasurementDeleteOne {
	return c.DeleteOneID(om.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutcomeMeasurementClient) DeleteOneID(id int) *OutcomeMeasurementDeleteOne {
	builder := c.Delete().Where(outcomemeasurement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutcomeMeasurementDeleteOne{builder}
}

// Query returns a query builder for OutcomeMeasurement.
func (c *OutcomeMeasurementClient) Query() *OutcomeMeasurementQuery {
	return &OutcomeMeasurementQuery{
		config: c.config,
	}
}

// Get returns a OutcomeMeasurement entity by its id.
func (c *OutcomeMeasurementClient) Get(ctx context.Context, id int) (*OutcomeMeasurement, error) {
	return c.Query().Where(outcomemeasurement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutcomeMeasurementClient) GetX(ctx context.Context, id int) *OutcomeMeasurement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OutcomeMeasurement.
func (c *OutcomeMeasurementClient) QueryParent(om *OutcomeMeasurement) *OutcomeCategoryQuery {
	query := &OutcomeCategoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := om.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasurement.Table, outcomemeasurement.FieldID, id),
			sqlgraph.To(outcomecategory.Table, outcomecategory.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomemeasurement.ParentTable, outcomemeasurement.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(om.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutcomeMeasurementClient) Hooks() []Hook {
	return c.hooks.OutcomeMeasurement
}

// OutcomeMeasuresModuleClient is a client for the OutcomeMeasuresModule schema.
type OutcomeMeasuresModuleClient struct {
	config
}

// NewOutcomeMeasuresModuleClient returns a client for the OutcomeMeasuresModule from the given config.
func NewOutcomeMeasuresModuleClient(c config) *OutcomeMeasuresModuleClient {
	return &OutcomeMeasuresModuleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outcomemeasuresmodule.Hooks(f(g(h())))`.
func (c *OutcomeMeasuresModuleClient) Use(hooks ...Hook) {
	c.hooks.OutcomeMeasuresModule = append(c.hooks.OutcomeMeasuresModule, hooks...)
}

// Create returns a create builder for OutcomeMeasuresModule.
func (c *OutcomeMeasuresModuleClient) Create() *OutcomeMeasuresModuleCreate {
	mutation := newOutcomeMeasuresModuleMutation(c.config, OpCreate)
	return &OutcomeMeasuresModuleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutcomeMeasuresModule entities.
func (c *OutcomeMeasuresModuleClient) CreateBulk(builders ...*OutcomeMeasuresModuleCreate) *OutcomeMeasuresModuleCreateBulk {
	return &OutcomeMeasuresModuleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutcomeMeasuresModule.
func (c *OutcomeMeasuresModuleClient) Update() *OutcomeMeasuresModuleUpdate {
	mutation := newOutcomeMeasuresModuleMutation(c.config, OpUpdate)
	return &OutcomeMeasuresModuleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutcomeMeasuresModuleClient) UpdateOne(omm *OutcomeMeasuresModule) *OutcomeMeasuresModuleUpdateOne {
	mutation := newOutcomeMeasuresModuleMutation(c.config, OpUpdateOne, withOutcomeMeasuresModule(omm))
	return &OutcomeMeasuresModuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutcomeMeasuresModuleClient) UpdateOneID(id int) *OutcomeMeasuresModuleUpdateOne {
	mutation := newOutcomeMeasuresModuleMutation(c.config, OpUpdateOne, withOutcomeMeasuresModuleID(id))
	return &OutcomeMeasuresModuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutcomeMeasuresModule.
func (c *OutcomeMeasuresModuleClient) Delete() *OutcomeMeasuresModuleDelete {
	mutation := newOutcomeMeasuresModuleMutation(c.config, OpDelete)
	return &OutcomeMeasuresModuleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutcomeMeasuresModuleClient) DeleteOne(omm *OutcomeMeasuresModule) *OutcomeMeasuresModuleDeleteOne {
	return c.DeleteOneID(omm.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutcomeMeasuresModuleClient) DeleteOneID(id int) *OutcomeMeasuresModuleDeleteOne {
	builder := c.Delete().Where(outcomemeasuresmodule.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutcomeMeasuresModuleDeleteOne{builder}
}

// Query returns a query builder for OutcomeMeasuresModule.
func (c *OutcomeMeasuresModuleClient) Query() *OutcomeMeasuresModuleQuery {
	return &OutcomeMeasuresModuleQuery{
		config: c.config,
	}
}

// Get returns a OutcomeMeasuresModule entity by its id.
func (c *OutcomeMeasuresModuleClient) Get(ctx context.Context, id int) (*OutcomeMeasuresModule, error) {
	return c.Query().Where(outcomemeasuresmodule.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutcomeMeasuresModuleClient) GetX(ctx context.Context, id int) *OutcomeMeasuresModule {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OutcomeMeasuresModule.
func (c *OutcomeMeasuresModuleClient) QueryParent(omm *OutcomeMeasuresModule) *ResultsDefinitionQuery {
	query := &ResultsDefinitionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := omm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasuresmodule.Table, outcomemeasuresmodule.FieldID, id),
			sqlgraph.To(resultsdefinition.Table, resultsdefinition.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, outcomemeasuresmodule.ParentTable, outcomemeasuresmodule.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(omm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOutcomeMeasureList queries the outcome_measure_list edge of a OutcomeMeasuresModule.
func (c *OutcomeMeasuresModuleClient) QueryOutcomeMeasureList(omm *OutcomeMeasuresModule) *OutcomeMeasureQuery {
	query := &OutcomeMeasureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := omm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasuresmodule.Table, outcomemeasuresmodule.FieldID, id),
			sqlgraph.To(outcomemeasure.Table, outcomemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomemeasuresmodule.OutcomeMeasureListTable, outcomemeasuresmodule.OutcomeMeasureListColumn),
		)
		fromV = sqlgraph.Neighbors(omm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutcomeMeasuresModuleClient) Hooks() []Hook {
	return c.hooks.OutcomeMeasuresModule
}

// OutcomeOverviewClient is a client for the OutcomeOverview schema.
type OutcomeOverviewClient struct {
	config
}

// NewOutcomeOverviewClient returns a client for the OutcomeOverview from the given config.
func NewOutcomeOverviewClient(c config) *OutcomeOverviewClient {
	return &OutcomeOverviewClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outcomeoverview.Hooks(f(g(h())))`.
func (c *OutcomeOverviewClient) Use(hooks ...Hook) {
	c.hooks.OutcomeOverview = append(c.hooks.OutcomeOverview, hooks...)
}

// Create returns a create builder for OutcomeOverview.
func (c *OutcomeOverviewClient) Create() *OutcomeOverviewCreate {
	mutation := newOutcomeOverviewMutation(c.config, OpCreate)
	return &OutcomeOverviewCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutcomeOverview entities.
func (c *OutcomeOverviewClient) CreateBulk(builders ...*OutcomeOverviewCreate) *OutcomeOverviewCreateBulk {
	return &OutcomeOverviewCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutcomeOverview.
func (c *OutcomeOverviewClient) Update() *OutcomeOverviewUpdate {
	mutation := newOutcomeOverviewMutation(c.config, OpUpdate)
	return &OutcomeOverviewUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutcomeOverviewClient) UpdateOne(oo *OutcomeOverview) *OutcomeOverviewUpdateOne {
	mutation := newOutcomeOverviewMutation(c.config, OpUpdateOne, withOutcomeOverview(oo))
	return &OutcomeOverviewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutcomeOverviewClient) UpdateOneID(id int) *OutcomeOverviewUpdateOne {
	mutation := newOutcomeOverviewMutation(c.config, OpUpdateOne, withOutcomeOverviewID(id))
	return &OutcomeOverviewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutcomeOverview.
func (c *OutcomeOverviewClient) Delete() *OutcomeOverviewDelete {
	mutation := newOutcomeOverviewMutation(c.config, OpDelete)
	return &OutcomeOverviewDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutcomeOverviewClient) DeleteOne(oo *OutcomeOverview) *OutcomeOverviewDeleteOne {
	return c.DeleteOneID(oo.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutcomeOverviewClient) DeleteOneID(id int) *OutcomeOverviewDeleteOne {
	builder := c.Delete().Where(outcomeoverview.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutcomeOverviewDeleteOne{builder}
}

// Query returns a query builder for OutcomeOverview.
func (c *OutcomeOverviewClient) Query() *OutcomeOverviewQuery {
	return &OutcomeOverviewQuery{
		config: c.config,
	}
}

// Get returns a OutcomeOverview entity by its id.
func (c *OutcomeOverviewClient) Get(ctx context.Context, id int) (*OutcomeOverview, error) {
	return c.Query().Where(outcomeoverview.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutcomeOverviewClient) GetX(ctx context.Context, id int) *OutcomeOverview {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a OutcomeOverview.
func (c *OutcomeOverviewClient) QueryParent(oo *OutcomeOverview) *OutcomeMeasureQuery {
	query := &OutcomeMeasureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := oo.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeoverview.Table, outcomeoverview.FieldID, id),
			sqlgraph.To(outcomemeasure.Table, outcomemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomeoverview.ParentTable, outcomeoverview.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(oo.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutcomeOverviewClient) Hooks() []Hook {
	return c.hooks.OutcomeOverview
}

// ParticipantFlowModuleClient is a client for the ParticipantFlowModule schema.
type ParticipantFlowModuleClient struct {
	config
}

// NewParticipantFlowModuleClient returns a client for the ParticipantFlowModule from the given config.
func NewParticipantFlowModuleClient(c config) *ParticipantFlowModuleClient {
	return &ParticipantFlowModuleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `participantflowmodule.Hooks(f(g(h())))`.
func (c *ParticipantFlowModuleClient) Use(hooks ...Hook) {
	c.hooks.ParticipantFlowModule = append(c.hooks.ParticipantFlowModule, hooks...)
}

// Create returns a create builder for ParticipantFlowModule.
func (c *ParticipantFlowModuleClient) Create() *ParticipantFlowModuleCreate {
	mutation := newParticipantFlowModuleMutation(c.config, OpCreate)
	return &ParticipantFlowModuleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ParticipantFlowModule entities.
func (c *ParticipantFlowModuleClient) CreateBulk(builders ...*ParticipantFlowModuleCreate) *ParticipantFlowModuleCreateBulk {
	return &ParticipantFlowModuleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ParticipantFlowModule.
func (c *ParticipantFlowModuleClient) Update() *ParticipantFlowModuleUpdate {
	mutation := newParticipantFlowModuleMutation(c.config, OpUpdate)
	return &ParticipantFlowModuleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ParticipantFlowModuleClient) UpdateOne(pfm *ParticipantFlowModule) *ParticipantFlowModuleUpdateOne {
	mutation := newParticipantFlowModuleMutation(c.config, OpUpdateOne, withParticipantFlowModule(pfm))
	return &ParticipantFlowModuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ParticipantFlowModuleClient) UpdateOneID(id int) *ParticipantFlowModuleUpdateOne {
	mutation := newParticipantFlowModuleMutation(c.config, OpUpdateOne, withParticipantFlowModuleID(id))
	return &ParticipantFlowModuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ParticipantFlowModule.
func (c *ParticipantFlowModuleClient) Delete() *ParticipantFlowModuleDelete {
	mutation := newParticipantFlowModuleMutation(c.config, OpDelete)
	return &ParticipantFlowModuleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ParticipantFlowModuleClient) DeleteOne(pfm *ParticipantFlowModule) *ParticipantFlowModuleDeleteOne {
	return c.DeleteOneID(pfm.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ParticipantFlowModuleClient) DeleteOneID(id int) *ParticipantFlowModuleDeleteOne {
	builder := c.Delete().Where(participantflowmodule.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ParticipantFlowModuleDeleteOne{builder}
}

// Query returns a query builder for ParticipantFlowModule.
func (c *ParticipantFlowModuleClient) Query() *ParticipantFlowModuleQuery {
	return &ParticipantFlowModuleQuery{
		config: c.config,
	}
}

// Get returns a ParticipantFlowModule entity by its id.
func (c *ParticipantFlowModuleClient) Get(ctx context.Context, id int) (*ParticipantFlowModule, error) {
	return c.Query().Where(participantflowmodule.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ParticipantFlowModuleClient) GetX(ctx context.Context, id int) *ParticipantFlowModule {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a ParticipantFlowModule.
func (c *ParticipantFlowModuleClient) QueryParent(pfm *ParticipantFlowModule) *ResultsDefinitionQuery {
	query := &ResultsDefinitionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pfm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(participantflowmodule.Table, participantflowmodule.FieldID, id),
			sqlgraph.To(resultsdefinition.Table, resultsdefinition.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, participantflowmodule.ParentTable, participantflowmodule.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(pfm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFlowGroupList queries the flow_group_list edge of a ParticipantFlowModule.
func (c *ParticipantFlowModuleClient) QueryFlowGroupList(pfm *ParticipantFlowModule) *FlowGroupQuery {
	query := &FlowGroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pfm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(participantflowmodule.Table, participantflowmodule.FieldID, id),
			sqlgraph.To(flowgroup.Table, flowgroup.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, participantflowmodule.FlowGroupListTable, participantflowmodule.FlowGroupListColumn),
		)
		fromV = sqlgraph.Neighbors(pfm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFlowPeriodList queries the flow_period_list edge of a ParticipantFlowModule.
func (c *ParticipantFlowModuleClient) QueryFlowPeriodList(pfm *ParticipantFlowModule) *FlowPeriodQuery {
	query := &FlowPeriodQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pfm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(participantflowmodule.Table, participantflowmodule.FieldID, id),
			sqlgraph.To(flowperiod.Table, flowperiod.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, participantflowmodule.FlowPeriodListTable, participantflowmodule.FlowPeriodListColumn),
		)
		fromV = sqlgraph.Neighbors(pfm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ParticipantFlowModuleClient) Hooks() []Hook {
	return c.hooks.ParticipantFlowModule
}

// PointOfContactClient is a client for the PointOfContact schema.
type PointOfContactClient struct {
	config
}

// NewPointOfContactClient returns a client for the PointOfContact from the given config.
func NewPointOfContactClient(c config) *PointOfContactClient {
	return &PointOfContactClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `pointofcontact.Hooks(f(g(h())))`.
func (c *PointOfContactClient) Use(hooks ...Hook) {
	c.hooks.PointOfContact = append(c.hooks.PointOfContact, hooks...)
}

// Create returns a create builder for PointOfContact.
func (c *PointOfContactClient) Create() *PointOfContactCreate {
	mutation := newPointOfContactMutation(c.config, OpCreate)
	return &PointOfContactCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PointOfContact entities.
func (c *PointOfContactClient) CreateBulk(builders ...*PointOfContactCreate) *PointOfContactCreateBulk {
	return &PointOfContactCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PointOfContact.
func (c *PointOfContactClient) Update() *PointOfContactUpdate {
	mutation := newPointOfContactMutation(c.config, OpUpdate)
	return &PointOfContactUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PointOfContactClient) UpdateOne(poc *PointOfContact) *PointOfContactUpdateOne {
	mutation := newPointOfContactMutation(c.config, OpUpdateOne, withPointOfContact(poc))
	return &PointOfContactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PointOfContactClient) UpdateOneID(id int) *PointOfContactUpdateOne {
	mutation := newPointOfContactMutation(c.config, OpUpdateOne, withPointOfContactID(id))
	return &PointOfContactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PointOfContact.
func (c *PointOfContactClient) Delete() *PointOfContactDelete {
	mutation := newPointOfContactMutation(c.config, OpDelete)
	return &PointOfContactDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PointOfContactClient) DeleteOne(poc *PointOfContact) *PointOfContactDeleteOne {
	return c.DeleteOneID(poc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PointOfContactClient) DeleteOneID(id int) *PointOfContactDeleteOne {
	builder := c.Delete().Where(pointofcontact.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PointOfContactDeleteOne{builder}
}

// Query returns a query builder for PointOfContact.
func (c *PointOfContactClient) Query() *PointOfContactQuery {
	return &PointOfContactQuery{
		config: c.config,
	}
}

// Get returns a PointOfContact entity by its id.
func (c *PointOfContactClient) Get(ctx context.Context, id int) (*PointOfContact, error) {
	return c.Query().Where(pointofcontact.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PointOfContactClient) GetX(ctx context.Context, id int) *PointOfContact {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a PointOfContact.
func (c *PointOfContactClient) QueryParent(poc *PointOfContact) *MoreInfoModuleQuery {
	query := &MoreInfoModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := poc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(pointofcontact.Table, pointofcontact.FieldID, id),
			sqlgraph.To(moreinfomodule.Table, moreinfomodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, pointofcontact.ParentTable, pointofcontact.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(poc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PointOfContactClient) Hooks() []Hook {
	return c.hooks.PointOfContact
}

// ResultsDefinitionClient is a client for the ResultsDefinition schema.
type ResultsDefinitionClient struct {
	config
}

// NewResultsDefinitionClient returns a client for the ResultsDefinition from the given config.
func NewResultsDefinitionClient(c config) *ResultsDefinitionClient {
	return &ResultsDefinitionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `resultsdefinition.Hooks(f(g(h())))`.
func (c *ResultsDefinitionClient) Use(hooks ...Hook) {
	c.hooks.ResultsDefinition = append(c.hooks.ResultsDefinition, hooks...)
}

// Create returns a create builder for ResultsDefinition.
func (c *ResultsDefinitionClient) Create() *ResultsDefinitionCreate {
	mutation := newResultsDefinitionMutation(c.config, OpCreate)
	return &ResultsDefinitionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ResultsDefinition entities.
func (c *ResultsDefinitionClient) CreateBulk(builders ...*ResultsDefinitionCreate) *ResultsDefinitionCreateBulk {
	return &ResultsDefinitionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ResultsDefinition.
func (c *ResultsDefinitionClient) Update() *ResultsDefinitionUpdate {
	mutation := newResultsDefinitionMutation(c.config, OpUpdate)
	return &ResultsDefinitionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ResultsDefinitionClient) UpdateOne(rd *ResultsDefinition) *ResultsDefinitionUpdateOne {
	mutation := newResultsDefinitionMutation(c.config, OpUpdateOne, withResultsDefinition(rd))
	return &ResultsDefinitionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ResultsDefinitionClient) UpdateOneID(id int) *ResultsDefinitionUpdateOne {
	mutation := newResultsDefinitionMutation(c.config, OpUpdateOne, withResultsDefinitionID(id))
	return &ResultsDefinitionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ResultsDefinition.
func (c *ResultsDefinitionClient) Delete() *ResultsDefinitionDelete {
	mutation := newResultsDefinitionMutation(c.config, OpDelete)
	return &ResultsDefinitionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ResultsDefinitionClient) DeleteOne(rd *ResultsDefinition) *ResultsDefinitionDeleteOne {
	return c.DeleteOneID(rd.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ResultsDefinitionClient) DeleteOneID(id int) *ResultsDefinitionDeleteOne {
	builder := c.Delete().Where(resultsdefinition.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ResultsDefinitionDeleteOne{builder}
}

// Query returns a query builder for ResultsDefinition.
func (c *ResultsDefinitionClient) Query() *ResultsDefinitionQuery {
	return &ResultsDefinitionQuery{
		config: c.config,
	}
}

// Get returns a ResultsDefinition entity by its id.
func (c *ResultsDefinitionClient) Get(ctx context.Context, id int) (*ResultsDefinition, error) {
	return c.Query().Where(resultsdefinition.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ResultsDefinitionClient) GetX(ctx context.Context, id int) *ResultsDefinition {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a ResultsDefinition.
func (c *ResultsDefinitionClient) QueryParent(rd *ResultsDefinition) *ClinicalTrialQuery {
	query := &ClinicalTrialQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := rd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(resultsdefinition.Table, resultsdefinition.FieldID, id),
			sqlgraph.To(clinicaltrial.Table, clinicaltrial.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, resultsdefinition.ParentTable, resultsdefinition.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(rd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParticipantFlowModule queries the participant_flow_module edge of a ResultsDefinition.
func (c *ResultsDefinitionClient) QueryParticipantFlowModule(rd *ResultsDefinition) *ParticipantFlowModuleQuery {
	query := &ParticipantFlowModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := rd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(resultsdefinition.Table, resultsdefinition.FieldID, id),
			sqlgraph.To(participantflowmodule.Table, participantflowmodule.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, resultsdefinition.ParticipantFlowModuleTable, resultsdefinition.ParticipantFlowModuleColumn),
		)
		fromV = sqlgraph.Neighbors(rd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBaselineCharacteristicsModule queries the baseline_characteristics_module edge of a ResultsDefinition.
func (c *ResultsDefinitionClient) QueryBaselineCharacteristicsModule(rd *ResultsDefinition) *BaselineCharacteristicsModuleQuery {
	query := &BaselineCharacteristicsModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := rd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(resultsdefinition.Table, resultsdefinition.FieldID, id),
			sqlgraph.To(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, resultsdefinition.BaselineCharacteristicsModuleTable, resultsdefinition.BaselineCharacteristicsModuleColumn),
		)
		fromV = sqlgraph.Neighbors(rd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOutcomeMeasuresModule queries the outcome_measures_module edge of a ResultsDefinition.
func (c *ResultsDefinitionClient) QueryOutcomeMeasuresModule(rd *ResultsDefinition) *OutcomeMeasuresModuleQuery {
	query := &OutcomeMeasuresModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := rd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(resultsdefinition.Table, resultsdefinition.FieldID, id),
			sqlgraph.To(outcomemeasuresmodule.Table, outcomemeasuresmodule.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, resultsdefinition.OutcomeMeasuresModuleTable, resultsdefinition.OutcomeMeasuresModuleColumn),
		)
		fromV = sqlgraph.Neighbors(rd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAdverseEventsModule queries the adverse_events_module edge of a ResultsDefinition.
func (c *ResultsDefinitionClient) QueryAdverseEventsModule(rd *ResultsDefinition) *AdverseEventsModuleQuery {
	query := &AdverseEventsModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := rd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(resultsdefinition.Table, resultsdefinition.FieldID, id),
			sqlgraph.To(adverseeventsmodule.Table, adverseeventsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, resultsdefinition.AdverseEventsModuleTable, resultsdefinition.AdverseEventsModuleColumn),
		)
		fromV = sqlgraph.Neighbors(rd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMoreInfoModule queries the more_info_module edge of a ResultsDefinition.
func (c *ResultsDefinitionClient) QueryMoreInfoModule(rd *ResultsDefinition) *MoreInfoModuleQuery {
	query := &MoreInfoModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := rd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(resultsdefinition.Table, resultsdefinition.FieldID, id),
			sqlgraph.To(moreinfomodule.Table, moreinfomodule.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, resultsdefinition.MoreInfoModuleTable, resultsdefinition.MoreInfoModuleColumn),
		)
		fromV = sqlgraph.Neighbors(rd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ResultsDefinitionClient) Hooks() []Hook {
	return c.hooks.ResultsDefinition
}

// ScheduleClient is a client for the Schedule schema.
type ScheduleClient struct {
	config
}

// NewScheduleClient returns a client for the Schedule from the given config.
func NewScheduleClient(c config) *ScheduleClient {
	return &ScheduleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `schedule.Hooks(f(g(h())))`.
func (c *ScheduleClient) Use(hooks ...Hook) {
	c.hooks.Schedule = append(c.hooks.Schedule, hooks...)
}

// Create returns a create builder for Schedule.
func (c *ScheduleClient) Create() *ScheduleCreate {
	mutation := newScheduleMutation(c.config, OpCreate)
	return &ScheduleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Schedule entities.
func (c *ScheduleClient) CreateBulk(builders ...*ScheduleCreate) *ScheduleCreateBulk {
	return &ScheduleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Schedule.
func (c *ScheduleClient) Update() *ScheduleUpdate {
	mutation := newScheduleMutation(c.config, OpUpdate)
	return &ScheduleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ScheduleClient) UpdateOne(s *Schedule) *ScheduleUpdateOne {
	mutation := newScheduleMutation(c.config, OpUpdateOne, withSchedule(s))
	return &ScheduleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ScheduleClient) UpdateOneID(id int) *ScheduleUpdateOne {
	mutation := newScheduleMutation(c.config, OpUpdateOne, withScheduleID(id))
	return &ScheduleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Schedule.
func (c *ScheduleClient) Delete() *ScheduleDelete {
	mutation := newScheduleMutation(c.config, OpDelete)
	return &ScheduleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ScheduleClient) DeleteOne(s *Schedule) *ScheduleDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ScheduleClient) DeleteOneID(id int) *ScheduleDeleteOne {
	builder := c.Delete().Where(schedule.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ScheduleDeleteOne{builder}
}

// Query returns a query builder for Schedule.
func (c *ScheduleClient) Query() *ScheduleQuery {
	return &ScheduleQuery{
		config: c.config,
	}
}

// Get returns a Schedule entity by its id.
func (c *ScheduleClient) Get(ctx context.Context, id int) (*Schedule, error) {
	return c.Query().Where(schedule.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ScheduleClient) GetX(ctx context.Context, id int) *Schedule {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ScheduleClient) Hooks() []Hook {
	return c.hooks.Schedule
}

// SeriousEventClient is a client for the SeriousEvent schema.
type SeriousEventClient struct {
	config
}

// NewSeriousEventClient returns a client for the SeriousEvent from the given config.
func NewSeriousEventClient(c config) *SeriousEventClient {
	return &SeriousEventClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `seriousevent.Hooks(f(g(h())))`.
func (c *SeriousEventClient) Use(hooks ...Hook) {
	c.hooks.SeriousEvent = append(c.hooks.SeriousEvent, hooks...)
}

// Create returns a create builder for SeriousEvent.
func (c *SeriousEventClient) Create() *SeriousEventCreate {
	mutation := newSeriousEventMutation(c.config, OpCreate)
	return &SeriousEventCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SeriousEvent entities.
func (c *SeriousEventClient) CreateBulk(builders ...*SeriousEventCreate) *SeriousEventCreateBulk {
	return &SeriousEventCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SeriousEvent.
func (c *SeriousEventClient) Update() *SeriousEventUpdate {
	mutation := newSeriousEventMutation(c.config, OpUpdate)
	return &SeriousEventUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SeriousEventClient) UpdateOne(se *SeriousEvent) *SeriousEventUpdateOne {
	mutation := newSeriousEventMutation(c.config, OpUpdateOne, withSeriousEvent(se))
	return &SeriousEventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SeriousEventClient) UpdateOneID(id int) *SeriousEventUpdateOne {
	mutation := newSeriousEventMutation(c.config, OpUpdateOne, withSeriousEventID(id))
	return &SeriousEventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SeriousEvent.
func (c *SeriousEventClient) Delete() *SeriousEventDelete {
	mutation := newSeriousEventMutation(c.config, OpDelete)
	return &SeriousEventDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SeriousEventClient) DeleteOne(se *SeriousEvent) *SeriousEventDeleteOne {
	return c.DeleteOneID(se.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SeriousEventClient) DeleteOneID(id int) *SeriousEventDeleteOne {
	builder := c.Delete().Where(seriousevent.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SeriousEventDeleteOne{builder}
}

// Query returns a query builder for SeriousEvent.
func (c *SeriousEventClient) Query() *SeriousEventQuery {
	return &SeriousEventQuery{
		config: c.config,
	}
}

// Get returns a SeriousEvent entity by its id.
func (c *SeriousEventClient) Get(ctx context.Context, id int) (*SeriousEvent, error) {
	return c.Query().Where(seriousevent.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SeriousEventClient) GetX(ctx context.Context, id int) *SeriousEvent {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a SeriousEvent.
func (c *SeriousEventClient) QueryParent(se *SeriousEvent) *AdverseEventsModuleQuery {
	query := &AdverseEventsModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := se.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(seriousevent.Table, seriousevent.FieldID, id),
			sqlgraph.To(adverseeventsmodule.Table, adverseeventsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, seriousevent.ParentTable, seriousevent.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(se.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySeriousEventStatsList queries the serious_event_stats_list edge of a SeriousEvent.
func (c *SeriousEventClient) QuerySeriousEventStatsList(se *SeriousEvent) *SeriousEventStatsQuery {
	query := &SeriousEventStatsQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := se.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(seriousevent.Table, seriousevent.FieldID, id),
			sqlgraph.To(seriouseventstats.Table, seriouseventstats.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, seriousevent.SeriousEventStatsListTable, seriousevent.SeriousEventStatsListColumn),
		)
		fromV = sqlgraph.Neighbors(se.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SeriousEventClient) Hooks() []Hook {
	return c.hooks.SeriousEvent
}

// SeriousEventStatsClient is a client for the SeriousEventStats schema.
type SeriousEventStatsClient struct {
	config
}

// NewSeriousEventStatsClient returns a client for the SeriousEventStats from the given config.
func NewSeriousEventStatsClient(c config) *SeriousEventStatsClient {
	return &SeriousEventStatsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `seriouseventstats.Hooks(f(g(h())))`.
func (c *SeriousEventStatsClient) Use(hooks ...Hook) {
	c.hooks.SeriousEventStats = append(c.hooks.SeriousEventStats, hooks...)
}

// Create returns a create builder for SeriousEventStats.
func (c *SeriousEventStatsClient) Create() *SeriousEventStatsCreate {
	mutation := newSeriousEventStatsMutation(c.config, OpCreate)
	return &SeriousEventStatsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SeriousEventStats entities.
func (c *SeriousEventStatsClient) CreateBulk(builders ...*SeriousEventStatsCreate) *SeriousEventStatsCreateBulk {
	return &SeriousEventStatsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SeriousEventStats.
func (c *SeriousEventStatsClient) Update() *SeriousEventStatsUpdate {
	mutation := newSeriousEventStatsMutation(c.config, OpUpdate)
	return &SeriousEventStatsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SeriousEventStatsClient) UpdateOne(ses *SeriousEventStats) *SeriousEventStatsUpdateOne {
	mutation := newSeriousEventStatsMutation(c.config, OpUpdateOne, withSeriousEventStats(ses))
	return &SeriousEventStatsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SeriousEventStatsClient) UpdateOneID(id int) *SeriousEventStatsUpdateOne {
	mutation := newSeriousEventStatsMutation(c.config, OpUpdateOne, withSeriousEventStatsID(id))
	return &SeriousEventStatsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SeriousEventStats.
func (c *SeriousEventStatsClient) Delete() *SeriousEventStatsDelete {
	mutation := newSeriousEventStatsMutation(c.config, OpDelete)
	return &SeriousEventStatsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SeriousEventStatsClient) DeleteOne(ses *SeriousEventStats) *SeriousEventStatsDeleteOne {
	return c.DeleteOneID(ses.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SeriousEventStatsClient) DeleteOneID(id int) *SeriousEventStatsDeleteOne {
	builder := c.Delete().Where(seriouseventstats.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SeriousEventStatsDeleteOne{builder}
}

// Query returns a query builder for SeriousEventStats.
func (c *SeriousEventStatsClient) Query() *SeriousEventStatsQuery {
	return &SeriousEventStatsQuery{
		config: c.config,
	}
}

// Get returns a SeriousEventStats entity by its id.
func (c *SeriousEventStatsClient) Get(ctx context.Context, id int) (*SeriousEventStats, error) {
	return c.Query().Where(seriouseventstats.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SeriousEventStatsClient) GetX(ctx context.Context, id int) *SeriousEventStats {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a SeriousEventStats.
func (c *SeriousEventStatsClient) QueryParent(ses *SeriousEventStats) *SeriousEventQuery {
	query := &SeriousEventQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ses.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(seriouseventstats.Table, seriouseventstats.FieldID, id),
			sqlgraph.To(seriousevent.Table, seriousevent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, seriouseventstats.ParentTable, seriouseventstats.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(ses.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SeriousEventStatsClient) Hooks() []Hook {
	return c.hooks.SeriousEventStats
}

// StudyEligibilityClient is a client for the StudyEligibility schema.
type StudyEligibilityClient struct {
	config
}

// NewStudyEligibilityClient returns a client for the StudyEligibility from the given config.
func NewStudyEligibilityClient(c config) *StudyEligibilityClient {
	return &StudyEligibilityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `studyeligibility.Hooks(f(g(h())))`.
func (c *StudyEligibilityClient) Use(hooks ...Hook) {
	c.hooks.StudyEligibility = append(c.hooks.StudyEligibility, hooks...)
}

// Create returns a create builder for StudyEligibility.
func (c *StudyEligibilityClient) Create() *StudyEligibilityCreate {
	mutation := newStudyEligibilityMutation(c.config, OpCreate)
	return &StudyEligibilityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of StudyEligibility entities.
func (c *StudyEligibilityClient) CreateBulk(builders ...*StudyEligibilityCreate) *StudyEligibilityCreateBulk {
	return &StudyEligibilityCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for StudyEligibility.
func (c *StudyEligibilityClient) Update() *StudyEligibilityUpdate {
	mutation := newStudyEligibilityMutation(c.config, OpUpdate)
	return &StudyEligibilityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StudyEligibilityClient) UpdateOne(se *StudyEligibility) *StudyEligibilityUpdateOne {
	mutation := newStudyEligibilityMutation(c.config, OpUpdateOne, withStudyEligibility(se))
	return &StudyEligibilityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StudyEligibilityClient) UpdateOneID(id int) *StudyEligibilityUpdateOne {
	mutation := newStudyEligibilityMutation(c.config, OpUpdateOne, withStudyEligibilityID(id))
	return &StudyEligibilityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for StudyEligibility.
func (c *StudyEligibilityClient) Delete() *StudyEligibilityDelete {
	mutation := newStudyEligibilityMutation(c.config, OpDelete)
	return &StudyEligibilityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *StudyEligibilityClient) DeleteOne(se *StudyEligibility) *StudyEligibilityDeleteOne {
	return c.DeleteOneID(se.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *StudyEligibilityClient) DeleteOneID(id int) *StudyEligibilityDeleteOne {
	builder := c.Delete().Where(studyeligibility.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StudyEligibilityDeleteOne{builder}
}

// Query returns a query builder for StudyEligibility.
func (c *StudyEligibilityClient) Query() *StudyEligibilityQuery {
	return &StudyEligibilityQuery{
		config: c.config,
	}
}

// Get returns a StudyEligibility entity by its id.
func (c *StudyEligibilityClient) Get(ctx context.Context, id int) (*StudyEligibility, error) {
	return c.Query().Where(studyeligibility.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StudyEligibilityClient) GetX(ctx context.Context, id int) *StudyEligibility {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a StudyEligibility.
func (c *StudyEligibilityClient) QueryParent(se *StudyEligibility) *ClinicalTrialQuery {
	query := &ClinicalTrialQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := se.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(studyeligibility.Table, studyeligibility.FieldID, id),
			sqlgraph.To(clinicaltrial.Table, clinicaltrial.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, studyeligibility.ParentTable, studyeligibility.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(se.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StudyEligibilityClient) Hooks() []Hook {
	return c.hooks.StudyEligibility
}

// StudyLocationClient is a client for the StudyLocation schema.
type StudyLocationClient struct {
	config
}

// NewStudyLocationClient returns a client for the StudyLocation from the given config.
func NewStudyLocationClient(c config) *StudyLocationClient {
	return &StudyLocationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `studylocation.Hooks(f(g(h())))`.
func (c *StudyLocationClient) Use(hooks ...Hook) {
	c.hooks.StudyLocation = append(c.hooks.StudyLocation, hooks...)
}

// Create returns a create builder for StudyLocation.
func (c *StudyLocationClient) Create() *StudyLocationCreate {
	mutation := newStudyLocationMutation(c.config, OpCreate)
	return &StudyLocationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of StudyLocation entities.
func (c *StudyLocationClient) CreateBulk(builders ...*StudyLocationCreate) *StudyLocationCreateBulk {
	return &StudyLocationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for StudyLocation.
func (c *StudyLocationClient) Update() *StudyLocationUpdate {
	mutation := newStudyLocationMutation(c.config, OpUpdate)
	return &StudyLocationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StudyLocationClient) UpdateOne(sl *StudyLocation) *StudyLocationUpdateOne {
	mutation := newStudyLocationMutation(c.config, OpUpdateOne, withStudyLocation(sl))
	return &StudyLocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StudyLocationClient) UpdateOneID(id int) *StudyLocationUpdateOne {
	mutation := newStudyLocationMutation(c.config, OpUpdateOne, withStudyLocationID(id))
	return &StudyLocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for StudyLocation.
func (c *StudyLocationClient) Delete() *StudyLocationDelete {
	mutation := newStudyLocationMutation(c.config, OpDelete)
	return &StudyLocationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *StudyLocationClient) DeleteOne(sl *StudyLocation) *StudyLocationDeleteOne {
	return c.DeleteOneID(sl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *StudyLocationClient) DeleteOneID(id int) *StudyLocationDeleteOne {
	builder := c.Delete().Where(studylocation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StudyLocationDeleteOne{builder}
}

// Query returns a query builder for StudyLocation.
func (c *StudyLocationClient) Query() *StudyLocationQuery {
	return &StudyLocationQuery{
		config: c.config,
	}
}

// Get returns a StudyLocation entity by its id.
func (c *StudyLocationClient) Get(ctx context.Context, id int) (*StudyLocation, error) {
	return c.Query().Where(studylocation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StudyLocationClient) GetX(ctx context.Context, id int) *StudyLocation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a StudyLocation.
func (c *StudyLocationClient) QueryParent(sl *StudyLocation) *ClinicalTrialQuery {
	query := &ClinicalTrialQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := sl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(studylocation.Table, studylocation.FieldID, id),
			sqlgraph.To(clinicaltrial.Table, clinicaltrial.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, studylocation.ParentTable, studylocation.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(sl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StudyLocationClient) Hooks() []Hook {
	return c.hooks.StudyLocation
}

// TaskClient is a client for the Task schema.
type TaskClient struct {
	config
}

// NewTaskClient returns a client for the Task from the given config.
func NewTaskClient(c config) *TaskClient {
	return &TaskClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `task.Hooks(f(g(h())))`.
func (c *TaskClient) Use(hooks ...Hook) {
	c.hooks.Task = append(c.hooks.Task, hooks...)
}

// Create returns a create builder for Task.
func (c *TaskClient) Create() *TaskCreate {
	mutation := newTaskMutation(c.config, OpCreate)
	return &TaskCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Task entities.
func (c *TaskClient) CreateBulk(builders ...*TaskCreate) *TaskCreateBulk {
	return &TaskCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Task.
func (c *TaskClient) Update() *TaskUpdate {
	mutation := newTaskMutation(c.config, OpUpdate)
	return &TaskUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TaskClient) UpdateOne(t *Task) *TaskUpdateOne {
	mutation := newTaskMutation(c.config, OpUpdateOne, withTask(t))
	return &TaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TaskClient) UpdateOneID(id string) *TaskUpdateOne {
	mutation := newTaskMutation(c.config, OpUpdateOne, withTaskID(id))
	return &TaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Task.
func (c *TaskClient) Delete() *TaskDelete {
	mutation := newTaskMutation(c.config, OpDelete)
	return &TaskDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TaskClient) DeleteOne(t *Task) *TaskDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TaskClient) DeleteOneID(id string) *TaskDeleteOne {
	builder := c.Delete().Where(task.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TaskDeleteOne{builder}
}

// Query returns a query builder for Task.
func (c *TaskClient) Query() *TaskQuery {
	return &TaskQuery{
		config: c.config,
	}
}

// Get returns a Task entity by its id.
func (c *TaskClient) Get(ctx context.Context, id string) (*Task, error) {
	return c.Query().Where(task.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TaskClient) GetX(ctx context.Context, id string) *Task {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TaskClient) Hooks() []Hook {
	return c.hooks.Task
}

// UseCaseClient is a client for the UseCase schema.
type UseCaseClient struct {
	config
}

// NewUseCaseClient returns a client for the UseCase from the given config.
func NewUseCaseClient(c config) *UseCaseClient {
	return &UseCaseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usecase.Hooks(f(g(h())))`.
func (c *UseCaseClient) Use(hooks ...Hook) {
	c.hooks.UseCase = append(c.hooks.UseCase, hooks...)
}

// Create returns a create builder for UseCase.
func (c *UseCaseClient) Create() *UseCaseCreate {
	mutation := newUseCaseMutation(c.config, OpCreate)
	return &UseCaseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UseCase entities.
func (c *UseCaseClient) CreateBulk(builders ...*UseCaseCreate) *UseCaseCreateBulk {
	return &UseCaseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UseCase.
func (c *UseCaseClient) Update() *UseCaseUpdate {
	mutation := newUseCaseMutation(c.config, OpUpdate)
	return &UseCaseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UseCaseClient) UpdateOne(uc *UseCase) *UseCaseUpdateOne {
	mutation := newUseCaseMutation(c.config, OpUpdateOne, withUseCase(uc))
	return &UseCaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UseCaseClient) UpdateOneID(id int) *UseCaseUpdateOne {
	mutation := newUseCaseMutation(c.config, OpUpdateOne, withUseCaseID(id))
	return &UseCaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UseCase.
func (c *UseCaseClient) Delete() *UseCaseDelete {
	mutation := newUseCaseMutation(c.config, OpDelete)
	return &UseCaseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UseCaseClient) DeleteOne(uc *UseCase) *UseCaseDeleteOne {
	return c.DeleteOneID(uc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UseCaseClient) DeleteOneID(id int) *UseCaseDeleteOne {
	builder := c.Delete().Where(usecase.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UseCaseDeleteOne{builder}
}

// Query returns a query builder for UseCase.
func (c *UseCaseClient) Query() *UseCaseQuery {
	return &UseCaseQuery{
		config: c.config,
	}
}

// Get returns a UseCase entity by its id.
func (c *UseCaseClient) Get(ctx context.Context, id int) (*UseCase, error) {
	return c.Query().Where(usecase.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UseCaseClient) GetX(ctx context.Context, id int) *UseCase {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UseCaseClient) Hooks() []Hook {
	return c.hooks.UseCase
}

// VaccineClient is a client for the Vaccine schema.
type VaccineClient struct {
	config
}

// NewVaccineClient returns a client for the Vaccine from the given config.
func NewVaccineClient(c config) *VaccineClient {
	return &VaccineClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `vaccine.Hooks(f(g(h())))`.
func (c *VaccineClient) Use(hooks ...Hook) {
	c.hooks.Vaccine = append(c.hooks.Vaccine, hooks...)
}

// Create returns a create builder for Vaccine.
func (c *VaccineClient) Create() *VaccineCreate {
	mutation := newVaccineMutation(c.config, OpCreate)
	return &VaccineCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Vaccine entities.
func (c *VaccineClient) CreateBulk(builders ...*VaccineCreate) *VaccineCreateBulk {
	return &VaccineCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Vaccine.
func (c *VaccineClient) Update() *VaccineUpdate {
	mutation := newVaccineMutation(c.config, OpUpdate)
	return &VaccineUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VaccineClient) UpdateOne(v *Vaccine) *VaccineUpdateOne {
	mutation := newVaccineMutation(c.config, OpUpdateOne, withVaccine(v))
	return &VaccineUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VaccineClient) UpdateOneID(id int) *VaccineUpdateOne {
	mutation := newVaccineMutation(c.config, OpUpdateOne, withVaccineID(id))
	return &VaccineUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Vaccine.
func (c *VaccineClient) Delete() *VaccineDelete {
	mutation := newVaccineMutation(c.config, OpDelete)
	return &VaccineDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VaccineClient) DeleteOne(v *Vaccine) *VaccineDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VaccineClient) DeleteOneID(id int) *VaccineDeleteOne {
	builder := c.Delete().Where(vaccine.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VaccineDeleteOne{builder}
}

// Query returns a query builder for Vaccine.
func (c *VaccineClient) Query() *VaccineQuery {
	return &VaccineQuery{
		config: c.config,
	}
}

// Get returns a Vaccine entity by its id.
func (c *VaccineClient) Get(ctx context.Context, id int) (*Vaccine, error) {
	return c.Query().Where(vaccine.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VaccineClient) GetX(ctx context.Context, id int) *Vaccine {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *VaccineClient) Hooks() []Hook {
	return c.hooks.Vaccine
}
