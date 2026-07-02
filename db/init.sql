CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL
);

INSERT INTO books(title, author)
VALUES
('Clean Code', 'Robert C. Martin'),
('The Pragmatic Programmer', 'Andrew Hunt');