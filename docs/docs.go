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
        "contact": {
            "name": "Timur Aliev",
            "url": "https://t.me/Aliev_Timur_M",
            "email": "alievtm@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/employee": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get list employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "Get Employee Lists",
                "operationId": "get-employee-lists",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.GetEmployeeListsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/queue": {
            "get": {
                "description": "get all queue list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queue"
                ],
                "summary": "Get Queue List",
                "operationId": "get-queue-list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.QueueItem"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/queue/service": {
            "get": {
                "description": "add new ticket (item queue) in the end of the queue",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queue"
                ],
                "summary": "Add New Queue Item",
                "operationId": "add-new-ticket",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.QueueItemNumber"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/responsibility": {
            "get": {
                "description": "get responsibility list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "responsibility"
                ],
                "summary": "Get Responsibility List",
                "operationId": "get-responsibility-lists",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.Responsibility"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/workstation": {
            "post": {
                "description": "get workstation data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workstation"
                ],
                "summary": "Get Workstation Data",
                "operationId": "get-workstation",
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "$ref": "#/definitions/handler.getWorkstationListsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/employee/client": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get an available client from the queue",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "Get New Client",
                "operationId": "get-new-client",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.GetNewClientResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/employee/confirmClient": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "confirms that the client has approached the workstation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "Confirm Client",
                "operationId": "confirm-client",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.confirmClientInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ConfirmClientResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/employee/endClient": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "complete the client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "End Client",
                "operationId": "end-client",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ConfirmClientResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/employee/getStatus": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get the current status of an employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "Get Employee Status",
                "operationId": "get-status-employee",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.EmployeeStatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "logout account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Logout",
                "operationId": "logout",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.logoutInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "$ref": "#/definitions/types.LogoutResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "description": "refresh AccessToken, RefreshToken, Employee, Workstation data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh",
                "operationId": "refresh",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.refreshTokenInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "$ref": "#/definitions/types.AuthorizationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "employee authorization",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign In",
                "operationId": "sign-in",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.signInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "$ref": "#/definitions/types.AuthorizationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "registration new account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign Up",
                "operationId": "registration-account",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Employee ID",
                        "schema": {
                            "$ref": "#/definitions/types.SignUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.confirmClientInput": {
            "type": "object",
            "required": [
                "numberQueue"
            ],
            "properties": {
                "numberQueue": {
                    "type": "string"
                }
            }
        },
        "handler.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.getWorkstationListsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Workstation"
                    }
                }
            }
        },
        "handler.logoutInput": {
            "type": "object",
            "required": [
                "employeeId"
            ],
            "properties": {
                "employeeId": {
                    "type": "string"
                }
            }
        },
        "handler.refreshTokenInput": {
            "type": "object",
            "required": [
                "employeeId",
                "refreshToken",
                "workstationId"
            ],
            "properties": {
                "employeeId": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                },
                "workstationId": {
                    "type": "string"
                }
            }
        },
        "handler.signInInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "types.AuthorizationResponse": {
            "type": "object",
            "required": [
                "accessToken",
                "employee",
                "refreshToken",
                "workstation"
            ],
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "employee": {
                    "$ref": "#/definitions/types.Employee"
                },
                "refreshToken": {
                    "type": "string"
                },
                "workstation": {
                    "$ref": "#/definitions/types.Workstation"
                }
            }
        },
        "types.ConfirmClientResponse": {
            "type": "object",
            "properties": {
                "number_queue": {
                    "type": "integer"
                }
            }
        },
        "types.Employee": {
            "type": "object",
            "properties": {
                "employee_id": {
                    "type": "integer"
                },
                "expires_at": {
                    "type": "integer"
                },
                "first_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "position": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                },
                "second_name": {
                    "type": "string"
                },
                "session_state": {
                    "type": "boolean"
                },
                "status": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                },
                "workstation_id": {
                    "type": "integer"
                }
            }
        },
        "types.EmployeeStatusResponse": {
            "type": "object",
            "properties": {
                "employee_status": {
                    "type": "integer"
                }
            }
        },
        "types.GetEmployeeListsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Employee"
                    }
                }
            }
        },
        "types.GetNewClientResponse": {
            "type": "object",
            "properties": {
                "employee_status": {
                    "type": "integer"
                },
                "number_queue": {
                    "type": "integer"
                },
                "number_ticket": {
                    "type": "integer"
                },
                "service_ticket": {
                    "type": "string"
                }
            }
        },
        "types.LogoutResponse": {
            "type": "object",
            "required": [
                "status_response"
            ],
            "properties": {
                "status_response": {
                    "type": "boolean"
                }
            }
        },
        "types.QueueItem": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "integer"
                },
                "Service": {
                    "type": "string"
                },
                "Status": {
                    "type": "integer"
                },
                "Time": {
                    "type": "string"
                },
                "Workstation": {
                    "type": "integer"
                }
            }
        },
        "types.QueueItemNumber": {
            "type": "object",
            "properties": {
                "TicketID": {
                    "type": "integer"
                }
            }
        },
        "types.Responsibility": {
            "type": "object",
            "properties": {
                "responsibility_id": {
                    "type": "integer"
                },
                "responsibility_name": {
                    "type": "string"
                }
            }
        },
        "types.SignUpResponse": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "types.Workstation": {
            "type": "object",
            "properties": {
                "employee_id": {
                    "type": "integer"
                },
                "workstation_id": {
                    "type": "integer"
                },
                "workstation_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Digital Queue API",
	Description:      "API Server for digital queue",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
