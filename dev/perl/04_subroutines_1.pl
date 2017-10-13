use strict;
use warnings;

sub print_first_three_arguments {
    my $first = shift;
    my $second = shift;
    my $third = shift;
    print($first, " ", $second, " ", $third, "\n");
}

# This prints: Amy Bob Cody
print_first_three_arguments "Amy", "Bob", "Cody", "Dan";
