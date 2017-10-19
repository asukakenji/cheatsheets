# Tensorflow Cheatsheet

## Sessions and Constants

### Version 1

```python
import tensorflow as tf

with tf.Session() as sess:
    x = tf.constant([1, 2])
    y = tf.constant([3, 4])
    z = x + y
    print(sess.run(z))
```

### Version 2

```python
import tensorflow as tf

with tf.Session():
    x = tf.constant([1, 2])
    y = tf.constant([3, 4])
    z = x + y
    print(z.eval())
```

### Version 3

```python
import tensorflow as tf

sess = tf.Session()
x = tf.constant([1, 2])
y = tf.constant([3, 4])
z = x + y
print(sess.run(z))
sess.close()
```
