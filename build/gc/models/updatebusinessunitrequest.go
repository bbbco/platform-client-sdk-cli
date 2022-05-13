package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatebusinessunitrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatebusinessunitrequestDud struct { 
    


    


    

}

// Updatebusinessunitrequest
type Updatebusinessunitrequest struct { 
    // Name - The name of the business unit
    Name string `json:"name"`


    // DivisionId - The ID of the division to which the business unit should be moved
    DivisionId string `json:"divisionId"`


    // Settings - Configuration for the business unit
    Settings Updatebusinessunitsettings `json:"settings"`

}

// String returns a JSON representation of the model
func (o *Updatebusinessunitrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatebusinessunitrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatebusinessunitrequest

    if UpdatebusinessunitrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatebusinessunitrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        DivisionId string `json:"divisionId"`
        
        Settings Updatebusinessunitsettings `json:"settings"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

