# Exercises in this month:

### 2021-02-16 builder-pattern

The main goal of builder pattern is simplify construction of complex types.

### 2021-02-18 Time overlapping

Given an array of time intervals (start, end) for classroom lectures (possibly overlapping), find the minimum number of rooms required.

For example, given [(30, 75), (0, 50), (60, 150)], you should return 2.

### 2021-02-18 Locking tree

Implement locking in a binary tree. A binary tree node can be locked or unlocked only if all of its descendants or ancestors are not locked.

Design a binary tree node class with the following methods:

is_locked, which returns whether the node is locked
lock, which attempts to lock the node. If it cannot be locked, then it should return false. Otherwise, it should lock it and return true.
unlock, which unlocks the node. If it cannot be unlocked, then it should return false. Otherwise, it should unlock it and return true.
You may augment the node to add parent pointers or any other property you would like. You may assume the class is used in a single-threaded program, so there is no need for actual locks or mutexes. Each method should run in O(h), where h is the height of the tree.


### 2021-02-18 Abstract factory

Suppose we want to build a global car factory. If it was factory design pattern, then it was suitable for a single location. But for this pattern, we need multiple locations and some critical design changes.


