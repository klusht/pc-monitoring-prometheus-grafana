package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

var tplinkPort = ":9999"
var tplinkSocket = ""
var exporterPort = ":"
var deviceAlias = "Unknown"
var responsePayloadVersion = "v0"

func main() {
	//SmartHome protocol runs on TCP port 9999
	if len(os.Args) < 4 {
		fmt.Println("ERROR: Missing start arguments.\n >>> Please pass the Smart-Plug IP to scrape, the port for the exporter AND the version of response payload(v1 or v2). /nEG  tplink-hs-prometheus-exporter.exe 192.168.0.100 1999 v1")
		os.Exit(0)
	}
	tplinkSocket = strings.Trim(os.Args[1], "") + tplinkPort
	exporterPort = exporterPort + strings.Trim(os.Args[2], "")
	deviceAlias = GetDeviceAlias()
	responsePayloadVersion = strings.Trim(os.Args[3], "")

	fmt.Println("Exporter will query " + deviceAlias + " Smart-plug on socket: " + tplinkSocket)
	fmt.Println("Exporter Server listening on localhost" + exporterPort + " /metrics and /info")

	if responsePayloadVersion == "v1" {
		http.HandleFunc("/metrics", scrapRealtimeMetricsV1)
		http.HandleFunc("/info", scrapTplinkInfoV1)
	}

	if responsePayloadVersion == "v2" {
		http.HandleFunc("/metrics", scrapRealtimeMetricsV2)
		http.HandleFunc("/info", scrapTplinkInfoV2)
	}

	if err := http.ListenAndServe(exporterPort, nil); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func GetDeviceAlias() string {
	command := `{"system":{"get_sysinfo":{}}}`
	jsonReceived := GetHSDetails(command)
	var res ResponseV1
	json.Unmarshal([]byte(jsonReceived), &res)
	return res.System.Get_sysinfo.Alias
}

func buildSendingMessage(command string) []byte {
	var msgLen = []byte{0, 0, 0, byte(len(command))}
	var key byte = 171
	var output []byte
	for i := 0; i < len(command); i++ {
		var step = key ^ byte(command[i])
		key = step
		output = append(output, byte(step))
	}
	return append(msgLen, output...) // prepend length
}

func decodeResponse(response []byte) string {
	var key byte = 171
	var output []byte
	for i := 0; i < len(response); i++ {
		var step = key ^ byte(response[i])
		key = byte(response[i])
		output = append(output, byte(step))
	}
	return string(output)
}

func GetHSDetails(command string) string {
	connection, err := net.Dial("tcp", tplinkSocket)
	if err != nil {
		fmt.Println(err)
	}

	defer connection.Close()
	message := buildSendingMessage(command)

	connection.Write(message)
	buff := make([]byte, 4096)
	bytesReceived, _ := connection.Read(buff)
	return decodeResponse(buff[4:bytesReceived])
}
