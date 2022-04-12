package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpenintegrationupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpenintegrationupdaterequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Openintegrationupdaterequest
type Openintegrationupdaterequest struct { 
    


    // Name - The name of the Open messaging integration.
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting
    MessagingSetting Messagingsettingreference `json:"messagingSetting"`


    // OutboundNotificationWebhookUrl - The outbound notification webhook URL for the Open messaging integration.
    OutboundNotificationWebhookUrl string `json:"outboundNotificationWebhookUrl"`


    // OutboundNotificationWebhookSignatureSecretToken - The outbound notification webhook signature secret token.
    OutboundNotificationWebhookSignatureSecretToken string `json:"outboundNotificationWebhookSignatureSecretToken"`


    // WebhookHeaders - The user specified headers for the Open messaging integration.
    WebhookHeaders map[string]string `json:"webhookHeaders"`


    

}

// String returns a JSON representation of the model
func (o *Openintegrationupdaterequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.WebhookHeaders = map[string]string{"": ""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openintegrationupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Openintegrationupdaterequest

    if OpenintegrationupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    OpenintegrationupdaterequestMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingreference `json:"messagingSetting"`
        
        OutboundNotificationWebhookUrl string `json:"outboundNotificationWebhookUrl"`
        
        OutboundNotificationWebhookSignatureSecretToken string `json:"outboundNotificationWebhookSignatureSecretToken"`
        
        WebhookHeaders map[string]string `json:"webhookHeaders"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        WebhookHeaders: map[string]string{"": ""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

