package storekeeper

import (
	//"errors"
	"log"
	"os"
	"reflect"
)

type Store struct {
	instance map[interface{}]interface{}
	binding  map[interface{}]interface{}
}

// Initialization or create
func New() *Store {
	storage := &Store{}
	storage.instance = make(map[interface{}]interface{})
	storage.binding = make(map[interface{}]interface{})
	return storage
}

func (store *Store) SetInstance(abstract interface{}, instance interface{}) *Store {
	store.instance[abstract] = instance
	return store
}

// extract from storage
func (store *Store) Extract(abstract interface{}) interface{} {
	if instance, exist := store.instance[abstract]; exist {
		return instance
	}
	if _, exist := store.binding[abstract]; exist {
		// TODO here it is necessary to determine two things : what is type (ptr... example), what first argument
		// TODO it is Store struct
		values, _ := store.call(store.binding, `name`, store)
		log.Println(`values`, values)
		os.Exit(2)
	}
	return nil
}

func (store *Store) Bind(abstract interface{}, concrete interface{}) *Store {
	store.binding[abstract] = concrete

	return store
}

// function call closure and return ready structure
func (store *Store) call(m map[interface{}]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	countArg := store.countArgumentsClosure(f)
	//if len(params) != f.Type().NumIn() {
	//	err = errors.New("The number of params is not adapted.")
	//	return
	//}

	in := make([]reflect.Value, countArg)
	if countArg > 0 {
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
		}
	}

	result = f.Call(in)
	// TODO need to come up with how return only first index which us need
	log.Println(result, `result`)
	os.Exit(2)
	return
}

// get count arguments in value, mean function (closure)
func (store *Store) countArgumentsClosure(function reflect.Value) int {
	return function.Type().NumIn()
}
