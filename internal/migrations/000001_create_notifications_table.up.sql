CREATE TABLE notifications (
    id INTEGER NOT NULL PRIMARY KEY,
    type TEXT NOT NULL,
    time DATE NOT NULL,
    message TEXT NOT NULL
);