use strict;
use warnings;

my $scalar1 = 8;
my $scalar2 = 42;
my $scalar3 = "8";

# Numeric Comparison:
# ==, !=, <, >, <=, >=, <=>
my $result1 = $scalar1 == $scalar2;   # undef
my $result2 = $scalar1 < $scalar2;    # 1
my $result3 = $scalar1 <=> $scalar2;  # -1

# String Comparison:
# eq, ne, lt, gt, le, ge, cmp
my $result4 = $scalar1 eq $scalar3;   # 1
my $result5 = $scalar1 lt $scalar2;   # undef
my $result6 = $scalar1 cmp $scalar2;  # 1

# Arithmetic Operators
# $scalar3 behaves like a number
print($scalar2 + $scalar3, "\n");   #       Addition: 50
print($scalar2 - $scalar3, "\n");   #    Subtraction: 34
print($scalar2 * $scalar3, "\n");   # Multiplication: 336
print($scalar2 / $scalar3, "\n");   #       Division: 5.25
print($scalar2 % $scalar3, "\n");   #      Remainder: 2
print($scalar2 ** $scalar3, "\n");  # Exponentiation: 9682651996416
print("\n");
print($scalar2++, "\n");            # Post-Increment: 42
print($scalar2, "\n");              #                 43
print(++$scalar2, "\n");            #  Pre-Increment: 44
print("\n");

# String Operators
# $scalar1 behaves like a string
print($scalar1 . $scalar3, "\n");  # Concatenation: 88
print($scalar1 x $scalar3, "\n");  #    Repetition: 88888888
print("\n");

# Single Quotes and Double Quotes
print('$scalar1, $scalar2', "\n");  # $scalar1, $scalar2
print("$scalar1, $scalar2\n");      # 8, 44
