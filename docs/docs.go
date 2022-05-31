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
        "/": {
            "get": {
                "description": "do ping",
                "tags": [
                    "example"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "{\"message\":\"hello world\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/problem/detail": {
            "get": {
                "tags": [
                    "问题"
                ],
                "summary": "问题详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "问题标识",
                        "name": "identity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status\":\"success\",\"data\":{\"id\": 1, \"identity\": \"\", \"title\": \"\", \"content\": \"\", \"total_num\": 0, \"problem_categories\": []}}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "status\":\"error\",\"error\":\"错误信息\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/problems": {
            "get": {
                "tags": [
                    "问题"
                ],
                "summary": "问题列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分页，默认1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页大小，默认15",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "关键字",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "分类标识",
                        "name": "category_identity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status\":\"success\",\"data\":{\"count\": 1, \"list\": []}}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "status\":\"error\",\"error\":\"错误信息\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user/detail": {
            "get": {
                "tags": [
                    "用户"
                ],
                "summary": "用户详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户标识",
                        "name": "identity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status\":\"success\",\"data\":{}}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "status\":\"error\",\"error\":\"错误信息\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
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
