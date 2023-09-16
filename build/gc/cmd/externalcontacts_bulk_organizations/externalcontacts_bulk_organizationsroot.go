package externalcontacts_bulk_organizations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_bulk_organizations_update"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_bulk_organizations_add"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_bulk_organizations_remove"
)

func init() {
	externalcontacts_bulk_organizationsCmd.AddCommand(externalcontacts_bulk_organizations_update.Cmdexternalcontacts_bulk_organizations_update())
	externalcontacts_bulk_organizationsCmd.AddCommand(externalcontacts_bulk_organizations_add.Cmdexternalcontacts_bulk_organizations_add())
	externalcontacts_bulk_organizationsCmd.AddCommand(externalcontacts_bulk_organizations_remove.Cmdexternalcontacts_bulk_organizations_remove())
	externalcontacts_bulk_organizationsCmd.Short = utils.GenerateCustomDescription(externalcontacts_bulk_organizationsCmd.Short, externalcontacts_bulk_organizations_update.Description, externalcontacts_bulk_organizations_add.Description, externalcontacts_bulk_organizations_remove.Description, )
	externalcontacts_bulk_organizationsCmd.Long = externalcontacts_bulk_organizationsCmd.Short
}
