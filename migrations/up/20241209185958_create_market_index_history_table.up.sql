CREATE TABLE market_index_history (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    index_id BIGINT NOT NULL,  -- Foreign key to indices table
    date DATE NOT NULL,        -- Date of the historical data
    open_price DECIMAL(10, 2), -- Opening price of the index
    close_price DECIMAL(10, 2),-- Closing price of the index
    high_price DECIMAL(10, 2), -- Highest price during the day
    low_price DECIMAL(10, 2),  -- Lowest price during the day
    volume BIGINT,             -- Trading volume
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uq_index_date (index_id, date)  -- Ensure one record per index per date
);
