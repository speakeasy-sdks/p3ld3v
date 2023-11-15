# User
(*User*)

## Overview

Operations about user

### Available Operations

* [CreateUserForm](#createuserform) - Create user
* [CreateUserJSON](#createuserjson) - Create user
* [CreateUserRaw](#createuserraw) - Create user
* [CreateUsersWithListInput](#createuserswithlistinput) - Creates list of users with given input array
* [DeleteUser](#deleteuser) - Delete user
* [GetUserByName](#getuserbyname) - Get user by user name
* [LoginUser](#loginuser) - Logs user into the system
* [LogoutUser](#logoutuser) - Logs out current logged in user session
* [UpdateUserForm](#updateuserform) - Update user
* [UpdateUserJSON](#updateuserjson) - Update user
* [UpdateUserRaw](#updateuserraw) - Update user

## CreateUserForm

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
)

func main() {
    s := p3ld3v.New(
        p3ld3v.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.CreateUserForm(ctx, &shared.User{
        Email: p3ld3v.String("john@email.com"),
        FirstName: p3ld3v.String("John"),
        ID: p3ld3v.Int64(10),
        LastName: p3ld3v.String("James"),
        Password: p3ld3v.String("12345"),
        Phone: p3ld3v.String("12345"),
        UserStatus: p3ld3v.Int(1),
        Username: p3ld3v.String("theUser"),
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.User](../../pkg/models/shared/user.md)        | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateUserFormResponse](../../pkg/models/operations/createuserformresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CreateUserJSON

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
)

func main() {
    s := p3ld3v.New(
        p3ld3v.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.CreateUserJSON(ctx, &shared.User{
        Email: p3ld3v.String("john@email.com"),
        FirstName: p3ld3v.String("John"),
        ID: p3ld3v.Int64(10),
        LastName: p3ld3v.String("James"),
        Password: p3ld3v.String("12345"),
        Phone: p3ld3v.String("12345"),
        UserStatus: p3ld3v.Int(1),
        Username: p3ld3v.String("theUser"),
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.User](../../pkg/models/shared/user.md)        | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateUserJSONResponse](../../pkg/models/operations/createuserjsonresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CreateUserRaw

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
)

func main() {
    s := p3ld3v.New(
        p3ld3v.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.CreateUserRaw(ctx, &[]byte("0xB4dDB1Eeed"))
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]byte](../../.md)                                   | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateUserRawResponse](../../pkg/models/operations/createuserrawresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CreateUsersWithListInput

Creates list of users with given input array

### Example Usage

```go
package main

import(
	"context"
	"log"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
)

func main() {
    s := p3ld3v.New(
        p3ld3v.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.CreateUsersWithListInput(ctx, &[]shared.User{
        shared.User{
            Email: p3ld3v.String("john@email.com"),
            FirstName: p3ld3v.String("John"),
            ID: p3ld3v.Int64(10),
            LastName: p3ld3v.String("James"),
            Password: p3ld3v.String("12345"),
            Phone: p3ld3v.String("12345"),
            UserStatus: p3ld3v.Int(1),
            Username: p3ld3v.String("theUser"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]shared.User](../../.md)                            | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateUsersWithListInputResponse](../../pkg/models/operations/createuserswithlistinputresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteUser

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/operations"
)

func main() {
    s := p3ld3v.New(
        p3ld3v.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.DeleteUser(ctx, operations.DeleteUserRequest{
        Username: "Demetris_Torphy",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteUserRequest](../../pkg/models/operations/deleteuserrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DeleteUserResponse](../../pkg/models/operations/deleteuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetUserByName

Get user by user name

### Example Usage

```go
package main

import(
	"context"
	"log"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/operations"
)

func main() {
    s := p3ld3v.New(
        p3ld3v.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.GetUserByName(ctx, operations.GetUserByNameRequest{
        Username: "Zachery_Schneider",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetUserByNameRequest](../../pkg/models/operations/getuserbynamerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetUserByNameResponse](../../pkg/models/operations/getuserbynameresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## LoginUser

Logs user into the system

### Example Usage

```go
package main

import(
	"context"
	"log"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/operations"
)

func main() {
    s := p3ld3v.New(
        p3ld3v.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.LoginUser(ctx, operations.LoginUserRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONRes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.LoginUserRequest](../../pkg/models/operations/loginuserrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.LoginUserResponse](../../pkg/models/operations/loginuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## LogoutUser

Logs out current logged in user session

### Example Usage

```go
package main

import(
	"context"
	"log"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
)

func main() {
    s := p3ld3v.New(
        p3ld3v.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.LogoutUser(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.LogoutUserResponse](../../pkg/models/operations/logoutuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateUserForm

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/operations"
)

func main() {
    s := p3ld3v.New(
        p3ld3v.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.UpdateUserForm(ctx, operations.UpdateUserFormRequest{
        User: &shared.User{
            Email: p3ld3v.String("john@email.com"),
            FirstName: p3ld3v.String("John"),
            ID: p3ld3v.Int64(10),
            LastName: p3ld3v.String("James"),
            Password: p3ld3v.String("12345"),
            Phone: p3ld3v.String("12345"),
            UserStatus: p3ld3v.Int(1),
            Username: p3ld3v.String("theUser"),
        },
        Username: "Bo_Lynch4",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateUserFormRequest](../../pkg/models/operations/updateuserformrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateUserFormResponse](../../pkg/models/operations/updateuserformresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateUserJSON

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/operations"
)

func main() {
    s := p3ld3v.New(
        p3ld3v.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.UpdateUserJSON(ctx, operations.UpdateUserJSONRequest{
        User: &shared.User{
            Email: p3ld3v.String("john@email.com"),
            FirstName: p3ld3v.String("John"),
            ID: p3ld3v.Int64(10),
            LastName: p3ld3v.String("James"),
            Password: p3ld3v.String("12345"),
            Phone: p3ld3v.String("12345"),
            UserStatus: p3ld3v.Int(1),
            Username: p3ld3v.String("theUser"),
        },
        Username: "Alanna_Waters81",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateUserJSONRequest](../../pkg/models/operations/updateuserjsonrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateUserJSONResponse](../../pkg/models/operations/updateuserjsonresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateUserRaw

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/operations"
)

func main() {
    s := p3ld3v.New(
        p3ld3v.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.UpdateUserRaw(ctx, operations.UpdateUserRawRequest{
        RequestBody: []byte("0xf4D36eFb83"),
        Username: "Eleonore2",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateUserRawRequest](../../pkg/models/operations/updateuserrawrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.UpdateUserRawResponse](../../pkg/models/operations/updateuserrawresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
