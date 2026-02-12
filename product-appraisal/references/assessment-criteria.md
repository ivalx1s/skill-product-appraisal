# Assessment Criteria Reference

Generalized criteria for evaluating complex products, bundles, and subscription offerings across seven dimensions. All criteria are industry-agnostic. Thresholds marked *calibrate per industry* require adjustment based on sector norms.

---

## Table of Contents

1. [How to Use This Reference](#how-to-use-this-reference)
2. [Dimension 1: Product-Market Fit (PMF)](#dimension-1-product-market-fit-pmf)
3. [Dimension 2: Pricing Adequacy (PRC)](#dimension-2-pricing-adequacy-prc)
4. [Dimension 3: Bundle Composition (BND)](#dimension-3-bundle-composition-bnd)
5. [Dimension 4: Competitive Positioning (CMP)](#dimension-4-competitive-positioning-cmp)
6. [Dimension 5: Financial Viability (FIN)](#dimension-5-financial-viability-fin)
7. [Dimension 6: Customer Experience (CX)](#dimension-6-customer-experience-cx)
8. [Dimension 7: Market Reach (MR)](#dimension-7-market-reach-mr)
9. [Evaluation Logic Flow](#evaluation-logic-flow)
10. [Decision Framework](#decision-framework)
11. [Calibration Guide](#calibration-guide)
12. [Source Index](#source-index)

---

## How to Use This Reference

Each criterion has:

- **ID** -- stable identifier for traceability across reports (e.g., PMF-1, FIN-3).
- **Name** -- what is being assessed.
- **Assessment Method** -- how to measure or evaluate it.
- **Pass Threshold** -- the bar to clear. Three types:
  - **Universal** -- applies across industries as stated.
  - **Calibrate per industry** -- the metric is universal but the number depends on sector. See [Calibration Guide](#calibration-guide).
  - **Red Flag** -- not a pass/fail threshold but an indicator that something is wrong.

Dimensions are evaluated sequentially. Each dimension has gate criteria that can trigger early exit (No-Go or Redesign). See [Evaluation Logic Flow](#evaluation-logic-flow).

---

## Dimension 1: Product-Market Fit (PMF)

**Core question:** Does the target segment exist, and does the product address their actual needs?

| ID | Name | Assessment Method | Pass Threshold |
|----|------|-------------------|----------------|
| PMF-1 | Addressable target segment size | Customer data segmented by income, behavior, and engagement level | >5% of total customer base |
| PMF-2 | Demonstrated demand for premium tier | Market research; existing premium or loyalty program adoption trends | Premium engagement penetration >15% and growing |
| PMF-3 | Willingness to pay at proposed prices | Conjoint analysis or Van Westendorp PSM on target segment | WTP >= proposed price for >30% of target segment |
| PMF-4 | Service-need alignment | Survey: which bundle components does the target segment actually want? | >60% of target segment wants >60% of bundle components |
| PMF-5 | Current unmet need | Analysis of target segment currently using fragmented or separate solutions | Evidence of fragmented premium consumption patterns |
| PMF-6 | Conversion funnel viability | Funnel modeling: total customers -> engaged -> eligible -> converters | 3-10% of eligible at entry tier; 1-5% at top tier. *Calibrate per industry.* |
| PMF-7 | Brand permission for premium | Brand assessment across three dimensions: Image, Communication, Execution | Brand perceived as credible in premium/lifestyle space |

**Gate criteria:** PMF-1 (<5% addressable) or PMF-3 (WTP below price) failing triggers No-Go.

### Notes

- **PMF-3** uses two complementary methods. Van Westendorp PSM establishes the acceptable price range; conjoint analysis decomposes WTP across components. See METHODOLOGY.md Section 6 for method details.
- **PMF-6** conversion benchmarks vary significantly: self-serve freemium averages 3-5%, sales-assisted 5-7%, top performers reach 8-15% ([First Page Sage, 2026](https://firstpagesage.com/seo-blog/saas-freemium-conversion-rates/)).
- **PMF-7** Image = how the brand is perceived; Communication = how the brand presents itself; Execution = how the brand delivers on promises.

---

## Dimension 2: Pricing Adequacy (PRC)

**Core question:** Are the price points justified by perceived value and supported by willingness to pay?

| ID | Name | Assessment Method | Pass Threshold |
|----|------|-------------------|----------------|
| PRC-1 | Bundle Value Ratio (BVR) | Sum of standalone component prices / bundle price | >1.5x at all tiers; >2.0x preferred. *Practitioner guidance.* |
| PRC-2 | Price-value ratio (customer-perceived) | Survey-based: perceived value / actual price | >1.0 |
| PRC-3 | Tier gap architecture | Price step analysis between tiers | Proportional gaps between tiers; avoid >50% jumps without matching value delta |
| PRC-4 | Anchoring effectiveness | Good-Better-Best tier dynamics analysis | ~66% of customers target on middle tier |
| PRC-5 | Price vs. market average | Premium price / market average for comparable products | Premium index explainable by demonstrable bundle value |
| PRC-6 | Affordability for target segment | Product price as % of target segment's relevant discretionary budget | *Calibrate per industry* |
| PRC-7 | Price floor clearance | Cost-plus analysis: total cost per customer per period | Price > cost floor at all tiers |
| PRC-8 | Behavioral pricing coherence | Decoy effect validation: entry tier pushes customers toward middle | Feature/price gap between tiers drives intended migration |

**Gate criteria:** BVR < 1.0 at any tier triggers pricing redesign. Disproportionate tier gaps trigger tier restructuring.

### Notes

- **PRC-1** BVR thresholds (>1.5x / >2.0x) are practitioner heuristics widely used in product management. No single academic source validates these specific numbers; optimal ratios depend on category and competitive context.
- **PRC-4** draws on the compromise effect: research consistently shows ~66% of customers choose the middle option in three-tier pricing ([HBR, Mohammed 2018](https://hbr.org/2018/09/the-good-better-best-approach-to-pricing)). Industry sources suggest removing the middle tier can reduce revenue by approximately 50%.
- **PRC-6** affordability varies by product category. Different products compete for different budget "wallets." The relevant budget denominator must be defined per industry.

---

## Dimension 3: Bundle Composition (BND)

**Core question:** Is each component a Leader, Filler, or Killer? What is the dead weight ratio?

### Component Classification

Components are classified using the Leaders / Fillers / Killers framework ([Simon-Kucher](https://www.simon-kucher.com/en/insights/future-proof-telco-marketing-6-strategies-best-practices)):

| Role | Definition | Target Count per Bundle |
|------|-----------|------------------------|
| **Leader** | High perceived value; drives purchase intent. The reason customers buy. | 2-3 |
| **Filler** | Adds perceived value at low marginal cost. Nice-to-have. | 3-5 |
| **Killer** | Reduces WTP, confuses the offer, or attracts wrong customers. Must be eliminated. | 0 |

### Criteria

| ID | Name | Assessment Method | Red Flag |
|----|------|-------------------|----------|
| BND-1 | Leader identification | Conjoint part-worth analysis: which components drive purchase intent? | No clear leader = confused value proposition |
| BND-2 | Dead weight ratio | Usage tracking: % of components used by <20% of customers within 3 months | >40% of components are dead weight |
| BND-3 | Dilution risk | Test: does removing a component increase bundle WTP? | Any component whose removal increases WTP is a Killer |
| BND-4 | Cross-subsidy balance | Cost per component vs. revenue contribution | Subsidized components cost >50% of incremental margin |
| BND-5 | Access/availability constraint ratio | % of components requiring specific conditions for use (physical presence, credentials, platform, geography) | >30% of value tied to access-constrained components |
| BND-6 | Complementarity | Do bundled services naturally reinforce each other? | Low complementarity = "random collection" perception |
| BND-7 | Customizability potential | Can customers swap components? | Fixed bundles with >8 components = high dead weight certainty |
| BND-8 | Switching cost creation | How many active services/features increase exit friction? | <3 active services per customer = weak lock-in |

**Gate criteria:** No clear Leader (BND-1) triggers bundle redesign. Dead weight >40% (BND-2) triggers Killer removal or addition of swappability. Access constraints >30% (BND-5) triggers component adaptation for broader reach.

### Notes

- **BND-2** threshold derives from the inverse of the finding that ~60% of customers respond positively to additional benefits ([Simon-Kucher 2024 Global Telecommunications Study](https://www.simon-kucher.com/en/insights/2024-telco-growth-strategies-brand-portfolio-and-pricing-insights)).
- **BND-3** is grounded in the dilution effect (Shaddy & Fishbach, 2017): consumers perceive bundles as gestalt units. Adding low-value components can reduce total WTP because consumers resist altering the "whole." The practical test is straightforward -- if removing a component increases WTP, that component is a Killer. Source: [Shaddy & Fishbach, Journal of Marketing Research](https://www.anderson.ucla.edu/documents/areas/fac/marketing/Seminars/Fall%202017/Shaddy%20%20Fishbach%20-%20How%20Bundling%20Affects%20Valuation%20(job%20market%20paper).pdf).
- Dead weight is not purely negative. Unused components can contribute **option value** -- perceived value from availability alone. However, if the dilution effect triggers (BND-3), option value is not sufficient justification. Always run the BND-3 test.
- Offering swappable benefits converts a portion of initially disinterested customers. The exact conversion rate is industry-specific.

---

## Dimension 4: Competitive Positioning (CMP)

**Core question:** How does this product sit in the market, and how defensible is it?

| ID | Name | Assessment Method | Pass Threshold |
|----|------|-------------------|----------------|
| CMP-1 | Uniqueness of offering | Feature-by-feature comparison with direct competitors | >3 features not replicable by competitors within 6 months |
| CMP-2 | Competitive price-value ratio | BVR analysis applied to competitor premium offerings | Own BVR >= competitor BVR |
| CMP-3 | Time to competitive imitation | Assessment of competitor capability to replicate the offering | >12 months for full replication |
| CMP-4 | Competitive response risk | Game theory analysis: likely competitor reactions | Competitors can match some but not all dimensions = sustainable |
| CMP-5 | Defensible differentiation | Exclusive partnerships, proprietary services, unique technology | At least 2 exclusive or hard-to-replicate components |
| CMP-6 | Segment ownership | Which player "owns" the premium segment in customer perception? | Brand recognized as premium leader (perception survey) |
| CMP-7 | Cross-competitive set | Bundle competes not just with direct competitors but also standalone alternatives | Bundle price < sum of best individual alternatives |

**Gate criteria:** <2 defensible components (CMP-5) combined with <6 months to full imitation (CMP-3) triggers differentiation rethink.

### Notes

- **CMP-7** is critical for bundles competing against unbundled alternatives. Mixed bundling (offering both bundle and standalone components) almost always strictly increases profits vs. pure bundling or pure component selling. This is the strongest theoretical result in bundling theory ([McAfee, McMillan & Whinston, 1989](https://academic.oup.com/qje/article-abstract/104/2/371/1854649)). Practical implication: always offer the bundle AND standalone components simultaneously.
- Bundles with diverse, heterogeneous components support higher prices because variety itself signals value ([Xia Wei et al., 2025](https://journals.sagepub.com/doi/10.1177/00472875231222263)).

---

## Dimension 5: Financial Viability (FIN)

**Core question:** Do the unit economics work after cannibalization and cross-subsidy?

| ID | Name | Assessment Method | Pass Threshold |
|----|------|-------------------|----------------|
| FIN-1 | Revenue per customer uplift | Premium revenue per customer vs. current base revenue per customer | *Calibrate per industry:* meaningful uplift required |
| FIN-2 | Unit economics per customer | Revenue minus all costs (licensing, partner fees, delivery, support) per customer | Positive margin per customer at all tiers |
| FIN-3 | Cannibalization rate | % of premium adopters migrating from existing offerings vs. net new | Net incremental revenue > 0 after cannibalization |
| FIN-4 | Partner/licensing cost ratio | Partner + licensing costs / premium revenue delta | <30% of premium revenue delta. *Practitioner guidance.* |
| FIN-5 | Lifetime value premium | Customer lifetime value (CLV) of premium customer vs. base customer | Premium CLV >= 2x base CLV. *Practitioner target.* |
| FIN-6 | Time to break-even | Months until cumulative revenue > cumulative cost per customer | *Calibrate per industry* |
| FIN-7 | Cross-subsidy sustainability | Stress test: partner costs +20%, customer growth -30% | Model remains margin-positive under stress |
| FIN-8 | Scale economics | Minimum customer count for financial viability | Breakeven customer count achievable within 12 months |
| FIN-9 | Standalone product cannibalization | Lost standalone product revenue resulting from bundling | Net revenue positive including standalone losses |

**Gate criteria:** Negative unit economics at all tiers (FIN-2) triggers repricing or cost reduction. Negative net revenue after cannibalization (FIN-3) triggers migration redesign. Stress test failure (FIN-7) triggers cost buffer creation.

### Notes

- **FIN-3 + FIN-9** together capture total cannibalization impact. FIN-3 measures migration from existing plans; FIN-9 measures lost standalone product sales. Both must be net positive.
- **FIN-4** threshold (<30%) is a practitioner heuristic. No externally validated benchmark exists.
- **FIN-5** the 2-4x CLV multiple for premium vs. base is a widely used practitioner target in subscription businesses, but not sourced from a single definitive study.
- **FIN-7** stress test parameters (+20% costs, -30% growth) are deliberately harsh. The logic: if the model survives pessimistic conditions, it's likely sustainable under normal variance.
- Bundling's primary financial lever is often **churn reduction**, not direct revenue uplift. Even thin or negative per-customer margins can be justified if churn reduction delivers sufficient lifetime value improvement. Empirical research in one industry found unbundling leads to ~10% profit decrease and ~17% consumer surplus decrease ([Luo, 2023](https://onlinelibrary.wiley.com/doi/10.1111/1756-2171.12437)). Model churn impact explicitly.
- Acquisition is 5-25x more expensive than retention ([HBR, Gallo 2014](https://hbr.org/2014/10/the-value-of-keeping-the-right-customers)). Factor retention economics into financial modeling.

---

## Dimension 6: Customer Experience (CX)

**Core question:** Will this reduce churn, improve satisfaction, and deliver on the experience promise?

| ID | Name | Assessment Method | Pass Threshold |
|----|------|-------------------|----------------|
| CX-1 | Churn/attrition reduction | Premium tier churn rate vs. base churn rate | Premium churn < 50% of base churn |
| CX-2 | NPS improvement | NPS for premium customers vs. standard customers | Premium NPS significantly above base. *Calibrate per industry.* |
| CX-3 | CSAT score | Satisfaction survey across all bundle components | >80% satisfaction |
| CX-4 | Feature utilization rate | % of premium features actively used per customer | >60% of features used by >60% of customers. *Practitioner target.* |
| CX-5 | WTP validation | Actual conversion rate vs. projected WTP | Price < 80% of measured WTP |
| CX-6 | Support experience | Priority support resolution time vs. standard | Premium resolution time < 50% of standard |
| CX-7 | Component engagement | Usage frequency and depth per bundle component per period | Engagement rates track with standalone service benchmarks |
| CX-8 | Disappointment/dilution risk | Post-purchase perceived value vs. pre-purchase expectation | <15% value perception drop after 3 months |

**Gate criteria:** WTP validation failure (CX-5: price > 80% of WTP) triggers repricing. High disappointment risk (CX-8: >15% drop) triggers component quality improvement.

### Notes

- **CX-1** churn reduction varies significantly by bundle design and industry:
  - Modest effect (5-15%): general bundling in mature markets ([Prince & Greenstein, 2014](https://host.kelley.iu.edu/riharbau/RePEc/iuk/wpaper/bepp2011-05-prince-greenstein.pdf) -- the authors characterize the effect as "modest").
  - Moderate effect (25-35%): well-designed multi-product bundles (industry consensus; conservative estimate).
  - Strong effect (50%+): tightly integrated entertainment bundles ([Ampere Analysis: Disney+ bundle subscribers 59% less likely to churn](https://www.nexttv.com/news/disney-bundlers-59-less-likely-to-churn-research-company-says-chart)).
- **CX-2** NPS benchmarks vary enormously by industry. Do not use a single number as a universal target. Research your specific sector's NPS distribution.
- **CX-4** feature utilization target (>60%) is a widely used product management heuristic but has no single externally validated academic source.
- **CX-8** threshold (~15% value perception drop) is based on hospitality/travel bundle research on all-inclusive packages. Directionally applicable across industries, but the exact tolerance may differ.
- McKinsey research shows personalization drives 10-15% revenue lift and 10-30% marketing ROI improvement ([McKinsey](https://www.mckinsey.com/capabilities/growth-marketing-and-sales/our-insights/the-value-of-getting-personalization-right-or-wrong-is-multiplying)). Personalizing the premium experience amplifies both CX-1 and CX-2.

---

## Dimension 7: Market Reach (MR)

**Core question:** Does this work beyond the primary market? What are the reach and affordability constraints?

| ID | Name | Assessment Method | Pass Threshold |
|----|------|-------------------|----------------|
| MR-1 | Access coverage of constrained components | Physical locations, platform availability, credential requirements | Service accessible in target markets covering >70% of addressable customers |
| MR-2 | Segment affordability | Product price as % of target segment's relevant budget | *Calibrate per industry:* within acceptable affordability threshold for each segment |
| MR-3 | Segment revenue ceiling | Maximum sustainable revenue per customer by segment | Product price < segment revenue ceiling |
| MR-4 | Segment competitive landscape | Premium offering competition by market segment | No dominant competitor premium offering in target segments |
| MR-5 | Delivery readiness by segment | Capability to deliver premium experience across all segments | Premium experience deliverable in all launch segments |
| MR-6 | Demand heterogeneity across segments | Segment sizing across markets, geographies, and demographics | Viable segment in at least primary + secondary markets |
| MR-7 | Product adaptation by segment | Can bundle components be swapped or adapted per segment? | Alternative components available in non-primary segments |

**Gate criteria:** Access coverage <70% (MR-1) triggers component adaptation for broader reach. No viable secondary markets (MR-6) triggers niche launch strategy consideration.

### Notes

- "Market" and "segment" here are intentionally broad. Depending on the product, segments may be geographic (countries, regions, cities), demographic (age, income), platform-based (iOS vs. Android vs. web), or channel-based (direct vs. partner).
- **MR-1** covers all access constraints: physical (store locations, venues), digital (platform availability), credential-based (qualifications, memberships), and geographic (service area). Any component that requires specific conditions to use is an access-constrained component.
- **MR-7** swappability across segments reduces dead weight and increases bundle relevance. If primary market customers get Component A but secondary market customers need Component B serving the same role, the bundle adapts rather than loses relevance.

---

## Evaluation Logic Flow

Dimensions are evaluated sequentially. Each stage has gate criteria that can stop progression.

```
1. DEMAND VALIDATION (PMF-1 through PMF-7)
   Does anyone want this? Is the segment real?
   GATE: PMF-1 fails (<5% addressable) or PMF-3 fails (WTP < price) -> No-Go

2. PRICE-VALUE MAPPING (PRC-1 through PRC-8)
   Is the price justified by perceived value?
   GATE: BVR < 1.0 at any tier -> Redesign pricing
   GATE: Tier gaps disproportionate -> Restructure tiers

3. BUNDLE DECOMPOSITION (BND-1 through BND-8)
   Is each component earning its place?
   GATE: No clear Leader (BND-1) -> Redesign bundle
   GATE: Dead weight >40% (BND-2) -> Remove Killers or add swappability
   GATE: Access constraints >30% (BND-5) -> Adapt for secondary markets

4. COMPETITIVE ANALYSIS (CMP-1 through CMP-7)
   Can it survive in the market?
   GATE: <2 defensible components (CMP-5) AND <6 months to imitation (CMP-3)
         -> Rethink differentiation

5. FINANCIAL MODELING (FIN-1 through FIN-9)
   Does the math work?
   GATE: Unit economics negative at all tiers (FIN-2) -> Reprice or reduce costs
   GATE: Net revenue negative after cannibalization (FIN-3) -> Redesign migration
   GATE: Stress test fails (FIN-7) -> Build cost buffers

6. CUSTOMER IMPACT (CX-1 through CX-8)
   Will it improve the customer relationship?
   GATE: Price > 80% of WTP (CX-5) -> Reprice
   GATE: Disappointment risk >15% (CX-8) -> Improve component quality

7. MARKET REACH (MR-1 through MR-7)
   Where can this actually launch?
   GATE: Access coverage <70% (MR-1) -> Adapt components
   GATE: No viable secondary markets (MR-6) -> Niche launch strategy

8. GO / NO-GO / REDESIGN DECISION
   Weighted scoring across all 7 dimensions
```

---

## Decision Framework

### Go/No-Go Scoring Matrix

| Dimension | Weight | Score (1-5) | Weighted Score |
|-----------|--------|-------------|----------------|
| Strategic fit (PMF) | 15% | | |
| Financial viability (FIN) | 25% | | |
| Customer demand (PRC + CX) | 20% | | |
| Competitive position (CMP) | 15% | | |
| Bundle composition (BND) | 10% | | |
| Market reach (MR) | 10% | | |
| Risk profile | 5% | | |
| **Total** | **100%** | | |

### Decision Thresholds

| Total Weighted Score | Decision | Action |
|---------------------|----------|--------|
| >= 4.0 | **Strong Go** | Proceed to launch planning |
| 3.0 - 3.9 | **Conditional Go** | Address specific gaps identified by failing criteria |
| 2.0 - 2.9 | **Redesign Required** | Fundamental changes needed to pricing, composition, or positioning |
| < 2.0 | **No-Go** | Do not launch in current form |

---

## Calibration Guide

The following thresholds must be calibrated when applying this framework to a specific industry. Listed defaults come from subscription/recurring product practice.

| Criterion | Default Threshold | Why Calibration Is Needed | Examples |
|-----------|------------------|---------------------------|----------|
| PMF-6: Conversion funnel | 3-10% entry, 1-5% top | Conversion varies dramatically by sales model and price point | SaaS self-serve: 3-5%; sales-assisted: 5-7%; top performers: 8-15% |
| PRC-6: Affordability | <3-5% of relevant budget | Different products compete for different "wallets" | Utilities: % of household income; SaaS: % of IT budget; luxury: % of discretionary spend |
| FIN-1: Revenue uplift | Meaningful uplift required | Magnitude depends on base price and margin structure | SaaS premium: 50-100%+; physical goods: 10-20%; media: 15-30% |
| FIN-6: Time to break-even | 6-18 months | Depends on investment level, margin structure, customer count | High-investment launches may need 24+ months |
| CX-2: NPS target | Significantly above base | NPS benchmarks vary enormously by sector | Tech: 40-60; retail: 50-70; financial services: 20-40 |
| MR-2: Segment affordability | Within acceptable threshold | Same product may be affordable in one segment, unaffordable in another | Calibrate per segment with local income/budget data |
| CAC payback | Varies | Depends on LTV:CAC norms in the industry | SaaS: 12-18 months; consumer apps: 1-3 months |
| Penetration of eligible base | 10-25% at maturity | Depends on market maturity, pricing, and segment definition | Early markets may see 2-5%; mature may exceed 25% |

---

## Source Index

All external benchmarks and frameworks referenced in this document, with verification status.

### Verified Academic Sources

| Source | Finding | URL |
|--------|---------|-----|
| Adams & Yellen (1976), *Quarterly Journal of Economics* | Bundling as price discrimination via reservation prices | [QJE](https://academic.oup.com/qje/article-abstract/90/3/475/1854397) |
| Schmalensee (1984), *Journal of Business* | Bundling reduces variance of consumer valuations | [JSTOR](https://www.jstor.org/stable/2352937) |
| McAfee, McMillan & Whinston (1989), *QJE* | Mixed bundling almost always strictly increases profits | [QJE](https://academic.oup.com/qje/article-abstract/104/2/371/1854649) |
| Wansink, Kent & Hoch (1998), *Journal of Marketing Research* | Bundle framing ("3 for $5") boosts sales ~32% | Original paper: JMR 35(1), 71-81 |
| Tversky & Kahneman (1974) | Price anchoring -- first price seen influences fairness perception | "Judgment Under Uncertainty: Heuristics and Biases", *Science* |
| Prince & Greenstein (2014), *J. of Economics & Management Strategy* | Bundled customers stay modestly longer (6.93 vs. 6.15 years for broadband) | [Working Paper](https://host.kelley.iu.edu/riharbau/RePEc/iuk/wpaper/bepp2011-05-prince-greenstein.pdf) |
| Shaddy & Fishbach (2017), *Journal of Marketing Research* | Dilution effect -- low-value components can reduce total bundle WTP | [UCLA Anderson](https://www.anderson.ucla.edu/documents/areas/fac/marketing/Seminars/Fall%202017/Shaddy%20%20Fishbach%20-%20How%20Bundling%20Affects%20Valuation%20(job%20market%20paper).pdf) |
| Luo (2023), *RAND Journal of Economics* | Unbundling leads to ~10% profit decrease, ~17% consumer surplus decrease | [Wiley](https://onlinelibrary.wiley.com/doi/10.1111/1756-2171.12437) |
| Blut et al. (2024), *Journal of Service Research* | Perceived value meta-analysis (687 articles, 357,247 customers) | [SAGE](https://journals.sagepub.com/doi/10.1177/10946705231222295) |
| Xia Wei, Yu & Li (2025), *Journal of Travel Research* | Heterogeneous bundles support higher prices | [SAGE](https://journals.sagepub.com/doi/10.1177/00472875231222263) |

### Verified Industry/Consulting Sources

| Source | Finding | URL |
|--------|---------|-----|
| Mohammed (2018), *Harvard Business Review* | Good-Better-Best pricing; ~66% choose middle tier | [HBR](https://hbr.org/2018/09/the-good-better-best-approach-to-pricing) |
| Gallo (2014), *Harvard Business Review* | Customer acquisition costs 5-25x more than retention | [HBR](https://hbr.org/2014/10/the-value-of-keeping-the-right-customers) |
| Ampere Analysis (Disney+ bundle data) | Bundle subscribers 59% less likely to churn | [NextTV](https://www.nexttv.com/news/disney-bundlers-59-less-likely-to-churn-research-company-says-chart) |
| Simon-Kucher (2024 Global Study) | Leaders/Fillers/Killers framework; ~60% respond positively to additional benefits | [Simon-Kucher](https://www.simon-kucher.com/en/insights/2024-telco-growth-strategies-brand-portfolio-and-pricing-insights) |
| McKinsey (Personalization research) | Personalization drives 10-15% revenue lift; 10-30% marketing ROI improvement | [McKinsey](https://www.mckinsey.com/capabilities/growth-marketing-and-sales/our-insights/the-value-of-getting-personalization-right-or-wrong-is-multiplying) |
| Deloitte TMT Predictions 2025 | ~4 streaming services per household peak; aggregation trend growing | [Deloitte](https://www.deloitte.com/us/en/insights/industry/technology/technology-media-and-telecom-predictions/2025/tmt-predictions-video-streaming-bundles-bigger-than-ever.html) |
| First Page Sage (2026) | Freemium conversion: 3-5% self-serve; 5-7% sales-assisted; 8-15% top performers | [First Page Sage](https://firstpagesage.com/seo-blog/saas-freemium-conversion-rates/) |
| Bain & Company (NPS case study) | NPS touchpoint scores improved by >30 points in implementation | [Bain](https://www.bain.com/client-results/dialing-up-customer-experience-in-telecommunications/) |

### Verified Research Methods

| Method | Source | URL |
|--------|--------|-----|
| Van Westendorp Price Sensitivity Meter | Peter Van Westendorp (1976) | [Wikipedia](https://en.wikipedia.org/wiki/Van_Westendorp's_Price_Sensitivity_Meter), [Sawtooth](https://sawtoothsoftware.com/resources/blog/posts/van-westendorp-pricing-sensitivity-meter) |
| Gabor-Granger Method | Andre Gabor & Clive Granger (1960s) | [Sawtooth](https://sawtoothsoftware.com/resources/blog/posts/gabor-granger-pricing-method) |
| Choice-Based Conjoint Analysis | Standard methodology | [BCG variant](https://www.bcg.com/publications/2014/telecommunications-pricing-pathways-conjoint-new-approach-pricing-mobile) |

### Practitioner Guidance (Not Externally Validated)

These thresholds are reasonable heuristics widely used in practice but not traceable to specific academic or consulting publications:

| Threshold | Value | Notes |
|-----------|-------|-------|
| BVR adequate / strong | >1.5x / >2.0x | Optimal ratio depends on category and competitive context |
| Premium CLV multiple | 2-4x base | Common subscription business target |
| Partner cost ratio | <30% of revenue delta | Industry-specific; depends on partner economics |
| Feature utilization | >60% used by >60% | Common product management heuristic |
| Premium churn target | <50% of base | Practitioner target; actual reduction varies 5-59% by bundle design |
| Multi-product churn reduction | 25-35% lower | Conservative estimate; documented range is 5-59% |
| Dead weight ceiling | <40% of components | Derived from ~60% positive response rate |
| Over-provisioning waste | 20-30% of feature costs | Based on expense management and cloud infrastructure analysis |
| Middle tier revenue loss | ~50% when removed | Industry sources; directionally consistent with GBB research |

---

*48 criteria across 7 dimensions. All IDs (PMF-1 through MR-7) are stable and should be referenced in appraisal reports for traceability. When applying to a specific industry, calibrate thresholds per the [Calibration Guide](#calibration-guide), select relevant benchmarks, and adjust definitions to match industry terminology.*
