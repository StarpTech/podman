package images

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

/*
This file is generated automatically by go generate.  Do not edit.
*/

// Changed
func (o *PushOptions) Changed(fieldName string) bool {
	r := reflect.ValueOf(o)
	value := reflect.Indirect(r).FieldByName(fieldName)
	return !value.IsNil()
}

// ToParams
func (o *PushOptions) ToParams() (url.Values, error) {
	params := url.Values{}
	if o == nil {
		return params, nil
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	s := reflect.ValueOf(o)
	if reflect.Ptr == s.Kind() {
		s = s.Elem()
	}
	sType := s.Type()
	for i := 0; i < s.NumField(); i++ {
		fieldName := sType.Field(i).Name
		if !o.Changed(fieldName) {
			continue
		}
		fieldName = strings.ToLower(fieldName)
		f := s.Field(i)
		if reflect.Ptr == f.Kind() {
			f = f.Elem()
		}
		switch f.Kind() {
		case reflect.Bool:
			params.Set(fieldName, strconv.FormatBool(f.Bool()))
		case reflect.String:
			params.Set(fieldName, f.String())
		case reflect.Int, reflect.Int64:
			// f.Int() is always an int64
			params.Set(fieldName, strconv.FormatInt(f.Int(), 10))
		case reflect.Uint, reflect.Uint64:
			// f.Uint() is always an uint64
			params.Set(fieldName, strconv.FormatUint(f.Uint(), 10))
		case reflect.Slice:
			typ := reflect.TypeOf(f.Interface()).Elem()
			switch typ.Kind() {
			case reflect.String:
				sl := f.Slice(0, f.Len())
				s, ok := sl.Interface().([]string)
				if !ok {
					return nil, errors.New("failed to convert to string slice")
				}
				for _, val := range s {
					params.Add(fieldName, val)
				}
			default:
				return nil, errors.Errorf("unknown slice type %s", f.Kind().String())
			}
		case reflect.Map:
			lowerCaseKeys := make(map[string][]string)
			iter := f.MapRange()
			for iter.Next() {
				lowerCaseKeys[iter.Key().Interface().(string)] = iter.Value().Interface().([]string)

			}
			s, err := json.MarshalToString(lowerCaseKeys)
			if err != nil {
				return nil, err
			}

			params.Set(fieldName, s)
		}
	}
	return params, nil
}

// WithAll
func (o *PushOptions) WithAll(value bool) *PushOptions {
	v := &value
	o.All = v
	return o
}

// GetAll
func (o *PushOptions) GetAll() bool {
	var all bool
	if o.All == nil {
		return all
	}
	return *o.All
}

// WithAuthfile
func (o *PushOptions) WithAuthfile(value string) *PushOptions {
	v := &value
	o.Authfile = v
	return o
}

// GetAuthfile
func (o *PushOptions) GetAuthfile() string {
	var authfile string
	if o.Authfile == nil {
		return authfile
	}
	return *o.Authfile
}

// WithCompress
func (o *PushOptions) WithCompress(value bool) *PushOptions {
	v := &value
	o.Compress = v
	return o
}

// GetCompress
func (o *PushOptions) GetCompress() bool {
	var compress bool
	if o.Compress == nil {
		return compress
	}
	return *o.Compress
}

// WithPassword
func (o *PushOptions) WithPassword(value string) *PushOptions {
	v := &value
	o.Password = v
	return o
}

// GetPassword
func (o *PushOptions) GetPassword() string {
	var password string
	if o.Password == nil {
		return password
	}
	return *o.Password
}

// WithSkipTLSVerify
func (o *PushOptions) WithSkipTLSVerify(value bool) *PushOptions {
	v := &value
	o.SkipTLSVerify = v
	return o
}

// GetSkipTLSVerify
func (o *PushOptions) GetSkipTLSVerify() bool {
	var skipTLSVerify bool
	if o.SkipTLSVerify == nil {
		return skipTLSVerify
	}
	return *o.SkipTLSVerify
}

// WithUsername
func (o *PushOptions) WithUsername(value string) *PushOptions {
	v := &value
	o.Username = v
	return o
}

// GetUsername
func (o *PushOptions) GetUsername() string {
	var username string
	if o.Username == nil {
		return username
	}
	return *o.Username
}
