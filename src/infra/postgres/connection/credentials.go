package connection

import (
	"fmt"
	"os"
	"tcms/src/utils/envs"
)

func getDatabaseURI() string {
	schema := os.Getenv(envs.DatabaseSchema)
	user := os.Getenv(envs.DatabaseUser)
	password := os.Getenv(envs.DatabasePassword)
	host := os.Getenv(envs.DatabaseHost)
	port := os.Getenv(envs.DatabasePort)
	name := os.Getenv(envs.DatabaseName)
	sslMode := os.Getenv(envs.DatabaseSSLMode)

	credentials := fmt.Sprintf("%s:%s", user, password)
	destination := fmt.Sprintf("%s:%s/%s", host, port, name)

	return fmt.Sprintf("%s://%s@%s?sslmode=%s", schema, credentials, destination, sslMode)
}
