package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TraininglistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TraininglistingDud struct { 
    


    


    


    

}

// Traininglisting
type Traininglisting struct { 
    // Entities
    Entities []Knowledgetraining `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Traininglisting) String() string {
     o.Entities = []Knowledgetraining{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Traininglisting) MarshalJSON() ([]byte, error) {
    type Alias Traininglisting

    if TraininglistingMarshalled {
        return []byte("{}"), nil
    }
    TraininglistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Knowledgetraining `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Knowledgetraining{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

