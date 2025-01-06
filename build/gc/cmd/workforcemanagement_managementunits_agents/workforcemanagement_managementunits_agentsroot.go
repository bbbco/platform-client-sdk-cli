package workforcemanagement_managementunits_agents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_agents_workplans"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_agents_shifttrades"
)

func init() {
	workforcemanagement_managementunits_agentsCmd.AddCommand(workforcemanagement_managementunits_agents_workplans.Cmdworkforcemanagement_managementunits_agents_workplans())
	workforcemanagement_managementunits_agentsCmd.AddCommand(workforcemanagement_managementunits_agents_shifttrades.Cmdworkforcemanagement_managementunits_agents_shifttrades())
	workforcemanagement_managementunits_agentsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_managementunits_agentsCmd.Short, workforcemanagement_managementunits_agents_workplans.Description, workforcemanagement_managementunits_agents_shifttrades.Description, )
	workforcemanagement_managementunits_agentsCmd.Long = workforcemanagement_managementunits_agentsCmd.Short
}
