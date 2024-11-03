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
        "/auth/login": {
            "post": {
                "description": "Login a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login Data",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Logout",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register a user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register User",
                "parameters": [
                    {
                        "description": "Login Data",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/status": {
            "get": {
                "description": "Check login status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Auth Status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/user/{uid}": {
            "delete": {
                "description": "delete a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/users": {
            "get": {
                "description": "List users on the sytem",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "List users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/auth/users/{uid}": {
            "put": {
                "description": "update users attributes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "update user",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Health check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "status"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/hosts/by-team/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "hosts"
                ],
                "summary": "Get all hosts by team",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/hosts/by-team/{tid}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "hosts"
                ],
                "summary": "Get hosts by team",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Team ID",
                        "name": "tid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/jobs": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "get jobs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/jobs/manager": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "get jobmanager state",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/jobs/nmap/{jid}": {
            "post": {
                "description": "upload new nmap scan",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "nmap scan",
                "parameters": [
                    {
                        "description": "nmap scan data",
                        "name": "scan",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Scan"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/jobs/{jobtype}/next": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "delete team",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Job Type",
                        "name": "jobtype",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/teams": {
            "get": {
                "description": "get all teams",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "teams"
                ],
                "summary": "Get Teams",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "create a team",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "teams"
                ],
                "summary": "Create team",
                "parameters": [
                    {
                        "description": "team data",
                        "name": "create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Team"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/teams/{tid}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "teams"
                ],
                "summary": "delete team",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Team ID",
                        "name": "tid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Host": {
            "type": "object"
        },
        "models.LoginReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "models.RegisterReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.Scan": {
            "type": "object"
        },
        "models.Team": {
            "type": "object",
            "properties": {
                "hosts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Host"
                    }
                },
                "iprange": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "tid": {
                    "type": "string"
                }
            }
        },
        "models.UserReq": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
