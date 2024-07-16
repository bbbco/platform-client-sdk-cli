package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhonebaseentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhonebaseentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Phonebaseentitylisting
type Phonebaseentitylisting struct { 
    // Entities
    Entities []Phonebase `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // TotalNumberOfEntities - The total organization-wide number of entities.
    TotalNumberOfEntities int `json:"totalNumberOfEntities"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Phonebaseentitylisting) String() string {
     o.Entities = []Phonebase{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phonebaseentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Phonebaseentitylisting

    if PhonebaseentitylistingMarshalled {
        return []byte("{}"), nil
    }
    PhonebaseentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Phonebase `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        TotalNumberOfEntities int `json:"totalNumberOfEntities"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Phonebase{{}},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

