package functions

import "sync"

var lock sync.Once

func init()  {
		lock.Do(func() {
				if function == nil {
						function= &FunctionBindWrapper{}
				}
		})
}
