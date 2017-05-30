package storekeeper

import (
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"reflect"
)

// list possible errors in current application
const (
	ERROR_NOT_SPECIFIED_STRUCT_IN_BIND = `expected struct in first parameter`
	ERROR_NOT_SPECIFIED_METHOD_IN_BIND = `the method is not a string and not found in the structure`
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

		switch reflect.TypeOf(bind).Kind().String() {
		case `func`:
			values, _ := store.call(store.binding, abstract.(string), store)
			instance := values[0].Interface()
			store.SetInstance(abstract, instance)
			return instance
		case `slice`:
			store.verifySliceBind(bind.([]interface{}))
			instance := bind.([]interface{})[0]
			store.callFS(instance, bind.([]interface{})[1].(string))
			store.SetInstance(abstract, instance)
			return instance
		}
		log.Println()
		os.Exit(2)
	}
	return nil
}

// bind some structure
func (store *Store) Bind(abstract interface{}, concrete interface{}) *Store {
	store.binding[abstract] = concrete
	return store
}

// call from structure
func (store *Store) callFS(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	return reflect.ValueOf(any).MethodByName(name).Call(inputs)
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
	firstArg := reflect.TypeOf(slice[0]).Kind().String()

	if firstArg != `struct` && firstArg != `ptr` {
		panic(ERROR_NOT_SPECIFIED_STRUCT_IN_BIND)
	}

	_, existMethod := reflect.TypeOf(slice[0]).MethodByName(slice[1].(string))

	if reflect.TypeOf(slice[1]).Kind().String() != `string` || existMethod == false {
		panic(ERROR_NOT_SPECIFIED_METHOD_IN_BIND)
	}
}

// --- INFO
func (store *Store) getStructTag(f reflect.StructField) string {
	return string(f.Name)
}
func (store *Store) State() bool {
	//types := []string{`instance`, `binding`}
	field, _ := reflect.TypeOf(store).Elem().FieldByName(`instance`)
	log.Println(store.getStructTag(field))
	os.Exit(2)
	data := [][]string{
		[]string{"A", "The Good", "500"},
		[]string{"B", "The Very very Bad Man", "288"},
		[]string{"C", "The Ugly", "120"},
		[]string{"D", "The Gopher", "800"},
	}
	data = [][]string{}

	for v, _ := range store.instance {
		log.Println(v.(string), `logger`)
		os.Exit(2)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Type", "Key", "Value"})
	table.SetBorder(true)  // Set Border to false
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
	return true
}
