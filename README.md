# Event Driven Commerce

## Overview

- An ecommerce application built applying concepts of distributed systems, microservices and event-driven architecture.
- The `api-gateway` serves REST endpoints to the `web-ui`, and internally performs gRPC calls to all others microservices. Additionally, there is an async communication flow for order payement where `order-svc` sends a message to RabbitMQ and `payment-svc` process it asynchronously.

## Architecture Diagram

![Diagram](docs/event-driven-commerce.svg)

## Technologies used

### Backend

- Golang
- PostgreSQL
- RabbitMQ
- gRPC
- REST

### Frontend

- TypeScript
- NextJS
- Tailwind CSS

## How to run

### Local

- The script below will start all the necessary databases and services

```bash
# Docker and docker-compose are required*
sh scripts/env_up.sh
```
