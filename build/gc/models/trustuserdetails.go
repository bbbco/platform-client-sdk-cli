package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrustuserdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrustuserdetailsDud struct { 
    DateCreated time.Time `json:"dateCreated"`


    CreatedBy Orguser `json:"createdBy"`

}

// Trustuserdetails
type Trustuserdetails struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Trustuserdetails) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trustuserdetails) MarshalJSON() ([]byte, error) {
    type Alias Trustuserdetails

    if TrustuserdetailsMarshalled {
        return []byte("{}"), nil
    }
    TrustuserdetailsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

