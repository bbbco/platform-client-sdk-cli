package flows_instances_settings_loglevels

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_instances_settings_loglevels_default"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_instances_settings_loglevels_characteristics"
)

func init() {
	flows_instances_settings_loglevelsCmd.AddCommand(flows_instances_settings_loglevels_default.Cmdflows_instances_settings_loglevels_default())
	flows_instances_settings_loglevelsCmd.AddCommand(flows_instances_settings_loglevels_characteristics.Cmdflows_instances_settings_loglevels_characteristics())
	flows_instances_settings_loglevelsCmd.Short = utils.GenerateCustomDescription(flows_instances_settings_loglevelsCmd.Short, flows_instances_settings_loglevels_default.Description, flows_instances_settings_loglevels_characteristics.Description, )
	flows_instances_settings_loglevelsCmd.Long = flows_instances_settings_loglevelsCmd.Short
}
