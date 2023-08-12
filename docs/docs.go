// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/args/{arg}": {
            "get": {
                "description": "返回路劲参数",
                "tags": [
                    "fiber-layout"
                ],
                "summary": "path arg",
                "parameters": [
                    {
                        "type": "string",
                        "description": "arg",
                        "name": "arg",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/bind/{arg}": {
            "post": {
                "description": "绑定uri和body参数，并进行参数校验",
                "tags": [
                    "fiber-layout"
                ],
                "summary": "bind json test",
                "parameters": [
                    {
                        "type": "string",
                        "description": "arg",
                        "name": "arg",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/body": {
            "post": {
                "description": "直接返回请求body内容",
                "tags": [
                    "fiber-layout"
                ],
                "summary": "test post body",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/error": {
            "get": {
                "description": "handler返回错误，打印access日志",
                "tags": [
                    "fiber-layout"
                ],
                "summary": "test error",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/error/wrap": {
            "get": {
                "description": "打印wrap error测试",
                "tags": [
                    "fiber-layout"
                ],
                "summary": "path arg",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/hello": {
            "get": {
                "description": "返回hello world，添加自定义响应头",
                "tags": [
                    "fiber-layout"
                ],
                "summary": "say hello",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/panic": {
            "get": {
                "description": "空指针panic，测试是否能recover，并返回响应",
                "tags": [
                    "fiber-layout"
                ],
                "summary": "test panic",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "fiber-layout 接口文档",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
