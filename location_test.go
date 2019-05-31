package location_test

import (
  "errors"
  "fmt"
  "github.com/joho/godotenv"
  "github.com/keloran/go-location"
  "github.com/stretchr/testify/assert"
  "os"
  "testing"
)

func TestLocation_GetLocation(t *testing.T) {
  if os.Getenv("GOOGLE_API_KEY") == "" {
    err := godotenv.Load()
    if err != nil {
      fmt.Println(fmt.Sprintf("godotenv err: %v", err))
    }
  }

  testPostCode1 := "BB11 1PX"
  testPostCode2 := "BB11 ZPX"

  tests := []struct {
  	request location.Location
  	expect  location.Location
  	err     error
  }{
  	{
  		request: location.Location{
  			PostCode: testPostCode1,
  		},
  		expect: location.Location{
  			PostCode:  testPostCode1,
  			Longitude: -2.2458968,
  			Latitude:  53.79071,
  			Street:    "Brick St",
  		},
  		err: nil,
  	},
  	{
  		request: location.Location{
  			PostCode: testPostCode2,
  		},
  		expect: location.Location{},
  		err:    errors.New("invalid postcode"),
  	},
  }

  for _, test := range tests {
  	resp, err := test.request.GetLocation()
  	if err != nil {
  		assert.Equal(t, test.err, err)
  	}
  	assert.Equal(t, test.expect, resp)
  }
}
