package workforcemanagement_timeoffrequests_integrationstatus

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_timeoffrequests_integrationstatus", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/timeoffrequests/integrationstatus")
	workforcemanagement_timeoffrequests_integrationstatusCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_timeoffrequests_integrationstatus"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_timeoffrequests_integrationstatusCmd)
}

func Cmdworkforcemanagement_timeoffrequests_integrationstatus() *cobra.Command {
	return workforcemanagement_timeoffrequests_integrationstatusCmd
}
