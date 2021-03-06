// Code generated by go-bindata.
// sources:
// assets/class.yaml
// assets/clusterrole.yaml
// assets/clusterrolebinding.yaml
// assets/deployment.yaml
// assets/role.yaml
// assets/rolebinding.yaml
// assets/serviceaccount.yaml
// DO NOT EDIT!

package generated

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

var _assetsClassYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcb\x41\x0a\xc2\x40\x0c\x85\xe1\xfd\x9c\x22\x17\x70\x8a\x3b\xc9\xd6\x23\x08\xee\x1f\x34\xad\xa1\x76\x32\xe4\x0d\xf5\xfa\x22\xd2\xf5\xff\x7f\x9b\xb7\x59\xe5\x31\x22\xb1\xda\xfd\x0d\xb2\xa0\xfb\xd3\x92\x1e\x4d\x85\xff\x50\xb7\x1b\xab\xc7\x74\x5c\xcb\x6e\x03\x33\x06\xb4\x88\x34\xec\xa6\x82\x0f\x2f\xb6\xb0\xf4\x8c\xc3\x7f\xcc\x52\x25\xba\x35\xbe\x7c\x19\x35\x72\x9d\xce\xe5\x1b\x00\x00\xff\xff\x8b\xd1\x6d\x3b\x6e\x00\x00\x00")

func assetsClassYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsClassYaml,
		"assets/class.yaml",
	)
}

func assetsClassYaml() (*asset, error) {
	bytes, err := assetsClassYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/class.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsClusterroleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x31\x6f\xdb\x40\x0c\x85\x77\xfd\x8a\x83\xe6\x4a\x45\xb7\x42\x6b\x87\x6e\x1d\x3a\x74\x29\x3c\xd0\xa7\x67\x8b\xf0\xe9\x78\x20\x79\x4a\x9c\x5f\x1f\x48\x11\x10\x20\xb1\x01\x23\xd3\x3d\x1c\x1e\xdf\xc7\xc7\x0b\xe7\x71\x08\xbf\x52\x35\x87\xfe\x95\x84\x86\x0a\xff\x83\x1a\x4b\x1e\x82\x1e\x29\xf6\x54\x7d\x12\xe5\x17\x72\x96\xdc\x5f\x7e\x5a\xcf\xf2\x7d\xf9\xd1\xcc\x70\x1a\xc9\x69\x68\x42\xc8\x34\x63\x08\x38\x59\x57\x54\x16\x5e\xa7\xa1\x9d\xd6\x9c\xa1\x8d\xd6\x04\x5b\x5d\x5d\xa0\xc2\xbf\x55\x6a\xb1\x21\xfc\x6f\xdb\x43\x13\x42\x08\x0a\x93\xaa\x11\xdb\x5f\x59\xd1\xe6\xc8\xbe\x48\xaa\x33\x6c\x37\x2d\xd0\xe3\x66\x38\xc3\xdb\x6f\xa1\x4d\x6c\xdb\xfb\x44\x1e\xa7\x55\x44\x05\x39\x56\x35\x22\xc1\xb1\xcd\x7d\x01\x18\x13\xf1\xfc\x18\xf5\x06\xc1\x5c\x94\xce\xd8\xaf\x74\x8b\xb7\x3b\x62\x22\xb3\x07\xdb\x3d\xd8\x04\x0b\xb2\x7f\x4a\x7c\xbf\x4b\x2d\xe3\xae\xca\xdd\xf5\x11\xab\xb2\x5f\x7b\x29\xc8\x36\xf1\xc9\xef\xb5\xd8\x8d\x51\xb2\xe3\xd9\xa3\x64\x73\x25\xbe\xc1\xaf\x86\x0f\x01\x7f\x68\x7e\x0b\x99\xc4\x7c\x96\x9a\xbd\xa3\x7c\xad\x3c\xb6\x87\xe6\x35\x00\x00\xff\xff\x99\x4a\xbc\xee\x8f\x02\x00\x00")

func assetsClusterroleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsClusterroleYaml,
		"assets/clusterrole.yaml",
	)
}

func assetsClusterroleYaml() (*asset, error) {
	bytes, err := assetsClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsClusterrolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcf\x3d\x6e\xc3\x30\x0c\x05\xe0\x5d\xa7\xe0\x05\xe4\xa2\x5b\xa1\xad\xed\x90\xdd\x01\xb2\xd3\x32\x9d\x30\xb6\x49\x81\x92\x3c\xe4\xf4\x81\x61\x78\xc9\xdf\xfc\xc8\xef\xe1\x8d\x2c\x7d\x80\xff\xa9\xe6\x42\xd6\xea\x44\x7f\x2c\x3d\xcb\xd9\x61\xe2\x13\x59\x66\x95\x00\xd6\x61\x6c\xb0\x96\x8b\x1a\xdf\xb0\xb0\x4a\x33\xfe\xe4\x86\xf5\x6b\xf9\x76\x33\x15\xec\xb1\x60\x70\x00\x82\x33\x05\xb0\x2a\x9e\x86\xec\x93\xe9\xc2\xab\x40\xe6\x72\xed\xae\x14\x4b\x5e\xaf\x3c\x6c\xad\x47\xb2\x85\x23\xfd\xc6\xa8\x55\x8a\x03\xd8\x81\xc7\xe7\x3d\xc9\x09\xe3\x8b\xd8\x74\xa2\x96\x86\x95\x7e\x9a\xe3\xde\x98\xde\xaa\x6c\x34\x26\x3e\x98\xd6\xf4\x61\xa6\xbb\x07\x00\x00\xff\xff\xa3\xb7\x8a\x39\x26\x01\x00\x00")

func assetsClusterrolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsClusterrolebindingYaml,
		"assets/clusterrolebinding.yaml",
	)
}

func assetsClusterrolebindingYaml() (*asset, error) {
	bytes, err := assetsClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\xcb\x8e\xdb\x3c\x0c\x85\xf7\x7e\x0a\xbe\x80\x6d\xe4\xff\x7b\x83\x76\x01\xc6\x2d\x0c\x4c\x92\x41\x52\x4c\xd1\x95\xc1\x6a\xe8\x44\xa8\x6e\x15\x69\x4f\xd3\xa7\x2f\xd4\x4c\x5a\x27\xe3\x68\x65\xc1\xdf\x39\x87\xa4\xf8\xdd\xf8\x27\x05\x77\x14\x6d\x38\x3a\xf2\x52\x60\x34\x8f\x94\xd8\x04\xaf\x00\x63\xe4\x7a\x5c\x14\x8e\x04\x9f\x50\x50\x15\x00\x1e\x1d\x29\xa0\x9e\xcb\x98\xc2\x68\x32\x48\xa9\x00\xb0\xf8\x8d\x2c\x67\x02\xb2\xee\x35\xc2\x91\x74\xfe\x9d\x28\x5a\xa3\x91\x15\x2c\x0a\x00\x26\x4b\x5a\x42\x3a\x09\x1d\x8a\x3e\xdc\x4f\x9c\x6e\x78\x01\xb0\x24\x14\xda\x1f\x4f\x98\x1c\x23\x29\xd8\x92\x4e\x84\x42\x50\x00\x08\xb9\x68\x51\xe8\xc5\x77\xd2\x40\x3e\xf6\x22\xe2\x66\x08\xc0\xb9\xe8\x3f\xdf\x94\x46\xa3\x69\xa9\x75\x18\xbc\xcc\xf3\x00\x3a\x78\x41\xe3\x29\x4d\xec\xcb\x9b\x53\x3b\x1f\xe3\x70\x4f\x0a\x7e\x0c\x78\xac\x4c\xa8\xe9\xa7\x50\xf2\x68\x3b\x96\x90\x70\x4f\xf5\x95\x52\xe5\xd6\x58\x26\x06\xe4\x47\x35\xb9\xfe\xcb\xfc\xd8\xde\x37\xdd\xee\xeb\xee\x73\xb3\xea\xda\xbb\x0b\x04\x60\x44\x3b\x90\x82\x9e\xcb\xc5\x7f\xff\xbf\x79\xfb\xee\xfd\x87\x59\x8f\xe5\x97\x5d\xb7\x6d\x3e\xb5\x9b\xf5\xbc\x7e\xe0\x92\x90\xa5\x5c\xcc\xaa\x1f\xb6\x9b\xc7\x76\xd7\x6e\xd6\xcd\xb6\x5b\x2f\x57\xcd\xbc\x47\x88\xe4\xf9\x60\x7a\xa9\x42\xda\xd7\xf8\xcc\x25\xf5\x3c\x41\xc7\x60\x07\x47\xab\x3c\x7a\x9e\xef\x34\x8e\xe5\x09\xba\x0a\x70\x59\xf3\x80\x72\x50\x50\xc7\xbc\xda\x2c\xe4\xe5\x84\x9e\x13\x5e\x6e\xaf\x9f\x6c\xce\xd4\xf7\x57\x15\xe4\xcd\xa0\x74\x31\xc7\x8a\x7a\xae\xfe\xce\xa5\x42\x87\xbf\x82\xc7\x67\xae\x74\x70\x17\xda\x78\xab\xb0\xdf\x01\x00\x00\xff\xff\x03\x53\xd7\x2f\x9a\x03\x00\x00")

func assetsDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDeploymentYaml,
		"assets/deployment.yaml",
	)
}

func assetsDeploymentYaml() (*asset, error) {
	bytes, err := assetsDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xcd\x31\x4e\x04\x31\x0c\x05\xd0\x3e\xa7\xb0\x52\xef\x0c\xa2\x43\xb9\x00\x3d\x05\x0d\xda\xc2\x9b\x7c\x76\xad\xc9\xc4\x91\x93\x0c\x12\xa7\x47\x13\x71\x83\xed\x2c\xff\xaf\xf7\x37\x29\x29\xd0\x87\x66\x38\xae\xf2\x09\x6b\xa2\x25\x90\xdd\x38\xae\x3c\xfa\x43\x4d\x7e\xb9\x8b\x96\x75\x7b\x6b\xab\xe8\xcb\xf1\xea\x76\x74\x4e\xdc\x39\x38\xa2\xc2\x3b\x02\x65\x70\x82\x2d\x59\xe3\x26\xe5\xbe\xe0\xbb\x2d\xd5\xf4\x90\x13\x83\x39\x1b\x19\xed\x6c\x2f\xc4\x55\xde\x4d\x47\x6d\x81\xbe\xbc\xbf\x3a\x22\x22\x43\xd3\x61\x11\xf3\x87\x92\xaa\x4a\xe9\xed\x3f\x3c\x60\xb7\x19\xdc\xd1\xfd\x85\xfc\xa8\x89\x3b\xce\xab\x72\x8f\x8f\xd9\x7a\x86\xcd\xd2\xa6\xfb\x33\xb5\x0b\xf9\x68\x38\x07\xae\xee\x2f\x00\x00\xff\xff\x55\x98\x00\x29\x1b\x01\x00\x00")

func assetsRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRoleYaml,
		"assets/role.yaml",
	)
}

func assetsRoleYaml() (*asset, error) {
	bytes, err := assetsRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsRolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x31\x6e\xc3\x30\x0c\x45\x77\x9d\x82\x17\x90\x8b\x6e\x85\xb6\x76\xe9\xee\x02\xdd\x69\x89\x76\x18\xcb\xa4\x40\x49\x1e\x72\xfa\xc0\x08\x82\x00\x19\x82\x1c\xe0\xbd\xf7\xff\xca\x92\x02\x8c\x9a\xe9\x87\x25\xb1\x2c\x0e\x0b\xff\x93\x55\x56\x09\x60\x13\xc6\x01\x7b\x3b\xa9\xf1\x05\x1b\xab\x0c\xeb\x57\x1d\x58\x3f\xf6\x4f\xb7\x51\xc3\x84\x0d\x83\x03\x10\xdc\x28\x40\x26\x4c\x64\x3e\x6b\x5c\x59\x16\x4f\x73\xf5\xc5\x74\xe7\x43\x46\xe6\x6a\x9f\xce\x14\x5b\x3d\x00\x0f\xb7\xf2\x1f\xd9\xce\x91\xbe\x63\xd4\x2e\xcd\x01\xdc\x5d\xcf\xb0\x69\xa6\x91\xe6\x83\x7d\x6c\x7e\xb7\x0c\x80\x85\x7f\x4d\x7b\x79\xf1\xc9\x5d\x03\x00\x00\xff\xff\x1d\xa0\x9b\xc4\x0c\x01\x00\x00")

func assetsRolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRolebindingYaml,
		"assets/rolebinding.yaml",
	)
}

func assetsRolebindingYaml() (*asset, error) {
	bytes, err := assetsRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x2c\xc8\x0c\x4b\x2d\x2a\xce\xcc\xcf\xb3\x52\x28\x33\xe4\xca\xce\xcc\x4b\xb1\x52\x08\x4e\x2d\x2a\xcb\x4c\x4e\x75\x4c\x4e\xce\x2f\xcd\x2b\xe1\xca\x4d\x2d\x49\x4c\x49\x2c\x49\xb4\xe2\x52\x50\xc8\x4b\xcc\x4d\xb5\x52\x48\x4d\x2b\xd6\x2d\x28\xca\x2f\xcb\x04\x69\x4c\x2d\xe2\x02\x04\x00\x00\xff\xff\x34\x5b\x29\x3a\x46\x00\x00\x00")

func assetsServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsServiceaccountYaml,
		"assets/serviceaccount.yaml",
	)
}

func assetsServiceaccountYaml() (*asset, error) {
	bytes, err := assetsServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"assets/class.yaml": assetsClassYaml,
	"assets/clusterrole.yaml": assetsClusterroleYaml,
	"assets/clusterrolebinding.yaml": assetsClusterrolebindingYaml,
	"assets/deployment.yaml": assetsDeploymentYaml,
	"assets/role.yaml": assetsRoleYaml,
	"assets/rolebinding.yaml": assetsRolebindingYaml,
	"assets/serviceaccount.yaml": assetsServiceaccountYaml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"class.yaml": &bintree{assetsClassYaml, map[string]*bintree{}},
		"clusterrole.yaml": &bintree{assetsClusterroleYaml, map[string]*bintree{}},
		"clusterrolebinding.yaml": &bintree{assetsClusterrolebindingYaml, map[string]*bintree{}},
		"deployment.yaml": &bintree{assetsDeploymentYaml, map[string]*bintree{}},
		"role.yaml": &bintree{assetsRoleYaml, map[string]*bintree{}},
		"rolebinding.yaml": &bintree{assetsRolebindingYaml, map[string]*bintree{}},
		"serviceaccount.yaml": &bintree{assetsServiceaccountYaml, map[string]*bintree{}},
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

