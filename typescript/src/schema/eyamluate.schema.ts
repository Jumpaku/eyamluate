export default {
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$ref": "#/definitions/Expr",
  "definitions": {
    "Expr": {
      "anyOf": [
        {
          "$ref": "#/definitions/Eval"
        },
        {
          "$ref": "#/definitions/Scalar"
        },
        {
          "$ref": "#/definitions/Obj"
        },
        {
          "$ref": "#/definitions/Arr"
        },
        {
          "$ref": "#/definitions/Json"
        },
        {
          "$ref": "#/definitions/RangeIter"
        },
        {
          "$ref": "#/definitions/GetElem"
        },
        {
          "$ref": "#/definitions/FunCall"
        },
        {
          "$ref": "#/definitions/Cases"
        },
        {
          "$ref": "#/definitions/OpUnary"
        },
        {
          "$ref": "#/definitions/OpBinary"
        },
        {
          "$ref": "#/definitions/OpVariadic"
        }
      ]
    },
    "Eval": {
      "type": "object",
      "properties": {
        "where": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/FunDef"
          }
        },
        "eval": {
          "$ref": "#/definitions/Expr"
        }
      },
      "required": [
        "eval"
      ],
      "additionalProperties": false
    },
    "FunDef": {
      "type": "object",
      "properties": {
        "def": {
          "type": "string"
        },
        "value": {
          "$ref": "#/definitions/Expr"
        },
        "with": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      },
      "required": [
        "def",
        "value"
      ],
      "additionalProperties": false
    },
    "Scalar": {
      "oneOf": [
        {
          "type": "number"
        },
        {
          "type": "boolean"
        },
        {
          "type": "string"
        }
      ]
    },
    "Obj": {
      "type": "object",
      "properties": {
        "obj": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/Expr"
          }
        }
      },
      "required": [
        "obj"
      ],
      "additionalProperties": false
    },
    "Arr": {
      "type": "object",
      "properties": {
        "arr": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Expr"
          }
        }
      },
      "required": [
        "arr"
      ],
      "additionalProperties": false
    },
    "Json": {
      "type": "object",
      "properties": {
        "json": {
          "$ref": "#/definitions/Json/definitions/NonNull"
        }
      },
      "required": [
        "json"
      ],
      "additionalProperties": false,
      "definitions": {
        "NonNull": {
          "oneOf": [
            {
              "type": "number"
            },
            {
              "type": "boolean"
            },
            {
              "type": "string"
            },
            {
              "type": "object",
              "additionalProperties": {
                "$ref": "#/definitions/Json/definitions/NonNull"
              }
            },
            {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Json/definitions/NonNull"
              }
            }
          ]
        }
      }
    },
    "RangeIter": {
      "type": "object",
      "properties": {
        "for": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "minItems": 2,
          "maxItems": 2
        },
        "in": {
          "$ref": "#/definitions/Expr"
        },
        "do": {
          "$ref": "#/definitions/Expr"
        },
        "if": {
          "$ref": "#/definitions/Expr"
        }
      },
      "required": [
        "for",
        "in",
        "do"
      ],
      "additionalProperties": false
    },
    "GetElem": {
      "type": "object",
      "properties": {
        "get": {
          "$ref": "#/definitions/Expr"
        },
        "from": {
          "$ref": "#/definitions/Expr"
        }
      },
      "required": [
        "get",
        "from"
      ],
      "additionalProperties": false
    },
    "FunCall": {
      "type": "object",
      "properties": {
        "ref": {
          "type": "string"
        },
        "with": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/Expr"
          }
        }
      },
      "required": [
        "ref"
      ],
      "additionalProperties": false
    },
    "Cases": {
      "type": "object",
      "properties": {
        "cases": {
          "type": "array",
          "items": {
            "oneOf": [
              {
                "$ref": "#/definitions/Cases/definitions/CasesWhenThen"
              },
              {
                "$ref": "#/definitions/Cases/definitions/CasesOtherwise"
              }
            ]
          },
          "contains": {
            "$ref": "#/definitions/Cases/definitions/CasesOtherwise"
          }
        }
      },
      "required": [
        "cases"
      ],
      "additionalProperties": false,
      "definitions": {
        "CasesWhenThen": {
          "type": "object",
          "properties": {
            "when": {
              "$ref": "#/definitions/Expr"
            },
            "then": {
              "$ref": "#/definitions/Expr"
            }
          },
          "required": [
            "when",
            "then"
          ],
          "additionalProperties": false
        },
        "CasesOtherwise": {
          "type": "object",
          "properties": {
            "otherwise": {
              "$ref": "#/definitions/Expr"
            }
          },
          "required": [
            "otherwise"
          ],
          "additionalProperties": false
        }
      }
    },
    "OpUnary": {
      "type": "object",
      "minProperties": 1,
      "maxProperties": 1,
      "additionalProperties": false,
      "properties": {
        "len": {
          "$ref": "#/definitions/Expr"
        },
        "not": {
          "$ref": "#/definitions/Expr"
        },
        "flat": {
          "$ref": "#/definitions/Expr"
        },
        "floor": {
          "$ref": "#/definitions/Expr"
        },
        "ceil": {
          "$ref": "#/definitions/Expr"
        },
        "abort": {
          "$ref": "#/definitions/Expr"
        }
      }
    },
    "OpBinary": {
      "type": "object",
      "minProperties": 1,
      "maxProperties": 1,
      "additionalProperties": false,
      "properties": {
        "sub": {
          "$ref": "#/definitions/OpBinary/definitions/OpBinaryOperand"
        },
        "div": {
          "$ref": "#/definitions/OpBinary/definitions/OpBinaryOperand"
        },
        "eq": {
          "$ref": "#/definitions/OpBinary/definitions/OpBinaryOperand"
        },
        "neq": {
          "$ref": "#/definitions/OpBinary/definitions/OpBinaryOperand"
        },
        "lt": {
          "$ref": "#/definitions/OpBinary/definitions/OpBinaryOperand"
        },
        "lte": {
          "$ref": "#/definitions/OpBinary/definitions/OpBinaryOperand"
        },
        "gt": {
          "$ref": "#/definitions/OpBinary/definitions/OpBinaryOperand"
        },
        "gte": {
          "$ref": "#/definitions/OpBinary/definitions/OpBinaryOperand"
        }
      },
      "definitions": {
        "OpBinaryOperand": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Expr"
          },
          "minItems": 2,
          "maxItems": 2
        }
      }
    },
    "OpVariadic": {
      "type": "object",
      "minProperties": 1,
      "maxProperties": 1,
      "additionalProperties": false,
      "properties": {
        "add": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Expr"
          }
        },
        "mul": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Expr"
          }
        },
        "and": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Expr"
          }
        },
        "or": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Expr"
          }
        },
        "cat": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Expr"
          }
        },
        "min": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Expr"
          },
          "minItems": 1
        },
        "max": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Expr"
          },
          "minItems": 1
        },
        "merge": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Expr"
          }
        }
      }
    }
  }
}