// Code generated by go-bindata.
// sources:
// data/bindata.go
// data/customers.json
// data/locations.json
// data/profiles.json
// data/rates.json
// DO NOT EDIT!

package data

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

var _dataBindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func dataBindataGoBytes() ([]byte, error) {
	return bindataRead(
		_dataBindataGo,
		"data/bindata.go",
	)
}

func dataBindataGo() (*asset, error) {
	bytes, err := dataBindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/bindata.go", size: 0, mode: os.FileMode(438), modTime: time.Unix(1569716771, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataCustomersJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xe6\xe5\x52\x50\xa8\x06\x11\x0a\x0a\x4a\x99\x29\x4a\x56\x0a\x86\x3a\x50\x5e\x62\x69\x49\x46\x48\x7e\x76\x6a\x9e\x92\x95\x82\x52\x98\xa3\x8f\xa7\x4b\x7c\x88\xbf\xb7\xab\x9f\x12\x48\xbe\x96\x97\x2b\x16\x10\x00\x00\xff\xff\x8b\xaa\x5d\xfe\x3c\x00\x00\x00")

func dataCustomersJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataCustomersJson,
		"data/customers.json",
	)
}

func dataCustomersJson() (*asset, error) {
	bytes, err := dataCustomersJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/customers.json", size: 60, mode: os.FileMode(438), modTime: time.Unix(1569427793, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataLocationsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xe6\xe5\x52\x50\xa8\x06\x11\x0a\x0a\x4a\x19\xf9\x25\xa9\x39\x9e\x29\x4a\x56\x0a\x4a\x86\x4a\x3a\x50\xc1\x9c\xc4\x12\x25\x2b\x05\x53\x43\x3d\x53\x03\x23\x4b\x73\x63\xb8\x70\x7e\x9e\x92\x95\x82\xae\x81\x9e\xa1\xa1\x89\xb9\x91\x31\x48\xb4\x56\x07\x87\x61\x46\xd4\x34\xcc\x98\x9a\x86\x99\x50\xd3\x30\x53\x6a\x1a\x66\x46\x8e\x61\xbc\x5c\xb1\x80\x00\x00\x00\xff\xff\xd2\xf6\x22\x67\xd1\x01\x00\x00")

func dataLocationsJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataLocationsJson,
		"data/locations.json",
	)
}

func dataLocationsJson() (*asset, error) {
	bytes, err := dataLocationsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/locations.json", size: 465, mode: os.FileMode(438), modTime: time.Unix(1569427817, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataProfilesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x96\xdf\x6e\x5b\xb7\x0f\xc7\xef\x0b\xf4\x1d\x08\x5f\xfd\x7e\x80\x5d\xf8\xcf\x39\x89\xd3\xbb\x38\x5d\xd3\x01\xf3\x10\xd4\xe9\x86\x6d\xd8\x05\x2d\xd1\x16\x17\x59\xf2\x28\x9e\xba\x07\x43\x81\xbd\xc6\x5e\x6f\x4f\x32\xe8\xd8\xce\x9f\xb9\x27\x99\xaf\xd2\x9b\x04\x3e\xa4\x24\x8a\x9f\x2f\x49\xfd\xf2\xf2\x05\xc0\x1f\xf9\x0f\x40\x87\x6d\xe7\x35\x74\x06\x9d\xee\xee\x77\xc0\x15\xe5\x2f\x17\x9e\x17\x0a\xef\xa2\x92\xbf\xb5\xad\x5d\x0c\xf4\x7d\xb5\x9a\x93\x64\x97\xff\x15\x83\xf2\xff\x70\x7a\x5a\xf6\x8a\xd3\x7e\xff\xd6\xcb\x52\x32\xc2\x6b\xe5\x18\xb2\xd7\x39\x9c\xf4\x56\x1c\x2a\x25\xd8\xa0\xbf\x81\x85\xc4\x15\x7c\x08\x1c\x03\xcc\x7e\xaf\x50\x08\x30\x58\x28\x60\xeb\x93\xb6\x76\x84\x69\x15\x18\xa6\xa4\x12\x21\x29\xe6\xcd\xba\xa0\x8e\x13\xf8\xea\x53\x25\x35\xb8\x1c\x19\x58\x4a\xbc\x0c\x64\x61\x5e\xc3\x95\x63\xcf\xeb\x35\xc1\x4c\x51\xcc\x0d\x2c\x08\xb5\x12\x4a\x80\x01\x50\x34\xd5\xb0\xa8\x24\x70\xfe\x06\x26\x7a\x4f\x26\xef\x0a\x1c\x40\x1d\x81\x8f\xf3\x79\xdd\x05\x0e\xc6\x57\x96\xc3\x12\x36\x51\x6e\xf2\xb6\x33\xf4\x1f\xd1\x46\x81\x37\xe8\xf9\xd5\xed\x2d\xd1\x5a\xa1\x94\x3a\xaf\xf7\x99\x04\xe8\x24\x15\x22\xbd\x4b\x50\x71\x56\xee\xfd\xef\xac\xbb\xfc\x5e\x12\x4a\x0d\x33\xbd\xe7\x60\x58\xeb\x6c\x9a\x61\x80\xb7\x82\xc1\x70\x32\xf1\xc1\x06\xa8\x5b\x36\xe7\xf7\x57\xc5\x2a\xa8\x34\x0b\x3f\x04\x56\xb2\xf9\xfe\x4a\xe9\x9e\xcb\x3a\x26\x45\x7f\x11\x6d\xb3\xfa\xac\x18\xf4\x87\x9d\xad\xf1\x73\xfe\xf7\xb9\x7b\x28\x88\xe1\x81\x20\x7e\x84\x2f\xc6\xd5\x26\x8a\xd3\x5e\x39\x6a\x17\xc5\x77\x94\x12\xa8\xcb\x64\x60\xee\xa3\xd9\xc9\x22\x83\xf8\x89\x64\x8e\x30\xa9\x28\x20\x5c\x50\x50\x12\x58\x44\x69\x4c\xe7\xa2\x69\xa7\x02\x15\x0a\x76\xaf\x02\x4e\x80\x30\x18\x3e\x2e\xb3\xe3\xd0\x0d\xc6\x83\x56\x74\x23\xb1\xcf\x06\x6e\xf4\x14\xb8\xd1\x01\xb8\xa6\x86\xe1\x67\x52\xc5\xc7\xa1\x95\xc5\xa8\x37\x2e\xcb\xb2\xbd\x92\x47\x87\x29\xce\x5c\xae\xe2\x86\xbc\x87\x59\x93\x25\x30\x38\xf7\xd4\x33\x28\xa0\x95\x04\x94\x58\x05\xdb\xd4\xf8\xe4\xfc\xfd\x35\x08\xb2\xff\x57\x45\x3b\x5e\xef\x40\x9e\x3d\xec\x02\x0f\xba\x84\x89\xab\x39\x07\xca\xee\x4b\xd7\x53\x32\x0e\x7c\xb4\xcb\xa6\x54\x59\xdd\xae\xc4\x35\x56\xc6\x51\x3a\x0e\x76\xd9\x5e\xa6\xa5\xba\xaf\x97\x75\xd1\xc2\xfa\x07\x56\xf4\xf4\x38\xec\xe1\xe9\xb8\x37\x7a\xa4\x6d\x5f\x67\x32\x1b\x54\x92\x85\xc4\xa0\x3b\x40\x4d\xa6\x27\x58\xc3\x44\xd8\x2e\x09\x3e\x32\x6d\x52\xae\xbf\xd1\xb6\x8c\xd3\x9d\x28\xde\x72\xc8\xd9\x41\x0f\x6f\x38\xa9\xb0\xd1\x46\x04\x08\xc5\x97\x45\xf4\x96\x44\x6a\x98\x54\xec\x73\xfb\x3d\x0e\xe0\xb8\x95\xdf\x94\x53\x6a\x44\xf4\x4c\x0c\xcb\xa7\x18\x96\x07\x0c\xaf\x5c\xa4\xc0\x9f\xfe\xdb\xec\x3d\xe9\x0d\x46\xe3\xf6\x36\x1b\x0d\xe6\x40\x77\x33\xee\x9a\x82\x25\xf1\x91\x03\x04\xe2\xa5\x9b\x47\x71\x31\xda\x6e\xee\x9e\xfd\x43\x2a\xd8\x5a\xb1\xd2\xcc\xe5\x55\xd4\x28\x4d\x11\x12\x38\x4c\xe0\x62\xca\x87\xad\x30\xd4\x20\xb9\xa7\xaf\xaa\xc4\x86\x31\xa4\x86\x7c\x54\x47\x02\x86\x3c\xcd\x85\x95\x29\x41\xe2\x60\xa8\x89\x6c\x70\x56\xf6\xd3\x2b\xf8\x56\xff\xfe\xf3\xaf\xd4\xae\x11\xc7\x49\xa3\xb0\x81\x4b\x21\x54\x38\x5f\x91\xb0\xc1\x00\xd3\x7c\x10\xbc\x43\xef\x21\xf0\xd2\xa9\xf1\xd5\xfc\x38\x05\x9d\xf4\xdb\xfb\xfd\x37\xd6\x3e\xdf\xa4\x3e\x7b\x4a\x40\x27\x07\x02\xba\x76\xf9\x0d\xf4\x0a\xde\xd3\x92\xd3\x31\x53\x7b\x38\x2e\x7a\x45\xbf\xbd\x27\xdc\x6d\x3a\xad\x12\x55\x2b\xb8\x8e\x1b\x92\xed\xf8\x2d\x86\xbd\xcc\xa6\xee\x42\x31\x2e\x60\xa1\x90\x6e\xea\x64\x04\xd7\xd9\x61\x2b\xbf\x59\xac\xd4\x41\x5c\xc0\x14\xe5\x86\x14\xec\xbe\x2f\xc4\xc5\xc3\x28\xbb\x70\x81\x9e\x17\x51\x02\x63\x17\xd0\xfe\x86\x86\x82\x82\xc6\x07\x8f\x83\x4b\x14\x4b\x21\x75\x61\x1a\x93\x89\x81\x76\xaf\x85\x2e\x5c\xa1\x99\xe4\x79\xb4\xef\x25\x8d\xfa\x9a\x00\xee\x1f\xb2\xbf\x43\x8e\x27\x5a\x92\x90\x9f\x17\x47\x3e\x12\x86\xed\x83\x63\x24\xf6\x6b\x10\xcc\xcb\x17\xbf\xfe\x13\x00\x00\xff\xff\xc3\x78\xf5\x2b\xef\x0b\x00\x00")

func dataProfilesJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataProfilesJson,
		"data/profiles.json",
	)
}

func dataProfilesJson() (*asset, error) {
	bytes, err := dataProfilesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/profiles.json", size: 3055, mode: os.FileMode(438), modTime: time.Unix(1569427852, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataRatesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x4f\xcb\x82\x40\x10\x87\xef\x82\xdf\x61\xd8\xb3\xca\xae\xef\x1b\x61\xb7\x28\x08\x11\x3a\x48\xb7\xe8\xa0\xee\x52\x4b\xdb\x8e\xe8\x1a\x54\xf8\xdd\x63\xd5\xb6\x08\xba\xcc\xc0\xf3\xfc\xe6\xcf\xde\xf7\x00\x1e\xb6\x00\x90\x13\x1a\xa1\x52\x4e\x16\x40\x18\x09\x26\x58\x21\x17\x96\xe4\xcb\x55\xe6\xa0\xd4\xeb\xc2\x0c\x38\xa6\x6c\x16\xd2\xff\x90\x26\x4e\x62\x67\xbe\x2d\xa3\xce\x36\x88\x97\xdd\xad\xb6\x7a\xba\x0b\x40\x4a\xc4\x73\x51\x2a\x91\x8f\x73\x8c\x26\x11\xa5\x81\xd3\xaf\x1f\xb2\xed\x86\xbc\x29\x17\x6d\xd5\xc8\xda\x48\xd4\x83\x94\xfa\x08\xad\xbc\x0b\x0e\xa5\xe0\x1f\x39\x83\xa6\x50\x3f\x36\x3b\x97\xea\x4a\x75\xad\xbc\x0e\xa1\xf8\x2f\x62\xf3\x31\xd3\xdb\xd6\xfb\xde\xe1\x19\x00\x00\xff\xff\x8a\xce\x6d\x63\x2a\x01\x00\x00")

func dataRatesJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataRatesJson,
		"data/rates.json",
	)
}

func dataRatesJson() (*asset, error) {
	bytes, err := dataRatesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/rates.json", size: 298, mode: os.FileMode(438), modTime: time.Unix(1569427870, 0)}
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
	"data/bindata.go": dataBindataGo,
	"data/customers.json": dataCustomersJson,
	"data/locations.json": dataLocationsJson,
	"data/profiles.json": dataProfilesJson,
	"data/rates.json": dataRatesJson,
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
	"data": &bintree{nil, map[string]*bintree{
		"bindata.go": &bintree{dataBindataGo, map[string]*bintree{}},
		"customers.json": &bintree{dataCustomersJson, map[string]*bintree{}},
		"locations.json": &bintree{dataLocationsJson, map[string]*bintree{}},
		"profiles.json": &bintree{dataProfilesJson, map[string]*bintree{}},
		"rates.json": &bintree{dataRatesJson, map[string]*bintree{}},
	}},
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

