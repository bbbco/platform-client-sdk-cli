package textbots_botflows_sessions_turns

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
	Description = utils.FormatUsageDescription("textbots_botflows_sessions_turns", "SWAGGER_OVERRIDE_/api/v2/textbots/botflows/sessions/{sessionId}/turns", )
	textbots_botflows_sessions_turnsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("textbots_botflows_sessions_turns"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(textbots_botflows_sessions_turnsCmd)
}

func Cmdtextbots_botflows_sessions_turns() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/textbots/botflows/sessions/{sessionId}/turns", utils.FormatPermissions([]string{ "textbots:botFlowSession:execute",  }), utils.GenerateDevCentreLink("POST", "Textbots", "/api/v2/textbots/botflows/sessions/{sessionId}/turns")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "turnRequest",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/TextBotFlowTurnRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TextBotFlowTurnResponse"
  }
}`)
	textbots_botflows_sessions_turnsCmd.AddCommand(createCmd)
	
	return textbots_botflows_sessions_turnsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [sessionId]",
	Short: "Issue a bot flow turn event",
	Long:  "Issue a bot flow turn event",
	Args:  utils.DetermineArgs([]string{ "sessionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Textbotflowturnrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/textbots/botflows/sessions/{sessionId}/turns"
		sessionId, args := args[0], args[1:]
		path = strings.Replace(path, "{sessionId}", fmt.Sprintf("%v", sessionId), -1)

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
