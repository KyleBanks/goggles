# storage
--
    import "github.com/KyleBanks/go-kit/storage/"

Package storage provides the ability to persist and retrieve structs.

## Usage

#### type FileStore

```go
type FileStore struct {
	Dirname  string
	Filename string
}
```

FileStore represents a simple on-disk file storage system.

#### func  NewFileStore

```go
func NewFileStore(d, f string) *FileStore
```
NewFileStore returns an initialized FileStore.

#### func (FileStore) Load

```go
func (f FileStore) Load(v interface{}) error
```
Load attempts to read and decode the storage file into the value provided.

#### func (FileStore) Save

```go
func (f FileStore) Save(v interface{}) error
```
Save attempts to write the value provided out to the storage file.
