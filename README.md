# Upfluence Mini Project

## Overview

This project simulates a basic influencer campaign system using Go and RabbitMQ. It demonstrates asynchronous communication between two services.

## Project Structure

- `types/types.go` → shared structs (e.g., Campaign)
- `campaign-service/publisher.go` → publishes campaign messages to RabbitMQ
- `analytics-service/consumer.go` → consumes and processes messages from RabbitMQ

## Running the Project

### 1. Start RabbitMQ (e.g., via Docker)

```bash
docker run -d --hostname my-rabbit --name some-rabbit -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

### 2. Start the consumer

```bash
cd analytics-service
go run consumer.go
```

### 3. Start the publisher

```bash
cd campaign-service
go run publisher.go
```

## Expected Behavior

You should see campaign messages being published and processed in real-time.
