package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InteractionstatsrulecontainerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InteractionstatsrulecontainerDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Interactionstatsrulecontainer
type Interactionstatsrulecontainer struct { 
    // Entities
    Entities []Interactionstatsrule `json:"entities"`


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
func (o *Interactionstatsrulecontainer) String() string {
     o.Entities = []Interactionstatsrule{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Interactionstatsrulecontainer) MarshalJSON() ([]byte, error) {
    type Alias Interactionstatsrulecontainer

    if InteractionstatsrulecontainerMarshalled {
        return []byte("{}"), nil
    }
    InteractionstatsrulecontainerMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Interactionstatsrule `json:"entities"`
        
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

        
        Entities: []Interactionstatsrule{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

