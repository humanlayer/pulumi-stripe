import * as pulumi from "@pulumi/pulumi";
import * as stripe from "@humanlayer/pulumi-stripe";

// Create a Stripe product
const product = new stripe.Product("example-product", {
    name: "Example Product",
    description: "A product created with Pulumi",
});

// Export the product ID
export const productId = product.id;
