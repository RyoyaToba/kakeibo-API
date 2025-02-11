CREATE TABLE IF NOT EXISTS category (
    user_id VARCHAR(10) NOT NULL,
    category_id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (user_id, category_id)
);
