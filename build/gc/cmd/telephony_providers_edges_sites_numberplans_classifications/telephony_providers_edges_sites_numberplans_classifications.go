package telephony_providers_edges_sites_numberplans_classifications

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
	Description = utils.FormatUsageDescription("telephony_providers_edges_sites_numberplans_classifications", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/sites/{siteId}/numberplans/classifications", )
	telephony_providers_edges_sites_numberplans_classificationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_sites_numberplans_classifications"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_sites_numberplans_classificationsCmd)
}

func Cmdtelephony_providers_edges_sites_numberplans_classifications() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "classification", "", "Classification")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/sites/{siteId}/numberplans/classifications", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/sites/{siteId}/numberplans/classifications")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "type" : "array",
    "uniqueItems" : true,
    "items" : {
      "type" : "string"
    }
  }
}`)
	telephony_providers_edges_sites_numberplans_classificationsCmd.AddCommand(listCmd)
	
	return telephony_providers_edges_sites_numberplans_classificationsCmd
}

var listCmd = &cobra.Command{
	Use:   "list [siteId]",
	Short: "Get a list of Classifications for this Site",
	Long:  "Get a list of Classifications for this Site",
	Args:  utils.DetermineArgs([]string{ "siteId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/sites/{siteId}/numberplans/classifications"
		siteId, args := args[0], args[1:]
		path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)

		classification := utils.GetFlag(cmd.Flags(), "string", "classification")
		if classification != "" {
			queryParams["classification"] = classification
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "list"
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
