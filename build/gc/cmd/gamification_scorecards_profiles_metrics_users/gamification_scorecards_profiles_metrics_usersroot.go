package gamification_scorecards_profiles_metrics_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_profiles_metrics_users_values"
)

func init() {
	gamification_scorecards_profiles_metrics_usersCmd.AddCommand(gamification_scorecards_profiles_metrics_users_values.Cmdgamification_scorecards_profiles_metrics_users_values())
	gamification_scorecards_profiles_metrics_usersCmd.Short = utils.GenerateCustomDescription(gamification_scorecards_profiles_metrics_usersCmd.Short, gamification_scorecards_profiles_metrics_users_values.Description, )
	gamification_scorecards_profiles_metrics_usersCmd.Long = gamification_scorecards_profiles_metrics_usersCmd.Short
}
