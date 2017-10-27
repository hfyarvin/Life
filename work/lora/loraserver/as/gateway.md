# service GateWay
## channel_configuration
- /api/gateways/channelconfigurations  --post 
- /api/gateways/channelconfigurations/{id}  --get
- /api/gateways/channelconfigurations/{id}  --put
- /api/gateways/channelconfigurations/{id}  --delete
- /api/gateways/channelconfigurations  --get (list)

## extra_channel
- /api/gateways/extrachannels  --post
- /api/gateways/extrachannels/{id}  --put
- /api/gateways/extrachannels/{id}  --delete
- /api/gateways/channelconfigurations/{id}/extrachannels  --get

## gateway
- /api/gateways  --post
- /api/gateways  --get  (list)
- /api/gateways/{mac}  --get  (for mac address)
- /api/gateways/{mac}  --put  
- /api/gateways/{mac}  --delete  
- /api/gateways/{mac}/stats  --get
- /api/gateways/{mac}/ping/last  --get