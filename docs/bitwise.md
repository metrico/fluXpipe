## `bitwise` package

Package bitwise provides functions for performing bitwise operations on integers.

Import the `bitwise` package:

```flux
import "bitwise"
```

### Functions

### `sand()`

sand performs the bitwise operation, `a AND b`, with integers.

#### Function type signature

```flux
(a: int, b: int) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| a | a: Left hand operand. | Yes |
| b | b: Right hand operand. | Yes |
### `sclear()`

sclear performs the bitwise operation `a AND NOT b`.
Both `a` and `b` are integers.

#### Function type signature

```flux
(a: int, b: int) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| a | a: Left hand operand. | Yes |
| b | b: Bits to clear. | Yes |
### `slshift()`

slshift shifts the bits in `a` left by `b` bits.
Both `a` and `b` are integers.

#### Function type signature

```flux
(a: int, b: int) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| a | a: Left hand operand. | Yes |
| b | b: Number of bits to shift. | Yes |
### `snot()`

snot inverts every bit in `a`, an integer.

#### Function type signature

```flux
(a: int) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| a | a: Integer to invert. | Yes |
### `sor()`

sor performs the bitwise operation, `a OR b`, with integers.

#### Function type signature

```flux
(a: int, b: int) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| a | a: Left hand operand. | Yes |
| b | b: Right hand operand. | Yes |
### `srshift()`

srshift shifts the bits in `a` right by `b` bits.
Both `a` and `b` are integers.

#### Function type signature

```flux
(a: int, b: int) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| a | a: Left hand operand. | Yes |
| b | b: Number of bits to shift. | Yes |
### `sxor()`

sxor performs the bitwise operation, `a XOR b`, with integers.

#### Function type signature

```flux
(a: int, b: int) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| a | a: Left hand operand. | Yes |
| b | b: Right hand operand. | Yes |
### `uand()`

uand performs the bitwise operation, `a AND b`, with unsigned integers.

#### Function type signature

```flux
(a: uint, b: uint) => uint
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| a | a: Left hand operand. | Yes |
| b | b: Right hand operand. | Yes |
### `uclear()`

uclear performs the bitwise operation `a AND NOT b`, with unsigned integers.

#### Function type signature

```flux
(a: uint, b: uint) => uint
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| a | a: Left hand operand. | Yes |
| b | b: Bits to clear. | Yes |
### `ulshift()`

ulshift shifts the bits in `a` left by `b` bits.
Both `a` and `b` are unsigned integers.

#### Function type signature

```flux
(a: uint, b: uint) => uint
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| a | a: Left hand operand. | Yes |
| b | b: Number of bits to shift. | Yes |
### `unot()`

unot inverts every bit in `a`, an unsigned integer.

#### Function type signature

```flux
(a: uint) => uint
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| a | a: Unsigned integer to invert. | Yes |
### `uor()`

uor performs the bitwise operation, `a OR b`, with unsigned integers.

#### Function type signature

```flux
(a: uint, b: uint) => uint
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| a | a: Left hand operand. | Yes |
| b | b: Right hand operand. | Yes |
### `urshift()`

urshift shifts the bits in `a` right by `b` bits.
Both `a` and `b` are unsigned integers.

#### Function type signature

```flux
(a: uint, b: uint) => uint
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| a | a: Left hand operand. | Yes |
| b | b: Number of bits to shift. | Yes |
### `uxor()`

uxor performs the bitwise operation, `a XOR b`, with unsigned integers.

#### Function type signature

```flux
(a: uint, b: uint) => uint
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| a | a: Left hand operand. | Yes |
| b | b: Right hand operand. | Yes |
