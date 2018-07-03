import sys
import decimal

decimal.getcontext().prec = 1010

total = 0
count = 0

def permutation_impl(f, d, s):
    if d == 0:
        f(s)
    else:
        for i in xrange(d, -1, -1):
            s[i], s[d] = s[d], s[i]
            permutation_impl(f, d-1, s)
            s[i], s[d] = s[d], s[i]

def permutation(f, s):
    if len(s) == 0:
        f(s)
    else:
        permutation_impl(f, len(s)-1, s)

def is_clashed(s):
    for i, v in enumerate(s):
        if i == v:
            return True
    return False

def callback(s):
    global total, count
    total += 1
    if is_clashed(s):
        count += 1

n = int(sys.argv[1])
permutation(callback, range(n))
print "Result: %d/%d = %s" % (count, total, decimal.Decimal(count)/total)
