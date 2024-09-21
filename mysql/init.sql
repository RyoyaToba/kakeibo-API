CREATE TABLE IF NOT EXISTS test (
    id INT AUTO_INCREMENT PRIMARY KEY,
    message VARCHAR(255)
);

INSERT INTO test (message) VALUES ('Hello world');
INSERT INTO test (message) VALUES ('Hello Golang');
