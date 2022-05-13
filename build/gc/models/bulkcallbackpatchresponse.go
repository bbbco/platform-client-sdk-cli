package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkcallbackpatchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkcallbackpatchresponseDud struct { 
    


    


    

}

// Bulkcallbackpatchresponse
type Bulkcallbackpatchresponse struct { 
    // Results - A list of the results from the bulk operation.
    Results []Bulkresult `json:"results"`


    // ErrorCount - The number of errors from the bulk operation.
    ErrorCount int `json:"errorCount"`


    // ErrorIndexes - An index of where the errors are in the listing.
    ErrorIndexes []int `json:"errorIndexes"`

}

// String returns a JSON representation of the model
func (o *Bulkcallbackpatchresponse) String() string {
     o.Results = []Bulkresult{{}} 
    
     o.ErrorIndexes = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkcallbackpatchresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkcallbackpatchresponse

    if BulkcallbackpatchresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkcallbackpatchresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkresult `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorIndexes []int `json:"errorIndexes"`
        *Alias
    }{

        
        Results: []Bulkresult{{}},
        


        


        
        ErrorIndexes: []int{0},
        

        Alias: (*Alias)(u),
    })
}

