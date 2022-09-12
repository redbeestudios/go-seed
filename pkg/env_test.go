package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEnv(t *testing.T) {

	t.Run("NewEnv returns object if environment is valid", func(t *testing.T) {
		env, err := NewEnv("local")
		assert.NoError(t, err)
		assert.Equal(t, LocalEnvironment.String(), env)
	})

	t.Run("NewEnv returns error if environment is not valid", func(t *testing.T) {
		_, err := NewEnv("nomeimportanada")
		assert.Error(t, err)
	})

}
