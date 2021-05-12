// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package i18n generated by go-bindata.// sources:
// locales/en-US/home.yml
// locales/zh-CN/home.yml
package i18n

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _localesEnUsHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x58\xcd\x72\xa3\x3a\x16\xde\xf3\x14\x0a\xae\xde\xf5\x74\xdd\x75\x76\x34\x26\x09\xd3\x36\x50\x80\x73\x6f\x66\x43\xc9\x70\x6c\x6b\x82\x25\x4a\x12\xc9\xf8\xee\xe6\xbd\xe6\x9d\xe6\x15\xa6\x8e\x24\x30\xd8\xce\xed\x74\xcd\xaa\x3b\x89\xce\x8f\xce\xcf\xf7\x7d\x62\x51\x8b\xe3\x51\x70\x2f\x09\xd6\x51\x15\xfd\x11\x17\x65\x71\x4f\xfc\x84\x1e\x81\xd0\x56\x02\x6d\x4e\x04\xfe\xc5\x94\x56\xbe\x17\x67\x55\x92\x96\xe7\x43\x59\x0b\x54\x01\xd9\xb1\xb6\x25\x8c\x13\x7d\x00\x22\x61\xcf\x94\x96\x27\x12\x67\x44\xd8\x5f\xa9\x93\xd2\x70\x24\x0a\xb4\x66\x7c\x4f\x3a\xba\x07\xdf\xf3\xbc\x45\xdd\xf6\x4a\x83\xf4\xc2\xd5\xa6\x28\xa3\xbc\x5a\x46\xab\xa8\x8c\xaa\x87\x20\x5e\x45\xcb\x7b\xe2\xd7\x94\x13\x2e\x34\x69\xa0\x05\x0d\xc4\x1d\xc7\x40\x75\x2f\x25\x70\x4d\x94\xa6\x1a\xfc\xd1\x41\x5c\x98\xf4\xf2\x4d\x92\xc4\xc9\xe3\x3d\xf1\xcb\xc3\xc4\x4c\x19\x67\xb2\xe7\x9c\xf1\xfd\x95\xd1\x2a\x0d\x83\xd5\x3d\xf1\xe3\x63\x27\xa4\x1e\xad\x6a\xca\xd1\x6a\x0b\xa4\xef\xf6\x92\x36\xd0\xf8\x98\xb9\x84\x06\xb8\x66\xb4\xf5\x66\x49\x57\x79\x54\xa4\x9b\x3c\x8c\xee\x89\xff\x40\x59\x0b\x0d\xd1\xc2\xe5\x7f\x47\xca\x03\x48\xc0\x3c\x28\x27\x54\x29\x51\x33\xaa\xa1\x21\x07\xa1\x34\xe9\x79\x03\x92\xe8\x03\x53\xe4\x15\x4e\xfe\x07\x6e\xab\x7f\xa4\xc9\x2f\xf9\xfe\x53\x70\xb8\xe1\xfb\x21\xd8\xac\xca\x2a\xcc\xa3\x65\x94\x94\x71\xb0\xaa\xc2\x20\x31\x55\xb0\x61\xef\x89\xbf\x84\x1d\xed\x5b\x4d\xce\x37\x9d\x94\xc2\x06\x6d\x7c\x3b\x32\xe1\x53\x14\xfe\x38\x77\xcd\xd4\xfc\x6c\xc5\x71\x8e\xce\xa6\x66\x1e\xcc\x68\x29\xf3\xff\x5e\x81\x34\x67\x7c\x2f\x0b\x8a\xe2\xf7\x34\x5f\x8e\xc9\x24\x9b\x15\x76\xa4\xa3\x4a\xbd\x0b\xd9\x90\x61\x1e\xb6\x40\xb6\x2d\xe5\xaf\xff\xfd\xcf\xbf\x7d\x2f\xcb\xe3\xe7\xa0\x8c\xaa\x1f\xd1\xcb\xa5\x21\x66\xd2\x49\xf6\x46\x35\xe0\xc5\x27\x59\x9c\xcd\xbd\x05\x96\xdf\x7b\x4a\x8b\xb2\x0a\x56\x79\x14\x2c\x5f\xce\xe3\xfd\x84\x9d\xb9\xdc\x01\xd7\x19\x63\x31\x5e\xfa\x66\x43\x6c\x67\xb1\x27\xce\xc5\xa4\x31\xef\x4c\x1f\x4c\x01\xdc\xa0\xdd\xf2\x5b\x7d\x7f\xa9\xb2\x3c\xfd\x7b\x14\x96\xff\x57\x88\x4e\x8a\x7f\x42\xad\x7d\xaf\x78\x29\xca\x68\x5d\xb9\x2d\x7e\x48\x37\xc9\xf2\xf6\x12\xb7\xa2\xa6\x2d\x6e\xf0\x8e\x49\xa5\x4d\xa1\x92\x14\xed\x82\xe7\x20\x5e\x05\xdf\x57\x38\x22\x89\x20\x71\x47\xe8\x1b\x65\x2d\xdd\xb6\xe0\x7b\x71\x61\xb7\xc8\xdc\x61\xb2\xbf\xcc\xae\x94\x75\x8a\x09\xfb\xb6\xde\xf1\x3a\x4b\xf3\xb2\x8a\xf2\x3c\xcd\x87\x9e\x25\x82\x34\x54\x53\xbc\xa6\x33\x7b\xa7\x8a\xec\x44\xcf\x9b\x3b\xe2\x32\xad\x0f\x50\xbf\x9a\x3c\xdd\x91\x1d\x6b\xe1\x6e\xee\x14\xdd\x55\xcf\xc1\x6a\x83\x99\x46\xc7\x4e\x9f\xac\x5f\xc1\x49\xcb\x38\x90\x2f\x6a\x7e\xfe\xf7\x3c\x4d\x1e\xab\x87\x34\x5f\x07\x98\x7a\xcc\x6b\x21\x25\xd4\x9a\xd8\x00\x42\x1e\xa9\xfe\xd0\x78\xb2\x48\xd3\xc2\x86\x93\x2d\x10\xda\x5e\xe2\x43\x1f\xae\xd1\xf3\xce\xd8\xc6\x7d\xc2\xda\x0d\x4c\xb2\x59\xdf\x13\x3f\x20\x5a\x68\xda\x12\xb1\x23\x5f\x14\x91\xe2\x5d\xe1\x7f\xcd\xf5\xa9\x04\x42\xb7\x1c\xaf\xd3\x7e\x25\xea\x95\x75\xce\xcf\x88\x30\xc3\xe8\xc5\xc9\x7c\xae\xb7\x8c\x3b\xa8\x92\xa0\x44\x2f\x6b\x4c\xe2\x2b\x91\x40\x95\xe0\xf7\x1f\xe4\x53\x04\xcf\x73\xb8\x52\xf4\xcd\xcd\xec\x85\xb1\xe7\x2d\x10\x0a\xce\x20\x80\x75\x58\x07\x65\xf8\x34\x2c\xf2\x80\x02\x4c\x11\x36\x74\xc7\xf7\xd2\x3c\x7e\x8c\x13\x57\xf8\xe9\x79\x21\xd9\x9e\x71\xda\x7e\x64\xb8\x29\xce\xd8\x1f\x84\x65\x6c\x12\x2d\x07\x44\x72\x64\x01\x1c\x27\x7b\x32\x79\x82\x6b\x5a\x6b\x33\x7b\xb4\x39\x32\x8e\x54\x47\xb5\x90\x77\xce\xe1\xb4\x7b\x89\x20\xaa\xaf\x0f\xc6\xa1\x59\xa1\x60\xb9\x8e\x93\x6b\xa8\xc5\xa0\x8d\x83\x5b\xe3\xd4\xa6\x70\x05\xb7\x77\xf3\xa4\xf3\x68\x15\x94\xd1\x72\x82\x10\x1b\x34\x3b\x50\x4c\x7d\x8a\x03\x6e\xfd\x4d\x0a\xab\x65\x90\x8d\x19\x6c\xb2\x65\x30\x66\xd0\x36\xb4\xbb\x0c\x0c\x0d\xb3\x71\x9f\xa3\x3c\x7e\x78\xa9\xc2\x74\x39\xa1\xe7\x67\x90\x6c\xc7\x6a\xaa\x99\xe0\xa4\x16\x0d\x10\x90\x52\x48\xdf\x8b\xd6\x41\xbc\xaa\x96\x71\xe1\x80\x62\x4d\x59\x3b\xb0\xbf\x32\x23\xd8\x30\xf5\xc9\xc2\x0e\xde\xa6\xed\x8d\x8e\xe8\xf0\x48\x75\x7d\x20\x3b\x33\x5a\x16\xa1\x90\x8c\xc6\xf9\x29\xf0\xa7\x31\x57\x2c\xcd\x5f\x30\xd1\x30\x23\x97\x4e\x0c\x34\xdd\x13\xff\x5d\x0a\xbe\x3f\x73\x15\x11\x72\x62\x62\x13\x34\xa4\x31\x26\x77\x49\x1a\xde\x02\x75\x91\xe0\x03\xca\xe7\xd1\x63\x9c\x26\x9f\x55\x0d\xc4\x1a\xff\x0c\xe7\x91\xec\x31\x14\xfe\x3b\x04\x42\xc1\xf0\xe9\x30\x46\x2d\xfc\x8c\x4c\x5a\xca\xe7\xe2\xc9\x02\x77\x68\x0b\xbb\x07\x3d\xa5\xb5\x1b\x98\x5d\x0b\xbe\x63\xfb\x5e\x9a\xb9\x31\x8d\x8b\xd7\xc1\x63\xf4\xb1\x2b\x76\xa4\x7b\xf8\x9c\xa3\xac\x2a\x9e\xd2\xdc\x02\xb8\xea\x77\x3b\x56\x33\x94\x89\x71\x87\x65\x11\x1d\x70\xa5\x69\xfd\xea\x3d\x46\xe5\xd0\x81\xa1\xc3\x89\x18\x8a\x6c\x80\x16\xcf\xbb\xbd\x59\xc3\x71\x0b\x72\x5c\xbd\x60\x89\xf3\xf4\x45\x91\x71\xdb\xb6\x00\x9c\xd0\xc6\x48\xc3\xe9\x82\x0e\x38\xf0\x45\xcd\x30\xc5\xf8\x77\xda\xc3\x85\x18\x15\xd9\x40\x02\x1f\xca\x31\x67\x70\x4b\x8b\x0d\xb6\x4f\x41\x51\xb9\xf6\x4c\x28\x64\xd2\xca\xb1\x35\xe1\x0d\x84\xf1\x16\x5c\x34\xe0\x25\xb8\xe9\x83\x1e\x72\x7a\xba\x2a\x83\xe2\x07\xd2\x4b\xd3\x10\x3c\x84\x5b\xe0\xa4\xb9\xf9\x71\x98\x1a\xa7\xb0\xbf\x76\xb6\x61\xef\x94\x69\xc2\x34\x69\x04\x87\x6f\x18\x60\x4b\xeb\xd7\xbe\x0b\xea\x5a\xf4\x5c\x7b\x59\x90\x07\xeb\x2a\x5a\x67\xe5\xcb\x65\xdb\x3a\x2a\xe9\x11\x34\x48\x85\xf2\xa3\xac\x8a\x4d\x96\xd9\xee\x6e\xb8\xea\x3b\x64\x66\x9c\xe1\x53\x87\x4f\x80\xb9\x08\x9d\x61\x93\xc5\x88\x51\x61\x7d\x0f\xc2\x1f\x9b\xac\x0a\xc2\x30\xdd\x24\xbf\xa2\xb5\x66\x89\x7f\x5a\x74\x79\x0b\x5c\x99\x0b\x41\xff\x89\x68\x68\xf5\x0b\x41\x5c\x57\xbf\x9b\x1c\x3d\x77\xc7\x87\x78\x15\x15\x53\x49\xec\x96\xc8\xb5\x4d\x8f\x97\x32\x12\x8a\x6c\x61\x27\x24\x10\xf5\xce\x74\x7d\xc0\xa7\xda\xe4\x00\xb5\xd7\x9e\xad\xbe\x8d\x72\xfd\xce\xda\x02\x1a\xa3\x21\x34\xa4\xef\xcc\xb0\x4f\xcc\xf2\xa8\x28\xd3\x3c\xba\xb6\x93\xa0\xb4\x90\x8c\xef\xe7\xeb\x91\x3b\xc1\x71\x5d\xc3\xc9\x35\x7f\x7a\xb9\xb3\x94\xbd\xad\xb4\xcf\x3b\x33\xea\xea\xa1\xf4\x5b\x68\x05\x52\x97\x16\x73\x6c\x2b\xf1\x49\x25\x3a\x90\x8e\x03\xc7\x7d\xea\x40\xa2\x64\x34\x1b\x65\x65\xd1\x15\x24\x3c\x39\xe1\x3e\x42\x82\xef\x8d\x1a\xcc\xc2\x4c\xe4\x4e\x0e\xd7\x37\xc3\x60\x61\x26\x5b\x05\xc9\x0d\x9f\x99\x1b\x99\x89\xcf\x8b\x61\xbf\xb6\xf9\x7e\x39\xd4\x13\x63\x6f\x81\xaa\xc0\x8a\x86\x11\xd9\xec\x20\x15\x27\x5e\x1f\xa4\xe0\xec\x4f\x6c\xb1\x02\x69\x89\xfd\x37\x27\x31\x56\xe9\x63\x9c\x5c\xda\x6c\xa6\xca\x0a\x89\xf1\xce\x9d\x3e\x4b\x85\xf2\xfc\xb9\xa0\x93\xe2\xc0\xb6\x4c\x2b\x82\x67\x5c\x8c\x9d\x14\x47\xd2\x8a\xfd\x1e\x07\x8c\xf1\x6f\x9f\x11\x66\xde\xa2\x66\xca\x0b\xe3\xc2\x00\xd8\x25\xaa\xe1\x3b\x85\x29\xa2\xa9\x7a\xbd\x44\x30\x34\x7d\x3b\x86\x86\x66\xbc\xe7\x75\x15\xa6\xc9\x43\xfc\x78\x7e\x1e\x86\x53\x02\xba\xa2\xfc\xb3\xc1\xe5\x77\x8d\xf2\x92\xbc\x3e\x1a\xb4\x06\xba\x56\x9c\x8e\x06\x0b\x5b\xca\x3f\x39\x70\xde\x82\x75\x48\x85\x63\x9e\x18\x0f\xb8\x06\x09\x0d\xbe\xe9\x14\xec\x8d\x4b\x4c\xa1\x65\xb5\x56\x67\x3c\x31\xb9\x63\x69\xcf\xc7\xbe\x92\x6e\xfe\x3e\xa4\x7b\xca\xf8\xf8\x29\x68\xfa\x20\x8c\x33\x7c\x6f\x60\x86\x75\xdd\x77\x0c\x1a\x42\x79\x33\x49\x52\x82\xf1\xd4\x18\xe3\x38\x79\x0e\x56\x31\x96\x23\xee\xac\x34\x7f\xa3\x2d\x6b\x06\x1e\x9f\x3c\xe1\x79\x8f\x04\x8c\x4f\x98\x3d\x70\xbc\xba\xbd\x06\x6d\x1a\x09\x4a\x81\x89\xf8\xdb\xb7\x6b\x7d\xa0\x34\x95\xe6\x32\xee\xa4\xc9\x46\xf5\x5b\x0e\x38\x78\xa6\x4c\x7f\xeb\x84\x68\x31\x5c\x96\xa6\xab\x9b\x7d\x8a\x33\x82\x67\x26\x44\x7f\x03\x86\xc7\x37\xb0\x55\x50\xf3\x5b\x8f\x04\x6d\x35\xa0\xd2\xf2\xe4\xa1\xf8\x28\xca\xfc\xe5\xfa\x93\x43\x39\xfd\x84\x26\x76\xf6\xab\x0d\x95\xf5\x81\x69\xa8\x75\x2f\xc1\xf0\xdd\x0d\x7d\xe9\x40\x69\xc4\xca\x01\x6f\xb3\x3c\x7d\x8e\x97\x51\x3e\x2a\x9d\x29\xe6\xd6\x12\xcc\x4d\xb0\xad\xbd\x16\x47\xaa\x59\x4d\x8e\x48\xe6\x2e\xff\x23\xe5\x3d\x6d\xdb\x13\xfe\x92\xed\x4e\xf3\x07\xa0\x9a\x40\x56\xf9\x92\x45\xb3\x10\x06\xa9\x6c\xb2\x6e\xe5\x47\xca\xbe\x23\x29\x6f\x4f\xc3\xcf\x8a\x20\x42\x7e\x25\x73\xa4\xfa\x39\x5b\x9f\xf1\x7a\xf7\x17\x6c\x4d\x0c\x0e\xc0\xfe\x16\x7f\x22\x53\xcd\xc8\xed\xcc\x5d\xb4\xd6\xe6\xbb\x91\xdd\x27\x05\x4a\xa1\x78\x2f\xa2\xa2\x40\xd1\xf8\x23\x7a\x99\xc1\xa8\xfb\xbb\xf9\xc8\xe4\xe6\xd8\x49\xc1\x2c\x4f\x91\x9c\x26\xd3\x3e\x9c\xb5\x4f\x2d\xc1\xdf\x40\xaa\xa9\x48\x31\x66\xa8\xe1\x92\x74\x2a\xd8\xfb\xc9\xdb\x6e\xa8\xbf\x19\x62\x64\x4c\xba\x07\xcf\xf4\x19\xb3\xc3\x56\xff\x51\xc4\x28\x2c\x0a\xfb\x37\x84\xd2\x37\xd6\x80\xbc\x1c\x9b\xff\x05\x00\x00\xff\xff\x00\x88\x06\x1f\xfa\x15\x00\x00")

func localesEnUsHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesEnUsHomeYml,
		"locales/en-US/home.yml",
	)
}

func localesEnUsHomeYml() (*asset, error) {
	bytes, err := localesEnUsHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/en-US/home.yml", size: 5626, mode: os.FileMode(420), modTime: time.Unix(1620699134, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesZhCnHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x58\x5b\x53\xdb\x48\x16\x7e\xd7\xaf\x70\xd9\x35\x6f\xb3\x5b\xfb\x3c\x6f\x8a\x2d\x40\x1b\x59\x52\x49\x32\x33\xec\x8b\x8a\x21\xde\x2c\x1b\xb0\x29\x0c\x5b\x35\xfb\x14\x73\xb3\x4d\x6c\xec\x49\x80\x0d\xe0\x04\x4c\x20\x78\x20\xbe\x00\x13\xec\xf8\x02\x7f\x46\xdd\x92\x9f\xf8\x0b\x5b\xdd\x2d\xb5\xdb\x76\x48\x78\x49\x55\xac\xf3\x9d\x3e\x7d\x2e\xdf\xf9\x9a\xc0\x4c\x7c\x7e\x3e\x1e\xe3\x64\x3e\x2c\x98\xc2\x2f\xa2\x6e\xe8\x3f\xf9\xfc\xa0\x90\xb3\xcf\xea\xa0\x71\x05\x2a\x6f\x41\xb1\xec\xe7\x44\xd5\x94\x15\xa3\x6f\xe0\xd4\x1a\xa0\x58\xb6\xaf\xdb\x76\xfb\xd0\xa9\xde\xda\xdd\x6a\xaf\xf4\xb9\xf7\xee\x18\x94\x2e\xc0\xc6\x9e\xd5\x7e\x03\x5a\x6f\x44\xd5\xcf\x71\x81\x99\xb9\xe5\xc4\x52\x74\x91\x0b\x4a\x11\xdd\x10\x34\x33\x24\x48\x82\x21\x98\x63\xbc\x28\x09\xa1\x9f\x7c\x7e\xf8\xbf\x23\x78\xbd\x03\xd2\x47\xbd\xbd\x13\xd0\x7d\x03\x32\x39\x7b\xf3\x06\xbe\x4c\xda\xfb\x6b\xbd\x83\x0d\xfb\xf6\xc4\x4f\xa1\xa2\x8e\x83\xd0\x22\xb2\x2c\xca\xe3\x3f\xf9\xfc\xc4\xc0\x6a\xe6\x40\xb1\xec\xdc\x15\x9c\x52\xd6\x6a\x56\xee\x3b\xc9\x11\x88\xa4\x04\x79\x09\xdd\xab\xd6\x01\xeb\xa7\x04\xe6\x1e\x9c\x4b\xd9\xad\x33\x1c\xe8\x62\xf4\x59\x34\xb6\x34\x3b\x3d\xc7\x0d\xc4\x68\x6a\x82\xae\x44\xb4\xa0\x80\xf0\x24\xcc\x93\x4b\xe7\xcf\xd3\xfb\x4e\xd2\xa9\x9d\xda\x67\x6f\x7b\xaf\x4f\xad\xe6\x2b\x58\xcc\x80\xf5\x6b\x27\xb9\x6d\x35\xdb\xb0\xd8\xf2\x3f\xe0\xc4\xfc\x87\x22\x3f\xd6\x13\xc8\xd7\xec\xed\x32\xc8\x62\x67\x63\x7c\x44\x32\xcc\xa0\x26\x84\x04\xd9\x10\x79\xc9\x0c\xf2\x32\xbe\x1b\x39\x07\x65\xa3\xfd\xd6\xa9\x9e\x80\x54\x05\xe6\xaa\x56\x33\xe7\xac\x76\xc9\x21\x38\x21\xb8\xbe\xc1\x09\x21\xf8\xb4\x9f\x7a\x72\x22\xa9\x35\x01\x58\xcd\x2d\x7b\xbb\x0c\xd3\x0d\xf4\xe3\x41\x13\x14\xb2\x7e\x4e\xe5\x75\xfd\x67\x45\x0b\xd1\x03\xe5\x88\x44\x72\xb9\x61\x1f\x25\x3d\x5c\xcb\xfe\xa3\x85\x0f\x52\x35\x71\x92\x37\x04\xf3\xa9\x30\x35\x8c\xf0\x6e\x38\x84\xe0\xb8\xc0\xbf\xe2\x89\x25\x6e\x42\xd1\x0d\x93\x97\x34\x81\x0f\x4d\xf5\x3b\x8d\xa4\x93\x69\x45\x37\xaf\xd8\x9a\x5e\x65\x34\x9d\x14\x67\xb7\xf3\x24\x9d\x5e\x3b\x8d\x3a\x30\x9f\x4c\x99\xaa\xa6\xfc\x5d\x08\x1a\x8f\xf5\x55\xfa\x62\x1f\x54\x71\xf8\xfa\x94\x6e\x08\x61\xd3\x9d\x90\x31\x25\x22\x87\xdc\x01\x59\x4f\x93\x71\x80\xc5\x4f\xb0\xd8\x12\x55\x52\x08\x05\x99\xf2\x93\xbc\x28\xf1\x4f\x24\x54\x37\x51\xf5\x39\x9f\xd7\x60\xab\x80\x32\x73\x73\xed\xe7\x44\x9d\x34\x2c\x0e\x11\xfb\x72\x23\xb0\x9a\x5b\x64\xb4\x7c\xa2\xea\x03\x1b\x57\xf6\x79\xf2\xbe\x93\x85\xc5\x73\x67\xb5\x0b\xd3\x05\xb0\x79\x08\x1b\x6d\xb0\x79\xe4\x27\xb9\x14\xc3\xaa\xa2\x19\xa6\xa0\x69\x8a\xe6\xd5\x00\x16\xcf\x61\xe6\x16\xa4\xeb\x20\x5f\x23\xd3\x60\xef\xaf\xc1\x9d\x3a\xcc\x55\xf1\x5d\x1b\xf0\xc3\x4b\x78\x78\x4a\x3e\xc1\xdd\x94\xd5\xbe\xc1\x61\xb3\x0e\x91\x2b\x73\x92\x97\x22\x28\xfa\x1f\x12\x3e\xa7\x94\x85\xc5\x8c\xfd\x47\x8b\xf8\x19\x34\xfe\x59\x53\xe4\x71\x73\x4c\xd1\xc2\xbc\x41\xcd\xed\x8b\x1a\xc8\x7f\x80\x47\x1d\xd0\xc9\x5b\xcd\x1c\xac\x7c\xb0\x4b\x43\x38\xa6\xd3\xd9\xbc\xba\xc7\x65\x6e\xd1\xd4\xa7\xeb\xa0\xb6\xd1\x7b\x7d\x3a\x88\x74\x2b\xf9\x2d\x18\x29\xdf\x20\xcc\x6d\x05\x39\x12\x46\x3d\xb0\x7e\xe9\x73\x31\xf8\x4e\x24\x48\xd0\x6c\xde\x77\xb2\x4e\xe3\xda\xb9\x4b\xb9\x60\x3a\xd7\x5e\x27\x89\xf8\x3c\x52\x2f\x52\x56\xe4\xc8\xba\x7b\x87\xfa\xd7\xed\xa9\x2c\xd8\x3a\x04\x07\x47\xf7\x9d\xfd\x1f\x12\x5f\x0d\x42\xe7\x27\x05\xea\xe5\x7b\x78\x8e\x0b\x2c\x27\xa2\x8b\xfd\x41\x45\x17\x0f\xf3\x46\x70\x82\x4e\x69\x6f\x7b\xcf\xa9\xd5\xfc\x9c\xa2\x89\xe3\xa2\xec\xa6\x94\x9a\x6c\x1d\x0e\x5a\x45\xf4\x3e\x73\xf2\x41\x43\xc4\xb1\x10\x5e\x80\xc5\x73\x50\x40\xac\x44\x9a\xc5\x49\x6e\xa3\x1d\x50\x2d\xd9\x85\x0d\xf0\xfb\x5b\xdc\x29\x18\xcd\xe6\x1e\x11\x6d\xe5\x84\xe0\xb1\x05\x1f\x0a\x8b\xf2\x43\xfc\xe5\x9b\x7e\x36\x3f\x1b\xf3\x11\x73\xc2\x15\xce\xf1\x05\xc3\x64\x6c\x74\x9a\x20\xf1\x86\x10\x62\x86\xd7\x0d\xf3\xaa\x44\x59\xd4\xab\xb5\x14\xe2\x55\x7a\x68\x44\x0d\xf1\xf8\x50\xf4\xeb\xc0\x61\xd6\x5d\x15\x6e\x7f\xc1\x27\x4d\x0a\x9a\x38\x36\x65\x06\x95\x10\xb3\xaf\x7a\xe7\x59\xa7\x96\x64\xb2\x25\x84\x79\x51\x32\x43\xa2\xee\x8e\x73\x6f\xa5\x6a\xb5\x6f\xc8\x52\x74\x8e\x2f\xec\x8f\xc9\x87\xd2\xe5\x61\xd9\x62\x10\x34\xc8\x7e\xe9\xad\xe7\x28\x05\xb9\xf4\x4d\x0b\xac\xa3\xff\xf5\x69\xdc\x63\x6c\xca\xe1\xa4\x9c\x1e\x81\x0f\x62\x31\x1b\xb0\x28\x98\xde\x1d\x2c\x3f\x09\x0a\x13\x30\x09\xc8\xae\x5e\x32\xfc\xcb\x71\x81\xc5\xe8\xf3\xd9\x78\xcc\x23\x52\x4d\x18\x17\x15\xf9\x51\xdb\x12\x64\x5b\xe0\xf0\x90\x25\x52\x66\xc7\x71\x81\xff\xc6\x63\x51\xcf\x2b\xda\x93\x8f\xf3\xe9\x79\x18\xe0\xe7\xd5\xb2\xdd\xbd\x72\xaa\x25\x90\x7e\x3d\xa8\x06\x08\x0d\x3a\x5b\x0d\x90\xdf\x75\xc9\x00\xef\x05\x96\xfd\x7a\xeb\x39\xbb\x4b\xd8\x5d\x0c\xf3\xe3\xc2\x43\xc0\x9d\x22\x58\xcd\x3f\x04\x54\x4d\x7d\x42\xd1\x50\x0a\xc5\x05\x9f\x47\xed\x1c\x17\x88\x2f\x44\x63\x89\xa5\xe9\x99\x17\xdc\xb8\x60\x78\xc9\xf3\xaa\xd2\x27\x36\x9c\x29\x94\x94\x85\xc5\xf8\xbf\xa3\x33\x4b\xe1\xe8\xfc\xaf\xd1\x45\xda\xfd\x7c\xc8\xa5\x35\xb7\x8e\xf8\xee\x1e\xfb\xb3\x23\xc2\x30\x20\x1d\x61\xb2\x01\xc8\xf2\xf5\xfc\x53\x89\xe1\xd1\xe7\x03\xf3\x49\xc6\x69\x44\x5f\x78\xa8\x09\x5e\x37\xdd\x74\x23\x08\x36\x66\xd7\xef\x7d\x27\x39\x82\xe5\x02\xb1\xf8\xb3\x28\x27\xa3\x39\xf3\xf6\xbf\x2b\xf2\x4c\x83\xd7\x9f\x62\x3e\xbe\xb1\xda\xbb\xce\xe6\x8a\xbd\xf2\x05\xee\xd4\x7b\xa9\x3c\x7c\x93\xb3\xba\x45\x44\xca\xc5\x32\xcc\x9c\x39\xa5\xec\x8f\x3e\xa7\xd6\xb0\x2b\x19\x70\xbb\x0e\xaa\xab\x56\xfb\x13\xf9\x19\x54\xb3\xb0\xb6\xf3\x57\x74\xcc\xaf\xd3\x33\x2f\x96\x17\xf8\x99\x99\xf8\x72\x6c\x89\x53\x79\x8d\x0f\x9b\x42\x58\x35\xa6\xd0\x09\xf9\x15\xb8\x53\xf7\xea\x84\x2e\xae\x47\x54\x95\x14\x10\xb1\xff\x76\x0d\x66\x91\x28\xb5\x2f\xdb\xe0\xfd\x2b\x3f\x37\xa4\xa5\xe0\x51\xa9\x77\x9e\x65\x26\xd6\x6d\xe3\x27\x7c\xf0\x69\x44\x35\xf9\x60\x50\x89\xc8\x8f\xd5\x1b\xe0\x24\x65\xb5\xbb\xce\x9f\x1f\x41\xbe\xf1\x80\xea\xe0\x02\x0b\x73\xd3\xb1\x21\x91\xf9\x1d\xb7\xec\x50\x8c\xba\x65\xc4\xfa\x13\x9c\x27\xce\x8d\x7d\x4c\x94\x04\x9d\xd5\x71\xae\xb6\x71\xfd\xa3\x50\x89\x50\x00\x1b\x39\x90\x4e\xc1\xdc\x31\x1b\xff\xc0\xf4\x11\x8f\x54\xb9\x93\xea\x11\xeb\xaf\x28\x77\x4d\xd0\x0d\x45\x13\x86\xcc\x61\xf2\x18\x9c\xe4\x3c\x73\xda\xbf\x5a\x34\x11\x5f\x5e\x9c\x89\x8e\xa6\x84\xb9\xc6\x37\x82\x67\x8b\x36\x24\x0d\xfb\xfd\x3c\x20\x04\x2f\xdf\x5b\xad\xad\xa1\xae\x76\xee\x0e\x90\x66\xa8\x9c\x90\xf6\x64\xa4\xd3\xc8\x34\x7a\x92\x2e\xe7\x11\x2b\x95\x12\x64\xbc\x05\x4f\xb7\x60\x11\x01\x1a\x57\xde\x74\xab\x12\x2f\x7f\xc5\x1f\x5b\x5d\xc6\xeb\x50\x03\x8e\xe2\xd8\x62\x31\x38\x2e\x30\xf7\x6c\x7a\x81\x6c\x4d\xca\x27\xae\xf2\x2f\x64\x61\xe5\x14\xa4\xeb\x68\x1c\xdc\xd5\xd9\xfa\x9b\xbb\x61\x25\x65\x5c\x94\x87\x11\x74\xc1\x12\xef\x38\x2d\xd8\xba\xbf\x36\xc9\x5b\xd2\xfe\x98\x84\x95\x63\xf4\x89\x40\xec\xbd\x76\x6f\x6f\xc3\xf7\x80\xd4\xe0\x02\x33\xb3\x09\x2e\x28\xea\x98\x26\x86\xb9\x03\x45\xea\xad\x2d\x98\xb9\x80\xf9\xbc\xd5\xac\xd8\xfb\x6b\x56\xbb\x0d\x36\x4b\x08\xfd\x9f\xf9\x60\x3c\xf6\xcf\xd9\xe7\xdc\x64\xd8\x0c\x2a\xf2\x98\x38\xde\x7f\x77\x10\x2a\x67\xf6\x5e\xdf\x66\xf8\x09\x4b\x4d\xfb\x5d\xc1\xd4\xe2\x9b\xbd\xc1\x05\x66\x17\xd0\xaa\xe8\xbf\xbc\xf1\x8b\xc1\xde\x5f\x13\x55\x58\xfd\x8c\x76\x79\xe3\x0a\x16\x33\xf8\x7f\x54\xf1\x3b\xb5\x46\x2f\x95\x83\xbb\x75\x62\x4d\x9f\xe7\x43\xaf\x0a\xd4\x3c\xc7\x17\x20\x77\x84\x05\x48\x96\x8a\xa9\x5e\x6a\x13\x6e\xdf\x62\x94\x28\x4f\xf2\x92\x18\xc2\xf6\xa0\x58\x07\xef\x5e\x52\x39\xee\x6d\x31\x5a\xc2\x43\x98\x2e\xe0\xc0\x88\x21\xa1\x61\x54\x79\x1f\xdd\x7f\xce\xe7\x06\x38\x7b\x45\xbe\x83\xd7\x59\x50\x29\xd8\xdd\xdf\xe9\x45\xff\xb2\x10\x8f\xcf\x21\x97\xaa\xa2\x48\x23\x59\x14\x17\x7c\xf0\xf2\xe8\xab\xda\x00\x3d\x72\x98\x3f\x13\xf8\x5d\x09\x92\x58\x5a\xfc\x8d\x43\x0b\x54\x37\xb4\xa9\xd1\x97\xa3\x53\x3b\x85\xef\x6f\xe0\x7b\x97\xb2\x71\xe5\xd1\xe3\xc9\x3d\x82\xf6\xb9\x4b\x7a\x94\x40\x3c\x06\x52\x35\x65\x52\x0c\x09\x1a\x5d\xce\x4e\xea\x1c\x6c\x96\x61\xb9\x04\x3a\x79\x90\x3e\x00\xed\x16\xfd\x23\x05\xc9\x2d\xcc\xbc\x02\x9b\x65\xa2\x20\xd9\x17\x00\x33\xdb\xc6\x94\x2a\x50\x87\x74\xaf\x90\xd1\xa6\xdb\x05\xb5\x4c\x7b\x9d\x7c\xf2\x21\xfa\xf8\x71\x70\x8a\xbf\xbf\x5d\xfa\x94\xf5\x9d\xed\x82\x13\x6d\xef\xaf\x91\x1f\xed\xca\xae\xbd\x73\x8a\x38\x19\x8b\xfc\x3e\xc9\x26\xa2\x89\x04\x92\x7c\xba\xa0\xeb\x48\xaf\xa0\xd7\x3d\xcb\x23\xee\x77\xdf\x8b\xe8\x6f\x3e\x54\xaa\x9d\xb4\xab\x42\x54\x4d\x41\xc4\xcb\x34\x9a\x67\x4a\x06\xdc\xe9\x7e\x02\xd9\x5d\x12\x9e\x0b\x41\x12\x42\x56\x58\xd1\xc7\x8a\x7a\x2f\xa3\x28\xa8\xa5\xf8\xe2\xf4\xf3\x28\x87\x0b\x85\xc2\x42\xb5\xfa\x45\x17\x0d\x76\xf2\xd1\xbf\x2b\x65\x98\x2f\x58\xb7\x07\x60\x67\xc3\xcf\xfd\x3f\x00\x00\xff\xff\xe4\xcd\xe5\x1d\xf2\x12\x00\x00")

func localesZhCnHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesZhCnHomeYml,
		"locales/zh-CN/home.yml",
	)
}

func localesZhCnHomeYml() (*asset, error) {
	bytes, err := localesZhCnHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/zh-CN/home.yml", size: 4850, mode: os.FileMode(420), modTime: time.Unix(1620699134, 0)}
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
	"locales/en-US/home.yml": localesEnUsHomeYml,
	"locales/zh-CN/home.yml": localesZhCnHomeYml,
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
	"locales": &bintree{nil, map[string]*bintree{
		"en-US": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesEnUsHomeYml, map[string]*bintree{}},
		}},
		"zh-CN": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesZhCnHomeYml, map[string]*bintree{}},
		}},
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
