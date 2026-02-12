# TASK-260213-3ubk2z: implement-dsl-parser

## Description
Go package: internal/query/parser.go. Implement tokenizer + recursive descent parser for DSL syntax: 'calc(module.function, params) { fields }'. Support batch queries via semicolons. Parse operation name, extract params as key=value pairs, extract field projection. Follow agent-facing-api assets/dsl-parser.go as reference. Grammar: query = operation '(' params ')' ['{' fields '}'], batch = query (';' query)*.

## Scope
(define task scope)

## Acceptance Criteria
(define acceptance criteria)
