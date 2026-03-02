CREATE TABLE employees (
   id UUID PRIMARY KEY default gen_random_uuid(),
   user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
   department_id UUID NOT NULL REFERENCES departments(id) ON DELETE CASCADE,
   position_id UUID NOT NULL REFERENCES positions(id) ON DELETE CASCADE,
   employee_code VARCHAR(255) NOT NULL UNIQUE,
   avatar_url TEXT,
   full_name VARCHAR(255) NOT NULL,
   gender VARCHAR(255),
   address VARCHAR(255),
   birth_date DATE,
   phone VARCHAR(255),
   is_active BOOLEAN DEFAULT TRUE,
   join_date TIMESTAMPTZ,
   leave_date TIMESTAMPTZ,
   status VARCHAR(255) NOT NULL DEFAULT 'active',
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ
)