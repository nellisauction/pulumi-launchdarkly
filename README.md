# LaunchDarkly Resource Provider

The LaunchDarkly Resource Provider lets you manage [LaunchDarkly](https://launchdarkly.com/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @nellisauction/pulumi-launchdarkly
```

or `yarn`:

```bash
yarn add @nellisauction/pulumi-launchdarkly
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_launchdarkly
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/nellisauction/pulumi-launchdarkly/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.LaunchDarkly
```

## Configuration

The following configuration points are available for the `launchdarkly` provider:

- `launchdarkly:accessToken` (environment: `LAUNCHDARKLY_ACCESS_TOKEN`) - a LaunchDarkly personal or service access token
- `launchdarkly:oauthToken` (environment: `LAUNCHDARKLY_OAUTH_TOKEN`) - a LaunchDarkly OAuth V2 token (alternative to `accessToken`)
- `launchdarkly:apiHost` (environment: `LAUNCHDARKLY_API_HOST`) - the LaunchDarkly API host (optional, defaults to `https://app.launchdarkly.com`)
- `launchdarkly:httpTimeout` - HTTP timeout in seconds (optional, defaults to 20)
- `launchdarkly:maxConcurrency` - maximum concurrent API requests (optional, defaults to 1)
- `launchdarkly:archiveFlagsOnDestroy` - archive feature flags instead of deleting them on destroy (optional)

## Upstream

This package bridges [launchdarkly/terraform-provider-launchdarkly](https://github.com/launchdarkly/terraform-provider-launchdarkly). The upstream repo tags `v3.x` releases without a `/v3` Go module path, so `provider/go.mod` pins it by commit pseudo-version (currently `v3.2.0` = `ebb166743915`). To upgrade:

```bash
cd provider && go get github.com/launchdarkly/terraform-provider-launchdarkly@<commit-sha> && go mod tidy
```

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/launchdarkly/api-docs/).
