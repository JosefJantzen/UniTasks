CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    e_mail STRING NOT NULL,
    pwd BYTEA NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS recurring_tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    name STRING NOT NULL, 
    description TEXT,
    start TIMESTAMP NOT NULL,
    end TIMESTAMP NOT NULL,
    interval INT NOT NULL, 
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    user_id UUID REFERENCES users(id) 
);

/*CREATE TABLE IF NOT EXISTS recurring_tasks_history ()*/

CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    name STRING NOT NULL, 
    description TEXT, 
    due TIMESTAMP NOT NULL,
    done BOOL DEFAULT false,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    user_id UUID REFERENCES users(id),
);

INSERT INTO users (id, e_mail, pwd) 
VALUES (
    'acde070d-8c4c-4f0d-9d8a-162843c10333', 
    'admin@admin.com', 
    '$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK'
); 

INSERT INTO recurring_tasks (id, name, description, start, interval, user_id)
VALUES (
    'acde070d-8c4c-4f0d-9d8a-162843c10444', 
    'Test-Recurring-Task', 
    'This is a recurring task.', 
    '2022-12-01 15:15:35',
    7, 
    'acde070d-8c4c-4f0d-9d8a-162843c10333'
);

INSERT INTO tasks (id, name, description, due, user_id)
VALUES (
    'acde070d-8c4c-4f0d-9d8a-162843c10555', 
    'Test-task', 
    'This is a normal task description', 
    '2023-01-01 15:15:35', 
    'acde070d-8c4c-4f0d-9d8a-162843c10333'
);