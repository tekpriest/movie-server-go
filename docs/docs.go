// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://#",
        "contact": {},
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/movies": {
            "get": {
                "tags": [
                    "Movie"
                ],
                "summary": "Lists all the movies",
                "operationId": "ListMovies",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/SuccessReponseForArrayOfMovies"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "Movie"
                ],
                "summary": "Create New Movie",
                "operationId": "CreateMovie",
                "parameters": [
                    {
                        "description": "New Movie",
                        "name": "Movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Movie"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/SuccessResponseForMovie"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    }
                }
            }
        },
        "/movies/{id}": {
            "get": {
                "tags": [
                    "Movie"
                ],
                "summary": "Get Movie by ID",
                "operationId": "GetMovieByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/SuccessResponseForMovie"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    }
                }
            },
            "put": {
                "tags": [
                    "Movie"
                ],
                "summary": "Update Movie Details",
                "operationId": "UpdateMovieByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Movie",
                        "name": "Movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Movie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/SuccessResponseForMovie"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Movie"
                ],
                "summary": "Delete Movie by ID",
                "operationId": "DeleteMovie",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Movie Deleted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Movie": {
            "type": "object",
            "required": [
                "genre",
                "imdbid",
                "link",
                "poster",
                "title",
                "type",
                "year"
            ],
            "properties": {
                "_id": {
                    "type": "string",
                    "readOnly": true
                },
                "createdAt": {
                    "type": "string",
                    "readOnly": true
                },
                "genre": {
                    "type": "string"
                },
                "imdbid": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "poster": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string",
                    "readOnly": true
                },
                "year": {
                    "type": "string"
                }
            }
        },
        "ResponseError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean",
                    "default": false
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "SuccessReponseForArrayOfMovies": {
            "type": "object",
            "properties": {
                "data:{movies}": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Movie"
                    }
                },
                "message": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "status": {
                    "type": "boolean",
                    "default": true
                },
                "statusCode": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "SuccessResponseForMovie": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/Movie"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean",
                    "default": true
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Movie Server API",
	Description: "Movie Server written in go",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
