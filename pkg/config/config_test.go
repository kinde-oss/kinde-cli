package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeDomain(t *testing.T) {
	valid := map[string]string{
		"mybiz.kinde.com":           "mybiz.kinde.com",
		"https://mybiz.kinde.com":   "mybiz.kinde.com",
		"https://MyBiz.kinde.com/":  "mybiz.kinde.com",
		"http://mybiz.kinde.com":    "mybiz.kinde.com",
		"  auth.example.com  ":      "auth.example.com",
		"https://auth.example.com/": "auth.example.com",
	}
	for input, want := range valid {
		got, err := NormalizeDomain(input)
		assert.NoError(t, err, input)
		assert.Equal(t, want, got, input)
	}

	for _, input := range []string{"", "https://", "mybiz.kinde.com/path", "my biz.kinde.com"} {
		_, err := NormalizeDomain(input)
		assert.Error(t, err, input)
	}
}

func TestUseDomain_DoesNotChangeSavedDefault(t *testing.T) {
	cfg := &Config{CurrentEnvironment: "a.kinde.com", Environments: map[string]Environment{}}

	assert.NoError(t, cfg.UseDomain("b.kinde.com"))

	assert.Equal(t, "b.kinde.com", cfg.GetEnvironment().DomainName)
	assert.Equal(t, "a.kinde.com", cfg.CurrentEnvironment)
	assert.False(t, cfg.changed, "selecting a domain for one run is not a change to save")
}

func TestSetDefaultDomain_MarksChanged(t *testing.T) {
	cfg := &Config{CurrentEnvironment: "a.kinde.com"}

	cfg.SetDefaultDomain("a.kinde.com")
	assert.False(t, cfg.changed, "setting the same default is not a change")

	cfg.SetDefaultDomain("b.kinde.com")
	assert.True(t, cfg.changed)
	assert.Equal(t, "b.kinde.com", cfg.CurrentEnvironment)
}
