# TASK-260213-28kao6: design-cli-commands-dsl

## Description
Design CLI command structure following agent-facing-api skill pattern. Commands: 'appraise q <query>' for DSL reads, 'appraise grep <pattern>' for text search, 'appraise calc <module> <operation> --input <json>' for direct calculations. DSL operations: calc(module.function, params) { fields }, batch(calc1; calc2), schema(). Define --format flag (json/compact). Design the Go project layout per agent-facing-api architecture.

## Scope
(define task scope)

## Acceptance Criteria
(define acceptance criteria)
