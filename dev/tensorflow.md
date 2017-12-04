# Tensorflow Cheatsheet

## Sessions and Constants

References:
- https://www.tensorflow.org/api_docs/python/tf/Session
- https://www.tensorflow.org/api_docs/python/tf/Session#run
- https://www.tensorflow.org/api_docs/python/tf/constant

### Version 1

```python
import tensorflow as tf

with tf.Session() as sess:
    x = tf.constant([1, 2])
    y = tf.constant([3, 4])
    z = x + y
    print(sess.run(z))  # [4 6]
```

### Version 2

```python
import tensorflow as tf

with tf.Session():
    x = tf.constant([1, 2])
    y = tf.constant([3, 4])
    z = x + y
    print(z.eval())     # [4 6]
```

### Version 3

```python
import tensorflow as tf

sess = tf.Session()
x = tf.constant([1, 2])
y = tf.constant([3, 4])
z = x + y
print(sess.run(z))      # [4 6]
sess.close()
```

## Matrices

### Creating Matrices

```python
import tensorflow as tf

with tf.Session() as sess:
    x0 = tf.constant([[1, 1 ,1], [1, 1, 1]])
    x1 = tf.constant(1.0, shape=[2, 3])
    x2 = tf.constant([[1., 1., 1.], [1., 1., 1.]])
    x3 = tf.constant([[1, 1 ,1], [1, 1, 1]], dtype=tf.float32)
    x4 = tf.constant([1, 1, 1, 1, 1, 1], dtype=tf.float32, shape=(2, 3))

    #print(sess.run(tf.reduce_all(tf.equal(x0, x1))))   # TypeError
    print(sess.run(tf.reduce_all(tf.equal(x1, x2))))    # True
    print(sess.run(tf.reduce_all(tf.equal(x1, x3))))    # True
    print(sess.run(tf.reduce_all(tf.equal(x1, x4))))    # True
```

In the official documentation, the `shape` parameter is an array.
However, a tuple also works, which may be a compatibility consideration with NumPy.

References:
- https://www.tensorflow.org/api_docs/python/tf/constant
- https://www.tensorflow.org/api_docs/python/tf/equal
- https://www.tensorflow.org/api_docs/python/tf/reduce_all

### Adding and Multiplying Matrices

```python
import tensorflow as tf

with tf.Session() as sess:
    a = tf.constant([[1, 0, 0, 1]])
    b = tf.constant([[1, 2], [3, 4], [5, 6], [7, 8]])
    c = tf.constant(1, shape=[4, 2])
    s1 = tf.add(b, c)
    s2 = b + c
    p = tf.matmul(a, b)
    # p = a @ b             # Only available for Python >= 3.5

    print(sess.run(s1))     # [[2 3] [4 5] [6 7] [8 9]]
    print(sess.run(s2))     # [[2 3] [4 5] [6 7] [8 9]]
    print(sess.run(p))      # [[8 10]]
```

Note that if `a = tf.constant([1, 0, 0, 1])`,
then the shape will become `[4]`, and a `ValueError` will be raised.
The extra pair of brackets turns the shape into `[1, 4]`,
so that it could be multiplied with a `[4, 2]` matrix.

References:
- https://www.tensorflow.org/api_docs/python/tf/add
- https://www.tensorflow.org/api_docs/python/tf/matmul

## Variables

```python
import tensorflow as tf

with tf.Session() as sess:
    increment = tf.constant([1, 2])
    total = tf.Variable([0, 0])

    update_total = tf.assign(total, total + increment)

    sess.run(tf.global_variables_initializer())
    #tf.global_variables_initializer().run()    # Same meaning
    for _ in range(5):
        v = sess.run(update_total)
        print(v)
```

Note that `constant` starts with a small letter,
while `Variable` starts with a capital letter.
In real applications, variables are trained and tuned by the system.

References:
- https://www.tensorflow.org/api_docs/python/tf/Variable

## Placeholders

```python
import tensorflow as tf

x = tf.placeholder(tf.float32, shape=(None, 2), name='x')
w = tf.Variable(initial_value=[[1, 0, 0, 1]], dtype=tf.float32, name='w')
b = tf.Variable(initial_value=[0, 10], dtype=tf.float32, name='b')
y = tf.matmul(w, x) + b

model = tf.global_variables_initializer()

with tf.Session() as sess:
    sess.run(model)
    print(sess.run(y, feed_dict={x: [[1, 2], [3, 4], [5, 6], [7, 8]]}))
```
