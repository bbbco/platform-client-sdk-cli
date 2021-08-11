package webdeployments_configurations_versions

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("webdeployments_configurations_versions", "SWAGGER_OVERRIDE_/api/v2/webdeployments/configurations/{configurationId}/versions", "SWAGGER_OVERRIDE_/api/v2/webdeployments/configurations/{configurationId}/versions", )
	webdeployments_configurations_versionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webdeployments_configurations_versions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webdeployments_configurations_versionsCmd)
}

func Cmdwebdeployments_configurations_versions() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/webdeployments/configurations/{configurationId}/versions/{versionId}", utils.FormatPermissions([]string{ "webDeployments:configuration:view",  }), utils.GenerateDevCentreLink("GET", "Web Deployments", "/api/v2/webdeployments/configurations/{configurationId}/versions/{versionId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebDeploymentConfigurationVersion&quot;
  }
}`)
	webdeployments_configurations_versionsCmd.AddCommand(getCmd)
	
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/webdeployments/configurations/{configurationId}/versions", utils.FormatPermissions([]string{ "webDeployments:configuration:view",  }), utils.GenerateDevCentreLink("GET", "Web Deployments", "/api/v2/webdeployments/configurations/{configurationId}/versions")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebDeploymentConfigurationVersionEntityListing&quot;
  }
}`)
	webdeployments_configurations_versionsCmd.AddCommand(listCmd)
	
	return webdeployments_configurations_versionsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [configurationId] [versionId]",
	Short: "Get a configuration version",
	Long:  "Get a configuration version",
	Args:  utils.DetermineArgs([]string{ "configurationId", "versionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/webdeployments/configurations/{configurationId}/versions/{versionId}"
		configurationId, args := args[0], args[1:]
		path = strings.Replace(path, "{configurationId}", fmt.Sprintf("%v", configurationId), -1)
		versionId, args := args[0], args[1:]
		path = strings.Replace(path, "{versionId}", fmt.Sprintf("%v", versionId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
var listCmd = &cobra.Command{
	Use:   "list [configurationId]",
	Short: "Get the versions of a configuration",
	Long:  "Get the versions of a configuration",
	Args:  utils.DetermineArgs([]string{ "configurationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/webdeployments/configurations/{configurationId}/versions"
		configurationId, args := args[0], args[1:]
		path = strings.Replace(path, "{configurationId}", fmt.Sprintf("%v", configurationId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
