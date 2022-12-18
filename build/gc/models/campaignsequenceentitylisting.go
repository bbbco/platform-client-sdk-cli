package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignsequenceentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignsequenceentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Campaignsequenceentitylisting
type Campaignsequenceentitylisting struct { 
    // Entities
    Entities []Campaignsequence `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Campaignsequenceentitylisting) String() string {
     o.Entities = []Campaignsequence{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignsequenceentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Campaignsequenceentitylisting

    if CampaignsequenceentitylistingMarshalled {
        return []byte("{}"), nil
    }
    CampaignsequenceentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Campaignsequence `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        PreviousUri string `json:"previousUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Campaignsequence{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

