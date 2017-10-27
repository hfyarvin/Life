# service Node
## node
- /api/nodes  --post create
- /api/nodes/{devEUI}  --get
- /api/nodes/{decEUI}  --delete
- /api/applications/{applicationID}/nodes  --get
- /api/nodes/{devEUI}  --put update

## activation
- /api/nodes/{decEUI}/activation  --post 激活
- /api/nodes/{devEUI}/activation  --get

## others
- /api/nodes/getRandomDevAddr  --get
- /api/nodes/{devEUI}/frames  --get