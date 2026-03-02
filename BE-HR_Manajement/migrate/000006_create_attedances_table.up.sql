CREATE TABLE attendances (
   id UUID PRIMARY KEY default gen_random_uuid(),
   employee_id UUID NOT NULL REFERENCES employees(id) ON DELETE CASCADE,
   check_in TIMESTAMPTZ NOT NULL,
   check_out TIMESTAMPTZ,
   date DATE NOT NULL,
   status VARCHAR(255) NOT NULL DEFAULT 'working',
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ
)