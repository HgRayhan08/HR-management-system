CREATE TABLE users (
   id UUID PRIMARY KEY default gen_random_uuid(),
   email VARCHAR(255) NOT NULL UNIQUE,
   password VARCHAR(255) NOT NULL,
   role_id UUID NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
   is_active BOOLEAN DEFAULT TRUE,
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ
)
