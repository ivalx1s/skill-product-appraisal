# TASK-260212-1detcq: generalize-assessment-dimensions

## Description
Take the 7 assessment dimensions (PMF, Pricing, Bundle, Competitive, Financial, CX, Regional) and replace telecom-specific criteria/thresholds with universal equivalents. Keep the structure, change the content.

Use .research/audit-inventory.md (from audit task) as input — focus on items marked 'adaptable'.

For each dimension:
- Replace industry-specific criteria with universal ones (e.g. 'ARPU uplift' → 'revenue per user uplift')
- Replace telecom thresholds with cross-industry benchmarks where available
- Keep criteria IDs (PMF-1, PRC-1, etc.) for traceability
- Flag criteria where no universal threshold exists — mark as 'calibrate per industry'

Output: .research/generalized-dimensions.md

## Scope
(define task scope)

## Acceptance Criteria
1. All 7 dimensions generalized. 2. No telecom-specific terms remain in criteria. 3. Each criterion has either a universal threshold or 'calibrate per industry' flag. 4. Stored in .research/generalized-dimensions.md.
