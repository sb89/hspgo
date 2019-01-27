package hsp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetServiceDetailsReturnsResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `
		{
			"serviceAttributesDetails": {
				"date_of_service": "2016-07-29",
				"toc_code": "GX",
				"rid": "123456789",
				"locations": [
					{
						"location": "BTN",
						"gbtt_ptd": "0712",
						"gbtt_pta": "",
						"actual_td": "0711",
						"actual_ta": "",
						"late_canc_reason": ""
					},
					{
						"location": "GTW",
						"gbtt_ptd": "0749",
						"gbtt_pta": "0747",
						"actual_td": "0751",
						"actual_ta": "0744",
						"late_canc_reason": ""
					},
					{
						"location": "VIC",
						"gbtt_ptd": "",
						"gbtt_pta": "0823",
						"actual_td": "",
						"actual_ta": "",
						"late_canc_reason": "CANC"
					}
					
				]
			}
		}`)
	}))
	defer server.Close()

	expected := ServiceDetailsResponseData{
		ServiceAttributesDetails{
			DateOfService: "2016-07-29",
			TocCode:       "GX",
			Rid:           "123456789",
			Locations: []Location{
				Location{
					Location:       "BTN",
					GbttPtd:        "0712",
					GbttPta:        "",
					ActualTD:       "0711",
					ActualTA:       "",
					LateCancReason: "",
				},
				Location{
					Location:       "GTW",
					GbttPtd:        "0749",
					GbttPta:        "0747",
					ActualTD:       "0751",
					ActualTA:       "0744",
					LateCancReason: "",
				},
				Location{
					Location:       "VIC",
					GbttPtd:        "",
					GbttPta:        "0823",
					ActualTD:       "",
					ActualTA:       "",
					LateCancReason: "CANC",
				},
			},
		},
	}

	c := NewClient("email", "password", BaseURL(server.URL))
	resp, _ := c.GetServiceDetails("123456789")

	if !reflect.DeepEqual(expected, *resp) {
		t.Errorf("Expected response to be %v but instead got %v!", expected, *resp)
	}
}
