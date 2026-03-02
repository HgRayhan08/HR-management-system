
CREATE TABLE payrolls (
   id UUID PRIMARY KEY default gen_random_uuid(),
   employee_id UUID NOT NULL REFERENCES employees(id) ON DELETE CASCADE,
   basic_salary VARCHAR(255) NOT NULL,
   piece VARCHAR(255) NOT NULL,
   allowance VARCHAR(255) NOT NULL,
   total_salary VARCHAR(255) NOT NULL,
   month VARCHAR(255) NOT NULL,
   year VARCHAR(255) NOT NULL,
   status VARCHAR(255) NOT NULL DEFAULT 'pending',
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ
)