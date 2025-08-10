# ​ E-Commerce Backend (Go)

Backend for a sample e-commerce application written in **Go**, featuring user registration and authentication, products management, cart, and orders workflow.

---

##  Description

This is a learning pet-project developed to practice building REST APIs with Go, PostgreSQL, and Gorilla Mux. It emphasizes clean architecture, modular design, and readiness for extension to real-world use.

---

##  Features

- User registration and login (authentication)
- CRUD operations for products
- Shopping cart management
- Order creation and viewing
- PostgreSQL database backend

---

##  Tech Stack

- **Go** v1.24.2  
- **Gorilla Mux** for routing  
- **REST API architecture**  
- **PostgreSQL** database  

---

##  Setup & Run

### Clone the repository
```bash
git clone https://github.com/yanpavel/ecom.git
cd ecom

---

##  Project Structure
├── cmd/ — entry point of the application (main package)
├── config/ — configuration handling (env, config files)
├── db/ — database initialization and migrations
├── service/ — core business logic
├── types/ — shared domain types and data models
├── utils/ — utility functions (helpers, middleware)
└── Makefile — convenient task runner (e.g. make run, make migrate)

---

##  API Endpoints Overview

| Method | Endpoint           | Description                     |
|--------|--------------------|---------------------------------|
| POST   | `/register`        | Register new users (returns JWT) |
| POST   | `/login` *(if exists)* | Authenticate and issue token |
| GET    | `/products`        | Retrieve all available products |
| POST   | `/products` *(if exists)* | Create a new product |
| PUT    | `/products/{id}`   | Update product information       |
| DELETE | `/products/{id}`   | Delete a product                |
| GET    | `/cart`            | View shopping cart contents     |
| POST   | `/cart/add`        | Add item to cart                |
| POST   | `/cart/remove`     | Remove item from cart           |
| POST   | `/orders`          | Create a new order              |
| GET    | `/orders/{id}`     | Get order details               |

---

## 📡 Server Address

The server is available at: 
http://localhost:8080

---

## 📄 API Request Examples

### 📝 User Registration
```bash
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{
        "username": "john",
        "password": "123456"
      }'

### 📦  Get product list
  curl -X GET http://localhost:8080//api/v1/products \

---

## 📌 Roadmap
- Add Docker Compose for running with a database
- Add Swagger for API documentation
- Write unit tests
- Implement payment system integration
