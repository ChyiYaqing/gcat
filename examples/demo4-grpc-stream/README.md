# Set Up

[Caddy] http2 file server: 
gRPC: cpu core to saturate the link when doing Put/Get

| Workload | Benefit | What to do |
|:------|:---|:---|
| Unary request | approx -10% CPU (will increase if proto msg are larger & more complex) | Use vtprotobuf codec |
| Egress stream | 2x reduction in CPU usgae | Use a recent grpc version & a CodecV2 implementation that uses memory pools |
| Ingress stream | 2.5x reduction in CPU usage | - Use a recent grpc version or enable internal memory pooling <br> - in the handler: pool received messages |