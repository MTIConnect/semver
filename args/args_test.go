package args_test

import (
	"testing"

	"github.com/MTIConnect/semver/args"
	"github.com/MTIConnect/semver/model"
)

func TestArgsParse(t *testing.T) {
	subject := args.Args{
		Version: model.Semver{
			Major:         1,
			Minor:         2,
			Patch:         3,
			Prerelease:    `beta`,
			BuildMetadata: `dev`,
		},
	}
	t.Run(`major`, func(t *testing.T) {
		t.Run(`noop`, func(t *testing.T) {
			expected := subject.Version.Major
			subject.Major = ``
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Major; expected != actual {
				t.Errorf(`unexpected major version: expected %d; found %d`, expected, actual)
			}
		})
		t.Run(`plus`, func(t *testing.T) {
			expected := subject.Version.Major + 1
			subject.Major = `+`
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Major; expected != actual {
				t.Errorf(`unexpected major version: expected %d; found %d`, expected, actual)
			}
		})
		t.Run(`minus`, func(t *testing.T) {
			expected := subject.Version.Major - 1
			subject.Major = `-`
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Major; expected != actual {
				t.Errorf(`unexpected major version: expected %d; found %d`, expected, actual)
			}
		})
		t.Run(`set`, func(t *testing.T) {
			expected := 7
			subject.Major = `7`
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Major; expected != actual {
				t.Errorf(`unexpected major version: expected %d; found %d`, expected, actual)
			}
		})
	})
	t.Run(`minor`, func(t *testing.T) {
		t.Run(`noop`, func(t *testing.T) {
			expected := subject.Version.Minor
			subject.Minor = ``
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Minor; expected != actual {
				t.Errorf(`unexpected minor version: expected %d; found %d`, expected, actual)
			}
		})
		t.Run(`plus`, func(t *testing.T) {
			expected := subject.Version.Minor + 1
			subject.Minor = `+`
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Minor; expected != actual {
				t.Errorf(`unexpected minor version: expected %d; found %d`, expected, actual)
			}
		})
		t.Run(`minus`, func(t *testing.T) {
			expected := subject.Version.Minor - 1
			subject.Minor = `-`
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Minor; expected != actual {
				t.Errorf(`unexpected minor version: expected %d; found %d`, expected, actual)
			}
		})
		t.Run(`set`, func(t *testing.T) {
			expected := 7
			subject.Minor = `7`
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Minor; expected != actual {
				t.Errorf(`unexpected minor version: expected %d; found %d`, expected, actual)
			}
		})
	})
	t.Run(`patch`, func(t *testing.T) {
		t.Run(`noop`, func(t *testing.T) {
			expected := subject.Version.Patch
			subject.Patch = ``
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Patch; expected != actual {
				t.Errorf(`unexpected patch version: expected %d; found %d`, expected, actual)
			}
		})
		t.Run(`plus`, func(t *testing.T) {
			expected := subject.Version.Patch + 1
			subject.Patch = `+`
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Patch; expected != actual {
				t.Errorf(`unexpected patch version: expected %d; found %d`, expected, actual)
			}
		})
		t.Run(`minus`, func(t *testing.T) {
			expected := subject.Version.Patch - 1
			subject.Patch = `-`
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Patch; expected != actual {
				t.Errorf(`unexpected patch version: expected %d; found %d`, expected, actual)
			}
		})
		t.Run(`set`, func(t *testing.T) {
			expected := 7
			subject.Patch = `7`
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Patch; expected != actual {
				t.Errorf(`unexpected patch version: expected %d; found %d`, expected, actual)
			}
		})
	})
	t.Run(`prerelease`, func(t *testing.T) {
		t.Run(`noop`, func(t *testing.T) {
			expected := subject.Version.Prerelease
			subject.Prerelease = ``
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Prerelease; expected != actual {
				t.Errorf(`unexpected prerelease version: expected %s; found %s`, expected, actual)
			}
		})
		t.Run(`minus`, func(t *testing.T) {
			expected := ``
			subject.Prerelease = `-`
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Prerelease; expected != actual {
				t.Errorf(`unexpected prerelease version: expected %s; found %s`, expected, actual)
			}
		})
		t.Run(`set`, func(t *testing.T) {
			expected := `minty`
			subject.Prerelease = `minty`
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.Prerelease; expected != actual {
				t.Errorf(`unexpected prerelease version: expected %s; found %s`, expected, actual)
			}
		})
	})
	t.Run(`build-metadata`, func(t *testing.T) {
		t.Run(`noop`, func(t *testing.T) {
			expected := subject.Version.BuildMetadata
			subject.BuildMetadata = ``
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.BuildMetadata; expected != actual {
				t.Errorf(`unexpected build-metadata version: expected %s; found %s`, expected, actual)
			}
		})
		t.Run(`minus`, func(t *testing.T) {
			expected := ``
			subject.BuildMetadata = `-`
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.BuildMetadata; expected != actual {
				t.Errorf(`unexpected build-metadata version: expected %s; found %s`, expected, actual)
			}
		})
		t.Run(`set`, func(t *testing.T) {
			expected := `minty`
			subject.BuildMetadata = `minty`
			if err := subject.Process(); err != nil {
				t.Errorf(`could not process request: %s`, err.Error())
			}
			if actual := subject.Version.BuildMetadata; expected != actual {
				t.Errorf(`unexpected build-metadata version: expected %s; found %s`, expected, actual)
			}
		})
	})
}
