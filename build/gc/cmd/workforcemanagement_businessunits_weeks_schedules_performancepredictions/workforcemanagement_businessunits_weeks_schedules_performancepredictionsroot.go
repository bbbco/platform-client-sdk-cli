package workforcemanagement_businessunits_weeks_schedules_performancepredictions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_weeks_schedules_performancepredictions_recalculations"
)

func init() {
	workforcemanagement_businessunits_weeks_schedules_performancepredictionsCmd.AddCommand(workforcemanagement_businessunits_weeks_schedules_performancepredictions_recalculations.Cmdworkforcemanagement_businessunits_weeks_schedules_performancepredictions_recalculations())
	workforcemanagement_businessunits_weeks_schedules_performancepredictionsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_businessunits_weeks_schedules_performancepredictionsCmd.Short, workforcemanagement_businessunits_weeks_schedules_performancepredictions_recalculations.Description, )
	workforcemanagement_businessunits_weeks_schedules_performancepredictionsCmd.Long = workforcemanagement_businessunits_weeks_schedules_performancepredictionsCmd.Short
}
