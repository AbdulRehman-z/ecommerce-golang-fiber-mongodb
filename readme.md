# E-commerce implemented with fiber(v2)



 This is e-commerce backend implemented with fiber(v2) and mongodb. It is a simple e-commerce backend with basic features.

## Environment Variables

Make sure to set the following environment variables:

- `JWT_SECRET` - Used in utils.CreateToken.
- `ADMIN_PASS` - Used in controllers.Signup.
- `ADMIN_EMAIL` - Used in controllers.Signup.
- `MONGO_URI` - Used in database.ConnectWithMongodb.

### Features

- [x] User jwt authentication with argon2 password hashing
- [x] User roles (admin, user)
- [x] User profile
- [x] User cart
- [x] User orders
- [x] Product CRUD
- [x] Admin routes

## Routes

#### Auth Routes

- [x] POST /api/v1/auth/signup (signup)
- [x] POST /api/v1/auth/signin (signin)
- [x] POST /api/v1/auth/signout (signout)
- [x] POST /api/v1/auth/profile (get user profile)

#### Product Routes

- [x] GET /api/v1/products (get all products)
- [x] GET /api/v1/products/:id (get product by id)
- [x] POST /api/v1/products/create (create product) -- admin only
- [x] PUT /api/v1/products/:id (update product) -- admin only
- [x] DELETE /api/v1/products/:id (delete product) -- admin only

#### Cart Routes

- [x] POST /api/v1/cart/remove/:id (remove product from cart)
- [x] POST /api/v1/cart/add/:id (add product to cart)

#### Order Routes

- [x] POST /api/v1/orders (get all orders)
- [x] POST /api/v1/orders/create (create order)

### Admin Routes

- [x] GET /api/v1/admin/getUsers (get all users) -- admin only
- [x] GET /api/v1/admin/getUser/:id (get user by id) -- admin only
- [x] GET /api/v1/admin/deleteUsers (delete all users) -- admin only
- [x] GET /api/v1/admin/delete/:id (delete user by id) -- admin only

### How to work with this repo

- [x] Clone the repo
- [x] Run `go mod tidy` to install all the dependencies
- [x] Run `go run .` to start the server

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
