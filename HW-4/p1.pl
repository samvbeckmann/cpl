sub union {
  my %hash = map { $_ => 1} (@{$_[0]}, @{$_[1]}); # Put arrays into hash
  my @keys = keys %hash; # Get keyset from hash
  return wantarray ? @keys : join(",", @keys); # Return result
}
