package utils_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
)

func Test_LOADENV_LOADS_FROM_TESTDATAFOLDER(t *testing.T) {
	tt := testutils.NewTester(t)

	// Check if 'testdata' folder exists.
	_, err := os.Stat("../testdata/")
	// os.IsNotExist(err)
	tt.CheckNotError(err)

	// Create abs path and change working directory
	var filename string
	_, filename, _, _ = runtime.Caller(0) //nolint:dogsled
	dir := filepath.Join(filepath.Dir(filename), "../testdata/")
	err = os.Chdir(dir)
	tt.CheckNotError(err)

	// Check if 'env.test' file exists
	absPath := filepath.Join(dir, "/test.env")
	_, err = os.Stat(absPath)
	tt.CheckNotError(err)

	utils.SetEnvVariablesFromAbs(absPath)

	neo4jURI, found := os.LookupEnv("NEO4J_URI")
	tt.CheckEqual("neo4j://localhost:7687", neo4jURI)
	tt.CheckEqual(true, found)

	neo4jUsername, found := os.LookupEnv("NEO4J_USERNAME")
	tt.CheckEqual("neo4j", neo4jUsername)
	tt.CheckEqual(true, found)

	neo4jPassword, found := os.LookupEnv("NEO4J_PASSWORD")
	tt.CheckEqual("testpass", neo4jPassword)
	tt.CheckEqual(true, found)
}

func Test_LOADENV_READWRITE(t *testing.T) {
	tt := testutils.NewTester(t)

	// Create an .env file within temp directory of OS
	file, err := os.CreateTemp("", "test-*.env")
	defer func() {
		// Have these files deleted on return
		err = os.Remove(file.Name())
		tt.CheckNotError(err)
	}()

	tt.CheckNotError(err)

	// Write to the file
	_, err = file.Write([]byte("test data\n"))
	tt.CheckNotError(err)

	// Close file.
	err = file.Close()
	tt.CheckNotError(err)

	// Check if we can read the file
	_, err = os.ReadFile(file.Name())
	tt.CheckNotError(err)
}
