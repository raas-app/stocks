CREATE TABLE stock_price_history (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,         -- Unique identifier for each price record
    stock_id BIGINT,                              -- Foreign key to the stock table
    price DECIMAL(10, 2) NOT NULL,              -- The stock price
    volume BIGINT NOT NULL,                     -- Trading volume (number of shares traded)
    open_price DECIMAL(10, 2),                  -- Opening price of the stock
    high_price DECIMAL(10, 2),                  -- Highest price during the day
    low_price DECIMAL(10, 2),                   -- Lowest price during the day
    close_price DECIMAL(10, 2),                 -- Closing price of the stock
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of the price record
    UNIQUE (stock_id, timestamp),              -- Ensure there are no duplicate entries for a stock at the same time
    FOREIGN KEY (stock_id) REFERENCES stock(id) ON DELETE CASCADE, -- Foreign key constraint
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX idx_stock_id ON stock_price_history(stock_id);