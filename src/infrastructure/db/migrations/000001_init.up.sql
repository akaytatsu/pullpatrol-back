CREATE TABLE if not exists users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(90) NOT NULL,
    email VARCHAR(120) NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_admin boolean DEFAULT false NOT NULL,
    git_name VARCHAR(70) DEFAULT '',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL
);

create table if not exists repositories (
    id BIGSERIAL PRIMARY KEY,
    repository VARCHAR(220) NOT NULL,
    active boolean DEFAULT true NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL
);

create table if not exists pullrequest (
    id BIGSERIAL PRIMARY KEY,
    number BIGINT NOT NULL,
    action VARCHAR(40) NOT NULL,
    repository_id BIGINT NOT NULL,
    status VARCHAR(40) NOT NULL,
    url VARCHAR(220) NOT NULL,
    title VARCHAR(220) NOT NULL,
    closed_at TIMESTAMP,
    additions int NOT NULL,
    deletions int NOT NULL,
    changed_files int NOT NULL,
    commits int NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL
);

create table if not exists groups (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(90) NOT NULL,
    description text NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL
);

create table if not exists group_user (
    id BIGSERIAL PRIMARY KEY,
    group_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL
);

create table if not exists pullrequest_role (
    id BIGSERIAL PRIMARY KEY,
    pullrequest_id BIGINT NOT NULL,
    role_type VARCHAR(40) NOT NULL,
    description text NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL
);

create table if not exists pullrequest_review (
    id BIGSERIAL PRIMARY KEY,
    pullrequest_id BIGINT NOT NULL,
    pullrequest_role_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    status VARCHAR(40) NOT NULL,
    comment text NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL
);
