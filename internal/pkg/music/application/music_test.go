package application_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"magical.dev/tranquility/internal/pkg/music/application"
)

func TestCreateMusicApplicationService(t *testing.T) {
	ap := application.NewMusicApplicationService()
	assert.NotNil(t, ap)
}
