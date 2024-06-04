package dbtest

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	_ "embed" //nolint

	"fmt"
	"os"
	"testing"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// StartContainerDatabase starts postgres container with default test settings, and migrates the db. Caller must terminate container.
func StartContainerDatabase(ctx context.Context, dbName string, t *testing.T, migrationsDirRelPath string) (db.Store, testcontainers.Container) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	settings := getTestDbSettings(dbName)
	pgPort := "5432/tcp"
	dbURL := func(_ string, port nat.Port) string {
		return fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", settings.DB.User, settings.DB.Password, port.Port(), settings.DB.Name)
	}
	cr := testcontainers.ContainerRequest{
		Image:        "postgres:12.9-alpine",
		Env:          map[string]string{"POSTGRES_USER": settings.DB.User, "POSTGRES_PASSWORD": settings.DB.Password, "POSTGRES_DB": settings.DB.Name},
		ExposedPorts: []string{pgPort},
		Cmd:          []string{"postgres", "-c", "fsync=off"},
		WaitingFor:   wait.ForSQL(nat.Port(pgPort), "postgres", dbURL),
	}

	pgContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: cr,
		Started:          true,
	})
	if err != nil {
		return handleContainerStartErr(ctx, err, pgContainer, t)
	}
	mappedPort, err := pgContainer.MappedPort(ctx, nat.Port(pgPort))
	if err != nil {
		return handleContainerStartErr(ctx, errors.Wrap(err, "failed to get container external port"), pgContainer, t)
	}
	fmt.Printf("postgres container session %s ready and running at port: %s \n", pgContainer.SessionID(), mappedPort)
	//defer pgContainer.Terminate(ctx) // this should be done by the caller

	settings.DB.Port = mappedPort.Port()
	pdb := db.NewDbConnectionForTest(ctx, &settings.DB, false)
	pdb.WaitForDB(logger)

	_, err = pdb.DBS().Writer.Exec(fmt.Sprintf(`
		grant usage on schema public to public;
		grant create on schema public to public;
		CREATE SCHEMA IF NOT EXISTS %s;
		ALTER USER postgres SET search_path = %s, public;
		SET search_path = %s, public;
		`, dbName, dbName, dbName))
	if err != nil {
		return handleContainerStartErr(ctx, errors.Wrapf(err, "failed to apply schema. session: %s, port: %s",
			pgContainer.SessionID(), mappedPort.Port()), pgContainer, t)
	}
	logger.Info().Msgf("set default search_path for user postgres to %s", dbName)
	// add truncate tables func
	_, err = pdb.DBS().Writer.Exec(fmt.Sprintf(`
CREATE OR REPLACE FUNCTION %s.truncate_tables() RETURNS void AS $$
DECLARE
    statements CURSOR FOR
        SELECT tablename FROM pg_tables
        WHERE schemaname = '%s' and tablename != 'migrations';
BEGIN
    FOR stmt IN statements LOOP
        EXECUTE 'TRUNCATE TABLE ' || quote_ident(stmt.tablename) || ' CASCADE;';
    END LOOP;
END;
$$ LANGUAGE plpgsql;
`, dbName, dbName))
	if err != nil {
		return handleContainerStartErr(ctx, errors.Wrap(err, "failed to create truncate func"), pgContainer, t)
	}

	goose.SetTableName(dbName + ".migrations")
	if err := goose.RunContext(ctx, "up", pdb.DBS().Writer.DB, migrationsDirRelPath); err != nil {
		return handleContainerStartErr(ctx, errors.Wrap(err, "failed to apply goose migrations for test"), pgContainer, t)
	}

	return pdb, pgContainer
}

// getTestDbSettings builds test db config.Settings object
func getTestDbSettings(dbName string) config.Settings {
	settings := config.Settings{
		LogLevel: "info",
		DB: db.Settings{
			Name:               dbName,
			Host:               "localhost",
			Port:               "6669",
			User:               "postgres",
			Password:           "postgres",
			MaxOpenConnections: 2,
			MaxIdleConnections: 2,
		},
		ServiceName: "vehicle-signal-decoding",
	}
	return settings
}

func BuildRequest(method, url, body string) *http.Request {
	req, _ := http.NewRequest(
		method,
		url,
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	return req
}

func handleContainerStartErr(ctx context.Context, err error, container testcontainers.Container, t *testing.T) (db.Store, testcontainers.Container) {
	if err != nil {
		fmt.Println("start container error: " + err.Error())
		if container != nil {
			_ = container.Terminate(ctx)
		}
		t.Fatal(err)
	}
	return db.Store{}, container
}

// TruncateTables truncates tables for the test db, useful to run as teardown at end of each DB dependent test.
func TruncateTables(db *sql.DB, dbName string, t *testing.T) {
	query := fmt.Sprintf(`SELECT %s.truncate_tables();`, dbName)
	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("Error truncating tables in schema '%s': %s\n", dbName, err)
		t.Fatal(err)
	}
	// TruncateTables not working, manually deleting for now

	deleteDeviceSettingsQuery := `DELETE FROM device_settings CASCADE;`
	_, err = db.Exec(deleteDeviceSettingsQuery)
	if err != nil {
		fmt.Println("error:", err)
	}

	deleteTemplateVehiclesQuery := `DELETE FROM template_vehicles CASCADE;`
	_, err = db.Exec(deleteTemplateVehiclesQuery)
	if err != nil {
		fmt.Println("error:", err)
	}

	deleteDBCFilesQuery := `DELETE FROM dbc_files CASCADE;`
	_, err = db.Exec(deleteDBCFilesQuery)
	if err != nil {
		fmt.Println("error:", err)
	}

	deletePidConfigsQuery := `DELETE FROM pid_configs CASCADE;`
	_, err = db.Exec(deletePidConfigsQuery)
	if err != nil {
		fmt.Println("error:", err)
	}

	deleteTemplateDeviceDefinitionsQuery := `DELETE FROM template_device_definitions CASCADE;`
	_, err = db.Exec(deleteTemplateDeviceDefinitionsQuery)
	if err != nil {
		fmt.Println("error:", err)
	}

	deleteTemplatesQuery := `DELETE FROM templates CASCADE;`
	_, err = db.Exec(deleteTemplatesQuery)
	if err != nil {
		fmt.Println("error:", err)
	}

}

func Logger() *zerolog.Logger {
	l := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "vehicle-signal-decoding").
		Logger()
	return &l
}
