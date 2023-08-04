package db

type SqlMigrator interface {
	Migration()
	Seeding(scriptDir string)
}
