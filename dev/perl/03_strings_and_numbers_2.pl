use strict;
use warnings;

my $scalar1 = 8;
my $scalar4 = "foo";
my $scalar5 = "bar";

# $scalar4 behaves like zero
print($scalar4 + $scalar1, "\n");    # 8, Warning: Argument "foo" isn't numeric in addition (+)
print($scalar4 - $scalar1, "\n");    # -8
print($scalar4 * $scalar1, "\n");    # 0
print($scalar4 / $scalar1, "\n");    # 0
print($scalar4 % $scalar1, "\n");    # 0
print($scalar4 ** $scalar1, "\n");   # 0
print("\n");
print($scalar1 + $scalar4, "\n");    # 8
print($scalar1 - $scalar4, "\n");    # 8
print($scalar1 * $scalar4, "\n");    # 0
# print($scalar1 / $scalar4, "\n");  # Error: Illegal division by zero
# print($scalar1 % $scalar4, "\n");  # Error: Illegal modulus zero
print($scalar1 ** $scalar4, "\n");   # 1
print("\n");

# String Increment
print($scalar5++, "\n");  # Post-Increment: bar
print($scalar5, "\n");    #                 bas
print(++$scalar5, "\n");  #  Pre-Increment: bat
print("\n");

# String Increment (Weird Behaviour)
# $scalar4 becomes a number after the warning above
# Tested on:
# - Perl v5.18.2 on darwin
# - Perl v5.24.2 on linux
print($scalar4++, "\n");  # Post-Increment: foo
print($scalar4, "\n");    #                 1
print(++$scalar4, "\n");  #  Pre-Increment: 2
print("\n");
