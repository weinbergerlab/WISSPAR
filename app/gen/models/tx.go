// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
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

	// lazily loaded.
	client     *Client
	clientOnce sync.Once

	// completion callbacks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook

	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Commit method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	tx.mu.Lock()
	hooks := append([]CommitHook(nil), tx.onCommit...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onCommit = append(tx.onCommit, f)
}

type (
	// Rollbacker is the interface that wraps the Rollback method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	tx.mu.Lock()
	hooks := append([]RollbackHook(nil), tx.onRollback...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onRollback = append(tx.onRollback, f)
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.AdverseEventsModule = NewAdverseEventsModuleClient(tx.config)
	tx.BaselineCategory = NewBaselineCategoryClient(tx.config)
	tx.BaselineCharacteristicsModule = NewBaselineCharacteristicsModuleClient(tx.config)
	tx.BaselineClass = NewBaselineClassClient(tx.config)
	tx.BaselineClassDenom = NewBaselineClassDenomClient(tx.config)
	tx.BaselineClassDenomCount = NewBaselineClassDenomCountClient(tx.config)
	tx.BaselineDenom = NewBaselineDenomClient(tx.config)
	tx.BaselineDenomCount = NewBaselineDenomCountClient(tx.config)
	tx.BaselineGroup = NewBaselineGroupClient(tx.config)
	tx.BaselineMeasure = NewBaselineMeasureClient(tx.config)
	tx.BaselineMeasureDenom = NewBaselineMeasureDenomClient(tx.config)
	tx.BaselineMeasureDenomCount = NewBaselineMeasureDenomCountClient(tx.config)
	tx.BaselineMeasurement = NewBaselineMeasurementClient(tx.config)
	tx.CertainAgreement = NewCertainAgreementClient(tx.config)
	tx.ClinicalTrial = NewClinicalTrialClient(tx.config)
	tx.ClinicalTrialMetadata = NewClinicalTrialMetadataClient(tx.config)
	tx.DoseDescription = NewDoseDescriptionClient(tx.config)
	tx.EventGroup = NewEventGroupClient(tx.config)
	tx.FlowAchievement = NewFlowAchievementClient(tx.config)
	tx.FlowDropWithdraw = NewFlowDropWithdrawClient(tx.config)
	tx.FlowGroup = NewFlowGroupClient(tx.config)
	tx.FlowMilestone = NewFlowMilestoneClient(tx.config)
	tx.FlowPeriod = NewFlowPeriodClient(tx.config)
	tx.FlowReason = NewFlowReasonClient(tx.config)
	tx.ImmunocompromisedGroups = NewImmunocompromisedGroupsClient(tx.config)
	tx.Manufacturer = NewManufacturerClient(tx.config)
	tx.MoreInfoModule = NewMoreInfoModuleClient(tx.config)
	tx.OtherEvent = NewOtherEventClient(tx.config)
	tx.OtherEventStats = NewOtherEventStatsClient(tx.config)
	tx.OutcomeAnalysis = NewOutcomeAnalysisClient(tx.config)
	tx.OutcomeAnalysisGroupID = NewOutcomeAnalysisGroupIDClient(tx.config)
	tx.OutcomeCategory = NewOutcomeCategoryClient(tx.config)
	tx.OutcomeClass = NewOutcomeClassClient(tx.config)
	tx.OutcomeClassDenom = NewOutcomeClassDenomClient(tx.config)
	tx.OutcomeClassDenomCount = NewOutcomeClassDenomCountClient(tx.config)
	tx.OutcomeDenom = NewOutcomeDenomClient(tx.config)
	tx.OutcomeDenomCount = NewOutcomeDenomCountClient(tx.config)
	tx.OutcomeGroup = NewOutcomeGroupClient(tx.config)
	tx.OutcomeMeasure = NewOutcomeMeasureClient(tx.config)
	tx.OutcomeMeasurement = NewOutcomeMeasurementClient(tx.config)
	tx.OutcomeMeasuresModule = NewOutcomeMeasuresModuleClient(tx.config)
	tx.OutcomeOverview = NewOutcomeOverviewClient(tx.config)
	tx.ParticipantFlowModule = NewParticipantFlowModuleClient(tx.config)
	tx.PointOfContact = NewPointOfContactClient(tx.config)
	tx.ResultsDefinition = NewResultsDefinitionClient(tx.config)
	tx.Schedule = NewScheduleClient(tx.config)
	tx.SeriousEvent = NewSeriousEventClient(tx.config)
	tx.SeriousEventStats = NewSeriousEventStatsClient(tx.config)
	tx.StudyEligibility = NewStudyEligibilityClient(tx.config)
	tx.StudyLocation = NewStudyLocationClient(tx.config)
	tx.Task = NewTaskClient(tx.config)
	tx.UseCase = NewUseCaseClient(tx.config)
	tx.Vaccine = NewVaccineClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: AdverseEventsModule.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
