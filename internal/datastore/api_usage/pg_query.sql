-- name: GetUserApiUsageStatistics :one
SELECT
    DATE_TRUNC('day', created_at) as usage_date,
    COUNT(*) as total_requests,
    COUNT(CASE WHEN status_code = 200 THEN 1 END) as successful_requests,
    AVG(response_time_ms) as avg_response_time,
    SUM(ai_tokens_used) as total_tokens
FROM api_usage_logs
WHERE user_id = $1
  AND created_at >= CURRENT_DATE - INTERVAL '30 days'
GROUP BY DATE_TRUNC('day', created_at)
ORDER BY usage_date DESC;
