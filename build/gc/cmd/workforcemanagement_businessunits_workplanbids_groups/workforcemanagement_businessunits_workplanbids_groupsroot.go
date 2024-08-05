package workforcemanagement_businessunits_workplanbids_groups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_workplanbids_groups_summary"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_workplanbids_groups_preferences"
)

func init() {
	workforcemanagement_businessunits_workplanbids_groupsCmd.AddCommand(workforcemanagement_businessunits_workplanbids_groups_summary.Cmdworkforcemanagement_businessunits_workplanbids_groups_summary())
	workforcemanagement_businessunits_workplanbids_groupsCmd.AddCommand(workforcemanagement_businessunits_workplanbids_groups_preferences.Cmdworkforcemanagement_businessunits_workplanbids_groups_preferences())
	workforcemanagement_businessunits_workplanbids_groupsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_businessunits_workplanbids_groupsCmd.Short, workforcemanagement_businessunits_workplanbids_groups_summary.Description, workforcemanagement_businessunits_workplanbids_groups_preferences.Description, )
	workforcemanagement_businessunits_workplanbids_groupsCmd.Long = workforcemanagement_businessunits_workplanbids_groupsCmd.Short
}
