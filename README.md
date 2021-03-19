# Daily Coding Problem: Problem #558 [Medium] 

## Problem statement

This problem was asked by Google.

The area of a circle is defined as pi\*r<sup>2</sup>.
Estimate pi to 3 decimal places using a Monte Carlo method.

Hint: The basic equation of a circle is x<sup>2</sup>+ y<sup>2</sup>= r<sup>2</sup>

## Program design

I'm going to use Go, a perfectly fine general purpose programming language.


The standard library function `rand.Float64()` gives back a pseudo-random
number in the range [0.0,1.0).

A quarter of a 1.0-unit radius circle fits inside a quarter of a 2x2 square.

By calling `rand.Float64()` twice, once for an X value, once for a Y value,
the program gets a point in the 1.0 x 1.0 quarter of a 2x2 square.
The square of the distance from the origin is X*X+Y*Y.
If that square of the distance is less than 1.0, the point is
inside the quarter of a circle,
If the square of the distance is greater than 1.0, the point is outside
the quarter of a circle.

If we assume that any such pseudo-randomly chosen point is uniformly
likely to be anywhere in the 1.0 x 1.0 quarter of a square,
and we choose enough random points, pi will be 4.0 * (inside/total).

I wrote a [program](pi1.go) to choose a large number of points
pseudo-randomly, calculate the square of the distance,
and count all the points that lie inside the quarter of a unit circle.

You just need a loop counter, a count of points inside the circle,
seed the random number generator, and make a simple floating point
calculation.

For 100,000,000 pseudo-randomly-chosen points:

```
$ time ./pi1
pi = 3.141572
./pi1  4.13s user 0.01s system 99% cpu 4.147 total
```

## Another method

A method from 1733,
(Buffon's needle)[https://arxiv.org/pdf/2103.09347.pdf],
can also be used.
You can compute e and pi using this method.

## Analysis

This isn't a super hard question,
if you've seen [Physics Girl](https://qa.pbs.org/video/physics-girl-pi-darts/)
try this in real life with darts.

I'm not sure this is a good interview question to find out if
a candidate can write a program.
The actual program I wrote ended up as 27 lines of leisurely-formatted Go.
It has a loop and an if-statement.
It makes 2 calls to a library function,
and a simple floating point calculation.

It may be a good question if the candidate hasn't seen
Physics Girl, or has not seen any of the large number of
other demonstrations of this that exist in the world.
It requires a grasp of pseudo-random numbers and plane geometry,
and maybe a small flash of insight to marry the two.

It's possible that interviewers want candidates to talk about
how this method converges on pi,
and candidates should discuss that.

Or maybe, since Google asked this, they just want to see
if the candidate knows about Buffon's needle.

## Around the web

[Another version](https://ahmadhamze.github.io/estimating-pie.html) in python.
