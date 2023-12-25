
package main

import(
  "container/list"
  "fmt"
    )

func main(){
  // this creates a new list 
  // := acts as both declaration of 
  // a variable and assignment 
  l := list.New()

  e4 := l.PushBack(4) 
  e1 := l.PushFront(1)

  l.InsertBefore(3, e4) 
  l.InsertAfter(2, e1)
  add(l, 7)
  // iterate throught the list ans print 
  // its contents 
  for e:=l.Front(); e !=nil; e=e.Next(){
    fmt.Println(e.Value)
  }
}

// Methods
//Go does not have classes. However, you can define methods on types.

//A method is a function with a special receiver argument.

//The receiver appears in its own argument list between the func keyword and the method name.

//In this example, the Abs method has a receiver of type Vertex named v.


//func (v Vertex) Abs() float64 {

func add(list *list.List, element int){

  list.PushBack(element)
  return 
}

func remove(list *list.List, element int){
  // left off here 
}

func peek(list *list.List) (frontElement int){



}

func isEmpty(list *list.List) (result bool){
  
}

func isFull(list *list.List) (result bool){
  
}


