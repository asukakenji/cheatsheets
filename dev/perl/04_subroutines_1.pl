use strict;
use warnings;

# Subroutine Forward Declaration
sub print_first_three_arguments;

# This prints: Amy Bob Cody
print_first_three_arguments("Amy", "Bob", "Cody", "Dan");

# Parentheses are optional
print_first_three_arguments "Amy", "Bob", "Cody", "Dan";

sub print_first_three_arguments {
    my $first = shift;
    my $second = shift;
    my $third = shift;
    print($first, " ", $second, " ", $third, "\n");
}

# Subroutine Reference (Function Pointer)
my $p3ptr = \&print_first_three_arguments;
$p3ptr->("Amy", "Bob", "Cody", "Dan");

# Old syntax, same result
&$p3ptr("Amy", "Bob", "Cody", "Dan");

# Anonymous Subroutine
my $p3anon = sub {
    my $first = shift;
    my $second = shift;
    my $third = shift;
    print($first, " ", $second, " ", $third, "\n");
};
$p3anon->("Amy", "Bob", "Cody", "Dan");

# Old syntax, same result
&$p3anon("Amy", "Bob", "Cody", "Dan");
