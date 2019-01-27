package hsp

type serviceDetailsService struct {
	client *Client
}

// Location contains data from the GetServiceDetails response
type Location struct {
	Location       string
	GbttPtd        string `json:"gbtt_ptd"`
	GbttPta        string `json:"gbtt_pta"`
	ActualTD       string `json:"actual_td"`
	ActualTA       string `json:"actual_ta"`
	LateCancReason string `json:"late_canc_reason"`
}

// ServiceAttributesDetails contains data from the GetServiceDetails response
type ServiceAttributesDetails struct {
	DateOfService string `json:"date_of_service"`
	TocCode       string `json:"toc_code"`
	Rid           string
	Locations     []Location
}

// ServiceDetailsResponseData contains the response from ServiceMetrics API request
type ServiceDetailsResponseData struct {
	ServiceAttributesDetails `json:"serviceAttributesDetails"`
}

func (s *serviceDetailsService) GetServiceDetails(rid string) (*ServiceDetailsResponseData, *Error) {
	reqData := struct {
		Rid string `json:"rid"`
	}{Rid: rid}

	req, err := s.client.newRequest("/serviceDetails", reqData)
	if err != nil {
		return nil, err
	}

	var resp ServiceDetailsResponseData
	_, err = s.client.do(req, &resp)

	return &resp, err
}
