CREATE TABLE notifications (
    id NOT NULL PRIMARY KEY,
    type TEXT NOT NULL,
    time DATE NOT NULL,
    message TEXT NOT NULL,
    is_view BOOLEAN NOT NULL
);