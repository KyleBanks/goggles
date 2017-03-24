package assets

import (
	"github.com/elazarl/go-bindata-assetfs"
)

// FS returns the assets FileSystem.
func FS() *assetfs.AssetFS {
	return assetFS()
}
