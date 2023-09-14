## `regexp` package

Package regexp provides tools for working with regular expressions.

Import the `regexp` package:

```flux
import "regexp"
```

### Functions

### `compile()`

compile parses a string into a regular expression and returns a regexp type
that can be used to match against strings.

#### Function type signature

```flux
(v: string) => regexp
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to parse into a regular expression. | Yes |
### `findString()`

findString returns the left-most regular expression match in a string.

#### Function type signature

```flux
(r: regexp, v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| r | r: Regular expression used to search `v`. | Yes |
| v | v: String value to search. | Yes |
### `findStringIndex()`

findStringIndex returns a two-element array of integers that represent the
beginning and ending indexes of the first regular expression match in a string.

#### Function type signature

```flux
(r: regexp, v: string) => [int]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| r | r: Regular expression used to search `v`. | Yes |
| v | v: String value to search. | Yes |
### `getString()`

getString returns the source string used to compile a regular expression.

#### Function type signature

```flux
(r: regexp) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| r | r: Regular expression object to convert to a string. | Yes |
### `matchRegexpString()`

matchRegexpString tests if a string contains any match to a regular expression.

#### Function type signature

```flux
(r: regexp, v: string) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| r | r: Regular expression used to search `v`. | Yes |
| v | v: String value to search. | Yes |
### `quoteMeta()`

quoteMeta escapes all regular expression metacharacters in a string.

#### Function type signature

```flux
(v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String that contains regular expression metacharacters to escape. | Yes |
### `replaceAllString()`

replaceAllString replaces all reguar expression matches in a string with a
specified replacement.

#### Function type signature

```flux
(r: regexp, t: string, v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| r | r: Regular expression used to search `v`. | Yes |
| v | v: String value to search. | Yes |
| t | t: Replacement for matches to `r`. | Yes |
### `splitRegexp()`

splitRegexp splits a string into substrings separated by regular expression
matches and returns an array of `i` substrings between matches.

#### Function type signature

```flux
(i: int, r: regexp, v: string) => [string]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| r | r: Regular expression used to search `v`. | Yes |
| v | v: String value to be searched. | Yes |
| i | i: Maximum number of substrings to return. | Yes |
