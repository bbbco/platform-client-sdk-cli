package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScheduleuploadprocessingresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScheduleuploadprocessingresponseDud struct { 
    


    


    

}

// Scheduleuploadprocessingresponse
type Scheduleuploadprocessingresponse struct { 
    // Status - The status of the operation
    Status string `json:"status"`


    // OperationId - The ID for the operation
    OperationId string `json:"operationId"`


    // Result - The result of the operation.  Null unless status == Complete
    Result Buschedulemetadata `json:"result"`

}

// String returns a JSON representation of the model
func (o *Scheduleuploadprocessingresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scheduleuploadprocessingresponse) MarshalJSON() ([]byte, error) {
    type Alias Scheduleuploadprocessingresponse

    if ScheduleuploadprocessingresponseMarshalled {
        return []byte("{}"), nil
    }
    ScheduleuploadprocessingresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        OperationId string `json:"operationId"`
        
        Result Buschedulemetadata `json:"result"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

