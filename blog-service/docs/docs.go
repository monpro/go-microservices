// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-11-13 21:02:13.168918 +1100 AEDT m=+0.065462601

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/go-microservices",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/articles": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get a list of articles",
                "parameters": [
                    {
                        "type": "string",
                        "description": "articleName",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "tagId",
                        "name": "tag_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "state",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page_size",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.ArticleSwagger"
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/articles/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get an article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id for the article",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "get a list of articles",
                "parameters": [
                    {
                        "description": "tag_id",
                        "name": "tag_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "title",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "description for the article",
                        "name": "desc",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "cover_image_url",
                        "name": "cover_image_url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "content",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "modifier",
                        "name": "modified_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.ArticleSwagger"
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "get a list of articles",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id for the article",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/string"
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/tags": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get a list of tags",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "tagName",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "status",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page_size",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.TagSwapper"
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "create a tag",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "tagName",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "status",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "creator",
                        "name": "created_by",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.Tag"
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/tags/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "update  atag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id for tag",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "tag name",
                        "name": "name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "state",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "modifier",
                        "name": "modified_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.Tag"
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "delete a tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id for tag",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/string"
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Pager": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total_rows": {
                    "type": "integer"
                }
            }
        },
        "errcode.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "details": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.Article": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "cover_image_url": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "created_on": {
                    "type": "integer"
                },
                "deleted_on": {
                    "type": "integer"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_del": {
                    "type": "integer"
                },
                "modified_by": {
                    "type": "string"
                },
                "state": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ArticleSwagger": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Article"
                    }
                },
                "pager": {
                    "type": "object",
                    "$ref": "#/definitions/app.Pager"
                }
            }
        },
        "model.Tag": {
            "type": "object",
            "properties": {
                "created_by": {
                    "type": "string"
                },
                "created_on": {
                    "type": "integer"
                },
                "deleted_on": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_del": {
                    "type": "integer"
                },
                "modified_by": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "integer"
                }
            }
        },
        "model.TagSwapper": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Tag"
                    }
                },
                "pager": {
                    "type": "object",
                    "$ref": "#/definitions/app.Pager"
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
	BasePath:    "",
	Schemes:     []string{},
	Title:       "blog microservice",
	Description: "Go microservices",
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
