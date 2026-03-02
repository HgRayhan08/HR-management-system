CREATE TABLE departments (
   id UUID PRIMARY KEY default gen_random_uuid(),
   name VARCHAR(255) NOT NULL,
   description VARCHAR(255)
)