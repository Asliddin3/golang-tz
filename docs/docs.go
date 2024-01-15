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
        "/api/person": {
            "get": {
                "description": "this api is to get person by filter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person"
                ],
                "summary": "Get person",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "age",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "gender",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "multiSearch",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "nation",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Person"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/api/person/": {
            "post": {
                "description": "this api is create new person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person"
                ],
                "summary": "Create new person",
                "parameters": [
                    {
                        "description": "data body",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.PersonRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Person"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/api/person/{id}": {
            "get": {
                "description": "this api is to get person by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person"
                ],
                "summary": "Get person by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Person"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "this api is Update person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person"
                ],
                "summary": "Update person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "update id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "data body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PersonUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Person"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "this api is to delete person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person"
                ],
                "summary": "Delete person",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Person": {
            "type": "object",
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Patronymic": {
                    "type": "string"
                },
                "Surname": {
                    "type": "string"
                },
                "age": {
                    "type": "integer"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nation": {
                    "type": "string"
                }
            }
        },
        "models.PersonRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "models.PersonUpdateRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "gender": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nation": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "AUTH.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Golan tz documentation",
	Description:      "This is a sample server CRM server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
