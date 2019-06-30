# PC monitoring prometheus and grafana
Project that helps PC users to monitor their PC using prometheus and grafana.

- This project needs a proper versioning and releases process, will be added over time.

The current goal is to create a nice Grafana dashboard that will collect details of the running applications and their real power consumption
I am a windows user with linux VM development environment, hence the first main goal is to have this project running on windows OS

The centric piece of the puzzle is the TP-Link Smart plug HS 110, which provides an web API to collect real time power usage of your PC.
This project contains the exporter, written in go.

This project contains the instruction how to download and set up all other applications. Please make sure you have 7zip to extract tar.gz archives

## How to set up windows monitoring
First crate a new directory to hold all third party applications. EG: pc_monitoring
Copy entire configs directory under pc_monitoring
I did not figure out how to use relative paths in cmd scripts, hence please update the start-monitoring.cmd file with the absolute path of your pc_monitoring folder :)

Copy start-monitoring cmd file in pc_monitoring 

##### Prometheus server 
- download windows version at https://prometheus.io/download/  EG: prometheus-[VERSION].windows-amd64.tar.gz
- using 7zip extract content in pc_monitoring under `prometheus`  folder
You can change the prometheus server port number adding  `--web.listen-address="0.0.0.0:19090"` on start arguments

##### Grafana
- download Standalone Windows Binaries from page: https://grafana.com/grafana/download?platform=windows 
- extract content in pc_monitoring under `grafana`  folder
- place grafana_data_sources_prometheus.yaml file in `pc_monitoring\grafana\conf\provisioning\datasources`

##### Windows exporter 
- get the latest version of wmi exporter at: https://github.com/martinlindhe/wmi_exporter/releases  eg: wmi_exporter-0.7.999-amd64.msi
- using 7zip extract the executable  wmi_exporter.exe in pc_monitoring under `wmi` folder

##### tp-link hs exporter 
- you can follow the instructions under tplink-hs-prometheus-exporter README file or for the moment download the executable from the executable directory
- tplink-hs-prometheus-exporter.exe was build for windows x64
- plink-hs-prometheus-exporter was build for linux x64
- pleace the plink-hs-prometheus-exporter.exe in pc_monitoring under `tplink` folder
- update IP address in start-monitoring.cmd file (last line)


#### Start all services using the start-monitoring.cmd script and access http://localhost:3000 to add the JSON dashboard
Dashboard not ready yet
 



