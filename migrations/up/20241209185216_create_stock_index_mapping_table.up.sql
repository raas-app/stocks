CREATE TABLE stock_index_mapping (
    stock_id BIGINT NOT NULL,
    index_id BIGINT NOT NULL,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (stock_id, index_id),
    FOREIGN KEY (stock_id) REFERENCES stock(id) ON DELETE CASCADE,
    FOREIGN KEY (index_id) REFERENCES market_indices(id) ON DELETE CASCADE
);

CREATE INDEX idx_index_id ON stock_index_mapping(index_id);