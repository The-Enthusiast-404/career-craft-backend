CREATE TABLE jobs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    company VARCHAR(255) NOT NULL,
    url TEXT,
    UNIQUE(company, title)
);

-- Create index
CREATE INDEX idx_company ON jobs(company);
