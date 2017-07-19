# Links3
A solution to an interview question I was once asked.

## Background

A long time ago, I was asked a an algorithmic question in an interview that made
a huge impression on me. I'll tell the full story some other time on my blog.
The question started off very simple, and then ramped up in difficulty:

* implement a linked list
* show how you can add/subtract elements from this list
* write a function to reverse this linked list

and finally, the part where I got stuck:

* write a function to reverse every group of three within a linked list

Although I ended up getting the job, this question has haunted me for many
years. I've been able to solve it in my own time, but it was never _elegant_.
Each time I'd have to struggle with dangling pointers and weird edge cases until
I got it right. Then, a few days ago, I tried to solve the problem again and
found a way of breaking up the problem into a better solution.

## link.go

This github repo contains the package which defines link.go, a go package that
defines a linked list of integers.

`Cons` is the function that is used to construct and add to a linked list by
creating a new element as the head of the list.

As a warm up and an aid to testing, I also defined the functions Len, Reverse
and String, which respectively returns the length of the list, reverses the list
in place and returns the string representation of the list.

## ReverseBy3s

Finally, the package defines the function ReverseBy3s, which does the job of
reversing every group of 3 elements in the linked list. The trick which makes
this easy to implement is the helper function ReverseFirstN. ReverseBy3s
repeatedly calls ReverseFirstN to split off the first 3 elements and reverse
them. It joins the pieces as it goes, and then returns the completed list.

Have a look at link/link_test.go for example code. You may run it with

```
! go test github.com/danverbraganza/Links3/link/
```

This code is not meant to be used in production. It just demonstrates a way to
simplify the algorithm of reversing by 3s. I'm planning to tell the whole story
and write a more thorough explanation of the code on my own blog soon.