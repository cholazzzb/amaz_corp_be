-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    timezone VARCHAR(50) DEFAULT 'UTC',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS energy_levels (
    id SERIAL PRIMARY KEY,
    level_name VARCHAR(20) UNIQUE NOT NULL, -- 'High', 'Medium', 'Low'
    level_value INTEGER NOT NULL, -- 3, 2, 1 for scoring
    description TEXT
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS time_slots (
    id SERIAL PRIMARY KEY,
    slot_name VARCHAR(50) UNIQUE NOT NULL, -- 'Morning', 'Afternoon', 'Evening', etc.
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    sort_order INTEGER NOT NULL
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS user_energy_patterns (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    time_slot_id INTEGER NOT NULL REFERENCES time_slots(id),
    energy_level_id INTEGER NOT NULL REFERENCES energy_levels(id),
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    -- Ensure one energy level per user per date per time slot
    UNIQUE(user_id, date, time_slot_id)
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS task_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    color_hex VARCHAR(7) -- For UI display
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    energy_demand_id INTEGER NOT NULL REFERENCES energy_levels(id),
    category_id INTEGER REFERENCES task_categories(id),
    estimated_duration INTEGER, -- in minutes
    priority INTEGER DEFAULT 1, -- 1-5 scale
    is_flexible BOOLEAN DEFAULT TRUE, -- can be rescheduled
    deadline DATE,
    is_completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS schedules (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    schedule_date DATE NOT NULL,
    generated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    ai_model_version VARCHAR(50), -- Track which AI model was used
    optimization_score DECIMAL(3,2), -- 0.00 to 1.00 confidence score
    raw_ai_response JSONB, -- Store full AI response for debugging
    is_active BOOLEAN DEFAULT TRUE, -- Allow multiple schedules per day
    user_rating INTEGER, -- 1-5 star rating from user
    user_feedback TEXT
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS schedule_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    schedule_id UUID NOT NULL REFERENCES schedules(id) ON DELETE CASCADE,
    task_id UUID REFERENCES tasks(id) ON DELETE SET NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    title VARCHAR(255) NOT NULL, -- Task title or custom activity
    description TEXT,
    item_type VARCHAR(50) DEFAULT 'task', -- 'task', 'break', 'buffer', 'focus_block'
    sort_order INTEGER NOT NULL,
    is_completed BOOLEAN DEFAULT FALSE,
    actual_start_time TIMESTAMP WITH TIME ZONE,
    actual_end_time TIMESTAMP WITH TIME ZONE
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS user_preferences (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    prefer_morning_focus BOOLEAN DEFAULT NULL,
    prefer_afternoon_focus BOOLEAN DEFAULT NULL,
    prefer_evening_focus BOOLEAN DEFAULT NULL,
    break_frequency INTEGER DEFAULT 60, -- minutes between breaks
    break_duration INTEGER DEFAULT 15, -- minutes per break
    work_start_time TIME DEFAULT '09:00:00',
    work_end_time TIME DEFAULT '17:00:00',
    buffer_time INTEGER DEFAULT 15, -- minutes buffer between tasks
    max_high_energy_blocks INTEGER DEFAULT 3, -- per day
    notifications_enabled BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(user_id)
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS schedule_feedback (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    schedule_id UUID NOT NULL REFERENCES schedules(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feedback_date DATE NOT NULL,
    energy_accuracy_rating INTEGER, -- 1-5 how accurate was energy prediction
    task_completion_rate DECIMAL(3,2), -- percentage of tasks completed
    overall_satisfaction INTEGER, -- 1-5 overall satisfaction
    suggestions TEXT, -- user improvement suggestions
    would_use_again BOOLEAN,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS api_usage_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    endpoint VARCHAR(255) NOT NULL,
    method VARCHAR(10) NOT NULL,
    status_code INTEGER,
    response_time_ms INTEGER,
    ai_tokens_used INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Up
CREATE INDEX idx_users_email ON users(email);
-- +migrate Up
CREATE INDEX idx_users_created_at ON users(created_at);
-- +migrate Up
CREATE INDEX idx_energy_patterns_user_date ON user_energy_patterns(user_id, date);
-- +migrate Up
CREATE INDEX idx_tasks_user_id ON tasks(user_id);
-- +migrate Up
CREATE INDEX idx_tasks_deadline ON tasks(deadline);
-- +migrate Up
CREATE INDEX idx_schedules_user_date ON schedules(user_id, schedule_date);
-- +migrate Up
CREATE INDEX idx_schedule_items_schedule_id ON schedule_items(schedule_id);
-- +migrate Up
CREATE INDEX idx_schedule_items_task_id ON schedule_items(task_id);
-- +migrate Up
CREATE INDEX idx_api_logs_user_created ON api_usage_logs(user_id, created_at);

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';
-- +migrate StatementEnd

-- Create triggers for updated_at timestamps
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_tasks_updated_at BEFORE UPDATE ON tasks
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_user_preferences_updated_at BEFORE UPDATE ON user_preferences
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Add some helpful views for common queries
CREATE VIEW user_energy_summary AS
SELECT
    uep.user_id,
    uep.date,
    ts.slot_name,
    el.level_name,
    el.level_value,
    uep.notes
FROM user_energy_patterns uep
JOIN time_slots ts ON uep.time_slot_id = ts.id
JOIN energy_levels el ON uep.energy_level_id = el.id
ORDER BY uep.date DESC, ts.sort_order;

CREATE VIEW task_summary AS
SELECT
    t.id,
    t.user_id,
    t.title,
    t.description,
    el.level_name as energy_demand,
    tc.name as category,
    t.estimated_duration,
    t.priority,
    t.deadline,
    t.is_completed,
    t.created_at
FROM tasks t
JOIN energy_levels el ON t.energy_demand_id = el.id
LEFT JOIN task_categories tc ON t.category_id = tc.id;

CREATE VIEW schedule_overview AS
SELECT
    s.id as schedule_id,
    s.user_id,
    s.schedule_date,
    s.generated_at,
    s.optimization_score,
    s.user_rating,
    COUNT(si.id) as total_items,
    COUNT(CASE WHEN si.is_completed = true THEN 1 END) as completed_items,
    CASE
        WHEN COUNT(si.id) > 0 THEN
            ROUND(COUNT(CASE WHEN si.is_completed = true THEN 1 END)::decimal / COUNT(si.id) * 100, 2)
        ELSE 0
    END as completion_percentage
FROM schedules s
LEFT JOIN schedule_items si ON s.id = si.schedule_id
WHERE s.is_active = true
GROUP BY s.id, s.user_id, s.schedule_date, s.generated_at, s.optimization_score, s.user_rating
ORDER BY s.schedule_date DESC;
