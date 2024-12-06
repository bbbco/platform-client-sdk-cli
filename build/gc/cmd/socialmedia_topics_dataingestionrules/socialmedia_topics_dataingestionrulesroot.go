package socialmedia_topics_dataingestionrules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_topics_dataingestionrules_facebook"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_topics_dataingestionrules_open"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_topics_dataingestionrules_twitter"
)

func init() {
	socialmedia_topics_dataingestionrulesCmd.AddCommand(socialmedia_topics_dataingestionrules_facebook.Cmdsocialmedia_topics_dataingestionrules_facebook())
	socialmedia_topics_dataingestionrulesCmd.AddCommand(socialmedia_topics_dataingestionrules_open.Cmdsocialmedia_topics_dataingestionrules_open())
	socialmedia_topics_dataingestionrulesCmd.AddCommand(socialmedia_topics_dataingestionrules_twitter.Cmdsocialmedia_topics_dataingestionrules_twitter())
	socialmedia_topics_dataingestionrulesCmd.Short = utils.GenerateCustomDescription(socialmedia_topics_dataingestionrulesCmd.Short, socialmedia_topics_dataingestionrules_facebook.Description, socialmedia_topics_dataingestionrules_open.Description, socialmedia_topics_dataingestionrules_twitter.Description, )
	socialmedia_topics_dataingestionrulesCmd.Long = socialmedia_topics_dataingestionrulesCmd.Short
}
