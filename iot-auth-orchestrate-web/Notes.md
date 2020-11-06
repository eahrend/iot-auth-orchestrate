# React notes


## Challenge 
1. Create a clear all button
1. Create an active friends list
    1. Submitting a friend adds users to this list
    1. Create a remove button that removes the friend
    1. Create a deactivate button that moves them over to inactive list
1. Create an inactive friends list
    1. Create an activate button that moves them over to an active friend list

## arrow functions
```
// Traditional Function
function (a){
  return a + 100;
}

// Arrow Function Break Down

// 1. Remove the word "function" and place arrow between the argument and opening body bracket
(a) => {
  return a + 100;
}

// 2. Remove the body brackets and word "return" -- the return is implied.
(a) => a + 100;

// 3. Remove the argument parentheses
a => a + 100;
```


## js map
Definition and Usage
The map() method creates a new array with the results of calling a function for every array element.

The map() method calls the provided function once for each element in an array, in order.

Note: map() does not execute the function for array elements without values.

Note: this method does not change the original array.



## div
html container

## DOM
document object model, when a page is loaded it creates a tree of objects

## Elements
Object representation of the DOM

## Components
A component is a function/class that takes input, and provides a react
element