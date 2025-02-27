package dnac

import (
	"github.com/go-resty/resty/v2"
)

// ConfigurationArchiveService is the service to communicate with the ConfigurationArchive API endpoint
type ConfigurationArchiveService service

// ExportDeviceConfigurations exportDeviceConfigurations
/* Export Device configurations to an encrypted zip file.
@param Content-Type
*/
func (s *ConfigurationArchiveService) ExportDeviceConfigurations() (string, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device-archive/cleartext"

	var operationResult string
	response, err := s.client.R().
		SetResult(&operationResult).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return "", nil, err
	}
	return operationResult, response, err
}
