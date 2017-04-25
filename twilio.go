package main
 
import (
  "net/http"
  "net/url"
  "fmt"
  "strings"
  "io/ioutil"
  "encoding/json"
)
 
func sendText(str string) {
  // Set initial variables
  accountSid := config.Twilio.AccountSid
  authToken := config.Twilio.AuthToken
  urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
 
  // Build out the data for our message
  v := url.Values{}
  v.Set("To",config.Twilio.ToPhoneNumber)
  v.Set("From",config.Twilio.FromPhoneNumber)
  v.Set("Body",str)
  rb := *strings.NewReader(v.Encode())
 
  // Create client
  client := &http.Client{}
 
  req, _ := http.NewRequest("POST", urlStr, &rb)
  req.SetBasicAuth(accountSid, authToken)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
 
  // Make request
  resp, _ := client.Do(req)
  if( resp.StatusCode >= 200 && resp.StatusCode < 300 ) {
    var data map[string]interface{}
    bodyBytes, _ := ioutil.ReadAll(resp.Body)
    err := json.Unmarshal(bodyBytes, &data)
    if( err == nil ) {
      fmt.Println(data["sid"])
    }
  } else {
    fmt.Println(resp.Status);
  }
}
