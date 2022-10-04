package workforcemanagement_businessunits_weeks_schedules_reschedule

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
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_weeks_schedules_reschedule", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule", )
	workforcemanagement_businessunits_weeks_schedules_rescheduleCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_weeks_schedules_reschedule"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_weeks_schedules_rescheduleCmd)
}

func Cmdworkforcemanagement_businessunits_weeks_schedules_reschedule() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule", utils.FormatPermissions([]string{ "wfm:schedule:generate",  }), utils.GenerateDevCentreLink("POST", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "body",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/BuRescheduleRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/BuAsyncScheduleRunResponse"
      }
    }
  }
}`)
	workforcemanagement_businessunits_weeks_schedules_rescheduleCmd.AddCommand(createCmd)
	return workforcemanagement_businessunits_weeks_schedules_rescheduleCmd
}

var createCmd = &cobra.Command{
	Use:   "create [businessUnitId] [weekId] [scheduleId]",
	Short: "Start a rescheduling run",
	Long:  "Start a rescheduling run",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", "weekId", "scheduleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Bureschedulerequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
		weekId, args := args[0], args[1:]
		path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
		scheduleId, args := args[0], args[1:]
		path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)

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
