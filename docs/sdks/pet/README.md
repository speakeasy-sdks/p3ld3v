# Pet

## Overview

Everything about your Pets

Find out more
<http://swagger.io>
### Available Operations

* [AddPetForm](#addpetform) - Add a new pet to the store
* [AddPetJSON](#addpetjson) - Add a new pet to the store
* [AddPetRaw](#addpetraw) - Add a new pet to the store
* [DeletePet](#deletepet) - Deletes a pet
* [FindPetsByStatus](#findpetsbystatus) - Finds Pets by status
* [FindPetsByTags](#findpetsbytags) - Finds Pets by tags
* [GetPetByID](#getpetbyid) - Find pet by ID
* [UpdatePetWithForm](#updatepetwithform) - Updates a pet in the store with form data
* [UpdatePetForm](#updatepetform) - Update an existing pet
* [UpdatePetJSON](#updatepetjson) - Update an existing pet
* [UpdatePetRaw](#updatepetraw) - Update an existing pet
* [UploadFile](#uploadfile) - uploads an image

## AddPetForm

Add a new pet to the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/p3ld3v"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/operations"
)

func main() {
    s := bestapievermade.New()
    operationSecurity := operations.AddPetFormSecurity{
            PetstoreAuth: "",
        }

    ctx := context.Background()
    res, err := s.Pet.AddPetForm(ctx, shared.Pet{
        Category: &shared.Category{
            ID: bestapievermade.Int64(1),
            Name: bestapievermade.String("Dogs"),
        },
        ID: bestapievermade.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "ipsam",
        },
        Status: shared.PetStatusSold.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: bestapievermade.Int64(778157),
                Name: bestapievermade.String("Teri Strosin"),
            },
            shared.Tag{
                ID: bestapievermade.Int64(799159),
                Name: bestapievermade.String("Erik Lebsack"),
            },
            shared.Tag{
                ID: bestapievermade.Int64(118274),
                Name: bestapievermade.String("Luke McCullough"),
            },
            shared.Tag{
                ID: bestapievermade.Int64(944669),
                Name: bestapievermade.String("Everett Breitenberg"),
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.Pet](../../models/shared/pet.md)                                       | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.AddPetFormSecurity](../../models/operations/addpetformsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.AddPetFormResponse](../../models/operations/addpetformresponse.md), error**


## AddPetJSON

Add a new pet to the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/p3ld3v"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/operations"
)

func main() {
    s := bestapievermade.New()
    operationSecurity := operations.AddPetJSONSecurity{
            PetstoreAuth: "",
        }

    ctx := context.Background()
    res, err := s.Pet.AddPetJSON(ctx, shared.Pet{
        Category: &shared.Category{
            ID: bestapievermade.Int64(1),
            Name: bestapievermade.String("Dogs"),
        },
        ID: bestapievermade.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "qui",
            "impedit",
        },
        Status: shared.PetStatusSold.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: bestapievermade.Int64(216550),
                Name: bestapievermade.String("Brandon Auer"),
            },
            shared.Tag{
                ID: bestapievermade.Int64(149675),
                Name: bestapievermade.String("Curtis Morissette"),
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.Pet](../../models/shared/pet.md)                                       | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.AddPetJSONSecurity](../../models/operations/addpetjsonsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.AddPetJSONResponse](../../models/operations/addpetjsonresponse.md), error**


## AddPetRaw

Add a new pet to the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/p3ld3v"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/operations"
)

func main() {
    s := bestapievermade.New()
    operationSecurity := operations.AddPetRawSecurity{
            PetstoreAuth: "",
        }

    ctx := context.Background()
    res, err := s.Pet.AddPetRaw(ctx, []byte("saepe"), operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [[]byte](../../models//.md)                                                  | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.AddPetRawSecurity](../../models/operations/addpetrawsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.AddPetRawResponse](../../models/operations/addpetrawresponse.md), error**


## DeletePet

delete a pet

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
    operationSecurity := operations.DeletePetSecurity{
            PetstoreAuth: "",
        }

    ctx := context.Background()
    res, err := s.Pet.DeletePet(ctx, operations.DeletePetRequest{
        APIKey: bestapievermade.String("fuga"),
        PetID: 449950,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.DeletePetRequest](../../models/operations/deletepetrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.DeletePetSecurity](../../models/operations/deletepetsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.DeletePetResponse](../../models/operations/deletepetresponse.md), error**


## FindPetsByStatus

Multiple status values can be provided with comma separated strings

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
    operationSecurity := operations.FindPetsByStatusSecurity{
            PetstoreAuth: "",
        }

    ctx := context.Background()
    res, err := s.Pet.FindPetsByStatus(ctx, operations.FindPetsByStatusRequest{
        Status: operations.FindPetsByStatusStatusPending.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.FindPetsByStatusRequest](../../models/operations/findpetsbystatusrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.FindPetsByStatusSecurity](../../models/operations/findpetsbystatussecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.FindPetsByStatusResponse](../../models/operations/findpetsbystatusresponse.md), error**


## FindPetsByTags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

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
    operationSecurity := operations.FindPetsByTagsSecurity{
            PetstoreAuth: "",
        }

    ctx := context.Background()
    res, err := s.Pet.FindPetsByTags(ctx, operations.FindPetsByTagsRequest{
        Tags: []string{
            "iure",
            "saepe",
            "quidem",
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.FindPetsByTagsRequest](../../models/operations/findpetsbytagsrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.FindPetsByTagsSecurity](../../models/operations/findpetsbytagssecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.FindPetsByTagsResponse](../../models/operations/findpetsbytagsresponse.md), error**


## GetPetByID

Returns a single pet

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
    operationSecurity := operations.GetPetByIDSecurity{
            APIKey: bestapievermade.String(""),
        }

    ctx := context.Background()
    res, err := s.Pet.GetPetByID(ctx, operations.GetPetByIDRequest{
        PetID: 99280,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetPetByIDRequest](../../models/operations/getpetbyidrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.GetPetByIDSecurity](../../models/operations/getpetbyidsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.GetPetByIDResponse](../../models/operations/getpetbyidresponse.md), error**


## UpdatePetWithForm

Updates a pet in the store with form data

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
    operationSecurity := operations.UpdatePetWithFormSecurity{
            PetstoreAuth: "",
        }

    ctx := context.Background()
    res, err := s.Pet.UpdatePetWithForm(ctx, operations.UpdatePetWithFormRequest{
        Name: bestapievermade.String("Lela Orn"),
        PetID: 170909,
        Status: bestapievermade.String("dolorem"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdatePetWithFormRequest](../../models/operations/updatepetwithformrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.UpdatePetWithFormSecurity](../../models/operations/updatepetwithformsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.UpdatePetWithFormResponse](../../models/operations/updatepetwithformresponse.md), error**


## UpdatePetForm

Update an existing pet by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/p3ld3v"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/operations"
)

func main() {
    s := bestapievermade.New()
    operationSecurity := operations.UpdatePetFormSecurity{
            PetstoreAuth: "",
        }

    ctx := context.Background()
    res, err := s.Pet.UpdatePetForm(ctx, shared.Pet{
        Category: &shared.Category{
            ID: bestapievermade.Int64(1),
            Name: bestapievermade.String("Dogs"),
        },
        ID: bestapievermade.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "explicabo",
            "nobis",
        },
        Status: shared.PetStatusAvailable.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: bestapievermade.Int64(363711),
                Name: bestapievermade.String("Velma Batz"),
            },
            shared.Tag{
                ID: bestapievermade.Int64(988374),
                Name: bestapievermade.String("Juan O'Hara"),
            },
            shared.Tag{
                ID: bestapievermade.Int64(161309),
                Name: bestapievermade.String("Shaun McCullough"),
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [shared.Pet](../../models/shared/pet.md)                                             | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.UpdatePetFormSecurity](../../models/operations/updatepetformsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.UpdatePetFormResponse](../../models/operations/updatepetformresponse.md), error**


## UpdatePetJSON

Update an existing pet by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/p3ld3v"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/operations"
)

func main() {
    s := bestapievermade.New()
    operationSecurity := operations.UpdatePetJSONSecurity{
            PetstoreAuth: "",
        }

    ctx := context.Background()
    res, err := s.Pet.UpdatePetJSON(ctx, shared.Pet{
        Category: &shared.Category{
            ID: bestapievermade.Int64(1),
            Name: bestapievermade.String("Dogs"),
        },
        ID: bestapievermade.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "molestiae",
            "velit",
        },
        Status: shared.PetStatusPending.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: bestapievermade.Int64(338007),
                Name: bestapievermade.String("Kayla O'Kon"),
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [shared.Pet](../../models/shared/pet.md)                                             | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.UpdatePetJSONSecurity](../../models/operations/updatepetjsonsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.UpdatePetJSONResponse](../../models/operations/updatepetjsonresponse.md), error**


## UpdatePetRaw

Update an existing pet by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/p3ld3v"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/operations"
)

func main() {
    s := bestapievermade.New()
    operationSecurity := operations.UpdatePetRawSecurity{
            PetstoreAuth: "",
        }

    ctx := context.Background()
    res, err := s.Pet.UpdatePetRaw(ctx, []byte("quo"), operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [[]byte](../../models//.md)                                                        | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.UpdatePetRawSecurity](../../models/operations/updatepetrawsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.UpdatePetRawResponse](../../models/operations/updatepetrawresponse.md), error**


## UploadFile

uploads an image

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
    operationSecurity := operations.UploadFileSecurity{
            PetstoreAuth: "",
        }

    ctx := context.Background()
    res, err := s.Pet.UploadFile(ctx, operations.UploadFileRequest{
        RequestBody: []byte("sequi"),
        AdditionalMetadata: bestapievermade.String("tenetur"),
        PetID: 368725,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.APIResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.UploadFileRequest](../../models/operations/uploadfilerequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.UploadFileSecurity](../../models/operations/uploadfilesecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.UploadFileResponse](../../models/operations/uploadfileresponse.md), error**

