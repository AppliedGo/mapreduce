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

Originally, `map` and `reduce` are two functions often seen in functional programming languages. The first applies a function to a list of objects, producing a list of results. The second reduces this list to a single output.

This does not sound too fancy, except that these two functions do nothing less than making `for` loops obsolete. Furthermore, if the mapped function calls do not interact with each other, they can be trivially executed concurrently. That is, split up your input list into *n* pieces and pass them to *n* independently running mappers. On hardware that supports parallelization, the `map` operation suddenly gets multiple times faster! And harware parallelization can happen at different scales (think "multi-core CPU", "multi-CPU node", "multi-node rack", "multi-rack datacenter").

> Lemme repeat that. By abstracting away the very concept of looping, you can implement looping any way you want, including implementing it in a way that scales nicely with extra hardware.
>
> Joel Spolsky, [Can Your Programming Language Do This?](http://www.joelonsoftware.com/items/2006/08/01.html) (2006)


## The next step: MapReduce




## The code
*/

// ## Imports and globals
package main
