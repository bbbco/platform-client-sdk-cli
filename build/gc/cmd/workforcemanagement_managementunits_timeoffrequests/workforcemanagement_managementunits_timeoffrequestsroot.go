package workforcemanagement_managementunits_timeoffrequests

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_timeoffrequests_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_timeoffrequests_integrationstatus"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_timeoffrequests_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_timeoffrequests_waitlistpositions"
)

func init() {
	workforcemanagement_managementunits_timeoffrequestsCmd.AddCommand(workforcemanagement_managementunits_timeoffrequests_users.Cmdworkforcemanagement_managementunits_timeoffrequests_users())
	workforcemanagement_managementunits_timeoffrequestsCmd.AddCommand(workforcemanagement_managementunits_timeoffrequests_integrationstatus.Cmdworkforcemanagement_managementunits_timeoffrequests_integrationstatus())
	workforcemanagement_managementunits_timeoffrequestsCmd.AddCommand(workforcemanagement_managementunits_timeoffrequests_query.Cmdworkforcemanagement_managementunits_timeoffrequests_query())
	workforcemanagement_managementunits_timeoffrequestsCmd.AddCommand(workforcemanagement_managementunits_timeoffrequests_waitlistpositions.Cmdworkforcemanagement_managementunits_timeoffrequests_waitlistpositions())
	workforcemanagement_managementunits_timeoffrequestsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_managementunits_timeoffrequestsCmd.Short, workforcemanagement_managementunits_timeoffrequests_users.Description, workforcemanagement_managementunits_timeoffrequests_integrationstatus.Description, workforcemanagement_managementunits_timeoffrequests_query.Description, workforcemanagement_managementunits_timeoffrequests_waitlistpositions.Description, )
	workforcemanagement_managementunits_timeoffrequestsCmd.Long = workforcemanagement_managementunits_timeoffrequestsCmd.Short
}
