package hsp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetServiceMetricsReturnsResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `
		{
			"header": {
				"from_location": "BTN",
				"to_location": "VIC"
			},
			"Services": [
				{
					"serviceAttributesMetrics": {
						"origin_location": "BTN",
						"destination_location": "VIC",
						"gbtt_ptd": "0712",
						"gbtt_pta": "0823",
						"toc_code": "GX",
						"matched_services": "22",
						"rids": [ "201607013361753", "201607043443704"]
					},
					"Metrics": [
						{
							"tolerance_value": "0",
							"num_not_tolerance": "5",
							"num_tolerance": "17",
							"percent_tolerance": "77",
							"global_tolerance": true
						}
					]
				},
				{
					"serviceAttributesMetrics": {
					  "origin_location": "BTN",
					  "destination_location": "VIC",
					  "gbtt_ptd": "0729",
					  "gbtt_pta": "0839",
					  "toc_code": "GX",
					  "matched_services": "22",
					  "rids": [ "201607013361763", "201607043443714"]
					},
					"Metrics": [
					  {
						"tolerance_value": "0",
						"num_not_tolerance": "7",
						"num_tolerance": "15",
						"percent_tolerance": "68",
						"global_tolerance": false
					  }
					]
				  }
			]
		}`)
	}))
	defer server.Close()

	expected := ServiceMetricsResponseData{
		Header: Header{
			FromLocation: "BTN",
			ToLocation:   "VIC",
		},
		Services: []Service{
			Service{
				ServiceAttributesMetrics: ServiceAttributesMetrics{
					OriginLocation:      "BTN",
					DestinationLocation: "VIC",
					GbttPtd:             "0712",
					GbttPta:             "0823",
					TocCode:             "GX",
					MatchedServices:     "22",
					RIDS:                []string{"201607013361753", "201607043443704"},
				},
				Metrics: []Metric{
					Metric{
						ToleranceValue:   "0",
						NumNotTolerance:  "5",
						NumTolerance:     "17",
						PercentTolerance: "77",
						GlobalTolerance:  true,
					},
				},
			},
			Service{
				ServiceAttributesMetrics: ServiceAttributesMetrics{
					OriginLocation:      "BTN",
					DestinationLocation: "VIC",
					GbttPtd:             "0729",
					GbttPta:             "0839",
					TocCode:             "GX",
					MatchedServices:     "22",
					RIDS:                []string{"201607013361763", "201607043443714"},
				},
				Metrics: []Metric{
					Metric{
						ToleranceValue:   "0",
						NumNotTolerance:  "7",
						NumTolerance:     "15",
						PercentTolerance: "68",
						GlobalTolerance:  false,
					},
				},
			},
		},
	}

	c := NewClient("email", "password", BaseURL(server.URL))
	resp, _ := c.GetServiceMetrics(ServiceMetricsRequestData{})

	if !reflect.DeepEqual(expected, *resp) {
		t.Errorf("Expected response to be %v but instead got %v!", expected, *resp)
	}
}
