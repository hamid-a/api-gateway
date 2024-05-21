# API Gateway
### Overview

This document outlines the implementation details of an API Gateway in Go using the Gin framework. The gateway serves HTTP requests and forwards them to gRPC upstream services. Key features include:

* Circuit breaker
* Authentication capability
* Multiple backend support with load balancing
* Standard zap logging
* YAML configuration
* Dynamic routing

### Prerequisites
* Go 1.21+
* Gin framework
* Zap logging library
* Koanf for configuration management

### Golang Standard Project Structure
This project is defined based on the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

### YAML Configuration
You can run the project by replacing `configs/configs.sample.yaml` with `configs/configs.yaml`. We are using the `koanf` library to parse the configuration file. It is more lightweight and has fewer dependencies compared to `Viper`.

#### Rules:
Each routing rule consists of rule objects:

```yaml
rules:
    - name: "service-a-router"
      path: "/service-a-path"
      auth: true
      upstream: "ServiceA"
      url: "upstream-a-path"
      methods: ["GET", "POST", "OPTIONS"]
```

- In each rule, you can define the `url` prefix, and the requests will be routed to the defined `upstream`.
- You can define any method for rules.
- The authentication section is not implemented yet, but the middleware and the logic places are defined.
- Rules should have a unique `name` and `path`.
- `path` is the URL that is called in the service, and its request will be routed to the specified `upstream`.

#### Upstreams:
```yaml
upstreams:
    - name: ServiceA
      backends:
        - name: "ServiceA-1"
          connection: http
          addr: "http://service-a1.svc.cluster.local"
          timeout: 3s
          cb:
            enabled: true
            resetInterval: "60s"
            openTimeout: "60s"
            maxRequests: 2
            minRequests: 3
            failureRatioThreshold: 0.6
```

- Each upstream should have a unique name.
- You can define multiple backends for each upstream. Load will be distributed using a round-robin algorithm.
- Both gRPC and HTTP are supported for backends, and each upstream's backend should have the same protocol.
- Each backend has a specified circuit breaker.
- Examples for each protocol are defined in the `config.sample.yaml` file.

*For gRPC usage, you should define the `.proto` file and implement it in the upstream package. Every implementation should implement the `Service` interface.*

*HTTP protocols are the same and can be used by implementing the same as the `service-a.go` upstream.*

### Run
After replacing `configs.sample.yaml` with `configs.yaml` and defining your specific configs, you can run the application with the command below:

```bash
go run cmd/api-gateway/api-gateway.go
```

You should have Go 1.21+ to run the service, but the Dockerfile is included, and you can run it with the Dockerfile.

```bash
docker build -t api-gateway:latest .
```

```bash
docker run -p 8080:8080 -v $(pwd)/configs/configs.yaml:/app/configs/configs.yaml api-gateway:latest
```

For testing the application, you can run the command below:

```bash
curl -X GET 'http://127.0.0.1:8080/service-a-path'
```

This request will be routed to the service-a upstream with the HTTP protocol.
