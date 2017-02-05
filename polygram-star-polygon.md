# Calculating the exact coordinates of the vertices of a polygram (star polygon)

Let `p` and `q` be variables in the Schläfli symbol `{p/q}` representing a polygram.

## The Vertices

Assume that the `0`-th vertex is located at `(r⋅cos(0), r⋅sin(0))`.

Then, the `k`-th vertex, where `0 <= k < p`, is located at:

```
(r⋅cos(2π⋅k/p), r⋅sin(2π⋅k/p))
```

In the trace, a line is drawn by joining the `k`-th vertex and the `(k+q)`-th vertex, located at:

```
(r⋅cos(2π⋅(k+q)/p), r⋅sin(2π⋅(k+q)/p))
```

In the outline, the above line is intersected by the line joining the `k+1`-th vertex and the `(k-q+1)`-th vertex. The `k+1`-th vertex is located at:

```
(r⋅cos(2π⋅(k+1)/p), r⋅sin(2π⋅(k+1)/p))
```

The `(k-q+1)`-th vertex is located at:

```
(r⋅cos(2π⋅(k-q+1)/p), r⋅sin(2π⋅(k-q+1)/p))
```

## The Lines

The equation of the 1st line (`k`-th to `(k+q)`-th vertex) is:

```
           y - y₁   y₂ - y₁
           ------ = -------
           x - x₁   x₂ - x₁

y - r⋅sin(2π⋅k/p)   r⋅sin(2π⋅(k+q)/p) - r⋅sin(2π⋅k/p)
----------------- = ---------------------------------
x - r⋅cos(2π⋅k/p)   r⋅cos(2π⋅(k+q)/p) - r⋅cos(2π⋅k/p)

                    sin(2π⋅(k+q)/p) - sin(2π⋅k/p)
                  = -----------------------------
                    cos(2π⋅(k+q)/p) - cos(2π⋅k/p)

                     2⋅cos(π⋅(2k+q)/p)⋅sin(π⋅q/p)
                  = -----------------------------
                    -2⋅sin(π⋅(2k+q)/p)⋅sin(π⋅q/p)

                      cos(π⋅(2k+q)/p)
                  = - ---------------
                      sin(π⋅(2k+q)/p)

                      cos(π⋅(2k+q)/p)
y - r⋅sin(2π⋅k/p) = - --------------- ⋅ (x - r⋅cos(2π⋅k/p))
                      sin(π⋅(2k+q)/p)

                       cos(π⋅(2k+q)/p)        r⋅cos(2π⋅k/p))⋅cos(π⋅(2k+q)/p)
                  = (- ---------------) ⋅ x + ------------------------------
                       sin(π⋅(2k+q)/p)               sin(π⋅(2k+q)/p)

                       cos(π⋅(2k+q)/p)        r⋅cos(2π⋅k/p))⋅cos(π⋅(2k+q)/p)
                y = (- ---------------) ⋅ x + ------------------------------ + r⋅sin(2π⋅k/p)
                       sin(π⋅(2k+q)/p)               sin(π⋅(2k+q)/p)

                       cos(π⋅(2k+q)/p)        r⋅cos(2π⋅k/p))⋅cos(π⋅(2k+q)/p) + r⋅sin(2π⋅k/p)⋅sin(π⋅(2k+q)/p)
                  = (- ---------------) ⋅ x + --------------------------------------------------------------
                       sin(π⋅(2k+q)/p)                               sin(π⋅(2k+q)/p)

                       cos(π⋅(2k+q)/p)        r⋅cos(π⋅(2k+q)/p - 2π⋅k/p))
                  = (- ---------------) ⋅ x + ---------------------------
                       sin(π⋅(2k+q)/p)              sin(π⋅(2k+q)/p)

                       cos(π⋅(2k+q)/p)         r⋅cos(π⋅q/p))
                  = (- ---------------) ⋅ x + ---------------
                       sin(π⋅(2k+q)/p)        sin(π⋅(2k+q)/p)
```

The equation of the 2nd line (`k+1`-th to `(k-q+1)`-th vertex) is:

```
               y - y₁   y₂ - y₁
               ------ = -------
               x - x₁   x₂ - x₁

y - r⋅sin(2π⋅(k+1)/p)   r⋅sin(2π⋅(k-q+1)/p) - r⋅sin(2π⋅(k+1)/p)
--------------------- = ---------------------------------------
x - r⋅cos(2π⋅(k+1)/p)   r⋅cos(2π⋅(k-q+1)/p) - r⋅cos(2π⋅(k+1)/p)

                        sin(2π⋅(k-q+1)/p) - sin(2π⋅(k+1)/p)
                      = -----------------------------------
                        cos(2π⋅(k-q+1)/p) - cos(2π⋅(k+1)/p)

                         2⋅cos(π⋅(2k-q+2)/p)⋅sin(π⋅q/p)
                      = -------------------------------
                        -2⋅sin(π⋅(2k-q+2)/p)⋅sin(π⋅q/p)

                          cos(π⋅(2k-q+2)/p)
                      = - -----------------
                          sin(π⋅(2k-q+2)/p)

                          cos(π⋅(2k-q+2)/p)
y - r⋅sin(2π⋅(k+1)/p) = - ----------------- ⋅ (x - r⋅cos(2π⋅(k+1)/p))
                          sin(π⋅(2k-q+2)/p)

                           cos(π⋅(2k-q+2)/p)        r⋅cos(2π⋅(k+1)/p))⋅cos(π⋅(2k-q+2)/p)
                      = (- -----------------) ⋅ x + ------------------------------------
                           sin(π⋅(2k-q+2)/p)                 sin(π⋅(2k-q+2)/p)

                           cos(π⋅(2k-q+2)/p)        r⋅cos(2π⋅(k+1)/p))⋅cos(π⋅(2k-q+2)/p)
                    y = (- -----------------) ⋅ x + ------------------------------------ + r⋅sin(2π⋅(k+1)/p)
                           sin(π⋅(2k-q+2)/p)                 sin(π⋅(2k-q+2)/p)

                           cos(π⋅(2k-q+2)/p)        r⋅cos(2π⋅(k+1)/p))⋅cos(π⋅(2k-q+2)/p) + r⋅sin(2π⋅(k+1)/p)⋅sin(π⋅(2k-q+2)/p)
                      = (- -----------------) ⋅ x + --------------------------------------------------------------------------
                           sin(π⋅(2k-q+2)/p)                                    sin(π⋅(2k-q+2)/p)

                           cos(π⋅(2k-q+2)/p)        r⋅cos(π⋅(2k-q+2)/p - 2π⋅(k+1)/p))
                      = (- -----------------) ⋅ x + ---------------------------------
                           sin(π⋅(2k-q+2)/p)                sin(π⋅(2k-q+2)/p)

                           cos(π⋅(2k-q+2)/p)          r⋅cos(π⋅q/p))
                      = (- -----------------) ⋅ x + -----------------
                           sin(π⋅(2k-q+2)/p)        sin(π⋅(2k-q+2)/p)
```
