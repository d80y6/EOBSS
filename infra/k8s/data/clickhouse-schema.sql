-- ClickHouse Schema for Telecom CDRs and Analytics
CREATE TABLE IF NOT EXISTS cdr (
    id String,
    usage_type Enum8('Data' = 1, 'Voice' = 2, 'SMS' = 3),
    quantity Float64,
    unit String,
    amount Float64,
    currency String,
    timestamp DateTime,
    service_id String,
    customer_id String,
    node_id String
) ENGINE = MergeTree()
PARTITION BY toYYYYMM(timestamp)
ORDER BY (customer_id, timestamp);

-- Materialized View for daily usage analytics
CREATE MATERIALIZED VIEW IF NOT EXISTS daily_usage_mv
ENGINE = SummingMergeTree()
PARTITION BY toYYYYMM(day)
ORDER BY (customer_id, usage_type, day)
AS SELECT
    customer_id,
    usage_type,
    toStartOfDay(timestamp) AS day,
    sum(quantity) AS total_quantity,
    sum(amount) AS total_amount
FROM cdr
GROUP BY customer_id, usage_type, day;
