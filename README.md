# Pulumi Stripe Provider

A Pulumi provider for managing Stripe resources, built by bridging the official [Stripe Terraform Provider](https://docs.stripe.com/terraform/install) using [pulumi-terraform-bridge](https://github.com/pulumi/pulumi-terraform-bridge).

## Installation

```bash
npm install @humanlayer/pulumi-stripe
```

## Usage

```typescript
import * as stripe from "@humanlayer/pulumi-stripe";

const product = new stripe.Product("my-product", {
    name: "My Product",
    description: "A product managed with Pulumi",
});

export const productId = product.id;
```

## Configuration

Set your Stripe API key:

```bash
export STRIPE_API_KEY=sk_test_...
```

Or configure via Pulumi config:

```bash
pulumi config set stripe:apiKey sk_test_... --secret
```

## Building from Source

```bash
make build
```

This will:
1. Generate the Pulumi schema from the Terraform provider
2. Build the provider binary
3. Generate and compile the TypeScript SDK

## Resources

- [Stripe Terraform Provider](https://docs.stripe.com/terraform/install) - Upstream provider
- [pulumi-terraform-bridge](https://github.com/pulumi/pulumi-terraform-bridge) - Bridge library
- [pulumi-tf-provider-boilerplate](https://github.com/pulumi/pulumi-tf-provider-boilerplate) - Project template

## License

MIT
