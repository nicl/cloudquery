# Table: k8s_core_replication_controllers

This table shows data for Kubernetes (K8s) Core Replication Controllers.

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|context|String|
|kind|String|
|api_version|String|
|name|String|
|namespace|String|
|uid (PK)|String|
|resource_version|String|
|generation|Int|
|deletion_grace_period_seconds|Int|
|labels|JSON|
|annotations|JSON|
|owner_references|JSON|
|finalizers|StringArray|
|spec_replicas|Int|
|spec_min_ready_seconds|Int|
|spec_selector|JSON|
|spec_template|JSON|
|status_replicas|Int|
|status_fully_labeled_replicas|Int|
|status_ready_replicas|Int|
|status_available_replicas|Int|
|status_observed_generation|Int|
|status_conditions|JSON|