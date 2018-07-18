String utilities for Golang. Nothing major. No rocket science.

This package has no external dependencies beyond the standard library.

Why this package exists:

- A reluctance to use regular expressions in XML processing (they be slow).
- XML allows both single and double quotes (and the code to process it can
get very ugly):
  ```
  "A quoted string."
  'A quoted string.'
  <mytag myatt1="Ain't trip'd up" myatt2='by "weird" quoting'>
  ```

This package takes a liberal view towards trimming leading and trailing
whitespace. Whitespace treatment is one of those very irritating things
about XML. 
