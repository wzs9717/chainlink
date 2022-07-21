package chainlink

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSecretGenerator_sessionSecret(t *testing.T) {
	t.Parallel()
	rootDir := path.Join("/tmp/chainlink_test", "TestSecretGenerator_sessionSecret")
	err := os.MkdirAll(rootDir, os.FileMode(0770))
	require.NoError(t, err)
	defer os.RemoveAll(rootDir)

	var secretGenerator FilePersistedSecretGenerator

	initial, err := secretGenerator.Generate(rootDir)
	require.NoError(t, err)
	require.NotEqual(t, "", initial)
	require.NotEqual(t, "clsession_test_secret", initial)

	second, err := secretGenerator.Generate(rootDir)
	require.NoError(t, err)
	require.Equal(t, initial, second)
}
