# Daily Coding Problem: Problem #558 [Medium] 

## Problem statement

This problem was asked by Google.

The area of a circle is defined as pi\*r<sup>2</sup>.
Estimate pi to 3 decimal places using a Monte Carlo method.

Hint: The basic equation of a circle is x<sup>2</sup>+ y<sup>2</sup>= r<sup>2</sup>

## Analysis

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

For 100,000,000 pseudo-randomly-chosen points:

```
$ time ./pi1
pi = 3.141572
./pi1  4.13s user 0.01s system 99% cpu 4.147 total
```
