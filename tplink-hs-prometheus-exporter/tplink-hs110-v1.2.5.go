package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func scrapRealtimeMetricsV1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(GetRealtimeUsageV1()))
}

func scrapTplinkInfoV1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(GetSystemInfoV1()))
}

func GetRealtimeUsageV1() string {
	command := `{"emeter":{"get_realtime":{},"get_vgain_igain":{}}}`

	jsonReceived := GetHSDetails(command)
	var res ResponseV1
	json.Unmarshal([]byte(jsonReceived), &res)

	response := "# HELP emeter_realtime_usage  usage statistic from HS 110\n"
	response += "# TYPE emeter_realtime_usage gauge\n"
	response += fmt.Sprintf("emeter_realtime_usage{device=\"%v\", metric=\"current\", unit=\"amper\"}  %f\n", deviceAlias, res.Emeter.Get_realtime.Current)
	response += fmt.Sprintf("emeter_realtime_usage{device=\"%v\", metric=\"voltage\", unit=\"volt\"}  %f\n", deviceAlias, res.Emeter.Get_realtime.Voltage)
	response += fmt.Sprintf("emeter_realtime_usage{device=\"%v\", metric=\"power\", unit=\"watt\"}  %f\n", deviceAlias, res.Emeter.Get_realtime.Power)
	response += fmt.Sprintf("emeter_realtime_usage{device=\"%v\", metric=\"total\"}  %v\n", deviceAlias, res.Emeter.Get_realtime.Total)
	response += "# HELP emeter_gains usage statistic from HS110\n"
	response += "# TYPE emeter_gains gauge\n"
	response += fmt.Sprintf("emeter_gains{device=\"%v\", metric=\"vgain\"}  %v\n", deviceAlias, res.Emeter.Get_vgain_igain.Vgain)
	response += fmt.Sprintf("emeter_gains{device=\"%v\", metric=\"igain\"}  %v\n", deviceAlias, res.Emeter.Get_vgain_igain.Igain)
	return response
}

func GetSystemInfoV1() string {
	command := `{"system":{"get_sysinfo":{}},"cnCloud":{"get_info":{},"get_intl_fw_list":{}},"time":{"get_time":{}},"schedule":{"get_next_action":{}},"count_down":{"get_rules":{}},"anti_theft":{"get_rules":{}}}`

	jsonReceived := GetHSDetails(command)
	var res ResponseV1
	json.Unmarshal([]byte(jsonReceived), &res)

	response := "# HELP system_info usage data from HS 110\n"
	response += "# TYPE system_info gauge\n"
	response += fmt.Sprintf("system_info{device=\"%v\", sw_ver=\"%v\"} 1\n", deviceAlias, res.System.Get_sysinfo.Sw_ver)
	response += fmt.Sprintf("system_info{device=\"%v\", hw_ver=\"%v\"} 1\n", deviceAlias, res.System.Get_sysinfo.Hw_ver)
	response += fmt.Sprintf("system_info{device=\"%v\", type=\"%v\"} 1\n", deviceAlias, res.System.Get_sysinfo.Type)
	response += fmt.Sprintf("system_info{device=\"%v\", model=\"%v\"} 1\n", deviceAlias, res.System.Get_sysinfo.Model)
	response += fmt.Sprintf("system_info{device=\"%v\", deviceId=\"%v\"} 1\n", deviceAlias, res.System.Get_sysinfo.DeviceId)
	response += fmt.Sprintf("system_info{device=\"%v\", hwId=\"%v\"} 1\n", deviceAlias, res.System.Get_sysinfo.HwId)
	response += fmt.Sprintf("system_info{device=\"%v\", fwId=\"%v\"} 1\n", deviceAlias, res.System.Get_sysinfo.FwId)
	response += fmt.Sprintf("system_info{device=\"%v\", oemId=\"%v\"} 1\n", deviceAlias, res.System.Get_sysinfo.OemId)
	response += fmt.Sprintf("system_info{device=\"%v\", alias=\"%v\"} 1\n", deviceAlias, res.System.Get_sysinfo.Alias)
	response += fmt.Sprintf("system_info{device=\"%v\", dev_name=\"%v\"} 1\n", deviceAlias, res.System.Get_sysinfo.Dev_name)
	response += fmt.Sprintf("system_info{device=\"%v\", icon_hash=\"%v\"} 1\n", deviceAlias, res.System.Get_sysinfo.Icon_hash)
	response += fmt.Sprintf("system_info_relay_state{device=\"%v\"} %v\n", deviceAlias, res.System.Get_sysinfo.Relay_state)
	response += fmt.Sprintf("system_info_on_time{device=\"%v\"} %v\n", deviceAlias, res.System.Get_sysinfo.On_time)
	response += fmt.Sprintf("system_info{device=\"%v\", active_mode=\"%v\"} 1\n", deviceAlias, res.System.Get_sysinfo.Active_mode)
	response += fmt.Sprintf("system_info{device=\"%v\", feature=\"%v\"} 1\n", deviceAlias, res.System.Get_sysinfo.Feature)
	response += fmt.Sprintf("system_info_updating{device=\"%v\"} %v\n", deviceAlias, res.System.Get_sysinfo.Updating)
	response += fmt.Sprintf("system_info_rssi{device=\"%v\"} %v\n", deviceAlias, res.System.Get_sysinfo.Rssi)
	response += fmt.Sprintf("system_info_led_off{device=\"%v\"} %v\n", deviceAlias, res.System.Get_sysinfo.Led_off)
	response += fmt.Sprintf("system_info_latitude{device=\"%v\"} %v\n", deviceAlias, res.System.Get_sysinfo.Latitude)
	response += fmt.Sprintf("system_info_longitude{device=\"%v\"} %v\n", deviceAlias, res.System.Get_sysinfo.Longitude)
	response += "# HELP cnCloud_info usage data from HS 110\n"
	response += "# TYPE cnCloud_info gauge\n"
	response += fmt.Sprintf("cnCloud_info{device=\"%v\", username=\"%v\"} 1\n", deviceAlias, res.CnCloud.Get_info.Username)
	response += fmt.Sprintf("cnCloud_info{device=\"%v\", server=\"%v\"} 1\n", deviceAlias, res.CnCloud.Get_info.Server)
	response += fmt.Sprintf("cnCloud_info_cld_binded{device=\"%v\"} %v\n", deviceAlias, res.CnCloud.Get_info.Binded)
	response += fmt.Sprintf("cnCloud_info_cld_connection{device=\"%v\"} %v\n", deviceAlias, res.CnCloud.Get_info.Cld_connection)
	response += fmt.Sprintf("cnCloud_info_cld_illegalType{device=\"%v\"} %v\n", deviceAlias, res.CnCloud.Get_info.IllegalType)
	response += fmt.Sprintf("cnCloud_info_cld_stopConnect{device=\"%v\"} %v\n", deviceAlias, res.CnCloud.Get_info.StopConnect)
	response += fmt.Sprintf("cnCloud_info_cld_tcspStatus{device=\"%v\"} %v\n", deviceAlias, res.CnCloud.Get_info.TcspStatus)
	response += fmt.Sprintf("cnCloud_info{device=\"%v\", fwDlPage=\"%v\"} 1\n", deviceAlias, res.CnCloud.Get_info.FwDlPage)
	response += fmt.Sprintf("cnCloud_info{device=\"%v\", tcspInfo=\"%v\"} 1\n", deviceAlias, res.CnCloud.Get_info.TcspInfo)
	response += fmt.Sprintf("cnCloud_info_cld_fwNotifyType{device=\"%v\"} %v\n", deviceAlias, res.CnCloud.Get_info.FwNotifyType)
	response += "# HELP schedule_get_next_action data from HS 110\n"
	response += "# TYPE schedule_get_next_action gauge\n"
	response += fmt.Sprintf("schedule_get_next_action{device=\"%v\", type=\"%v\"} 1\n", deviceAlias, res.Schedule.Get_next_action.Type)
	return response
}

type ResponseV1 struct {
	Emeter   EmeterV1
	System   SystemV1
	CnCloud  CnCloudV1
	Schedule ScheduleV1
}

type ScheduleV1 struct {
	Get_next_action Get_next_actionV1
}

type Get_next_actionV1 struct {
	Type    string
	ErrCode int
}

type CnCloudV1 struct {
	Get_info         Get_infoV1
	Get_intl_fw_list Get_intl_fw_listV1
}

type Get_intl_fw_listV1 struct {
	Fw_list  string
	Err_code int
}

type Get_infoV1 struct {
	Username       string
	Server         string
	Binded         int
	Cld_connection int
	IllegalType    int
	StopConnect    int
	TcspStatus     int
	FwDlPage       string
	TcspInfo       string
	FwNotifyType   int
	ErrCode        int
}

type SystemV1 struct {
	Get_sysinfo Get_sysinfoV1
}

type Get_sysinfoV1 struct {
	Err_code    int
	Sw_ver      string
	Hw_ver      float32
	Type        string
	Model       string
	Mac         string
	DeviceId    string
	HwId        string
	FwId        string
	OemId       string
	Alias       string
	Dev_name    string
	Icon_hash   string
	Relay_state int
	On_time     int64
	Active_mode string
	Feature     string
	Updating    int
	Rssi        int
	Led_off     int
	Latitude    float64
	Longitude   float64
}

type EmeterV1 struct {
	Get_realtime    GetRealtimeV1
	Get_vgain_igain Get_vgain_igainV1
}

type Get_vgain_igainV1 struct {
	Vgain    int
	Igain    int
	Err_code int
}

type GetRealtimeV1 struct {
	Current  float64
	Voltage  float64
	Power    float64
	Total    float64
	Err_code int
}
