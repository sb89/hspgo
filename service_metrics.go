package hsp

type serviceMetricsService struct {
	client *Client
}

// ServiceMetricsRequestData is used to send request data to ServiceMetrics API
type ServiceMetricsRequestData struct {
	FromLoc  string `json:"from_loc"`
	ToLoc    string `json:"to_loc"`
	FromTime string `json:"from_time"`
	ToTime   string `json:"to_time"`
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
	Days     string `json:"days"`
}

// Header contains data from the GetServiceMetrics response
type Header struct {
	FromLocation string `json:"from_location"`
	ToLocation   string `json:"to_location"`
}

// Service contains data from the GetServiceMetrics response
type Service struct {
	ServiceAttributesMetrics `json:"serviceAttributesMetrics"`
	Metrics                  []Metric
}

// ServiceAttributesMetrics contains data from the GetServiceMetrics response
type ServiceAttributesMetrics struct {
	OriginLocation      string   `json:"origin_location"`
	DestinationLocation string   `json:"destination_location"`
	GbttPtd             string   `json:"gbtt_ptd"`
	GbttPta             string   `json:"gbtt_pta"`
	TocCode             string   `json:"toc_code"`
	MatchedServices     string   `json:"matched_services"`
	RIDS                []string `json:"rids"`
}

// Metric contains data from the GetServiceMetrics response
type Metric struct {
	ToleranceValue   string `json:"tolerance_value"`
	NumNotTolerance  string `json:"num_not_tolerance"`
	NumTolerance     string `json:"num_tolerance"`
	PercentTolerance string `json:"percent_tolerance"`
	GlobalTolerance  bool   `json:"global_tolerance"`
}

// ServiceMetricsResponseData contains the response from ServiceMetrics API request
type ServiceMetricsResponseData struct {
	Header   `json:"header"`
	Services []Service `json:"Services"`
}

func (s *serviceMetricsService) GetServiceMetrics(reqData ServiceMetricsRequestData) (*ServiceMetricsResponseData, *Error) {
	req, err := s.client.newRequest("/serviceMetrics", reqData)
	if err != nil {
		return nil, err
	}

	var resp ServiceMetricsResponseData
	_, err = s.client.do(req, &resp)

	return &resp, err
}
