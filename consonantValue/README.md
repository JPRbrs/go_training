Given a lowercase string that has alphabetic characters only and no spaces, return the
highest value of consonant substrings. Consonants are any letters of the alpahabet except
"aeiou".

We shall assign the following values: a = 1, b = 2, c = 3, .... z = 26.

For example, for the word "zodiacs", let's remove the vowels. We get: "z d cs"

-- The consonant substrings are: "z", "d" and "cs"
values are
 z = 26, d = 4
 cs = 3 + 19 = 22.

The highest is 26.=> solve("zodiacs") = 26

For the word "strength"
consonant substrings are:
 "str" = 19 + 20 + 18 = 57
 "ngth" = 14 + 7 + 20 + 8 = 49.

The highest is 57 => solve("strength") = 57
