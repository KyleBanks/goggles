# git
--
    import "github.com/KyleBanks/go-kit/git/"

Package git provides git source control functionality.

## Usage

#### func  InstallPreCommitHook

```go
func InstallPreCommitHook(hook string, path string) error
```
InstallPreCommitHook takes a string to set as the pre-commit git hook for the
.git directory specified by the path provided.

The path should be to the parent of the .git directory. For example:
<gopath>/src/github.com/KyleBanks/go-kit
