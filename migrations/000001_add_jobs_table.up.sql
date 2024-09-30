-- 000001_create_jobs_table.up.sql

CREATE TABLE IF NOT EXISTS jobs (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    company TEXT NOT NULL,
    location TEXT,
    salary TEXT,
    role TEXT,
    skills TEXT,
    remote BOOLEAN DEFAULT false,
    experience TEXT,
    education TEXT,
    department TEXT,
    job_type TEXT,
    url TEXT,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(company, title)
);

CREATE INDEX idx_company ON jobs(company);
