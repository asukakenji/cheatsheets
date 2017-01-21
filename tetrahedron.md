# Calculating the exact coordinates of a regular tetrahedron circumscribed by the unit sphere at the origin

First, draw the first point `O` at the origin `(0, 0)`:


Then, draw the second point `A` at `(1, 0)`:


Then, draw the third point `B` so that △BOA is a equilateral triangle.

Since `∠BOA = 60°`, and `OA = OB = 1`:
```
B = (cos 60°, sin 60°)
  = (1/2, √3/2)
```


Let point `C` be the centroid of △BOA.
The coordinates of point `C` could be found be "averaging" the coordinates of `O`, `A`, and `B`.

```
C = ((0 + 1 + 1/2) / 3, (0 + 0 + √3/2) / 3)
  = ((3/2) / 3, (√3/2) / 3)
  = (1/2, 1/(2√3))
```


By extending the 2D space to 3D, here are what we got so far:
```
O = (0,   0,       0)
A = (1,   0,       0)
B = (1/2, √3/2,    0)
C = (1/2, 1/(2√3), 0)
```

Since the forth point `D` is "above" point `C`, let `D = (1/2, 1/(2√3), d)`.
Since `OA = OB = AB = OD = AD = BD = 1`, we take `OD = 1`:
```
OD² = (1/2)² + (1/(2√3))² + d²
  1 = 1/4 + 1/12 + d²
    = 3/12 + 1/12 + d²
    = 4/12 + d²
    = 1/3 + d²
2/3 = d²
  d = ±√(2/3)
```

By picking the positive root for easier calculation, we have `D = (1/2, 1/(2√3), √(2/3))`.


Let point `E` be the center of the tetrahedron.
Since point `E` is on the line `CD`, let `E = (1/2, 1/(2√3), e)`.
Since `OE = AE = BE = DE`, we take `OE = DE`:
```
                      OE = DE
                     OE² = DE²
(1/2)² + (1/(2√3))² + e² = 0² + 0² + (e - √(2/3))²
                1/3 + e² = e² - 2√(2/3) e + (2/3)
                     1/3 = -2√(2/3) e + (2/3)
                    -1/3 = -2√(2/3) e
                     1/3 = 2√(2/3) e
                 1/(2*3) = √(2/3) e
                       e = √3/(2*3*√2)
                         = 1/(2√6)
```

So, here are what we got so far:
```
O = (0,   0,       0)
A = (1,   0,       0)
B = (1/2, √3/2,    0)
D = (1/2, 1/(2√3), √(2/3))
E = (1/2, 1/(2√3), 1/(2√6))
```

Since we need the vertices to be on a unit circle, the distance from the center to a vertex must be 1.
That means, we must scale the points so that `OE = AE = BE = DE = 1`.
Let `k` be the scaling factor:
```
                                 k * OE = 1
                               k² * OE² = 1
k² * [(1/2)² + (1/(2√3))² + (1/(2√6))²] = 1
               k² * (1/4 + 1/12 + 1/24) = 1
              k² * (6/24 + 2/24 + 1/24) = 1
                            k² * (9/24) = 1
                             k² * (3/8) = 1
                                     k² = (8/3)
                                      k = (2√2)/(√3), since k is positive
```

So:
```
O' = (0, 0, 0)
A' = (2√(2/3), 0, 0)
B' = (√(2/3), 2, 0)
D' = (√(2/3), √2/3, 4/3)
E' = (√(2/3), √2/3, 1/3)
```







