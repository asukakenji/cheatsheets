# Permutation and Probability

`n` people join a party, and exchange presents.
Each of them brings a unique present to the party.
Each present is given a label (from `1` to `n`).
For each present, a corresponding label is written on a piece of paper and put in a box.
Then, each of the people draws randomly exactly one piece of paper from the box.
The person gets the present with the label written on the paper.
What is the chance of at least one people getting the present he/she brings?

| `n` | Permutations | Percentage |
| --- | --- | --- |
| `1` | <code><strong>1</strong></code>\* | `1/1` = `1.0` |
| `2` | <code><strong>12</strong></code>\*<br/><code>21</code> | `1/2` = `0.5` |
| `3` | <code><strong>123</strong></code>\*<br/><code><strong>1</strong>32</code>\*<br/><code>21<strong>3</strong></code>\*<br/><code>231</code><br/><code>312</code><br/><code>3<strong>2</strong>1</code>\* | `4/6` = `0.666...` |
| `4` | <code><strong>1234</strong></code>\*<br/><code><strong>12</strong>43</code>\*<br/><code><strong>1</strong>32<strong>4</strong></code>\*<br/><code><strong>1</strong>342</code>\*<br/><code><strong>1</strong>423</code>\*<br/><code><strong>1</strong>4<strong>3</strong>2</code>\*<br/><code>21<strong>34</strong></code>\*<br/><code>2143</code><br/><code>231<strong>4</strong></code>\*<br/><code>2341</code><br/><code>2413</code><br/><code>24<strong>3</strong>1</code>\*<br/><code>312<strong>4</strong></code>\*<br/><code>3142</code><br/><code>3<strong>2</strong>1<strong>4</strong></code>\*<br/><code>3<strong>2</strong>41</code>\*<br/><code>3412</code><br/><code>3421</code><br/><code>4123</code><br/><code>41<strong>3</strong>2</code>\*<br/><code>4<strong>2</strong>13</code>\*<br/><code>4<strong>23</strong>1</code>\*<br/><code>4312</code><br/><code>4321</code> | `15/24` = `0.625` |
