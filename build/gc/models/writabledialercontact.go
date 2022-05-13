package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WritabledialercontactMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WritabledialercontactDud struct { 
    


    


    


    LatestSmsEvaluations map[string]Messageevaluation `json:"latestSmsEvaluations"`


    


    

}

// Writabledialercontact
type Writabledialercontact struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // ContactListId - The identifier of the contact list containing this contact.
    ContactListId string `json:"contactListId"`


    // Data - An ordered map of the contact's columns and corresponding values.
    Data map[string]interface{} `json:"data"`


    


    // Callable - Indicates whether or not the contact can be called.
    Callable bool `json:"callable"`


    // PhoneNumberStatus - A map of phone number columns to PhoneNumberStatuses, which indicate if the phone number is callable or not.
    PhoneNumberStatus map[string]Phonenumberstatus `json:"phoneNumberStatus"`

}

// String returns a JSON representation of the model
func (o *Writabledialercontact) String() string {
    
    
     o.Data = map[string]interface{}{"": Interface{}} 
    
     o.PhoneNumberStatus = map[string]Phonenumberstatus{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Writabledialercontact) MarshalJSON() ([]byte, error) {
    type Alias Writabledialercontact

    if WritabledialercontactMarshalled {
        return []byte("{}"), nil
    }
    WritabledialercontactMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        ContactListId string `json:"contactListId"`
        
        Data map[string]interface{} `json:"data"`
        
        Callable bool `json:"callable"`
        
        PhoneNumberStatus map[string]Phonenumberstatus `json:"phoneNumberStatus"`
        *Alias
    }{

        


        


        
        Data: map[string]interface{}{"": Interface{}},
        


        


        


        
        PhoneNumberStatus: map[string]Phonenumberstatus{"": {}},
        

        Alias: (*Alias)(u),
    })
}

