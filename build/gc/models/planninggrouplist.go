package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PlanninggrouplistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PlanninggrouplistDud struct { 
    

}

// Planninggrouplist - List of planning groups
type Planninggrouplist struct { 
    // Entities
    Entities []Planninggroup `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Planninggrouplist) String() string {
     o.Entities = []Planninggroup{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Planninggrouplist) MarshalJSON() ([]byte, error) {
    type Alias Planninggrouplist

    if PlanninggrouplistMarshalled {
        return []byte("{}"), nil
    }
    PlanninggrouplistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Planninggroup `json:"entities"`
        *Alias
    }{

        
        Entities: []Planninggroup{{}},
        

        Alias: (*Alias)(u),
    })
}

