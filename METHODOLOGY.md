# Product Appraisal Methodology

A universal framework for evaluating complex products, bundles, and subscription offerings. Derived from bundling theory, behavioral pricing research, and consulting frameworks, then generalized from telecom-specific application to work across industries.

---

## Table of Contents

1. [Evaluation Dimensions](#1-evaluation-dimensions)
2. [Key Metrics and How They Are Calculated](#2-key-metrics-and-how-they-are-calculated)
3. [Benchmarks and Sources](#3-benchmarks-and-sources)
4. [Evaluation Logic Flow](#4-evaluation-logic-flow)
5. [Theoretical Foundations](#5-theoretical-foundations)
6. [Research Methods](#6-research-methods)
7. [Decision Framework](#7-decision-framework)

---

## 1. Evaluation Dimensions

The appraisal operates across seven dimensions, applied sequentially. Each dimension has criteria with IDs for traceability, assessment methods, and pass thresholds.

### Dimension 1: Product-Market Fit

**Core question:** Does the target segment exist, and does the bundle address their actual needs?

| ID | Criterion | Method | Threshold |
|----|-----------|--------|-----------|
| PMF-1 | Addressable target segment size | Customer data segmented by income, behavior, engagement | >5% of total customer base |
| PMF-2 | Demonstrated demand for premium tier | Market research, existing premium adoption trends | Premium engagement penetration >15% and growing |
| PMF-3 | Willingness to pay at proposed prices | Conjoint analysis / Van Westendorp | WTP >= proposed price for >30% of target |
| PMF-4 | Service-need alignment | Survey: which components does the target want? | >60% of target wants >60% of components |
| PMF-5 | Current unmet need | Analysis of fragmented service consumption | Evidence of fragmented premium consumption |
| PMF-6 | Conversion funnel viability | Funnel modeling: total -> engaged -> eligible -> converters | 3-10% of eligible at entry tier; 1-5% at top tier |
| PMF-7 | Brand permission for premium | Brand assessment: Image, Communication, Execution | Brand perceived as credible in premium space |

### Dimension 2: Pricing Adequacy

**Core question:** Are the price points justified by perceived value and supported by WTP?

| ID | Criterion | Method | Threshold |
|----|-----------|--------|-----------|
| PRC-1 | Bundle Value Ratio (BVR) | Sum of standalone prices / bundle price | >1.5x at all tiers; >2.0x preferred |
| PRC-2 | Price-value ratio (perceived) | Survey: perceived value / actual price | >1.0 |
| PRC-3 | Tier gap architecture | Price step analysis between tiers | Proportional gaps; avoid >50% jumps without matching value delta |
| PRC-4 | Anchoring effectiveness | Good-Better-Best dynamics | ~66% target on middle tier |
| PRC-5 | Price vs. market average | Premium price / market average | Premium index justified by bundle value |
| PRC-6 | Affordability for target segment | Price as % of discretionary budget | *Calibrate per industry* |
| PRC-7 | Price floor clearance | Cost-plus analysis: total cost per customer | Price > cost floor at all tiers |
| PRC-8 | Behavioral pricing coherence | Decoy effect validation | Entry tier drives migration to middle tier |

### Dimension 3: Bundle Composition

**Core question:** Is each component a Leader, Filler, or Killer? What is the dead weight ratio?

Components are classified using the **Leaders / Fillers / Killers** framework:

| Role | Definition | Target Count |
|------|-----------|-------------|
| **Leader** | High perceived value, drives purchase intent | 2-3 per bundle |
| **Filler** | Adds perceived value at low marginal cost | 3-5 per bundle |
| **Killer** | Reduces WTP, confuses the offer, attracts wrong customers | 0 per bundle |

| ID | Criterion | Method | Red Flag |
|----|-----------|--------|----------|
| BND-1 | Leader identification | Conjoint part-worth analysis | No clear leader = confused value proposition |
| BND-2 | Dead weight ratio | Usage tracking: components used by <20% within 3 months | >40% dead weight |
| BND-3 | Dilution risk | Test: does removing a component increase WTP? | Any component whose removal increases WTP |
| BND-4 | Cross-subsidy balance | Cost vs. revenue contribution per component | Subsidized components >50% of incremental margin |
| BND-5 | Access constraint ratio | % of components requiring specific conditions (geography, platform, credentials) | >30% of value tied to access-constrained components |
| BND-6 | Complementarity | Do components reinforce each other? | Low complementarity = "random collection" |
| BND-7 | Customizability | Can customers swap components? | Fixed bundles with >8 components = high dead weight certainty |
| BND-8 | Switching cost creation | Active services increasing exit friction | <3 active services = weak lock-in |

**Key insight from research:** Dead weight is not purely negative. Unused components can contribute "option value" -- perceived value from availability alone. However, if low-value components trigger the dilution effect (Shaddy & Fishbach, 2017), they reduce total bundle WTP. The test is BND-3: does removing the component increase or decrease WTP?

### Dimension 4: Competitive Positioning

**Core question:** How does this sit in the market, and how defensible is it?

| ID | Criterion | Method | Threshold |
|----|-----------|--------|-----------|
| CMP-1 | Uniqueness | Feature-by-feature competitor comparison | >3 features not replicable within 6 months |
| CMP-2 | Competitive BVR | BVR analysis applied to competitor offerings | Own BVR >= competitor BVR |
| CMP-3 | Time to imitation | Competitor capability assessment | >12 months for full replication |
| CMP-4 | Response risk | Game theory: likely competitor reactions | Competitors can match some but not all dimensions |
| CMP-5 | Defensible differentiation | Exclusive partnerships, proprietary elements | At least 2 exclusive components |
| CMP-6 | Segment ownership | Brand perception in premium segment | Recognized as premium leader |
| CMP-7 | Cross-competitive set | Bundle vs. standalone alternatives | Bundle price < sum of best individual alternatives |

### Dimension 5: Financial Viability

**Core question:** Do the unit economics work after cannibalization and cross-subsidy?

| ID | Criterion | Method | Threshold |
|----|-----------|--------|-----------|
| FIN-1 | Revenue per customer uplift | Premium vs. base revenue per customer | *Calibrate per industry* |
| FIN-2 | Unit economics | Revenue minus all costs per customer | Positive margin at all tiers |
| FIN-3 | Cannibalization rate | Migration from existing offerings vs. net new | Net incremental revenue > 0 |
| FIN-4 | Partner cost ratio | Partner + licensing costs / premium revenue delta | <30% (practitioner guidance) |
| FIN-5 | Lifetime value premium | Premium CLV vs. base CLV | Premium CLV >= 2x base (practitioner target) |
| FIN-6 | Time to break-even | Cumulative revenue > cumulative cost | *Calibrate per industry* |
| FIN-7 | Cross-subsidy sustainability | Stress test: costs +20%, growth -30% | Model remains margin-positive under stress |
| FIN-8 | Scale economics | Minimum customer count for viability | Breakeven count achievable within 12 months |
| FIN-9 | Standalone product cannibalization | Lost standalone revenue from bundling | Net revenue positive including losses |

**Critical financial insight:** Bundling's primary financial lever is often churn reduction, not direct revenue uplift. Even thin or negative per-customer margins can be justified if churn reduction delivers sufficient lifetime value improvement. Model churn reduction explicitly.

### Dimension 6: Customer Experience

**Core question:** Will this improve satisfaction and reduce churn?

| ID | Criterion | Method | Threshold |
|----|-----------|--------|-----------|
| CX-1 | Churn reduction | Premium vs. base churn rate | Premium churn < 50% of base |
| CX-2 | NPS improvement | Premium vs. standard NPS | Premium NPS significantly above base (*calibrate per industry*) |
| CX-3 | CSAT | Satisfaction survey across components | >80% |
| CX-4 | Feature utilization | % of features actively used | >60% of features by >60% of customers (practitioner target) |
| CX-5 | WTP validation | Actual conversion vs. projected WTP | Price < 80% of measured WTP |
| CX-6 | Support experience | Premium vs. standard resolution time | Premium < 50% of standard |
| CX-7 | Component engagement | Usage frequency per component | Engagement tracks standalone benchmarks |
| CX-8 | Disappointment risk | Post-purchase vs. pre-purchase value perception | <15% value perception drop after 3 months |

### Dimension 7: Market Reach

**Core question:** Does this work beyond the primary market/segment?

| ID | Criterion | Method | Threshold |
|----|-----------|--------|-----------|
| MR-1 | Access coverage | Availability of constrained components | Accessible to >70% of addressable customers |
| MR-2 | Segment affordability | Price as % of segment's relevant budget | *Calibrate per industry* |
| MR-3 | Revenue ceiling | Max sustainable revenue per customer by segment | Price < segment ceiling |
| MR-4 | Segment competition | Premium competition by segment | No dominant competitor in target segments |
| MR-5 | Delivery readiness | Premium experience deliverable across segments | Deliverable in all launch segments |
| MR-6 | Demand heterogeneity | Segment sizing across markets/demographics | Viable in primary + secondary markets |
| MR-7 | Product adaptation | Can components be swapped per segment? | Alternatives available in non-primary segments |

---

## 2. Key Metrics and How They Are Calculated

### Revenue Metrics

| Metric | Formula | Purpose |
|--------|---------|---------|
| Revenue Per Customer (RPC) | Total product revenue / avg customers | Core revenue health indicator |
| Revenue uplift | Premium RPC - Base RPC | Measures premium's incremental value |
| Blended revenue impact | (Premium subs x Premium RPC + Base subs x Base RPC) / Total subs | Net effect including cannibalization |
| Gross margin per customer | RPC - COGS per customer | Unit economics viability |

### Bundle Economics Metrics

| Metric | Formula | Purpose |
|--------|---------|---------|
| Bundle Value Ratio (BVR) | Sum of standalone prices / bundle price | Core bundle value indicator |
| Cross-subsidy efficiency | High-margin component margin - Low-margin component costs | Sustainability of internal subsidies |
| Partner cost ratio | (Partner + licensing costs per customer) / Premium revenue delta | Partner economics health |
| Incremental revenue | Bundle RPC - Lost standalone revenue per migrating customer | Net cannibalization impact |
| Multi-component usage rate | Customers using 3+ components / Total bundle customers | Bundle engagement depth |

### Customer Metrics

| Metric | Formula | Purpose |
|--------|---------|---------|
| CLV (Customer Lifetime Value) | RPC x Gross Margin % x Avg Customer Lifespan (months/years) | Long-term customer value |
| CAC (Customer Acquisition Cost) | Total acquisition spend / New customers | Acquisition efficiency |
| CAC Payback | CAC / (Monthly RPC x Gross Margin %) | Months to recover acquisition cost |
| Churn rate | Customers lost / Total customers per period | Retention health |
| NPS | % Promoters - % Detractors | Customer advocacy |

### Segment Metrics

| Metric | Formula | Purpose |
|--------|---------|---------|
| Penetration rate | Premium customers / Total customer base | Adoption depth |
| Upgrade rate | Customers upgrading / Eligible base per period | Migration velocity |
| Share of wallet | Company revenue / Customer's total category spend | Competitive position per customer |
| Dead weight ratio | Components with <20% monthly usage / Total components | Bundle efficiency |

---

## 3. Benchmarks and Sources

Every external benchmark is tagged: **[Verified]** with URL, **[Practitioner Guidance]** for unverified but reasonable heuristics, or **[Calibrate]** for industry-dependent thresholds.

### Bundle Pricing Theory

| Benchmark | Value | Source | Status |
|-----------|-------|--------|--------|
| Mixed bundling profit superiority | Strictly increases profits vs. pure bundling | McAfee, McMillan & Whinston (1989), *QJE* | [Verified] [URL](https://academic.oup.com/qje/article-abstract/104/2/371/1854649) |
| Bundle variance reduction | Bundling reduces valuation variance, enabling better surplus extraction | Schmalensee (1984), *Journal of Business* | [Verified] [URL](https://www.jstor.org/stable/2352937) |
| Unbundling profit loss | ~10% profit decrease from unbundling (telecom empirical) | Luo (2023), *RAND Journal of Economics* | [Verified] [URL](https://onlinelibrary.wiley.com/doi/10.1111/1756-2171.12437) |
| Bundle framing effect | Multi-unit framing ("3 for $5") boosts sales ~32% | Wansink et al. (1998), *Journal of Marketing Research* | [Verified] |
| Heterogeneous bundles | Diverse component bundles support higher prices | Xia Wei et al. (2025), *Journal of Travel Research* | [Verified] [URL](https://journals.sagepub.com/doi/10.1177/00472875231222263) |

### Behavioral Pricing

| Benchmark | Value | Source | Status |
|-----------|-------|--------|--------|
| GBB middle tier selection | ~66% of customers choose middle option | Multiple sources; HBR (Mohammed, 2018) | [Verified] [URL](https://hbr.org/2018/09/the-good-better-best-approach-to-pricing) |
| Removing middle tier revenue loss | ~50% revenue reduction | Industry sources (multiple) | [Practitioner Guidance] |
| Price anchoring effect | Displaying component prices alongside bundle price significantly increases perceived value | Tversky & Kahneman (1974) anchoring theory | [Verified -- mechanism. No verified specific % for bundles.] |
| Dilution effect | Low-value components can reduce total bundle WTP | Shaddy & Fishbach (2017), *Journal of Marketing Research* | [Verified] [URL](https://www.anderson.ucla.edu/documents/areas/fac/marketing/Seminars/Fall%202017/Shaddy%20%20Fishbach%20-%20How%20Bundling%20Affects%20Valuation%20(job%20market%20paper).pdf) |

### Churn and Retention

| Benchmark | Value | Source | Status |
|-----------|-------|--------|--------|
| Bundle churn reduction (entertainment) | 59% less likely to churn | Ampere Analysis (Disney+ bundle data) | [Verified] [URL](https://www.nexttv.com/news/disney-bundlers-59-less-likely-to-churn-research-company-says-chart) |
| Bundle churn reduction (broadband) | Modest: 6.93 vs. 6.15 years retention | Prince & Greenstein (2014) | [Verified] [URL](https://host.kelley.iu.edu/riharbau/RePEc/iuk/wpaper/bepp2011-05-prince-greenstein.pdf) |
| Multi-product churn reduction | 25-35% lower churn (conservative range) | Industry consensus; varies by design | [Practitioner Guidance] |
| Acquisition vs. retention cost | 5-25x more expensive to acquire than retain | HBR (Gallo, 2014) | [Verified] [URL](https://hbr.org/2014/10/the-value-of-keeping-the-right-customers) |
| Personalization revenue lift | 10-15% revenue lift; 10-30% marketing ROI improvement | McKinsey | [Verified] [URL](https://www.mckinsey.com/capabilities/growth-marketing-and-sales/our-insights/the-value-of-getting-personalization-right-or-wrong-is-multiplying) |

### Bundle Composition

| Benchmark | Value | Source | Status |
|-----------|-------|--------|--------|
| Dead weight threshold | <40% of components should be dead weight | Simon-Kucher (2024 Telco Study), derived from "~60% respond positively" | [Verified -- inverse] [URL](https://www.simon-kucher.com/en/insights/2024-telco-growth-strategies-brand-portfolio-and-pricing-insights) |
| Post-purchase disappointment | ~15% value perception drop when expectations unmet | Hospitality bundle research | [Partially Verified -- domain-specific] |
| Over-provisioning waste | 20-30% of feature costs wasted on unvalued capabilities | Industry analysis (TEM, cloud) | [Practitioner Guidance] |

### Conversion and Adoption

| Benchmark | Value | Source | Status |
|-----------|-------|--------|--------|
| Freemium conversion (self-serve) | 3-5% average; 6-8% top performers | First Page Sage (2026) | [Verified] [URL](https://firstpagesage.com/seo-blog/saas-freemium-conversion-rates/) |
| Freemium conversion (sales-assisted) | 5-7% average; 10-15% top performers | First Page Sage (2026) | [Verified] [URL](https://firstpagesage.com/seo-blog/saas-freemium-conversion-rates/) |
| Streaming aggregation trend | ~4 SVOD services per household peak; 20% -> 25% telco-mediated by 2028 | Deloitte TMT Predictions 2025 | [Verified] [URL](https://www.deloitte.com/us/en/insights/industry/technology/technology-media-and-telecom-predictions/2025/tmt-predictions-video-streaming-bundles-bigger-than-ever.html) |

### Financial Targets

| Benchmark | Value | Source | Status |
|-----------|-------|--------|--------|
| BVR threshold | >1.5x adequate; >2.0x strong | Not externally validated | [Practitioner Guidance] |
| Premium CLV multiple | 2-4x base CLV | Industry practice | [Practitioner Guidance] |
| Partner cost ratio | <30% of premium revenue delta | Not externally validated | [Practitioner Guidance] |
| Feature utilization target | >60% of features used | Product management practice | [Practitioner Guidance] |
| Premium churn target | <50% of base churn | Not externally validated | [Practitioner Guidance] |

---

## 4. Evaluation Logic Flow

The seven dimensions are evaluated sequentially. Each stage builds on the previous and can trigger early exit.

```
1. DEMAND VALIDATION (PMF-1 through PMF-7)
   Does anyone want this? Is the segment real?

   GATE: If PMF-1 fails (<5% addressable) or PMF-3 fails (WTP below price) -> No-Go
   |
   v
2. PRICE-VALUE MAPPING (PRC-1 through PRC-8)
   Is the price justified by perceived value?

   GATE: If BVR < 1.0 at any tier -> Redesign pricing
   GATE: If tier gaps are disproportionate -> Restructure tiers
   |
   v
3. BUNDLE DECOMPOSITION (BND-1 through BND-8)
   Is each component earning its place?

   GATE: If no clear Leader (BND-1) -> Redesign bundle
   GATE: If dead weight >40% (BND-2) -> Remove Killers or add swappability
   GATE: If access constraints >30% (BND-5) -> Adapt for secondary markets
   |
   v
4. COMPETITIVE ANALYSIS (CMP-1 through CMP-7)
   Can it survive in the market?

   GATE: If <2 defensible components (CMP-5) and <6 months to imitation (CMP-3) -> Rethink differentiation
   |
   v
5. FINANCIAL MODELING (FIN-1 through FIN-9)
   Does the math work?

   GATE: If unit economics negative at all tiers (FIN-2) -> Reprice or reduce costs
   GATE: If net revenue negative after cannibalization (FIN-3) -> Redesign migration
   GATE: If stress test fails (FIN-7) -> Build cost buffers
   |
   v
6. CUSTOMER IMPACT (CX-1 through CX-8)
   Will it improve the customer relationship?

   GATE: If WTP validation fails (CX-5: price > 80% WTP) -> Reprice
   GATE: If disappointment risk high (CX-8: >15% drop) -> Improve component quality
   |
   v
7. MARKET REACH (MR-1 through MR-7)
   Where can this actually launch?

   GATE: If access coverage <70% (MR-1) -> Adapt components for broader reach
   GATE: If no viable secondary markets (MR-6) -> Consider niche launch strategy
   |
   v
8. GO / NO-GO / REDESIGN DECISION
   Weighted scoring across all 7 dimensions
```

---

## 5. Theoretical Foundations

### Why Bundling Works

**Variance reduction** (Schmalensee, 1984): Consumer valuations for a bundle have lower variance than valuations for individual components. This makes the consumer population more homogeneous in their bundle valuation, enabling more effective surplus extraction through a single price point. The effect strengthens with more components (law of large numbers).

**Price discrimination** (Adams & Yellen, 1976): Bundling sorts customers by their reservation price combinations for multiple goods. Mixed bundling (offering both bundles and standalone components) captures surplus from both high-value-for-all customers (who buy the bundle) and high-value-for-one customers (who buy standalone).

**Mixed bundling optimality** (McAfee, McMillan & Whinston, 1989): Mixed bundling almost always strictly increases profits compared to pure bundling or pure component selling. This is the strongest theoretical result in bundling theory. **Practical implication: always offer the bundle AND standalone components simultaneously.**

### How Consumers Evaluate Bundles

**Highest-priced item heuristic**: Consumers mentally evaluate bundles based on the most valued component. Everything else is processed as a bonus. The Leader component determines perception.

**Dilution effect** (Shaddy & Fishbach, 2017): Consumers perceive bundles as gestalt units. Adding low-value components can reduce total WTP because consumers resist altering the "whole." Removing a component increases perceived loss (consumers demand more compensation); adding a component yields diminished perceived gain (consumers offer less WTP).

**Perceived heterogeneity** (Xia Wei et al., 2025): Bundles with diverse components (high perceived heterogeneity) support higher prices because variety itself signals value.

**Option value**: Unused bundle components can still contribute perceived value through availability -- the knowledge they could be used if needed. This partially offsets dead weight concerns for premium/aspirational features.

### Key Behavioral Pricing Principles

**Good-Better-Best architecture**: Three-tier pricing exploits the compromise effect (~66% choose the middle option). The top tier serves as a price anchor making the middle tier feel reasonable. The bottom tier serves as a decoy making the middle tier feel like better value.

**Price anchoring** (Tversky & Kahneman, 1974): The first price consumers encounter anchors their fairness assessment. Displaying individual component prices before the bundle price creates a powerful anchoring effect that increases perceived bundle value.

**Bundle framing** (Wansink et al., 1998): Presenting items as a bundle ("3 for $5") rather than individually ("$1.67 each") boosts sales by approximately 32%, even at mathematically equivalent prices. The framing itself creates perceived value.

---

## 6. Research Methods

### Recommended Sequence for Price Validation

| Stage | Method | Purpose | Sample Size |
|-------|--------|---------|-------------|
| 1 | **Van Westendorp PSM** | Establish acceptable price range | 500-1,000 |
| 2 | **Choice-Based Conjoint (CBC)** | Decompose WTP across bundle components | 2,000+ |
| 3 | **Gabor-Granger** | Fine-tune specific price points within range | 500-1,000 |
| 4 | **MaxDiff Scaling** | Rank component importance (must-have vs. dead weight) | 500+ |
| 5 | **A/B Testing** | In-market validation | Live traffic |
| 6 | **Pilot Launch** | Limited geography/segment before full rollout | Real customers |

### Van Westendorp Price Sensitivity Meter
Four questions identifying the acceptable price range:
1. At what price is this too expensive to consider?
2. At what price is this expensive but worth considering?
3. At what price is this a bargain?
4. At what price is this so cheap you'd question quality?

Outputs: Optimal Price Point (OPP), Point of Marginal Cheapness (PMC), Point of Marginal Expensiveness (PME), Indifference Price Point (IDP).

Source: Peter Van Westendorp, 1976. [Wikipedia](https://en.wikipedia.org/wiki/Van_Westendorp's_Price_Sensitivity_Meter)

### Choice-Based Conjoint Analysis
Decomposes bundle value into component-level WTP by presenting trade-off choices between product configurations. Produces part-worth utilities, market simulators, and demand curves.

Source: Standard methodology. BCG's "Pathways Conjoint" is a telecom-specific variant. [BCG](https://www.bcg.com/publications/2014/telecommunications-pricing-pathways-conjoint-new-approach-pricing-mobile)

### Gabor-Granger Method
Sequential price presentation: if customer accepts a price, show higher; if rejects, show lower. Produces a demand curve tied to specific price points.

Source: Andre Gabor and Clive W.J. Granger, 1960s. [Sawtooth Software](https://sawtoothsoftware.com/resources/blog/posts/gabor-granger-pricing-method)

---

## 7. Decision Framework

### Go/No-Go Scoring Matrix

| Dimension | Weight | Score (1-5) | Weighted |
|-----------|--------|-------------|----------|
| Strategic fit (PMF) | 15% | | |
| Financial viability (FIN) | 25% | | |
| Customer demand (PRC + CX) | 20% | | |
| Competitive position (CMP) | 15% | | |
| Bundle composition (BND) | 10% | | |
| Market reach (MR) | 10% | | |
| Risk profile | 5% | | |
| **Total** | **100%** | | |

### Decision Thresholds

| Score | Decision | Action |
|-------|----------|--------|
| >= 4.0 | **Strong Go** | Proceed to launch planning |
| 3.0 - 3.9 | **Conditional Go** | Address specific gaps identified by failing criteria |
| 2.0 - 2.9 | **Redesign Required** | Fundamental changes needed to pricing, composition, or positioning |
| < 2.0 | **No-Go** | Do not launch in current form |

### Risk Assessment Template

For each identified risk:

| Risk | Probability (1-5) | Impact (1-5) | Score (P x I) | Mitigation |
|------|-------------------|-------------|----------------|------------|
| Low adoption | | | | |
| Excessive cannibalization | | | | |
| Partner failure | | | | |
| Competitive response | | | | |
| Brand perception damage | | | | |
| Regulatory intervention | | | | |
| Delivery quality shortfall | | | | |
| Access constraint backlash | | | | |
| Partner cost escalation | | | | |
| Dead weight > 40% | | | | |

---

## Calibration Notes

The following thresholds must be calibrated for each industry application. Default values come from telecom/subscription industry practice.

| Threshold | Default | Calibration Guidance |
|-----------|---------|---------------------|
| Revenue uplift target | +12-18% blended | SaaS: 50-100%+; physical goods: 10-20%; media: 15-30% |
| NPS premium target | >30 | Varies by industry: tech 40-60, retail 50-70, financial 20-40 |
| CAC payback | <6 months | SaaS: 12-18 months; consumer apps: 1-3 months |
| Penetration of eligible base | 10-25% | Depends on market maturity and pricing |
| Affordability threshold | <3-5% of income | Differs by product category and wallet share norms |
| Time to break-even | 6-18 months | Depends on investment level and margin structure |
| Freemium conversion | 3-5% self-serve | Sales-assisted: 5-7%; top performers: 8-15% |

---

*This methodology is a working tool. When applying to a specific industry, calibrate thresholds, select relevant case studies, and adjust KPI definitions to match industry terminology. All criteria IDs (PMF-1 through MR-7) should be referenced when documenting findings for traceability.*
