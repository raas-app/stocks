CREATE TABLE stock_exchange (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  symbol VARCHAR(255) NOT NULL,
  country_id BIGINT,
  FOREIGN KEY (country_id) REFERENCES country(id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX idx_symbol ON stock_exchange(symbol);
CREATE INDEX idx_country_id ON stock_exchange(country_id);
