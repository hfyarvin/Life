## node-session
- CreateNodeSession(CreateNodeSessionRequest) returns (CreateNodeSessionResponse) {} //NwkID
- GetNodeSession(GetNodeSessionRequest) returns (GetNodeSessionResponse) {} //DevEUI
- UpdateNodeSession(UpdateNodeSessionRequest) returns (UpdateNodeSessionResponse) {}
- DeleteNodeSession(DeleteNodeSessionRequest) returns (DeleteNodeSessionResponse) {} //DevAddr

## DevAddr
- GetRandomDevAddr(GetRandomDevAddrRequest) returns (GetRandomDevAddrResponse) {}

## 
- EnqueueDataDownMACCommand(EnqueueDataDownMACCommandRequest) returns (EnqueueDataDownMACCommandResponse) {}

## Data
- PushDataDown(PushDataDownRequest) returns (PushDataDownResponse) {} //downlink payload to the node --(class-c)
- SendProprietaryPayload(SendProprietaryPayloadRequest) returns (SendProprietaryPayloadResponse) {}

## Gateway
- CreateGateway(CreateGatewayRequest) returns (CreateGatewayResponse) {}
- GetGateway(GetGatewayRequest) returns (GetGatewayResponse) {}
- UpdateGateway(UpdateGatewayRequest) returns (UpdateGatewayResponse) {}
- ListGateways(ListGatewayRequest) returns (ListGatewayResponse) {} //列表
- DeleteGateway(DeleteGatewayRequest) returns (DeleteGatewayResponse) {}
- GenerateGatewayToken(GenerateGatewayTokenRequest) returns (GenerateGatewayTokenResponse) {} //a JWT token
- GetGatewayStats(GetGatewayStatsRequest) returns (GetGatewayStatsResponse) {} //of an existing gateway

## 
- GetFrameLogsForDevEUI(GetFrameLogsForDevEUIRequest) returns (GetFrameLogsResponse) {} //un/downlink frame

## channel-configuration
- CreateChannelConfiguration(CreateChannelConfigurationRequest) returns (CreateChannelConfigurationResponse) {}
- GetChannelConfiguration(GetChannelConfigurationRequest) returns (GetChannelConfigurationResponse) {} //for given ID
- UpdateChannelConfiguration(UpdateChannelConfigurationRequest) returns (UpdateChannelConfigurationResponse) {}
- DeleteChannelConfiguration(DeleteChannelConfigurationRequest) returns (DeleteChannelConfigurationResponse) {}
- ListChannelConfigurations(ListChannelConfigurationsRequest) returns (ListChannelConfigurationsResponse) {}

## extra channel
- CreateExtraChannel(CreateExtraChannelRequest) returns (CreateExtraChannelResponse) {}
- UpdateExtraChannel(UpdateExtraChannelRequest) returns (UpdateExtraChannelResponse) {}
- DeleteExtraChannel(DeleteExtraChannelRequest) returns (DeleteExtraChannelResponse) {}
- GetExtraChannelsForChannelConfigurationID(GetExtraChannelsForChannelConfigurationIDRequest) returns (GetExtraChannelsForChannelConfigurationIDResponse) {}