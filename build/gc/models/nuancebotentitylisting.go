package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NuancebotentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NuancebotentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Nuancebotentitylisting - Model for a paged listing of Nuance bots
type Nuancebotentitylisting struct { 
    // Entities
    Entities []Nuancebot `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Nuancebotentitylisting) String() string {
     o.Entities = []Nuancebot{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nuancebotentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Nuancebotentitylisting

    if NuancebotentitylistingMarshalled {
        return []byte("{}"), nil
    }
    NuancebotentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Nuancebot `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Nuancebot{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

