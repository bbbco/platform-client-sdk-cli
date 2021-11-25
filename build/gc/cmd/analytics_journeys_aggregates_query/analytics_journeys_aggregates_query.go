package analytics_journeys_aggregates_query

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
	Description = utils.FormatUsageDescription("analytics_journeys_aggregates_query", "SWAGGER_OVERRIDE_/api/v2/analytics/journeys/aggregates/query", )
	analytics_journeys_aggregates_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_journeys_aggregates_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_journeys_aggregates_queryCmd)
}

func Cmdanalytics_journeys_aggregates_query() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/analytics/journeys/aggregates/query", utils.FormatPermissions([]string{ "analytics:journeyAggregate:view",  }), utils.GenerateDevCentreLink("POST", "Analytics", "/api/v2/analytics/journeys/aggregates/query")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "query",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/JourneyAggregationQuery"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/JourneyAggregateQueryResponse"
  }
}`)
	analytics_journeys_aggregates_queryCmd.AddCommand(createCmd)
	
	return analytics_journeys_aggregates_queryCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Query for journey aggregates",
	Long:  "Query for journey aggregates",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Journeyaggregationquery{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/analytics/journeys/aggregates/query"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "create"
		const httpMethod = "POST"
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
