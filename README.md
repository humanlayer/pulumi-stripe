# Pulumi Stripe Provider

A Pulumi provider for managing Stripe resources, built by bridging the official [Stripe Terraform Provider](https://docs.stripe.com/terraform/install) using [pulumi-terraform-bridge](https://github.com/pulumi/pulumi-terraform-bridge).

## Installation

```bash
npm install @humanlayer/pulumi-stripe
```

## Usage

```typescript
import * as stripe from "@humanlayer/pulumi-stripe";

// Create a product
const product = new stripe.Product("my-product", {
    name: "My SaaS",
    description: "A product managed with Pulumi",
});

// Create a price for the product
const price = new stripe.Price("my-price", {
    product: product.id,
    currency: "usd",
    unitAmount: 1999, // $19.99
    recurring: {
        interval: "month",
    },
});

// Create a coupon
const coupon = new stripe.Coupon("launch-discount", {
    percentOff: 25,
    duration: "repeating",
    durationInMonths: 3,
});

// Create a webhook endpoint
const webhook = new stripe.WebhookEndpoint("my-webhook", {
    url: "https://api.example.com/stripe/webhooks",
    enabledEvents: [
        "customer.subscription.created",
        "customer.subscription.deleted",
        "invoice.paid",
    ],
});

export const productId = product.id;
export const priceId = price.id;
```

## Available Resources

- `Product` - Products and services
- `Price` - Pricing for products
- `Customer` - Customer records
- `Coupon` - Discount coupons
- `PromotionCode` - Promotion codes for coupons
- `WebhookEndpoint` - Webhook endpoints
- `TaxRate` - Tax rates
- `ShippingRate` - Shipping rates
- `BillingMeter` - Usage-based billing meters
- `EntitlementsFeature` - Entitlement features

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
