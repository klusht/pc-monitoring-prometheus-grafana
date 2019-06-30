START /b /d ".\prometheus\" prometheus.exe --config.file="E:\projects\pc_monitoring\configs\prometheus.yaml"

START /b /d ".\grafana\bin\" grafana-server.exe --config="E:\projects\pc_monitoring\configs\grafana.ini"

START /b /d  ".\wmi\" wmi_exporter.exe --collectors.enabled="cpu,cs,memory,os,process,system"

START /b /d ".\tplink" tplink-hs-prometheus-exporter.exe 192.168.0.23 9999 v1