// Code generated by swaggo/swag at 2023-09-21 14:45:52.734202 -0400 EDT m=+1.000653834. DO NOT EDIT.

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
        "/device-config/eth-addr/{ethAddr}/urls": {
            "get": {
                "description": "Retrieve the URLs for PID, DeviceSettings, and DBC configuration based on device's Ethereum Address",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vehicle-signal-decoding"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ethereum Address",
                        "name": "ethAddr",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved configuration URLs",
                        "schema": {
                            "$ref": "#/definitions/internal_controllers.DeviceConfigResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found - No templates available for the given parameters"
                    }
                }
            }
        },
        "/device-config/vin/{vin}/urls": {
            "get": {
                "description": "Retrieve the URLs for PID, DeviceSettings, and DBC configuration based on a given VIN",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vehicle-signal-decoding"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "vehicle identification number (VIN)",
                        "name": "vin",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved configuration URLs",
                        "schema": {
                            "$ref": "#/definitions/internal_controllers.DeviceConfigResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found - No templates available for the given parameters"
                    }
                }
            }
        },
        "/device-config/{templateName}/dbc": {
            "get": {
                "description": "Fetches the DBC file from the dbc_files table given a template name",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "vehicle-signal-decoding"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "template name",
                        "name": "templateName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved DBC file",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "No DBC file found for the given template name."
                    }
                }
            }
        },
        "/device-config/{templateName}/deviceSettings": {
            "get": {
                "description": "Fetches the device settings configurations from device_settings table given a template name",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vehicle-signal-decoding"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "template name",
                        "name": "templateName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved Device Settings",
                        "schema": {
                            "$ref": "#/definitions/internal_controllers.DeviceSetting"
                        }
                    },
                    "404": {
                        "description": "No Device Settings data found for the given template name."
                    }
                }
            }
        },
        "/device-config/{templateName}/pids": {
            "get": {
                "description": "Retrieves a list of PID configurations from the database given a template name",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vehicle-signal-decoding"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "template name",
                        "name": "templateName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved PID Configurations",
                        "schema": {
                            "$ref": "#/definitions/github_com_DIMO-Network_vehicle-signal-decoding_pkg_grpc.PIDRequests"
                        }
                    },
                    "404": {
                        "description": "No PID Config data found for the given template name."
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_DIMO-Network_vehicle-signal-decoding_pkg_grpc.PIDConfig": {
            "type": "object",
            "properties": {
                "formula": {
                    "type": "string"
                },
                "header": {
                    "type": "integer"
                },
                "interval_seconds": {
                    "type": "integer"
                },
                "mode": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "pid": {
                    "type": "integer"
                },
                "protocol": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "github_com_DIMO-Network_vehicle-signal-decoding_pkg_grpc.PIDRequests": {
            "type": "object",
            "properties": {
                "requests": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_DIMO-Network_vehicle-signal-decoding_pkg_grpc.PIDConfig"
                    }
                },
                "template_name": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "internal_controllers.DeviceConfigResponse": {
            "type": "object",
            "properties": {
                "dbcURL": {
                    "type": "string"
                },
                "deviceSettingUrl": {
                    "type": "string"
                },
                "pidUrl": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "internal_controllers.DeviceSetting": {
            "type": "object",
            "properties": {
                "battery_critical_level_voltage": {
                    "type": "string"
                },
                "safety_cut_out_voltage": {
                    "type": "string"
                },
                "sleep_timer_event_driven_interval": {
                    "type": "string"
                },
                "sleep_timer_event_driven_period": {
                    "type": "string"
                },
                "sleep_timer_inactivity_after_sleep_interval": {
                    "type": "string"
                },
                "sleep_timer_inactivity_fallback_interval": {
                    "type": "string"
                },
                "template_name": {
                    "type": "string"
                },
                "wake_trigger_voltage_level": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "DIMO Vehicle-Signal-Decoding",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
