# github.com/dubinc/dub-go

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/dubinc/dub-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example 1

```go
package main

import (
	"context"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
	)
	var request *operations.CreateLinkRequestBody = &operations.CreateLinkRequestBody{
		URL:        "https://google/com",
		ExternalID: dubgo.String("123456"),
		TagIds: operations.CreateTagIdsArrayOfStr(
			[]string{
				"clux0rgak00011...",
			},
		),
	}
	ctx := context.Background()
	res, err := s.Links.Create(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

### Example 2

```go
package main

import (
	"context"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
	)
	var request *operations.UpsertLinkRequestBody = &operations.UpsertLinkRequestBody{
		URL:        "https://google/com",
		ExternalID: dubgo.String("123456"),
		TagIds: operations.CreateUpsertLinkTagIdsArrayOfStr(
			[]string{
				"clux0rgak00011...",
			},
		),
	}
	ctx := context.Background()
	res, err := s.Links.Upsert(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Links](docs/sdks/links/README.md)

* [List](docs/sdks/links/README.md#list) - Retrieve a list of links
* [Create](docs/sdks/links/README.md#create) - Create a new link
* [Count](docs/sdks/links/README.md#count) - Retrieve links count
* [Get](docs/sdks/links/README.md#get) - Retrieve a link
* [Delete](docs/sdks/links/README.md#delete) - Delete a link
* [Update](docs/sdks/links/README.md#update) - Update a link
* [CreateMany](docs/sdks/links/README.md#createmany) - Bulk create links
* [UpdateMany](docs/sdks/links/README.md#updatemany) - Bulk update links
* [Upsert](docs/sdks/links/README.md#upsert) - Upsert a link

### [QRCodes](docs/sdks/qrcodes/README.md)

* [Get](docs/sdks/qrcodes/README.md#get) - Retrieve a QR code

### [Analytics](docs/sdks/analytics/README.md)

* [Retrieve](docs/sdks/analytics/README.md#retrieve) - Retrieve analytics for a link, a domain, or the authenticated workspace.

### [Workspaces](docs/sdks/workspaces/README.md)

* [Get](docs/sdks/workspaces/README.md#get) - Retrieve a workspace
* [Update](docs/sdks/workspaces/README.md#update) - Update a workspace

### [Tags](docs/sdks/tags/README.md)

* [List](docs/sdks/tags/README.md#list) - Retrieve a list of tags
* [Create](docs/sdks/tags/README.md#create) - Create a new tag
* [Update](docs/sdks/tags/README.md#update) - Update a tag

### [Domains](docs/sdks/domains/README.md)

* [List](docs/sdks/domains/README.md#list) - Retrieve a list of domains
* [Create](docs/sdks/domains/README.md#create) - Create a domain
* [Delete](docs/sdks/domains/README.md#delete) - Delete a domain
* [Update](docs/sdks/domains/README.md#update) - Update a domain

### [Track](docs/sdks/track/README.md)

* [Lead](docs/sdks/track/README.md#lead) - Track a lead
* [Sale](docs/sdks/track/README.md#sale) - Track a sale
* [Customer](docs/sdks/track/README.md#customer) - Track a customer

### [Metatags](docs/sdks/metatags/README.md)

* [Get](docs/sdks/metatags/README.md#get) - Retrieve the metatags for a URL
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequest          | 400                           | application/json              |
| sdkerrors.Unauthorized        | 401                           | application/json              |
| sdkerrors.Forbidden           | 403                           | application/json              |
| sdkerrors.NotFound            | 404                           | application/json              |
| sdkerrors.Conflict            | 409                           | application/json              |
| sdkerrors.InviteExpired       | 410                           | application/json              |
| sdkerrors.UnprocessableEntity | 422                           | application/json              |
| sdkerrors.RateLimitExceeded   | 429                           | application/json              |
| sdkerrors.InternalServerError | 500                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

### Example

```go
package main

import (
	"context"
	"errors"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"github.com/dubinc/dub-go/models/sdkerrors"
	"log"
)

func main() {
	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
	)
	request := operations.GetLinksRequest{}
	ctx := context.Background()
	res, err := s.Links.List(ctx, request)
	if err != nil {

		var e *sdkerrors.BadRequest
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.Unauthorized
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.Forbidden
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.NotFound
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.Conflict
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.InviteExpired
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.UnprocessableEntity
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.RateLimitExceeded
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.InternalServerError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.dub.co` | None |

#### Example

```go
package main

import (
	"context"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
	s := dubgo.New(
		dubgo.WithServerIndex(0),
		dubgo.WithSecurity("DUB_API_KEY"),
	)
	request := operations.GetLinksRequest{}
	ctx := context.Background()
	res, err := s.Links.List(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
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
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
	s := dubgo.New(
		dubgo.WithServerURL("https://api.dub.co"),
		dubgo.WithSecurity("DUB_API_KEY"),
	)
	request := operations.GetLinksRequest{}
	ctx := context.Background()
	res, err := s.Links.List(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
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
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name        | Type        | Scheme      |
| ----------- | ----------- | ----------- |
| `Token`     | http        | HTTP Bearer |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	dubgo "github.com/dubinc/dub-go"
	"github.com/dubinc/dub-go/models/operations"
	"log"
)

func main() {
	s := dubgo.New(
		dubgo.WithSecurity("DUB_API_KEY"),
	)
	request := operations.GetLinksRequest{}
	ctx := context.Background()
	res, err := s.Links.List(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
