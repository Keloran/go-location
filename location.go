package location

import (
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "strings"
)

// GetLocation gets the longitude, latitude and address
func (l Location) GetLocation() (r Location, err error) {
  client := &http.Client{}
  reqAdd := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?key=%s&address=%s", os.Getenv("GOOGLE_API_KEY"), strings.Replace(l.PostCode, " ", "%20", -1))
  req, err := http.NewRequest("GET", reqAdd, nil)
  if err != nil {
    fmt.Printf("Request: %v", err)
    return r, err
  }

  resp, err := client.Do(req)
  if err != nil {
    fmt.Printf("Do: %v", err)
    return r, err
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Printf("Body: %v", err)
    return r, err
  }

  j := GoogleResponse{}
  err = json.Unmarshal(body, &j)
  if err != nil {
    fmt.Printf("Unmarshal: %v", err)
    return r, err
  }

  if len(j.Results) >= 1 {
    r := Location{
      PostCode:  l.PostCode,
      Longitude: j.Results[0].Geometry.Location.Longitude,
      Latitude:  j.Results[0].Geometry.Location.Latitude,
      Street:    j.Results[0].AddressComponents[1].ShortName,
    }
    return r, nil
  }

  return r, errors.New("invalid postcode")
}
