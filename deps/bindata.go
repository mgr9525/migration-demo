package deps

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __000001_init_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x28\x49\x4c\xca\x49\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x28\x89\x2f\x2d\x4e\x2d\x4a\xb0\x06\x04\x00\x00\xff\xff\x1c\xb3\xee\x9d\x1e\x00\x00\x00")

func _000001_init_down_sql() ([]byte, error) {
	return bindata_read(
		__000001_init_down_sql,
		"000001_init.down.sql",
	)
}

var __000001_init_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xd0\xc1\x6b\xc2\x30\x14\x06\xf0\xbb\xe0\xff\xf0\x8e\x2d\xec\x50\xcb\x0a\x83\xd1\x43\x5a\xdf\x5c\x58\x8d\x23\x26\x07\x4f\xcd\xb3\x66\x2e\x6c\xc6\x91\x46\xff\xfe\xd1\x09\x13\xa4\xe0\xed\x1d\x7e\xef\xe3\xe3\xab\x25\x32\x85\xa0\x58\xd5\x20\x98\xd8\x9e\x7a\x1b\x0c\x24\xd3\x09\x80\x71\x3b\x03\x5b\xb7\x77\x3e\x26\x79\x96\x82\x58\x29\x10\xba\x69\x80\x69\xb5\x6a\xb9\xa8\x25\x2e\x51\xa8\x87\x3f\xeb\xe9\x60\x0d\x9c\x29\x74\x9f\x14\x92\x59\x96\xa5\x30\xc7\x17\xa6\x9b\xcb\xcf\x05\xfd\x50\xdf\x5f\x51\x5e\x14\x63\xc8\xbb\xee\xeb\x6e\x12\x9d\x29\x52\xb8\xb2\x62\x9c\x75\xc1\x52\xb4\x6d\x74\x43\xb9\x1d\x45\x3b\x5c\x23\xee\xfb\xb8\x77\xfe\x2e\x7b\x97\x7c\xc9\xe4\x06\xde\x70\x03\xc9\xb0\x4e\x0a\x7a\xcd\xc5\x02\x2a\x25\x11\xa7\x93\x14\x50\x2c\xb8\xc0\x92\x7b\x7f\x9c\x57\x37\x3b\x95\xb3\xfc\x3f\xb2\x7e\x65\x72\x8d\xaa\x3c\xc5\x8f\xa7\xc3\xf6\xf1\xf9\x37\x00\x00\xff\xff\x14\x3b\x70\xcc\x86\x01\x00\x00")

func _000001_init_up_sql() ([]byte, error) {
	return bindata_read(
		__000001_init_up_sql,
		"000001_init.up.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"000001_init.down.sql": _000001_init_down_sql,
	"000001_init.up.sql":   _000001_init_up_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"000001_init.down.sql": &_bintree_t{_000001_init_down_sql, map[string]*_bintree_t{}},
	"000001_init.up.sql":   &_bintree_t{_000001_init_up_sql, map[string]*_bintree_t{}},
}}
