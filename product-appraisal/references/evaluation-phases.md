# Evaluation Phases Reference

Each evaluation runs as 9 phases using a MapReduce pattern. Phases that can
run independently execute in parallel (Map); synthesis steps merge results
(Reduce). Each phase produces a standalone document. Agents read only what
they need — no full-context accumulation.

**Naming convention:** `{slug}-p{N}-{phase}.md` where slug is a short product
identifier (e.g., `bosch-glm100`, `mts-one-plus`, `notion-teams`).

**Storage:** All phase docs go to `.research/` in the project directory.

---

## Phase 0: Research & Data Collection

**Goal:** Gather ALL raw data before analysis begins. No scoring, no opinions.

Phase 0 splits into 4 sub-phases. Sub-phases 0a, 0b, and 0c run in parallel
(different agents). Sub-phase 0d synthesizes everything into the final p0 doc.

**Source attribution rule:** Every data point gets a URL. No exceptions.

**Search log:** Start `{slug}-search-log.md` — log every web search query
and URL visited with what was found. Later phase agents check this log
before doing new searches (avoids duplication, saves context).
Each sub-phase agent appends to the same log file.

### Sub-phase 0a: Target Product

**Goal:** Full specs of the product being evaluated.

**Agent:** 1 agent.

**Reads:** Product URL/brief from user.

**Does:**
- Official specs, features, price, included accessories
- SKU variants, regional pricing differences
- Official positioning and claims
- Customer reviews + ratings from 2-3 major sources (aggregated)

**Output:** `{slug}-p0a-product.md`

**Structure:**
```markdown
# {Product}: Product Data
## Official Specification
## Price & SKU Variants
## Included Accessories / Components
## Customer Reviews (aggregated: rating, count, top complaints)
## Source Index
```

**Size target:** 60-120 lines.

### Sub-phase 0b: Competitor Research (parallel)

**Goal:** Deep data on each competitor product — one agent per competitor.

**Agent:** 1 agent PER competitor (5-8 agents in parallel).

The orchestrator identifies the top 5-8 competitors from initial research
or user input, then launches one agent per competitor simultaneously.

**Each agent does:**
- Official specs, features, price
- Key differentiators vs. target product
- Customer reviews + ratings (aggregated)
- Market positioning and target segment

**Each agent outputs:** `{slug}-competitor-{competitor-slug}.md`

**Structure (per competitor):**
```markdown
# {Competitor}: Research Summary
## Specification
## Price & Variants
## Key Differentiators
## Customer Reviews (aggregated)
## Market Positioning
## Source Index
```

**Size target:** 40-80 lines per competitor.

**Naming examples:**
- `bosch-glm100-competitor-leica-disto-d2.md`
- `bosch-glm100-competitor-dewalt-dw099s.md`
- `bosch-glm100-competitor-makita-ld050p.md`

### Sub-phase 0c: Market Context

**Goal:** Market-level data — not product-specific.

**Agent:** 1 agent (can run in parallel with 0a and 0b).

**Does:**
- Market size, growth trends, segment breakdown
- Pricing landscape: price range in category, premium vs budget positioning
- Distribution channels
- Regulatory or standards context (if relevant)
- Industry trends affecting the category

**Output:** `{slug}-p0c-market.md`

**Structure:**
```markdown
# {Category}: Market Context
## Market Size & Growth
## Segment Breakdown
## Pricing Landscape (range, tiers, positioning)
## Distribution Channels
## Trends & Regulatory Context
## Source Index
```

**Size target:** 60-120 lines.

### Sub-phase 0d: Synthesis

**Goal:** Merge all sub-phase outputs into the canonical p0 doc.

**Agent:** 1 agent (or orchestrator). Runs AFTER 0a + 0b + 0c complete.

**Reads:** `{slug}-p0a-product.md`, all `{slug}-competitor-*.md`,
`{slug}-p0c-market.md`

**Does:**
- Merge into unified `{slug}-p0-research.md`
- Build competitor comparison table (specs + prices side by side)
- Identify gaps — if any competitor data is missing, flag it
- Compile complete source index

**Output:** `{slug}-p0-research.md` (canonical — used by all later phases)

Also maintains: `{slug}-search-log.md` (URL log across all phases)

**Structure:**
```markdown
# {Product}: Research & Data Collection
## Product Specification (from p0a)
## Competitor Comparison Table
## Competitor Detail (summary per competitor, link to full docs)
## Market Context (from p0c)
## Customer Signals (from p0a reviews)
## Pricing Landscape (from p0c)
## Data Gaps & Flags
## Source Index (merged from all sub-phases)
```

**Size target:** 200-400 lines. Raw data, tables, URLs. No analysis.

---

## Phase 1: Product-Market Fit

**Goal:** Does the target segment exist and want this product?

**Reads:** p0-research.md, `references/assessment-criteria.md` (PMF section)

**Evaluates:** PMF-1 through PMF-7
- Addressable segment size (>5% of customer base)
- Premium demand evidence
- WTP at proposed price
- Service-need alignment
- Conversion funnel viability
- Brand permission

**CLI:** `appraise calc product penetration_rate`, `appraise calc product trial_conversion`

**Gate:** segment <5% OR WTP below price → **No-Go**

**Output:** `{slug}-p1-pmf.md`

**Structure:**
```markdown
# {Product}: Phase 1 — Product-Market Fit
## Target Segment Definition
## Criteria Evaluation (PMF-1 through PMF-7)
## Gate Check
## Dimension Score (1-5) + Rationale
```

**Size target:** 80-150 lines.

---

## Phase 2: Pricing Adequacy

**Goal:** Is the price justified by value and supported by WTP?

**Reads:** p0-research.md, p1-pmf.md (segment context), `references/pricing-methods.md`

**Evaluates:** PRC-1 through PRC-8
- Bundle Value Ratio (BVR)
- Price-value perception
- Tier gap architecture (if multi-tier)
- GBB anchoring effectiveness
- Affordability
- Price floor clearance
- Behavioral pricing coherence

**CLI:** `appraise calc pricing bvr`, `appraise calc pricing tier_gap`,
`appraise calc pricing cost_floor`, `appraise calc pricing price_value_ratio`,
`appraise calc pricing premium_price_index`, `appraise calc pricing bundle_discount`

**Gate:** BVR <1.0 → Redesign pricing. Tier gaps disproportionate → Restructure.

**Output:** `{slug}-p2-pricing.md`

**Structure:**
```markdown
# {Product}: Phase 2 — Pricing Adequacy
## Price-Value Analysis
## BVR Calculation (with CLI output)
## Tier Gap Analysis (if applicable)
## Competitive Price Positioning
## Gate Check
## Dimension Score (1-5) + Rationale
```

**Size target:** 80-150 lines.

---

## Phase 3: Bundle Composition

**Goal:** Is each component earning its place?

**Reads:** p0-research.md, `references/bundle-valuation.md`

**Evaluates:** BND-1 through BND-8
- Leaders / Fillers / Killers classification
- Dead weight ratio (<40%)
- Dilution risk
- Cross-subsidy balance
- Access constraints
- Complementarity
- Customizability
- Switching cost creation

**Adaptation for non-bundle products:** Reinterpret "components" as features,
included accessories, ecosystem integrations, software/app, warranty, service.
Classify each as Leader (drives purchase), Filler (nice bonus), or Killer
(reduces perceived value).

**CLI:** `appraise calc bundle classify`, `appraise calc bundle dead_weight`,
`appraise calc bundle cross_subsidy`, `appraise calc bundle component_activation`

**Gate:** No clear Leader → Redesign. Dead weight >40% → Remove Killers.

### MapReduce (always)

Always split into Map + Reduce, regardless of component count.

**Map:** 1 agent per component (or component group of 2-3 if >8). Each agent:
- Researches standalone value, usage data, market comparables
- Classifies L/F/K for its component(s)
- Outputs `{slug}-p3-component-{name}.md`

**Reduce:** 1 agent reads all component docs, builds unified L/F/K table,
calculates dead weight, assesses cross-subsidies, checks complementarity.

**Output:** `{slug}-p3-bundle.md`

**Structure:**
```markdown
# {Product}: Phase 3 — Bundle Composition
## Component Classification (L/F/K table)
## Dead Weight Analysis
## Cross-Subsidy Assessment (if applicable)
## Gate Check
## Dimension Score (1-5) + Rationale
```

**Size target:** 80-150 lines.

---

## Phase 4: Competitive Positioning

**Goal:** How defensible is this in the market?

**Reads:** p0-research.md (competitor data), individual competitor docs from
Phase 0b, p2-pricing.md (price context)

**Evaluates:** CMP-1 through CMP-7
- Feature uniqueness (>3 non-replicable within 6 months)
- Competitive BVR comparison
- Time to full imitation
- Competitive response risk
- Defensible differentiation (>=2 exclusive components)
- Segment ownership
- Cross-competitive set comparison

**CLI:** `appraise calc pricing premium_price_index`

**Gate:** <2 defensible + <6mo imitation → Rethink

### MapReduce (always)

Always split into Map + Reduce, regardless of competitor count.

**Map:** 1 agent per competitor. Each agent reads the target product data
(p0a) + its competitor doc (from 0b) + p2-pricing. Produces:
- Feature-by-feature comparison (target vs. this competitor)
- BVR comparison
- Imitation time estimate
- Likely competitive response
- Output: `{slug}-p4-vs-{competitor-slug}.md` (30-50 lines each)

**Reduce:** 1 agent reads all p4-vs-* docs. Builds:
- Combined feature matrix
- Overall defensibility assessment
- Cross-competitive set analysis
- Gate check

**Output:** `{slug}-p4-competitive.md`

**Structure:**
```markdown
# {Product}: Phase 4 — Competitive Position
## Feature Comparison Matrix (with source URLs)
## Defensibility Assessment
## Competitive Response Analysis
## Gate Check
## Dimension Score (1-5) + Rationale
```

**Size target:** 100-200 lines. Feature matrix can be large.

---

## Phase 5: Financial Viability

**Goal:** Do the unit economics work?

**Reads:** p0-research.md, p2-pricing.md, `references/kpi-catalog.md`

**Evaluates:** FIN-1 through FIN-9
- Revenue per customer uplift
- Unit economics per tier (must be positive)
- Cannibalization rate
- Partner cost ratio (<30% of premium delta)
- CLV premium (>=2x base)
- Time to break-even
- Stress test (costs +20%, growth -30%)
- Scale economics
- Standalone cannibalization

**CLI:** `appraise calc financial unit_economics`, `appraise calc financial clv`,
`appraise calc financial cac_payback`, `appraise calc financial break_even`,
`appraise calc financial stress_test`, `appraise calc financial cannibalization`,
`appraise calc financial revenue_uplift`

**Gate:** Unit economics negative → Reprice. Stress test fails → Build buffers.

**Output:** `{slug}-p5-financial.md`

**Structure:**
```markdown
# {Product}: Phase 5 — Financial Viability
## Unit Economics Table
## Cannibalization Analysis
## Stress Test Results (CLI output)
## Break-even Timeline
## Gate Check
## Dimension Score (1-5) + Rationale
```

**Size target:** 80-150 lines.

---

## Phase 6: Customer Experience

**Goal:** Will this improve satisfaction and reduce churn?

**Reads:** p0-research.md (reviews/ratings), p1-pmf.md (segment context)

**Evaluates:** CX-1 through CX-8
- Churn reduction (premium < 50% of base)
- NPS improvement
- CSAT (>80%)
- Feature utilization (>60% used by >60%)
- WTP validation (price < 80% of WTP)
- Support experience
- Component engagement
- Disappointment risk (<15% drop after 3 months)

**CLI:** `appraise calc customer churn_rate`, `appraise calc customer nps`,
`appraise calc customer csat`, `appraise calc product feature_utilization`

**Gate:** Price >80% WTP → Reprice. Disappointment risk >15% → Improve quality.

**Output:** `{slug}-p6-cx.md`

**Structure:**
```markdown
# {Product}: Phase 6 — Customer Experience
## Satisfaction Signals (reviews, NPS proxy)
## Feature Utilization Assessment
## Disappointment Risk
## Gate Check
## Dimension Score (1-5) + Rationale
```

**Size target:** 80-120 lines.

---

## Phase 7: Market Reach

**Goal:** Where can this actually launch/sell?

**Reads:** p0-research.md (market context), p1-pmf.md (segment definition)

**Evaluates:** MR-1 through MR-7
- Access coverage (>70% addressable)
- Segment affordability
- Revenue ceiling per segment
- Segment competition
- Delivery readiness
- Demand heterogeneity (primary + secondary markets)
- Product adaptation (component swapping)

**Gate:** Coverage <70% → Adapt. No viable secondary markets → Niche strategy.

**Output:** `{slug}-p7-market.md`

**Structure:**
```markdown
# {Product}: Phase 7 — Market Reach
## Geographic / Segment Coverage
## Affordability Analysis
## Expansion Potential
## Gate Check
## Dimension Score (1-5) + Rationale
```

**Size target:** 60-100 lines.

---

## Phase 8: Scoring & Decision

**Goal:** Synthesize all dimensions into a Go/No-Go decision.

**Reads:** p1 through p7 docs — ONLY the "Dimension Score" sections (score +
rationale). Do NOT re-read full analysis. Each phase doc ends with a score
and rationale — that's all this phase needs.

**Does:**
- Collect 7 dimension scores from p1-p7
- Compute weighted total via `appraise calc scoring go_no_go`
- Build risk matrix via `appraise calc scoring risk_matrix`
- Write executive summary + recommendation
- Top 3 strengths, top 3 risks

**CLI:** `appraise calc scoring go_no_go --input scoring.json`

**Output:** `{slug}-final-report.md`

**Structure:**
```markdown
# {Product}: Final Report
## Executive Summary (1 page)
## Scorecard (7 dimensions + weighted total)
## Gate Pass/Fail Summary
## Risk Matrix
## Top 3 Strengths
## Top 3 Risks
## Recommendation: Go / Conditional Go / Redesign / No-Go
## Phase Document Index (links to p0-p7)
```

**Size target:** 80-120 lines. Concise synthesis, not a rehash.

---

## Orchestration: MapReduce Pattern

The evaluation uses a MapReduce pattern throughout. Phases that are
data-heavy or repeat across entities (competitors, components) split
into parallel Map agents + a Reduce agent that synthesizes.

### Dependency Graph

```
         ┌─────────┐
         │  INPUT   │  (product URL / brief)
         └────┬─────┘
              │
    ┌─────────┼──────────┐
    v         v          v
  [P0a]    [P0b x N]   [P0c]     ← MAP (parallel)
  product  competitors  market
    │         │          │
    └─────────┼──────────┘
              v
           [P0d]                  ← REDUCE → p0-research.md
              │
       ┌──────┴──────┐
       v              v
     [P1]           [P3]         ← MAP (parallel, both read p0 only)
     PMF            Bundle(M+R)
       │              │
       v              │
     [P2]             │          ← sequential (needs P1)
     Pricing          │
       │              │
       ├──────────────┘
       v
     [P4]                        ← MAP per competitor → REDUCE
     Competitive(M+R)
       │
       ├──────────┐
       v          v
     [P5]       [P6]  [P7]      ← P5 needs P2+P4; P6+P7 need p0+p1 only
     Financial   CX   Market
       │          │     │
       └────┬─────┴─────┘
            v
          [P8]                   ← REDUCE → final-report.md
          Scoring
```

**Key dependency:** P6 and P7 depend on p0 + p1 only (reviews, segment).
They do NOT depend on P5 (financial). This means P5, P6, P7 can all
run in the same wave after P4 completes.

P3 always splits into per-component Map agents + Reduce.
P4 always splits into per-competitor Map agents + Reduce.

### Execution Waves

| Wave | Phases | Pattern | Agents |
|------|--------|---------|--------|
| 1 | P0a + P0b(xN) + P0c | Map | 2 + N competitors |
| 2 | P0d | Reduce | 1 |
| 3 | P1 + P3-map(xM) | Map | 1 + M components |
| 4 | P3-reduce + P2 | Reduce + Sequential | 2 |
| 5 | P4-map(xN) | Map | N competitors |
| 6 | P4-reduce + P5 + P6 + P7 | Reduce + Parallel | 4 |
| 7 | P8 | Reduce | 1 |

**Typical evaluation** (6 competitors, 5 components): ~25 agents across 7 waves.
Every phase that has entities (competitors, components) always splits into
parallel Map agents + Reduce. No conditional logic.

### Board Integration

Every evaluation is tracked via `task-board` (requires `project-management` skill).
The orchestrator creates tasks wave-by-wave, not upfront — because later waves
depend on results from earlier ones (e.g., component list from P0, competitor
count from P0b).

**Per-wave orchestrator loop:**
1. Create story for the wave: `task-board create story --epic EPIC-XX --name "wave-{N}-..."`
2. Create task per agent: `task-board create task --story STORY-XX --name "..."`
3. Set dependencies (reduce tasks blocked by map tasks)
4. Assign agents + set status `development`
5. Monitor: `task-board q --format compact 'list(type=task, status=development) { overview }'`
6. Review outputs — read `.research/` files, check quality
7. Mark tasks `done` (or return to `development` if issues)
8. Plan next wave based on results

**Task naming:** Each task name starts with its phase ID for clarity:
- `p0a-product-specs`, `p0b-leica-disto-d2`, `p0c-market-context`
- `p1-product-market-fit`, `p3-map-bluetooth-module`, `p4-map-vs-leica`
- `p0d-research-synthesis`, `p3-reduce-bundle`, `p4-reduce-competitive`

**Agent notes:** After each agent completes, log the output path:
```bash
task-board progress notes TASK-XX "Output: .research/{slug}-{output}.md"
```

### Agent Prompt Pattern

Each phase agent gets this structure:
```
You are evaluating {product} — Phase {N}: {phase name}.
Board task: {TASK-ID}

Read these files:
- product-appraisal/SKILL.md (methodology overview)
- product-appraisal/references/{relevant-ref}.md
- .research/{slug}-p0-research.md (raw data)
- .research/{slug}-p{N-1}-*.md (previous phase, if needed)
- .research/{slug}-search-log.md (CHECK BEFORE new web searches)

Evaluate criteria {IDs}. Use `appraise` CLI for calculations.
If you do web searches, APPEND to the search log immediately.

Output: .research/{slug}-p{N}-{phase}.md
End with: Dimension Score (1-5) and 1-paragraph rationale.

When done: task-board progress notes {TASK-ID} "Output: .research/{slug}-p{N}-{phase}.md"
```

### Map Agent Prompt Pattern (for parallel sub-tasks)
```
You are researching {entity} as part of evaluating {product}.
Phase {N}, sub-task: {description}.
Board task: {TASK-ID}

Read these files:
- .research/{slug}-p0a-product.md (target product data)
- .research/{slug}-search-log.md (CHECK BEFORE new web searches)

Research {entity}. Focus ONLY on this entity — do not research others.
If you do web searches, APPEND to the search log immediately.

Output: .research/{slug}-{output-name}.md

When done: task-board progress notes {TASK-ID} "Output: .research/{slug}-{output-name}.md"
```

### Reduce Agent Prompt Pattern (for synthesis sub-tasks)
```
You are synthesizing {phase description} for {product}.
Phase {N}, synthesis step.
Board task: {TASK-ID}

Read these files:
- .research/{slug}-{map-output-pattern}-*.md (all Map outputs)
- product-appraisal/references/{relevant-ref}.md

Merge all Map outputs into a unified document.
Build comparison tables. Flag any data gaps.

Output: .research/{slug}-p{N}-{phase}.md
End with: Dimension Score (1-5) and 1-paragraph rationale.

When done: task-board progress notes {TASK-ID} "Output: .research/{slug}-p{N}-{phase}.md"
```
