# Order Management System (OMS) - Microservices in Go

## Overview

This project is a **Microservices-based Order Management System (OMS)** built with **Go (Golang)** and **gRPC**. The goal is to provide a simple yet extensible solution for managing orders, integrating various services like **Order Service**, **Payment Service**, **Kitchen Service**, and **Stock Management** through **gRPC** communication.

Each service is independent and focuses on its respective domain, enabling flexibility, scalability, and better resource management. The architecture showcases a **modular, loosely-coupled system** with microservices principles in action.

## Architecture

* **Gateway Service**: Acts as a communication gateway for all microservices. It handles HTTP requests and forwards them to the appropriate gRPC service.
* **Order Service**: Responsible for managing and processing orders.
* **Payment Service**: Handles payment processing.
* **Kitchen Service**: Manages the preparation and status of food orders.
* **Stock Service**: Handles inventory management and updates stock levels.
* **Common Module**: Contains shared code, gRPC Protobuf definitions, environment variables, and other utilities.

## Technologies

* **Go (Golang)** for microservice development.
* **gRPC** for fast and efficient communication between services.
* **Protocol Buffers** for defining gRPC service methods and message types.
* **Makefile** for automating build processes.

## File Structure

```
.
├── common
│   ├── api
│   │   ├── oms.pb.go
│   │   ├── oms.proto
│   │   └── oms_grpc.pb.go
│   ├── env.go
│   ├── go.mod
│   ├── go.sum
│   ├── json.go
│   └── Makefile
├── gateway
│   ├── go.mod
│   ├── go.sum
│   ├── http_handler.go
│   ├── main.go
│   └── tmp
│       └── main
├── kitchen
│   └── go.mod
├── orders
│   ├── go.mod
│   ├── main.go
│   ├── service.go
│   ├── store.go
│   └── types.go
├── payments
│   └── go.mod
├── stock
│   └── go.mod
└── tmp
    └── build-errors.log
```

## Setup

### Prerequisites

* Go 1.18+ installed.
* gRPC and Protocol Buffers installed. You can follow the installation steps here:
  [gRPC Installation Guide](https://grpc.io/docs/protoc-installation/)

### Install Dependencies

To install all necessary dependencies, run the following command:

```bash
go mod tidy
```

### Running the Services

1. **Start Order Service**:
   Navigate to the `orders` directory and run the following command:

   ```bash
   go run main.go
   ```

2. **Start Payment Service**:
   Navigate to the `payments` directory and run the following command:

   ```bash
   go run main.go
   ```

3. **Start Kitchen Service**:
   Navigate to the `kitchen` directory and run the following command:

   ```bash
   go run main.go
   ```

4. **Start Stock Service**:
   Navigate to the `stock` directory and run the following command:

   ```bash
   go run main.go
   ```

5. **Start Gateway**:
   Navigate to the `gateway` directory and run the following command to start the HTTP-GRPC gateway:

   ```bash
   go run main.go
   ```

The services will communicate via gRPC, with the Gateway service routing incoming HTTP requests to the appropriate microservice.

### Running with Makefile

For convenience, you can also use the `Makefile` to automate building and running services.

```bash
make gen         # Generate gRPC files using protoc
make run         # Run all services
```

## API Documentation

The services expose various gRPC endpoints defined in the `oms.proto` file. You can use a gRPC client (e.g., [grpcurl](https://github.com/fullstorydev/grpcurl)) to interact with the services.

For example:

* **Create Order**: `POST /api/orders`
* **Process Payment**: `POST /api/payments`
* **Update Kitchen Status**: `POST /api/kitchen`
* **Manage Stock**: `POST /api/stock`

Refer to the `oms.proto` file in `common/api` for more details.

## Contributing

Feel free to fork the repository and create pull requests! All contributions are welcome.

1. Fork the repository.
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add new feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request.

## License

This project is licensed under the MIT License.
