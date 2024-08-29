CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    website VARCHAR(255),
    industry VARCHAR(255),
    size VARCHAR(50),
    founded INTEGER
);

-- Insert some sample data
INSERT INTO companies (name, description, website, industry, size, founded)
VALUES
('Flipkart', 'India''s leading e-commerce marketplace', 'https://www.flipkart.com', 'E-commerce', '10,000+ employees', 2007),
('PhonePe', 'Digital payments and financial services company', 'https://www.phonepe.com', 'Financial Technology', '5,000-10,000 employees', 2015);
