# Bus Schedule REST Go
RESTful website built with Go to support [Bus Schedule React](https://github.com/fannyhasbi/bus-schedule-react)

---

* [Entity Relationship Diagram](#entity-relationship-diagram)
* [Pre-Installation](#pre-installation)
* [Installation](#installation)
* [Request](#request)
  * [Bus](#bus)
  * [Add Bus](#add-bus)
  * [Place](#place)
  * [Departure](#departure)
  * [Add Departure](#add-departure)
  * [Arrival](#arrival)
  * [Add Arrival](#add-arrival)

---

## Entity Relationship Diagram
Here I create my own ERD for the application

![Bus Schedule REST](./bus-schedule-erd.png "Bus Schedule REST")

## Pre-Installation
In this project, I used some packages to maximize the application. So, you must install some required packages.

```bash
go get github.com/go-sql-driver/mysql
go get github.com/gorilla/mux
```

## Installation
1. Get the project from this repository
   ```bash
   go get github.com/fannyhasbi/bus-schedule-rest-go
   ```

2. Move to the cloned project (`yourgopathfolder/src/github.com/fannyhasbi/bus-schedule-rest-go`), then run this command:
   ```bash
   go run main.go
   ```
   
   Last, access `localhost:8080` from your browser

## Request
### Bus
**GET** `/api/bus`

Endpoint to get bus data

Success response example
```json
{
  "status": 200,
  "message": "OK",
  "data": [
    {
      "id": 1,
      "nama": "Sinar Jaya"
    },
    {
      "id": 2,
      "nama": "Nusantara"
    }
  ]
}
```

### Add Bus
**POST** `/api/add-bus`

Endpoint to add new bus data

Post parameters

| Property | Type | Description |
| --- | --- | --- |
| perusahaan | String | Name of bus company |

Success response example
```json
{
  "status": 200,
  "message": "OK",
  "data": null
}
```

### Place
**GET** `/api/place`

Endpoint to get place data

Success response example
```json
{
  "status": 200,
  "message": "OK",
  "data": [
    {
      "id": 1,
      "nama": "Semarang"
    },
    {
      "id": 2,
      "nama": "Tegal"
    }
  ]
}
```

### Departure
**GET** `/api/departure`

Endpoint to get departure data

Success response example
```json
{
  "status": 200,
  "message": "OK",
  "data": [
    {
      "id": 1,
      "id_perusahaan": 1,
      "nama_perusahaan": "Sinar Jaya",
      "id_tujuan": 2,
      "nama_tujuan": "Tegal",
      "id_asal": 1,
      "nama_asal": "Semarang",
      "berangkat": "2018-12-01 12:00:00",
      "sampai": "2018-12-01 15:00:00"
    }
  ]
}
```

### Add Departure
**POST** `/api/add-departure`

Endpoint to add new departure data

Post parameters

| Property | Type | Description |
| --- | --- | --- |
| id_perusahaan | Number | Company ID |
| id_tujuan | Number | Place ID |
| id_asal | Number | Place ID |
| berangkat | String | Format: `HH:MM` |
| sampai | String | Format: `HH:MM` |

Success response example
```json
{
  "status": 200,
  "message": "OK",
  "data": null
}
```

### Arrival
**GET** `/api/arrival`

Endpoint to get arrival data

Success response example
```json
{
  "status": 200,
  "message": "OK",
  "data": [
    {
      "id": 3,
      "id_perusahaan": 1,
      "nama_perusahaan": "Sinar Jaya",
      "id_tujuan": 1,
      "nama_tujuan": "Semarang",
      "id_asal": 4,
      "nama_asal": "Pekalongan",
      "berangkat": "2018-11-29 10:00:00",
      "datang": "2018-11-29 12:00:00"
    }
  ]
}
```

### Add Arrival
**POST** `/api/add-arrival`

Endpoint to add new arrival data

Post parameters

| Property | Type | Description |
| --- | --- | --- |
| id_perusahaan | Number | Company ID |
| id_tujuan | Number | Place ID |
| id_asal | Number | Place ID |
| berangkat | String | Format: `HH:MM` |
| datang | String | Format: `HH:MM` |

Success response example
```json
{
  "status": 200,
  "message": "OK",
  "data": null
}
```