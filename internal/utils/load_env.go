package utils

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

// LoadEnvFromLocalRepo loads env vars from .env file.
func LoadEnvFromLocalRepo(configFolder string, envFile string) {
	cwd, _ := os.Getwd()
	// absPath, _ := filepath.Abs("../" + projectDirName + "/" + envFile)
	// nwd := filepath.Join(cwd, "..")
	absPath := filepath.Join(cwd, "../"+configFolder+"/"+envFile)
	SetEnvVariablesFromAbs(absPath)
}

// SetEnvVariablesFromAbs sets environment variables from an absolute path.
func SetEnvVariablesFromAbs(absPath string) {
	err := godotenv.Load(absPath)
	if err != nil {
		log.WithFields(log.Fields{
			"cause": err,
		}).Fatal("Problem loading .env file.")
		os.Exit(-1)
	}
}

// GetQueryFromTxt return neo4j cypher query from a text file.
// QueryFolder within local wyrth-io repo, github.com/wyrth-io/<QueryFolder>.
func GetQueryFromTxt(queryFolder string, filename string) string {
	absPath, _ := filepath.Abs("../" + queryFolder + "/" + filename)
	query, err := os.ReadFile(filepath.Clean(absPath))
	if err != nil {
		panic(err)
	}
	return string(query)
}
