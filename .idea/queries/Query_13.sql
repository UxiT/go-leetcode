SELECT notification.uuid          AS "uuid",
       notification.customer_uuid AS "customer_uuid",
       notification.level         AS "level",
       notification.category_uuid AS "category_uuid",
       notification.channel_uuid  AS "channel_uuid",
       notification.payload       AS "payload",
       notification.message       AS "message",
       notification.viewed_at     AS "viewed_at",
       notification.sent_at       AS "sent_at",
       notification.created_at    AS "created_at",
       channels.uuid              AS "categories.uuid",
       channels.name              AS "categories.name",
       channels.created_at        AS "categories.created_at",
       channels.deleted_at        AS "categories.deleted_at",
       categories.uuid            AS "categories.uuid",
       categories.key             AS "categories.key",
       categories.parent_uuid     AS "categories.parent_uuid",
       categories.created_at      AS "categories.created_at"
FROM customer_notifications AS notification
         INNER JOIN channels ON notification.channel_uuid = channels.uuid
         INNER JOIN categories ON notification.category_uuid = categories.uuid
WHERE categories.key = 'registered'
  AND channels.name = 'email'
ORDER BY notification.created_at
LIMIT 4 OFFSET 0