CREATE TABLE orders (
    id VARCHAR(50) PRIMARY KEY,
    facility_code VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (facility_code) REFERENCES facilities(code)
);
