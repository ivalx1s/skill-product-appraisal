# Agent-Facing DSL

## Description
Implement mini-query DSL for agents to call calculations token-efficiently. Follow agent-facing-api skill pattern (structured reads + scoped grep). Agents send evaluation data, get structured results back without loading calculation logic into context.

## Scope
(define story scope)

## Acceptance Criteria
1. DSL syntax works. 2. Batch queries supported. 3. Compact output format. 4. Schema introspection (schema() command).
