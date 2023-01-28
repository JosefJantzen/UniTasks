CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    e_mail STRING NOT NULL,
    pwd BYTEA NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS recurring_tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    name STRING NOT NULL, 
    description TEXT,
    start TIMESTAMP NOT NULL DEFAULT now(),
    ending TIMESTAMP,
    interval INT NOT NULL, 
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    user_id UUID REFERENCES users(id) NOT NULL
);

CREATE TABLE IF NOT EXISTS recurring_tasks_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    description TEXT,
    done BOOL DEFAULT false,
    done_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    user_id UUID REFERENCES users(id) NOT NULL,
    recurring_task_id UUID REFERENCES recurring_tasks(id) NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    name STRING NOT NULL, 
    description TEXT, 
    due TIMESTAMP NOT NULL,
    done BOOL DEFAULT false,
    done_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    user_id UUID REFERENCES users(id) NOT NULL
);