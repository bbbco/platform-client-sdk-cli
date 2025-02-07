package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WeekshifttraderesponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WeekshifttraderesponseDud struct { 
    


    

}

// Weekshifttraderesponse
type Weekshifttraderesponse struct { 
    // Trade - The shift trade details
    Trade Shifttraderesponse `json:"trade"`


    // MatchReview - A preview of what the schedule would look like if the shift trade is approved plus any violations
    MatchReview Shifttradematchreviewresponse `json:"matchReview"`

}

// String returns a JSON representation of the model
func (o *Weekshifttraderesponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Weekshifttraderesponse) MarshalJSON() ([]byte, error) {
    type Alias Weekshifttraderesponse

    if WeekshifttraderesponseMarshalled {
        return []byte("{}"), nil
    }
    WeekshifttraderesponseMarshalled = true

    return json.Marshal(&struct {
        
        Trade Shifttraderesponse `json:"trade"`
        
        MatchReview Shifttradematchreviewresponse `json:"matchReview"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

