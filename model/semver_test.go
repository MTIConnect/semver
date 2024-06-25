package model_test

import (
	"testing"

	"github.com/MTIConnect/semver/model"
)

func TestSemver(t *testing.T) {
	subject := new(model.Semver)
	for _, version := range []string{
		`1.2.3`,
		`3.5.7-beta`,
		`5.7.11+2024-06-25`,
		`7.11.13-beta+2024-06-25`,
	} {
		if err := subject.Parse(version); err != nil {
			t.Errorf(`could not parse version %s: %s`, version, err.Error())
		}
		if expected, actual := version, subject.String(); expected != actual {
			t.Errorf(`version not rendered as original: expected %q; actual: %q`, expected, actual)
		}
	}
}
