package workforcemanagement_businessunits_timeofflimits_values

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_timeofflimits_values_query"
)

func init() {
	workforcemanagement_businessunits_timeofflimits_valuesCmd.AddCommand(workforcemanagement_businessunits_timeofflimits_values_query.Cmdworkforcemanagement_businessunits_timeofflimits_values_query())
	workforcemanagement_businessunits_timeofflimits_valuesCmd.Short = utils.GenerateCustomDescription(workforcemanagement_businessunits_timeofflimits_valuesCmd.Short, workforcemanagement_businessunits_timeofflimits_values_query.Description, )
	workforcemanagement_businessunits_timeofflimits_valuesCmd.Long = workforcemanagement_businessunits_timeofflimits_valuesCmd.Short
}
