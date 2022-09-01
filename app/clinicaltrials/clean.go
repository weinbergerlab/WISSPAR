package clinicaltrials

import (
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/eventgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowachievement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowdropwithdraw"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowmilestone"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowreason"
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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriousevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriouseventstats"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/run"
)

func (t *ResultsContext) cleanClinicalTrial(trial *models.ClinicalTrial) error {
	if trial == nil {
		return nil
	}

	rd, err := trial.Edges.ResultsDefinitionOrErr()
	if rd != nil && err == nil {
		if err := run.Go(t.ctx,
			func() error { return t.cleanParticipantFlowModule(rd) },
			func() error { return t.cleanBaselineCharacteristicsModule(rd) },
			func() error { return t.cleanOutcomeMeasuresModule(rd) },
			func() error { return t.cleanAdverseEventsModule(rd) },
			func() error { return t.cleanMoreInfoModule(rd) },
		); err != nil {
			return err
		}

		if err := t.client.ResultsDefinition.DeleteOne(rd).Exec(t.ctx); err != nil {
			return err
		}
	}

	return t.client.ClinicalTrial.DeleteOne(trial).Exec(t.ctx)
}

func (t *ResultsContext) cleanParticipantFlowModule(rd *models.ResultsDefinition) error {
	pfm, err := rd.Edges.ParticipantFlowModuleOrErr()
	if pfm == nil || err != nil {
		return nil
	}

	if len(pfm.Edges.FlowGroupList) > 0 {
		flowGroupIDs := []int{}
		for _, flowGroup := range pfm.Edges.FlowGroupList {
			flowGroupIDs = append(flowGroupIDs, flowGroup.ID)
		}
		if _, err := t.client.FlowGroup.Delete().Where(flowgroup.IDIn(flowGroupIDs...)).Exec(t.ctx); err != nil {
			return err
		}
	}

	if len(pfm.Edges.FlowPeriodList) > 0 {
		flowPeriodIDs := []int{}
		for _, flowPeriod := range pfm.Edges.FlowPeriodList {
			flowPeriodIDs = append(flowPeriodIDs, flowPeriod.ID)


			if len(flowPeriod.Edges.FlowDropWithdrawList) > 0 {
				flowDropWithdrawIDs := []int{}
				for _, flowDropWithdraw := range flowPeriod.Edges.FlowDropWithdrawList {
					flowDropWithdrawIDs = append(flowDropWithdrawIDs, flowDropWithdraw.ID)

					if len(flowDropWithdraw.Edges.FlowReasonList) > 0 {
						flowReasonIDs := []int{}
						for _, flowReason := range flowDropWithdraw.Edges.FlowReasonList {
							flowReasonIDs = append(flowReasonIDs, flowReason.ID)
						}
						if _, err := t.client.FlowReason.Delete().Where(flowreason.IDIn(flowReasonIDs...)).Exec(t.ctx); err != nil {
							return err
						}
					}
				}
				if _, err := t.client.FlowDropWithdraw.Delete().Where(flowdropwithdraw.IDIn(flowDropWithdrawIDs...)).Exec(t.ctx); err != nil {
					return err
				}
			}

			if len(flowPeriod.Edges.FlowMilestoneList) > 0 {
				flowMilestoneIDs := []int{}
				for _, flowMilestone := range flowPeriod.Edges.FlowMilestoneList {
					flowMilestoneIDs = append(flowMilestoneIDs, flowMilestone.ID)

					if len(flowMilestone.Edges.FlowAchievementList) > 0 {
						flowAchievementIDs := []int{}
						for _, flowAchievement := range flowMilestone.Edges.FlowAchievementList {
							flowAchievementIDs = append(flowAchievementIDs, flowAchievement.ID)
						}
						if _, err := t.client.FlowAchievement.Delete().Where(flowachievement.IDIn(flowAchievementIDs...)).Exec(t.ctx); err != nil {
							return err
						}
					}
				}
				if _, err := t.client.FlowMilestone.Delete().Where(flowmilestone.IDIn(flowMilestoneIDs...)).Exec(t.ctx); err != nil {
					return err
				}
			}

		}
		if _, err := t.client.FlowPeriod.Delete().Where(flowperiod.IDIn(flowPeriodIDs...)).Exec(t.ctx); err != nil {
			return err
		}
	}

	return t.client.ParticipantFlowModule.DeleteOne(pfm).Exec(t.ctx)
}

func (t *ResultsContext) cleanBaselineCharacteristicsModule(rd *models.ResultsDefinition) error {
	bcm, err := rd.Edges.BaselineCharacteristicsModuleOrErr()
	if bcm == nil || err != nil {
		return nil
	}

	if len(bcm.Edges.BaselineGroupList) > 0 {
		baselineGroupIDs := []int{}
		for _, baselineGroup := range bcm.Edges.BaselineGroupList {
			baselineGroupIDs = append(baselineGroupIDs, baselineGroup.ID)
		}
		if _, err := t.client.BaselineGroup.Delete().Where(baselinegroup.IDIn(baselineGroupIDs...)).Exec(t.ctx); err != nil {
			return err
		}
	}

	if len(bcm.Edges.BaselineDenomList) > 0 {
		baselineDenomIDs := []int{}
		for _, baselineDenom := range bcm.Edges.BaselineDenomList {
			baselineDenomIDs = append(baselineDenomIDs, baselineDenom.ID)

			if len(baselineDenom.Edges.BaselineDenomCountList) > 0 {
				baselineDenomCountIDs := []int{}
				for _, baselineDenomCount := range baselineDenom.Edges.BaselineDenomCountList {
					baselineDenomCountIDs = append(baselineDenomCountIDs, baselineDenomCount.ID)
				}
				if _, err := t.client.BaselineDenomCount.Delete().Where(baselinedenomcount.IDIn(baselineDenomCountIDs...)).Exec(t.ctx); err != nil {
					return err
				}
			}

		}
		if _, err := t.client.BaselineDenom.Delete().Where(baselinedenom.IDIn(baselineDenomIDs...)).Exec(t.ctx); err != nil {
			return err
		}
	}

	if len(bcm.Edges.BaselineMeasureList) > 0 {
		baselineMeasureIDs := []int{}
		for _, baselineMeasure := range bcm.Edges.BaselineMeasureList {
			baselineMeasureIDs = append(baselineMeasureIDs, baselineMeasure.ID)

			if len(baselineMeasure.Edges.BaselineMeasureDenomList) > 0 {
				baselineMeasureDenomIDs := []int{}
				for _, baselineMeasureDenom := range baselineMeasure.Edges.BaselineMeasureDenomList {
					baselineMeasureDenomIDs = append(baselineMeasureDenomIDs, baselineMeasureDenom.ID)

					if len(baselineMeasureDenom.Edges.BaselineMeasureDenomCountList) > 0 {
						baselineMeasureDenomCountIDs := []int{}
						for _, baselineMeasureDenomCount := range baselineMeasureDenom.Edges.BaselineMeasureDenomCountList {
							baselineMeasureDenomCountIDs = append(baselineMeasureDenomCountIDs, baselineMeasureDenomCount.ID)
						}
						if _, err := t.client.BaselineMeasureDenomCount.Delete().Where(baselinemeasuredenomcount.IDIn(baselineMeasureDenomCountIDs...)).Exec(t.ctx); err != nil {
							return err
						}
					}

				}
				if _, err := t.client.BaselineMeasureDenom.Delete().Where(baselinemeasuredenom.IDIn(baselineMeasureDenomIDs...)).Exec(t.ctx); err != nil {
					return err
				}
			}

		}
		if _, err := t.client.BaselineMeasure.Delete().Where(baselinemeasure.IDIn(baselineMeasureIDs...)).Exec(t.ctx); err != nil {
			return err
		}
	}

	return t.client.BaselineCharacteristicsModule.DeleteOne(bcm).Exec(t.ctx)
}

func (t *ResultsContext) cleanOutcomeMeasuresModule(rd *models.ResultsDefinition) error {
	om, err := rd.Edges.OutcomeMeasuresModuleOrErr()
	if om == nil || err != nil {
		return nil
	}

	if len(om.Edges.OutcomeMeasureList) > 0 {
		outcomeMeasureIDs := []int{}
		for _, outcomeMeasure := range om.Edges.OutcomeMeasureList {
			outcomeMeasureIDs = append(outcomeMeasureIDs, outcomeMeasure.ID)

			if len(outcomeMeasure.Edges.OutcomeGroupList) > 0 {
				outcomeGroupIDs := []int{}
				for _, outcomeGroup := range outcomeMeasure.Edges.OutcomeGroupList {
					outcomeGroupIDs = append(outcomeGroupIDs, outcomeGroup.ID)
				}
				if _, err := t.client.OutcomeGroup.Delete().Where(outcomegroup.IDIn(outcomeGroupIDs...)).Exec(t.ctx); err != nil {
					return err
				}
			}

			if len(outcomeMeasure.Edges.OutcomeDenomList) > 0 {
				outcomeDenomIDs := []int{}
				for _, outcomeDenom := range outcomeMeasure.Edges.OutcomeDenomList {
					outcomeDenomIDs = append(outcomeDenomIDs, outcomeDenom.ID)

					if len(outcomeDenom.Edges.OutcomeDenomCountList) > 0 {
						outcomeDenomCountIDs := []int{}
						for _, outcomeDenomCount := range outcomeDenom.Edges.OutcomeDenomCountList {
							outcomeDenomCountIDs = append(outcomeDenomCountIDs, outcomeDenomCount.ID)
						}
						if _, err := t.client.OutcomeDenomCount.Delete().Where(outcomedenomcount.IDIn(outcomeDenomCountIDs...)).Exec(t.ctx); err != nil {
							return err
						}
					}

				}
				if _, err := t.client.OutcomeDenom.Delete().Where(outcomedenom.IDIn(outcomeDenomIDs...)).Exec(t.ctx); err != nil {
					return err
				}
			}

			if len(outcomeMeasure.Edges.OutcomeClassList) > 0 {
				outcomeClassIDs := []int{}
				for _, outcomeClass := range outcomeMeasure.Edges.OutcomeClassList {
					outcomeClassIDs = append(outcomeClassIDs, outcomeClass.ID)

					if len(outcomeClass.Edges.OutcomeCategoryList) > 0 {
						outcomeCategoryIDs := []int{}
						for _, outcomeCategory := range outcomeClass.Edges.OutcomeCategoryList {
							outcomeCategoryIDs = append(outcomeCategoryIDs, outcomeCategory.ID)

							if len(outcomeCategory.Edges.OutcomeMeasurementList) > 0 {
								outcomeMeasurementIDs := []int{}
								for _, outcomeMeasurement := range outcomeCategory.Edges.OutcomeMeasurementList {
									outcomeMeasurementIDs = append(outcomeMeasurementIDs, outcomeMeasurement.ID)
								}
								if _, err := t.client.OutcomeMeasurement.Delete().Where(outcomemeasurement.IDIn(outcomeMeasurementIDs...)).Exec(t.ctx); err != nil {
									return err
								}
							}

						}
						if _, err := t.client.OutcomeCategory.Delete().Where(outcomecategory.IDIn(outcomeCategoryIDs...)).Exec(t.ctx); err != nil {
							return err
						}
					}

					if len(outcomeClass.Edges.OutcomeClassDenomList) > 0 {
						outcomeClassDenomIDs := []int{}
						for _, outcomeClassDenom := range outcomeClass.Edges.OutcomeClassDenomList {
							outcomeClassDenomIDs = append(outcomeClassDenomIDs, outcomeClassDenom.ID)

							if len(outcomeClassDenom.Edges.OutcomeClassDenomCountList) > 0 {
								outcomeClassDenomCountIDs := []int{}
								for _, outcomeClassDenomCount := range outcomeClassDenom.Edges.OutcomeClassDenomCountList {
									outcomeClassDenomCountIDs = append(outcomeClassDenomCountIDs, outcomeClassDenomCount.ID)
								}
								if _, err := t.client.OutcomeClassDenomCount.Delete().Where(outcomeclassdenomcount.IDIn(outcomeClassDenomCountIDs...)).Exec(t.ctx); err != nil {
									return err
								}
							}

						}
						if _, err := t.client.OutcomeClassDenom.Delete().Where(outcomeclassdenom.IDIn(outcomeClassDenomIDs...)).Exec(t.ctx); err != nil {
							return err
						}
					}

				}
				if _, err := t.client.OutcomeClass.Delete().Where(outcomeclass.IDIn(outcomeClassIDs...)).Exec(t.ctx); err != nil {
					return err
				}
			}

			if len(outcomeMeasure.Edges.OutcomeAnalysisList) > 0 {
				outcomeAnalyisIDs := []int{}
				for _, outcomeAnalyis := range outcomeMeasure.Edges.OutcomeAnalysisList {
					outcomeAnalyisIDs = append(outcomeAnalyisIDs, outcomeAnalyis.ID)

					if len(outcomeAnalyis.Edges.OutcomeAnalysisGroupIDList) > 0 {
						outcomeAnalyisGroupIDIDs := []int{}
						for _, outcomeAnalyisGroupID := range outcomeAnalyis.Edges.OutcomeAnalysisGroupIDList {
							outcomeAnalyisGroupIDIDs = append(outcomeAnalyisGroupIDIDs, outcomeAnalyisGroupID.ID)
						}
						if _, err := t.client.OutcomeAnalysisGroupID.Delete().Where(outcomeanalysisgroupid.IDIn(outcomeAnalyisGroupIDIDs...)).Exec(t.ctx); err != nil {
							return err
						}
					}

				}
				if _, err := t.client.OutcomeAnalysis.Delete().Where(outcomeanalysis.IDIn(outcomeAnalyisIDs...)).Exec(t.ctx); err != nil {
					return err
				}
			}

		}
		if _, err := t.client.OutcomeMeasure.Delete().Where(outcomemeasure.IDIn(outcomeMeasureIDs...)).Exec(t.ctx); err != nil {
			return err
		}
	}

	return t.client.OutcomeMeasuresModule.DeleteOne(om).Exec(t.ctx)
}

func (t *ResultsContext) cleanAdverseEventsModule(rd *models.ResultsDefinition) error {
	aem, err := rd.Edges.AdverseEventsModuleOrErr()
	if aem == nil || err != nil {
		return nil
	}

	if len(aem.Edges.SeriousEventList) > 0 {
		seriousEventIDs := []int{}
		for _, seriousEvent := range aem.Edges.SeriousEventList {
			seriousEventIDs = append(seriousEventIDs, seriousEvent.ID)

			if len(seriousEvent.Edges.SeriousEventStatsList) > 0 {
				seriousEventStatsIDs := []int{}
				for _, seriousEventStats := range seriousEvent.Edges.SeriousEventStatsList {
					seriousEventStatsIDs = append(seriousEventStatsIDs, seriousEventStats.ID)
				}
				if _, err := t.client.SeriousEventStats.Delete().Where(seriouseventstats.IDIn(seriousEventStatsIDs...)).Exec(t.ctx); err != nil {
					return err
				}
			}

		}
		if _, err := t.client.SeriousEvent.Delete().Where(seriousevent.IDIn(seriousEventIDs...)).Exec(t.ctx); err != nil {
			return err
		}
	}

	if len(aem.Edges.EventGroupList) > 0 {
		eventGroupIDs := []int{}
		for _, eventGroup := range aem.Edges.EventGroupList {
			eventGroupIDs = append(eventGroupIDs, eventGroup.ID)
		}
		if _, err := t.client.EventGroup.Delete().Where(eventgroup.IDIn(eventGroupIDs...)).Exec(t.ctx); err != nil {
			return err
		}
	}

	if len(aem.Edges.OtherEventList) > 0 {
		otherEventIDs := []int{}
		for _, otherEvent := range aem.Edges.OtherEventList {
			otherEventIDs = append(otherEventIDs, otherEvent.ID)

			if len(otherEvent.Edges.OtherEventStatsList) > 0 {
				otherEventStatsIDs := []int{}
				for _, otherEventStats := range otherEvent.Edges.OtherEventStatsList {
					otherEventStatsIDs = append(otherEventStatsIDs, otherEventStats.ID)
				}
				if _, err := t.client.OtherEventStats.Delete().Where(othereventstats.IDIn(otherEventStatsIDs...)).Exec(t.ctx); err != nil {
					return err
				}
			}
		}
		if _, err := t.client.OtherEvent.Delete().Where(otherevent.IDIn(otherEventIDs...)).Exec(t.ctx); err != nil {
			return err
		}
	}

	return t.client.AdverseEventsModule.DeleteOne(aem).Exec(t.ctx)
}

func (t *ResultsContext) cleanMoreInfoModule(rd *models.ResultsDefinition) error {
	mi, err := rd.Edges.MoreInfoModuleOrErr()
	if mi == nil || err != nil {
		return nil
	}

	poc, err := mi.Edges.PointOfContactOrErr()
	if poc != nil && err == nil {
		if err := t.client.PointOfContact.DeleteOne(poc).Exec(t.ctx); err != nil {
			return err
		}
	}

	ca, err := mi.Edges.CertainAgreementOrErr()
	if ca != nil && err == nil {
		if err := t.client.CertainAgreement.DeleteOne(ca).Exec(t.ctx); err != nil {
			return err
		}
	}

	return t.client.MoreInfoModule.DeleteOne(mi).Exec(t.ctx)
}
