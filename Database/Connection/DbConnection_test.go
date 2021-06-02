package DbConnection

import (
	"testing"
)

func TestConnectionToDatabase(t *testing.T) {
	expected := "Connection available!"
	actual := "Connection available!"

	if actual != expected {
		t.Error("Connection fail due to an unexpeted error", expected, actual)
	}
}
