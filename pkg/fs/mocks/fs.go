// Code generated by mockery v1.1.1. DO NOT EDIT.

package mocks

import (
	billy "github.com/go-git/go-billy/v5"

	iofs "io/fs"

	mock "github.com/stretchr/testify/mock"
)

// FS is an autogenerated mock type for the FS type
type FS struct {
	mock.Mock
}

// CheckExistsOrWrite provides a mock function with given fields: path, data
func (_m *FS) CheckExistsOrWrite(path string, data []byte) (bool, error) {
	ret := _m.Called(path, data)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, []byte) bool); ok {
		r0 = rf(path, data)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte) error); ok {
		r1 = rf(path, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Chroot provides a mock function with given fields: path
func (_m *FS) Chroot(path string) (billy.Filesystem, error) {
	ret := _m.Called(path)

	var r0 billy.Filesystem
	if rf, ok := ret.Get(0).(func(string) billy.Filesystem); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(billy.Filesystem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: filename
func (_m *FS) Create(filename string) (billy.File, error) {
	ret := _m.Called(filename)

	var r0 billy.File
	if rf, ok := ret.Get(0).(func(string) billy.File); ok {
		r0 = rf(filename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(billy.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exists provides a mock function with given fields: path
func (_m *FS) Exists(path string) (bool, error) {
	ret := _m.Called(path)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExistsOrDie provides a mock function with given fields: path
func (_m *FS) ExistsOrDie(path string) bool {
	ret := _m.Called(path)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Join provides a mock function with given fields: elem
func (_m *FS) Join(elem ...string) string {
	_va := make([]interface{}, len(elem))
	for _i := range elem {
		_va[_i] = elem[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(...string) string); ok {
		r0 = rf(elem...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Lstat provides a mock function with given fields: filename
func (_m *FS) Lstat(filename string) (iofs.FileInfo, error) {
	ret := _m.Called(filename)

	var r0 iofs.FileInfo
	if rf, ok := ret.Get(0).(func(string) iofs.FileInfo); ok {
		r0 = rf(filename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iofs.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MkdirAll provides a mock function with given fields: filename, perm
func (_m *FS) MkdirAll(filename string, perm iofs.FileMode) error {
	ret := _m.Called(filename, perm)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, iofs.FileMode) error); ok {
		r0 = rf(filename, perm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Open provides a mock function with given fields: filename
func (_m *FS) Open(filename string) (billy.File, error) {
	ret := _m.Called(filename)

	var r0 billy.File
	if rf, ok := ret.Get(0).(func(string) billy.File); ok {
		r0 = rf(filename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(billy.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenFile provides a mock function with given fields: filename, flag, perm
func (_m *FS) OpenFile(filename string, flag int, perm iofs.FileMode) (billy.File, error) {
	ret := _m.Called(filename, flag, perm)

	var r0 billy.File
	if rf, ok := ret.Get(0).(func(string, int, iofs.FileMode) billy.File); ok {
		r0 = rf(filename, flag, perm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(billy.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, iofs.FileMode) error); ok {
		r1 = rf(filename, flag, perm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadDir provides a mock function with given fields: path
func (_m *FS) ReadDir(path string) ([]iofs.FileInfo, error) {
	ret := _m.Called(path)

	var r0 []iofs.FileInfo
	if rf, ok := ret.Get(0).(func(string) []iofs.FileInfo); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]iofs.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Readlink provides a mock function with given fields: link
func (_m *FS) Readlink(link string) (string, error) {
	ret := _m.Called(link)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(link)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(link)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: filename
func (_m *FS) Remove(filename string) error {
	ret := _m.Called(filename)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(filename)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rename provides a mock function with given fields: oldpath, newpath
func (_m *FS) Rename(oldpath string, newpath string) error {
	ret := _m.Called(oldpath, newpath)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(oldpath, newpath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Root provides a mock function with given fields:
func (_m *FS) Root() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Stat provides a mock function with given fields: filename
func (_m *FS) Stat(filename string) (iofs.FileInfo, error) {
	ret := _m.Called(filename)

	var r0 iofs.FileInfo
	if rf, ok := ret.Get(0).(func(string) iofs.FileInfo); ok {
		r0 = rf(filename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iofs.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Symlink provides a mock function with given fields: target, link
func (_m *FS) Symlink(target string, link string) error {
	ret := _m.Called(target, link)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(target, link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TempFile provides a mock function with given fields: dir, prefix
func (_m *FS) TempFile(dir string, prefix string) (billy.File, error) {
	ret := _m.Called(dir, prefix)

	var r0 billy.File
	if rf, ok := ret.Get(0).(func(string, string) billy.File); ok {
		r0 = rf(dir, prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(billy.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(dir, prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}