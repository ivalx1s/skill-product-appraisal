# Generalized Assessment Dimensions

**Date:** 2026-02-12
**Source:** 7 Assessment Dimensions from MTS methodology framework
**Purpose:** Replace telecom-specific criteria with universal equivalents for any complex product/bundle evaluation

---

## Evaluation Logic Flow (Generalized)

```
Market Demand Validation (Does anyone want this?)
  -> Price-Value Mapping (Is the price justified?)
    -> Bundle Decomposition (Is each piece earning its place?)
      -> Competitive Analysis (Can it survive in the market?)
        -> Financial Modeling (Does the math work?)
          -> Customer Impact Projection (Will it improve or damage customer relationship?)
            -> Market Reach Assessment (Where/who can this actually serve?)
              -> Go / No-Go / Redesign Decision
```

---

## Dimension 1: Product-Market Fit

**Core Question:** Does the target segment exist, and does the bundle address their actual needs?

| # | Criterion | Assessment Method | Pass Threshold |
|---|-----------|-------------------|----------------|
| PMF-1 | Addressable target segment size | Customer data segmented by income, behavior, engagement | >5% of total customer base |
| PMF-2 | Demonstrated demand for premium tier | Market research, existing premium/loyalty program adoption trends | Premium engagement penetration >15% and growing |
| PMF-3 | Segment willingness to pay at proposed prices | Conjoint analysis / Van Westendorp on target segment | WTP >= proposed price for >30% of target segment |
| PMF-4 | Service-need alignment | Survey: which bundle components does the target segment actually want? | >60% of target segment wants >60% of bundle components |
| PMF-5 | Current unmet need | Analysis of target segment currently using separate services / fragmented solutions | Evidence of fragmented premium service consumption |
| PMF-6 | Conversion funnel viability | Funnel: total customers -> engaged users -> premium-eligible -> converters | Realistic conversion: 3-10% of premium-eligible at entry tier, 1-5% at top tier |
| PMF-7 | Brand permission for premium | ICE Model assessment (Image, Communication, Execution) | Brand perceived as credible in premium/lifestyle space |

**Changes from telecom version:**
- PMF-1: "subscriber base" -> "customer base" (universal)
- PMF-2: Removed MTS-specific ecosystem user numbers
- PMF-6: Removed MTS-specific funnel numbers (82.2M, 17.5M); kept structure with generic labels

---

## Dimension 2: Pricing Adequacy

**Core Question:** Are the price points justified by perceived value and supported by WTP?

| # | Criterion | Assessment Method | Pass Threshold |
|---|-----------|-------------------|----------------|
| PRC-1 | Bundle Value Ratio | Sum of standalone prices / bundle price | >1.5x at all tiers (>2.0x preferred) |
| PRC-2 | Price-value ratio (customer-perceived) | Survey-based: perceived value / actual price | >1.0 |
| PRC-3 | Tier gap architecture | Price step analysis between tiers | Proportional gaps between tiers (avoid >50% jumps without matching value delta) |
| PRC-4 | Anchoring effectiveness | Three-tier Good-Better-Best dynamics | ~66% target on middle tier |
| PRC-5 | Price vs. market average | Premium plan price / market average for comparable products | Premium index explainable by bundle value |
| PRC-6 | Affordability for target segment | Plan price as % of target segment discretionary budget | *Calibrate per industry:* <3-5% of relevant budget category |
| PRC-7 | Price floor clearance | Cost-plus analysis: total cost per customer per period | Price > cost floor at all tiers |
| PRC-8 | Behavioral pricing coherence | Decoy effect works as intended (entry tier pushes to middle) | Feature/price gap between tiers drives intended migration |

**Changes from telecom version:**
- PRC-5: "market average (~543 RUB)" -> "market average for comparable products"
- PRC-6: "monthly disposable income" -> "relevant budget category" (different products take share from different wallets)
- PRC-6 threshold: Marked "calibrate per industry" -- telecom uses 3-5% of income; SaaS, consumer goods, etc. differ

---

## Dimension 3: Bundle Composition

**Core Question:** Is each component a Leader, Filler, or Killer? What is the dead weight ratio?

**Component Classification (Simon-Kucher Leaders / Fillers / Killers):**

| Role | Definition | Ideal Ratio |
|------|-----------|-------------|
| **Leader** | High perceived value, drives purchase intent. The reason people buy. | 2-3 per bundle |
| **Filler** | Adds perceived value at low marginal cost. Nice-to-have. | 3-5 per bundle |
| **Killer** | Reduces WTP, confuses the offer, or attracts wrong customers. Must be eliminated. | 0 per bundle |

| # | Criterion | Assessment Method | Red Flag |
|---|-----------|-------------------|----------|
| BND-1 | Leader identification | Conjoint part-worth analysis: which components drive purchase intent? | No clear leader = confused value proposition |
| BND-2 | Dead weight ratio | Usage tracking: % of components used by <20% of customers within 3 months | >40% of components are dead weight |
| BND-3 | Dilution risk | Test: does removing a component increase bundle WTP? | Any component whose removal increases WTP is a Killer |
| BND-4 | Cross-subsidy balance | Cost per component vs. revenue contribution | Subsidized components cost >50% of incremental margin |
| BND-5 | Access/availability constraint ratio | % of components requiring specific conditions for use (physical presence, credentials, platform, geography) | >30% of value tied to access-constrained components |
| BND-6 | Complementarity | Do bundled services naturally reinforce each other? | Low complementarity = "random collection" perception |
| BND-7 | Customizability potential | Can customers swap components? | Fixed bundles with >8 components = high dead weight certainty |
| BND-8 | Switching cost creation | How many active services/features increase exit friction? | <3 active services per customer = weak lock-in |

**Changes from telecom version:**
- BND-5: "Geographic constraint ratio" + "physical presence" -> "Access/availability constraint ratio" covering physical, digital, credential, platform, and geographic constraints. Threshold kept at >30%.

---

## Dimension 4: Competitive Positioning

**Core Question:** How does this sit in the market, and how defensible is it?

| # | Criterion | Assessment Method | Pass Threshold |
|---|-----------|-------------------|----------------|
| CMP-1 | Uniqueness of offering | Feature-by-feature comparison with direct competitors | >3 features not replicable within 6 months |
| CMP-2 | Competitive price-value ratio | Same BVR analysis applied to competitor premium offerings | Own BVR >= competitor BVR |
| CMP-3 | Time to competitive imitation | Assessment of competitor capability to replicate | >12 months for full replication |
| CMP-4 | Competitive response risk | Game theory: likely competitor reactions | Competitor matching on some but not all dimensions = sustainable |
| CMP-5 | Defensible differentiation | Exclusive partnerships, proprietary services, unique technology | At least 2 exclusive or hard-to-replicate components |
| CMP-6 | Market segment ownership | Which player "owns" the premium segment in customer perception? | Recognized as premium leader (brand/perception survey) |
| CMP-7 | Cross-competitive set | Bundle competes not just with direct competitors but with standalone alternatives | Bundle price < sum of best individual alternatives |

**Changes from telecom version:**
- CMP-1: "MegaFon, Beeline, Tele2" -> "direct competitors"
- CMP-4: "Competitors match on content but not lifestyle" -> generic formulation
- All operator-specific references removed

---

## Dimension 5: Financial Viability

**Core Question:** Do the unit economics work after cannibalization and cross-subsidy?

| # | Criterion | Assessment Method | Pass Threshold |
|---|-----------|-------------------|----------------|
| FIN-1 | Revenue per customer uplift | Premium revenue per customer vs. current base | *Calibrate per industry:* meaningful uplift (telecom benchmark: +12-18% blended) |
| FIN-2 | Unit economics per customer | Revenue per customer minus all costs (licensing, partner fees, delivery, support) | Positive margin per customer at all tiers |
| FIN-3 | Cannibalization rate | % of premium adopters migrating from existing offerings vs. net new | Net incremental revenue > 0 after cannibalization |
| FIN-4 | Partner/licensing cost ratio | Partner + licensing costs / premium revenue delta | <30% of premium revenue delta |
| FIN-5 | Lifetime value premium | CLV of premium customer vs. base customer | Premium CLV >= 2x base CLV |
| FIN-6 | Time to break-even | Months until cumulative revenue > cumulative cost per customer | 6-18 months |
| FIN-7 | Cross-subsidy sustainability | Stress test: partner costs +20%, customer growth -30% | Model remains margin-positive under stress |
| FIN-8 | Scale economics | Minimum customer count for financial viability | Breakeven customer count achievable within 12 months |
| FIN-9 | Cannibalization of standalone products | Lost standalone product revenue from bundling | Net revenue positive including standalone losses |

**Changes from telecom version:**
- FIN-1: "ARPU" -> "Revenue per customer"; "blended ARPU impact" -> "blended revenue impact"
- FIN-2: "AMPU (Average Margin Per User)" -> "Positive margin per customer"
- FIN-4: "content/VAS licensing" -> "partner/licensing"
- FIN-5: "subscriber" -> "customer"
- FIN-9: "KION, MTS Bank, World Class" -> "standalone product revenue"
- FIN-1 threshold: Marked "calibrate per industry" since +12-18% is telecom-specific

---

## Dimension 6: Customer Experience Metrics

**Core Question:** Will this reduce churn, improve satisfaction, and deliver on the experience promise?

| # | Criterion | Assessment Method | Pass Threshold |
|---|-----------|-------------------|----------------|
| CX-1 | Churn/attrition reduction | Premium tier churn vs. base churn | Premium churn < 50% of base churn |
| CX-2 | NPS improvement | NPS for premium customers vs. standard | Premium NPS > 30 (general industry avg varies by sector) |
| CX-3 | CSAT score | Satisfaction survey across bundle components | >80% satisfaction |
| CX-4 | Feature utilization rate | % of premium features actively used per customer | >60% of features used by >60% of customers |
| CX-5 | WTP validation | Actual conversion vs. projected WTP | Price < 80% of measured WTP |
| CX-6 | Support experience | Priority support resolution time delta | Premium resolution time < 50% of standard |
| CX-7 | Component engagement | Usage frequency/depth per bundle component per period | Engagement rates track with standalone service benchmarks |
| CX-8 | Dilution/disappointment risk | Post-purchase perceived value vs. pre-purchase expectation | <15% value perception drop after 3 months |

**Changes from telecom version:**
- CX-1: "Subscribers" -> "Customers"; "churn" generalized (works for SaaS, subscription, etc.)
- CX-2: NPS threshold kept at >30 but noted "varies by sector" -- NPS benchmarks differ greatly by industry
- CX-7: "KION hours, fitness visits, restaurant visits" -> "Usage frequency/depth per bundle component"

---

## Dimension 7: Market Reach / Regional Viability

**Core Question:** Does this work beyond the primary market? What are the reach and affordability constraints?

| # | Criterion | Assessment Method | Pass Threshold |
|---|-----------|-------------------|----------------|
| MR-1 | Access coverage of constrained components | Physical locations, platform availability, credential requirements | Service accessible in target markets covering >70% of addressable customers |
| MR-2 | Segment affordability | Product price as % of target segment's relevant budget | *Calibrate per industry:* within acceptable affordability threshold for each segment |
| MR-3 | Segment revenue ceiling | Maximum sustainable revenue per customer by segment | Product price < segment revenue ceiling |
| MR-4 | Segment competitive landscape | Premium offering competition by market segment | No dominant competitor premium offering in target segments |
| MR-5 | Delivery readiness by segment | Capability to deliver premium experience across segments | Premium experience deliverable in all launch segments |
| MR-6 | Demand heterogeneity across segments | Segment sizing across markets/geographies/demographics | Viable segment in at least primary + secondary markets |
| MR-7 | Product adaptation by segment | Can bundle components be swapped or adapted per segment? | Alternative components available in non-primary segments |

**Changes from telecom version (renamed from "Regional Viability"):**
- Dimension renamed: "Regional Viability" -> "Market Reach" -- not all products have geographic constraints; some have demographic, platform, or channel constraints
- REG-1: "World Class gym locations" -> "Access coverage of constrained components"
- REG-2: "Regional average income" -> "Target segment's relevant budget"
- REG-3: "Regional ARPU ceiling" -> "Segment revenue ceiling"
- REG-4: "Regional competitive landscape" -> "Segment competitive landscape"
- REG-5: "Network QoS capability" -> "Delivery readiness"
- REG-6: "Moscow vs SPb vs millionniki" -> "Primary + secondary markets"
- REG-7: "Regional bundle adaptation" -> "Product adaptation by segment"
- All Russia-specific geography removed

---

## Evaluation Templates (Generalized)

### Strategic Fit Scorecard

| Criterion | Weight | Score (1-5) | Weighted | Notes |
|-----------|--------|-------------|----------|-------|
| Market demand for premium tier | 15% | | | |
| Brand permission (ICE Model) | 10% | | | |
| Competitive gap (unserved premium segment) | 15% | | | |
| Strategic alignment with company direction | 10% | | | |
| Delivery readiness for premium experience | 10% | | | |
| Partner ecosystem completeness | 10% | | | |
| Market reach (beyond primary segment) | 15% | | | |
| Regulatory/compliance risk profile | 15% | | | |
| **Total Strategic Fit** | **100%** | | | |

### Go/No-Go Decision Matrix

| Dimension | Weight | Score (1-5) | Weighted Score |
|-----------|--------|-------------|----------------|
| Strategic fit | 15% | | |
| Financial viability | 25% | | |
| Customer demand validation | 20% | | |
| Competitive position | 15% | | |
| Bundle composition quality | 10% | | |
| Market reach | 10% | | |
| Risk profile | 5% | | |
| **Total** | **100%** | | |

**Decision thresholds:**

| Score | Decision | Action |
|-------|----------|--------|
| >= 4.0 | **Strong Go** | Proceed to launch planning |
| 3.0 - 3.9 | **Conditional Go** | Address specific gaps before launch |
| 2.0 - 2.9 | **Redesign Required** | Fundamental changes to pricing, composition, or positioning needed |
| < 2.0 | **No-Go** | Do not launch in current form |

---

## Traceability: Criteria ID Mapping

| Original ID | Generalized ID | Change Summary |
|-------------|----------------|----------------|
| PMF-1 through PMF-7 | PMF-1 through PMF-7 | Terminology generalized, numbers removed |
| PRC-1 through PRC-8 | PRC-1 through PRC-8 | Market-specific references removed; PRC-6 marked "calibrate" |
| BND-1 through BND-8 | BND-1 through BND-8 | BND-5 broadened from geographic to all access constraints |
| CMP-1 through CMP-7 | CMP-1 through CMP-7 | Operator names removed |
| FIN-1 through FIN-9 | FIN-1 through FIN-9 | ARPU/AMPU -> Revenue/Margin per customer; FIN-1 marked "calibrate" |
| CX-1 through CX-8 | CX-1 through CX-8 | Subscriber -> Customer; telecom features -> generic components |
| REG-1 through REG-7 | MR-1 through MR-7 | Renamed dimension; geography -> segment; all Russia-specific removed |
