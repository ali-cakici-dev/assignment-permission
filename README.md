<details>
  <summary>Assignment</summary>

## Acceptance criteria

The objective of this project is to develop a permission management system that connects to a mongoDB database, retrieves permitted resources for a user, determines the user's role for a specific resource, and checks whether the user has permission to perform a particular action on that resource.

Please fork the project repository on GitHub, make the necessary changes as outlined in the assignment, commit your modifications with clear messages, and submit the URL of your forked repository along with any additional documentation or instructions.

## Objective

In this task, you are required to fill in the empty functions with appropriate names to achieve the desired functionality. These functions will play a crucial role in connecting to the database, fetching permitted resources for users, determining user roles for specific resources, and performing authorization checks.

## Empty Functions:

"initDB": This function is responsible for establishing a connection to the database system. It should configure the necessary parameters and handle any errors that occur during the connection process.

"fetchDocument": This function is responsible for fetching a specific document from the database.

"fetchPermittedResources": Implement this function to retrieve permitted resources for a given user from the database. It should query the database based on the user's credentials and return a list of resources that the user is allowed to access.

"getRole": This function will determine the role of the user for a specific resource. The function should take the user's identifier and the resource ID as input parameters and return the corresponding role assigned to the user for that resource.

"checkPermission": This function is responsible for performing authorization checks to determine whether a user has permission to perform a specified action on a resource.

## Example Documents

### Role

```json
{
    "_id":"admin",
    "name":"Admin",
    "description":"Admin privilages",
    "action":["read","write","update","delete"]
}
```

### Permission

```json
{
    "resource":"some_resource_id",
    "users":[
        {
            "user_id":"user_1_id",
            "role":"admin"
        },
        {
            "user_id":"user_2_id",
            "role":"readonly"
        }
    ],
    "groups":[
        {
            "group_id":"group_1_id",
            "role":"admin",
            "members":["user_2_id","user_3_id"]
        },
    ]
}
```

NOTE: In the event of a permission conflict, precedence shall be given to the higher permission level for resolution.
</details>


How to run:
```bash
make up
```

Endpoints:

| /all-permissions                     | GET  | Retrieves all stored permissions                            |
|--------------------------------------|------|-------------------------------------------------------------|
| /permissions                         | POST | Creates a new permission entry                              |
| /add-role                            | POST | Creates a new role                                          |
| /fetch-permitted-resources/{user_id} | GET  | Fetches the list of resources a user is permitted to access |
| /get-role/{group_id}/{user_id}       | GET  | Gets the role of a user for a specific resource             |