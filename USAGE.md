<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/p3ld3v"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/operations"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/shared"
)

func main() {
    s := p3ld3v.New()
    operationSecurity := operations.AddPetFormSecurity{
            PetstoreAuth: "",
        }

    ctx := context.Background()
    res, err := s.Pet.AddPetForm(ctx, shared.Pet{
        Category: &shared.Category{
            ID: p3ld3v.Int64(1),
            Name: p3ld3v.String("Dogs"),
        },
        ID: p3ld3v.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "corrupti",
        },
        Status: shared.PetStatusPending.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: p3ld3v.Int64(715190),
                Name: p3ld3v.String("Stuart Stiedemann"),
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
<!-- End SDK Example Usage -->