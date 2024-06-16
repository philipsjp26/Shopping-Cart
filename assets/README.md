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
- Description : Create category of product
    ```json
    {
        "name": "Indomie"
    }
    ```

- Feature name: `Create product`
- `POST: {{base_url}}/api/v1/category/products`
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
- `POST: {{base_url}}/api/v1/category/products/:category_id`
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


