// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
)

// Option function to configure the client.
type Option func(*config)

// Config is the configuration for the client and its builder.
type config struct {
	// driver used for executing database requests.
	driver dialect.Driver
	// debug enable a debug logging.
	debug bool
	// log used for logging on debug mode.
	log func(...interface{})
	// hooks to execute on mutations.
	hooks *hooks
}

// hooks per client, for fast access.
type hooks struct {
	AdverseEventsModule           []ent.Hook
	BaselineCategory              []ent.Hook
	BaselineCharacteristicsModule []ent.Hook
	BaselineClass                 []ent.Hook
	BaselineClassDenom            []ent.Hook
	BaselineClassDenomCount       []ent.Hook
	BaselineDenom                 []ent.Hook
	BaselineDenomCount            []ent.Hook
	BaselineGroup                 []ent.Hook
	BaselineMeasure               []ent.Hook
	BaselineMeasureDenom          []ent.Hook
	BaselineMeasureDenomCount     []ent.Hook
	BaselineMeasurement           []ent.Hook
	CertainAgreement              []ent.Hook
	ClinicalTrial                 []ent.Hook
	ClinicalTrialMetadata         []ent.Hook
	DoseDescription               []ent.Hook
	EventGroup                    []ent.Hook
	FlowAchievement               []ent.Hook
	FlowDropWithdraw              []ent.Hook
	FlowGroup                     []ent.Hook
	FlowMilestone                 []ent.Hook
	FlowPeriod                    []ent.Hook
	FlowReason                    []ent.Hook
	ImmunocompromisedGroups       []ent.Hook
	Manufacturer                  []ent.Hook
	MoreInfoModule                []ent.Hook
	OtherEvent                    []ent.Hook
	OtherEventStats               []ent.Hook
	OutcomeAnalysis               []ent.Hook
	OutcomeAnalysisGroupID        []ent.Hook
	OutcomeCategory               []ent.Hook
	OutcomeClass                  []ent.Hook
	OutcomeClassDenom             []ent.Hook
	OutcomeClassDenomCount        []ent.Hook
	OutcomeDenom                  []ent.Hook
	OutcomeDenomCount             []ent.Hook
	OutcomeGroup                  []ent.Hook
	OutcomeMeasure                []ent.Hook
	OutcomeMeasurement            []ent.Hook
	OutcomeMeasuresModule         []ent.Hook
	OutcomeOverview               []ent.Hook
	ParticipantFlowModule         []ent.Hook
	PointOfContact                []ent.Hook
	ResultsDefinition             []ent.Hook
	Schedule                      []ent.Hook
	SeriousEvent                  []ent.Hook
	SeriousEventStats             []ent.Hook
	StudyEligibility              []ent.Hook
	StudyLocation                 []ent.Hook
	Task                          []ent.Hook
	UseCase                       []ent.Hook
	Vaccine                       []ent.Hook
}

// Options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...interface{})) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}
