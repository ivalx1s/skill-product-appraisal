# TASK-260212-3s5bk1: generalize-kpis

## Description
Replace telecom KPIs (ARPU, Revenue per GB, prepaid/postpaid gap) with universal product KPIs. Keep structure: Revenue, Customer, Product Performance, Segment, Bundle Economics, Market Context.

Use .research/audit-inventory.md as input — focus on KPI tables marked 'adaptable'.

For each KPI:
- Replace telecom-specific definition with universal equivalent
- Keep formulas where applicable (CLV = ARPU x Margin x Lifespan → CLV = RPU x Margin x Lifespan)
- Replace telecom benchmark ranges with cross-industry ranges where known
- Flag KPIs where benchmarks are highly industry-dependent — mark as 'benchmark per industry'

Output: .research/generalized-kpis.md

## Scope
(define task scope)

## Acceptance Criteria
1. All 6 KPI categories generalized. 2. No telecom-specific KPIs remain. 3. Each KPI has definition + formula + benchmark (or 'benchmark per industry' flag). 4. Stored in .research/generalized-kpis.md.
