## Quick Start
- Feature name: `Register Customer`
- `POST : {{base_url}}/api/v1/customers/register`
- Description : Register customer first for login to authenticate
    ```json
    {
        "name": "Vera",
        "username": "vera426",
        "password": "administrator",
        "email": "vera@gmail.com"
    }
    ```

- Feature name: `Login Customer`
- `POST: {{base_url}}/api/v1/customers/login`
- Description: Login authenticate & authorization
    ```json
    {
        "username": "",
        "password": ""
    }
    ```

- Feature name: `Create product category`
- `POST: {{base_url}}/api/v1/category/create`
- Headers: 

    | Key | Value | Description |
    |----------|----------|----------|
    | Content-Type | application/json | - |
    | Authorization | Bearer Token | access token after login |
- Headers: 

    | Key | Value | Description |
    |----------|----------|----------|
    | Content-Type | application/json | - |
    | Authorization | Bearer Token | access token after login |
- Description : Create category of product
    ```json
    {
        "name": "Indomie"
    }
    ```

- Feature name: `Create product`
- `POST: {{base_url}}/api/v1/category/products`
- Headers: 

    | Key | Value | Description |
    |----------|----------|----------|
    | Content-Type | application/json | - |
    | Authorization | Bearer Token | access token after login |
- Description : Create product
    ```json
    {
        "category_id": 1,
        "name": "Mie Kuah",
        "description": "Indofood products",
        "quantity": 100,
        "price": 50000
    }
    ```

- Feature name: `List category`
- `POST: {{base_url}}/api/v1/category`
- Description : Show all category
- Response : 
    ```json
    {
        "code": 200,
        "data": [
            {
                "id": 2,
                "name": "Indomie"
            }
        ],
        "message": "success retrieve categories"
    }
    ```

- Feature name: `Show products by category`
- `GET: {{base_url}}/api/v1/category/products/:category_id`
- Headers: 

    | Key | Value | Description |
    |----------|----------|----------|    
    | Authorization | Bearer Token | access token after login |
- Description : Show all products by category
- Response : 
    ```json
    {
        "code": 200,
        "data": [
            {
                "id": 1,
                "category_id": 2,
                "name": "Mie Goreng",
                "description": "Indofood products",
                "quantity": 100,
                "price": 50000,
                "created_at": "2024-06-14T19:11:48.962431Z",
                "updated_at": "2024-06-14T19:11:48.962431Z"
            }
        ],
        "message": "success retrieve products"
    }
    ```

- Feature name: `Show list products of customer`
- `GET: {{base_url}}/api/v1/customers/carts`
- Headers: 

    | Key | Value | Description |
    |----------|----------|----------|    
    | Authorization | Bearer Token | access token after login |
- Description: `Customers can see a list of products that have been added to the shopping cart`
- Response: 
    ```json
    {
        "code": 200,
        "data": [
            {
                "name": "Philips Jose Patric",
                "carts": [
                    {
                        "id": 1,
                        "products": [
                            {
                                "name": "Mie Kuah",
                                "description": "Indofood products"
                            },
                            {
                                "name": "Mie Goreng",
                                "description": "test"
                            }
                        ]
                    }
                ]
            }
        ],
        "message": "success retrieve customers"
    }
    ```

- Feature name: `Customer create a cart first`
- Headers: 

    | Key | Value | Description |
    |----------|----------|----------|
    | Content-Type | application/json | - |
    | Authorization | Bearer Token | access token after login |
- Description : The customer takes a basket to store their shopping items.
- `POST: {{base_url}}/api/v1/carts`

- Feature name: `Add products to cart`
- Headers: 

    | Key | Value | Description |
    |----------|----------|----------|
    | Content-Type | application/json | - |
    | Authorization | Bearer Token | access token after login |
- Description : The customer takes a products to put into cart.
- `POST: {{base_url}}/api/v1/carts/:cart_id`
    ```json
    {
        "products": [
            {
                "id": 1,
                "quantity": 10
            },
            {
                "id": 2,
                "quantity": 10
            }
        ]
    }
    ```

- Feature name: `Remove items from cart`
- Headers: 

    | Key | Value | Description |
    |----------|----------|----------|
    | Content-Type | application/json | - |
    | Authorization | Bearer Token | access token after login |
- Description : The customer can remove product from cart.
- `POST: {{base_url}}/api/v1/carts/:cart_id`
    ```json
    {
        "cart_id": 1,
        "product_ids": [
            2
        ]
    }
    ```
- Feature name: `Checkout item`
- Headers: 

    | Key | Value | Description |
    |----------|----------|----------|
    | Content-Type | application/json | - |
    | Authorization | Bearer Token | access token after login |
- Description : The customer can checkout cart .
- `POST: {{base_url}}/api/v1/carts/:cart_id`
    ```json
    {
        "cart_id": 13,
        "payment_method": "transfer"
    }
    ```
