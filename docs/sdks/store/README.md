# Store

## Overview

Access to Petstore orders

Find out more about our store
<http://swagger.io>
### Available Operations

* [DeleteOrder](#deleteorder) - Delete purchase order by ID
* [GetInventory](#getinventory) - Returns pet inventories by status
* [GetOrderByID](#getorderbyid) - Find purchase order by ID
* [PlaceOrderForm](#placeorderform) - Place an order for a pet
* [PlaceOrderJSON](#placeorderjson) - Place an order for a pet
* [PlaceOrderRaw](#placeorderraw) - Place an order for a pet

## DeleteOrder

For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/p3ld3v"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/operations"
)

func main() {
    s := bestapievermade.New()

    ctx := context.Background()
    res, err := s.Store.DeleteOrder(ctx, operations.DeleteOrderRequest{
        OrderID: 414662,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.DeleteOrderRequest](../../models/operations/deleteorderrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.DeleteOrderResponse](../../models/operations/deleteorderresponse.md), error**


## GetInventory

Returns a map of status codes to quantities

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/p3ld3v"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/operations"
)

func main() {
    s := bestapievermade.New()
    operationSecurity := operations.GetInventorySecurity{
            APIKey: "",
        }

    ctx := context.Background()
    res, err := s.Store.GetInventory(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetInventory200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `security`                                                                         | [operations.GetInventorySecurity](../../models/operations/getinventorysecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.GetInventoryResponse](../../models/operations/getinventoryresponse.md), error**


## GetOrderByID

For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/p3ld3v"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/operations"
)

func main() {
    s := bestapievermade.New()

    ctx := context.Background()
    res, err := s.Store.GetOrderByID(ctx, operations.GetOrderByIDRequest{
        OrderID: 473600,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetOrderByIDRequest](../../models/operations/getorderbyidrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetOrderByIDResponse](../../models/operations/getorderbyidresponse.md), error**


## PlaceOrderForm

Place a new order in the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/p3ld3v"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/pkg/types"
)

func main() {
    s := bestapievermade.New()

    ctx := context.Background()
    res, err := s.Store.PlaceOrderForm(ctx, shared.Order{
        Complete: bestapievermade.Bool(false),
        ID: bestapievermade.Int64(10),
        PetID: bestapievermade.Int64(198772),
        Quantity: bestapievermade.Int(7),
        ShipDate: types.MustTimeFromString("2022-10-24T23:43:42.369Z"),
        Status: shared.OrderStatusApproved.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Order](../../models/shared/order.md)          | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PlaceOrderFormResponse](../../models/operations/placeorderformresponse.md), error**


## PlaceOrderJSON

Place a new order in the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/p3ld3v"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/pkg/types"
)

func main() {
    s := bestapievermade.New()

    ctx := context.Background()
    res, err := s.Store.PlaceOrderJSON(ctx, shared.Order{
        Complete: bestapievermade.Bool(false),
        ID: bestapievermade.Int64(10),
        PetID: bestapievermade.Int64(198772),
        Quantity: bestapievermade.Int(7),
        ShipDate: types.MustTimeFromString("2020-10-16T01:47:24.760Z"),
        Status: shared.OrderStatusApproved.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Order](../../models/shared/order.md)          | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PlaceOrderJSONResponse](../../models/operations/placeorderjsonresponse.md), error**


## PlaceOrderRaw

Place a new order in the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/p3ld3v"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/pkg/types"
)

func main() {
    s := bestapievermade.New()

    ctx := context.Background()
    res, err := s.Store.PlaceOrderRaw(ctx, []byte("esse"))
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]byte](../../models//.md)                           | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PlaceOrderRawResponse](../../models/operations/placeorderrawresponse.md), error**

