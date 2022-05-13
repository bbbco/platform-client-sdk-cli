package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainentityDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Domainentity
type Domainentity struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Domainentity) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainentity) MarshalJSON() ([]byte, error) {
    type Alias Domainentity

    if DomainentityMarshalled {
        return []byte("{}"), nil
    }
    DomainentityMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

