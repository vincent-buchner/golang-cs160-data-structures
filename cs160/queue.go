
package main 

import(
  //"container/list"
  "fmt"
)



func main(){
  queue := Queue{} 

  queue.add(1)
  queue.add(2)


  for _, value := range queue.items {
    fmt.Println(value)
  }
}


type Queue struct {
  items []int
}

func (queue *Queue) add(element int){
  queue.items = append(queue.items, element)
  return 
}

func (queue *Queue) remove(element int){
  //first_item := queue.items[0]
  queue.items = queue.items[1:]
}

func (queue *Queue) peek() (front_element int){
  front_element = queue.items[0] 
  return front_element 
}

func (queue *Queue) isEmpty() (result bool){
  if len(queue.items) == 0{
    result = true 
    return result 
  }
  result = false 
  return result
}





