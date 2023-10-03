## `zenoss` package

Package zenoss provides functions that send events to [Zenoss](https://www.zenoss.com/).

Import the `zenoss` package:

```flux
import "zenoss"
```

### Functions

### `endpoint()`

endpoint sends events to Zenoss using data from input rows.

#### Function type signature

```flux
(
    url: string,
    ?action: A,
    ?apiKey: B,
    ?method: C,
    ?password: string,
    ?tid: D,
    ?type: E,
    ?username: string,
) => (
    mapFn: (
        r: F,
    ) => {
        G with
        summary: O,
        severity: N,
        message: M,
        eventClassKey: L,
        eventClass: K,
        device: J,
        component: I,
        collector: H,
    },
) => (<-tables: stream[F]) => stream[{F with _sent: string}] where B: Equatable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: Zenoss [router endpoint URL](https://help.zenoss.com/zsd/RM/configuring-resource-manager/enabling-access-to-browser-interfaces/creating-and-changing-public-endpoints). | Yes |
| username | username: Zenoss username to use for HTTP BASIC authentication.
  Default is `""` (no authentication). | No |
| password | password: Zenoss password to use for HTTP BASIC authentication.
  Default is `""` (no authentication). | No |
| apiKey | apiKey: Zenoss cloud API key.
  Default is `""` (no API key). | No |
| action | action: Zenoss router name.
  Default is `"EventsRouter"`. | No |
| method | method: EventsRouter method.
  Default is `"add_event"`. | No |
| type | type: Event type. Default is `"rpc"`. | No |
| tid | tid: Temporary request transaction ID.
  Default is `1`. | No |
### `event()`

event sends an event to [Zenoss](https://www.zenoss.com/).

#### Function type signature

```flux
(
    severity: A,
    url: string,
    ?action: B,
    ?apiKey: C,
    ?collector: D,
    ?component: E,
    ?device: F,
    ?eventClass: G,
    ?eventClassKey: H,
    ?message: I,
    ?method: J,
    ?password: string,
    ?summary: K,
    ?tid: L,
    ?type: M,
    ?username: string,
) => int where C: Equatable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: Zenoss [router endpoint URL](https://help.zenoss.com/zsd/RM/configuring-resource-manager/enabling-access-to-browser-interfaces/creating-and-changing-public-endpoints). | Yes |
| username | username: Zenoss username to use for HTTP BASIC authentication.
  Default is `""` (no authentication). | No |
| password | password: Zenoss password to use for HTTP BASIC authentication.
  Default is `""` (no authentication). | No |
| apiKey | apiKey: Zenoss cloud API key.
  Default is `""` (no API key). | No |
| action | action: Zenoss router name.
  Default is "EventsRouter". | No |
| method | method: [EventsRouter method](https://help.zenoss.com/dev/collection-zone-and-resource-manager-apis/codebase/routers/router-reference/eventsrouter).
  Default is "add_event". | No |
| type | type: Event type.
  Default is "rpc". | No |
| tid | tid: Temporary request transaction ID.
  Default is `1`. | No |
| summary | summary: Event summary.
  Default is `""`. | No |
| device | device: Related device.
  Default is `""`. | No |
| component | component: Related component.
  Default is `""`. | No |
| severity | severity: [Event severity level](https://help.zenoss.com/zsd/RM/administering-resource-manager/event-management/event-severity-levels). | Yes |
| eventClass | eventClass: [Event class](https://help.zenoss.com/zsd/RM/administering-resource-manager/event-management/understanding-event-classes).
  Default is `""`. | No |
| eventClassKey | eventClassKey: Event [class key](https://help.zenoss.com/zsd/RM/administering-resource-manager/event-management/event-fields).
  Default is `""`. | No |
| collector | collector: Zenoss [collector](https://help.zenoss.com/zsd/RM/administering-resource-manager/event-management/event-fields).
  Default is `""`. | No |
| message | message: Related message.
  Default is `""`. | No |
