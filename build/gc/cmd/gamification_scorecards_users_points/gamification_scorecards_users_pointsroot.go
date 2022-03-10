package gamification_scorecards_users_points

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_users_points_alltime"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_users_points_average"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_users_points_trends"
)

func init() {
	gamification_scorecards_users_pointsCmd.AddCommand(gamification_scorecards_users_points_alltime.Cmdgamification_scorecards_users_points_alltime())
	gamification_scorecards_users_pointsCmd.AddCommand(gamification_scorecards_users_points_average.Cmdgamification_scorecards_users_points_average())
	gamification_scorecards_users_pointsCmd.AddCommand(gamification_scorecards_users_points_trends.Cmdgamification_scorecards_users_points_trends())
	gamification_scorecards_users_pointsCmd.Short = utils.GenerateCustomDescription(gamification_scorecards_users_pointsCmd.Short, gamification_scorecards_users_points_alltime.Description, gamification_scorecards_users_points_average.Description, gamification_scorecards_users_points_trends.Description, )
	gamification_scorecards_users_pointsCmd.Long = gamification_scorecards_users_pointsCmd.Short
}
