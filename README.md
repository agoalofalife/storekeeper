


### Storekeeper


**This is an attempt to organize the container for Golang, the idea is that we have a repository structures to which we can flexibly apply.**



- [Installation](#Installation)
- [Fast Example](#Fast_Example)




<a name="Fast_Example"></a>
## Installation

    go get github.com/agoalofalife/storekeeper
    
<a name="Example_fast"></a>
## Fast_Example

When we need to put structure and method constructor
    
    
    import ( 
    store "github.com/agoalofalife/storekeeper"
	      "fmt"
    )
       type People struct {
    	name string
    	Address string
    }
    // this is any constructor for People
    func (people *People) Constructor(){
    	people.name = name{}
    	people.Address = `Some Street some House ...`
    }
    // always initialize first
    storage := store.New()
    // Method Bind get first argument any type (preferably string)
    
    // Two argumetn slice where first it is struct and second method constructor
    storage.Bind(`People`, []interface{}{&People{}, `Constructor`})
    
    // get result
	p := storage.Extract(`People`)
	
    // p and has initialized struct and run method Constructor 
    // &{{} Some Street some House ...}
    // for further use it is necessary to give type
    p.(*People)
    // ! please note when handling the type depend of pointer or not

 




