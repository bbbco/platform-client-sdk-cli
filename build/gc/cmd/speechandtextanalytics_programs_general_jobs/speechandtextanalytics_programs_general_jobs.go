package speechandtextanalytics_programs_general_jobs

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
	Description = utils.FormatUsageDescription("speechandtextanalytics_programs_general_jobs", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/programs/general/jobs", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/programs/general/jobs", )
	speechandtextanalytics_programs_general_jobsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_programs_general_jobs"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_programs_general_jobsCmd)
}

func Cmdspeechandtextanalytics_programs_general_jobs() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/speechandtextanalytics/programs/general/jobs", utils.FormatPermissions([]string{ "speechAndTextAnalytics:program:add", "speechAndTextAnalytics:program:edit", "speechAndTextAnalytics:topic:add", "speechAndTextAnalytics:topic:edit",  }), utils.GenerateDevCentreLink("POST", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/programs/general/jobs")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "The general programs job to create",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/GeneralProgramJobRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/GeneralProgramJob"
  }
}`)
	speechandtextanalytics_programs_general_jobsCmd.AddCommand(createCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/speechandtextanalytics/programs/general/jobs/{jobId}", utils.FormatPermissions([]string{ "speechAndTextAnalytics:program:add", "speechAndTextAnalytics:program:edit", "speechAndTextAnalytics:topic:add", "speechAndTextAnalytics:topic:edit",  }), utils.GenerateDevCentreLink("GET", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/programs/general/jobs/{jobId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/GeneralProgramJob"
  }
}`)
	speechandtextanalytics_programs_general_jobsCmd.AddCommand(getCmd)
	
	return speechandtextanalytics_programs_general_jobsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new Speech and Text Analytics general program job",
	Long:  "Create new Speech and Text Analytics general program job",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Generalprogramjobrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/programs/general/jobs"

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
var getCmd = &cobra.Command{
	Use:   "get [jobId]",
	Short: "Get a Speech and Text Analytics general program job by id",
	Long:  "Get a Speech and Text Analytics general program job by id",
	Args:  utils.DetermineArgs([]string{ "jobId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/programs/general/jobs/{jobId}"
		jobId, args := args[0], args[1:]
		path = strings.Replace(path, "{jobId}", fmt.Sprintf("%v", jobId), -1)

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
