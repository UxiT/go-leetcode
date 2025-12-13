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
       channels.uuid              AS "channels.uuid",
       channels.name              AS "channels.name",
       channels.created_at        AS "channels.created_at",
       channels.deleted_at        AS "channels.deleted_at",
       categories.uuid            AS "categories.uuid",
       categories.key             AS "categories.key",
       categories.parent_uuid     AS "categories.parent_uuid",
       categories.created_at      AS "categories.created_at"
FROM customer_notifications AS notification
         INNER JOIN channels ON notification.channel_uuid = channels.uuid
         INNER JOIN categories ON notification.category_uuid = categories.uuid
WHERE categories.key = $1
  AND channels.name = $2
ORDER BY notification.created_at ASC
LIMIT $3 OFFSET $4