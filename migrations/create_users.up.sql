CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(50) NOT NULL UNIQUE,
	encrypted_password VARCHAR(255) NOT NULL
);