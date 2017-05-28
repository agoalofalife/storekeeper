package storekeeper

import (
	"errors"
	"log"
	"os"
	"reflect"
)

type store struct {
	instance map[interface{}]interface{}
	binding  map[interface{}]interface{}
}

// Initialization or create
func New() *store {
	storage := &store{}
	storage.instance = make(map[interface{}]interface{})
	storage.binding = make(map[interface{}]interface{})
	return storage
}

func (store *store) SetInstance(abstract interface{}, instance interface{}) *store {
	store.instance[abstract] = instance
	return store
}

// extract from storage
func (store *store) Extract(abstract interface{}) interface{} {
	if instance, exist := store.instance[abstract]; exist {
		return instance
	}
	if _, exist := store.binding[abstract]; exist {
		values, _ := store.call(store.binding, `name`, store)
		log.Println(`values`, values)
		os.Exit(2)
	}
	return nil
}

func (store *store) Bind(abstract interface{}, concrete interface{}) *store {
	store.binding[abstract] = concrete

	return store
}

// function call closure and return ready structure
func (store *store) call(m map[interface{}]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	result = f.Call(in)
	log.Println(result, `result`)
	os.Exit(2)
	return
}
