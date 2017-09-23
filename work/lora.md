# LoRa(Long Range)  
## 1. 特点
- 距离
- 抗干扰
- 组网容量大
- 带宽小

## 2. LoRaWAN (Long Range Wild Area Network)
- star topology
- bidirectional: 终端通讯双向
- 资料率： 0.3 kbps - 50 kbps

## 3. LoRaWAN优点
- 长距离
- 大容量
- 安全
- 低功耗
- 一网络多网关
- 三类终端
- 低成本

## 4.LoRaWAN对Server四层架构
- NS(Network Server)
- AS(Application Server)
- NC(Network controller)
- CS(Customer Server)
### - 注：LoRaWAN规定GW与NS的接口协议是json(MQTT-->JSON转换 ora-gateway-bridge)

## 术语
- 终端设备(end-point devices)
- 上行传输(uplink transmission)
- 下行传输(downlink transmission)
- 接入点(AP access point)
- 网关(GW gateway)