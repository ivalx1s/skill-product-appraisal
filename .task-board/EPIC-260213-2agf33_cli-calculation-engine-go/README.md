# CLI Calculation Engine (Go)

## Description
Offload all calculation methodology from the skill into a CLI tool written in Go. The skill becomes lightweight (workflow + references), the CLI handles all computations: BVR calculation, tier gap analysis, scorecard scoring, price-value mapping, dead weight ratios, cannibalization modeling, etc.

Architecture:
- Go CLI tool with agent-facing API (DSL for structured queries, like task-board)
- Use go-testing-tools skill patterns (tuitestkit library from remote)
- Use agent-facing-api skill pattern (mini-query DSL + scoped grep)
- CLI reads evaluation data (JSON/YAML), runs calculations, outputs structured results
- Agents call CLI instead of doing math in-context (token-efficient, deterministic)

Benefits:
- Deterministic calculations (no LLM math errors)
- Token-efficient (agent sends data, gets results — no intermediate reasoning)
- Testable (Go tests for every formula)
- Reusable across projects
- Skill stays lean (<500 lines) — just workflow and references

## Scope
Go CLI tool with: BVR calculator, tier gap analyzer, scorecard scorer, price-value mapper, dead weight calculator, cannibalization model, competitive positioning. Agent-facing DSL API. Tests via tuitestkit.

## Acceptance Criteria
1. CLI builds and runs. 2. All formulas from METHODOLOGY.md implemented. 3. Agent-facing DSL API works. 4. Tests pass. 5. Skill updated to reference CLI for calculations.
