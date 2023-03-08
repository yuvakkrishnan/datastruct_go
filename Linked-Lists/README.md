In Go language, a linked list is a collection of nodes, where each node contains a value and a reference to the next node in the list. Here are some of the methods that can be used with linked lists in Go:

1.Adding elements:

func (list *LinkedList) add(data int) - appends a new node with the given data to the end of the list
2.Removing elements:

func (list *LinkedList) remove(data int) - removes the first node with the given data from the list
3.Searching for elements:

func (list *LinkedList) contains(data int) bool - returns true if the list contains a node with the given data
4.Traversing the list:

func (list *LinkedList) print() - iterates through the list and prints out the data in each node
func (list *LinkedList) toSlice() []int - returns a new slice containing the data from each node in the list, in order
Note that these are just some examples of methods that can be used with linked lists in Go. The specific methods and their implementation may vary depending on the use case and requirements of your program.