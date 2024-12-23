-- Delete existing data from the users table if the migration is rerun
DO $$
BEGIN
   IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users') THEN
      DELETE FROM users;
   END IF;
END$$;

-- Create ENUM types (only if they do not already exist)
DO $$
BEGIN
   IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'loan_status') THEN
      CREATE TYPE loan_status AS ENUM ('proposed', 'approved', 'invested', 'disbursed');
   END IF;
END$$;

DO $$
BEGIN
   IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_type') THEN
      CREATE TYPE user_type AS ENUM ('borrower', 'lender');
   END IF;
END$$;

-- Create users table
CREATE TABLE IF NOT EXISTS users (
   id bigserial PRIMARY KEY,
   full_name varchar(255) NOT NULL,
   email varchar(255) NOT NULL UNIQUE,
   type user_type NOT NULL
);

-- Create loans table
CREATE TABLE IF NOT EXISTS loans (
   id bigserial PRIMARY KEY,
   borrower_id int8 NOT NULL,
   principal_amount numeric(19,3) NOT NULL,
   rate numeric(5,2) NOT NULL,
   roi numeric(5,2) NOT NULL,
   agreement_letter text,
   status loan_status DEFAULT 'proposed' NOT NULL,
   field_validator_id varchar(255),
   visit_doc_proof text,
   approval_date timestamp,
   collector_id varchar(255),
   disbursement_date timestamp,
   CONSTRAINT fk_borrower FOREIGN KEY (borrower_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create marketplace table
CREATE TABLE IF NOT EXISTS marketplace (
   id bigserial PRIMARY KEY,
   loan_id int8 NOT NULL,
   lender_id int8 NOT NULL,
   funding_amount numeric(19,3) NOT NULL,
   CONSTRAINT fk_loan FOREIGN KEY (loan_id) REFERENCES loans(id) ON DELETE CASCADE,
   CONSTRAINT fk_lender FOREIGN KEY (lender_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_loans_borrower_id ON loans(borrower_id);
CREATE INDEX IF NOT EXISTS idx_marketplace_loan_id ON marketplace(loan_id);
CREATE INDEX IF NOT EXISTS idx_marketplace_lender_id ON marketplace(lender_id);

-- Insert sample users data
INSERT INTO users (full_name, email, type)
VALUES
   ('John Doe', 'mrwnolii+1@gmail.com', 'borrower'),
   ('Jane Smith', 'mrwnolii+2@gmail.com', 'lender'),
   ('Alice Johnson', 'mrwnolii+3@gmail.com', 'borrower'),
   ('Bob Brown', 'mrwnolii+4@gmail.com', 'lender'),
   ('Charlie Green', 'mrwnolii+5@gmail.com', 'borrower');
