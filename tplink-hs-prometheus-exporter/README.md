#Simple scraper/exporter of tp-link Smart-plug for prometheus 
The exporter has been tested on HS 110.

Many thanks for this indispensable source

`https://github.com/softScheck/tplink-smartplug/blob/master/tplink_smartplug.py`

###The exporter requires passing the IP address of the Smart-plug, a localhost port number and a response version 


#####Smart-plug IP should be reserved on your router.Access your router and reserve an IP based for the Smart-plug MAC address.
- After successfully setting up the plug you will be able to see the MAC address in your router UI under "Wireless Attached devices" using the name "HS110". 
- Take a note of its MAC address of form 00:00:00:00:00:00
- Go to DHCP and "Add reserved rule", pass the MAC and either choose a new IP or use the one that has been already assigned (If you choose a new one, you must reconnect your plug) 

##### The exporter requires the port to listen to. It is done like this to have multiple exporters running on the same machine and scrape multiple plugs.

##### The Response version available are v1 or v2
There are different JSON fields name between some versions and also some values are under a different metric scale.

If response from localhost:[port]/metrics returns 0 for v1, please update the start argument to v2... and vice versa  

###Development
There are no tests present. For the moment I feel is too small to get any value from test coverage.
run the application using `go run *.go [IP] [PORT] [v1/v2]`

EG: `go run *.go 192.168.0.100 1999 v1`


###Build for different platforms
Clone this repository and execute the commands 

####Build for windows 
- command `env GOOS=windows GOARCH=amd64 go build`
- this will generate tplink-hs-prometheus-exporter.exe which requires the HS110 IP to be passed and a port number to listen on. 

Either define a start-exporter.bat file and use a START command:
`START /b /d ".\" tplink-hs-prometheus-exporter.exe 192.168.0.100 1999 v1`

Or create a shortcut of the app and UPDATE Target to `C:\Users\user\Desktop\tplink-hs-prometheus-exporter.exe 192.168.0.100 1999 v1`


####Build for Linux 
- command `GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build`
- this will generate tplink-hs-prometheus-exporter in the same directory 

Run it using `./tplink-hs-prometheus-exporter 192.168.0.100 1999 v1`