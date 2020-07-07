// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package rdb

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/internal/core"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		rdbRoot(),
		rdbBackup(),
		rdbEngine(),
		rdbInstance(),
		rdbACL(),
		rdbSettings(),
		rdbPrivilege(),
		rdbUser(),
		rdbDatabase(),
		rdbNodeType(),
		rdbLog(),
		rdbCertificate(),
		rdbBackupList(),
		rdbBackupCreate(),
		rdbBackupGet(),
		rdbBackupUpdate(),
		rdbBackupDelete(),
		rdbBackupRestore(),
		rdbBackupExport(),
		rdbInstanceClone(),
		rdbEngineList(),
		rdbInstanceUpgrade(),
		rdbInstanceList(),
		rdbInstanceGet(),
		rdbInstanceCreate(),
		rdbInstanceUpdate(),
		rdbInstanceDelete(),
		rdbCertificateGet(),
		rdbLogPrepare(),
		rdbLogList(),
		rdbInstanceMetrics(),
		rdbSettingsAdd(),
		rdbSettingsDelete(),
		rdbACLList(),
		rdbACLAdd(),
		rdbACLDelete(),
		rdbUserList(),
		rdbUserCreate(),
		rdbUserUpdate(),
		rdbUserDelete(),
		rdbDatabaseList(),
		rdbDatabaseCreate(),
		rdbDatabaseDelete(),
		rdbPrivilegeList(),
		rdbPrivilegeSet(),
		rdbNodeTypeList(),
	)
}
func rdbRoot() *core.Command {
	return &core.Command{
		Short:     `Database RDB API`,
		Long:      ``,
		Namespace: "rdb",
	}
}

func rdbBackup() *core.Command {
	return &core.Command{
		Short: ``,
		Long: `Save and restore backups of your database instance.
`,
		Namespace: "rdb",
		Resource:  "backup",
	}
}

func rdbEngine() *core.Command {
	return &core.Command{
		Short: ``,
		Long: `Software that stores and retrieves data from a database. Each database engine has a name and versions.
`,
		Namespace: "rdb",
		Resource:  "engine",
	}
}

func rdbInstance() *core.Command {
	return &core.Command{
		Short: ``,
		Long: `A Database Instance is composed of one or more Nodes, depending of the is_ha_cluster setting. Autohealing is enabled by default for HA clusters. Database automated backup is enabled by default in order to prevent data loss.
`,
		Namespace: "rdb",
		Resource:  "instance",
	}
}

func rdbACL() *core.Command {
	return &core.Command{
		Short: ``,
		Long: `Network Access Control List allows to control network in and out traffic by setting up ACL rules. ACL rules could be created, edited, deleted
`,
		Namespace: "rdb",
		Resource:  "acl",
	}
}

func rdbSettings() *core.Command {
	return &core.Command{
		Short: ``,
		Long: `Instance Settings are tunables of Database Engines. Available settings depend on the database engine and its version.
`,
		Namespace: "rdb",
		Resource:  "settings",
	}
}

func rdbPrivilege() *core.Command {
	return &core.Command{
		Short: ``,
		Long: `Define some privileges to a user on a specific database.
`,
		Namespace: "rdb",
		Resource:  "privilege",
	}
}

func rdbUser() *core.Command {
	return &core.Command{
		Short: ``,
		Long: `Manage users on your instance
`,
		Namespace: "rdb",
		Resource:  "user",
	}
}

func rdbDatabase() *core.Command {
	return &core.Command{
		Short: ``,
		Long: `Manage logical databases on your instance
`,
		Namespace: "rdb",
		Resource:  "database",
	}
}

func rdbNodeType() *core.Command {
	return &core.Command{
		Short: ``,
		Long: `Node types powering your instance
`,
		Namespace: "rdb",
		Resource:  "node-type",
	}
}

func rdbLog() *core.Command {
	return &core.Command{
		Short:     ``,
		Long:      ``,
		Namespace: "rdb",
		Resource:  "log",
	}
}

func rdbCertificate() *core.Command {
	return &core.Command{
		Short:     ``,
		Long:      ``,
		Namespace: "rdb",
		Resource:  "certificate",
	}
}

func rdbBackupList() *core.Command {
	return &core.Command{
		Short:     `List database backups`,
		Long:      `List database backups.`,
		Namespace: "rdb",
		Resource:  "backup",
		Verb:      "list",
		ArgsType:  reflect.TypeOf(rdb.ListDatabaseBackupsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the database backups`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Criteria to use when ordering database backups listing`,
				Required:   false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc", "status_asc", "status_desc"},
			},
			{
				Name:       "instance-id",
				Short:      `UUID of the instance`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Organization ID the database backups belongs to`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListDatabaseBackupsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			resp, err := api.ListDatabaseBackups(request, scw.WithAllPages())
			if err != nil {
				return nil, err
			}
			return resp.DatabaseBackups, nil

		},
	}
}

func rdbBackupCreate() *core.Command {
	return &core.Command{
		Short:     `Create a database backup`,
		Long:      `Create a database backup.`,
		Namespace: "rdb",
		Resource:  "backup",
		Verb:      "create",
		ArgsType:  reflect.TypeOf(rdb.CreateDatabaseBackupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "database-name",
				Short:      `Name of the database you want to make a backup of`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the backup`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "expires-at",
				Short:      `Expiration date (Format ISO 8601)`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CreateDatabaseBackupRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CreateDatabaseBackup(request)

		},
	}
}

func rdbBackupGet() *core.Command {
	return &core.Command{
		Short:     `Get a database backup`,
		Long:      `Get a database backup.`,
		Namespace: "rdb",
		Resource:  "backup",
		Verb:      "get",
		ArgsType:  reflect.TypeOf(rdb.GetDatabaseBackupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "database-backup-id",
				Short:      `UUID of the database backup`,
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.GetDatabaseBackupRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.GetDatabaseBackup(request)

		},
	}
}

func rdbBackupUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a database backup`,
		Long:      `Update a database backup.`,
		Namespace: "rdb",
		Resource:  "backup",
		Verb:      "update",
		ArgsType:  reflect.TypeOf(rdb.UpdateDatabaseBackupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "database-backup-id",
				Short:      `UUID of the database backup to update`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the Database Backup`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "expires-at",
				Short:      `Expiration date (Format ISO 8601)`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.UpdateDatabaseBackupRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.UpdateDatabaseBackup(request)

		},
	}
}

func rdbBackupDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a database backup`,
		Long:      `Delete a database backup.`,
		Namespace: "rdb",
		Resource:  "backup",
		Verb:      "delete",
		ArgsType:  reflect.TypeOf(rdb.DeleteDatabaseBackupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "database-backup-id",
				Short:      `UUID of the database backup to delete`,
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.DeleteDatabaseBackupRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.DeleteDatabaseBackup(request)

		},
	}
}

func rdbBackupRestore() *core.Command {
	return &core.Command{
		Short:     `Restore a database backup`,
		Long:      `Restore a database backup.`,
		Namespace: "rdb",
		Resource:  "backup",
		Verb:      "restore",
		ArgsType:  reflect.TypeOf(rdb.RestoreDatabaseBackupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "database-name",
				Short:      `Defines the destination database in order to restore into a specified database, the default destination is set to the origin database of the backup`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "database-backup-id",
				Short:      `Backup of a logical database`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "instance-id",
				Short:      `Defines the rdb instance where the backup has to be restored`,
				Required:   true,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.RestoreDatabaseBackupRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.RestoreDatabaseBackup(request)

		},
	}
}

func rdbBackupExport() *core.Command {
	return &core.Command{
		Short:     `Export a database backup`,
		Long:      `Export a database backup.`,
		Namespace: "rdb",
		Resource:  "backup",
		Verb:      "export",
		ArgsType:  reflect.TypeOf(rdb.ExportDatabaseBackupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "database-backup-id",
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ExportDatabaseBackupRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.ExportDatabaseBackup(request)

		},
	}
}

func rdbInstanceClone() *core.Command {
	return &core.Command{
		Short:     `Clone an instance`,
		Long:      `Clone an instance.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "clone",
		ArgsType:  reflect.TypeOf(rdb.CloneInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance you want to clone`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the clone instance`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "node-type",
				Short:      `Node type of the clone`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CloneInstanceRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CloneInstance(request)

		},
	}
}

func rdbEngineList() *core.Command {
	return &core.Command{
		Short:     `List available database engines`,
		Long:      `List available database engines.`,
		Namespace: "rdb",
		Resource:  "engine",
		Verb:      "list",
		ArgsType:  reflect.TypeOf(rdb.ListDatabaseEnginesRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListDatabaseEnginesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			resp, err := api.ListDatabaseEngines(request, scw.WithAllPages())
			if err != nil {
				return nil, err
			}
			return resp.Engines, nil

		},
	}
}

func rdbInstanceUpgrade() *core.Command {
	return &core.Command{
		Short:     `Upgrade an instance to an higher instance type`,
		Long:      `Upgrade your current ` + "`" + `node_type` + "`" + ` or enable high availability on your standalone database instance.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "upgrade",
		ArgsType:  reflect.TypeOf(rdb.UpgradeInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Required:   true,
				Positional: true,
			},
			{
				Name:       "node-type",
				Required:   false,
				Positional: false,
			},
			{
				Name:       "enable-ha",
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.UpgradeInstanceRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.UpgradeInstance(request)

		},
	}
}

func rdbInstanceList() *core.Command {
	return &core.Command{
		Short:     `List instances`,
		Long:      `List instances.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "list",
		ArgsType:  reflect.TypeOf(rdb.ListInstancesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "tags.{index}",
				Short:      `List instance that have a given tags`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `List instance that match a given name pattern`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Criteria to use when ordering instance listing`,
				Required:   false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc", "region", "status_asc", "status_desc"},
			},
			{
				Name:       "organization-id",
				Short:      `Organization ID to list the instance of`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListInstancesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			resp, err := api.ListInstances(request, scw.WithAllPages())
			if err != nil {
				return nil, err
			}
			return resp.Instances, nil

		},
	}
}

func rdbInstanceGet() *core.Command {
	return &core.Command{
		Short:     `Get an instance`,
		Long:      `Get an instance.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "get",
		ArgsType:  reflect.TypeOf(rdb.GetInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance`,
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.GetInstanceRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.GetInstance(request)

		},
	}
}

func rdbInstanceCreate() *core.Command {
	return &core.Command{
		Short:     `Create an instance`,
		Long:      `Create an instance.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "create",
		ArgsType:  reflect.TypeOf(rdb.CreateInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the instance`,
				Required:   true,
				Positional: false,
			},
			{
				Name:       "engine",
				Short:      `Database engine of the database (PostgreSQL, MySQL, ...)`,
				Required:   true,
				Positional: false,
			},
			{
				Name:       "user-name",
				Short:      `Name of the user created when the instance is created`,
				Required:   true,
				Positional: false,
			},
			{
				Name:       "password",
				Short:      `Password of the user`,
				Required:   true,
				Positional: false,
			},
			{
				Name:       "node-type",
				Short:      `Type of node to use for the instance`,
				Required:   true,
				Positional: false,
			},
			{
				Name:       "is-ha-cluster",
				Short:      `Whether or not High-Availability is enabled`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "disable-backup",
				Short:      `Whether or not backups are disabled`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags to apply to the instance`,
				Required:   false,
				Positional: false,
			},
			core.OrganizationIDArgSpec(),
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CreateInstanceRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CreateInstance(request)

		},
	}
}

func rdbInstanceUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an instance`,
		Long:      `Update an instance.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "update",
		ArgsType:  reflect.TypeOf(rdb.UpdateInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "backup-schedule-frequency",
				Short:      `In hours`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "backup-schedule-retention",
				Short:      `In days`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "is-backup-schedule-disabled",
				Short:      `Whether or not the backup schedule is disabled`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the instance`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "instance-id",
				Short:      `UUID of the instance to update`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags of a given instance`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.UpdateInstanceRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.UpdateInstance(request)

		},
	}
}

func rdbInstanceDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an instance`,
		Long:      `Delete an instance.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "delete",
		ArgsType:  reflect.TypeOf(rdb.DeleteInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance to delete`,
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.DeleteInstanceRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.DeleteInstance(request)

		},
	}
}

func rdbCertificateGet() *core.Command {
	return &core.Command{
		Short:     `Get the TLS certificate of an instance`,
		Long:      `Get the TLS certificate of an instance.`,
		Namespace: "rdb",
		Resource:  "certificate",
		Verb:      "get",
		ArgsType:  reflect.TypeOf(rdb.GetInstanceCertificateRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance`,
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.GetInstanceCertificateRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.GetInstanceCertificate(request)

		},
	}
}

func rdbLogPrepare() *core.Command {
	return &core.Command{
		Short:     `Prepare logs of a given instance`,
		Long:      `Prepare your instance logs. Logs will be grouped on a minimum interval of a day.`,
		Namespace: "rdb",
		Resource:  "log",
		Verb:      "prepare",
		ArgsType:  reflect.TypeOf(rdb.PrepareInstanceLogsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance you want logs of`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "start-date",
				Short:      `Start datetime of your log. Format: ` + "`" + `{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z` + "`" + ``,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "end-date",
				Short:      `End datetime of your log. Format: ` + "`" + `{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z` + "`" + ``,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.PrepareInstanceLogsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.PrepareInstanceLogs(request)

		},
	}
}

func rdbLogList() *core.Command {
	return &core.Command{
		Short:     `List available logs of a given instance`,
		Long:      `List available logs of a given instance.`,
		Namespace: "rdb",
		Resource:  "log",
		Verb:      "list",
		ArgsType:  reflect.TypeOf(rdb.ListInstanceLogsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Required:   true,
				Positional: true,
			},
			{
				Name:       "order-by",
				Required:   false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListInstanceLogsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.ListInstanceLogs(request)

		},
	}
}

func rdbInstanceMetrics() *core.Command {
	return &core.Command{
		Short:     `Get instance metrics`,
		Long:      `Get database instance metrics.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "metrics",
		ArgsType:  reflect.TypeOf(rdb.GetInstanceMetricsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "start-date",
				Short:      `Start date to gather metrics from`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "end-date",
				Short:      `End date to gather metrics from`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "metric-name",
				Short:      `Name of the metric to gather`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.GetInstanceMetricsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.GetInstanceMetrics(request)

		},
	}
}

func rdbSettingsAdd() *core.Command {
	return &core.Command{
		Short:     `Add an instance setting`,
		Long:      `Add an instance setting.`,
		Namespace: "rdb",
		Resource:  "settings",
		Verb:      "add",
		ArgsType:  reflect.TypeOf(rdb.AddInstanceSettingsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance you want to add settings to`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "settings.{index}.name",
				Required:   false,
				Positional: false,
			},
			{
				Name:       "settings.{index}.value",
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.AddInstanceSettingsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.AddInstanceSettings(request)

		},
	}
}

func rdbSettingsDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an instance setting`,
		Long:      `Delete an instance setting.`,
		Namespace: "rdb",
		Resource:  "settings",
		Verb:      "delete",
		ArgsType:  reflect.TypeOf(rdb.DeleteInstanceSettingsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance to delete settings from`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "setting-names.{index}",
				Short:      `Settings names to delete`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.DeleteInstanceSettingsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.DeleteInstanceSettings(request)

		},
	}
}

func rdbACLList() *core.Command {
	return &core.Command{
		Short:     `List ACL rules of a given instance`,
		Long:      `List ACL rules of a given instance.`,
		Namespace: "rdb",
		Resource:  "acl",
		Verb:      "list",
		ArgsType:  reflect.TypeOf(rdb.ListInstanceACLRulesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance`,
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListInstanceACLRulesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			resp, err := api.ListInstanceACLRules(request, scw.WithAllPages())
			if err != nil {
				return nil, err
			}
			return resp.Rules, nil

		},
	}
}

func rdbACLAdd() *core.Command {
	return &core.Command{
		Short:     `Add an ACL instance to a given instance`,
		Long:      `Add an ACL instance to a given instance.`,
		Namespace: "rdb",
		Resource:  "acl",
		Verb:      "add",
		ArgsType:  reflect.TypeOf(rdb.AddInstanceACLRulesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance you want to add acl rules to`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "rules.{index}.ip",
				Required:   false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.description",
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.AddInstanceACLRulesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.AddInstanceACLRules(request)

		},
	}
}

func rdbACLDelete() *core.Command {
	return &core.Command{
		Short:     `Delete ACL rules of a given instance`,
		Long:      `Delete ACL rules of a given instance.`,
		Namespace: "rdb",
		Resource:  "acl",
		Verb:      "delete",
		ArgsType:  reflect.TypeOf(rdb.DeleteInstanceACLRulesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance you want to delete an ACL rules from`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "acl-rule-ips.{index}",
				Short:      `ACL rules IP present on the instance`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.DeleteInstanceACLRulesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.DeleteInstanceACLRules(request)

		},
	}
}

func rdbUserList() *core.Command {
	return &core.Command{
		Short:     `List users of a given instance`,
		Long:      `List users of a given instance.`,
		Namespace: "rdb",
		Resource:  "user",
		Verb:      "list",
		ArgsType:  reflect.TypeOf(rdb.ListUsersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the user`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Criteria to use when ordering users listing`,
				Required:   false,
				Positional: false,
				EnumValues: []string{"name_asc", "name_desc", "is_admin_asc", "is_admin_desc"},
			},
			{
				Name:       "instance-id",
				Short:      `UUID of the instance`,
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListUsersRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			resp, err := api.ListUsers(request, scw.WithAllPages())
			if err != nil {
				return nil, err
			}
			return resp.Users, nil

		},
	}
}

func rdbUserCreate() *core.Command {
	return &core.Command{
		Short:     `Create an user in a given instance`,
		Long:      `Create an user in a given instance.`,
		Namespace: "rdb",
		Resource:  "user",
		Verb:      "create",
		ArgsType:  reflect.TypeOf(rdb.CreateUserRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance you want to create a user in`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the user you want to create`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "password",
				Short:      `Password of the user you want to create`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "is-admin",
				Short:      `Whether the user you want to create will have administrative privileges`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CreateUserRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CreateUser(request)

		},
	}
}

func rdbUserUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an user in a given instance`,
		Long:      `Update an user in a given instance.`,
		Namespace: "rdb",
		Resource:  "user",
		Verb:      "update",
		ArgsType:  reflect.TypeOf(rdb.UpdateUserRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance the user belongs to`,
				Required:   true,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the database user`,
				Required:   true,
				Positional: false,
			},
			{
				Name:       "password",
				Short:      `Password of the database user`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "is-admin",
				Short:      `Whether or not this user got administrative privileges`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.UpdateUserRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.UpdateUser(request)

		},
	}
}

func rdbUserDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an user in a given instance`,
		Long:      `Delete an user in a given instance.`,
		Namespace: "rdb",
		Resource:  "user",
		Verb:      "delete",
		ArgsType:  reflect.TypeOf(rdb.DeleteUserRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance to delete a user from`,
				Required:   true,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the user`,
				Required:   true,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.DeleteUserRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			e = api.DeleteUser(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "user",
				Verb:     "delete",
			}, nil
		},
	}
}

func rdbDatabaseList() *core.Command {
	return &core.Command{
		Short:     `List all database in a given instance`,
		Long:      `List all database in a given instance.`,
		Namespace: "rdb",
		Resource:  "database",
		Verb:      "list",
		ArgsType:  reflect.TypeOf(rdb.ListDatabasesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the database`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "managed",
				Short:      `Whether or not the database is managed`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "owner",
				Short:      `User that owns this database`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Criteria to use when ordering database listing`,
				Required:   false,
				Positional: false,
				EnumValues: []string{"name_asc", "name_desc", "size_asc", "size_desc"},
			},
			{
				Name:       "instance-id",
				Short:      `UUID of the instance to list database of`,
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListDatabasesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			resp, err := api.ListDatabases(request, scw.WithAllPages())
			if err != nil {
				return nil, err
			}
			return resp.Databases, nil

		},
	}
}

func rdbDatabaseCreate() *core.Command {
	return &core.Command{
		Short:     `Create a database in a given instance`,
		Long:      `Create a database in a given instance.`,
		Namespace: "rdb",
		Resource:  "database",
		Verb:      "create",
		ArgsType:  reflect.TypeOf(rdb.CreateDatabaseRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance where to create the database`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the database`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CreateDatabaseRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CreateDatabase(request)

		},
	}
}

func rdbDatabaseDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a database in a given instance`,
		Long:      `Delete a database in a given instance.`,
		Namespace: "rdb",
		Resource:  "database",
		Verb:      "delete",
		ArgsType:  reflect.TypeOf(rdb.DeleteDatabaseRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance where to delete the database`,
				Required:   true,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the database to delete`,
				Required:   true,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.DeleteDatabaseRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			e = api.DeleteDatabase(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "database",
				Verb:     "delete",
			}, nil
		},
	}
}

func rdbPrivilegeList() *core.Command {
	return &core.Command{
		Short:     `List privileges of a given user in a given database in a given instance`,
		Long:      `List privileges of a given user in a given database in a given instance.`,
		Namespace: "rdb",
		Resource:  "privilege",
		Verb:      "list",
		ArgsType:  reflect.TypeOf(rdb.ListPrivilegesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Criteria to use when ordering privileges listing`,
				Required:   false,
				Positional: false,
				EnumValues: []string{"user_name_asc", "user_name_desc", "database_name_asc", "database_name_desc"},
			},
			{
				Name:       "database-name",
				Short:      `Name of the database`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "instance-id",
				Short:      `UUID of the instance`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "user-name",
				Short:      `Name of the user`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListPrivilegesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			resp, err := api.ListPrivileges(request, scw.WithAllPages())
			if err != nil {
				return nil, err
			}
			return resp.Privileges, nil

		},
	}
}

func rdbPrivilegeSet() *core.Command {
	return &core.Command{
		Short:     `Set privileges of a given user in a given database in a given instance`,
		Long:      `Set privileges of a given user in a given database in a given instance.`,
		Namespace: "rdb",
		Resource:  "privilege",
		Verb:      "set",
		ArgsType:  reflect.TypeOf(rdb.SetPrivilegeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the instance`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "database-name",
				Short:      `Name of the database`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "user-name",
				Short:      `Name of the user`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "permission",
				Short:      `Permission to set (Read, Read/Write, All, Custom)`,
				Required:   false,
				Positional: false,
				EnumValues: []string{"readonly", "readwrite", "all", "custom", "none"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.SetPrivilegeRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.SetPrivilege(request)

		},
	}
}

func rdbNodeTypeList() *core.Command {
	return &core.Command{
		Short:     `List available node types`,
		Long:      `List available node types.`,
		Namespace: "rdb",
		Resource:  "node-type",
		Verb:      "list",
		ArgsType:  reflect.TypeOf(rdb.ListNodeTypesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "include-disabled-types",
				Short:      `Whether or not to include disabled types`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListNodeTypesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			resp, err := api.ListNodeTypes(request, scw.WithAllPages())
			if err != nil {
				return nil, err
			}
			return resp.NodeTypes, nil

		},
	}
}
