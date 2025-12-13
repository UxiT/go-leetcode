SELECT t.*, channel.*
FROM notification_service.customer_settings as t
    JOIN notification_service.channels channel on t.channel_uuid = channel.uuid
WHERE t.customer_uuid = 'a8223e81-a3d1-4378-8863-e50f0806bf99' AND t.category_uuid = 'dfcf1486-e627-46ca-96fe-f459ccbc785e'