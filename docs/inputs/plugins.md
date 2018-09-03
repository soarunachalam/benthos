Input Plugins
=============

This document has been generated, do not edit it directly.

This document lists any input plugins that this flavour of Benthos offers beyond
the standard set.

### Contents

1. [`mysql`](#mysql)

## `mysql`

``` yaml
type: mysql
plugin:
  batch_size: 1
  buffer_timeout: 1s
  cache: ""
  consumer_id: 0
  databases:
  - my_db
  host: localhost
  key_prefix: ""
  latest: false
  mysqldump_path: ""
  password: ""
  prefetch_count: 0
  port: 3306
  sync_interval: 30s
  tables:
  - my_table
  username: ""
  view: NEW_AND_OLD_IMAGES
```

Streams a MySQL binary replication log as a series of JSON events. Each event
includes top level fields containing the binlog file name, event offset, schema
name, table name, event timestamp, event type, map of primary keys, and an
optional row summary. A sample event is shown below:

``` json
{
	"file": "binlog.000003",
	"keys": {
		"id": 1
	},
	"pos": 670,
	"row": {
		"after": {
			"created_at": "2018-08-23T21:32:05.839348Z",
			"id": 1,
			"title": "foo"
		},
		"before": {
			"created_at": "2018-08-23T21:32:05.839348Z",
			"id": 1,
			"title": "bar"
		}
	},
	"schema": "test",
	"table": "foo",
	"timestamp": "2018-08-23T15:32:05Z",
	"type":"update"
}
```

This input type requires a durable cache for periodically checkpointing its
position in the replication stream. The checkpointing frequency can be tuned
using the `sync_interval` config field. It is also recommended to
disable any cache TTL functionality.

This input supports both single (default) and batch message types. To enable
batch mode, set the `batch_size` config field to a value that is
greater than 1. When operating in batch mode, the buffering window can be
adjusted by tuning the `buffer_timeout` config field (default value is
`1s`).

### Row Summary

This input supports the following row summary view types that can be set
using the `view` config field.
- **KEYS_ONLY** - No row summary, primary key map only.
- **NEW_AND_OLD_IMAGES** - Includes an image of the record both before
	and after it was modified.
- **NEW_IMAGE** - Includes an image of the record after it was modified.
- **OLD_IMAGE** - Includes an image of the record before it was modified.

### Starting Position & Content filters

This input supports the following starting positions:
- **last synchronized**
	- The default starting position uses the last synchronized checkpoint.
	This position will take priority over all others. To disable, clear the
	cache for this consumer.
- **dump**
	- This method begins by dumping the current db(s) using `mysqldump`.
	This requires the `mysqldump` executable to be available on the
	host machine. To use this starting position, the `mysqldump_path`
	config field must contain the path/to/mysqldump,
- **latest**
	- This method will begin streaming at the latest binlog position. Enable
	this position using the `latest` config field.

A list of databases to subscribe to can be specified using the `databases`
config field. An optional table filter can be specified using the `tables`
config field as long as only a single database is specified.

### Metadata

This input adds the following metadata fields to each message:

```
- mysql_event_keys
- mysql_event_schema
- mysql_event_table
- mysql_log_position
- mysql_event_type
- mysql_file
- mysql_next_pos
- mysql_server_id
```

You can access these metadata fields using
[function interpolation](../config_interpolation.md#metadata).

