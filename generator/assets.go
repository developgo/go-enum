// Code generated by go-bindata.
// sources:
// enum.tmpl
// DO NOT EDIT!

package generator

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _enumTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x4d\x6b\xdb\x40\x10\x3d\x7b\x7e\xc5\xab\x48\x41\xa2\xfe\x48\xae\x6e\xd5\x4b\xd3\x43\x2e\x49\x20\xa6\x17\x63\xca\xd6\x1a\xb9\x4b\xac\x95\x58\xcb\xc6\x66\x33\xff\xbd\x8c\x24\x13\x5b\x55\xa1\x37\xed\xce\xbc\xf7\xf6\xcd\x1b\x85\x30\x41\xc6\xb9\x75\x8c\xe8\x37\x9b\x8c\x7d\x24\x42\xb3\x19\xbe\x95\x19\x63\xc3\x8e\xbd\xa9\x39\xc3\xaf\x13\x36\xe5\x84\xdd\xbe\xd0\xe2\xfd\x13\x1e\x9f\x16\xf8\x7e\xff\xb0\xf8\x40\x54\x99\xf5\xab\xd9\x30\x42\x98\x76\x9f\x22\x44\xb6\xa8\x4a\x5f\x23\x26\x00\x88\xf2\xa2\x8e\x28\xa1\x10\xd8\x65\x98\x68\xfd\x52\x59\x79\x55\x77\x5d\xba\x9d\x42\x42\xc0\x8d\xde\x3d\x9a\x82\x31\x4f\x31\xd5\xc3\x54\x4f\x22\x5a\xf4\xc6\x6d\x18\x37\xfe\xc1\x65\x7c\x1c\xe3\xe6\x60\xb6\xfb\x8b\xc6\x1f\x7a\xdc\x41\x84\x46\xb3\x19\x42\x68\xeb\xd3\x67\xcf\xb9\x3d\x72\xd6\xf2\xc0\xee\x60\xb4\x78\x16\x12\x41\x99\xa3\x3e\x55\xfc\x0e\xe9\x24\xd5\xc2\xbf\x68\xae\x19\x52\xd8\xb2\x36\x8d\x39\x76\x99\x08\x25\xd4\xb9\xfa\x19\xc2\xa5\x8d\xc6\x5a\x8a\x28\x84\x5d\xed\xad\xdb\xd8\xfc\xd4\xbe\x1e\x22\x11\xd1\xc1\xf8\x3e\xa2\x31\x8b\x14\x21\x58\xfd\xba\x04\x10\xe5\x7b\xb7\x46\x6c\x71\x0d\x49\xf0\xd2\x90\xc7\x09\x5a\x15\x04\x1a\xd9\x1c\x16\x5f\x70\x8b\xb7\x37\x58\x7c\x4d\x7b\x98\x78\xcb\x2e\x1e\x92\x4e\x26\x77\x89\xe2\x47\x9e\xeb\xbd\x77\xc8\x8b\x7a\xfa\x52\x79\xeb\xea\x3c\x8e\x7a\x1c\x1f\xb3\x24\x1a\xc3\x26\x34\x12\x3a\xf7\x0f\x0c\x60\x39\xa4\xb3\xb4\xab\xf9\xf0\xfd\xa7\xbb\xd5\x8a\x64\x70\x38\x4d\xe4\x48\x51\x98\x6a\xd9\x5a\x5d\x9d\x1b\x16\xa7\x8a\x45\x42\x17\xe2\x7f\xee\xce\xa4\x4b\xbd\x59\xde\xde\x3a\x44\xf3\xf7\x65\x68\xda\x45\xc6\x1d\x7b\x9b\xf9\x39\x8f\x67\xe3\x77\xdc\x9b\x8c\xd3\xdc\xdb\x17\x26\xbd\xc9\xeb\x70\x0f\x66\xab\x6f\xe9\x81\x6e\x93\x26\xb6\xe3\x18\xe5\xab\x96\x87\xcc\x2f\x95\x79\xf5\x59\x3b\x34\x24\x25\xfa\x8b\xe7\x78\x15\xc8\xc1\x6c\xa9\xf9\x11\xdb\x57\xff\x09\x00\x00\xff\xff\xae\x5d\xdb\x3c\x0b\x04\x00\x00")

func enumTmplBytes() ([]byte, error) {
	return bindataRead(
		_enumTmpl,
		"enum.tmpl",
	)
}

func enumTmpl() (*asset, error) {
	bytes, err := enumTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "enum.tmpl", size: 1035, mode: os.FileMode(420), modTime: time.Unix(1502404542, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"enum.tmpl": enumTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"enum.tmpl": &bintree{enumTmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

