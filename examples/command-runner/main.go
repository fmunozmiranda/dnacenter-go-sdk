package main

import (
	"encoding/json"
	"fmt"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
)

// client is DNA Center API client
var client *dnac.Client

func main() {
	var err error
	fmt.Println("Authenticating...")
	client, err = dnac.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"true", "false")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Printing device list  ... PlatformID is C9300-48U")
	getDeviceListQueryParams := &dnac.GetDeviceListQueryParams{
		PlatformID: []string{"C9300-48U"},
	}

	devices, _, err := client.Devices.GetDeviceList(getDeviceListQueryParams)
	if err != nil {
		fmt.Println(err)
	}

	var deviceIDs []string
	for id, device := range devices.Response {
		deviceIDs = append(deviceIDs, device.ID)
		fmt.Println("GET:", id, device.ID, device.MacAddress, device.ManagementIPAddress, device.PlatformID)
	}

	commands := []string{"show version", "show ip interface brief"}
	runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest := &dnac.RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest{
		Commands:    commands,
		DeviceUUIDs: deviceIDs,
		Timeout:     0,
	}
	response, _, err := client.CommandRunner.RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration(runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest)
	taskID := response.Response.TaskID

	fmt.Println(taskID)
	time.Sleep(20 * time.Second)

	taskResponse, _, err := client.Task.GetTaskByID(taskID)
	if err != nil {
		fmt.Println(err)
	}
	type fileIDJSON struct {
		FileID string `json:"fileId"`
	}
	fmt.Println(taskResponse.Response.Progress)
	var fileID fileIDJSON
	err = json.Unmarshal([]byte(taskResponse.Response.Progress), &fileID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileResponse, _, err := client.File.DownloadAFileByFileID(fileID.FileID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileResponse)

}
