# Dog Service Documentation

## Table of Contents

- [Drinks Service Documentation](#drinks-service-documentation)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Event topics](#event-topics)
    - [UserCreated](#usercreated)
    - [UserUpdated](#userupdated)
    - [UserDeleted](#userdeleted)
    - [UserReaded](#userreaded)
  - [Endpoints](#endpoints)
    - [GET /v1/dogs](#get-v1dogs)
    - [GET /v1/dog/:userId/:dogId](#get-v1doguserIddogId)
  - [Status Codes](#status-codes)

## Introduction

"Welcome to the Dog Breed Information Service! This service utilizes text from Kafka and responds with dog breed data, managing events for each topic, storing messages in a database, and fetching information from https://api.thedogapi.com/v1/breeds."

## Event topics

### UserCreated

When the user makes an event `UserCreated`, it will create a user in the database if the user does not exist.

**Message**:

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "example@email.com",
  "time_stamp": "Time Stamp"
}
```

### UserUpdated

When the user makes an event `UserUpdated`, it will update user data

**Message**:

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "example@email.com",
  "time_stamp": "Time Stamp"
}
```

### UserDeleted

Deleted user from data base.

**Message**:

```json
{
  "id": 1
}
```

### UserReaded

When the user receives dog breed information with `dogId` it will send message to topic `UserReaded`

**Message**:

```json
{
  "user_id": "uint",
  "item_id": "uint",
  "item_details": {},
  "time_stamp": "time Stamp"
}
```

## Endpoints

### GET /v1/dogs

get dog breeds.

**Response**:

- Code: `200`

```json
{
   "message": "Succeed",
    "status": "OK",
    "status_code": 200,
    "data": [
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "id": 1,
            "name": "Affenpinscher",
            "temperament": "Stubborn, Curious, Playful, Adventurous, Active, Fun-loving",
            "life_span": "10 - 12 years",
            "origin": "Germany, France",
            "weight_imperial": "6 - 13",
            "weight_metric": "3 - 6",
            "height_imperial": "9 - 11.5",
            "height_metric": "23 - 29",
            "bred_for": "Small rodent hunting, lapdog",
            "breed_group": "Toy",
            "reference_image_id": "BJa4kxc4X"
        },{.........}
        ]
}
```

### GET /v1/dog/:userId/:dogId (UserReaded)

User get dog breed information with dog id. and sent data to topic `UserReaded`

**Parameters**:

- `dogId`: (require) type uint.
- `userId` : (require) type uint.

**Response**:

```json
{
  "message": "Succeed",
    "status": "OK",
    "status_code": 200,
    "data": {
        "ID": 0,
        "CreatedAt": "2023-12-03T10:02:12.45842Z",
        "UpdatedAt": "2023-12-03T10:02:12.45842Z",
        "DeletedAt": null,
        "id": 1,
        "name": "Affenpinscher",
        "temperament": "Stubborn, Curious, Playful, Adventurous, Active, Fun-loving",
        "life_span": "10 - 12 years",
        "origin": "Germany, France",
        "weight_imperial": "6 - 13",
        "weight_metric": "3 - 6",
        "height_imperial": "9 - 11.5",
        "height_metric": "23 - 29",
        "bred_for": "Small rodent hunting, lapdog",
        "breed_group": "Toy",
        "reference_image_id": "BJa4kxc4X"
    }
}
```

**Response Error**:

```json
{
  "message": "No data found",
  "status": "Not Found",
  "status_code": 404
}
```

```json
{
  "message": "Invalid parameter type",
  "status": "Bad Request",
  "status_code": 400
}
```

## Status Codes

<ul>
  <li>200 : OK. Request was successful.</li>
  <li>201 : Created. Resource was successfully created.</li>
  <li>202 : Accepted response status code indicates that the request has been accepted for processing, but the processing has not been completed.</li>
  <li>400 : Bad request. The request was invalid or cannot be served.</li>
 authentication credentials.</li>
  <li>404 : No data found</li>
</ul>
