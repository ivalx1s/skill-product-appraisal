# Product Appraisal: Evaluation Template

> **Product/Bundle:** [Name]
> **Evaluator:** [Name / Agent ID]
> **Date:** [YYYY-MM-DD]
> **Industry:** [e.g., SaaS, Telecom, Retail, Media]
> **Version:** [1.0]

---

## How to Use This Template

1. Work through dimensions 1-7 sequentially. Each dimension has a **gate** -- if a gate fails, stop and address the issue before continuing.
2. For every data point, fill in the **Source** column. Acceptable source types: `[Verified]` (URL/citation), `[Internal Data]`, `[Practitioner Guidance]`, `[Estimate]`, `[Calibrated]`.
3. Use `appraise` CLI for all calculations:
   - `appraise calc pricing bvr --input data.json` for BVR
   - `appraise calc pricing tier_gap --input data.json` for tier gap analysis
   - `appraise calc bundle classify --input data.json` for L/F/K classification
   - `appraise calc financial stress_test --input data.json` for stress tests
   - `appraise calc scoring go_no_go --input scoring.json` for final weighted score
4. Score each criterion 1-5 (1 = fails badly, 3 = meets threshold, 5 = exceeds significantly).
5. After all dimensions, fill the condensed scorecard at the end.
6. Thresholds marked `*Calibrate*` must be set for your specific industry before evaluation begins.

---

## Dimension 1: Product-Market Fit

**Core question:** Does the target segment exist, and does the bundle address their actual needs?

**Gate:** If PMF-1 fails (<5% addressable) or PMF-3 fails (WTP below price) --> **No-Go**.

### 1.1 Data Collection

| ID | Criterion | Data / Finding | Source | Method Used |
|----|-----------|---------------|--------|-------------|
| PMF-1 | Addressable target segment size | | | |
| PMF-2 | Demonstrated demand for premium tier | | | |
| PMF-3 | Willingness to pay at proposed prices | | | |
| PMF-4 | Service-need alignment (>60% wants >60% of components) | | | |
| PMF-5 | Current unmet need (fragmented consumption evidence) | | | |
| PMF-6 | Conversion funnel viability | | | |
| PMF-7 | Brand permission for premium (ICE assessment) | | | |

### 1.2 Conversion Funnel Detail

| Funnel Stage | Count / Estimate | % of Previous | Source |
|-------------|-----------------|---------------|--------|
| Total customer base | | 100% | |
| Engaged / active users | | | |
| Premium-eligible (income, behavior) | | | |
| Projected converters (entry tier) | | | |
| Projected converters (top tier) | | | |

**Benchmark:** 3-10% of eligible at entry tier; 1-5% at top tier.

### 1.3 Assessment

| ID | Threshold | Pass / Fail | Score (1-5) | Notes |
|----|-----------|-------------|-------------|-------|
| PMF-1 | >5% of total customer base | | | |
| PMF-2 | Premium engagement >15% and growing | | | |
| PMF-3 | WTP >= price for >30% of target | | | |
| PMF-4 | >60% wants >60% of components | | | |
| PMF-5 | Evidence of fragmented consumption | | | |
| PMF-6 | 3-10% entry; 1-5% top tier | | | |
| PMF-7 | Brand perceived as credible in premium | | | |

**Dimension 1 Score:** ___/5
**Gate status:** [PASS / FAIL / CONDITIONAL]
**Key findings:**

[Write 2-3 sentences summarizing the product-market fit assessment.]

---

## Dimension 2: Pricing Adequacy

**Core question:** Are the price points justified by perceived value and supported by WTP?

**Gate:** If BVR < 1.0 at any tier --> **Redesign pricing**. If tier gaps disproportionate --> **Restructure tiers**.

### 2.1 Tier Structure

| Tier | Name | Price | Standalone Value Sum | BVR | Source |
|------|------|-------|---------------------|-----|--------|
| Entry | | | | | |
| Middle | | | | | |
| Top | | | | | |

### 2.2 Data Collection

| ID | Criterion | Data / Finding | Source | Method Used |
|----|-----------|---------------|--------|-------------|
| PRC-1 | Bundle Value Ratio per tier | | | |
| PRC-2 | Perceived price-value ratio (survey) | | | |
| PRC-3 | Tier gap architecture (price steps) | | | |
| PRC-4 | Anchoring effectiveness (% choosing middle) | | | |
| PRC-5 | Price vs. market average | | | |
| PRC-6 | Affordability (price as % of budget) | | | |
| PRC-7 | Price floor clearance (cost-plus) | | | |
| PRC-8 | Behavioral pricing coherence (decoy effect) | | | |

### 2.3 Price Sensitivity Research (if conducted)

| Method | Sample Size | Key Output | Result | Source |
|--------|-------------|------------|--------|--------|
| Van Westendorp PSM | | OPP / PMC / PME / IDP | | |
| Choice-Based Conjoint | | Part-worth utilities | | |
| Gabor-Granger | | Demand curve | | |
| MaxDiff | | Component ranking | | |
| A/B Test | | Conversion by price point | | |

### 2.4 Assessment

| ID | Threshold | Pass / Fail | Score (1-5) | Notes |
|----|-----------|-------------|-------------|-------|
| PRC-1 | BVR >1.5x all tiers; >2.0x preferred | | | |
| PRC-2 | Perceived value / price > 1.0 | | | |
| PRC-3 | No >50% jumps without matching value | | | |
| PRC-4 | ~66% choosing middle tier | | | |
| PRC-5 | Premium index justified by value | | | |
| PRC-6 | *Calibrate:* <___% of relevant budget | | | |
| PRC-7 | Price > cost floor at all tiers | | | |
| PRC-8 | Entry tier drives migration to middle | | | |

**Dimension 2 Score:** ___/5
**Gate status:** [PASS / FAIL / CONDITIONAL]
**Key findings:**

[Write 2-3 sentences summarizing pricing adequacy.]

---

## Dimension 3: Bundle Composition

**Core question:** Is each component a Leader, Filler, or Killer? What is the dead weight ratio?

**Gate:** If no clear Leader (BND-1) --> **Redesign bundle**. If dead weight >40% (BND-2) --> **Remove Killers or add swappability**. If access constraints >30% (BND-5) --> **Adapt for secondary markets**.

### 3.1 Component Classification

| # | Component | Role (Leader / Filler / Killer) | Standalone Price | Marginal Cost | Usage Rate (%) | Source |
|---|-----------|-------------------------------|-----------------|---------------|----------------|--------|
| 1 | | | | | | |
| 2 | | | | | | |
| 3 | | | | | | |
| 4 | | | | | | |
| 5 | | | | | | |
| 6 | | | | | | |
| 7 | | | | | | |
| 8 | | | | | | |

**Ideal composition:** 2-3 Leaders, 3-5 Fillers, 0 Killers.

### 3.2 Dead Weight and Dilution Analysis

| Component | Used by <20% of customers? | Removing it increases WTP? | Classification | Action |
|-----------|---------------------------|---------------------------|----------------|--------|
| | | | | |
| | | | | |
| | | | | |

**Dead weight ratio:** ___% (threshold: <40%)

### 3.3 Data Collection

| ID | Criterion | Data / Finding | Source | Method Used |
|----|-----------|---------------|--------|-------------|
| BND-1 | Leader identification (conjoint part-worth) | | | |
| BND-2 | Dead weight ratio | | | |
| BND-3 | Dilution risk (removal increases WTP?) | | | |
| BND-4 | Cross-subsidy balance | | | |
| BND-5 | Access constraint ratio | | | |
| BND-6 | Complementarity between components | | | |
| BND-7 | Customizability (can customers swap?) | | | |
| BND-8 | Switching cost creation (active services) | | | |

### 3.4 Assessment

| ID | Threshold / Red Flag | Pass / Fail | Score (1-5) | Notes |
|----|---------------------|-------------|-------------|-------|
| BND-1 | No clear leader = confused VP | | | |
| BND-2 | >40% dead weight | | | |
| BND-3 | Any component removal increases WTP | | | |
| BND-4 | Subsidized components >50% of margin | | | |
| BND-5 | >30% value tied to constrained access | | | |
| BND-6 | Low complementarity = random collection | | | |
| BND-7 | >8 fixed components = high dead weight | | | |
| BND-8 | <3 active services = weak lock-in | | | |

**Dimension 3 Score:** ___/5
**Gate status:** [PASS / FAIL / CONDITIONAL]
**Key findings:**

[Write 2-3 sentences summarizing bundle composition quality.]

---

## Dimension 4: Competitive Positioning

**Core question:** How does this sit in the market, and how defensible is it?

**Gate:** If <2 defensible components (CMP-5) and <6 months to imitation (CMP-3) --> **Rethink differentiation**.

### 4.1 Data Collection

| ID | Criterion | Data / Finding | Source | Method Used |
|----|-----------|---------------|--------|-------------|
| CMP-1 | Uniqueness (features not replicable in 6mo) | | | |
| CMP-2 | Competitive BVR comparison | | | |
| CMP-3 | Time to competitive imitation | | | |
| CMP-4 | Competitive response risk (game theory) | | | |
| CMP-5 | Defensible differentiation (exclusive components) | | | |
| CMP-6 | Segment ownership (brand perception) | | | |
| CMP-7 | Cross-competitive set (bundle vs. standalone alts) | | | |

### 4.2 Assessment

| ID | Threshold | Pass / Fail | Score (1-5) | Notes |
|----|-----------|-------------|-------------|-------|
| CMP-1 | >3 features not replicable within 6 months | | | |
| CMP-2 | Own BVR >= competitor BVR | | | |
| CMP-3 | >12 months for full replication | | | |
| CMP-4 | Competitors match some but not all dimensions | | | |
| CMP-5 | At least 2 exclusive components | | | |
| CMP-6 | Recognized as premium leader | | | |
| CMP-7 | Bundle price < sum of best standalone alternatives | | | |

**Dimension 4 Score:** ___/5
**Gate status:** [PASS / FAIL / CONDITIONAL]
**Key findings:**

[Write 2-3 sentences summarizing competitive positioning.]

---

## Dimension 5: Financial Viability

**Core question:** Do the unit economics work after cannibalization and cross-subsidy?

**Gate:** If unit economics negative at all tiers (FIN-2) --> **Reprice or reduce costs**. If net revenue negative after cannibalization (FIN-3) --> **Redesign migration**. If stress test fails (FIN-7) --> **Build cost buffers**.

### 5.1 Unit Economics per Tier

| Metric | Entry Tier | Middle Tier | Top Tier | Source |
|--------|-----------|-------------|----------|--------|
| Price (per period) | | | | |
| COGS per customer | | | | |
| Partner/licensing costs | | | | |
| Delivery/support costs | | | | |
| **Gross margin per customer** | | | | |
| Base product RPC (pre-premium) | | | | |
| **Revenue uplift** | | | | |

### 5.2 Cannibalization Model

| Metric | Value | Source |
|--------|-------|--------|
| Projected premium customers (12mo) | | |
| Of which: migrating from existing products | | |
| Of which: net new customers | | |
| Lost standalone revenue per migrating customer | | |
| Incremental revenue per premium customer | | |
| **Net incremental revenue** | | |

### 5.3 Stress Test

| Scenario | Revenue Impact | Margin Impact | Still Viable? | Source |
|----------|---------------|---------------|---------------|--------|
| Partner costs +20% | | | | |
| Customer growth -30% | | | | |
| Combined: costs +20%, growth -30% | | | | |

### 5.4 Data Collection

| ID | Criterion | Data / Finding | Source | Method Used |
|----|-----------|---------------|--------|-------------|
| FIN-1 | Revenue per customer uplift | | | |
| FIN-2 | Unit economics (margin per customer) | | | |
| FIN-3 | Cannibalization rate (net incremental revenue) | | | |
| FIN-4 | Partner/licensing cost ratio | | | |
| FIN-5 | Lifetime value: premium CLV vs. base CLV | | | |
| FIN-6 | Time to break-even (months) | | | |
| FIN-7 | Cross-subsidy stress test result | | | |
| FIN-8 | Scale economics (breakeven customer count) | | | |
| FIN-9 | Standalone product cannibalization | | | |

### 5.5 Assessment

| ID | Threshold | Pass / Fail | Score (1-5) | Notes |
|----|-----------|-------------|-------------|-------|
| FIN-1 | *Calibrate:* meaningful uplift | | | |
| FIN-2 | Positive margin at all tiers | | | |
| FIN-3 | Net incremental revenue > 0 | | | |
| FIN-4 | Partner costs <30% of revenue delta | | | |
| FIN-5 | Premium CLV >= 2x base CLV | | | |
| FIN-6 | *Calibrate:* break-even within ___mo | | | |
| FIN-7 | Margin-positive under stress | | | |
| FIN-8 | Breakeven count achievable in 12mo | | | |
| FIN-9 | Net revenue positive including losses | | | |

**Dimension 5 Score:** ___/5
**Gate status:** [PASS / FAIL / CONDITIONAL]
**Key findings:**

[Write 2-3 sentences summarizing financial viability. Note: churn reduction may be the primary financial lever, not direct revenue uplift.]

---

## Dimension 6: Customer Experience

**Core question:** Will this improve satisfaction and reduce churn?

**Gate:** If WTP validation fails (CX-5: price > 80% WTP) --> **Reprice**. If disappointment risk high (CX-8: >15% drop) --> **Improve component quality**.

### 6.1 Data Collection

| ID | Criterion | Data / Finding | Source | Method Used |
|----|-----------|---------------|--------|-------------|
| CX-1 | Churn reduction (premium vs. base) | | | |
| CX-2 | NPS improvement (premium vs. standard) | | | |
| CX-3 | CSAT score across components | | | |
| CX-4 | Feature utilization rate | | | |
| CX-5 | WTP validation (conversion vs. projected) | | | |
| CX-6 | Support experience (resolution time delta) | | | |
| CX-7 | Component engagement (vs. standalone benchmarks) | | | |
| CX-8 | Disappointment risk (post vs. pre value perception) | | | |

### 6.2 Assessment

| ID | Threshold | Pass / Fail | Score (1-5) | Notes |
|----|-----------|-------------|-------------|-------|
| CX-1 | Premium churn < 50% of base churn | | | |
| CX-2 | Premium NPS > 30 (*calibrate per industry*) | | | |
| CX-3 | CSAT > 80% | | | |
| CX-4 | >60% of features used by >60% of customers | | | |
| CX-5 | Price < 80% of measured WTP | | | |
| CX-6 | Premium resolution < 50% of standard | | | |
| CX-7 | Engagement tracks standalone benchmarks | | | |
| CX-8 | <15% value perception drop at 3 months | | | |

**Dimension 6 Score:** ___/5
**Gate status:** [PASS / FAIL / CONDITIONAL]
**Key findings:**

[Write 2-3 sentences summarizing customer experience impact.]

---

## Dimension 7: Market Reach

**Core question:** Does this work beyond the primary market/segment?

**Gate:** If access coverage <70% (MR-1) --> **Adapt components**. If no viable secondary markets (MR-6) --> **Consider niche launch**.

### 7.1 Market/Segment Reach Analysis

| Market / Segment | Addressable Size | Access Coverage (%) | Affordability (price as % budget) | Competitive Intensity | Delivery Ready? | Source |
|-----------------|-----------------|--------------------|---------------------------------|----------------------|----------------|--------|
| Primary: | | | | | | |
| Secondary: | | | | | | |
| Tertiary: | | | | | | |

### 7.2 Data Collection

| ID | Criterion | Data / Finding | Source | Method Used |
|----|-----------|---------------|--------|-------------|
| MR-1 | Access coverage of constrained components | | | |
| MR-2 | Segment affordability | | | |
| MR-3 | Revenue ceiling by segment | | | |
| MR-4 | Segment competitive landscape | | | |
| MR-5 | Delivery readiness by segment | | | |
| MR-6 | Demand heterogeneity across segments | | | |
| MR-7 | Product adaptation (component swaps per segment) | | | |

### 7.3 Assessment

| ID | Threshold | Pass / Fail | Score (1-5) | Notes |
|----|-----------|-------------|-------------|-------|
| MR-1 | Accessible to >70% of addressable customers | | | |
| MR-2 | *Calibrate:* within affordability threshold | | | |
| MR-3 | Price < segment revenue ceiling | | | |
| MR-4 | No dominant competitor in target segments | | | |
| MR-5 | Deliverable in all launch segments | | | |
| MR-6 | Viable in primary + secondary markets | | | |
| MR-7 | Alternatives available in non-primary segments | | | |

**Dimension 7 Score:** ___/5
**Gate status:** [PASS / FAIL / CONDITIONAL]
**Key findings:**

[Write 2-3 sentences summarizing market reach assessment.]

---

## Competitive Comparison

### Feature-by-Feature Comparison

| Feature / Component | Our Product | Competitor A | Competitor B | Competitor C | Notes |
|--------------------|-------------|-------------|-------------|-------------|-------|
| | | | | | |
| | | | | | |
| | | | | | |
| | | | | | |
| | | | | | |
| **Bundle Price** | | | | | |
| **BVR** | | | | | |
| **Standalone Value Sum** | | | | | |

### Competitive KPI Comparison

| KPI | Our Product | Competitor A | Competitor B | Competitor C | Source |
|-----|-------------|-------------|-------------|-------------|--------|
| Price (entry tier) | | | | | |
| Price (middle tier) | | | | | |
| Price (top tier) | | | | | |
| BVR (best tier) | | | | | |
| # of components | | | | | |
| # of Leaders | | | | | |
| Exclusive components | | | | | |
| Estimated time to imitate | | | | | |
| NPS (if available) | | | | | |
| Churn rate (if available) | | | | | |

### Competitive Positioning Summary

[Write 3-5 sentences: Where does this product sit in the competitive landscape? What is defensible? What is vulnerable?]

---

## Risk Assessment

| Risk | Probability (1-5) | Impact (1-5) | Score (P x I) | Mitigation | Source |
|------|-------------------|-------------|----------------|------------|--------|
| Low adoption / demand miss | | | | | |
| Excessive cannibalization | | | | | |
| Partner failure / cost escalation | | | | | |
| Competitive response | | | | | |
| Brand perception damage | | | | | |
| Regulatory / compliance | | | | | |
| Delivery quality shortfall | | | | | |
| Access constraint backlash | | | | | |
| Dead weight > 40% | | | | | |
| [Add product-specific risks] | | | | | |

**Top 3 risks by score:**
1. [Risk] -- [P x I] -- [Mitigation summary]
2. [Risk] -- [P x I] -- [Mitigation summary]
3. [Risk] -- [P x I] -- [Mitigation summary]

---

## Recommendations

### Strengths (Preserve)

| # | Strength | Supporting Evidence | Criteria IDs |
|---|----------|-------------------|--------------|
| 1 | | | |
| 2 | | | |
| 3 | | | |

### Weaknesses (Address Before Launch)

| # | Weakness | Severity (H/M/L) | Recommended Action | Criteria IDs |
|---|----------|-------------------|-------------------|--------------|
| 1 | | | | |
| 2 | | | | |
| 3 | | | | |

### Opportunities (Explore)

| # | Opportunity | Potential Impact | Next Step | Criteria IDs |
|---|------------|-----------------|-----------|--------------|
| 1 | | | | |
| 2 | | | | |
| 3 | | | | |

### Critical Path to Launch

| # | Action Item | Owner | Deadline | Blocking? | Status |
|---|------------|-------|----------|-----------|--------|
| 1 | | | | | |
| 2 | | | | | |
| 3 | | | | | |

---

## Condensed Scorecard (1-Page Summary)

> **Product:** [Name]
> **Date:** [YYYY-MM-DD]
> **Overall Verdict:** [Strong Go / Conditional Go / Redesign Required / No-Go]

### Dimension Scores

| Dimension | Weight | Score (1-5) | Weighted Score | Go / No-Go | Gate Status | Rationale |
|-----------|--------|-------------|---------------|-------------|-------------|-----------|
| 1. Product-Market Fit | 15% | | | | | |
| 2. Pricing Adequacy | 20% | | | | | |
| 3. Bundle Composition | 10% | | | | | |
| 4. Competitive Positioning | 15% | | | | | |
| 5. Financial Viability | 25% | | | | | |
| 6. Customer Experience | 10% | | | | | |
| 7. Market Reach | 5% | | | | | |
| **Total** | **100%** | | **___** | | | |

### Decision Thresholds

| Weighted Total | Decision | Action |
|---------------|----------|--------|
| >= 4.0 | **Strong Go** | Proceed to launch planning |
| 3.0 - 3.9 | **Conditional Go** | Address gaps listed below before launch |
| 2.0 - 2.9 | **Redesign Required** | Fundamental changes needed (see recommendations) |
| < 2.0 | **No-Go** | Do not launch in current form |

### Gate Failures (if any)

| Dimension | Gate | Status | Required Action |
|-----------|------|--------|----------------|
| | | | |

### Top 3 Strengths

1. [One-liner]
2. [One-liner]
3. [One-liner]

### Top 3 Risks

1. [Risk] -- P:_ x I:_ = _ -- [Mitigation]
2. [Risk] -- P:_ x I:_ = _ -- [Mitigation]
3. [Risk] -- P:_ x I:_ = _ -- [Mitigation]

### Recommended Next Steps

1. [Action]
2. [Action]
3. [Action]

---

*Template version 1.0. Based on Product Appraisal Methodology (7 dimensions, 54 criteria). All criteria IDs (PMF-1 through MR-7) reference METHODOLOGY.md for full definitions, thresholds, and theoretical foundations.*
