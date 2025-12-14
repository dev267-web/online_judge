-- USERS
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE,
    password_hash TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- LANGUAGES
CREATE TABLE languages (
    id SMALLSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    slug VARCHAR(20) UNIQUE NOT NULL, -- c, cpp, python
    file_ext VARCHAR(10) NOT NULL,
    time_limit_ms INT NOT NULL DEFAULT 2000,
    memory_limit_kb INT NOT NULL DEFAULT 65536,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

-- PROBLEMS
CREATE TABLE problems (
    id BIGSERIAL PRIMARY KEY,
    code VARCHAR(20) UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    statement TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- TEST CASES
CREATE TABLE test_cases (
    id BIGSERIAL PRIMARY KEY,
    problem_id BIGINT NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    input TEXT NOT NULL,
    expected_output TEXT NOT NULL,
    score INT NOT NULL DEFAULT 1,
    is_sample BOOLEAN NOT NULL DEFAULT FALSE
);

-- SUBMISSIONS
CREATE TABLE submissions (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(id),
    problem_id BIGINT NOT NULL REFERENCES problems(id),
    language_id SMALLINT NOT NULL REFERENCES languages(id),
    source_code TEXT NOT NULL,
    status VARCHAR(30) NOT NULL DEFAULT 'queued',
    time_ms INT,
    memory_kb INT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    finished_at TIMESTAMPTZ
);

-- PER TEST CASE RESULT
CREATE TABLE submission_runs (
    id BIGSERIAL PRIMARY KEY,
    submission_id BIGINT NOT NULL REFERENCES submissions(id) ON DELETE CASCADE,
    test_case_id BIGINT NOT NULL REFERENCES test_cases(id),
    status VARCHAR(30) NOT NULL, -- accepted, wa, tle, re
    time_ms INT,
    memory_kb INT
);

-- INDEXES (IMPORTANT)
CREATE INDEX idx_submissions_user ON submissions(user_id);
CREATE INDEX idx_submissions_problem ON submissions(problem_id);
CREATE INDEX idx_submissions_status ON submissions(status);
CREATE INDEX idx_runs_submission ON submission_runs(submission_id);
