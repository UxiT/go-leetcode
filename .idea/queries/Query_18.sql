select ch.name, cs.status, c.key from notification_service.customer_settings cs
join notification_service.categories c on cs.category_uuid = c.uuid
join notification_service.channels ch on cs.channel_uuid = ch.uuid
where ch.name != 'telegram' and cs.customer_uuid is null;

select * from customer_service.customers
where uuid = '7db780fa-fc5f-4ca3-9ad7-1e6748f4a08a';

select * from notification_service.customer_notifications
where customer_uuid = '7db780fa-fc5f-4ca3-9ad7-1e6748f4a08a'
order by created_at desc;