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
SERVER_PORT=8080
```

## 4. Run Server
```bash
go run ./cmd/api
```

Server starts at `http://localhost:8080`