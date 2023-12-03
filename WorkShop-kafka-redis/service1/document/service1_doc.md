# User Service API Documentation

## Table of Contents

- [User Service API Documentation](#user-service-api-documentation)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Endpoints](#endpoints)
    - [POST /v1/user/register](#post-v1userregister)
    - [PUT /v1/user/update/:id](#put-v1userupdateid)
    - [DELETE /v1/user/delete/:id](#delete-v1userdeleteid)
  - [Status Codes](#status-codes)
  - [Event topics](#event-topics)
    - [UserReaded](#userreaded)
    - [UserCreted](#usercreted)
    - [UserUpdated](#userupdated)
    - [UserDeleted](#userdeleted)

## Introduction

"Welcome to the "User Services" API! When you apply for our services You'll get free dog breed information."

## Endpoints

### POST /v1/user/register

Create a new user.

**Request Body:**

```json
{
  "name": "John Doe",
  "email": "example@email.com"
}
```

**Response**:

``Json
{
    "message": {
        "id": 1,
        "name": "John Doe",
        "email": "example@email.com"
    }
}
```

**Response Error**:

```json
{
    "error": "Invalid request payload"
}
```

### PUT /v1/user/update/:id

update user data

**Paremeters**:

- `id` : (require) type: uint

**Request Body :**

```json
{
  "name": "John Doe",
  "email": "example@email.com"
}
```

**Response**:

```json
{
    "message": {
        "id": 1,
        "name": "John Doe",
        "email": "example@email.com"
    }
}
```

**Response Error**:

```json
{
    "error": "Invalid request payload"
}
```
```json
{
    "error": "Invalid user ID"
}
```
```json
{
    "error": {}
}
```

### DELETE /v1/user/delete/:id

delete user from database

**Paremeters**:

- `id` : (require) type: uint


**Response**:

```json
{
    "message": "Delete ID _ Successfuly"
}
```

**Response Error**:

```json
{
    "error": {}
}
```

## Status Codes

<ul>
  <li>200 : OK. Request was successful.</li>
  <li>201 : Created. Resource was successfully created.</li>
  <li>202 : Accepted response status code indicates that the request has been accepted for processing, but the processing has not been completed.</li>
  <li>400 : Bad request. The request was invalid or cannot be served.</li>
  <li>401 : Unauthorized. The request lacks valid authentication credentials.</li>
  <li>404 : No data found</li>
</ul>

## Event topics

### UserReaded

Subscribe to the `UserReaded` topic to receive information that users have requested from other services.

**Message**:

```json
{
  "user_id": "uint",
  "dog_id": "uint",
  "dog_details": {},
  "time_stamp": "time Stamp"
}
```

### UserCreted

When user is created. Then send message to `UserCreated`

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

When user is updated. Then send message to `UserUpdated`

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

When user is deleted. Then send message to `UserDeleted`

**Message**:

```json
{
  "id": 1
}
```
