### protoc --go_out=plugins=grpc:. helloworld.proto
### rpc GetFeature(Point) returns (Feature) {}
### rpc ListFeatures(Rectangle) returns (stream Feature) {}
### rpc RecordRoute(stream Point) returns (RouteSummary) {}
### rpc RouteChat(stream RouteNote) returns (stream RouteNote) {}