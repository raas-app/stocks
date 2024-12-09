CREATE TABLE market_indices (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  country_id BIGINT,
  FOREIGN KEY (country_id) REFERENCES country(id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX idx_name ON market_indices(name);
CREATE INDEX idx_country_id ON market_indices(country_id);