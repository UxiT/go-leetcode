INSERT INTO notification_service.channels (uuid, name, created_at, deleted_at) VALUES ('28997dea-c7e8-4fcf-9e5b-a9ad1c1c0385', 'email', '2025-10-10 12:20:56', null);
INSERT INTO notification_service.channels (uuid, name, created_at, deleted_at) VALUES ('d992fc2e-f6dd-4200-a570-2c5c4ea99a2a', 'telegram', '2025-10-10 12:21:02', null);
INSERT INTO notification_service.channels (uuid, name, created_at, deleted_at) VALUES ('60c54def-4c26-42a2-804e-9a48562ea7ff', 'push', '2025-10-10 12:21:06', null);

INSERT INTO notification_service.categories (uuid, key, parent_uuid, created_at) VALUES ('dfcf1486-e627-46ca-96fe-f459ccbc785e', 'REGISTERED', null, '2025-10-10 12:21:33');

INSERT INTO notification_service.customer_settings (customer_uuid, category_uuid, channel_uuid, status, created_at) VALUES ('a8223e81-a3d1-4378-8863-e50f0806bf99', 'dfcf1486-e627-46ca-96fe-f459ccbc785e', '28997dea-c7e8-4fcf-9e5b-a9ad1c1c0385', 'disabled', '2025-10-10 12:25:34');
INSERT INTO notification_service.customer_settings (customer_uuid, category_uuid, channel_uuid, status, created_at) VALUES ('a8223e81-a3d1-4378-8863-e50f0806bf99', 'dfcf1486-e627-46ca-96fe-f459ccbc785e', 'd992fc2e-f6dd-4200-a570-2c5c4ea99a2a', 'disabled', '2025-10-10 12:25:34');

