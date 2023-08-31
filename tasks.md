> sorted by priority

## URL scheme

### Signed in

- /
  - explore page of products
- /@username
  - `Add product`
  - show `solana wallet address`
  - allow to send `solana` to external solana address (cash out)
- /@username/add-product
  - enter all metadata for product, description, add files
  - - `publish` button
- /products/<product-id>
  - how many sales were made

### Unsigned in

- /
  - explore of products
- /products/<product-id>
  - can buy product

## UI

- create product
  - uploading of file(s)
- view product
- pay product

## Go backend

- allow to receive money from externally
- endpoint for show `solana wallet address`
- endpoint to allow to send `solana` to external solana address (cash out)
- support multiple currencies
  - dollar, euro, japan yen
- product should have files attached
  - aws s3. Product `files`: aws_s3_url[]
- user_goods (current table, improve)
  - rename user_goods to `products_bought`
  - `created_by` is used to keep the ID of who created product
  - used for ownership of who bought which product
- error handling
  - send transaction but user has no money
- connect Hanko auth to Go backend to authorise requests
- count how many users bought this product
- if transaction succesfull, `products_bought` has to be updated
  - `user` who bought product. can access `files` of the product
- potentially add other types
  - `Newsletter`
    - `user` will have access to past paid articles + new ones
  - `Article`: `content: string`
