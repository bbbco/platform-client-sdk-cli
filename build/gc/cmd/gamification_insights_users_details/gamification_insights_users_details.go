package gamification_insights_users_details

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
	Description = utils.FormatUsageDescription("gamification_insights_users_details", "SWAGGER_OVERRIDE_/api/v2/gamification/insights/users/{userId}/details", )
	gamification_insights_users_detailsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_insights_users_details"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_insights_users_detailsCmd)
}

func Cmdgamification_insights_users_details() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "filterType", "", "Filter type for the query request. - REQUIRED Valid values: PerformanceProfile, Division")
	utils.AddFlag(getCmd.Flags(), "string", "filterId", "", "ID for the filter type. - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "string", "granularity", "", "Granularity - REQUIRED Valid values: Weekly, Monthly")
	utils.AddFlag(getCmd.Flags(), "time.Time", "comparativePeriodStartWorkday", "", "The start work day of comparative period. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "time.Time", "primaryPeriodStartWorkday", "", "The start work day of primary period. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/gamification/insights/users/{userId}/details", utils.FormatPermissions([]string{ "gamification:insights:viewAll",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/insights/users/{userId}/details")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("filterType")
	getCmd.MarkFlagRequired("filterId")
	getCmd.MarkFlagRequired("granularity")
	getCmd.MarkFlagRequired("comparativePeriodStartWorkday")
	getCmd.MarkFlagRequired("primaryPeriodStartWorkday")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/InsightsDetails"
      }
    }
  }
}`)
	gamification_insights_users_detailsCmd.AddCommand(getCmd)
	return gamification_insights_users_detailsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [userId]",
	Short: "Get insights details for the user",
	Long:  "Get insights details for the user",
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/insights/users/{userId}/details"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

		filterType := utils.GetFlag(cmd.Flags(), "string", "filterType")
		if filterType != "" {
			queryParams["filterType"] = filterType
		}
		filterId := utils.GetFlag(cmd.Flags(), "string", "filterId")
		if filterId != "" {
			queryParams["filterId"] = filterId
		}
		granularity := utils.GetFlag(cmd.Flags(), "string", "granularity")
		if granularity != "" {
			queryParams["granularity"] = granularity
		}
		comparativePeriodStartWorkday := utils.GetFlag(cmd.Flags(), "time.Time", "comparativePeriodStartWorkday")
		if comparativePeriodStartWorkday != "" {
			queryParams["comparativePeriodStartWorkday"] = comparativePeriodStartWorkday
		}
		primaryPeriodStartWorkday := utils.GetFlag(cmd.Flags(), "time.Time", "primaryPeriodStartWorkday")
		if primaryPeriodStartWorkday != "" {
			queryParams["primaryPeriodStartWorkday"] = primaryPeriodStartWorkday
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
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
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
