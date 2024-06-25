package args

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MTIConnect/semver/model"
	"github.com/spf13/pflag"
)

// Args represents the command line arguments for the executable
type Args struct {
	Version       model.Semver
	Major         string
	Minor         string
	Patch         string
	Prerelease    string
	BuildMetadata string
}

// ParseArgs handles the command line arguments
func ParseArgs() (*Args, error) {
	var args Args
	pflag.StringVarP(&args.Major, `major`, `M`, ``, `Set the major version. "+" or "-" to inc/decrement`)
	pflag.StringVarP(&args.Minor, `minor`, `m`, ``, `Set the minor version. "+" or "-" to inc/decrement`)
	pflag.StringVarP(&args.Patch, `patch`, `p`, ``, `Set the patch version. "+" or "-" to inc/decrement`)
	pflag.StringVarP(&args.Prerelease, `prerelease`, `r`, ``, `Set the prerelease info. "-" to clear the value`)
	pflag.StringVarP(&args.BuildMetadata, `build-metadata`, `b`, ``, `Set the build metadata info. "-" to clear the value`)
	pflag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage")
		fmt.Fprintf(os.Stderr, "%s [FLAGS] version\n", os.Args[0])
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "FLAGS:")
		pflag.PrintDefaults()
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "EXAMPLE:")
		fmt.Fprintf(os.Stderr, "    %s -p+ -r some-feature 1.2.3\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "    1.2.4-some-feature")
	}
	pflag.Parse()

	if pflag.NArg() > 0 {
		return &args, args.Version.Parse(pflag.Arg(0))
	}
	return &args, nil
}

func (a *Args) processMajor() error {
	switch a.Major {
	case ``:
		return nil
	case `+`:
		a.Version.Major = a.Version.Major + 1
	case `-`:
		a.Version.Major = a.Version.Major - 1
	default:
		var err error
		a.Version.Major, err = strconv.Atoi(a.Major)
		if err != nil {
			return fmt.Errorf(`unable to set major version: %w`, err)
		}
	}
	return nil
}
func (a *Args) processMinor() error {
	switch a.Minor {
	case ``:
		return nil
	case `+`:
		a.Version.Minor = a.Version.Minor + 1
	case `-`:
		a.Version.Minor = a.Version.Minor - 1
	default:
		var err error
		a.Version.Minor, err = strconv.Atoi(a.Minor)
		if err != nil {
			return fmt.Errorf(`unable to set minor version: %w`, err)
		}
	}
	return nil
}
func (a *Args) processPatch() error {
	switch a.Patch {
	case ``:
		return nil
	case `+`:
		a.Version.Patch = a.Version.Patch + 1
	case `-`:
		a.Version.Patch = a.Version.Patch - 1
	default:
		var err error
		a.Version.Patch, err = strconv.Atoi(a.Patch)
		if err != nil {
			return fmt.Errorf(`unable to set patch version: %w`, err)
		}
	}
	return nil
}

func (a *Args) processPrerelease() error {
	switch a.Prerelease {
	case ``:
		return nil
	case `-`:
		a.Version.Prerelease = ``
	default:
		a.Version.Prerelease = a.Prerelease
	}
	return nil
}

func (a *Args) processBuildMetadata() error {
	switch a.BuildMetadata {
	case ``:
		return nil
	case `-`:
		a.Version.BuildMetadata = ``
	default:
		a.Version.BuildMetadata = a.BuildMetadata
	}
	return nil
}

// Process executes the specified operation on each of the command line arguments
func (a *Args) Process() error {
	if err := a.processMajor(); err != nil {
		return err
	}
	if err := a.processMinor(); err != nil {
		return err
	}
	if err := a.processPatch(); err != nil {
		return err
	}
	if err := a.processPrerelease(); err != nil {
		return err
	}
	if err := a.processBuildMetadata(); err != nil {
		return err
	}
	return nil
}
