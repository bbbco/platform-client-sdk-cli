package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SttengineentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SttengineentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Sttengineentitylisting
type Sttengineentitylisting struct { 
    // Entities
    Entities []Sttengineentity `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Sttengineentitylisting) String() string {
     o.Entities = []Sttengineentity{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sttengineentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Sttengineentitylisting

    if SttengineentitylistingMarshalled {
        return []byte("{}"), nil
    }
    SttengineentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Sttengineentity `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Sttengineentity{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

