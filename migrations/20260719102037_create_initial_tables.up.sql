CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
                       id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                       name text NOT NULL,
                       email text NOT NULL UNIQUE,
                       password_hash text NOT NULL,
                       created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE companies (
                           id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                           name text NOT NULL,
                           site_url text,
                           country text,
                           city text,
                           job_aggregator_url text,
                           created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE vacancies (
                           id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                           company_id uuid NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
                           title text NOT NULL,
                           start_salary int,
                           end_salary int,
                           currency text,
                           work_type text,
                           url text,
                           description text,
                           skills text,
                           public_date timestamptz,
                           created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE job_applications (
                                  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                                  user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                                  vacancy_id uuid NOT NULL REFERENCES vacancies(id) ON DELETE CASCADE,
                                  application_date timestamptz NOT NULL DEFAULT now(),
                                  status text NOT NULL,
                                  note text,
                                  recruiter_name text,
                                  recruiter_email text,
                                  recruiter_tg text,
                                  created_at timestamptz NOT NULL DEFAULT now()
);