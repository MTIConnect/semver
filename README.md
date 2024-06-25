# semver
Just a tiny tool for manipulating semantic version numbers

```
Usage
./semver [FLAGS] version

FLAGS:
  -b, --build-metadata string   Set the build metadata info. "-" to clear the value
  -M, --major string            Set the major version. "+" or "-" to inc/decrement
  -m, --minor string            Set the minor version. "+" or "-" to inc/decrement
  -p, --patch string            Set the patch version. "+" or "-" to inc/decrement
  -r, --prerelease string       Set the prerelease info. "-" to clear the value

EXAMPLE:
    ./semver -p+ -r some-feature 1.2.3
    1.2.4-some-feature
```
