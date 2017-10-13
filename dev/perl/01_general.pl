# Tell the compiler to generate errors if variables are not declared
use strict;

# Tell the compiler to show warnings
use warnings;

# my creates block scoped variables
my $value = 0;

# Print the Perl version
# It prints something like 5.018002, which means 5.18.2
print("$]\n");

# Print the operating sysyem
# It prints something like darwin, linux, MSWin32, etc.
print("$^O\n");
