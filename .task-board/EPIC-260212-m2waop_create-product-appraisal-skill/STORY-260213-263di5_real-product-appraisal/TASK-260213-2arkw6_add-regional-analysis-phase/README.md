# TASK-260213-2arkw6: add-regional-analysis-phase

## Description
Add P0e: Regional Research as a new sub-phase inside Phase 0.

PLACEMENT: P0e runs in Wave 1 alongside P0a (product), P0b (competitors), P0c (market). One agent per target region/market. P0d synthesis merges regional data into p0-research.md. No new waves or dependencies added.

WHY P0 (not a separate phase): Regional research is data collection, not analysis. It belongs with P0a/P0b/P0c. The analysis of regional data happens downstream: P2 (regional price spreads), P4 (local competitive landscape), P5 (region-weighted financial projections), P7 (regional viability scoring).

P0e AGENT SCOPE (per region):
1. Regional pricing — local prices for target product + key competitors, currency, VAT/tax, price spread across channels
2. Distribution channels — which retailers/marketplaces carry the product, online vs offline split, availability/stock status
3. Competitor purchasing patterns — how people actually buy competitor products in this region: which channels dominate, direct sales vs resellers, B2B vs B2C split, marketplace presence, service/warranty availability, after-sales support
4. Local competitive landscape — region-specific competitors not covered in global P0b (e.g. local/regional brands), local market share if available
5. Regulatory/certification — local standards, mandatory certifications, import restrictions, warranty requirements
6. Population & purchasing power — regional population, urbanization, average income/salary, relevant industry size, consumer spending patterns for the product category
7. Demand seasonality — buying seasons, regional demand patterns, YoY trends if available

OUTPUT: {slug}-p0e-region-{region-slug}.md (per region)
SIZE TARGET: 60-120 lines per region.

P0d UPDATE: P0d synthesis must merge P0e outputs — add "Regional Summary" section to p0-research.md with cross-region comparison table.

P7 UPDATE: P7 (Market Reach) evaluates regional data from P0e instead of doing its own research. P7 becomes pure scoring/analysis.

DOWNSTREAM IMPACT:
- P2 (Pricing): use regional price spreads for positioning analysis
- P4 (Competitive): reference region-specific competitors + purchasing patterns from P0e
- P5 (Financial): use region-weighted volumes, price points, purchasing power data
- P7 (Market Reach): score regional viability using P0e data
- P8 (Scoring): regional risk factors feed into go_no_go

FILES TO UPDATE:
- evaluation-phases.md: add P0e sub-phase, update P0d, update P7, update dependency graph, update wave table
- SKILL.md: update phase count reference, wave diagram
- domain/types.go: add RegionalData input fields
- Relevant calculators: accept optional regional context

## Scope
(define task scope)

## Acceptance Criteria
(define acceptance criteria)
