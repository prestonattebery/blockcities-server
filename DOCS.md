# api docs

> list of endpoints from our rest api! :neckbeard:

## `users`

endpoints pertaining to a user.

### get user by id

**URL** : `/users/{user id}`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

### Success Response

**Code** : `200 OK`

**Content examples**

For a User with ID 48 on the local database where that User has saved an
email address and name information.

```json
{
    "Name": "devin",
    "Username": "devinROX",
    "Owned": [
        {
            "BuildingID": 15,
            "Title": "BIG BUILDING",
            "Address": "0x32be343b94f860124dc4fee278fdcbd38c102d88"
        }
    ]
}
```

### get all users

**URL** : `/users/`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

#### Success Response

**Code** : `200 OK`

**Content examples**

```json
[
    {
        "ID": 48,
        "CreatedAt": "2018-07-30T10:12:33.024834-07:00",
        "UpdatedAt": "2018-07-30T10:12:33.024834-07:00",
        "DeletedAt": null,
        "Name": "devin",
        "Username": "devo11",
        "Email": "roche@fadsf.com",
        "Buildings": null
    }
    ...
]
```

### create a new user

**URL** : `/users/`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

#### Success Response

**Code** : `200 OK`

**Content examples**

```javascript
axios.post(url, {
    "name": "devin",
    "username": "devinROX_LOL",
    "email": "roche@coolman.com"
})
```

#### notes
emails and usernames are unique so if you reuse on it wont work

## `buildings`

operations allowed throught building endpoings

### get building by id

**URL** : `/buildings/{building id}`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

### Success Response

**Code** : `200 OK`

**Content examples**

For a User with ID 48 on the local database where that User has saved an
email address and name information.

```json
{
    "ID": 15,
    "CreatedAt": "2018-07-30T10:12:33.025933-07:00",
    "UpdatedAt": "2018-07-30T10:12:33.025933-07:00",
    "DeletedAt": null,
    "Title": "BIG BUILDING",
    "Address": "0x32be343b94f860124dc4fee278fdcbd38c102d88"
}
```

### create a new building

**URL** : `/buildings/`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

#### Success Response

**Code** : `200 OK`

**Content examples**

```javascript
axios.post(url, {
	"title": "BUILDING NAME",
	"address": "ETH_HASH"
})
```

### get all buildings

**URL** : `/buildings/`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

#### Success Response

**Code** : `200 OK`

**Content examples**

```json
[
    {
        "ID": 15,
        "CreatedAt": "2018-07-30T10:12:33.025933-07:00",
        "UpdatedAt": "2018-07-30T10:12:33.025933-07:00",
        "DeletedAt": null,
        "Title": "BIG AZZ BUILDING",
        "Address": "0x32be343b94f860124dc4fee278fdcbd38c102d88"
    }
    ...
]
```