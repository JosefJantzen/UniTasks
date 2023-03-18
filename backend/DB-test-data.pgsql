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
    '2023-03-01 15:15:35',
    '2023-06-01 15:15:35', 
    7, 
    'acde070d-8c4c-4f0d-9d8a-162843c10333'
) ON CONFLICT DO NOTHING;

INSERT INTO recurring_tasks_history (id, description, due, done, done_at, user_id, recurring_task_id)
VALUES (
    'acde070d-8c4c-4f0d-9d8a-162843c10455', 
    'This is a recurring task history entry.', 
    '2023-03-15 23:59:59',
    true, 
    '2022-03-13 15:15:35',
    'acde070d-8c4c-4f0d-9d8a-162843c10333',
    'acde070d-8c4c-4f0d-9d8a-162843c10444'
) ON CONFLICT DO NOTHING;

INSERT INTO recurring_tasks_history (id, description, due, user_id, recurring_task_id)
VALUES (
    'acde070d-8c4c-4f0d-9d8a-162843c10456', 
    'This is a recurring task history entry not done', 
    '2023-03-22 23:59:59',
    'acde070d-8c4c-4f0d-9d8a-162843c10333',
    'acde070d-8c4c-4f0d-9d8a-162843c10444'
) ON CONFLICT DO NOTHING;

INSERT INTO tasks (id, name, description, due, user_id)
VALUES (
    'acde070d-8c4c-4f0d-9d8a-162843c10555', 
    'Test-task1', 
    'This is a normal task description', 
    '2023-05-13 15:15:35', 
    'acde070d-8c4c-4f0d-9d8a-162843c10333'
) ON CONFLICT DO NOTHING;

INSERT INTO tasks (id, name, description, due, done, done_at, user_id)
VALUES (
    'acde070d-8c4c-4f0d-9d8a-162843c10556', 
    'Test-task2', 
    'This is a normal task description2', 
    '2023-01-01 15:15:35', 
    true,
    '2022-12-31 12:45:43',
    'acde070d-8c4c-4f0d-9d8a-162843c10333'
) ON CONFLICT DO NOTHING;