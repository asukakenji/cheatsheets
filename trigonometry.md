# Deriving trigonometry formulae

Start with Euler's Formula:

```
cos(x) + i⋅sin(x) = eⁱˣ
```

Substitute `x = A + B`:

```
cos(A + B) + i⋅sin(A + B) = eⁱ⁽ᴬ⁺ᴮ⁾
                          = eⁱᴬ⋅eⁱᴮ
                          = [cos(A) + i⋅sin(A)]⋅[cos(B) + i⋅sin(B)]
                          = cos(A)cos(B) - sin(A)sin(B) + i⋅[sin(A)cos(B) + cos(A)sin(B)]
```

Therefore...

## Compund Angle Formulae

```
cos(A + B) = cos(A)cos(B) - sin(A)sin(B)
cos(A - B) = cos(A)cos(B) + sin(A)sin(B)
sin(A + B) = sin(A)cos(B) + cos(A)sin(B)
sin(A - B) = sin(A)cos(B) - cos(A)sin(B)
```

From the above:

```
             cos(A - B) = cos(A)cos(B) + sin(A)sin(B)
+            cos(A + B) = cos(A)cos(B) - sin(A)sin(B)
---------------------------------------------------------
cos(A - B) + cos(A + B) = 2⋅cos(A)cos(B)
           cos(A)cos(B) = (1/2)⋅[cos(A - B) + cos(A + B)]


             cos(A - B) = cos(A)cos(B) + sin(A)sin(B)
-            cos(A + B) = cos(A)cos(B) - sin(A)sin(B)
---------------------------------------------------------
cos(A - B) - cos(A + B) = 2⋅sin(A)sin(B)
           sin(A)sin(B) = (1/2)⋅[cos(A - B) - cos(A + B)]

             sin(A + B) = sin(A)cos(B) + cos(A)sin(B)
+            sin(A - B) = sin(A)cos(B) - cos(A)sin(B)
---------------------------------------------------------
sin(A + B) + sin(A - B) = 2⋅sin(A)cos(B)
           sin(A)cos(B) = (1/2)⋅[sin(A + B) + sin(A - B)]


             sin(A + B) = sin(A)cos(B) + cos(A)sin(B)
-            sin(A - B) = sin(A)cos(B) - cos(A)sin(B)
---------------------------------------------------------
sin(A + B) - sin(A - B) = 2⋅cos(A)sin(B)
           cos(A)sin(B) = (1/2)⋅[sin(A + B) - sin(A - B)]
```

Therefore...

## Product-to-Sum Formulae

```
cos(A)cos(B) = (1/2)⋅[cos(A - B) + cos(A + B)]
sin(A)sin(B) = (1/2)⋅[cos(A - B) - cos(A + B)]
cos(A)sin(B) = (1/2)⋅[sin(A + B) - sin(A - B)]
sin(A)cos(B) = (1/2)⋅[sin(A + B) + sin(A - B)]
```

Substitute `C = A + B`, `D = A - B`. Then, `A = (C + D) / 2`, `B = (C - D) / 2`:

```
cos((C + D) / 2)cos((C - D) / 2) = (1/2)⋅[cos(D) + cos(C)]
sin((C + D) / 2)sin((C - D) / 2) = (1/2)⋅[cos(D) - cos(C)]
cos((C + D) / 2)sin((C - D) / 2) = (1/2)⋅[sin(C) - sin(D)]
sin((C + D) / 2)cos((C - D) / 2) = (1/2)⋅[sin(C) + sin(D)]

cos(D) + cos(C) = 2⋅cos((C + D) / 2)⋅cos((C - D) / 2)
cos(D) - cos(C) = 2⋅sin((C + D) / 2)⋅sin((C - D) / 2)
sin(C) - sin(D) = 2⋅cos((C + D) / 2)⋅sin((C - D) / 2)
sin(C) + sin(D) = 2⋅sin((C + D) / 2)⋅cos((C - D) / 2)

cos(C) + cos(D) = 2⋅cos((C + D) / 2)⋅cos((D - C) / 2)
cos(C) - cos(D) = 2⋅sin((C + D) / 2)⋅sin((D - C) / 2)
sin(C) + sin(D) = 2⋅sin((C + D) / 2)⋅cos((C - D) / 2)
sin(C) - sin(D) = 2⋅cos((C + D) / 2)⋅sin((C - D) / 2)

cos(C) + cos(D) =  2⋅cos((C + D) / 2)⋅cos((C - D) / 2)
cos(C) - cos(D) = -2⋅sin((C + D) / 2)⋅sin((C - D) / 2)
sin(C) + sin(D) =  2⋅sin((C + D) / 2)⋅cos((C - D) / 2)
sin(C) - sin(D) =  2⋅cos((C + D) / 2)⋅sin((C - D) / 2)
```

Therefore...

## Sum-to-Product Formulae

```
cos(A) + cos(B) =  2⋅cos((A + B) / 2)⋅cos((A - B) / 2)
cos(A) - cos(B) = -2⋅sin((A + B) / 2)⋅sin((A - B) / 2)
sin(A) + sin(B) =  2⋅sin((A + B) / 2)⋅cos((A - B) / 2)
sin(A) - sin(B) =  2⋅cos((A + B) / 2)⋅sin((A - B) / 2)
```
