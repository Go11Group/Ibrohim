// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/user": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Adds a new user",
                "tags": [
                    "admin"
                ],
                "summary": "Adds user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.NewUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.NewUserResp"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/user/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves user info by ID",
                "tags": [
                    "admin"
                ],
                "summary": "Gets user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.UserInfo"
                        }
                    },
                    "400": {
                        "description": "Invalid user id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates user info by ID",
                "tags": [
                    "admin"
                ],
                "summary": "Updates user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New user data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.NewDataNoID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.NewDataResp"
                        }
                    },
                    "400": {
                        "description": "Invalid user id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Deletes user info by ID",
                "tags": [
                    "admin"
                ],
                "summary": "Deletes user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid user id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves users from the database by filtering",
                "tags": [
                    "admin"
                ],
                "summary": "Fetches users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Full name",
                        "name": "full_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Location",
                        "name": "location",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Role",
                        "name": "role",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of users per page",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.Users"
                        }
                    },
                    "400": {
                        "description": "Invalid pagination parameters",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves user profile by ID",
                "tags": [
                    "user"
                ],
                "summary": "Gets user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.Profile"
                        }
                    },
                    "400": {
                        "description": "Invalid user id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates user info by ID",
                "tags": [
                    "user"
                ],
                "summary": "Updates user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New user data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.NewData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.UpdateResp"
                        }
                    },
                    "400": {
                        "description": "Invalid user id or data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Deletes user info by ID",
                "tags": [
                    "user"
                ],
                "summary": "Deletes user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid user id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "admin.NewDataNoID": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "image": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "postal_code": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.NewDataResp": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.NewUser": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "postal_code": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.NewUserResp": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.UserDetails": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.UserInfo": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "postal_code": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.Users": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/admin.UserDetails"
                    }
                }
            }
        },
        "user.NewData": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "image": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "postal_code": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.Profile": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "phone_number": {
                    "type": "string"
                },
                "postal_code": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.UpdateResp": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "phone_number": {
                    "type": "string"
                },
                "postal_code": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/e-commerce",
	Schemes:          []string{"http"},
	Title:            "E-Commerce",
	Description:      "API Gateway of E-Commerce",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
