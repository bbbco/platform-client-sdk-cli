package externalcontacts_bulk_notes_remove

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
	Description = utils.FormatUsageDescription("externalcontacts_bulk_notes_remove", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/bulk/notes/remove", )
	externalcontacts_bulk_notes_removeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_bulk_notes_remove"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_bulk_notes_removeCmd)
}

func Cmdexternalcontacts_bulk_notes_remove() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/externalcontacts/bulk/notes/remove", utils.FormatPermissions([]string{ "externalContacts:contact:delete", "externalContacts:externalOrganization:delete",  }), utils.GenerateDevCentreLink("POST", "External Contacts", "/api/v2/externalcontacts/bulk/notes/remove")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Note ids",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/BulkIdsRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/BulkDeleteResponse"
  }
}`)
	externalcontacts_bulk_notes_removeCmd.AddCommand(createCmd)
	
	return externalcontacts_bulk_notes_removeCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Bulk remove notes",
	Long:  "Bulk remove notes",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Bulkidsrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/bulk/notes/remove"

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
