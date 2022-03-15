package conversations_participants

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_participants_codes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_participants_wrapupcodes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_participants_wrapup"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_participants_flaggedreason"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_participants_callbacks"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_participants_attributes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_participants_replace"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_participants_digits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_participants_secureivrsessions"
)

func init() {
	conversations_participantsCmd.AddCommand(conversations_participants_codes.Cmdconversations_participants_codes())
	conversations_participantsCmd.AddCommand(conversations_participants_wrapupcodes.Cmdconversations_participants_wrapupcodes())
	conversations_participantsCmd.AddCommand(conversations_participants_wrapup.Cmdconversations_participants_wrapup())
	conversations_participantsCmd.AddCommand(conversations_participants_flaggedreason.Cmdconversations_participants_flaggedreason())
	conversations_participantsCmd.AddCommand(conversations_participants_callbacks.Cmdconversations_participants_callbacks())
	conversations_participantsCmd.AddCommand(conversations_participants_attributes.Cmdconversations_participants_attributes())
	conversations_participantsCmd.AddCommand(conversations_participants_replace.Cmdconversations_participants_replace())
	conversations_participantsCmd.AddCommand(conversations_participants_digits.Cmdconversations_participants_digits())
	conversations_participantsCmd.AddCommand(conversations_participants_secureivrsessions.Cmdconversations_participants_secureivrsessions())
	conversations_participantsCmd.Short = utils.GenerateCustomDescription(conversations_participantsCmd.Short, conversations_participants_codes.Description, conversations_participants_wrapupcodes.Description, conversations_participants_wrapup.Description, conversations_participants_flaggedreason.Description, conversations_participants_callbacks.Description, conversations_participants_attributes.Description, conversations_participants_replace.Description, conversations_participants_digits.Description, conversations_participants_secureivrsessions.Description, )
	conversations_participantsCmd.Long = conversations_participantsCmd.Short
}
