# github.com/speakeasy-sdks/p3ld3v

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/p3ld3v
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
### Example

```go
package main

import (
	"context"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
	"log"
)

func main() {
	s := p3ld3v.New(
		p3ld3v.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Pet.AddPetForm(ctx, shared.Pet{
		Category: &shared.Category{
			ID:   p3ld3v.Int64(1),
			Name: p3ld3v.String("Dogs"),
		},
		ID:   p3ld3v.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"string",
		},
		Tags: []shared.Tag{
			shared.Tag{},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Pet != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Pet](docs/sdks/pet/README.md)

* [AddPetForm](docs/sdks/pet/README.md#addpetform) - Add a new pet to the store
* [AddPetJSON](docs/sdks/pet/README.md#addpetjson) - Add a new pet to the store
* [AddPetRaw](docs/sdks/pet/README.md#addpetraw) - Add a new pet to the store
* [DeletePet](docs/sdks/pet/README.md#deletepet) - Deletes a pet
* [FindPetsByStatus](docs/sdks/pet/README.md#findpetsbystatus) - Finds Pets by status
* [FindPetsByTags](docs/sdks/pet/README.md#findpetsbytags) - Finds Pets by tags
* [GetPetByID](docs/sdks/pet/README.md#getpetbyid) - Find pet by ID
* [UpdatePetWithForm](docs/sdks/pet/README.md#updatepetwithform) - Updates a pet in the store with form data
* [UpdatePetForm](docs/sdks/pet/README.md#updatepetform) - Update an existing pet
* [UpdatePetJSON](docs/sdks/pet/README.md#updatepetjson) - Update an existing pet
* [UpdatePetRaw](docs/sdks/pet/README.md#updatepetraw) - Update an existing pet
* [UploadFile](docs/sdks/pet/README.md#uploadfile) - uploads an image

### [Store](docs/sdks/store/README.md)

* [DeleteOrder](docs/sdks/store/README.md#deleteorder) - Delete purchase order by ID
* [GetInventory](docs/sdks/store/README.md#getinventory) - Returns pet inventories by status
* [GetOrderByID](docs/sdks/store/README.md#getorderbyid) - Find purchase order by ID
* [PlaceOrderForm](docs/sdks/store/README.md#placeorderform) - Place an order for a pet
* [PlaceOrderJSON](docs/sdks/store/README.md#placeorderjson) - Place an order for a pet
* [PlaceOrderRaw](docs/sdks/store/README.md#placeorderraw) - Place an order for a pet

### [User](docs/sdks/user/README.md)

* [CreateUserForm](docs/sdks/user/README.md#createuserform) - Create user
* [CreateUserJSON](docs/sdks/user/README.md#createuserjson) - Create user
* [CreateUserRaw](docs/sdks/user/README.md#createuserraw) - Create user
* [CreateUsersWithListInput](docs/sdks/user/README.md#createuserswithlistinput) - Creates list of users with given input array
* [DeleteUser](docs/sdks/user/README.md#deleteuser) - Delete user
* [GetUserByName](docs/sdks/user/README.md#getuserbyname) - Get user by user name
* [LoginUser](docs/sdks/user/README.md#loginuser) - Logs user into the system
* [LogoutUser](docs/sdks/user/README.md#logoutuser) - Logs out current logged in user session
* [UpdateUserForm](docs/sdks/user/README.md#updateuserform) - Update user
* [UpdateUserJSON](docs/sdks/user/README.md#updateuserjson) - Update user
* [UpdateUserRaw](docs/sdks/user/README.md#updateuserraw) - Update user
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Error Handling -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
	"log"
)

func main() {
	s := p3ld3v.New(
		p3ld3v.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Pet.AddPetForm(ctx, shared.Pet{
		Category: &shared.Category{
			ID:   p3ld3v.Int64(1),
			Name: p3ld3v.String("Dogs"),
		},
		ID:   p3ld3v.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"string",
		},
		Tags: []shared.Tag{
			shared.Tag{},
		},
	})
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://petstore3.swagger.io/api/v3` | None |

#### Example

```go
package main

import (
	"context"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
	"log"
)

func main() {
	s := p3ld3v.New(
		p3ld3v.WithServerIndex(0),
		p3ld3v.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Pet.AddPetForm(ctx, shared.Pet{
		Category: &shared.Category{
			ID:   p3ld3v.Int64(1),
			Name: p3ld3v.String("Dogs"),
		},
		ID:   p3ld3v.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"string",
		},
		Tags: []shared.Tag{
			shared.Tag{},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Pet != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
	"log"
)

func main() {
	s := p3ld3v.New(
		p3ld3v.WithServerURL("https://petstore3.swagger.io/api/v3"),
		p3ld3v.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Pet.AddPetForm(ctx, shared.Pet{
		Category: &shared.Category{
			ID:   p3ld3v.Int64(1),
			Name: p3ld3v.String("Dogs"),
		},
		ID:   p3ld3v.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"string",
		},
		Tags: []shared.Tag{
			shared.Tag{},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Pet != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->



<!-- Start Authentication -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name           | Type           | Scheme         |
| -------------- | -------------- | -------------- |
| `PetstoreAuth` | oauth2         | OAuth2 token   |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
	"log"
)

func main() {
	s := p3ld3v.New(
		p3ld3v.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Pet.AddPetForm(ctx, shared.Pet{
		Category: &shared.Category{
			ID:   p3ld3v.Int64(1),
			Name: p3ld3v.String("Dogs"),
		},
		ID:   p3ld3v.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"string",
		},
		Tags: []shared.Tag{
			shared.Tag{},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Pet != nil {
		// handle response
	}
}

```

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
	"context"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v2"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/operations"
	"log"
)

func main() {
	s := p3ld3v.New()

	operationSecurity := operations.GetPetByIDSecurity{
		APIKey: p3ld3v.String(""),
	}

	ctx := context.Background()
	res, err := s.Pet.GetPetByID(ctx, operations.GetPetByIDRequest{
		PetID: 504151,
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.Pet != nil {
		// handle response
	}
}

```
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
