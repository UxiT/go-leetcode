SET search_path TO notification_service;

SELECT notification.uuid          AS "notification.uuid",
       notification.customer_uuid AS "notification.customer_uuid",
       notification.level         AS "notification.level",
       notification.category_uuid AS "notification.category_uuid",
       notification.channel_uuid  AS "notification.channel_uuid",
       notification.payload       AS "notification.payload",
       notification.message       AS "notification.message",
       notification.viewed_at     AS "notification.viewed_at",
       notification.sent_at       AS "notification.sent_at",
       notification.created_at    AS "notification.created_at",
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
WHERE notification.uuid = 'd240aef2-293c-4342-a199-8a217425cf37'
LIMIT 1