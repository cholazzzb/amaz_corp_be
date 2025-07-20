-- +migrate Up
-- Insert default energy levels
INSERT INTO energy_levels (level_name, level_value, description) VALUES
('High', 3, 'Peak energy and focus - ideal for demanding cognitive tasks'),
('Medium', 2, 'Moderate energy - good for routine tasks and planning'),
('Low', 1, 'Low energy - best for light tasks, breaks, and reflection');

-- Insert default time slots
INSERT INTO time_slots (slot_name, start_time, end_time, sort_order) VALUES
('Early Morning', '06:00:00', '09:00:00', 1),
('Morning', '09:00:00', '12:00:00', 2),
('Afternoon', '12:00:00', '15:00:00', 3),
('Late Afternoon', '15:00:00', '18:00:00', 4),
('Evening', '18:00:00', '21:00:00', 5),
('Night', '21:00:00', '23:59:59', 6);

-- Insert default task categories
INSERT INTO task_categories (name, description, color_hex) VALUES
('Work', 'Professional tasks and meetings', '#3B82F6'),
('Personal', 'Personal errands and activities', '#10B981'),
('Learning', 'Education and skill development', '#F59E0B'),
('Health', 'Exercise and wellness activities', '#EF4444'),
('Creative', 'Creative projects and hobbies', '#8B5CF6'),
('Administrative', 'Paperwork and routine tasks', '#6B7280');
