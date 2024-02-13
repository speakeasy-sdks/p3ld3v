<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	p3ld3v "github.com/speakeasy-sdks/p3ld3v/v3"
	"github.com/speakeasy-sdks/p3ld3v/v3/pkg/models/shared"
	"log"
)

func main() {
	s := p3ld3v.New(
		p3ld3v.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Pet.AddPetForm(ctx, shared.Pet{
		ID:   p3ld3v.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"string",
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
<!-- End SDK Example Usage [usage] -->