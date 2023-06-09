package boostrap

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/osmait/blog-go/internal/creating"
	"github.com/osmait/blog-go/internal/find"
	"github.com/osmait/blog-go/internal/platfrom/server"
	"github.com/osmait/blog-go/internal/platfrom/storage/postgres"
	"github.com/rs/zerolog/log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	host       = "localhost"
	port       = 3000
	dbUser     = "osmait"
	dbPassword = "admin123"
	dbName     = "my_store"
	dbHost     = "localhost"
	dbPort     = 5432
)

func Run() error {
	postgresURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		return err
	}
	runDBMigration("file://db/migration", postgresURI)

	userRepository := postgres.NewUserRepository(db)
	postRepository := postgres.NewPostRepository(db)
	creatingUserService := creating.NewCreateUseService(userRepository)
	creatingPostService := creating.NewCreatePostService(postRepository)
	findUser := find.NewFind(userRepository)

	srv := server.NewServer(host, port, *creatingUserService, *creatingPostService, *findUser)
	return srv.Run()

}

func runDBMigration(migrationURL string, dbSource string) {
	fmt.Println(migrationURL)
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}
