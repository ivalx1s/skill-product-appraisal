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

---

## Evaluation Workflow

Seven dimensions, evaluated sequentially. Each dimension has gate conditions
that can trigger early exit (No-Go) or forced redesign before proceeding.

```
1. PRODUCT-MARKET FIT  ─── Does anyone want this?
   GATE: segment <5% of base OR WTP below price → No-Go
   │
   v
2. PRICING ADEQUACY  ──── Is the price justified?
   GATE: BVR <1.0 at any tier → Redesign pricing
   GATE: tier gaps disproportionate → Restructure tiers
   │
   v
3. BUNDLE COMPOSITION  ── Is each component earning its place?
   GATE: no clear Leader → Redesign bundle
   GATE: dead weight >40% → Remove Killers or add swappability
   GATE: access constraints >30% → Adapt for markets
   │
   v
4. COMPETITIVE POSITION ─ Can it survive in the market?
   GATE: <2 defensible components AND <6mo to imitation → Rethink
   │
   v
5. FINANCIAL VIABILITY ── Does the math work?
   GATE: unit economics negative at all tiers → Reprice or cut costs
   GATE: net revenue negative after cannibalization → Redesign migration
   GATE: stress test fails (costs +20%, growth -30%) → Build buffers
   │
   v
6. CUSTOMER EXPERIENCE ── Will it improve the relationship?
   GATE: price >80% of WTP → Reprice
   GATE: disappointment risk >15% drop → Improve quality
   │
   v
7. MARKET REACH  ──────── Where can this actually launch?
   GATE: access coverage <70% → Adapt components
   GATE: no viable secondary markets → Consider niche strategy
   │
   v
8. GO / NO-GO / REDESIGN DECISION
   Weighted scoring across all 7 dimensions
```

**Critical rule:** Do NOT skip dimensions. A product can score 5/5 on financials
but fail on bundle composition. The sequential flow catches structural problems
before you waste time modeling revenue on a broken bundle.

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

## Theoretical Foundations (Summary)

**Variance reduction** (Schmalensee, 1984): Bundling reduces valuation variance,
enabling better surplus extraction through a single price point.

**Mixed bundling optimality** (McAfee, McMillan & Whinston, 1989): Mixed bundling
almost always strictly increases profits vs. pure bundling or pure component
selling. Always offer the bundle AND standalone components simultaneously.

**Dilution effect** (Shaddy & Fishbach, 2017): Consumers perceive bundles as
gestalt units. Adding low-value components can reduce total WTP. Removing a
component increases perceived loss; adding yields diminished perceived gain.

**Good-Better-Best** (multiple sources; HBR Mohammed, 2018): Three-tier pricing
exploits the compromise effect. ~66% choose the middle option. The top tier
anchors; the bottom tier serves as a decoy.

**Bundle framing** (Wansink et al., 1998): Presenting items as a bundle rather
than individually boosts sales by ~32%, even at equivalent prices.

See `references/pricing-methods.md` for detailed theory and application guidance.

---

## Output Format

The evaluation deliverable follows this structure:

### 1. Executive Summary (1 page)
- Product/bundle being evaluated
- Target segment and market context
- Overall score and Go/No-Go recommendation
- Top 3 strengths and top 3 risks

### 2. Dimension Scores (1-2 pages)
- Score per dimension (1-5) with 1-paragraph justification each
- Gate pass/fail status for each dimension
- Weighted total score

### 3. Detailed Findings (per dimension)
- Each criterion evaluated with evidence and source URLs
- Red flags and gate failures highlighted
- Comparison to benchmarks (tagged as Verified / Practitioner Guidance / Calibrate)

### 4. Financial Model Summary
- Unit economics table (revenue, costs, margin per customer per tier)
- Cannibalization analysis
- Stress test results (costs +20%, growth -30%)
- Break-even timeline

### 5. Competitive Landscape
- Feature-by-feature comparison matrix (with source URLs for every data point)
- BVR comparison across competitors
- Defensibility assessment

### 6. Risk Matrix
- Each risk: probability (1-5) x impact (1-5) = score
- Mitigation plan for top risks

### 7. Recommendation
- Go / Conditional Go / Redesign / No-Go with specific next steps
- If Conditional Go or Redesign: list exactly which criteria must be addressed

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

## Quick Start

1. **Define scope:** What product/bundle? What market? What target segment?
2. **Gather data:** Collect all product characteristics WITH source URLs.
   Apply the source attribution rule -- no unsourced data.
3. **Walk the 7 dimensions sequentially.** Do not skip. Stop at any gate failure.
4. **Score each dimension** (1-5). Compute weighted total.
5. **Write the deliverable** following the output format above.
6. **Decide:** Go / Conditional Go / Redesign / No-Go.

For the full criteria checklist, start with `references/assessment-criteria.md`.
For a worked example, see `references/example-telecom-appraisal.md`.
