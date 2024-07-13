package conversations_calls_participants

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_participants_replace"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_participants_monitor"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_participants_wrapupcodes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_participants_attributes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_participants_wrapup"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_participants_communications"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_participants_coach"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_participants_consult"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_participants_barge"
)

func init() {
	conversations_calls_participantsCmd.AddCommand(conversations_calls_participants_replace.Cmdconversations_calls_participants_replace())
	conversations_calls_participantsCmd.AddCommand(conversations_calls_participants_monitor.Cmdconversations_calls_participants_monitor())
	conversations_calls_participantsCmd.AddCommand(conversations_calls_participants_wrapupcodes.Cmdconversations_calls_participants_wrapupcodes())
	conversations_calls_participantsCmd.AddCommand(conversations_calls_participants_attributes.Cmdconversations_calls_participants_attributes())
	conversations_calls_participantsCmd.AddCommand(conversations_calls_participants_wrapup.Cmdconversations_calls_participants_wrapup())
	conversations_calls_participantsCmd.AddCommand(conversations_calls_participants_communications.Cmdconversations_calls_participants_communications())
	conversations_calls_participantsCmd.AddCommand(conversations_calls_participants_coach.Cmdconversations_calls_participants_coach())
	conversations_calls_participantsCmd.AddCommand(conversations_calls_participants_consult.Cmdconversations_calls_participants_consult())
	conversations_calls_participantsCmd.AddCommand(conversations_calls_participants_barge.Cmdconversations_calls_participants_barge())
	conversations_calls_participantsCmd.Short = utils.GenerateCustomDescription(conversations_calls_participantsCmd.Short, conversations_calls_participants_replace.Description, conversations_calls_participants_monitor.Description, conversations_calls_participants_wrapupcodes.Description, conversations_calls_participants_attributes.Description, conversations_calls_participants_wrapup.Description, conversations_calls_participants_communications.Description, conversations_calls_participants_coach.Description, conversations_calls_participants_consult.Description, conversations_calls_participants_barge.Description, )
	conversations_calls_participantsCmd.Long = conversations_calls_participantsCmd.Short
}
