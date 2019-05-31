package location

// Location the model for input
type Location struct {
  Longitude float32
  Latitude  float32
  PostCode  string
  Street    string
}

// GoogleResponse the response from google geocoder
type GoogleResponse struct {
  Results []struct {
    AddressComponents []struct {
      LongName  string `json:"long_name"`
      ShortName string `json:"short_name"`
      Types     []string
    } `json:"address_components"`
    Geometry struct {
      Bounds struct {
        NorthEast struct {
          Latitude  float32 `json:"lat"`
          Longitude float32 `json:"lng"`
        } `json:"northeast"`
        SouthWest struct {
          Latitude  float32 `json:"lat"`
          Longitude float32 `json:"lng"`
        } `json:"southwest"`
      } `json:"bounds"`
      Location struct {
        Latitude  float32 `json:"lat"`
        Longitude float32 `json:"lng"`
      } `json:"location"`
      LocationType string `json:"location_type"`
      ViewPort     struct {
        NorthEast struct {
          Latitude  float32 `json:"lat"`
          Longitude float32 `json:"lng"`
        } `json:"northeast"`
        SouthWest struct {
          Latitude  float32 `json:"lat"`
          Longitude float32 `json:"lng"`
        } `json:"southwest"`
      } `json:"viewport"`
    } `json:"geometry"`
    FormattedAddress string   `json:"formatted_address"`
    PlaceID          string   `json:"place_id"`
    Types            []string `json:"types"`
  }
  Status string `json:"status"`
}
