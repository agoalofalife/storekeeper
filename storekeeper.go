package storekeeper

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
	return nil
}
