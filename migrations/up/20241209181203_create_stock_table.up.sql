CREATE TABLE stock (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    symbol VARCHAR(255) NOT NULL,
    market_cap DECIMAL(10, 2) DEFAULT 0,
    shares DECIMAL(10, 2) DEFAULT 0,
    free_float DECIMAL(10, 2) DEFAULT 0,
    book_value DECIMAL(10, 2) DEFAULT 0,
    face_value DECIMAL(10, 2) DEFAULT 0, 
    sector_id BIGINT,
    company_profile JSON,
    FOREIGN KEY (sector_id) REFERENCES sector(id) ON DELETE CASCADE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE (symbol)
);

CREATE INDEX idx_symbol ON stock(symbol);
CREATE INDEX idx_sector_id ON stock(sector_id);