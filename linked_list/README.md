<!-- vim-markdown-toc GFM -->

-  [Restrictions](#restrictions)
-  [Implementation](#implementation)
   -  [Iterator](#iterator)
-  [Part 3: Recursive function (5 marks)](#part-3-recursive-function-5-marks)
   -  [Online visualizer](#online-visualizer)

<!-- vim-markdown-toc -->

## Restrictions

**No STL containers classes(such as vector, lists, etc.) or library functions(such as sort) may be used for this assignment. Please ask if you aren't sure when using any part of the Standard Library.**

## Implementation

Implement a template for a doubly linked linked list.

Nodes in this list will have a pointer not only to the next node but also a pointer to the previous node.

The data type of the data stored in this list (T) supports the following operator:

-  ==
-  !=
-  =
-  <
   This means that you can use any of the above operators on the data stored in the nodes of your linked list. You do not need to implement these operators. It is up to the program using the linked list to implement these operators for the data type.

The DList class has the following member functions:

```c++
DList();
```

-  constructor, creates empty DList

---

```c++
iterator insert(iterator it, const T& data);
```

-  adds new node containing **_data_** before the node refer to by **_it_**
-  function returns iterator to newly added node

---

```c++
iterator begin();
```

-  returns iterator to Node containing the first piece of data in the linked list

---

```c++
iterator end();
```

-  returns iterator to the Node **_after_** the node containing the last piece of data of the linked list

---

```c++
const_iterator cbegin() const;
```

-  returns const_iterator to to Node containing the first piece of data in the linked list

---

```c++
const_iterator cend() const;
```

-  returns const*iterator to the Node \*\*\_after*\*\* the node containing the last piece of data of the linked list

---

```c++
iterator erase(iterator it);
```

-  removes the node referred to by **_it_**
-  returns iterator to node that follows the node that was removed

---

```c++
void sort(iterator first, iterator last);
```

-  This function sorts all nodes from **_first_** to **_last_**, including **_first_** but not including **_last_**. This process must not create any new nodes or deallocate any old nodes. Furthermore, the data within each node is not to be swapped around... the point is to be able to move nodes around to sort. **Recreating the list, using temporary arrays to sort or swapping data without moving nodes will be graded as incorrect and you will need to resubmit your assignment if you do this**

**For example:**

In the example below

Initial current list: {1, 3, 4, 8, 2, 4, 7, 3, 6, 5, 9, 10}, suppose that **_first_** refers to the node with 8, **_last_** refers to the node with 5.

The function would sort the nodes with the bolded values:{1, 3, 4, **8,2,4,7,3,6,\*** 5, 9, 10}

The final list would end up with values in the following order: {1, 3, 4, 2, 3, 4, 6, 7, 8, 5, 9, 10}

---

```c++
iterator search(const T& data);
const_iterator search(const T& data) const;
```

-  returns iterator to the node containing data. If data is not found, returns end()

---

```c++
bool empty() const;
```

-  function returns true if the list is empty, false otherwise

---

```c++
int size() const;
```

-  function returns number of pieces of data stored in the list

---

```c++
~DList();
DList(const DList&);
DList& operator=(const DList&);
DList(DList&&);
DList& operator=(DList&&);
```

-  Your sorted linked list must also implement destructor, copy constructor, assignment operator, move constructor, move operator.

### Iterator

The idea of an iterator is to provide a means to traverse your class. In the STL, different classes like vectors or list all have iterators that help you iterate through your list. For the sorted list, you will also write an iterator class.

You will need two iterators. a const_iterator and an iterator which is derived from const_iterator. For both operators, the following public members/operators are needed. You are welcome to add any other private or protected members as you see fit:

```c
iterator();
const_iterator();
```

constructors, returns iterators to nullptrs

---

```c
iterator& operator++();
const_iterator& operator++() const;
iterator operator++(int);
const_iterator operator++(int);
```

-  prefix and postfix ++ operator
-  makes iterator point to the next node in the list.
-  return as appropriate for each operator

---

```c
iterator& operator--();
const_iterator& operator--() const;
iterator operator--(int);
const_iterator operator--(int);

```

-  prefix and postfix operator --
-  increments iterator so that it points to previous node in list
-  return as appropriate for each operator

---

```
operator==
```

returns true if two iterators point at the same node, false otherwise

---

```c
operator !=
```

-  returns true if two iterators point at different nodes, false otherwise

---

```c
const T& operator*()const; (in both iterator and const_iterator)
T& operator*(); (only in iterator)
```

-  dereferencing operator, returns the data in the node pointed to by iterator

## Part 3: Recursive function (5 marks)

We describe a maze as having row X col cells. For example if row was 3, and col was 4, then we would have a grid of cells as follows. We describe a wall by the two cell numbers the wall separates (smaller number first). If every single wall existed, there would be (row-1)(col) + (col-1)(row) walls.

```
  0 |  1 |  2 |  3
-------------------
  4 |  5 |  6 |  7
--------------------
  8 |  9 | 10 | 11
```

A Maze class (which you do not need to implement) describes a maze using the method described It has member functions that you can use to travel through the maze (ie figure out where you are, know what cells you can reach, etc.)

Write a recursive maze runner:

```c++
int runMaze(Maze& theMaze, int path[], int fromCell, int toCell);
```

The runMaze() function will find a path from cell number **_fromCell_** to cell number **_toCell_**.
function will "markup" theMaze with the path and pass back an array containing the path (using the cell numbers) from the starting cell to ending cell. The function will return the number of cells along the pathway from the starting cell to the ending cell inclusive.

For example, suppose the fromCell was 0 and the toCell was 3 using the maze below:

```
  0 |  1    2    3
         ----------
  4    5 |  6    7
----
  8    9   10 | 11
```

runMaze() function would put the following into path: {0,4,5,1,2,3} and return 6

### Online visualizer

To help you debug your program, when you run the tester, the tester will create a "path" file based on the path your runMaze() generates and its return value. You can see what is happening by going to this site:

https://seneca-dsa555-f21.github.io/dsa555-f21/

-  Use the radio buttons to select the test in question (see your error message)
-  then load the corresponding testpath file.
