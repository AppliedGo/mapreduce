/*
<!--
Copyright (c) 2016 Christoph Berger. Some rights reserved.
Use of this text is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

<!--
NOTE: The comments in this file are NOT godoc compliant. This is not an oversight.

Comments and code in this file are used for describing and explaining a particular topic to the reader. While this file is a syntactically valid Go source file, its main purpose is to get converted into a blog article. The comments were created for learning and not for code documentation.
-->

+++
title = "MapReduce - munching through Big Data"
description = "The essence of the MapReduce algorithm, explained in Go"
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2016-11-10"
publishdate = "2016-11-10"
draft = "true"
domains = ["Big Data"]
tags = ["mapreduce", "parallelization", "data processing", "performance"]
categories = ["Tutorial"]
+++

How Google tackled the problem of processing enormous amounts of data, and how you can do the same with Go.

<!--more-->

## Map and Reduce

This is going to be a boring article about two boring functions, `map()` and `reduce()`. Here is the story:

You have a list with elements of type, say, string.

```go
var list []string
```

You have a function that takes an int and produces a string.

```go
func process(n int) (s string) {
	return len(s)
}
```

Now you define a function called `map()` that takes this function and applies it to each of the elements in the list and returns a list of all results.

```go
func map(list []string, fn func(string)int) []int {
	res := make([]string, len(list))
	for i, elem := range list {
		res[i] = fn(elem)
	}
	return res
}
```

Finally, you define another function `reduce()` that takes the result list and boils it down to a single result..

```go
func reduce(list []int) (res int) {
	for _, elem := range list {
		res += elem
	}
	return res
}
```

HYPE[Map and Reduce](mapandreduce.html)

That's it. End of the story. Pretty boring, eh?


## But wait! ...

... what's this?

**Looks like we just abstracted away the concept of `for` loops!**

What does this buy us?

First, no more one-off index errors.

Second, and more importantly, if the mapped function `fn` does not depend on previous results, it can be trivially called in a concurrent manner.

How to do this? Simple: Split the list into *n* pieces and pass them to *n* independently running mappers. Next, have the mappers run on separate CPU cores, or even on separate CPU's.

Imagine the speed boost you'll get. Map and reduce, as it seems, form a fundamental concept for distributed loops.

> Lemme repeat that. By abstracting away the very concept of looping, you can implement looping any way you want, including implementing it in a way that scales nicely with extra hardware.
>
> Joel Spolsky, [Can Your Programming Language Do This?](http://www.joelonsoftware.com/items/2006/08/01.html) (2006)


## From map and reduce to MapReduce

Google researchers were reportedly the first who took the map/reduce concept and scaled it up to "Web search engine level" (I leave the exact definition of "Web search engine level" as an exercise for the reader). MapReduce was born.

Here is how it works.







## The code
*/

// ## Imports and globals
package main
