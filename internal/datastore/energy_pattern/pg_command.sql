-- name: RecordEnergyPattern :one
INSERT INTO user_energy_patterns (user_id, date, time_slot_id, energy_level_id, notes)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (user_id, date, time_slot_id)
DO UPDATE SET
    energy_level_id = EXCLUDED.energy_level_id,
    notes = EXCLUDED.notes,
    created_at = CURRENT_TIMESTAMP
RETURNING id, user_id, date, time_slot_id, energy_level_id;

-- name: DeleteEnergyPattern :one
DELETE FROM user_energy_patterns
WHERE user_id = $1 AND date = $2 AND time_slot_id = $3
RETURNING id;
