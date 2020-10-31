package dnac

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// ClientsService is the service to communicate with the Clients API endpoint
type ClientsService service

// ConnectionInfo is the ConnectionInfo definition
type ConnectionInfo struct {
	Band          string `json:"band,omitempty"`          //
	Channel       string `json:"channel,omitempty"`       //
	ChannelWidth  string `json:"channelWidth,omitempty"`  //
	HostType      string `json:"hostType,omitempty"`      //
	NwDeviceMac   string `json:"nwDeviceMac,omitempty"`   //
	NwDeviceName  string `json:"nwDeviceName,omitempty"`  //
	Protocol      string `json:"protocol,omitempty"`      //
	SpatialStream string `json:"spatialStream,omitempty"` //
	Timestamp     int    `json:"timestamp,omitempty"`     //
	Uapsd         string `json:"uapsd,omitempty"`         //
	Wmm           string `json:"wmm,omitempty"`           //
}

// Detail is the Detail definition
type Detail struct {
	ApGroup          string        `json:"apGroup,omitempty"`          //
	AuthType         string        `json:"authType,omitempty"`         //
	AvgRssi          string        `json:"avgRssi,omitempty"`          //
	AvgSnr           string        `json:"avgSnr,omitempty"`           //
	Channel          string        `json:"channel,omitempty"`          //
	ClientConnection string        `json:"clientConnection,omitempty"` //
	ClientType       string        `json:"clientType,omitempty"`       //
	ConnectedDevice  []string      `json:"connectedDevice,omitempty"`  //
	ConnectionStatus string        `json:"connectionStatus,omitempty"` //
	DataRate         string        `json:"dataRate,omitempty"`         //
	DNSFailure       string        `json:"dnsFailure,omitempty"`       //
	DNSSuccess       string        `json:"dnsSuccess,omitempty"`       //
	Frequency        string        `json:"frequency,omitempty"`        //
	HealthScore      []HealthScore `json:"healthScore,omitempty"`      //
	HostIPV4         string        `json:"hostIpV4,omitempty"`         //
	HostIPV6         []string      `json:"hostIpV6,omitempty"`         //
	HostMac          string        `json:"hostMac,omitempty"`          //
	HostName         string        `json:"hostName,omitempty"`         //
	HostOs           string        `json:"hostOs,omitempty"`           //
	HostType         string        `json:"hostType,omitempty"`         //
	HostVersion      string        `json:"hostVersion,omitempty"`      //
	ID               string        `json:"id,omitempty"`               //
	IosCapable       bool          `json:"iosCapable,omitempty"`       //
	IssueCount       int           `json:"issueCount,omitempty"`       //
	LastUpdated      int           `json:"lastUpdated,omitempty"`      //
	Location         string        `json:"location,omitempty"`         //
	Onboarding       Onboarding    `json:"onboarding,omitempty"`       //
	OnboardingTime   string        `json:"onboardingTime,omitempty"`   //
	Port             string        `json:"port,omitempty"`             //
	Rssi             string        `json:"rssi,omitempty"`             //
	RxBytes          string        `json:"rxBytes,omitempty"`          //
	Snr              string        `json:"snr,omitempty"`              //
	SSID             string        `json:"ssid,omitempty"`             //
	SubType          string        `json:"subType,omitempty"`          //
	TxBytes          string        `json:"txBytes,omitempty"`          //
	UserID           string        `json:"userId,omitempty"`           //
	VlanID           string        `json:"vlanId,omitempty"`           //
	Vnid             string        `json:"vnid,omitempty"`             //
}

// GetClientDetailResponse is the GetClientDetailResponse definition
type GetClientDetailResponse struct {
	ConnectionInfo ConnectionInfo `json:"connectionInfo,omitempty"` //
	Detail         Detail         `json:"detail,omitempty"`         //
	Topology       Topology       `json:"topology,omitempty"`       //
}

// GetClientEnrichmentDetailsResponse is the GetClientEnrichmentDetailsResponse definition
type GetClientEnrichmentDetailsResponse struct {
	ConnectedDevice []ConnectedDevice `json:"connectedDevice,omitempty"` //
	IssueDetails    IssueDetails      `json:"issueDetails,omitempty"`    //
	UserDetails     UserDetails       `json:"userDetails,omitempty"`     //
}

// GetOverallClientHealthResponse is the GetOverallClientHealthResponse definition
type GetOverallClientHealthResponse struct {
	Response []ClientsResponse `json:"response,omitempty"` //
}

// ImpactedHosts is the ImpactedHosts definition
type ImpactedHosts struct {
	ConnectedInterface string   `json:"connectedInterface,omitempty"` //
	FailedAttempts     int      `json:"failedAttempts,omitempty"`     //
	HostName           string   `json:"hostName,omitempty"`           //
	HostOs             string   `json:"hostOs,omitempty"`             //
	HostType           string   `json:"hostType,omitempty"`           //
	Location           Location `json:"location,omitempty"`           //
	MacAddress         string   `json:"macAddress,omitempty"`         //
	SSID               string   `json:"ssid,omitempty"`               //
	Timestamp          int      `json:"timestamp,omitempty"`          //
}

// Location is the Location definition
type Location struct {
	ApsImpacted []string `json:"apsImpacted,omitempty"` //
	Area        string   `json:"area,omitempty"`        //
	Building    string   `json:"building,omitempty"`    //
	Floor       string   `json:"floor,omitempty"`       //
	SiteID      string   `json:"siteId,omitempty"`      //
	SiteType    string   `json:"siteType,omitempty"`    //
}

// Onboarding is the Onboarding definition
type Onboarding struct {
	AaaRootcauseList     []string `json:"aaaRootcauseList,omitempty"`     //
	AaaServerIP          string   `json:"aaaServerIp,omitempty"`          //
	AssocDoneTime        string   `json:"assocDoneTime,omitempty"`        //
	AssocRootcauseList   []string `json:"assocRootcauseList,omitempty"`   //
	AuthDoneTime         string   `json:"authDoneTime,omitempty"`         //
	AverageAssocDuration string   `json:"averageAssocDuration,omitempty"` //
	AverageAuthDuration  string   `json:"averageAuthDuration,omitempty"`  //
	AverageDhcpDuration  string   `json:"averageDhcpDuration,omitempty"`  //
	AverageRunDuration   string   `json:"averageRunDuration,omitempty"`   //
	DhcpDoneTime         string   `json:"dhcpDoneTime,omitempty"`         //
	DhcpRootcauseList    []string `json:"dhcpRootcauseList,omitempty"`    //
	DhcpServerIP         string   `json:"dhcpServerIp,omitempty"`         //
	MaxAssocDuration     string   `json:"maxAssocDuration,omitempty"`     //
	MaxAuthDuration      string   `json:"maxAuthDuration,omitempty"`      //
	MaxDhcpDuration      string   `json:"maxDhcpDuration,omitempty"`      //
	MaxRunDuration       string   `json:"maxRunDuration,omitempty"`       //
	OtherRootcauseList   []string `json:"otherRootcauseList,omitempty"`   //
}

// ClientsResponse is the Response definition
type ClientsResponse struct {
	ScoreDetail []ScoreDetail `json:"scoreDetail,omitempty"` //
	SiteID      string        `json:"siteId,omitempty"`      //
}

// ScoreCategory is the ScoreCategory definition
type ScoreCategory struct {
	ScoreCategory string `json:"scoreCategory,omitempty"` //
	Value         string `json:"value,omitempty"`         //
}

// ScoreDetail is the ScoreDetail definition
type ScoreDetail struct {
	ClientCount       int           `json:"clientCount,omitempty"`       //
	ClientUniqueCount int           `json:"clientUniqueCount,omitempty"` //
	Endtime           int           `json:"endtime,omitempty"`           //
	ScoreCategory     ScoreCategory `json:"scoreCategory,omitempty"`     //
	ScoreList         []ScoreList   `json:"scoreList,omitempty"`         //
	ScoreValue        int           `json:"scoreValue,omitempty"`        //
	Starttime         int           `json:"starttime,omitempty"`         //
}

// ScoreList is the ScoreList definition
type ScoreList struct {
	ClientCount       int           `json:"clientCount,omitempty"`       //
	ClientUniqueCount string        `json:"clientUniqueCount,omitempty"` //
	Endtime           int           `json:"endtime,omitempty"`           //
	ScoreCategory     ScoreCategory `json:"scoreCategory,omitempty"`     //
	ScoreValue        int           `json:"scoreValue,omitempty"`        //
	Starttime         int           `json:"starttime,omitempty"`         //
}

// GetClientDetailQueryParams defines the query parameters for this request
type GetClientDetailQueryParams struct {
	Timestamp  string `url:"timestamp,omitempty"`  // Epoch time(in milliseconds) when the Client health data is required
	MacAddress string `url:"macAddress,omitempty"` // MAC Address of the client
}

// GetClientDetail getClientDetail
/* Returns detailed Client information retrieved by Mac Address for any given point of time.
@param timestamp Epoch time(in milliseconds) when the Client health data is required
@param macAddress MAC Address of the client
*/
func (s *ClientsService) GetClientDetail(getClientDetailQueryParams *GetClientDetailQueryParams) (*GetClientDetailResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/client-detail"

	queryString, _ := query.Values(getClientDetailQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetClientDetailResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetClientDetailResponse)
	return result, response, err

}

// GetClientEnrichmentDetails getClientEnrichmentDetails
/* Enriches a given network End User context (a network user-id or end user’s device Mac Address) with details about the user, the devices that the user is connected to and the assurance issues that the user is impacted by
@param entity_type Client enrichment details can be fetched based on either User ID or Client MAC address. This parameter value must either be network_user_id/mac_address
@param entity_value Contains the actual value for the entity type that has been defined
@param issueCategory The category of the DNA event based on which the underlying issues need to be fetched
*/
func (s *ClientsService) GetClientEnrichmentDetails() (*GetClientEnrichmentDetailsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/client-enrichment-details"

	response, err := RestyClient.R().
		SetResult(&GetClientEnrichmentDetailsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetClientEnrichmentDetailsResponse)
	return result, response, err

}

// GetOverallClientHealthQueryParams defines the query parameters for this request
type GetOverallClientHealthQueryParams struct {
	Timestamp string `url:"timestamp,omitempty"` // Epoch time(in milliseconds) when the Client health data is required
}

// GetOverallClientHealth getOverallClientHealth
/* Returns Overall Client Health information by Client type (Wired and Wireless) for any given point of time
@param timestamp Epoch time(in milliseconds) when the Client health data is required
*/
func (s *ClientsService) GetOverallClientHealth(getOverallClientHealthQueryParams *GetOverallClientHealthQueryParams) (*GetOverallClientHealthResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/client-health"

	queryString, _ := query.Values(getOverallClientHealthQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetOverallClientHealthResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetOverallClientHealthResponse)
	return result, response, err

}
