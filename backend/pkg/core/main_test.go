package core_test

import (
	"os"
	"testing"

	hansip "github.com/asasmoyo/pq-hansip"
	"github.com/awanku/awanku/pkg/testhelper"
)

var db *hansip.Cluster

func TestMain(m *testing.M) {
	var dbDone func()
	db, dbDone = testhelper.DBCluster()
	defer dbDone()

	os.Exit(m.Run())
}
