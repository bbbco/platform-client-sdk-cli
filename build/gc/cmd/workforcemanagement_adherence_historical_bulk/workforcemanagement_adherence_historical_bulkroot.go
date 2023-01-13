package workforcemanagement_adherence_historical_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_adherence_historical_bulk_jobs"
)

func init() {
	workforcemanagement_adherence_historical_bulkCmd.AddCommand(workforcemanagement_adherence_historical_bulk_jobs.Cmdworkforcemanagement_adherence_historical_bulk_jobs())
	workforcemanagement_adherence_historical_bulkCmd.Short = utils.GenerateCustomDescription(workforcemanagement_adherence_historical_bulkCmd.Short, workforcemanagement_adherence_historical_bulk_jobs.Description, )
	workforcemanagement_adherence_historical_bulkCmd.Long = workforcemanagement_adherence_historical_bulkCmd.Short
}
