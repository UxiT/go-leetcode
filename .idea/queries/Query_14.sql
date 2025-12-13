INSERT INTO notification_service.categories (uuid, key, parent_uuid, created_at)
values (gen_random_uuid(), 'SUBSCRIPTION_EXPIRED', null, now()),
       (gen_random_uuid(), 'SUBSCRIPTION_IS_EXPIRING_IN_1_DAY', null, now()),
       (gen_random_uuid(), 'SUBSCRIPTION_IS_EXPIRING_IN_4_DAYS', null, now()),
       (gen_random_uuid(), 'SUBSCRIPTION_IS_EXPIRING_IN_7_DAYS', null, now());