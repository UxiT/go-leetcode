SELECT customer_notifications.uuid,
       customer_notifications.customer_uuid,
       customer_notifications.level,
       customer_notifications.category_uuid,
       customer_notifications.channel_uuid,
       customer_notifications.payload,
       customer_notifications.message,
       customer_notifications.viewed_at,
       customer_notifications.sent_at,
       customer_notifications.created_at,
       channels.uuid,
       channels.name,
       channels.created_at,
       channels.deleted_at,
       categories.uuid,
       categories.key,
       categories.parent_uuid,
       categories.created_at
FROM customer_notifications AS notification
         INNER JOIN channels ON notification.channel_uuid = channels.uuid
         INNER JOIN categories ON notification.category_uuid = categories.uuid
WHERE notification.customer_uuid = $1
  AND notification.viewed_at IS NULL
ORDER BY t.created_at DESC
LIMIT $2 OFFSET $3