// Code generated by go-bindata. DO NOT EDIT.
// sources:
// pkg/security/tests/schemas/chmod.schema.json
// pkg/security/tests/schemas/chown.schema.json
// pkg/security/tests/schemas/container.json
// pkg/security/tests/schemas/container_context.json
// pkg/security/tests/schemas/container_event.json
// pkg/security/tests/schemas/datetime.json
// pkg/security/tests/schemas/event.json
// pkg/security/tests/schemas/exec.schema.json
// pkg/security/tests/schemas/file.json
// pkg/security/tests/schemas/host_event.json
// pkg/security/tests/schemas/link.schema.json
// pkg/security/tests/schemas/open.schema.json
// pkg/security/tests/schemas/process.json
// pkg/security/tests/schemas/process_context.json
// pkg/security/tests/schemas/rename.schema.json
// pkg/security/tests/schemas/selinux.schema.json
// pkg/security/tests/schemas/usr.json
// +build functionaltests

package schemas

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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataChmodSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x5f\x6e\x83\x30\x0c\xc6\xdf\x39\x45\x14\xed\xb1\x6d\x3a\x1e\x73\x89\x1d\x60\x42\x53\x46\xcc\x48\x05\x31\x73\xbc\x49\x55\xc5\xdd\xa7\x40\x41\x6d\x09\x5d\xeb\x27\xe3\xdf\xe7\x3f\x1f\xca\x29\x13\x42\x08\xf9\x12\xca\x1a\x5a\x23\xb5\x90\x35\x73\x17\xb4\x52\x87\x80\x7e\x3b\x96\x77\x48\x5f\xca\x92\xa9\x58\xe5\xfb\x7c\xbf\x7d\xcd\xd5\x59\xbf\x39\xb7\x3b\x1b\x5b\xcb\xba\x45\xbb\x8b\x8d\x13\xe0\x63\x07\x91\xe0\xe7\x01\x4a\x9e\xaa\xc6\x1f\xdf\x2a\xa9\xc5\xfb\xf0\x19\xe3\x34\x67\xe3\x3c\x82\xc8\x65\xe5\x1a\xd0\x4a\xa9\x12\x3d\x1b\xe7\x81\x3e\xe0\x17\x3c\x8f\x2b\xe6\x96\x7e\xf3\xe8\x9c\x1a\x03\xa7\x47\x0c\x59\x31\xdd\xd7\x34\xf7\xef\xeb\x08\x3b\x20\x76\x10\xa4\xbe\x61\x03\x8f\xfb\x92\x64\xfd\xa7\x2c\x54\x04\xdf\x3f\x8e\xc0\x5e\xdd\xb1\x50\x59\x08\xec\xbc\x61\x77\xe9\xe6\x32\x8a\x95\xf1\xff\x58\x48\x2e\xb8\x27\x7c\xdc\xd9\x73\x0e\x67\x75\x8b\x16\xd2\x16\xa7\x28\x56\x69\x9f\x24\xcb\xea\x75\xa5\xbf\x7d\x1e\x59\x9f\xfd\x05\x00\x00\xff\xff\x28\x58\x9f\x39\x30\x03\x00\x00")

func bindataChmodSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataChmodSchemaJson,
		"/chmod.schema.json",
	)
}

func bindataChmodSchemaJson() (*asset, error) {
	bytes, err := bindataChmodSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/chmod.schema.json",
		size:        816,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataChownSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x5f\x4e\xc3\x30\x0c\xc6\xdf\x73\x8a\x28\xe2\x71\x9b\x47\x1f\x73\x09\x0e\x80\x26\x14\x12\x97\x64\x2a\x49\x70\x0c\x68\x9a\x7a\x77\x94\xfe\x99\x0a\x2d\x88\x3c\xd9\xfe\xf9\xb3\xbf\xf8\x2a\xa4\x94\x52\xdd\x15\xeb\xf1\xd5\x28\x2d\x95\x67\xce\x45\x03\x9c\x4b\x8a\xfb\xb1\x7c\x48\xf4\x02\x8e\x4c\xcb\xd0\x1c\x9b\xe3\xfe\xbe\x81\xa9\x7f\x37\xc9\x83\xab\x52\xeb\xd3\x67\x3c\x54\xe1\x0c\xf8\x92\xb1\x92\xf4\x7c\x46\xcb\x73\xd5\xc4\xcb\x43\xab\xb4\x7c\x1c\xd2\xfa\xae\xb7\x68\x9c\x47\x58\xb9\x6a\x43\x87\x1a\x00\x6c\x8a\x6c\x42\x44\x7a\xc2\x0f\x8c\x3c\xae\xb8\x49\xfa\xdd\x7f\xe7\xf8\x54\x78\x7b\xc4\x10\x9d\x66\x7f\x5d\xf7\xb7\xbf\x4c\x29\x23\x71\xc0\xa2\xf4\x0f\x36\xf0\xba\x6f\x93\xfc\x7e\x94\x55\x17\xe1\xdb\x7b\x20\x74\xdf\x7c\xac\xba\x1c\x16\x0e\xd1\x70\x58\xfe\x66\xf9\x4e\xab\x6a\x2f\xb6\xb3\xe9\x06\xa2\x17\x5f\x01\x00\x00\xff\xff\x02\x48\x3f\x29\x15\x02\x00\x00")

func bindataChownSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataChownSchemaJson,
		"/chown.schema.json",
	)
}

func bindataChownSchemaJson() (*asset, error) {
	bytes, err := bindataChownSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/chown.schema.json",
		size:        533,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataContainerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xc1\xae\xc2\x20\x10\x45\xf7\xfd\x8a\x09\x79\xcb\xd7\x52\x59\xf2\x2b\xc6\x05\xd2\xb1\xa5\x89\x80\xc3\xb8\x30\x86\x7f\x37\x50\xaa\xe9\x8e\x7b\x2f\xe7\x64\xde\x1d\x00\x80\xf8\x4b\x76\xc1\xbb\x11\x1a\xc4\xc2\x1c\x93\x96\x72\x4d\xc1\xf7\x5b\x3d\x04\x9a\xe5\x44\xe6\xc6\x52\x8d\x6a\xec\x4f\x4a\xb6\xff\xff\x0d\x77\x53\x41\x6d\xf0\x6c\x9c\x47\x1a\x0a\xbc\x8f\xfc\x8a\x58\xd6\x70\x5d\xd1\xf2\xde\x46\x0a\x11\x89\x1d\x26\xa1\x61\xbb\xa2\xf6\xd5\xf4\xcb\x07\x43\x62\x72\x7e\x16\xdf\x31\xd7\x57\x6e\x46\xc2\xc7\xd3\x11\x16\xfe\x7c\xf0\xd5\x70\xe9\xf2\x27\x00\x00\xff\xff\x12\xe6\xde\xdb\xec\x00\x00\x00")

func bindataContainerJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataContainerJson,
		"/container.json",
	)
}

func bindataContainerJson() (*asset, error) {
	bytes, err := bindataContainerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/container.json",
		size:        236,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataContainercontextJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xd1\x0a\x82\x40\x10\x45\xdf\xfd\x8a\x61\xe9\x31\x1d\xf3\x71\x7f\x25\x22\x4c\xc7\x5c\x29\x77\x9b\x9d\xa0\x08\xff\x3d\xc6\x55\x21\xe8\x6d\x38\x9c\xb9\xf7\x7e\x32\x00\x00\xb3\x8b\x4d\x4f\xf7\xda\x58\x30\xbd\x48\x88\x16\x71\x88\x7e\xcc\x13\x2e\x3c\x5f\xb1\xe5\xba\x13\xac\xca\xaa\xcc\x0f\x15\x2e\xfe\x7e\x79\x77\xad\xbe\x36\x7e\x94\xda\x8d\xc4\x67\xbd\xe8\x25\x85\x86\xac\x92\xbc\x03\xa9\xe5\x2f\x03\x35\xb2\xd2\xc0\x3e\x10\x8b\xa3\x68\x2c\xa4\x35\x33\xdf\xb2\x7e\x70\x6a\x63\xea\x34\xa8\x73\x37\xb2\x88\xb8\xa9\xa9\x6e\x93\xa7\xf9\x9a\x96\x22\xa6\xc7\xd3\x31\xe9\xd0\xe3\xbf\x9a\x99\x9d\xb2\xe9\x1b\x00\x00\xff\xff\xb9\xfa\x42\x76\x12\x01\x00\x00")

func bindataContainercontextJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataContainercontextJson,
		"/container_context.json",
	)
}

func bindataContainercontextJson() (*asset, error) {
	bytes, err := bindataContainercontextJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/container_context.json",
		size:        274,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataContainereventJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\x50\x50\x52\x29\x4e\xce\x48\xcd\x4d\x54\xb2\x52\x50\xca\x28\x29\x29\x28\xb6\xd2\xd7\xcf\x2a\xce\xcf\xd3\x85\x08\xeb\xe5\x17\xa5\xeb\xa7\x14\x25\xa6\x95\xe8\x1b\x19\x18\x19\xe8\x1a\x1a\xe9\x43\xd5\xeb\x40\xb5\x67\xa6\x80\xb4\x26\xe7\xe7\x95\x24\x66\xe6\xa5\x16\xc5\xa7\x96\xa5\xe6\x95\xe8\x81\x8c\x80\x29\x49\xcc\xc9\xf1\x4f\x53\xb2\x52\x88\x06\x73\x41\xa0\x1a\xce\x82\x98\x51\x94\x0a\x92\x57\x4a\xcb\xcc\x49\xb5\xd2\xd7\xd7\xcf\xc8\x2f\x2e\x41\x36\x07\xae\xba\x56\x87\x58\x23\x10\xee\x01\xb1\x52\x2b\x30\x4c\x02\xb3\x62\xb9\x6a\x01\x01\x00\x00\xff\xff\xef\xab\x6c\xf9\x06\x01\x00\x00")

func bindataContainereventJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataContainereventJson,
		"/container_event.json",
	)
}

func bindataContainereventJson() (*asset, error) {
	bytes, err := bindataContainereventJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/container_event.json",
		size:        262,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataDatetimeJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\x4d\xca\x83\x30\x10\x86\xf7\x9e\x22\x0c\xdf\xd2\x98\x31\x7c\x9b\xe6\x12\x3d\x80\x58\x08\x35\xfe\x14\x8d\x92\xcc\xa6\x48\xee\x5e\xa2\xd6\x62\x4b\x69\xc8\xe2\xcd\xcc\xf3\x90\x77\x4e\x18\x63\x0c\xfe\xfc\xb5\x35\x83\x06\xc5\xa0\x25\x9a\xbc\x12\xe2\xe6\x47\xcb\xd7\x71\x36\xba\x46\x54\x4e\xd7\x24\x24\x4a\xe4\xb9\x14\x1b\x9f\x6e\x7a\x57\x45\xb5\xd2\x64\xa8\x1b\x4c\x16\xdd\xe7\x4e\xf7\xfd\xb9\x06\xc5\x8a\xe5\x19\xcf\xbc\xa7\x05\xa0\xfb\x64\xa2\xed\xc9\x75\xb6\xd9\xb4\x7d\x3b\x69\x22\xe3\x6c\x04\x2e\x05\xf2\x53\x39\xff\x07\xbe\x06\xf9\x0a\xb0\x4b\x21\xfd\xf6\x8f\x1d\x09\xd4\xdb\xf0\x77\x81\x8f\x12\x88\x98\xf3\xe5\xc2\x01\x0c\xc9\x31\x95\x49\x78\x04\x00\x00\xff\xff\x99\x6a\xef\xa9\x5d\x01\x00\x00")

func bindataDatetimeJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataDatetimeJson,
		"/datetime.json",
	)
}

func bindataDatetimeJson() (*asset, error) {
	bytes, err := bindataDatetimeJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/datetime.json",
		size:        349,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataEventJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x52\xcb\x6e\xc2\x30\x10\xbc\xfb\x2b\xac\x15\x47\x60\x69\x8e\xf9\x95\xaa\x07\x37\xd9\x04\xa3\xc6\x76\xd7\x0b\x12\xaa\xf8\xf7\xca\x38\x20\xb9\x79\x5c\xca\x75\x66\x67\x76\xc6\xde\x1f\xa5\xb5\xd6\xb0\x89\xcd\x91\x06\x03\xb5\x86\xa3\x48\x88\x35\xe2\x29\x7a\xb7\xcb\xf0\xde\x73\x8f\x2d\x9b\x4e\xb0\x3a\x54\x87\xdd\x5b\x85\xe3\xfc\x76\x94\xdb\x36\x49\xe9\x42\x4e\xf6\x49\xf8\x20\xe4\x1a\x28\x31\xfe\xf3\x44\x8d\x3c\xd0\xc0\x3e\x10\x8b\xa5\x08\xb5\xce\x09\xee\x38\x5d\xa4\x00\x96\x2d\x9e\xec\x82\xd5\x93\x77\x66\xa0\x59\xa6\xf0\x8e\xc2\xd6\xf5\x30\x19\xba\x6d\xa7\x8e\x8d\x11\xea\x3d\x5f\x5f\xeb\xea\xcf\xd2\xf8\x7f\x45\x55\x2b\x2b\x80\xe9\xfb\x6c\x99\xd2\x2f\xbd\x2f\x3c\xd2\x5a\xd5\x95\xc0\x05\xf3\xa1\x66\x02\x40\x6b\x64\x5a\x0c\x36\x4c\x5d\x2a\xd4\xd9\x2f\xaa\x11\x31\x4d\x89\x1d\x28\xdf\x8f\x2a\x7b\x8d\x76\xf3\x3d\xee\x77\xf3\x77\x9f\xca\x79\x6e\xbf\x01\x00\x00\xff\xff\x77\xb8\x98\xe7\xe1\x02\x00\x00")

func bindataEventJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataEventJson,
		"/event.json",
	)
}

func bindataEventJson() (*asset, error) {
	bytes, err := bindataEventJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/event.json",
		size:        737,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataExecSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x51\x6e\xc2\x30\x10\x44\xff\x73\x0a\xcb\xea\x27\xb0\x34\x9f\xbe\x44\x0f\x50\x21\x94\x3a\x93\xc6\x28\xd8\xe9\x7a\x5b\x81\x50\xee\x5e\x39\x21\x48\x80\x53\x35\x5f\x9b\x79\x33\xe3\xd5\x5e\x0a\xa5\x94\xd2\x2f\xd1\xb6\x38\x56\xda\x28\xdd\x8a\xf4\xd1\x10\x1d\x62\xf0\xeb\x49\xde\x04\xfe\xa4\x9a\xab\x46\xa8\xdc\x96\xdb\xf5\x6b\x49\x57\xff\xea\x1a\x77\x75\x8a\xe2\x04\xbb\x49\xb9\x59\x97\x73\x8f\x04\xc2\xc7\x01\x56\x66\xb5\xf2\xe7\xb7\x46\x1b\xf5\x3e\xfe\xa6\xef\x72\x9b\xa6\x3a\x46\xe2\xba\x71\x1d\x0c\x11\xd9\xe0\xa5\x72\x1e\xbc\xc7\x0f\xbc\x4c\x4f\xdc\x22\xc3\xea\xbf\x3d\x6d\x88\x92\xaf\x18\xa7\xdd\xbc\x5f\xd7\xfd\xbd\x5f\xcf\xa1\x07\x8b\x43\xd4\xe6\x81\xcd\xdc\x22\xe6\xe1\xf2\x5d\x9e\x5c\x3e\xc8\x62\xc5\x68\x60\x7c\x7d\x3b\x46\x7d\xb7\x6b\xd6\x89\x93\x93\xbd\xb8\x23\xf4\xa2\x6f\x97\x25\xc3\x93\x7a\xaf\x0c\x8f\x37\x2c\x86\xe2\x37\x00\x00\xff\xff\x84\xb4\xfe\x2c\x54\x02\x00\x00")

func bindataExecSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataExecSchemaJson,
		"/exec.schema.json",
	)
}

func bindataExecSchemaJson() (*asset, error) {
	bytes, err := bindataExecSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/exec.schema.json",
		size:        596,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataFileJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x41\x6e\xf2\x30\x10\x85\xf7\x9c\x22\xb2\x58\x02\xe6\x67\x99\x1b\xfc\x17\xe8\xa6\xaa\xd0\x34\x9e\x38\x83\xb0\x9d\x8e\x27\x0b\x84\xb8\x7b\x65\x0c\x54\x86\x80\xaa\xa8\x59\x3e\xfb\x7b\x19\xcf\xbc\x39\xce\xaa\xaa\xaa\xd4\x3c\x36\x1d\x3a\x50\x75\xa5\x3a\x91\x3e\xd6\x5a\xef\x62\xf0\xcb\x2c\xaf\x02\x5b\x6d\x18\x5a\xd1\x9b\xf5\x66\xbd\xfc\xb7\xd1\x97\xfb\x8b\x0b\x4e\x26\xa1\x2d\xed\x71\x95\xb8\xab\x2e\x87\x1e\xd3\x41\xf8\xdc\x61\x23\x57\xb5\xe7\xd0\x23\x0b\x61\x54\x75\x95\x0b\xc8\x3a\x48\x57\x28\x85\x47\x14\x26\x6f\xd5\xed\xf0\xb4\xf8\x21\x3d\x38\x9c\x46\x92\x0f\xe6\x05\x4a\x5e\xd0\x22\xab\x45\x79\xdc\x06\x76\x20\xe9\xc2\x1b\xec\xc9\xfc\x3f\x9b\x8c\xfa\xbb\x5f\xd9\x3f\x41\x07\x2f\xdb\x73\x67\xa7\xe0\x69\x18\xf1\x10\x05\xdd\xb4\xce\x0c\x11\x79\x1a\x69\x39\x0c\xfd\x34\xd4\x05\x43\x2d\x35\x20\x14\xfc\x56\x68\x6c\xaa\x73\xc6\xf6\x1a\xb6\x5a\x6b\x6d\x40\x30\xdd\xcc\xc1\x1b\x75\x6d\x3a\xf0\x16\xff\xce\xaf\xdd\x83\x8d\xcf\x1f\x08\xcc\x70\xb8\x8f\x0c\x09\xba\xc4\xbc\x17\x72\xfa\x8e\x0f\xca\xcb\x76\xdd\x0a\x2a\x94\x8f\x59\xa9\x5f\xca\x55\x8c\x5f\x03\x31\x9a\xe2\xcf\x79\xd3\xee\xf7\xe7\x61\x2b\xee\x63\x3c\x96\xcd\xf1\xc0\xbd\x9c\xe8\x93\xc1\xcc\xf2\x33\x4e\xdf\x01\x00\x00\xff\xff\x64\x93\x2e\xfd\x92\x04\x00\x00")

func bindataFileJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataFileJson,
		"/file.json",
	)
}

func bindataFileJson() (*asset, error) {
	bytes, err := bindataFileJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/file.json",
		size:        1170,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataHosteventJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x51\x0e\x82\x30\x0c\x86\xdf\x39\x45\xb3\xf8\x08\x14\x79\xdc\x25\x3c\x80\x31\x64\x81\x22\x18\x60\x64\xab\xc6\x84\x70\x77\xb3\x81\x44\x50\x13\x78\xa1\x6b\xff\xef\xdb\x3a\x04\x00\x00\xe2\x60\xf3\x8a\x5a\x25\x24\x88\x8a\xb9\xb7\x12\xf1\x66\x75\x17\x4d\xed\x58\x9b\x2b\x16\x46\x95\x8c\x69\x92\x26\xd1\x31\xc5\x39\x1f\xce\x78\x5d\x78\x54\x5b\xce\xe8\x41\x1d\xc7\x8e\x7e\x4f\x55\xd3\x9c\x4a\x21\xe1\xec\x8f\xee\x1b\x96\x6a\xc2\x0d\xb9\xb9\x28\xeb\x86\x24\x22\x7e\x28\x96\xe0\x18\xfe\xa3\x1d\x25\x24\x0c\xdf\x1e\xf7\x9f\x34\x30\xee\x10\x6d\xf1\xbb\x35\xbb\x1f\xb1\x65\x7b\xa3\x73\xb2\x36\xcb\x75\xc7\xf4\xdc\xbf\x4c\xa1\xd8\x2f\xb3\xea\xfe\xbc\xc1\x25\xb9\x6e\x69\xa3\xf6\xfa\x60\x5d\x5d\x82\xf1\x15\x00\x00\xff\xff\xab\xb6\xb3\x64\xe8\x01\x00\x00")

func bindataHosteventJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataHosteventJson,
		"/host_event.json",
	)
}

func bindataHosteventJson() (*asset, error) {
	bytes, err := bindataHosteventJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/host_event.json",
		size:        488,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataLinkSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x51\x5d\x6e\xc2\x30\x0c\x7e\xef\x29\xa2\x68\x8f\x80\x59\x1f\x73\x89\x1d\x60\x42\x53\xd6\xba\x6b\x50\x49\x32\xc7\x9b\x84\x50\xee\x3e\xa5\xa5\x08\xd2\xc0\xf0\x4b\xd3\xef\xcf\x8e\x73\xaa\x84\x10\x42\xbe\x84\xa6\xc7\x83\x96\x4a\xc8\x9e\xd9\x07\x05\xb0\x0f\xce\xae\x27\x78\xe3\xe8\x0b\x5a\xd2\x1d\x43\xbd\xad\xb7\xeb\xd7\x1a\xce\xfa\xd5\xd9\x6e\xda\x64\x25\xb4\xfa\x80\x9b\xe4\x9c\x19\x3e\x7a\x4c\x94\xfb\xdc\x63\xc3\x33\xaa\xed\xf1\xad\x93\x4a\xbc\x8f\xbf\xa9\x4e\x97\xd3\x14\x48\x98\x78\xd9\x99\x01\x15\x00\x34\xce\xb2\x36\x16\xe9\x03\x7f\xd1\xf2\xd4\xe2\x62\x89\xab\x67\x73\x7a\x17\xb8\x1c\x31\x9e\x76\xf3\x7c\xc3\xf0\x78\x3e\x4f\xce\x23\xb1\xc1\x20\x55\xc6\x8d\x7c\xea\x57\x64\xee\x2f\x65\xa1\x22\xfc\xfe\x31\x84\xed\xcd\x1c\x0b\x55\x8b\x81\x8d\xd5\x6c\xae\x6f\x73\x5d\xbb\x3b\xf1\xff\x5c\xa1\xd8\xe0\x91\x50\x94\xf6\x9d\xbe\xd9\xa6\xf3\x8a\x45\x66\x89\xde\x22\x31\x7f\xb9\x2a\x56\x7f\x01\x00\x00\xff\xff\x66\xe2\x62\xf4\xcc\x02\x00\x00")

func bindataLinkSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataLinkSchemaJson,
		"/link.schema.json",
	)
}

func bindataLinkSchemaJson() (*asset, error) {
	bytes, err := bindataLinkSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/link.schema.json",
		size:        716,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataOpenSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xdd\x4e\xc3\x30\x0c\x85\xef\xfb\x14\x51\xc4\xe5\x36\x8f\x5e\xe6\x25\x78\x00\x34\xa1\x90\xba\x6b\xa6\x90\x04\xc7\x20\xaa\xaa\xef\x8e\xd2\x3f\x09\x1a\xd0\x7a\x65\x9f\xcf\xc7\x3e\xcd\x50\x09\x21\x84\x7c\x48\xa6\xc3\x37\x2d\x95\x90\x1d\x73\x4c\x0a\xe0\x96\x82\x3f\xce\xf2\x29\xd0\x15\x1a\xd2\x2d\x43\x7d\xae\xcf\xc7\xc7\x1a\x96\xf9\xc3\x62\xb7\x4d\xb6\xe2\x17\x9a\x53\xf6\xad\x3a\xf7\x11\x33\x08\xaf\x37\x34\xbc\xaa\xda\xf7\x4f\xad\x54\xe2\x79\x6a\xf3\x37\x6c\xd5\xbc\x8e\x30\x73\xd9\x5a\x87\x0a\x00\x4c\xf0\xac\xad\x47\x7a\xc1\x4f\xf4\x3c\x9f\xd8\x2c\xe3\xe1\xde\x3d\x5d\x48\x5c\x5e\x31\x55\x97\x35\x9f\x73\xff\xe7\x8b\x14\x22\x12\x5b\x4c\x52\xfd\x62\x13\xcf\xf7\x8a\xe4\xef\x47\xd9\x4d\x11\xbe\x7f\x58\xc2\xe6\x47\x8e\xfd\x25\xa7\xaf\x49\x16\xf9\x65\xa7\x8e\x55\xb9\x5b\xfe\xbe\x1a\xbf\x03\x00\x00\xff\xff\xc6\x03\x8c\xaa\x0d\x02\x00\x00")

func bindataOpenSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataOpenSchemaJson,
		"/open.schema.json",
	)
}

func bindataOpenSchemaJson() (*asset, error) {
	bytes, err := bindataOpenSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/open.schema.json",
		size:        525,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataProcessJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xc1\x6e\xe3\x20\x10\xbd\xfb\x2b\x2c\xab\xc7\xb6\x6e\x73\xcc\x1f\xec\x69\x3f\x60\x55\x45\x14\x0f\xc9\x74\x0d\xb8\x30\x5e\x6d\x14\xe5\xdf\x57\x76\x12\xa7\x31\xe0\xd8\x31\xd1\x36\xa7\x08\x98\x37\xc3\xf3\xbc\x19\x60\x97\xa4\x69\x9a\x66\x0f\x96\x6f\x40\xb2\x6c\x99\x66\x1b\xa2\xca\x2e\xf3\xfc\xc3\x6a\xf5\x74\x18\x7e\xd6\x66\x9d\x17\x86\x09\xca\x17\x2f\x8b\x97\xa7\xd7\x45\x7e\x5c\xff\x78\x34\xc7\xa2\x31\xad\x8c\xe6\x60\xed\x73\x63\x7a\x9a\xa2\x6d\x05\xcd\x9c\x7e\xff\x00\x4e\xa7\xd1\xca\xe8\x0a\x0c\x21\xd8\x6c\x99\x1e\x62\x38\xac\x6e\x81\xce\x03\x17\x10\xa8\x08\xd6\x60\xb2\x6e\x76\xff\x78\xb6\xac\x6f\xb6\x5c\xdf\xee\xd3\x82\x09\x9b\x5a\x32\xa8\xd6\x01\x9f\x46\xd7\xd5\x6d\xa6\x5c\x4b\x79\x9b\x25\xd1\xd6\x35\x64\x65\xf9\x53\x64\xcb\xf4\xd7\xc5\x70\xf3\xdb\x39\x23\x83\x9e\x3c\x1e\xaf\x20\x29\x4d\x4e\x3c\x43\xae\x5c\xe0\x6e\x25\xa8\x5a\x7a\x37\x71\xe9\xb0\x2e\x4b\x37\xe0\xd3\xef\xcd\x3b\xb3\x77\x37\x98\xf8\xad\xbe\x52\x2d\xb4\xf9\xbd\x22\x94\xe0\x12\xfe\x60\xa0\xe1\x3b\x13\x58\xc2\x32\xcf\xf3\x82\x11\x34\x2b\x0f\xaa\xf1\xa2\xc1\x5f\xe0\xf1\xd0\xb8\x81\x02\x14\x21\x2b\x6d\x38\x8f\x2e\xe4\xda\xcd\x06\x64\xdb\xcd\xfb\x44\xe8\x40\x3b\x92\xf2\x04\x79\x46\xf4\x49\xcc\x81\x9c\x90\x8b\x5e\xb9\xcf\x8b\xd1\x2f\xe6\x59\x41\x42\x7c\x26\x21\x3e\x95\x10\x9f\x4b\xb8\x03\x99\xc2\xc6\x67\x53\xd8\xf8\x74\x0a\x1b\x9f\x4f\x61\xef\x40\x28\x67\xd5\x0a\x84\x00\x4e\xf8\xc7\xad\x49\x0e\x34\x33\x86\x6d\x03\xd5\x3b\x43\x02\xe9\xaf\x26\xa3\x43\x0c\x84\xd9\xda\xd6\x0a\x3f\x6b\xf8\x71\x74\x42\xa6\x86\xd1\x3b\xac\xc0\x48\x24\x82\x11\x9f\xe4\xbb\xef\x30\x19\x00\xca\x0c\x7c\xd6\x68\xda\x6d\xba\xfd\xb3\x2d\xe9\x81\x2a\x1a\xaa\x5b\xa1\x4a\x11\x94\x66\x50\x0b\x57\x73\xef\xea\xa7\x1b\xd1\xab\x9b\xee\x5a\x13\x7b\x2f\x3d\xed\x75\x5e\x3b\xac\x18\x6d\xe2\x2a\x4f\x31\xcf\x21\x60\x16\x22\x2a\x5d\x8c\x80\x9c\x54\x74\xe4\x3d\x20\x6b\x45\xab\xf8\xe5\x11\x4b\xb0\x5b\x4b\xe0\x1e\xa9\x1d\xe0\x29\xac\xc6\x6f\x38\xf1\x5b\x83\xd4\x05\x0a\xe4\x8c\x50\x2b\xff\xe9\xb2\x5b\x3a\xf2\x94\x39\xe4\x8d\x6f\x98\x5a\x43\x64\x3f\x33\xea\x5a\xab\xcd\x90\xc2\x82\x3a\x09\x25\xfb\x50\xc6\x0e\xa7\x5d\x20\x79\xc6\x97\x5c\xf7\x33\x5e\x61\x7f\x44\x49\xe4\x5a\x11\x43\xe5\xbb\xdd\xf6\x3f\x51\xb7\xb4\x7f\xe3\x48\xbe\x80\x66\x5a\x41\xef\x96\xb9\x9b\x58\x4a\xa7\xa4\x7f\xa0\x53\x5e\xbd\x24\xbe\x7a\x67\xdc\x7b\xe1\xac\xbc\xc3\x31\x4d\xe9\x3f\xb0\x23\x51\xa1\x6c\x09\x5a\x8c\x12\x74\x35\xc9\x69\x74\x0e\x43\x11\x85\xb8\x6d\xff\xbd\x1d\xd3\xd1\x0b\xdf\x3e\x3e\xf5\x5e\x94\x7a\xcf\x44\xfd\x67\x98\xc0\xad\xda\x7f\xb8\x38\x44\x90\xec\x93\x7f\x01\x00\x00\xff\xff\xfb\xfe\xd0\xb1\x74\x13\x00\x00")

func bindataProcessJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataProcessJson,
		"/process.json",
	)
}

func bindataProcessJson() (*asset, error) {
	bytes, err := bindataProcessJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/process.json",
		size:        4980,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataProcesscontextJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\x41\x6e\xc3\x20\x10\x45\xf7\x3e\x05\x42\x5d\x26\x21\xf5\x92\x4b\xf4\x00\x55\x55\x51\x3c\x6e\x88\x1c\xa0\xc3\x54\x6a\x14\xe5\xee\x15\x06\x5b\x8d\x8c\x9d\x34\xf1\xf2\xf3\xff\xd7\x1b\x18\x9f\x2a\xc6\x18\xe3\x4f\x41\xef\xe0\xa0\xb8\x64\x7c\x47\xe4\x83\x14\x62\x1f\x9c\x5d\x27\x79\xe3\xf0\x53\x34\xa8\x5a\x12\xf5\xb6\xde\xae\x9f\x6b\x91\xfd\xab\x1c\x37\x4d\x8c\x7a\x74\x1a\x42\x78\xd7\xce\x12\xfc\xd0\x26\x56\x0c\x16\x3a\x7a\x88\x1e\xf7\xb1\x07\x4d\x83\xea\xd1\x79\x40\x32\x10\xb8\x64\x89\x65\xd0\x63\xd3\x85\xd8\x1f\xa8\xae\x7b\x69\xb9\x64\xaf\x17\x72\xfc\x4e\x13\x25\xa1\x21\x44\x3f\x6f\x4d\x07\x52\x08\x91\x9b\x13\xdb\x24\x72\x5e\xdd\xda\x3b\x43\x3e\xf5\x29\x04\x4b\x8b\x9e\x3b\x38\x17\x78\xc7\x46\x65\x35\x04\x72\xb8\x0c\xc8\xfe\x3e\x8e\x42\x54\x47\x3e\xdf\xd9\x9b\x0d\xc1\xe1\x7a\xe7\x03\x43\xf5\x83\xdd\x35\x72\x5c\x3c\x65\x2c\xe0\xff\xef\x7b\x8c\x5e\xbb\xf1\xe2\xc9\x0c\x14\x47\xf8\xfa\x36\x08\x4d\x71\x63\x47\x57\xde\x91\x9b\xde\xb2\xe8\x79\x9b\x2e\x72\x55\x3e\x4f\x7a\xc6\x2d\xe3\x8d\x3f\x5f\x95\xa2\xe7\xdf\x00\x00\x00\xff\xff\x69\x16\x74\x58\x24\x04\x00\x00")

func bindataProcesscontextJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataProcesscontextJson,
		"/process_context.json",
	)
}

func bindataProcesscontextJson() (*asset, error) {
	bytes, err := bindataProcesscontextJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/process_context.json",
		size:        1060,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataRenameSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x51\x5d\x6e\xc2\x30\x0c\x7e\xe7\x14\x51\xb4\x47\xc0\xac\x8f\xb9\xc4\x0e\x30\xa1\x29\x6b\xdd\x35\xa8\x24\x99\xe3\x4d\x42\x28\x77\x9f\xd2\x52\x04\x69\xe8\xf0\x4b\xd3\xef\xcf\x8e\x73\x5e\x09\x21\x84\x7c\x09\x75\x87\x47\x2d\x95\x90\x1d\xb3\x0f\x0a\xe0\x10\x9c\xdd\x8c\xf0\xd6\xd1\x17\x34\xa4\x5b\x86\x6a\x57\xed\x36\xaf\x15\x5c\xf4\xeb\x8b\xdd\x34\xc9\x4a\x68\xf5\x11\xb7\xc9\x39\x31\x7c\xf2\x98\x28\xf7\x79\xc0\x9a\x27\x54\xdb\xd3\x5b\x2b\x95\x78\x1f\x7e\x53\x9d\xaf\xa7\x31\x90\x30\xf1\xb2\x35\x3d\x2a\x00\xa8\x9d\x65\x6d\x2c\xd2\x07\xfe\xa2\xe5\xb1\xc5\xd5\x12\xd7\xcf\xe6\x74\x2e\x70\x39\x62\x38\xed\xa7\xf9\xfa\x7e\x79\x3e\x4f\xce\x23\xb1\xc1\x20\x55\xc6\x0d\x7c\xea\x57\x64\x1e\x2f\x65\xa6\x22\xfc\xfe\x31\x84\xcd\xdd\x1c\x33\x55\x83\x81\x8d\xd5\x6c\x6e\x6f\x73\x5b\xfb\x07\xf1\xff\x5c\xa1\xd8\x60\x49\x28\x4a\xfb\x4e\xdf\x6c\xd3\x79\xc5\x22\x33\x47\xef\x91\x98\xbf\xdc\x2a\xfe\x05\x00\x00\xff\xff\x3a\x47\x3d\xb9\xcb\x02\x00\x00")

func bindataRenameSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataRenameSchemaJson,
		"/rename.schema.json",
	)
}

func bindataRenameSchemaJson() (*asset, error) {
	bytes, err := bindataRenameSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/rename.schema.json",
		size:        715,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataSelinuxSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x56\x4b\x6e\xdb\x30\x10\xdd\xeb\x14\x04\x9b\x65\x13\xa6\x5e\xea\x12\x3d\x40\x11\x04\xb4\x34\x8c\x18\x48\xa4\x4a\x8e\x82\x06\x85\xef\x5e\x50\xbf\x58\xe6\xc7\x34\x1a\x1b\x05\xea\x95\xc1\x79\x7c\xf3\x38\x7c\x33\xe2\xef\x82\x10\x42\xe8\x9d\xad\x1a\xe8\x38\x2d\x09\x6d\x10\x7b\x5b\x32\xf6\x6a\xb5\xba\x9f\x96\x1f\xb4\x79\x61\xb5\xe1\x02\xd9\xee\x71\xf7\x78\xff\x6d\xc7\x66\xfc\xd7\x79\xbb\xac\xdd\x56\x0b\xad\x54\xc3\xaf\x07\xb7\x75\x09\xd5\x20\xa4\x92\x28\xb5\xb2\xb4\x24\x53\xba\x31\xb0\xd7\xba\xdd\xac\x8c\xab\xf8\xde\x83\xa3\xd2\xfb\x57\xa8\x70\x26\x59\xa3\xbd\xd1\x3d\x18\x94\x60\xbd\x9d\x71\xce\x3c\xee\xdc\x1c\x2b\x4e\xf1\x0e\x92\x88\x4d\x4e\x8b\x46\xaa\x17\x1a\x05\x1f\xc2\x6a\x46\x0e\x8b\x1c\x33\x52\x81\x1a\x3a\x5a\x92\x1f\x49\xd4\x88\x5c\x6f\x27\x8d\x12\x22\xae\xd7\xfd\x9e\xe2\xa7\x09\x46\x22\x67\xa4\x06\x7e\x0e\xd2\x40\x9d\x14\x3f\xd5\xfb\x6c\x95\x82\x71\x5f\xe8\x56\xe0\x89\xb0\xb4\xa0\xc9\x64\x45\x98\xff\x88\x89\x82\x12\xda\x54\xfe\xc5\xfd\xa5\xc7\x63\xb4\x79\xf4\xb9\x69\x56\x9c\x2b\xeb\x90\xc6\x90\xcb\xdc\x37\x1d\xc0\x75\x43\x86\x09\x7b\x30\x9d\xb4\x56\xbe\xa5\xee\x7e\x45\xd7\xd2\xf2\x7d\x0b\xf5\xbf\xe4\xdb\xb9\x80\x37\x30\xe6\xe2\x8c\x0c\x6f\x3a\x0f\x3f\x57\xba\xeb\x24\x5e\x63\x06\xc7\xa8\xf3\x52\xe4\xa6\x5a\x71\x99\x03\x72\x49\xea\x14\x02\x57\x89\x61\x7c\x25\x1b\xdc\x6c\x3c\x2d\xf5\x8f\x39\xa1\x38\xa2\x0d\xdf\x05\xe5\xea\xfd\xbb\xd8\x64\x38\x71\xc9\x9d\x01\x17\xa7\x42\xb6\x50\x32\xc6\x2a\xad\x90\x4b\x05\xe6\x19\xde\x40\xe1\xf4\x08\x08\x99\xef\x0c\x4f\xa3\x2d\x86\x29\xc6\x7f\x4f\xb3\xbe\x88\x37\x96\x37\xc8\xa5\xae\xd6\x0a\x4e\x8e\x1b\x96\xeb\xc9\xfe\xc2\x8e\x5e\x38\xcc\xff\x36\x90\xb0\x73\x2e\x62\x0d\x36\xf6\x67\x10\x47\xbd\x42\x3c\x13\xc6\xdc\xc3\xdb\x36\xed\x13\x29\xc2\x73\x22\xa3\xb9\x47\x4f\xa4\x5b\x3f\x6b\x90\x90\xec\x36\x1d\x91\x1f\x4e\xee\x39\x36\xf1\x31\x11\xfe\x8e\xf8\xc3\x23\xdd\xcd\xd8\x80\x0a\x57\xe8\xac\xe2\x0f\xa5\xbe\xc8\xad\xb8\x43\x46\x23\xc6\x2e\xea\x33\x65\x64\x9f\xfd\xbf\x75\x47\x64\xdc\x05\xf5\xad\xc3\x6e\x82\x16\x87\xe2\x4f\x00\x00\x00\xff\xff\x09\x35\xe4\x29\xcf\x0d\x00\x00")

func bindataSelinuxSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataSelinuxSchemaJson,
		"/selinux.schema.json",
	)
}

func bindataSelinuxSchemaJson() (*asset, error) {
	bytes, err := bindataSelinuxSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/selinux.schema.json",
		size:        3535,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataUsrJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8e\x41\x0e\x83\x20\x10\x45\xf7\x9c\x62\x42\xba\x54\xb1\x2e\xbd\x8d\x55\xaa\x98\x54\xc8\xcc\xb0\x68\x1a\xee\xde\x20\x6a\x62\xa8\xab\xee\xe0\xfd\xff\x1f\x7c\x04\x00\x80\xbc\x51\x3f\xe9\x57\x27\x5b\x90\x13\xb3\xa3\x56\xa9\x99\xec\x52\x26\x5c\x59\x1c\xd5\x80\xdd\x93\x55\x53\x37\x75\x79\x6f\xd4\xd6\x2f\xb6\xb9\x19\xe2\xd4\x13\x56\x71\xb6\x63\x7e\x3b\x1d\xb9\x7d\xcc\xba\xe7\x9d\x3a\xb4\x4e\x23\x1b\x4d\xb2\x85\xf4\xfe\xca\x3d\xe1\x09\x5c\x2b\x8e\xf4\x42\x75\xe4\xeb\xb7\x72\x7e\x32\x13\xa3\x59\x46\x99\x95\x42\x91\xfb\x46\xb4\xde\xfd\xa1\x14\xbf\x6f\xe9\x14\x44\x10\xdf\x00\x00\x00\xff\xff\x86\xf5\x71\x42\x8f\x01\x00\x00")

func bindataUsrJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataUsrJson,
		"/usr.json",
	)
}

func bindataUsrJson() (*asset, error) {
	bytes, err := bindataUsrJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/usr.json",
		size:        399,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"/chmod.schema.json":      bindataChmodSchemaJson,
	"/chown.schema.json":      bindataChownSchemaJson,
	"/container.json":         bindataContainerJson,
	"/container_context.json": bindataContainercontextJson,
	"/container_event.json":   bindataContainereventJson,
	"/datetime.json":          bindataDatetimeJson,
	"/event.json":             bindataEventJson,
	"/exec.schema.json":       bindataExecSchemaJson,
	"/file.json":              bindataFileJson,
	"/host_event.json":        bindataHosteventJson,
	"/link.schema.json":       bindataLinkSchemaJson,
	"/open.schema.json":       bindataOpenSchemaJson,
	"/process.json":           bindataProcessJson,
	"/process_context.json":   bindataProcesscontextJson,
	"/rename.schema.json":     bindataRenameSchemaJson,
	"/selinux.schema.json":    bindataSelinuxSchemaJson,
	"/usr.json":               bindataUsrJson,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op:   "open",
					Path: name,
					Err:  os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"": {Func: nil, Children: map[string]*bintree{
		"chmod.schema.json":      {Func: bindataChmodSchemaJson, Children: map[string]*bintree{}},
		"chown.schema.json":      {Func: bindataChownSchemaJson, Children: map[string]*bintree{}},
		"container.json":         {Func: bindataContainerJson, Children: map[string]*bintree{}},
		"container_context.json": {Func: bindataContainercontextJson, Children: map[string]*bintree{}},
		"container_event.json":   {Func: bindataContainereventJson, Children: map[string]*bintree{}},
		"datetime.json":          {Func: bindataDatetimeJson, Children: map[string]*bintree{}},
		"event.json":             {Func: bindataEventJson, Children: map[string]*bintree{}},
		"exec.schema.json":       {Func: bindataExecSchemaJson, Children: map[string]*bintree{}},
		"file.json":              {Func: bindataFileJson, Children: map[string]*bintree{}},
		"host_event.json":        {Func: bindataHosteventJson, Children: map[string]*bintree{}},
		"link.schema.json":       {Func: bindataLinkSchemaJson, Children: map[string]*bintree{}},
		"open.schema.json":       {Func: bindataOpenSchemaJson, Children: map[string]*bintree{}},
		"process.json":           {Func: bindataProcessJson, Children: map[string]*bintree{}},
		"process_context.json":   {Func: bindataProcesscontextJson, Children: map[string]*bintree{}},
		"rename.schema.json":     {Func: bindataRenameSchemaJson, Children: map[string]*bintree{}},
		"selinux.schema.json":    {Func: bindataSelinuxSchemaJson, Children: map[string]*bintree{}},
		"usr.json":               {Func: bindataUsrJson, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
