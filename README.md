String utilities for Golang. Nothing major. No rocket science.

This package has no external dependencies beyond the standard library.

Why this package exists:

- A reluctance to use regular expressions in XML processing.
- XML allows both single and double quotes:
  ```
  "A quoted string."
  'A quoted string.'
  <mytag myatt1="ain't trip'd up" myatt2='by "weird" quotes'>
  ```
