# Plan: CLI Calculation Engine (Go)

**Date:** 2026-02-13 01:00
**Epic:** EPIC-260213-2agf33 — CLI Calculation Engine (Go)
**Repo:** skill-product-appraisal

---

## Phase Overview

### Phase 1: Design (sequential, 1 agent)
**Agent:** agent-design (opus, foreground)
**Story:** STORY-260213-2bq84r (design-cli-api-data-model)

| Order | Task | ID |
|-------|------|----|
| 1 | define-input-data-schema | TASK-260213-xr01kp |
| 2 | design-calculator-modules | TASK-260213-30s4ae |
| 3 | design-cli-commands-dsl | TASK-260213-28kao6 |
| 4 | scaffold-go-project | TASK-260213-1g3gvv |

Why 1 agent: sequential dependencies (schema → modules → CLI → scaffold). Design choices cascade. Splitting loses context.

### Phase 2: Implement Calculators (parallel, 3 agents)
**Story:** STORY-260213-1mc5zg (implement-calculators)

| Agent | Tasks | IDs |
|-------|-------|-----|
| agent-calc-1 | pricing + bundle calculators | TASK-260213-1cvvxs, TASK-260213-11yhty |
| agent-calc-2 | financial + customer calculators | TASK-260213-3fk3pg, TASK-260213-1kge2e |
| agent-calc-3 | product + scoring calculators | TASK-260213-1pg8l6, TASK-260213-24hyuo |

Why 3 agents: 6 calculator modules, each independent. Group by domain affinity (pricing+bundle, financial+customer, product+scoring). Each agent writes tests for its calculators.

### Phase 3: Agent-Facing DSL (sequential, 1 agent)
**Agent:** agent-dsl (opus, foreground)
**Story:** STORY-260213-xm37vr (agent-facing-dsl)

| Order | Task | ID |
|-------|------|----|
| 1a | implement-dsl-parser | TASK-260213-3ubk2z |
| 1b | implement-field-projection | TASK-260213-q9ftsp |
| 2 | implement-query-dispatch | TASK-260213-3b827u |
| 3a | implement-schema-introspection | TASK-260213-vdxkt7 |
| 3b | implement-scoped-grep | TASK-260213-1pghmi |

Why 1 agent: parser/projection/dispatch are tightly coupled. Agent-facing-api skill assets provide reference implementations.

### Phase 4: Integration (coordinator, no agent)
**Story:** STORY-260213-u2mzj7 (integrate-with-skill)

| Order | Task | ID |
|-------|------|----|
| 1 | update-skill-md-cli-refs | TASK-260213-131rpp |
| 2 | update-references-cli-commands | TASK-260213-rtv6p1 |
| 3 | package-cli-binary | TASK-260213-200tp6 |

Why coordinator: skill file updates need full context of what was built. Quick edits, not heavy coding.

---

## Critical Path

```
design (4 tasks seq) → calculators (6 tasks, 3 parallel agents) → DSL (5 tasks seq) → integration (3 tasks seq)
```

## Agent Allocation

- Phase 1: 1 opus agent (sequential design chain)
- Phase 2: 3 opus agents (parallel calculator implementation + tests)
- Phase 3: 1 opus agent (DSL + query layer)
- Phase 4: coordinator

## Key Skills

- `agent-facing-api` — DSL design pattern, parser/projection/grep assets
- `go-testing-tools` — test patterns for calculator modules
- `product-appraisal` — source of truth for all formulas

## Calculator Module Inventory (32+ formulas)

### Pricing (6 formulas)
- BVR, TierGapAnalysis, CostFloor, PriceValueRatio, PremiumPriceIndex, BundleDiscount

### Bundle (5 formulas)
- ClassifyComponents (L/F/K), DeadWeightRatio, CrossSubsidyAnalysis, ComponentActivation, MultiComponentUsage

### Financial (9 formulas)
- UnitEconomics, GrossMarginPerCustomer, CLV, CACPayback, BreakEven, CannibalizationNet, StressTest, IncrementalRevenue, RevenueUplift

### Customer (7 formulas)
- ChurnRate, RetentionRate, NPS, CSAT, ChurnReductionImpact, RevenueGrowthRate, ServiceRevenueShare

### Product (8 formulas)
- PenetrationRate, MigrationRate, CannibalizationRate, CrossSellRate, FeatureUtilizationRate, ComponentActivationRate, AttachRate, TrialConversion

### Scoring (3 formulas)
- GoNoGo (weighted scoring), RiskMatrix, DimensionScore

## Output

CLI binary: `appraise` (in ~/.local/bin/)
Go module: `tools/appraise/` within skill-product-appraisal repo
