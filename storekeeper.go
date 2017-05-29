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
	if bind, exist := store.binding[abstract]; exist {
		// TODO here it is necessary to determine two things : what is type (ptr... example), what first argument
		// TODO it is Store struct
		//log.Println(reflect.TypeOf(bind).Kind().String())
		//os.Exit(2)
		switch reflect.TypeOf(bind).Kind().String() {
		case `func`:
			values, _ := store.call(store.binding, abstract.(string), store)
			instance := values[0].Interface()
			store.SetInstance(abstract, instance)
			return instance
		case `slice`:
			log.Println(bind.([]interface{}))
			os.Exit(2)
			store.verifySliceBind(bind.([]interface{}))
			log.Println(bind)
			os.Exit(2)

		}

	}
	return nil
}

// bind some structure
func (store *Store) Bind(abstract interface{}, concrete interface{}) *Store {
	store.binding[abstract] = concrete
	return store
}

// function call closure and return ready structure
func (store *Store) call(m map[interface{}]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	countArg := store.countArgumentsClosure(f)

	in := make([]reflect.Value, countArg)
	if countArg > 0 {
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
		}
	}

	result = f.Call(in)
	return
}

// get count arguments in value, mean function (closure)
func (store *Store) countArgumentsClosure(function reflect.Value) int {
	return function.Type().NumIn()
}

// get slice check what :
// first parameter this is structure
// second method or so string
func (store *Store) verifySliceBind(slice []interface{}) {
	//log.Println(reflect.TypeOf(slice))
	//log.Println(reflect.TypeOf(slice))
	os.Exit(2)
}
