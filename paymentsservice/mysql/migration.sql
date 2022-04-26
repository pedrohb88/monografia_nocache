CREATE TABLE IF NOT EXISTS invoices (
    id INTEGER PRIMARY KEY AUTO_INCREMENT, 
    code VARCHAR(255) NOT NULL, 
    link VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS payments (
    id INTEGER PRIMARY KEY AUTO_INCREMENT, 
    amount FLOAT NOT NULL, 
    invoice_id INTEGER NULL, 
    FOREIGN KEY(invoice_id) REFERENCES invoices(id)
);

CREATE TABLE IF NOT EXISTS benchmark (
    test INTEGER, 
    resource ENUM('net', 'cpu', 'time'),
    x FLOAT, 
    y FLOAT
);

CREATE INDEX resource ON benchmark (resource);
