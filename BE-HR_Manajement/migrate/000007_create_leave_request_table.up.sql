


CREATE TABLE leave_requests (
   id UUID PRIMARY KEY default gen_random_uuid(),
   employee_id UUID NOT NULL REFERENCES employees(id) ON DELETE CASCADE,
   start_date DATE NOT NULL,
   end_date DATE NOT NULL,
   type VARCHAR(255) NOT NULL,
   reason VARCHAR(255),
   proof_url TEXT,
   status VARCHAR(255) NOT NULL DEFAULT 'pending',
   approver_id UUID REFERENCES employees(id) ON DELETE CASCADE,
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ
)