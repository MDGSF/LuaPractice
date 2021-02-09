test = "int x; /* x */ int y; /* y */"
print((string.gsub(test, "/%*.*%*/", ""))) -- int x;

test = "int x; /* x */ int y; /* y */"
print((string.gsub(test, "/%*.-%*/", ""))) -- int x;  int y;
