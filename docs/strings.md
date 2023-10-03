## `strings` package

Package strings provides functions to operate on UTF-8 encoded strings.

Import the `strings` package:

```flux
import "strings"
```

### Functions

### `compare()`

compare compares the lexicographical order of two strings.

#### Function type signature

```flux
(t: string, v: string) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to compare. | Yes |
| t | t: String value to compare against. | Yes |
### `containsAny()`

containsAny reports whether a specified string contains characters from another string.

#### Function type signature

```flux
(chars: string, v: string) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to search. | Yes |
| chars | chars: Characters to search for. | Yes |
### `containsStr()`

containsStr reports whether a string contains a specified substring.

#### Function type signature

```flux
(substr: string, v: string) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to search. | Yes |
| substr | substr: Substring value to search for. | Yes |
### `countStr()`

countStr counts the number of non-overlapping instances of a substring appears in a string.

#### Function type signature

```flux
(substr: string, v: string) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to search. | Yes |
| substr | substr: Substring to count occurences of. | Yes |
### `equalFold()`

equalFold reports whether two UTF-8 strings are equal under Unicode case-folding.

#### Function type signature

```flux
(t: string, v: string) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to compare. | Yes |
| t | t: String value to compare against. | Yes |
### `hasPrefix()`

hasPrefix indicates if a string begins with a specified prefix.

#### Function type signature

```flux
(prefix: string, v: string) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to search. | Yes |
| prefix | prefix: Prefix to search for. | Yes |
### `hasSuffix()`

hasSuffix indicates if a string ends with a specified suffix.

#### Function type signature

```flux
(suffix: string, v: string) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to search. | Yes |
| suffix | suffix: Suffix to search for. | Yes |
### `index()`

index returns the index of the first instance of a substring in a string.
If the substring is not present, it returns `-1`.

#### Function type signature

```flux
(substr: string, v: string) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to search. | Yes |
| substr | substr: Substring to search for. | Yes |
### `indexAny()`

indexAny returns the index of the first instance of specified characters in a string.
If none of the specified characters are present, it returns `-1`.

#### Function type signature

```flux
(chars: string, v: string) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to search. | Yes |
| chars | chars: Characters to search for. | Yes |
### `isDigit()`

isDigit tests if a single-character string is a digit (0-9).

#### Function type signature

```flux
(v: string) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Single-character string to test. | Yes |
### `isLetter()`

isLetter tests if a single character string is a letter (a-z, A-Z).

#### Function type signature

```flux
(v: string) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Single-character string to test. | Yes |
### `isLower()`

isLower tests if a single-character string is lowercase.

#### Function type signature

```flux
(v: string) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Single-character string value to test. | Yes |
### `isUpper()`

isUpper tests if a single character string is uppercase.

#### Function type signature

```flux
(v: string) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Single-character string value to test. | Yes |
### `joinStr()`

joinStr concatenates elements of a string array into a single string using a specified separator.

#### Function type signature

```flux
(arr: [string], v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| arr | arr: Array of strings to concatenate. | Yes |
| v | v: Separator to use in the concatenated value. | Yes |
### `lastIndex()`

lastIndex returns the index of the last instance of a substring in a string.
If the substring is not present, the function returns -1.

#### Function type signature

```flux
(substr: string, v: string) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to search. | Yes |
| substr | substr: Substring to search for. | Yes |
### `lastIndexAny()`

lastIndexAny returns the index of the last instance of any specified
characters in a string.
If none of the specified characters are present, the function returns `-1`.

#### Function type signature

```flux
(chars: string, v: string) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to search. | Yes |
| chars | chars: Characters to search for. | Yes |
### `repeat()`

repeat returns a string consisting of `i` copies of a specified string.

#### Function type signature

```flux
(i: int, v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to repeat. | Yes |
| i | i: Number of times to repeat `v`. | Yes |
### `replace()`

replace replaces the first `i` non-overlapping instances of a substring with
a specified replacement.

#### Function type signature

```flux
(i: int, t: string, u: string, v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to search. | Yes |
| t | t: Substring value to replace. | Yes |
| u | u: Replacement for `i` instances of `t`. | Yes |
| i | i: Number of non-overlapping `t` matches to replace. | Yes |
### `replaceAll()`

replaceAll replaces all non-overlapping instances of a substring with a specified replacement.

#### Function type signature

```flux
(t: string, u: string, v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to search. | Yes |
| t | t: Substring to replace. | Yes |
| u | u: Replacement for all instances of `t`. | Yes |
### `split()`

split splits a string on a specified separator and returns an array of substrings.

#### Function type signature

```flux
(t: string, v: string) => [string]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to split. | Yes |
| t | t: String value that acts as the separator. | Yes |
### `splitAfter()`

splitAfter splits a string after a specified separator and returns an array of substrings.
Split substrings include the separator, `t`.

#### Function type signature

```flux
(t: string, v: string) => [string]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to split. | Yes |
| t | t: String value that acts as the separator. | Yes |
### `splitAfterN()`

splitAfterN splits a string after a specified separator and returns an array of `i` substrings.
Split substrings include the separator, `t`.

#### Function type signature

```flux
(i: int, t: string, v: string) => [string]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to split. | Yes |
| t | t: String value that acts as the separator. | Yes |
| i | i: Maximum number of split substrings to return. | Yes |
### `splitN()`

splitN splits a string on a specified separator and returns an array of `i` substrings.

#### Function type signature

```flux
(i: int, t: string, v: string) => [string]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to split. | Yes |
| t | t: String value that acts as the separator. | Yes |
| i | i: Maximum number of split substrings to return. | Yes |
### `strlen()`

strlen returns the length of a string. String length is determined by the number of UTF code points a string contains.

#### Function type signature

```flux
(v: string) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to measure. | Yes |
### `substring()`

substring returns a substring based on start and end parameters. These parameters are represent indices of UTF code points in the string.

#### Function type signature

```flux
(end: int, start: int, v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to search for. | Yes |
| start | start: Starting inclusive index of the substring. | Yes |
| end | end: Ending exclusive index of the substring. | Yes |
### `title()`

title converts a string to title case.

#### Function type signature

```flux
(v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to convert. | Yes |
### `toLower()`

toLower converts a string to lowercase.

#### Function type signature

```flux
(v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to convert. | Yes |
### `toTitle()`

toTitle converts all characters in a string to title case.

#### Function type signature

```flux
(v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to convert. | Yes |
### `toUpper()`

toUpper converts a string to uppercase.

#### Function type signature

```flux
(v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String value to convert. | Yes |
### `trim()`

trim removes leading and trailing characters specified in the cutset from a string.

#### Function type signature

```flux
(cutset: string, v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to remove characters from. | Yes |
| cutset | cutset: Leading and trailing characters to remove from the string. | Yes |
### `trimLeft()`

trimLeft removes specified leading characters from a string.

#### Function type signature

```flux
(cutset: string, v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to to remove characters from. | Yes |
| cutset | cutset: Leading characters to trim from the string. | Yes |
### `trimPrefix()`

trimPrefix removes a prefix from a string. Strings that do not start with the prefix are returned unchanged.

#### Function type signature

```flux
(prefix: string, v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to trim. | Yes |
| prefix | prefix: Prefix to remove. | Yes |
### `trimRight()`

trimRight removes trailing characters specified in the cutset from a string.

#### Function type signature

```flux
(cutset: string, v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to to remove characters from. | Yes |
| cutset | cutset: Trailing characters to trim from the string. | Yes |
### `trimSpace()`

trimSpace removes leading and trailing spaces from a string.

#### Function type signature

```flux
(v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to remove spaces from. | Yes |
### `trimSuffix()`

trimSuffix removes a suffix from a string.

#### Function type signature

```flux
(suffix: string, v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to trim. | Yes |
| suffix | suffix: Suffix to remove. | Yes |
