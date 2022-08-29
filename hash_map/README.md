<!-- vim-markdown-toc GFM -->

-  [Overview](#overview)
   -  [LPTable](#lptable)
   -  [Chaining Table](#chaining-table)

<!-- vim-markdown-toc -->

## Overview

The records for this table consists of key-value pairs. The key is a string. The value can be any data type (T) that supports the following:

-  assignment (=)
-  instantiation without arguments (you can declare variables of type T, the value of which is undefined)

In the description below the word Record refer to a key-value pair. The way you store it (whether you make a template struct called Record with a data member for key and value, or you store them separately in parallel structures is entirely up to you).

Here are the operations on Table:

```c++
Table();
```

Initializes a Table to an empty table

---

```c++
bool update(const string& key, const TYPE& value);
```

This function is passed a key-value pair. If your table already has a record with a matching key, the record's value is replaced by the value passed to this function. If no record exists, a record with key-value pair is added to the table as a new record. function returns true if successful, false otherwise

---

```c++
bool find(const string& key,TYPE& value );
```

This function is passed a key and a reference for passing back a found value. If your table contains a record with a matching key, the function returns true, and passes back the value from the record. If it does not find a record with a matching key, function returns false.

---

```c++
bool remove(const string& key);
```

This function is passed a key. If your table contains a record with a matching key, the record (both the key and the value) is removed from the table

---

```c++
int numRecords() const
```

This function returns the number of records in the table.

---

```c++
bool isEmpty() const
```

This function returns the number of records in the table.

---

Aside from the above functions, all tables must have proper destructors, copy constructors/assignment operators, move constructor/operator that allow for resources to be properly managed.

---

A hash table organizes the records as an array, where records are placed
according to a hash value that you get from a hash function. Please use the
hash function from the C++ STL. You can find more info about how to do it here:

[[http://en.cppreference.com/w/cpp/utility/hash]]

Here is a simple example of how to use the hash function with string keys:

```cpp
#include <iostream>
#include <string>
//you need to include functional to use hash functions
#include <functional>
const int cap = 100;
int main(void){
    //this is just a string... any string can be a key
    std::string key = "this is my key";

    //declare an hash function object.  hashFunction is a variable but it is also a function
    std::hash<std::string> hashFunction;

    //you can now call hashFunction and it will return an unsigned whole number
    size_t hash = hashFunction(key);
    std::cout << hash << std::endl;

    //if you need to change it so that it is a proper index for a table with capacity of cap
    size_t idx = hash % cap;
    std::cout << idx << std::endl;
}
```

You will implement two hash tables for this assignment. One will use linear
probing for collision resolution, the other will use chaining.

### LPTable

LPTable is a hash table that uses linear probing as the collision resolution method.

The constructor for the LPTable accepts one argument

```
LPTable(size_t capacity);
```

**capacity** is the capacity (number of elements for the array) of the table.

---

```c++
bool update(const string& key, const TYPE& value);
```

This function is passed a key-value pair. If your table already has a record with a matching key, the record's value is replaced by the value passed to this function. If no record exists, a record with key-value pair is added to the table.

The table can only hold capacity - 1 records (need 1 open spot minimum). adding a new record to a table that already has capacity - 1 records will result in an unsuccessful operation. The new record is not to be added.

Function returns true if update() was successful, false otherwise.

---

```c++
bool find(const string& key,TYPE& value );
```

This function is passed a key and a reference for passing back a found value. If your table contains a record with a matching key, the function returns true, and passes back the value from the record. If it does not find a record with a matching key, function returns false.

---

```c++
bool remove(const string& key);
```

This function is passed a key. If your table contains a record with a matching key, the record (both the key and the value) is removed from the table

---

```c++
int numRecords() const
```

This function returns the number of records in the table.

---

```c++
bool isEmpty() const
```

This function returns the number of records in the table.

---

```c++
size_t capacity() const
```

This function returns capacity of the table.

---

```c++
double loadFactor() const
```

This function returns the load factor of the table. (number of records divided by capacity of table)

---

**Aside from the above functions, all tables must have proper destructors, copy constructors/assignment operators, move constructor/operator that allow for resources to be properly managed.**

### Chaining Table

ChainingTable is a hash table that uses chaining as the collision resolution method.

It is recommended that you use your assignment 1 linked list to help you in this part of the assignment. To do so, set up a struct for to combine your key-value pairs into a single object. Then add ==, !=, and < comparison operators to compare these according to its key

The constructor for the ChainingTable accepts one argument

```
ChainingTable(size_t capacity);
```

**capacity** is the capacity (number of elements for the array) of the table.

---

```c++
bool update(const string& key, const TYPE& value);
```

This function is passed a key-value pair. If your table already has a record with a matching key, the record's value is replaced by the value passed to this function. If no record exists, a record with key-value pair is added to the table.

This function always returns true if operation was successful, false otherwise

---

```c++
bool find(const string& key,TYPE& value );
```

This function is passed a key and a reference for passing back a found value. If your table contains a record with a matching key, the function returns true, and passes back the value from the record. If it does not find a record with a matching key, function returns false.

---

```c++
bool remove(const string& key);
```

This function is passed a key. If your table contains a record with a matching key, the record (both the key and the value) is removed from the table

---

```c++
int numRecords() const
```

This function returns the number of records in the table.

---

```c++
bool isEmpty() const
```

This function returns the number of records in the table.

---

```c++
size_t capacity() const
```

This function returns capacity of the table.

---

```c++
double loadFactor() const
```

This function returns the load factor of the table. (number of records divided by capacity of table)

---

**Aside from the above functions, all tables must have proper destructors, copy constructors/assignment operators, move constructor/operator that allow for resources to be properly managed.**
