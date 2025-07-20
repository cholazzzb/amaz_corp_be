-- name: GetEnergyPatterns :one
SELECT uep.date, ts.slot_name, ts.start_time, ts.end_time,
       el.level_name, el.level_value, uep.notes
FROM user_energy_patterns uep
JOIN time_slots ts ON uep.time_slot_id = ts.id
JOIN energy_levels el ON uep.energy_level_id = el.id
WHERE uep.user_id = $1
  AND uep.date BETWEEN $2 AND $3
ORDER BY uep.date DESC, ts.sort_order;

-- GetAverageEnergy :one
SELECT ts.slot_name, ts.start_time, ts.end_time,
       AVG(el.level_value) as avg_energy,
       COUNT(*) as data_points
FROM user_energy_patterns uep
JOIN time_slots ts ON uep.time_slot_id = ts.id
JOIN energy_levels el ON uep.energy_level_id = el.id
WHERE uep.user_id = $1
  AND uep.date >= CURRENT_DATE - INTERVAL '30 days'
GROUP BY ts.id, ts.slot_name, ts.start_time, ts.end_time, ts.sort_order
ORDER BY ts.sort_order;

-- GetEnergyPatternByDate :one
SELECT ts.slot_name, el.level_name, el.level_value, uep.notes
FROM user_energy_patterns uep
JOIN time_slots ts ON uep.time_slot_id = ts.id
JOIN energy_levels el ON uep.energy_level_id = el.id
WHERE uep.user_id = $1 AND uep.date = $2
ORDER BY ts.sort_order;
