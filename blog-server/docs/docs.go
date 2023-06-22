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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/adverts": {
            "get": {
                "description": "广告列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "广告列表",
                "parameters": [
                    {
                        "type": "string",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-models_AdvertModel"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "创建广告",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "创建广告",
                "parameters": [
                    {
                        "description": "表示多个参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/advert_api.AdvertRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "批量删除广告",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "批量删除广告",
                "parameters": [
                    {
                        "description": "广告id列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RemoveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/adverts/:id": {
            "put": {
                "description": "更新广告",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "更新广告",
                "parameters": [
                    {
                        "description": "广告的一些参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/advert_api.AdvertRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/image_names": {
            "get": {
                "description": "图片名称列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "图片名称列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/images_api.ImageResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/images": {
            "get": {
                "description": "图片列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "图片列表",
                "parameters": [
                    {
                        "type": "string",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-models_BannerModel"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "图片更新",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "图片更新",
                "parameters": [
                    {
                        "description": "图片修改内容",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/images_api.ImageUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "图片上传",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "图片上传",
                "parameters": [
                    {
                        "description": "本应使用表单多图片文件参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/image_ser.FileUploadResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "批量删除图片",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "批量删除图片",
                "parameters": [
                    {
                        "description": "图片id列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RemoveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/menu_names": {
            "get": {
                "description": "菜单名称列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "菜单名称列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/menu_api.MenuNameResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/menus": {
            "get": {
                "description": "菜单列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "菜单列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/menu_api.MenuResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "添加菜单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "添加菜单",
                "parameters": [
                    {
                        "description": "添加菜单",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/menu_api.menuRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "菜单删除",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "菜单删除",
                "parameters": [
                    {
                        "description": "删除菜单",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RemoveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/menus/:id": {
            "get": {
                "description": "菜单详情",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "菜单详情",
                "parameters": [
                    {
                        "description": "菜单详情",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/menu_api.MenuResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "菜单更新",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "菜单更新",
                "parameters": [
                    {
                        "description": "菜单更新",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/menu_api.menuRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/settings/:name": {
            "get": {
                "description": "配置信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "系统配置"
                ],
                "summary": "配置信息",
                "parameters": [
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "数据类型应为对应数据的结构体",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "修改配置信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "系统配置"
                ],
                "summary": "修改配置信息",
                "parameters": [
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "数据类型应为对应数据的结构体",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "advert_api.AdvertRequest": {
            "type": "object",
            "required": [
                "href",
                "images",
                "title"
            ],
            "properties": {
                "href": {
                    "description": "跳转链接",
                    "type": "string"
                },
                "images": {
                    "description": "图片",
                    "type": "string"
                },
                "is_show": {
                    "description": "是否提示",
                    "type": "boolean"
                },
                "title": {
                    "description": "显示的标题",
                    "type": "string"
                }
            }
        },
        "ctype.ImageType": {
            "type": "integer",
            "enum": [
                1,
                2
            ],
            "x-enum-comments": {
                "Local": "本地",
                "QiNiu": "七牛云"
            },
            "x-enum-varnames": [
                "Local",
                "QiNiu"
            ]
        },
        "image_ser.FileUploadResponse": {
            "type": "object",
            "properties": {
                "file_name": {
                    "type": "string"
                },
                "is_success": {
                    "description": "图片是否上传成功",
                    "type": "boolean"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "images_api.ImageResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "images_api.ImageUpdateRequest": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "menu_api.Banner": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "menu_api.ImageSort": {
            "type": "object",
            "properties": {
                "image_id": {
                    "type": "integer"
                },
                "sort": {
                    "type": "integer"
                }
            }
        },
        "menu_api.MenuNameResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "menu_api.MenuResponse": {
            "type": "object",
            "properties": {
                "abstract": {
                    "description": "简介",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "abstract_time": {
                    "description": "简介的切换时间",
                    "type": "integer"
                },
                "banner_time": {
                    "description": "菜单图片的切换时间，为0表示不切换",
                    "type": "integer"
                },
                "banners": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/menu_api.Banner"
                    }
                },
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "主键id",
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "slogan": {
                    "type": "string"
                },
                "sort": {
                    "description": "菜单的顺序",
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "update_at": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "menu_api.menuRequest": {
            "type": "object",
            "required": [
                "path",
                "sort",
                "title"
            ],
            "properties": {
                "abstract": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "abstract_time": {
                    "description": "切换的时间",
                    "type": "integer"
                },
                "banner_time": {
                    "description": "切换的时间",
                    "type": "integer"
                },
                "image_sort_list": {
                    "description": "具体图片的顺序",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/menu_api.ImageSort"
                    }
                },
                "path": {
                    "type": "string"
                },
                "slogan": {
                    "type": "string"
                },
                "sort": {
                    "description": "菜单的序号",
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.AdvertModel": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "href": {
                    "description": "跳转链接",
                    "type": "string"
                },
                "id": {
                    "description": "主键id",
                    "type": "integer"
                },
                "images": {
                    "description": "图片",
                    "type": "string"
                },
                "is_show": {
                    "description": "是否提示",
                    "type": "boolean"
                },
                "title": {
                    "description": "显示的标题",
                    "type": "string"
                },
                "update_at": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "models.BannerModel": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "hash": {
                    "description": "图片的hash值，用于判断重复图片",
                    "type": "string"
                },
                "id": {
                    "description": "主键id",
                    "type": "integer"
                },
                "image_type": {
                    "description": "图片类型，本地还是七牛",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ctype.ImageType"
                        }
                    ]
                },
                "name": {
                    "description": "图片名称",
                    "type": "string"
                },
                "path": {
                    "description": "图片路径",
                    "type": "string"
                },
                "update_at": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "models.RemoveRequest": {
            "type": "object",
            "properties": {
                "id_list": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "res.ListResponse-models_AdvertModel": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "$ref": "#/definitions/models.AdvertModel"
                }
            }
        },
        "res.ListResponse-models_BannerModel": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "$ref": "#/definitions/models.BannerModel"
                }
            }
        },
        "res.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.01:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "gvb_server API文档",
	Description:      "API文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}