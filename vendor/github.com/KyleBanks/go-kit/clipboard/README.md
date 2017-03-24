# clipboard
--
    import "github.com/KyleBanks/go-kit/clipboard/"

Package clipboard provides the ability to read and write to the system
clipboard.

Note: Currently only supports Mac OS.

## Usage

#### func  Read

```go
func Read() (io.Reader, error)
```
Read returns the current contents of the system clipboard.

#### func  ReadString

```go
func ReadString() (string, error)
```
ReadString returns the current contents of the system clipboard as a string.

#### func  Write

```go
func Write(r io.Reader) error
```
Write stores the contents of the reader provided in the system clipboard.

#### func  WriteString

```go
func WriteString(s string) error
```
WriteString stores a string in the system clipboard.
