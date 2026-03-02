CREATE TABLE positions (
   id UUID PRIMARY KEY default gen_random_uuid(),
   department_id UUID NOT NULL REFERENCES departments(id) ON DELETE CASCADE,
   name VARCHAR(255) NOT NULL,
   is_active BOOLEAN DEFAULT TRUE,
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)