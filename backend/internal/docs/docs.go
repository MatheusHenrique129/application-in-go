// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/v1/user": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create a new user.",
                "parameters": [
                    {
                        "description": "User Information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/v1/user/{user_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Find user by ID.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the user",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update a new user.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the user",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User Information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete user by ID.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the user",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.CreateUser": {
            "type": "object",
            "required": [
                "address",
                "birth_date",
                "cpf",
                "email",
                "gender",
                "name",
                "password",
                "phone_number"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "birth_date": {
                    "type": "string"
                },
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "gender": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 70,
                    "minLength": 2,
                    "example": "john"
                },
                "password": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 5
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "domain.UpdateUser": {
            "type": "object",
            "required": [
                "address",
                "birth_date",
                "cpf",
                "email",
                "gender",
                "name",
                "password",
                "phone_number"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "birth_date": {
                    "type": "string"
                },
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "gender": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 70,
                    "minLength": 2,
                    "example": "john"
                },
                "password": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 5
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birth_date": {
                    "type": "string"
                },
                "cpf": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
