package users_queues

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
	users_queuesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_queues"),
		Short: utils.FormatUsageDescription("Manages Genesys Cloud users_queues", "SWAGGER_OVERRIDE_queue users", "", "SWAGGER_OVERRIDE_queue users", "SWAGGER_OVERRIDE_queue users", "", ),
		Long:  utils.FormatUsageDescription(`Manages Genesys Cloud users_queues`, "SWAGGER_OVERRIDE_queue users", "", "SWAGGER_OVERRIDE_queue users", "SWAGGER_OVERRIDE_queue users", "", ),
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_queuesCmd)
}

func Cmdusers_queues() *cobra.Command { 
	activateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", activateCmd.UsageTemplate(), "PATCH", "/api/v2/routing/queues/{queueId}/users", utils.FormatPermissions([]string{ "routing:queue:edit", "routing:queueMember:manage",  })))
	utils.AddFileFlagIfUpsert(activateCmd.Flags(), "PATCH", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Queue Members&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;$ref&quot; : &quot;#/definitions/QueueMember&quot;
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(activateCmd.Flags(), "PATCH", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/QueueMemberEntityListing&quot;
  }
}`)
	users_queuesCmd.AddCommand(activateCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/routing/queues/{queueId}/users/{memberId}", utils.FormatPermissions([]string{ "routing:queue:edit", "routing:queueMember:manage",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Operation was successful.&quot;
}`)
	users_queuesCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "int", "pageSize", "25", "Page size [max 100]")
	utils.AddFlag(getCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(getCmd.Flags(), "string", "sortBy", "name", "Sort by")
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand. Valid values: routingStatus, presence, conversationSummary, outOfOffice, geolocation, station, authorization, lasttokenissued, authorization.unusedRoles, team, profileSkills, certifications, locations, groups, skills, languages, languagePreference, employerInfo, biography")
	utils.AddFlag(getCmd.Flags(), "bool", "joined", "", "Filter by joined status")
	utils.AddFlag(getCmd.Flags(), "string", "name", "", "Filter by queue member name")
	utils.AddFlag(getCmd.Flags(), "[]string", "profileSkills", "", "Filter by profile skill")
	utils.AddFlag(getCmd.Flags(), "[]string", "skills", "", "Filter by skill")
	utils.AddFlag(getCmd.Flags(), "[]string", "languages", "", "Filter by language")
	utils.AddFlag(getCmd.Flags(), "[]string", "routingStatus", "", "Filter by routing status")
	utils.AddFlag(getCmd.Flags(), "[]string", "presence", "", "Filter by presence")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/routing/queues/{queueId}/users", utils.FormatPermissions([]string{ "routing:queue:view", "routing:queueMember:manage",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/QueueMemberEntityListing&quot;
  }
}`)
	users_queuesCmd.AddCommand(getCmd)
	
	utils.AddFlag(moveCmd.Flags(), "bool", "delete", "false", "True to delete queue members")
	moveCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", moveCmd.UsageTemplate(), "POST", "/api/v2/routing/queues/{queueId}/users", utils.FormatPermissions([]string{ "routing:queue:edit", "routing:queueMember:manage",  })))
	utils.AddFileFlagIfUpsert(moveCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Queue Members&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;$ref&quot; : &quot;#/definitions/WritableEntity&quot;
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(moveCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;The request could not be understood by the server due to malformed syntax.&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ErrorBody&quot;
  },
  &quot;x-inin-error-codes&quot; : {
    &quot;bad.request&quot; : &quot;Invalid request data.  Make sure you submit a valid number of queue members.&quot;,
    &quot;response.entity.too.large&quot; : &quot;The response is over the size limit. Reduce pageSize or expand list to reduce response size if applicable&quot;,
    &quot;invalid.date&quot; : &quot;Dates must be specified as ISO-8601 strings. For example: yyyy-MM-ddTHH:mm:ss.SSSZ&quot;,
    &quot;queue.size.limit&quot; : &quot;Adding all requested members would exceed queue member limit.&quot;,
    &quot;invalid.value&quot; : &quot;Value [%s] is not valid for field type [%s]. Allowable values are: %s&quot;
  }
}`)
	users_queuesCmd.AddCommand(moveCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/routing/queues/{queueId}/users/{memberId}", utils.FormatPermissions([]string{ "routing:queue:edit", "routing:queueMember:manage",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Queue Member&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/QueueMember&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/QueueMember&quot;
  }
}`)
	users_queuesCmd.AddCommand(updateCmd)
	
	return users_queuesCmd
}

var activateCmd = &cobra.Command{
	Use:   "activate [queueId]",
	Short: "DEPRECATED: use PATCH /routing/queues/{queueId}/members.  Join or unjoin a set of users for a queue.",
	Long:  `DEPRECATED: use PATCH /routing/queues/{queueId}/members.  Join or unjoin a set of users for a queue.`,
	Args:  utils.DetermineArgs([]string{ "queueId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/queues/{queueId}/users"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PATCH", urlString, cmd.Flags())
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
var deleteCmd = &cobra.Command{
	Use:   "delete [queueId] [memberId]",
	Short: "DEPRECATED: use DELETE /routing/queues/{queueId}/members/{memberId}.  Delete queue member.",
	Long:  `DEPRECATED: use DELETE /routing/queues/{queueId}/members/{memberId}.  Delete queue member.`,
	Args:  utils.DetermineArgs([]string{ "queueId", "memberId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/queues/{queueId}/users/{memberId}"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)
		memberId, args := args[0], args[1:]
		path = strings.Replace(path, "{memberId}", fmt.Sprintf("%v", memberId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", urlString, cmd.Flags())
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
	Use:   "get [queueId]",
	Short: "DEPRECATED: use GET /routing/queues/{queueId}/members.  Get the members of this queue.",
	Long:  `DEPRECATED: use GET /routing/queues/{queueId}/members.  Get the members of this queue.`,
	Args:  utils.DetermineArgs([]string{ "queueId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/queues/{queueId}/users"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		joined := utils.GetFlag(cmd.Flags(), "bool", "joined")
		if joined != "" {
			queryParams["joined"] = joined
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		profileSkills := utils.GetFlag(cmd.Flags(), "[]string", "profileSkills")
		if profileSkills != "" {
			queryParams["profileSkills"] = profileSkills
		}
		skills := utils.GetFlag(cmd.Flags(), "[]string", "skills")
		if skills != "" {
			queryParams["skills"] = skills
		}
		languages := utils.GetFlag(cmd.Flags(), "[]string", "languages")
		if languages != "" {
			queryParams["languages"] = languages
		}
		routingStatus := utils.GetFlag(cmd.Flags(), "[]string", "routingStatus")
		if routingStatus != "" {
			queryParams["routingStatus"] = routingStatus
		}
		presence := utils.GetFlag(cmd.Flags(), "[]string", "presence")
		if presence != "" {
			queryParams["presence"] = presence
		}
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
var moveCmd = &cobra.Command{
	Use:   "move [queueId]",
	Short: "DEPRECATED: use POST /routing/queues/{queueId}/members.  Bulk add or delete up to 100 queue members.",
	Long:  `DEPRECATED: use POST /routing/queues/{queueId}/members.  Bulk add or delete up to 100 queue members.`,
	Args:  utils.DetermineArgs([]string{ "queueId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/queues/{queueId}/users"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)

		delete := utils.GetFlag(cmd.Flags(), "bool", "delete")
		if delete != "" {
			queryParams["delete"] = delete
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", urlString, cmd.Flags())
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
var updateCmd = &cobra.Command{
	Use:   "update [queueId] [memberId]",
	Short: "DEPRECATED: use PATCH /routing/queues/{queueId}/members/{memberId}.  Update the ring number OR joined status for a User in a Queue.",
	Long:  `DEPRECATED: use PATCH /routing/queues/{queueId}/members/{memberId}.  Update the ring number OR joined status for a User in a Queue.`,
	Args:  utils.DetermineArgs([]string{ "queueId", "memberId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/queues/{queueId}/users/{memberId}"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)
		memberId, args := args[0], args[1:]
		path = strings.Replace(path, "{memberId}", fmt.Sprintf("%v", memberId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PATCH", urlString, cmd.Flags())
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
