# Facility Order API

A simple REST API built with Go and Gin to manage facilities and create orders.
The service stores data in MySQL and uses SQL-based migrations for schema management.


## Features

### Facility Management

* Retrieve all facilities
* Retrieve facility by code
* Create a facility

### Order Management

* Create orders linked to a facility
* Retrieve all orders
* Retrieve order by ID

## Tech Stack

* Go
* Gin Web Framework
* MySQL
* SQL Migrations

---

# Setup

## Prerequisites
- Go 1.25+
- MySQL 5.7+

## 1. Install Dependencies
```bash
go mod tidy
```

## 2. Setup Database

### Login to MySQL
```bash
mysql -u <username> -p
```

### Create Database
```sql
CREATE DATABASE facility_order_service;
```

## 3. Environment Variables

Create `.env` file in the root directory:
```env
DB_USER=root
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=facility_order_service
SERVER_PORT=8081
KAFKA_BROKERS=localhost:9092
```

## 4. Run Server
```bash
go run ./cmd/api
```

Server starts at `http://localhost:8080`

---

# Kafka Setup (Ubuntu)

This section covers setting up Apache Kafka for event streaming on Ubuntu.

## Prerequisites
- Docker and Docker Compose installed
- Port 9092 available for Kafka
- Port 2181 available for Zookeeper

## 1. Install Docker on Ubuntu

### Using Official Docker Installation Script
```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Use Docker without sudo (optional)
sudo usermod -aG docker $USER
newgrp docker
```

### Verify Installation
```bash
docker --version
docker compose version
```

## 2. Update Environment Variables

Ensure Kafka broker address is added to `.env`:
```env
KAFKA_BROKERS=localhost:9092
```

## 3. Start Kafka

The `docker-compose.yml` file contains Zookeeper and Kafka configuration.

### Start Services
```bash
docker compose up -d
```

### Check Status
```bash
docker compose ps
```

You should see:
- `zookeeper` - RUNNING
- `kafka` - RUNNING
- `kafka-ui` - RUNNING

### View Kafka UI (Optional)
Open browser: `http://localhost:8080`

## 4. Create Kafka Topic

Create the `order_created` topic:
```bash
docker compose exec kafka kafka-topics \
  --create \
  --topic order_created \
  --bootstrap-server localhost:9092 \
  --partitions 1 \
  --replication-factor 1
```

### Verify Topic Creation
```bash
docker compose exec kafka kafka-topics --list --bootstrap-server localhost:9092
```

You should see `order_created` in the output.

## 5. Run Application with Kafka

Start the application:
```bash
go run ./cmd/api
```

You should see:
```
✓ Consumer started, waiting for messages...
```

## 6. Test Event Publishing

### Create an Order
```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{
    "id": "ORD001",
    "facility_code": "VND-2404"
  }'
```

### Check Application Logs

In the terminal where your app is running, you should see:

```
✓ Published order created: order_id=ORD001, facility_code=VND-2404
📦 Order Created Event Received:
   - Order ID: ORD001
   - Facility Code: VND-2404
   ─────────────────────────────────
```

## 7. Stop Kafka

```bash
docker compose down
```
