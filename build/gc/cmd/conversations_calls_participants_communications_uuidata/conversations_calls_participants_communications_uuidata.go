package conversations_calls_participants_communications_uuidata

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
	Description = utils.FormatUsageDescription("conversations_calls_participants_communications_uuidata", "SWAGGER_OVERRIDE_/api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata", )
	conversations_calls_participants_communications_uuidataCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_calls_participants_communications_uuidata"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_calls_participants_communications_uuidataCmd)
}

func Cmdconversations_calls_participants_communications_uuidata() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PUT", "Conversations", "/api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "UUIData Request",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/SetUuiDataRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Empty"
  }
}`)
	conversations_calls_participants_communications_uuidataCmd.AddCommand(updateCmd)
	
	return conversations_calls_participants_communications_uuidataCmd
}

var updateCmd = &cobra.Command{
	Use:   "update [conversationId] [participantId] [communicationId]",
	Short: "Set uuiData to be sent on future commands.",
	Long:  "Set uuiData to be sent on future commands.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "participantId", "communicationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Setuuidatarequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		participantId, args := args[0], args[1:]
		path = strings.Replace(path, "{participantId}", fmt.Sprintf("%v", participantId), -1)
		communicationId, args := args[0], args[1:]
		path = strings.Replace(path, "{communicationId}", fmt.Sprintf("%v", communicationId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "update"
		const httpMethod = "PUT"
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
