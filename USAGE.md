<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/p3ld3v"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/shared"
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