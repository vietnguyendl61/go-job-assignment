# Go Job Assignment

This is a Go application for managing job assignments between customers and helpers.

## Prerequisites

Before running this application, ensure you have the following installed on your machine:

- Go (version >= 1.19)
- PostgreSQL

## Installation
* Clone repo:
```bash
git clone https://github.com/vietnguyendl61/go-job-assignment.git
```

* Navigate to service directory:
```bash
cd booking-service 
```
* Install dependencies:
```bash
go mod tidy
```

* Do the same thing with other services.
## Configuration
1. Set up your PostgreSQL database.
2. Config file `.env` in each service:
```env
PORT=4000
DB_PORT=5432
DB_HOST=localhost
DB_NAME=booking
DB_USER=postgres
DB_PASSWORD=123456
GRPC_PORT=10000
PRICING_GRPC_HOST=0.0.0.0
PRICING_GRPC_PORT=10001
SENDING_GRPC_HOST=0.0.0.0
SENDING_GRPC_PORT=10002
```
**NOTE**:
* `PRICING_GPRC_PORT` in this example is `GRPC_PORT` in `.env` of `pricing-service` and so on with other (`SENDING_GRPC_PORT`,`BOOKING_GRPC_PORT`,`USER_GRPC_PORT`)

## Usage
1. Navigate to service directory:
```bash
cd booking-service 
```
2. Run the service:
```bash
go run main.go 
```
3. Do same thing to other services.

## API Documentation
### Endpoint
### 1. Migrate
- **URL:** `/migrate`
- **Method:** `GET`
- **Description:** Migrate database in each service.
- **Response:**
  * Status Code: `200 OK`

### 2.Register
- **URL:** `/user/register`
- **Method:** `POST`
- **Description:** Create users account or helpers account.
- **Request Body:**
```json
{
    "name":"Nguyen Bao Viet",
    "user_name":"vietnguyendl",
    "password":"vietnguyen",
    "phone_number":"0382634581",
    "address":"ho chi minh city, viet nam",
    "is_helper":false
}
```
- **Response:**
    * Status Code: `201 Created`
    * Body:
```json
{
    "id": "caed8d38-17ae-42f4-81ea-9e0613f63564",
    "created_at": "2024-02-24T07:50:52.230879+07:00",
    "updated_at": "2024-02-24T07:50:52.230879+07:00",
    "name": "Nguyen Bao Viet",
    "user_name": "vietnguyendl",
    "password": "vietnguyen",
    "phone_number": "0382634581",
    "address": "ho chi minh city, viet nam",
    "is_helper": false
}
```
- **Note:**
  * `user_name` is unique

### 3.Login
- **URL:** `/user/login`
- **Method:** `POST`
- **Description:** Login by username/password to get user id.
- **Request Body:**
```json
{
    "user_name":"vietnguyendl",
    "password":"vietnguyen"
}
```
- **Response:**
    * Status Code: `200 OK`
    * Body:
```json
"caed8d38-17ae-42f4-81ea-9e0613f63564"
```
- **Note:**
  * This user id use for api `price/get-list?date=` and `/job/create`

### 4. Create job
- **URL:** `/job/create`
- **Method:** `POST`
- **Description:** Create job and price, auto send to helper available.
- **Header:**
    * Add:
```
x-user-id=caed8d38-17ae-42f4-81ea-9e0613f63564
```
- **Request Body:**
```json
{
  "book_date": "2024-04-23T11:17:05.360024+00:00",
  "description": "day la viec nha",
  "price": 100399
}
```
- **Response:**
    * Status Code: `200 Created`
    * Body:
```json
{
  "id": "0f0462d7-1408-486d-b322-5685de7765e5",
  "creator_id": "caed8d38-17ae-42f4-81ea-9e0613f63564",
  "created_at": "2024-02-24T08:03:15.522508+07:00",
  "updated_at": "2024-02-24T08:03:15.522508+07:00",
  "book_date": "2024-04-23T11:17:05.360024Z",
  "description": "day la viec nha"
}
```
- **Note:**
  * Api `/price/get-list?date=` available to login user so if no header was added then field creator_id is null and cannot get price of this job.
  * `book_date` field must have same format with example request body.
  
### 5. Get price by time
- **URL:** `/price/get-list?date=`
- **Method:** `GET`
- **Description:** Get price by time.
**Header:**
    * Add:
```
x-user-id=caed8d38-17ae-42f4-81ea-9e0613f63564
```
- **Response:**
    * Status Code: `200 OK`
    * Body:
```json
[
  {
    "id": "31c6af44-fdc4-4cc5-8799-4748f068e95a",
    "creator_id": "caed8d38-17ae-42f4-81ea-9e0613f63564",
    "created_at": "2024-02-24T08:03:15.561566+07:00",
    "updated_at": "2024-02-24T08:03:15.561566+07:00",
    "job_id": "0f0462d7-1408-486d-b322-5685de7765e5",
    "price": 100399
  }
]
```
- **Note:**
    * Param `date` must have same format with `book_date` (`2024-02-24T08:03:15.522508`)
    * This API will get all the price of job for logined user with `x-user-id` header

### 6. Get One Job
- **URL:** `/job/get-one/{id}`
- **Method:** `GET`
- **Description:** Get one job by id.
- **Response:**
    * Status Code: `200 OK`
    * Body:
```json
{
  "id": "0f0462d7-1408-486d-b322-5685de7765e5",
  "creator_id": "91703bda-9c62-40d7-be0a-939069c61ece",
  "created_at": "2024-02-24T08:03:15.522508+07:00",
  "updated_at": "2024-02-24T08:03:15.522508+07:00",
  "book_date": "2024-04-23T18:17:05.360024+07:00",
  "description": "day la viec nha"
}
```

### 7. Get one job assignment
- **URL:** `/job-assignment/get-one/{job_id}`
- **Method:** `GET`
- **Description:** Get one job assignment by job id.
- **Response:**
    * Status Code: `200 OK`
    * Body:
```json
{
  "id": "1bf4d377-1988-4190-b09f-7ac63c64fe38",
  "creator_id": "91703bda-9c62-40d7-be0a-939069c61ece",
  "created_at": "2024-02-24T08:03:15.533276+07:00",
  "updated_at": "2024-02-24T08:03:15.533276+07:00",
  "job_id": "0f0462d7-1408-486d-b322-5685de7765e5",
  "helper_id": "495d6f88-f72b-4bc6-8fb6-56f9294e17c9",
  "job_status": "Processing"
}
```
## Some Rule
- Mỗi helper chỉ nhận được 1 việc 1 ngày
  - có thể mở rộng bằng cách estimate thời gian dọn dẹp thông qua hình ảnh, video của nơi cần dọn dẹp, mô tả của công việc hoặc giảm thành 2 công việc/ ngày ( buổi sáng và buổi chiều)
- Khi có nhiều helper đang rảnh thì chọn random 1 trong những helper đó
    - có thể mở rộng bằng cách chọn những user ở gần chỗ người đăng công việc hoặc chưa nhận được việc trong 1 khoảng thời gian nào đó

