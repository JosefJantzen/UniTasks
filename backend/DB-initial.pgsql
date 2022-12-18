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
    done_at TIMESTAMP NOT NULL DEFAULT now(),
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
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    user_id UUID REFERENCES users(id) NOT NULL
);

INSERT INTO users (id, e_mail, pwd) 
VALUES (
    'acde070d-8c4c-4f0d-9d8a-162843c10333', 
    'admin@admin.com', 
    '$2a$14$do7vbh0rgDNJCi1Uvz60zOE6dannxKht9yRG6bAXM4NFIrSHEKEZ.' /* secret */
) ON CONFLICT DO NOTHING;

INSERT INTO recurring_tasks (id, name, description, start, ending, interval, user_id)
VALUES (
    'acde070d-8c4c-4f0d-9d8a-162843c10444', 
    'Test-Recurring-Task', 
    'This is a recurring task.', 
    '2022-12-01 15:15:35',
    '2023-01-01 15:15:35', 
    7, 
    'acde070d-8c4c-4f0d-9d8a-162843c10333'
) ON CONFLICT DO NOTHING;

INSERT INTO recurring_tasks_history (id, description, done, done_at, user_id, recurring_task_id)
VALUES (
    'acde070d-8c4c-4f0d-9d8a-162843c10455', 
    'This is a recurring task history entry.', 
    true, 
    '2022-12-05 15:15:35',
    'acde070d-8c4c-4f0d-9d8a-162843c10333',
    'acde070d-8c4c-4f0d-9d8a-162843c10444'
) ON CONFLICT DO NOTHING;

INSERT INTO recurring_tasks_history (id, description, user_id, recurring_task_id)
VALUES (
    'acde070d-8c4c-4f0d-9d8a-162843c10456', 
    'This is a recurring task history entry not done', 
    'acde070d-8c4c-4f0d-9d8a-162843c10333',
    'acde070d-8c4c-4f0d-9d8a-162843c10444'
) ON CONFLICT DO NOTHING;

INSERT INTO tasks (id, name, description, due, user_id)
VALUES (
    'acde070d-8c4c-4f0d-9d8a-162843c10555', 
    'Test-task', 
    'This is a normal task description', 
    '2023-01-01 15:15:35', 
    'acde070d-8c4c-4f0d-9d8a-162843c10333'
) ON CONFLICT DO NOTHING;