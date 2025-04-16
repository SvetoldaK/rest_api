ALTER TABLE tasks ADD COLUMN user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE;
CREATE INDEX idx_tasks_user_id ON tasks(user_id); 