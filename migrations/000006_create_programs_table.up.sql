CREATE TABLE programs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    faculty_id UUID NOT NULL REFERENCES faculties(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    degree TEXT NOT NULL,
    duration_years NUMERIC(3,1),
    description TEXT,
    tuition_fee NUMERIC(12,2),
    is_published BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
