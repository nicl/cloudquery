# Table: digitalocean_balances

This table shows data for DigitalOcean Balances.

https://docs.digitalocean.com/reference/api/api-reference/#operation/balance_get

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|month_to_date_balance|String|
|account_balance|String|
|month_to_date_usage|String|
|generated_at|Timestamp|