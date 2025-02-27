# Table: digitalocean_database_replicas

This table shows data for DigitalOcean Database Replicas.

https://docs.digitalocean.com/reference/api/api-reference/#operation/databases_list_replicas

The primary key for this table is **_cq_id**.

## Relations

This table depends on [digitalocean_databases](digitalocean_databases).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|id|String|
|name|String|
|connection|JSON|
|private_connection|JSON|
|region|String|
|status|String|
|created_at|Timestamp|
|private_network_uuid|String|
|tags|StringArray|