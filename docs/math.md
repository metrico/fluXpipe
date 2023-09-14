## `math` package

Package math provides basic constants and mathematical functions.

Import the `math` package:

```flux
import "math"
```

### Functions

### `NaN()`

NaN returns a IEEE 754 "not-a-number" value.

#### Function type signature

```flux
() => float
```

#### Parameters

No parameters provided.

### `abs()`

abs returns the absolute value of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `acos()`

acos returns the acosine of `x` in radians.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `acosh()`

acosh returns the inverse hyperbolic cosine of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `asin()`

asin returns the arcsine of `x` in radians.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `asinh()`

asinh returns the inverse hyperbolic sine of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `atan()`

atan returns the arctangent of `x` in radians.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `atan2()`

atan2 returns the artangent of `x/y`, using the signs
of the two to determine the quadrant of the return value.

#### Function type signature

```flux
(x: float, y: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| y | y: y-coordinate to use in the operation. | Yes |
| x | x: x-corrdinate to use in the operation. | Yes |
### `atanh()`

atanh returns the inverse hyperbolic tangent of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `cbrt()`

cbrt returns the cube root of x.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `ceil()`

ceil returns the least integer value greater than or equal to `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `copysign()`

copysign returns a value with the magnitude `x` and the sign of `y`.

#### Function type signature

```flux
(x: float, y: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Magnitude to use in the operation. | Yes |
| y | y: Sign to use in the operation. | Yes |
### `cos()`

cos returns the cosine of the radian argument `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `cosh()`

cosh returns the hyperbolic cosine of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `dim()`

dim returns the maximum of `x - y` or `0`.

#### Function type signature

```flux
(x: float, y: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: x-value to use in the operation. | Yes |
| y | y: y-value to use in the operation. | Yes |
### `e()`

e represents the base of the natural logarithm, also known as Euler's number.

#### Function type signature

```flux
float
```

#### Parameters

No parameters provided.

### `erf()`

erf returns the error function of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `erfc()`

erfc returns the complementary error function of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `erfcinv()`

erfcinv returns the inverse of `math.erfc()`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `erfinv()`

erfinv returns the inverse error function of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `exp()`

exp returns `e**x`, the base-e exponential of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `exp2()`

exp2 returns `2**x`, the base-2 exponential of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `expm1()`

expm1 returns `e**x - 1`, the base-e exponential of `x` minus 1.
It is more accurate than `math.exp(x:x) - 1` when `x` is near zero.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `float64bits()`

float64bits returns the IEEE 754 binary representation of `f`,
with the sign bit of `f` and the result in the same bit position.

#### Function type signature

```flux
(f: float) => uint
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| f | f: Value to operate on. | Yes |
### `float64frombits()`

float64frombits returns the floating-point number corresponding to the IEE
754 binary representation `b`, with the sign bit of `b` and the result in the
same bit position.

#### Function type signature

```flux
(b: uint) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| b | b: Value to operate on. | Yes |
### `floor()`

floor returns the greatest integer value less than or equal to `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `frexp()`

frexp breaks `f` into a normalized fraction and an integral part of two.

#### Function type signature

```flux
(f: float) => {frac: float, exp: int}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| f | f: Value to operate on. | Yes |
### `gamma()`

gamma returns the gamma function of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `hypot()`

hypot returns the square root of `p*p + q*q`, taking care to avoid overflow
and underflow.

#### Function type signature

```flux
(p: float, q: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| p | p: p-value to use in the operation. | Yes |
| q | q: q-value to use in the operation. | Yes |
### `ilogb()`

ilogb returns the binary exponent of `x` as an integer.

#### Function type signature

```flux
(x: float) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `isInf()`

isInf reports whether `f` is an infinity, according to `sign`.

#### Function type signature

```flux
(f: float, sign: int) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| f | f: is the value used in the evaluation. | Yes |
| sign | sign: is the sign used in the eveluation. | Yes |
### `isNaN()`

isNaN reports whether `f` is an IEEE 754 "not-a-number" value.

#### Function type signature

```flux
(f: float) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| f | f: Value to operate on. | Yes |
### `j0()`

j0 returns the order-zero Bessel function of the first kind.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `j1()`

j1 is a funciton that returns the order-one Bessel function for the first kind.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `jn()`

jn returns the order-n Bessel funciton of the first kind.

#### Function type signature

```flux
(n: int, x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Order number. | Yes |
| x | x: Value to operate on. | Yes |
### `ldexp()`

ldexp is the inverse of `math.frexp()`. It returns `frac x 2**exp`.

#### Function type signature

```flux
(exp: int, frac: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| frac | frac: Fraction to use in the operation. | Yes |
| exp | exp: Exponent to use in the operation. | Yes |
### `lgamma()`

lgamma returns the natural logarithm and sign (-1 or +1) of `math.gamma(x:x)`.

#### Function type signature

```flux
(x: float) => {sign: int, lgamma: float}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `ln10()`

ln10 represents the natural logarithm of 10.

#### Function type signature

```flux
float
```

#### Parameters

No parameters provided.

### `ln2()`

ln2 represents the natural logarithm of 2.

#### Function type signature

```flux
float
```

#### Parameters

No parameters provided.

### `log()`

log returns the natural logarithm of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `log10()`

log10 returns the decimal logarithm of x.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `log10e()`

log10e represents the base 10 logarithm of **e** (`math.e`).

#### Function type signature

```flux
float
```

#### Parameters

No parameters provided.

### `log1p()`

log1p returns the natural logarithm of 1 plus `x`.
This operation is more accurate than `math.log(x: 1 + x)` when `x` is
near zero.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `log2()`

log2 is a function returns the binary logarithm of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: the value used in the operation. | Yes |
### `log2e()`

log2e represents the base 2 logarithm of **e** (`math.e`).

#### Function type signature

```flux
float
```

#### Parameters

No parameters provided.

### `logb()`

logb returns the binary exponent of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `mInf()`

mInf returns positive infinity if `sign >= 0`, negative infinity
if `sign < 0`.

#### Function type signature

```flux
(sign: int) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| sign | sign: Value to operate on. | Yes |
### `mMax()`

mMax returns the larger of `x` or `y`.

#### Function type signature

```flux
(x: float, y: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: x-value to use in the operation. | Yes |
| y | y: y-value to use in the operation. | Yes |
### `mMin()`

mMin is a function taht returns the lessser of `x` or `y`.

#### Function type signature

```flux
(x: float, y: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: x-value to use in the operation. | Yes |
| y | y: y-value to use in the operation. | Yes |
### `maxfloat()`

maxfloat represents the maximum float value.

#### Function type signature

```flux
float
```

#### Parameters

No parameters provided.

### `maxint()`

maxint represents the maximum integer value (`2^63 - 1`).

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `maxuint()`

maxuint representes the maximum unsigned integer value  (`2^64 - 1`).

#### Function type signature

```flux
uint
```

#### Parameters

No parameters provided.

### `minint()`

minint represents the minimum integer value (`-2^63`).

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `mod()`

mod returns a floating-point remainder of `x/y`.

#### Function type signature

```flux
(x: float, y: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: x-value to use in the operation. | Yes |
| y | y: y-value to use in the operation. | Yes |
### `modf()`

modf returns integer and fractional floating-point numbers that sum to `f`.

#### Function type signature

```flux
(f: float) => {int: float, frac: float}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| f | f: Value to operate on. | Yes |
### `nextafter()`

nextafter returns the next representable float value after `x` towards `y`.

#### Function type signature

```flux
(x: float, y: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: x-value to use in the operation. | Yes |
| y | y: y-value to use in the operation. | Yes |
### `phi()`

phi represents the [Golden Ratio](https://www.britannica.com/science/golden-ratio).

#### Function type signature

```flux
float
```

#### Parameters

No parameters provided.

### `pi()`

pi represents pi (π).

#### Function type signature

```flux
float
```

#### Parameters

No parameters provided.

### `pow()`

pow returns `x**y`, the base-x exponential of `y`.

#### Function type signature

```flux
(x: float, y: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Base value to operate on. | Yes |
| y | y: Exponent value. | Yes |
### `pow10()`

pow10 returns 10**n, the base-10 exponential of `n`.

#### Function type signature

```flux
(n: int) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Exponent value. | Yes |
### `remainder()`

remainder returns the IEEE 754 floating-point remainder of `x/y`.

#### Function type signature

```flux
(x: float, y: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Numerator to use in the operation. | Yes |
| y | y: Denominator to use in the operation. | Yes |
### `round()`

round returns the nearest integer, rounding half away from zero.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `roundtoeven()`

roundtoeven returns the nearest integer, rounding ties to even.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `signbit()`

signbit reports whether `x` is negative or negative zero.

#### Function type signature

```flux
(x: float) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to evaluate. | Yes |
### `sin()`

sin returns the sine of the radian argument `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Radian value to use in the operation. | Yes |
### `sincos()`

sincos returns the values of `math.sin(x:x)` and `math.cos(x:x)`.

#### Function type signature

```flux
(x: float) => {sin: float, cos: float}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `sinh()`

sinh returns the hyperbolic sine of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `smallestNonzeroFloat()`

smallestNonzeroFloat represents the smallest nonzero float value.

#### Function type signature

```flux
float
```

#### Parameters

No parameters provided.

### `sqrt()`

sqrt returns the square root of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `sqrt2()`

sqrt2 represents the square root of 2.

#### Function type signature

```flux
float
```

#### Parameters

No parameters provided.

### `sqrte()`

sqrte represents the square root of **e** (`math.e`).

#### Function type signature

```flux
float
```

#### Parameters

No parameters provided.

### `sqrtphi()`

sqrtphi represents the square root of phi (`math.phi`), the Golden Ratio.

#### Function type signature

```flux
float
```

#### Parameters

No parameters provided.

### `sqrtpi()`

sqrtpi represents the square root of pi (π).

#### Function type signature

```flux
float
```

#### Parameters

No parameters provided.

### `tan()`

tan returns the tangent of the radian argument `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `tanh()`

tanh returns the hyperbolic tangent of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `trunc()`

trunc returns the integer value of `x`.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `y0()`

y0 returns the order-zero Bessel function of the second kind.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `y1()`

y1 returns the order-one Bessel function of the second kind.

#### Function type signature

```flux
(x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: Value to operate on. | Yes |
### `yn()`

yn returns the order-n Bessel function of the second kind.

#### Function type signature

```flux
(n: int, x: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Order number to use in the operation. | Yes |
| x | x: Value to operate on. | Yes |
