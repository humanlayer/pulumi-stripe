# Stripe Pulumi Provider Examples

## Prerequisites

1. Install the Pulumi CLI: https://www.pulumi.com/docs/install/
2. Set your Stripe API key: `export STRIPE_API_KEY=sk_test_...`
3. Build the provider: `make build` (from repository root)
4. Add the provider to your PATH: `export PATH=$PWD/bin:$PATH`

## Running the Basic Example

```bash
cd basic-ts
npm install
pulumi stack init dev
pulumi preview
# If preview looks good:
pulumi up
```

## Cleaning Up

```bash
pulumi destroy
pulumi stack rm dev
```
