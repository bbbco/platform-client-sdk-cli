package languageunderstanding_domains_versions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_domains_versions_publish"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_domains_versions_report"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_domains_versions_train"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_domains_versions_detect"
)

func init() {
	languageunderstanding_domains_versionsCmd.AddCommand(languageunderstanding_domains_versions_publish.Cmdlanguageunderstanding_domains_versions_publish())
	languageunderstanding_domains_versionsCmd.AddCommand(languageunderstanding_domains_versions_report.Cmdlanguageunderstanding_domains_versions_report())
	languageunderstanding_domains_versionsCmd.AddCommand(languageunderstanding_domains_versions_train.Cmdlanguageunderstanding_domains_versions_train())
	languageunderstanding_domains_versionsCmd.AddCommand(languageunderstanding_domains_versions_detect.Cmdlanguageunderstanding_domains_versions_detect())
	languageunderstanding_domains_versionsCmd.Short = utils.GenerateCustomDescription(languageunderstanding_domains_versionsCmd.Short, languageunderstanding_domains_versions_publish.Description, languageunderstanding_domains_versions_report.Description, languageunderstanding_domains_versions_train.Description, languageunderstanding_domains_versions_detect.Description, )
	languageunderstanding_domains_versionsCmd.Long = languageunderstanding_domains_versionsCmd.Short
}
