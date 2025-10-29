package main

import (
	"flag"
)

var (
	DefaultMigrationDir = "."

	flags        = flag.NewFlagSet("goose", flag.ExitOnError)

	dir          = flags.String("dir", DefaultMigrationDir, "directory with migration files, (GOOSE_MIGRATION_DIR env variable supported)")
	table        = flags.String("table", "", "migrations table name")
	verbose      = flags.Bool("v", false, "enable verbose mode")
	help         = flags.Bool("h", false, "print help")
	versionFlag  = flags.Bool("version", false, "print version")
	certfile     = flags.String("certfile", "", "file path to root CA's certificates in pem format (only support on mysql)")
	sequential   = flags.Bool("s", false, "use sequential numbering for new migrations")
	allowMissing = flags.Bool("allow-missing", false, "applies missing (out-of-order) migrations")
	sslcert      = flags.String("ssl-cert", "", "file path to SSL certificates in pem format (only support on mysql)")
	sslkey       = flags.String("ssl-key", "", "file path to SSL key in pem format (only support on mysql)")
	noVersioning = flags.Bool("no-versioning", false, "apply migration commands with no versioning, in file order, from directory pointed to")
	noColor      = flags.Bool("no-color", false, "disable color output (NO_COLOR env variable supported)")
	timeout      = flags.Duration("timeout", 0, "maximum allowed duration for queries to run; e.g., 1h13m")
	envFile      = flags.String("env", "", "load environment variables from file (default .env)")
)

