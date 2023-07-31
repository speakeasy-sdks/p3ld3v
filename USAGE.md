<!-- Start SDK Example Usage -->


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
            "provident",
            "distinctio",
            "quibusdam",
        },
        Status: shared.PetStatusPending.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: bestapievermade.Int64(544883),
                Name: bestapievermade.String("Ben Mueller"),
            },
            shared.Tag{
                ID: bestapievermade.Int64(437587),
                Name: bestapievermade.String("Raquel Bednar"),
            },
            shared.Tag{
                ID: bestapievermade.Int64(383441),
                Name: bestapievermade.String("Alexandra Schulist"),
            },
            shared.Tag{
                ID: bestapievermade.Int64(568045),
                Name: bestapievermade.String("Mrs. Sophie Smith MD"),
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