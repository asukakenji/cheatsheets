# Permutation and Probability

`n` people join a party, and exchange presents.
Each of them brings a unique present to the party.
Each present is given a label (from `1` to `n`).
For each present, a corresponding label is written on a piece of paper and put in a box.
Then, each of the people draws randomly exactly one piece of paper from the box.
The person gets the present with the label written on the paper.

What is the chance of at least one people getting the present he/she brings?

| `n` | Permutations | Chance |
| --- | --- | --- |
| `1` | <code><strong>1</strong></code>\* | `1/1` = `1.0` |
| `2` | <code><strong>12</strong></code>\*<br/><code>21</code> | `1/2` = `0.5` |
| `3` | <code><strong>123</strong></code>\*<br/><code><strong>1</strong>32</code>\*<br/><code>21<strong>3</strong></code>\*<br/><code>231</code><br/><code>312</code><br/><code>3<strong>2</strong>1</code>\* | `4/6` = `0.666...` |
| `4` | <code><strong>1234</strong></code>\*<br/><code><strong>12</strong>43</code>\*<br/><code><strong>1</strong>32<strong>4</strong></code>\*<br/><code><strong>1</strong>342</code>\*<br/><code><strong>1</strong>423</code>\*<br/><code><strong>1</strong>4<strong>3</strong>2</code>\*<br/><code>21<strong>34</strong></code>\*<br/><code>2143</code><br/><code>231<strong>4</strong></code>\*<br/><code>2341</code><br/><code>2413</code><br/><code>24<strong>3</strong>1</code>\*<br/><code>312<strong>4</strong></code>\*<br/><code>3142</code><br/><code>3<strong>2</strong>1<strong>4</strong></code>\*<br/><code>3<strong>2</strong>41</code>\*<br/><code>3412</code><br/><code>3421</code><br/><code>4123</code><br/><code>41<strong>3</strong>2</code>\*<br/><code>4<strong>2</strong>13</code>\*<br/><code>4<strong>23</strong>1</code>\*<br/><code>4312</code><br/><code>4321</code> | `15/24` = `0.625` |
| `5` | (skipped) | `76/120` = `0.6333...` |
| `6` | (skipped) | `455/720` = `0.6319444...` |
| `7` | (skipped) | `3186/5040` = `0.632 142857 142857 142857 ...` |
| `8` | (skipped) | `25487/40320` = `0.6321180555...` |
| `9` | (skipped) | `229384/362880` = `0.6321 208112874779541446 208112874779541446 208112874779541446 ...` |
| `10` | (skipped) | `2293839/3628800` = `0.6321205357 142857 142857 142857 ...` |
| `11` | (skipped) | `25232230/39916800` = `0.6321205 607663940997274330 607663940997274330 607663940997274330 ...` |
| `12` | (skipped) | `302786759/479001600` = `0.6321205586 787184009406231628453850676072898295120517342739564961 ...` |
| ... | ... | ... |
| `n` | (skipped) | `1 - (1 - 1/n) ** n` |
| ... | ... | ... |
| `∞` | (skipped) | `1 - 1/e` |
