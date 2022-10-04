package notifications_availabletopics

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
	Description = utils.FormatUsageDescription("notifications_availabletopics", "SWAGGER_OVERRIDE_/api/v2/notifications/availabletopics", )
	notifications_availabletopicsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("notifications_availabletopics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(notifications_availabletopicsCmd)
}

func Cmdnotifications_availabletopics() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: description, enforced, schema, visibility, transports, publicApiTemplateUriPaths, requiresPermissions, permissionDetails, topicParameters")
	utils.AddFlag(listCmd.Flags(), "bool", "includePreview", "true", "Whether or not to include Preview topics")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/notifications/availabletopics", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Notifications", "/api/v2/notifications/availabletopics")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AvailableTopicEntityListing"
      }
    }
  }
}`)
	notifications_availabletopicsCmd.AddCommand(listCmd)
	return notifications_availabletopicsCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get available notification topics.",
	Long:  "Get available notification topics.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/notifications/availabletopics"

		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		includePreview := utils.GetFlag(cmd.Flags(), "bool", "includePreview")
		if includePreview != "" {
			queryParams["includePreview"] = includePreview
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
