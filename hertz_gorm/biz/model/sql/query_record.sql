CREATE TABLE query_records (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    word TEXT NOT NULL,
    ip_address TEXT,
    user_agent TEXT
);