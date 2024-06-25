package model

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

var pattern = regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

// Semver represents a semantic versioning string, broken into its parts
type Semver struct {
	Major         int
	Minor         int
	Patch         int
	Prerelease    string
	BuildMetadata string
}

// Parse takes a semantic versioning string and splits it into its component parts
func (s *Semver) Parse(v string) error {
	matches := pattern.FindStringSubmatch(v)
	for i, group := range pattern.SubexpNames() {
		var err error
		match := matches[i]
		switch group {
		case ``:
			continue
		case `major`:
			s.Major, err = strconv.Atoi(match)
			if err != nil {
				return fmt.Errorf(`could not set major version: %w`, err)
			}
		case `minor`:
			s.Minor, err = strconv.Atoi(match)
			if err != nil {
				return fmt.Errorf(`could not set minor version: %w`, err)
			}
		case `patch`:
			s.Patch, err = strconv.Atoi(match)
			if err != nil {
				return fmt.Errorf(`could not set patch version: %w`, err)
			}
		case `prerelease`:
			s.Prerelease = match
		case `buildmetadata`:
			s.BuildMetadata = match
		default:
			return fmt.Errorf(`unknown group %q`, group)
		}
	}
	return nil
}

// String implements fmt.Stringer and assembles the component parts into a semantic versioning string
func (s Semver) String() string {
	buff := new(bytes.Buffer)
	fmt.Fprintf(buff, `%d.%d.%d`, s.Major, s.Minor, s.Patch)
	if s.Prerelease != `` {
		fmt.Fprintf(buff, `-%s`, s.Prerelease)
	}
	if s.BuildMetadata != `` {
		fmt.Fprintf(buff, `+%s`, s.BuildMetadata)
	}
	return buff.String()
}
