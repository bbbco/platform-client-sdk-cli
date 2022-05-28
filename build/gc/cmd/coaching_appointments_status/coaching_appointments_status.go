package coaching_appointments_status

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
	Description = utils.FormatUsageDescription("coaching_appointments_status", "SWAGGER_OVERRIDE_/api/v2/coaching/appointments/{appointmentId}/status", )
	coaching_appointments_statusCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("coaching_appointments_status"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(coaching_appointments_statusCmd)
}

func Cmdcoaching_appointments_status() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/coaching/appointments/{appointmentId}/status", utils.FormatPermissions([]string{ "coaching:appointmentStatus:edit",  }), utils.GenerateDevCentreLink("PATCH", "Coaching", "/api/v2/coaching/appointments/{appointmentId}/status")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "Updated status of the coaching appointment",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CoachingAppointmentStatusRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "The status is posted successfully",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CoachingAppointmentStatusResponse"
      }
    }
  }
}`)
	coaching_appointments_statusCmd.AddCommand(updateCmd)
	return coaching_appointments_statusCmd
}

var updateCmd = &cobra.Command{
	Use:   "update [appointmentId]",
	Short: "Update the status of a coaching appointment",
	Long:  "Update the status of a coaching appointment",
	Args:  utils.DetermineArgs([]string{ "appointmentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Coachingappointmentstatusrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/coaching/appointments/{appointmentId}/status"
		appointmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{appointmentId}", fmt.Sprintf("%v", appointmentId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "update"
		const httpMethod = "PATCH"
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

		utils.Render(results)
	},
}
