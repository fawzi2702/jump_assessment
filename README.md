# JUMP ASSESSMENT

This program, developed in Golang, provides administrative automation for Jump users.
Particularly for adding invoices and managing transactions.

**main features:**

- List users
- Create new invoices
- Add and validate a new transaction

**Bonus feature:**

- Retrieve a user by its user ID
- Create a user
- Retrieve user's invoices using its user ID

## Installation

**0. Requirements**
First of all, you must have _docker_ and _docker-compose_ installed on your machine.

**1. Environment variables**
This program need some environment variables that must be specified in _.env_ file, that's in the root of the project.
An _.env.example_ file is provided at the root of the project.
Here are the required environment variables:
| name | type | description |
|--|--|--|
| MODE | Enum(prod, dev) | Specifies the mode in which the application runs, either production (`prod`) or development (`dev`). |
| POSTGRES_DB | string | The name of the PostgreSQL database to connect to. |
| POSTGRES_USER | string | The username for authenticating with the PostgreSQL database. |
| POSTGRES_PASSWORD | string | The password for authenticating with the PostgreSQL database. |
| POSTGRES_HOST | string | The hostname or IP address of the PostgreSQL server. |
| POSTGRES_PORT | number | The port number on which the PostgreSQL server is listening. |
| API_PORT | number | The port number on which the application's API server is listening. |

## Run project

**Debug environment**

```
make dev
```

or

**Production environment**

```
make
```

## ENPOINTS

## Invoice Endpoints

### Get All Invoices _[BONUS]_

- **URL:** `/invoice/`
- **Method:** `GET`
- **Description:** Retrieves a list of all invoices.
- **Query Parameters:**
  - `user_id` (required): the User ID whose invoices you wish to retrieve.
    Example: /users?user_id=12345
- **Response:**
  - _status_: **200**
  ```json
  [
    {
        "invoice_id": integer,
        "user_id": integer,
        "status": Enum(paid, pending),
        "label": string,
        "amount": float,
        "user": {
            "id": integer,
            "first_name": string,
            "last_name": string,
            "balance": float
        }
    }
  ]
  ```
- **Error Responses**
  - _status_: **500**
  - `internal Server Error`

### Create Invoice

- **URL:** `/invoice/`
- **Method:** `POST`
- **Description:** Creates a new invoice.
- **Request Body:**
  ```json
  {
    "user_id": integer,
    "amount": float,
    "label": string
  }
  ```
- **Response:**
  - _status_: **204**
  - `No Content`
- **Error Responses**
  - _status_: **500**
  - `internal Server Error`
  - _status_: **400**
  - `bad request`

## User Endpoints

### Get All Users

- **URL:** `/users/`
- **Method:** `GET`
- **Description:** Retrieves a list of a user.
- **Response:**
  - _status_: **200**
  ```json
  [
    {
      "id": integer,
      "first_name": string,
      "last_name": string,
      "balance": float
    }
  ]
  ```
- **Error Responses**
  - _status_: **500**
  - `internal Server Error`

### Get User by ID _[BONUS]_

- **URL:** `/users/:user_id`
- **Method:** `GET`
- **Description:** Retrieves the details of a specific user by their ID.
- **URL Parameters:**
- `user_id` (required): the User ID of the user you want to retrieve.
  Example: /users/12345
- **Response:**
  - _status_: **200**
  ```json
    {
      "id": integer,
      "first_name": string,
      "last_name": string,
      "balance": float
    }
  ```
- **Error Responses**
  - _status_: **500**
  - `internal Server Error`
  - _status_: **404**
  - `Not Found`

### Create User _[BONUS]_

- **URL:** `/users/`
- **Method:** `POST`
- **Description:** Creates a new user.
- **Request Body:**
  ```json
  {
    "first_name": string,
    "last_name": string,
  }
  ```
- **Response:**
  - _status_: **204**
  - `No Content`
- **Error Responses**
  - _status_: **500**
  - `internal Server Error`
  - _status_: **400**
  - `invalid request body`

## Transaction Endpoints

### Create Transaction

- **URL:** `/transaction/`
- **Method:** `POST`
- **Description:** Creates a new transaction.
- **Request Body:**
  ```json
  {
    "invoice_id": integer,
    "amount": float,
    "reference": string
  }
  ```
- **Response:**
  - _status_: **204**
  - `No Content`
- **Error Responses**
  - _status_: **500**
  - `internal Server Error`
  - _status_: **404**
  - `Not Found`
  - _status_: **400**
  - `invalid amount`
  - _status_: **422**
  - `invoice already paid`
