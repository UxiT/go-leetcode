WITH RECURSIVE category_chain AS (
    SELECT
        c.uuid,
        c.key,
        c.parent_uuid,
        c.created_at
    FROM notification_service.categories c
    WHERE c.uuid = 'bb5a286e-285e-4bd6-80fd-8653aac249e4'

    UNION ALL

    SELECT
        pc.uuid,
        pc.key,
        pc.parent_uuid,
        pc.created_at
    FROM notification_service.categories pc
             JOIN category_chain cc ON cc.parent_uuid = pc.uuid
)
SELECT * FROM category_chain;