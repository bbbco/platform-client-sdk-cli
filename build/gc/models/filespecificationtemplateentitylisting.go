package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FilespecificationtemplateentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FilespecificationtemplateentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Filespecificationtemplateentitylisting
type Filespecificationtemplateentitylisting struct { 
    // Entities
    Entities []Filespecificationtemplate `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Filespecificationtemplateentitylisting) String() string {
     o.Entities = []Filespecificationtemplate{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Filespecificationtemplateentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Filespecificationtemplateentitylisting

    if FilespecificationtemplateentitylistingMarshalled {
        return []byte("{}"), nil
    }
    FilespecificationtemplateentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Filespecificationtemplate `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Filespecificationtemplate{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

