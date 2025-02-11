CREATE TABLE IF NOT EXISTS bank_account (
    user_id VARCHAR(10) NOT NULL,
    account_id INT NOT NULL AUTO_INCREMENT,
    type INT,
    name VARCHAR(255),
    balance INT,
    created_by VARCHAR(255),
    created_at TIMESTAMP NULL DEFAULT NULL,
    updated_by VARCHAR(255),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (user_id, account_id)
);
