# TASK-260213-24hyuo: implement-scoring-calculator

## Description
Go package: internal/calculators/scoring. Functions: GoNoGo(dimensionScores map[string]float64, weights map[string]float64) GoNoGoResult (weighted total, decision: StrongGo/ConditionalGo/Redesign/NoGo, per-dimension breakdown); RiskMatrix(risks []Risk) []RiskResult (probability x impact = score, sorted by severity); DimensionScore(criteria []CriterionResult) float64 (aggregate criterion pass/fail into 1-5 dimension score). Default weights from SKILL.md (PMF 15%, FIN 25%, PRC+CX 20%, CMP 15%, BND 10%, MR 10%, Risk 5%). Tests.

## Scope
(define task scope)

## Acceptance Criteria
(define acceptance criteria)
