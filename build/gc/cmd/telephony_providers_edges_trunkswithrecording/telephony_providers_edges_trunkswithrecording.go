package telephony_providers_edges_trunkswithrecording

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("telephony_providers_edges_trunkswithrecording", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/trunkswithrecording", )
	telephony_providers_edges_trunkswithrecordingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_trunkswithrecording"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_trunkswithrecordingCmd)
}

func Cmdtelephony_providers_edges_trunkswithrecording() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "trunkType", "", "The type of this trunk base. Valid values: EXTERNAL, PHONE, EDGE")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/trunkswithrecording", utils.FormatPermissions([]string{ "recording:retentionPolicy:view", "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/trunkswithrecording")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TrunkRecordingEnabledCount"
  }
}`)
	telephony_providers_edges_trunkswithrecordingCmd.AddCommand(getCmd)
	
	return telephony_providers_edges_trunkswithrecordingCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Counts of trunks that have recording disabled or enabled",
	Long:  "Get Counts of trunks that have recording disabled or enabled",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/trunkswithrecording"

		trunkType := utils.GetFlag(cmd.Flags(), "string", "trunkType")
		if trunkType != "" {
			queryParams["trunkType"] = trunkType
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "get"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
