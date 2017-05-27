package storekeeper

type store struct {
  instance map[interface{}]interface{}
  binding  map[interface{}]interface{}

}
// Initialization or create
func New() *store {
  storage := &store{}
  return storage
}

func (store *store) SetInstance(abstract interface{}, instance interface{}) {

}

