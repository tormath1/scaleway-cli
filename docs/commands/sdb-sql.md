<!-- DO NOT EDIT: this file is automatically generated using scw-doc-gen -->
# Documentation for `scw sdb-sql`
This API allows you to manage your Serverless SQL DB databases.
  
- [](#)
  - [Export a database backup](#export-a-database-backup)
  - [Get a database backup information](#get-a-database-backup-information)
  - [List your Serverless SQL Database backups](#list-your-serverless-sql-database-backups)
- [](#)
  - [Create a new Serverless SQL Database](#create-a-new-serverless-sql-database)
  - [Delete a database](#delete-a-database)
  - [Get a database information](#get-a-database-information)
  - [List your Serverless SQL Databases](#list-your-serverless-sql-databases)
  - [Restore a database from a backup](#restore-a-database-from-a-backup)
  - [Update database information](#update-database-information)

  
## 




### Export a database backup

Export a database backup providing a download link once the export process is completed. You must provide the `backup_id` parameter.

**Usage:**

```
scw sdb-sql backup export <backup-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| backup-id | Required | UUID of the Serverless SQL Database backup. |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Get a database backup information

Retrieve information about your Serverless SQL Database backup. You must provide the `backup_id` parameter.

**Usage:**

```
scw sdb-sql backup get <backup-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| backup-id | Required | UUID of the Serverless SQL Database backup. |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### List your Serverless SQL Database backups

List all Serverless SQL Database backups for a given Scaleway Project or Database. By default, the backups returned in the list are ordered by creation date in descending order, though this can be modified via the order_by field.

**Usage:**

```
scw sdb-sql backup list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Filter by the UUID of the Scaleway project. |
| database-id | Required | Filter by the UUID of the Serverless SQL Database. |
| order-by | One of: `created_at_desc`, `created_at_asc` | Sorting criteria. One of `created_at_asc`, `created_at_desc`. |
| organization-id |  | Filter by the UUID of the Scaleway organization. |
| region | Default: `fr-par`<br />One of: `fr-par`, `all` | Region to target. If none is passed will use default region from the config |



## 




### Create a new Serverless SQL Database

You must provide the following parameters: `organization_id`, `project_id`, `name`, `cpu_min`, `cpu_max`. You can also provide `from_backup_id` to create a database from a backup.

**Usage:**

```
scw sdb-sql database create [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| name | Required | The name of the Serverless SQL Database to be created. |
| cpu-min | Required | The minimum number of CPU units for your Serverless SQL Database. |
| cpu-max | Required | The maximum number of CPU units for your Serverless SQL Database. |
| from-backup-id |  | The ID of the backup to create the database from. |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Delete a database

Deletes a database. You must provide the `database_id` parameter. All data stored in the database will be permanently deleted.

**Usage:**

```
scw sdb-sql database delete <database-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| database-id | Required | UUID of the Serverless SQL Database. |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Get a database information

Retrieve information about your Serverless SQL Database. You must provide the `database_id` parameter.

**Usage:**

```
scw sdb-sql database get <database-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| database-id | Required | UUID of the Serverless SQL DB database. |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### List your Serverless SQL Databases

List all Serverless SQL Databases for a given Scaleway Organization or Scaleway Project. By default, the databases returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. For the `name` parameter, the value you include will be checked against the whole name string to see if it includes the string you put in the parameter.

**Usage:**

```
scw sdb-sql database list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| name |  | Filter by the name of the database |
| order-by | One of: `created_at_asc`, `created_at_desc`, `name_asc`, `name_desc` | Sorting criteria. One of `created_at_asc`, `created_at_desc`, `name_asc`, `name_desc` |
| organization-id |  | Filter by the UUID of the Scaleway organization |
| region | Default: `fr-par`<br />One of: `fr-par`, `all` | Region to target. If none is passed will use default region from the config |



### Restore a database from a backup

Restore a database from a backup. You must provide the `backup_id` parameter.

**Usage:**

```
scw sdb-sql database restore [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| database-id | Required | UUID of the Serverless SQL Database. |
| backup-id | Required | UUID of the Serverless SQL Database backup to restore. |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Update database information

Update CPU limits of your Serverless SQL Database. You must provide the `database_id` parameter.

**Usage:**

```
scw sdb-sql database update <database-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| database-id | Required | UUID of the Serverless SQL Database. |
| cpu-min |  | The minimum number of CPU units for your Serverless SQL Database. |
| cpu-max |  | The maximum number of CPU units for your Serverless SQL Database. |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



