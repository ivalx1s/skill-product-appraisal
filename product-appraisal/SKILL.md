---
name: product-appraisal
description: >
  Universal methodology for evaluating complex products, bundles, and subscription
  offerings across any industry. Covers: product evaluation, bundle assessment,
  pricing analysis, market positioning, competitive analysis, product appraisal,
  go/no-go decision, viability assessment, unit economics, bundle value ratio,
  dead weight analysis, willingness to pay, tier pricing.
  RU triggers: оценка продукта, анализ бандла, ценообразование, позиционирование,
  конкурентный анализ, оценка жизнеспособности, юнит-экономика, анализ пакета,
  оценка ценности, решение go/no-go.
---

# Product Appraisal

A universal framework for evaluating complex products, bundles, and subscription
offerings. Built on bundling theory (Schmalensee, McAfee-McMillan-Whinston),
behavioral pricing research (Tversky-Kahneman, Shaddy-Fishbach), and consulting
frameworks (Simon-Kucher Leaders/Fillers/Killers, Good-Better-Best). Generalized
from telecom-specific application to work across industries.

---

## Source Attribution Rule

**EVERY product characteristic -- price, feature, spec, term, limit, condition --
MUST have a reference to a source URL where the data was obtained.**

Unsourced data is invalid. If a data point cannot be traced to a public URL,
official document, or verified API response, it must be marked `[UNVERIFIED]`
and excluded from scoring calculations.

Tag every external claim with one of:
- `[Verified]` -- confirmed with direct source URL
- `[Practitioner Guidance]` -- reasonable heuristic, no single source
- `[Calibrate]` -- threshold depends on industry; default from telecom
- `[Calculated]` -- derived from sourced data using documented formula

This rule applies to BOTH the product being evaluated AND all competitor data.

### Web Search Log

During research phases, maintain a running search log:
`{slug}-search-log.md` in `.research/`. Every web search and fetch gets
logged immediately — URL visited, what was found, what was useful.

```markdown
## Search Log: {product}
| # | Agent | Query / URL | Found | Useful? | Used In |
|---|-------|-------------|-------|---------|---------|
| 1 | P0a | "bosch glm 100 specs" | Official specs page | Yes | p0a-product |
| 2 | P0b-leica | ozon.ru/product/... | Price 31,700 RUB | Yes | competitor-leica |
| 3 | P0c | "laser rangefinder market size 2025" | $3.2B global | Yes | p0c-market |
```

**Why:** Prevents re-searching the same thing in later phases. Agents in
phases 1-7 check the search log before doing new web searches. The log
also serves as a source index for the final report.

---

## Evaluation Workflow

**Nine phases, MapReduce pattern.** Each phase produces a separate document.
This prevents context overflow — each agent reads only the phase outputs it
needs, not the entire evaluation history. Independent phases run in parallel
(Map); synthesis steps merge results (Reduce). Gate failures stop the chain.

```
Phase 0: RESEARCH (MapReduce)
   0a: product ─┐
   0b: N competitors (parallel) ├─→ 0d: REDUCE → {slug}-p0-research.md
   0c: market ──┘
   │
   ├──────────────┐
   v              v
Phase 1: PMF    Phase 3: BUNDLE(M+R) (parallel, both read p0 only)
   │              │                   always per-component Map + Reduce
   v              │
Phase 2: PRICING  │
   │              │
   ├──────────────┘
   v
Phase 4: COMPETITIVE(M+R) ────────── always per-competitor Map + Reduce
   │
   v
Phase 5: FINANCIAL ─┐
   │                │
Phase 6: CX  Phase 7: MARKET  ──── (P5+P6+P7 parallel in Wave 6)
   │              │
   └──────┬───────┘
          v
Phase 8: SCORING & DECISION ──────── REDUCE → {slug}-final-report.md
```

Gate failures at any phase stop the chain. See `references/evaluation-phases.md`
for the full dependency graph and MapReduce orchestration details.

**Critical rules:**
- Do NOT skip phases. A product can score 5/5 on financials but fail on
  bundle composition. Sequential flow catches structural problems early.
- Each phase is a separate agent session. The agent reads the skill + its
  phase's reference files + previous phase outputs. This keeps context lean.
- Gate failures in phases 1-7 stop the chain. Document the failure and
  recommend specific fixes before proceeding.

### Phase Details

See `references/evaluation-phases.md` for the full phase specification:
per-phase inputs, outputs, what to read, what to write, and agent prompts.

---

## Project Tracking (task-board)

**Dependency:** Requires the `project-management` skill with `task-board` CLI.

Every evaluation is tracked on the task board. No untracked agents.
The orchestrator builds tasks wave-by-wave, reviews results after each wave,
then creates the next wave's tasks.

### Board Structure

```
EPIC: "Evaluate {product} {YYMMDD}"   ← date for re-evaluation tracking
├── STORY: "Wave 1: P0 Research (Map)"
│   ├── TASK: "P0a: {product} specs"
│   ├── TASK: "P0b: {competitor-1}"
│   ├── TASK: "P0b: {competitor-2}"
│   ├── ...
│   └── TASK: "P0c: market context"
├── STORY: "Wave 2: P0 Research (Reduce)"
│   └── TASK: "P0d: synthesize research"
├── STORY: "Wave 3: P1 PMF + P3 Bundle (Map)"
│   ├── TASK: "P1: product-market fit"
│   ├── TASK: "P3-map: {component-1}"
│   ├── TASK: "P3-map: {component-2}"
│   └── ...
├── STORY: "Wave 4: P3 Reduce + P2 Pricing"
│   ├── TASK: "P3-reduce: bundle synthesis"
│   └── TASK: "P2: pricing adequacy"
├── STORY: "Wave 5: P4 Competitive (Map)"
│   ├── TASK: "P4-map: vs {competitor-1}"
│   ├── TASK: "P4-map: vs {competitor-2}"
│   └── ...
├── STORY: "Wave 6: P4 Reduce + P5 + P6 + P7"
│   ├── TASK: "P4-reduce: competitive synthesis"
│   ├── TASK: "P5: financial viability"
│   ├── TASK: "P6: customer experience"
│   └── TASK: "P7: market reach"
└── STORY: "Wave 7: P8 Scoring"
    └── TASK: "P8: final report"
```

### Orchestrator Workflow

The orchestrator does NOT create all waves upfront. It works wave-by-wave:

1. **Create epic** for the evaluation
2. **Plan Wave 1** — create story + tasks, set dependencies
3. **Launch Wave 1 agents** — assign agents to tasks, set status `development`
4. **Monitor** — check task statuses via `task-board agents`
5. **Review Wave 1 results** — read output docs, verify quality, mark tasks `done`
6. **Plan Wave 2** — based on Wave 1 results (e.g., now we know the components
   for P3, the exact competitor list for P4)
7. **Repeat** until Wave 8 completes

**Why wave-by-wave:** Later waves depend on earlier results. You don't know
which components to analyze in P3 until P0 research is done. You don't know
competitor count for P4 until P0b results are in. Planning ahead would mean
guessing.

### Board Commands (quick reference)

```bash
task-board create epic --name "evaluate-{slug}-{YYMMDD}"
task-board create story --epic EPIC-XX --name "wave-1-p0-research-map"
task-board create task --story STORY-XX --name "p0a-product-specs" \
  --description "Phase 0a: research {product} specs, prices, reviews"
task-board assign TASK-XX --agent "p0a-agent"
task-board progress status TASK-XX development
task-board progress status TASK-XX done
task-board q --format compact 'list(type=task, status=development) { overview }'
```

### Research Artifacts

All research outputs go to `.research/` (persistent, not tied to board).
Link from task notes:
```bash
task-board progress notes TASK-XX "Output: .research/{slug}-p0a-product.md"
```

---

## The 7 Dimensions

### 1. Product-Market Fit (PMF-1 through PMF-7)

Does the target segment exist, and does the bundle address their actual needs?
Evaluate addressable segment size (>5% of customer base), demonstrated premium
demand (engagement penetration >15% and growing), willingness to pay at proposed
prices (WTP >= price for >30% of target), service-need alignment (>60% of target
wants >60% of components), evidence of fragmented consumption (unmet need),
conversion funnel viability (3-10% of eligible at entry tier), and brand
permission for premium. If there is no segment or no WTP, nothing else matters.

**Key criteria:** PMF-1 (segment size), PMF-3 (WTP validation), PMF-6 (funnel).
See `references/assessment-criteria.md` for full criteria table.

### 2. Pricing Adequacy (PRC-1 through PRC-8)

Are the price points justified by perceived value and supported by WTP?
Calculate Bundle Value Ratio (sum of standalone prices / bundle price; target
>1.5x, prefer >2.0x), validate price-value perception (survey: perceived value /
price > 1.0), analyze tier gap architecture (avoid >50% jumps without matching
value delta), test anchoring effectiveness (~66% should land on middle tier via
Good-Better-Best dynamics), check affordability for target segment, verify price
floor clearance (price > cost at all tiers), and validate behavioral pricing
coherence (entry tier drives migration to middle).

**Key criteria:** PRC-1 (BVR), PRC-4 (GBB anchoring), PRC-7 (price floor).
See `references/pricing-methods.md` for GBB, WTP, and anchoring details.

### 3. Bundle Composition (BND-1 through BND-8)

Is each component a Leader, Filler, or Killer? Classify using the Simon-Kucher
framework: Leaders (2-3 per bundle) drive purchase intent, Fillers (3-5) add
perceived value at low marginal cost, Killers (target: 0) reduce WTP or confuse
the offer. Measure dead weight ratio (components used by <20% of customers within
3 months; red flag if >40%). Test for dilution risk -- if removing a component
INCREASES WTP, it is a Killer (Shaddy & Fishbach, 2017). Assess cross-subsidy
balance, access constraints (<30% of value tied to constrained components),
complementarity, customizability, and switching cost creation.

**Key insight:** Dead weight is not purely negative. Unused components can
contribute "option value" -- perceived value from availability alone. But if
low-value components trigger the dilution effect, they reduce total bundle WTP.
The test is BND-3: does removing the component increase or decrease WTP?

**Key criteria:** BND-1 (Leaders), BND-2 (dead weight), BND-3 (dilution).
See `references/bundle-valuation.md` for theory and dead weight analysis.

### 4. Competitive Positioning (CMP-1 through CMP-7)

How does this sit in the market, and how defensible is it? Evaluate uniqueness
(>3 features not replicable within 6 months), competitive BVR (own >= competitor),
time to imitation (>12 months for full replication), competitive response risk
(game theory: likely reactions), defensible differentiation (at least 2 exclusive
components), segment ownership (brand perception), and cross-competitive set
(bundle price < sum of best individual alternatives from different providers).

**Key criteria:** CMP-3 (imitation time), CMP-5 (defensible differentiation).
See `references/assessment-criteria.md` for full competitive criteria.

### 5. Financial Viability (FIN-1 through FIN-9)

Do the unit economics work after cannibalization and cross-subsidy? Model revenue
per customer uplift, unit economics at each tier (must be positive), cannibalization
rate (net incremental revenue > 0), partner cost ratio (<30% of premium revenue
delta), lifetime value premium (target: premium CLV >= 2x base), time to
break-even, cross-subsidy sustainability under stress (costs +20%, growth -30%),
scale economics (breakeven count achievable within 12 months), and standalone
product cannibalization.

**Critical insight:** Bundling's primary financial lever is often churn reduction,
not direct revenue uplift. Even thin margins can be justified if churn reduction
delivers sufficient lifetime value improvement. Model churn reduction explicitly.

**Key criteria:** FIN-2 (unit economics), FIN-3 (cannibalization), FIN-7 (stress).
See `references/kpi-catalog.md` for metric formulas and benchmarks.

### 6. Customer Experience (CX-1 through CX-8)

Will this improve satisfaction and reduce churn? Measure churn reduction (premium
< 50% of base churn), NPS improvement (calibrate per industry), CSAT (>80%),
feature utilization (>60% of features used by >60% of customers), WTP validation
(price < 80% of measured WTP), support experience (premium resolution < 50% of
standard), component engagement (track vs. standalone benchmarks), and
disappointment risk (<15% value perception drop after 3 months).

**Key criteria:** CX-1 (churn), CX-4 (utilization), CX-8 (disappointment).
See `references/assessment-criteria.md` for full CX criteria.

### 7. Market Reach (MR-1 through MR-7)

Does this work beyond the primary market/segment? Assess access coverage of
constrained components (>70% of addressable customers), segment affordability
(calibrate per industry), revenue ceiling per segment, segment competition
(no dominant competitor in target segments), delivery readiness, demand
heterogeneity (viable in primary + secondary markets), and product adaptation
(can components be swapped per segment?).

**Key criteria:** MR-1 (access coverage), MR-6 (demand heterogeneity).
See `references/assessment-criteria.md` for full market reach criteria.

---

## Go/No-Go Decision

After scoring all 7 dimensions (1-5 scale), compute weighted total:

| Dimension | Weight |
|-----------|--------|
| Strategic fit (PMF) | 15% |
| Financial viability (FIN) | 25% |
| Customer demand (PRC + CX) | 20% |
| Competitive position (CMP) | 15% |
| Bundle composition (BND) | 10% |
| Market reach (MR) | 10% |
| Risk profile | 5% |

| Score | Decision | Action |
|-------|----------|--------|
| >= 4.0 | **Strong Go** | Proceed to launch planning |
| 3.0 - 3.9 | **Conditional Go** | Address specific gaps before launch |
| 2.0 - 2.9 | **Redesign** | Fundamental changes needed |
| < 2.0 | **No-Go** | Do not launch in current form |

See `references/evaluation-template.md` for a ready-to-fill scoring template.

---

## Theoretical Foundations

Key theories: variance reduction (Schmalensee), mixed bundling optimality
(McAfee-McMillan-Whinston), dilution effect (Shaddy-Fishbach), Good-Better-Best
pricing, bundle framing (Wansink). See `references/pricing-methods.md` for details.

---

## Output Format

Each phase produces a standalone document in `.research/`. The final
deliverable is `{slug}-final-report.md` (Phase 8), which synthesizes
all dimension scores into an executive summary.

**Per-phase docs:** `{slug}-p{N}-{phase}.md` — see `references/evaluation-phases.md`
for the structure of each phase document.

**Final report structure** (Phase 8 output):
- Executive Summary: product, segment, overall score, Go/No-Go, top 3 strengths + risks
- Scorecard: 7 dimensions + weighted total
- Gate Pass/Fail Summary
- Risk Matrix
- Recommendation with specific next steps
- Phase Document Index (links to p0-p7 docs)

### Final Deliverables Package

After Phase 8, the orchestrator MUST produce **4 standalone deliverables** and
copy them into a timestamped results folder:

```
results/{YYMMDD}_{HHmmss}_{slug}/
  01-research-summary.md    ← {slug}-p0-research.md
  02-scoring-breakdown.md   ← NEW: consolidated scoring rationale (see below)
  03-recommendations.md     ← NEW: extracted recommendations + risk matrix
  04-final-report.md        ← {slug}-final-report.md
```

**`02-scoring-breakdown.md`** (required, not in Phase 8 output):
- Weighted scorecard table (dimension, phase, raw score, weight, weighted, source doc)
- Per-dimension rationale (1 paragraph each: what scored well, deductions, gates)
- Full gate summary table (all gates from P1-P7 with PASS/WARN/FAIL)
- Calculation method (weights, how Customer Demand and Risk Profile are derived)
- This doc answers "how was the score calculated?" without requiring the reader
  to open 7 separate phase documents.

**`03-recommendations.md`** (required, extracted from final report + phase analyses):
- Decision (Go/Conditional Go/Redesign/No-Go) with score
- Conditions for full approval (if Conditional Go) with problem/actions/impact/source
- Strategic recommendations beyond conditions
- Risk mitigation matrix (risk, probability, impact, score, mitigation, owner)
- Score improvement roadmap (current vs projected scores if conditions met)

Both documents are produced by the orchestrator after Phase 8 completes, using
data already present in phase outputs. No new research or agent launches needed.

**Why separate docs:** The final report is an executive summary — concise by design.
Stakeholders who need "show your work" get the scoring breakdown. Stakeholders who
need "what do we do?" get the recommendations. The research summary provides raw
data for anyone who wants to verify claims.

---

## Calibration

Many thresholds default to subscription/telecom values. When applying to a
different industry, calibrate these:

| Threshold | Default | Calibration Examples |
|-----------|---------|---------------------|
| Revenue uplift | +12-18% | SaaS: 50-100%+; physical goods: 10-20% |
| NPS target | >30 | Tech: 40-60; retail: 50-70; financial: 20-40 |
| CAC payback | <6 months | SaaS: 12-18 months; consumer apps: 1-3 months |
| Penetration target | 10-25% | Depends on market maturity and pricing |
| Affordability | <3-5% of income | Differs by product category and wallet norms |
| Break-even | 6-18 months | Depends on investment and margin structure |
| Freemium conversion | 3-5% self-serve | Sales-assisted: 5-7%; top: 8-15% |

---

## References

Detailed materials are in `references/`. Read the overview here; dive into
reference files when you need the full criteria tables, formulas, or templates.

### `references/assessment-criteria.md`
Full criteria tables for all 7 dimensions (48 criteria total: PMF-1 through MR-7).
Each criterion has an ID, assessment method, and pass threshold. Use this when
you need the exact evaluation checklist.

### `references/bundle-valuation.md`
Bundle theory deep dive: Leaders/Fillers/Killers classification, dead weight
analysis, dilution effect mechanics, option value concept, cross-subsidy
economics. The theoretical backbone for Dimension 3 (Bundle Composition).

### `references/pricing-methods.md`
Pricing methodology: Good-Better-Best architecture, Van Westendorp PSM,
Choice-Based Conjoint, Gabor-Granger, anchoring and framing effects.
Includes the recommended 6-stage price validation sequence.

### `references/kpi-catalog.md`
Universal KPI reference across 6 categories: Revenue, Customer, Product
Performance, Premium Segment, Bundle Economics, Market Context. Each KPI has
a formula, target range, and calibration notes. 40+ metrics.

### `references/evaluation-phases.md`
Phase-by-phase evaluation workflow. Defines 9 phases (P0 research through P8
scoring), per-phase inputs/outputs, document naming, agent orchestration
pattern. **Read this first** when planning an evaluation.

### `references/evaluation-template.md`
Ready-to-fill evaluation template. Includes: Strategic Fit Scorecard, Go/No-Go
Decision Matrix, Risk Assessment Matrix, and per-dimension scoring sheets.
Copy this file and fill it in for each evaluation.

### `references/example-telecom-appraisal.md`
Worked example: a telecom premium bundle evaluation using this methodology.
Shows how each dimension is scored, what data is collected, how gates are
evaluated, and what the final deliverable looks like. Use as a pattern for
your own evaluations.

---

## Calculation CLI: `appraise`

All calculations (BVR, tier gaps, scoring, stress tests, etc.) are handled by
the `appraise` CLI tool. Do not calculate manually — use the tool.

**Install:** From the skill repo: `cd tools/appraise && make install`
Installs globally to `~/.local/bin/appraise` — works from any project directory.

### DSL Queries (primary agent interface)

```bash
# Discover all available calculations
appraise q 'schema()'

# Single calculation
appraise q 'calc(pricing.bvr)' --input data.json

# Batch multiple calculations in one call
appraise q 'calc(pricing.bvr); calc(pricing.tier_gap); calc(bundle.dead_weight)' --input data.json

# Compact output (fewer tokens)
appraise q 'calc(scoring.go_no_go)' --input data.json --format compact
```

### Direct Calculation (simpler interface)

```bash
appraise calc pricing bvr --input data.json
appraise calc scoring go_no_go --input scoring.json
appraise calc financial stress_test --input financials.json
```

### Available Modules (38 functions)

| Module | Functions | Key Calculations |
|--------|-----------|------------------|
| `pricing` | 6 | BVR, tier gap analysis, cost floor, price-value ratio, premium price index, bundle discount |
| `bundle` | 5 | L/F/K classification, dead weight ratio, cross-subsidy analysis, component activation, multi-component usage |
| `financial` | 9 | Unit economics, gross margin, CLV, CAC payback, break-even, cannibalization, stress test, incremental revenue, revenue uplift |
| `customer` | 7 | Churn rate, retention, NPS, CSAT, churn reduction impact, revenue growth, service revenue share |
| `product` | 8 | Penetration, migration, cannibalization rate, cross-sell, feature utilization, component activation, attach rate, trial conversion |
| `scoring` | 3 | Go/No-Go (weighted 7-dimension), risk matrix, dimension score |

### Input Data Format

Prepare a JSON file with the `AppraisalInput` schema. Run `appraise q 'schema()'`
to see the full schema. Minimal example for BVR:

```json
{
  "product": {
    "name": "Premium Bundle",
    "price": 2000,
    "components": [
      {"name": "Service A", "standalone_price": 1500},
      {"name": "Service B", "standalone_price": 800},
      {"name": "Service C", "standalone_price": 500}
    ]
  }
}
```

---

## Quick Start

1. **Create epic:** `task-board create epic --name "evaluate-{slug}-{YYMMDD}"`
   Date in name — product may be re-evaluated later when market changes.
2. **Wave 1 — P0 Map:** Create story + tasks for P0a (product), P0b (each competitor),
   P0c (market). Launch all agents in parallel. Every data point gets a URL.
3. **Review Wave 1:** Read outputs, verify quality, mark tasks done.
4. **Wave 2 — P0 Reduce:** Synthesize all research into `{slug}-p0-research.md`.
5. **Wave 3 — P1+P3 Map:** PMF + per-component bundle analysis in parallel.
6. **Wave 4 — P3 Reduce + P2:** Bundle synthesis + pricing (needs P1).
7. **Wave 5 — P4 Map:** Per-competitor competitive analysis in parallel.
8. **Wave 6 — P4 Reduce + P5 + P6 + P7:** Competitive synthesis + financial + CX + market (P6/P7 only need p0+p1, not P5).
9. **Wave 7 — P8:** Collect scores, `appraise calc scoring go_no_go`, final report.

Each wave: plan tasks on board → launch agents → review → plan next wave.
Use `appraise` CLI for all calculations — never calculate manually.

For the phased workflow details, see `references/evaluation-phases.md`.
For the full criteria checklist, see `references/assessment-criteria.md`.
For a worked example, see `references/example-telecom-appraisal.md`.
