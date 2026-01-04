package ximg_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) {
	t.Helper()

	err := os.MkdirAll("output", 0750)
	require.NoError(t, err)
}
