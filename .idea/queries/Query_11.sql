WITH RECURSIVE
    category_chain AS (
        SELECT
            uuid,
            parent_uuid,
            key,
            0 AS depth
        FROM notification_service.categories
        WHERE key ILIKE 'kyc_submitted'
        UNION ALL
        SELECT c.uuid, c.parent_uuid, c.key, cc.depth + 1
        FROM notification_service.categories c
            JOIN category_chain cc ON cc.parent_uuid = c.uuid
    ),
    disabled_categories AS (
        SELECT
            s.channel_uuid,
            cat.uuid AS category_uuid,
            cat.depth
        FROM notification_service.customer_settings s
            JOIN category_chain cat ON s.category_uuid = cat.uuid
        WHERE (s.customer_uuid = 'a8223e81-a3d1-4378-8863-e50f0806bf99' or s.customer_uuid is null) AND s.status = 'disabled'
    ),
    ranked_disabled AS (
        SELECT
            channel_uuid,
            category_uuid,
            depth,
            row_number() OVER (PARTITION BY channel_uuid ORDER BY depth) AS rnk
        FROM disabled_categories
    )
SELECT ch.*
FROM notification_service.channels ch
    LEFT JOIN ranked_disabled rd ON ch.uuid = rd.channel_uuid AND rd.rnk = 1
WHERE rd.category_uuid IS NULL