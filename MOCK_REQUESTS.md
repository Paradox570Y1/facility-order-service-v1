# Mock CRUD Requests for Facility & Order Service

Mock data for testing Facility-Order Service via Postman. Copy and paste the JSON into request bodies.

---

## 📋 Facility Service

### 1. POST /facilities - Create Facility

**Request Body:**

```json
{
  "code": "FAC-001",
  "name": "Central Distribution Hub",
  "address": "123 Main St, New York, NY 10001"
}
```

```json
{
  "code": "FAC-002",
  "name": "West Coast Warehouse",
  "address": "456 Pacific Ave, Los Angeles, CA 90001"
}
```

```json
{
  "code": "FAC-003",
  "name": "Midwest Processing Center",
  "address": "789 Industrial Blvd, Chicago, IL 60601"
}
```

**Expected Response (200 OK):**

```json
{
  "code": "FAC-001",
  "name": "Central Distribution Hub",
  "address": "123 Main St, New York, NY 10001"
}
```

---

### 2. GET /facilities - Read All Facilities

No request body needed

**Expected Response (200 OK):**

```json
[
  {
    "code": "FAC-001",
    "name": "Central Distribution Hub",
    "address": "123 Main St, New York, NY 10001"
  },
  {
    "code": "FAC-002",
    "name": "West Coast Warehouse",
    "address": "456 Pacific Ave, Los Angeles, CA 90001"
  },
  {
    "code": "FAC-003",
    "name": "Midwest Processing Center",
    "address": "789 Industrial Blvd, Chicago, IL 60601"
  }
]
```

---

### 3. GET /facilities/FAC-001 - Read Facility by Code

No request body needed

**Expected Response (200 OK):**

```json
{
  "code": "FAC-001",
  "name": "Central Distribution Hub",
  "address": "123 Main St, New York, NY 10001"
}
```

---

### 4. GET /facilities/FAC-999 - Read Non-existent Facility

No request body needed

**Expected Response (404 Not Found):**

```json
{
  "error": "facility not found"
}
```

---

## 📋 Order Service

### 1. POST /orders - Create Order

**Request Body:**

```json
{
  "id": "ORD-001",
  "facility_code": "FAC-001"
}
```

```json
{
  "id": "ORD-002",
  "facility_code": "FAC-002"
}
```

```json
{
  "id": "ORD-003",
  "facility_code": "FAC-003"
}
```

**Expected Response (200 OK):**

```json
{
  "id": "ORD-001",
  "facility_code": "FAC-001",
  "status": "PENDING",
  "created_at": "2026-03-15T10:30:45Z"
}
```

---

### 2. POST /orders - Create Order with Invalid Facility

**Request Body:**

```json
{
  "id": "ORD-004",
  "facility_code": "FAC-999"
}
```

**Expected Response (400 Bad Request):**

```json
{
  "error": "facility not found"
}
```

---

### 3. GET /orders - Read All Orders

No request body needed

**Expected Response (200 OK):**

```json
[
  {
    "id": "ORD-001",
    "facility_code": "FAC-001",
    "status": "PENDING",
    "created_at": "2026-03-15T10:30:45Z"
  },
  {
    "id": "ORD-002",
    "facility_code": "FAC-002",
    "status": "PENDING",
    "created_at": "2026-03-15T10:31:20Z"
  },
  {
    "id": "ORD-003",
    "facility_code": "FAC-003",
    "status": "PENDING",
    "created_at": "2026-03-15T10:32:10Z"
  }
]
```

---

### 4. GET /orders/ORD-001 - Read Order by ID

No request body needed

**Expected Response (200 OK):**

```json
{
  "id": "ORD-001",
  "facility_code": "FAC-001",
  "status": "PENDING",
  "created_at": "2026-03-15T10:30:45Z"
}
```

---

### 5. GET /orders/ORD-999 - Read Non-existent Order

No request body needed

**Expected Response (404 Not Found):**

```json
{
  "error": "order not found"
}
```

