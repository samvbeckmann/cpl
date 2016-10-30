local $/;
$infile = <>;
$infile =~ s/\/\*.*?\*\/|\/\/.*?\n//gs;
print($infile);
