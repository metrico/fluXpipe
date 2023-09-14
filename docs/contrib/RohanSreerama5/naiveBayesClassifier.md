## `naiveBayesClassifier` package

Package naiveBayesClassifier provides an implementation of
a naive Bayes classifier.

Import the `naiveBayesClassifier` package:

```flux
import "naiveBayesClassifier"
```

### Functions

### `naiveBayes()`

naiveBayes performs a naive Bayes classification.

#### Function type signature

```flux
(
    <-tables: stream[{C with _time: time, _measurement: E, _field: D}],
    myClass: string,
    myField: A,
    myMeasurement: B,
) => stream[F] where A: Equatable, B: Equatable, D: Equatable, E: Equatable, F: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| myMeasurement | myMeasurement: Measurement to use as training data. | Yes |
| myField | myField: Field to use as training data. | Yes |
| myClass | myClass: Class to classify against. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
