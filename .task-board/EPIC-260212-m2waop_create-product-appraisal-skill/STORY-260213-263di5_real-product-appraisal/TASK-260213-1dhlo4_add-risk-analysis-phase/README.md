# TASK-260213-1dhlo4: add-risk-analysis-phase

## Description
Add a dedicated Risk Analysis phase to the evaluation pipeline.

PLACEMENT: Runs in Wave 6 alongside P5 (Financial), P6 (CX), P7 (Market Reach). Depends on p0 + p1 + p4 (same deps as P5). Does NOT need P5 output — financial stress test is P5's job, this phase covers non-financial risks. P8 (Scoring) reads risk phase output for go_no_go.

WHY SEPARATE PHASE: Currently risks are scattered — P5 has stress test, P8 has risk_matrix, but no dedicated phase for systematic risk identification and assessment. A risk phase catches threats that slip between other phases.

SCOPE:
1. Supply chain risks — single-source dependency, component availability, manufacturing concentration, lead times
2. Technology obsolescence — how fast can this become irrelevant (phones replacing standalone devices, AI disruption, platform shifts)
3. Regulatory & compliance risks — pending legislation, certification changes, import/export restrictions, environmental regulations
4. Geopolitical risks — sanctions, trade wars, currency volatility affecting pricing/supply
5. IP & patent risks — patent exposure, freedom to operate, competitor patent walls
6. Competitive response risks — what competitors can do to neutralize this product (price war, feature copy, bundling counter)
7. Reputational risks — brand association risks, quality failure cascading, negative viral potential
8. Market timing risks — launching too early/late, seasonal misalignment, economic cycle position
9. Dependency risks — platform dependency (app stores, BT standards), partner dependency, ecosystem lock-in fragility

METHODOLOGY: Each risk gets: likelihood (1-5), impact (1-5), risk score (L x I), mitigation options. Build a risk matrix. Flag any score >= 15 as critical, >= 10 as high.

OUTPUT: {slug}-p{N}-risks.md
SIZE TARGET: 100-150 lines.

CLI: appraise calc scoring risk_matrix (already exists, may need input schema update for structured risk items)

DOWNSTREAM IMPACT:
- P8 (Scoring): risk phase dimension score feeds go_no_go, risk matrix feeds final report
- Risk matrix becomes a standalone section in final report

WAVE IMPACT: None — fits into existing Wave 6. No new waves needed.

FILES TO UPDATE (together with other phase tasks):
- evaluation-phases.md: add risk phase, update dependency graph, update wave table
- SKILL.md: update phase count, wave diagram
- assessment-criteria.md: add RISK-1 through RISK-9 criteria

## Scope
(define task scope)

## Acceptance Criteria
(define acceptance criteria)
