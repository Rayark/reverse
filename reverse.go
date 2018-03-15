// Copyright © 2015 Alexey Khalyapin - halyapin@gmail.com
//
// This program or package and any associated files are licensed under the
// GNU GENERAL PUBLIC LICENSE Version 2 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.gnu.org/licenses/gpl-2.0.html

package reverse

import (
	"errors"
	"fmt"
	"strings"
)

var Urls *URLStore

func init() {
	Urls = NewURLStore()
}

// Adds url to store
func Add(urlName, urlAddr string, params ...string) string {
	return Urls.MustAdd(urlName, urlAddr, params...)
}

// Reverse url by name
func Rev(urlName string, params ...string) string {
	return Urls.MustReverse(urlName, params...)
}

// Gets raw url by name
func Get(urlName string) string {
	return Urls.Get(urlName)
}

// Gets saved all urls
func GetAllUrls() map[string]string {
	out := map[string]string{}
	for key, value := range Urls.store {
		out[key] = value.url
	}
	return out
}

// Gets all params
func GetAllParams() map[string][]string {
	out := map[string][]string{}
	for key, value := range Urls.store {
		out[key] = value.params
	}
	return out
}

type URL struct {
	url    string
	params []string
}

type URLStore struct {
	store map[string]URL
}

// create a new URLStore
func NewURLStore() *URLStore {
	return &URLStore{store: make(map[string]URL)}
}

// Adds a Url to the Store
func (us *URLStore) Add(urlName, urlAddr string, params ...string) (string, error) {
	if _, ok := us.store[urlName]; ok {
		return "", errors.New("Url already exists. Try to use .Get() method.")
	}

	tmpUrl := URL{urlAddr, params}
	us.store[urlName] = tmpUrl
	return urlAddr, nil
}

// Adds a Url and panics if error
func (us *URLStore) MustAdd(urlName, urlAddr string, params ...string) string {
	addr, err := us.Add(urlName, urlAddr, params...)
	if err != nil {
		panic(err)
	}
	return addr
}

// Append a URLStore with prefix
func (us *URLStore) Append(prefix string, group *URLStore) error {
	for name, url := range group.store {
		_, err := us.Add(name, prefix+url.url, url.params...)
		if err != nil {
			return err
		}
	}
	return nil
}

// Append a URLStore with prefix
func (us *URLStore) MustAppend(prefix string, store *URLStore) {
	err := us.Append(prefix, store)
	if err != nil {
		panic(err)
	}
}

// Gets raw url string
func (us *URLStore) Get(urlName string) string {
	return us.store[urlName].url
}

// Gets reversed url
func (us *URLStore) Reverse(urlName string, params ...string) (string, error) {
	if len(params) != len(us.store[urlName].params) {
		return "", errors.New("Bad Url Reverse: mismatch params for URL: "+ urlName)
	}
	res := us.store[urlName].url
	for i, val := range params {
		res = strings.Replace(res, us.store[urlName].params[i], val, 1)
	}
	return res, nil
}

// Gets reversed url and panics if error
func (us *URLStore) MustReverse(urlName string, params ...string) string {
	res, err := us.Reverse(urlName, params...)
	if err != nil {
		panic(err)
	}
	return res
}

func (us *URLStore) Rev(urlName string, params ...string) string {
	return us.MustReverse(urlName, params...)
}

func (us *URLStore) Sting() string {
	return fmt.Sprint(us.store)
}

// For testing
func (us *URLStore) getParam(urlName string, num int) string {
	return us.store[urlName].params[num]
}
