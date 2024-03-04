package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialogflowcxenvironmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialogflowcxenvironmentDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Dialogflowcxenvironment
type Dialogflowcxenvironment struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Dialogflowcxenvironment) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialogflowcxenvironment) MarshalJSON() ([]byte, error) {
    type Alias Dialogflowcxenvironment

    if DialogflowcxenvironmentMarshalled {
        return []byte("{}"), nil
    }
    DialogflowcxenvironmentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

