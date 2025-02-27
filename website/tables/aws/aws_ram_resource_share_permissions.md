# Table: aws_ram_resource_share_permissions

This table shows data for RAM Resource Share Permissions.

https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceSharePermissionSummary.html

The composite primary key for this table is (**account_id**, **region**, **resource_share_arn**, **arn**, **version**).

## Relations

This table depends on [aws_ram_resource_shares](aws_ram_resource_shares).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|resource_share_arn (PK)|String|
|permission|JSON|
|arn (PK)|String|
|creation_time|Timestamp|
|default_version|Bool|
|is_resource_type_default|Bool|
|last_updated_time|Timestamp|
|name|String|
|resource_type|String|
|status|String|
|version (PK)|String|