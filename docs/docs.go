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
        "/dashboard": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    },
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get dashboard data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dashboard"
                ],
                "summary": "Get dashboard data",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.DashboardResponse"
                        }
                    }
                }
            }
        },
        "/images": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Upload an image",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Upload an image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Image file",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login user menggunakan email dan password, Admin: email: mail@mail.com, password: password123 dan user biasa, email:mail2@mail.com password: password123 Setelah login silahkan masukan token dengan format: Bearer \u003ctoken\u003e",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/reservations": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve all reservations in the system contoh: /reservations?status=cancel\u0026room_type=medium",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservations"
                ],
                "summary": "Get all reservations",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter by status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by room type",
                        "name": "room_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by start date (YYYY-MM-DD)",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by end date (YYYY-MM-DD)",
                        "name": "end_date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.Reservation"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Note: Untuk Booking date menggunakan format YYYY-MM-DD untuk time menggunakan format HH:MM:SS",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservations"
                ],
                "summary": "Create a new reservation",
                "parameters": [
                    {
                        "description": "Reservation details",
                        "name": "reservation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.ReservationInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/reservations/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve a reservation by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservations"
                ],
                "summary": "Get a reservation by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reservation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Reservation"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update the status of a reservation by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservations"
                ],
                "summary": "Update reservation status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reservation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Reservation status update",
                        "name": "reservation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateStatusInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request payload or missing status",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Reservation not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/rooms": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve all rooms in the system. Requires a valid Bearer token.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Get all rooms",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.responseRoom"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new room",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Create a new room",
                "parameters": [
                    {
                        "description": "Room details",
                        "name": "room",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Room"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handlers.Room"
                        }
                    }
                }
            }
        },
        "/rooms/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get a room by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Get a room by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Room ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.responseRoom"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update a room by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Update a room by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Room ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Room details",
                        "name": "room",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Room"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.responseRoom"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete a room by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Delete a room by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Room ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/rooms/{room_id}/price": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve the price of a room by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Get the price of a room by its ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Room ID",
                        "name": "room_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Room price",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "404": {
                        "description": "Room not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/snack": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get Snack",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Snack"
                ],
                "summary": "Get Snack",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create Snack",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Snack"
                ],
                "summary": "Create Snack",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/snacks/{snack_id}/price": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve the price of a snack by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "snacks"
                ],
                "summary": "Get the price of a snack by its ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Snack ID",
                        "name": "snack_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Snack price",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.response"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get user by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reservation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.response"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update user by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reservation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete user by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reservation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.DashboardResponse": {
            "type": "object",
            "properties": {
                "room_stats": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.RoomStats"
                    }
                },
                "total_reservations": {
                    "type": "integer"
                },
                "total_revenue": {
                    "type": "integer"
                },
                "total_rooms": {
                    "type": "integer"
                },
                "total_visitors": {
                    "type": "integer"
                }
            }
        },
        "handlers.Reservation": {
            "type": "object",
            "properties": {
                "booking_date": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "final_price": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "participants": {
                    "type": "integer"
                },
                "phone": {
                    "type": "string"
                },
                "room": {
                    "$ref": "#/definitions/handlers.Rooms"
                },
                "room_id": {
                    "type": "string"
                },
                "room_price": {
                    "type": "integer"
                },
                "snack": {
                    "$ref": "#/definitions/handlers.Snacks"
                },
                "snack_id": {
                    "type": "string"
                },
                "start_time": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "total_snack": {
                    "type": "integer"
                },
                "total_snack_price": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/handlers.UserShow"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "handlers.ReservationInput": {
            "type": "object",
            "properties": {
                "booking_date": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "participants": {
                    "type": "integer"
                },
                "phone": {
                    "type": "string"
                },
                "room_id": {
                    "type": "string",
                    "example": "6066f8a1-0a80-4299-86ca-99888912bbe5"
                },
                "snack_id": {
                    "type": "string",
                    "example": "b8f8cab4-9f0e-4d08-88aa-9fd465a52536"
                },
                "start_time": {
                    "type": "string"
                },
                "total_snack": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string",
                    "example": "21691490-6817-4bf4-9bf7-3bf624d210a7"
                }
            }
        },
        "handlers.Room": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "handlers.RoomStats": {
            "type": "object",
            "properties": {
                "room_name": {
                    "type": "string"
                },
                "room_revenue": {
                    "type": "integer"
                },
                "usage_percentage": {
                    "type": "number"
                }
            }
        },
        "handlers.Rooms": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "handlers.Snacks": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "package": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                }
            }
        },
        "handlers.UpdateStatusInput": {
            "type": "object",
            "required": [
                "status"
            ],
            "properties": {
                "status": {
                    "type": "string",
                    "example": "cancel"
                }
            }
        },
        "handlers.UserInput": {
            "type": "object",
            "required": [
                "confirm_password",
                "email",
                "password",
                "username"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "handlers.UserLogin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "handlers.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_id": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "handlers.UserShow": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_id": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "handlers.response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.UserResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handlers.responseRoom": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.Room"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
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
	Title:            "Swagger MeetApp By Sinau Koding API",
	Description:      "This is documentation API from Swagger. Jika ada masalah token silahkan coba pakai postman",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
