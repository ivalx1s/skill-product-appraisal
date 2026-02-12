# Audit Inventory: Source Methodology Classification

**Date:** 2026-02-12
**Source documents:**
- `260212_methodology-framework.md` (MF)
- `260212_bundle-valuation-methods.md` (BVM)
- `260212_telecom-assessment-frameworks.md` (TAF)
- `260212_telecom-pricing-methodologies.md` (TPM)

Classification: **(a) Universal**, **(b) Adaptable**, **(c) Telecom-only**

---

## 1. Theoretical Frameworks

### (a) Universal -- Transfer 1:1

| Framework | Source Doc | Description | Why Universal |
|-----------|-----------|-------------|---------------|
| Adams-Yellen Bundling Theory (1976) | BVM, TPM | Pure/mixed/component bundling as price discrimination; reservation prices | Applies to any multi-component product |
| Schmalensee Variance Reduction (1984) | BVM, TPM | Bundling reduces variance of consumer valuations via law of large numbers | Mathematical property, product-agnostic |
| McAfee-McMillan-Whinston Mixed Bundling | BVM, TPM | Mixed bundling almost always strictly increases profits vs. pure bundling | Proven across industries |
| Leaders / Fillers / Killers (Simon-Kucher) | MF, BVM, TAF | Classify bundle components by value contribution; ideal ratios 2-3/3-5/0 | Bundle composition framework, any industry |
| Dead Weight Analysis | MF, BVM | Components used by <20% of customers; ~40% neutral-to-low interest benchmark | Any bundle with multiple components |
| Dilution Effect | BVM | Bundling expensive + inexpensive items can lower perceived value below expensive item alone | Cognitive bias, product-agnostic |
| Bundle Value Ratio (BVR) | MF, TPM | Sum of standalone prices / bundle price; thresholds >1.5x adequate, >2.0x strong | Arithmetic ratio, any bundle |
| Good-Better-Best (GBB) Tiered Pricing | MF, TPM | Three-tier architecture; ~20/66/14% distribution; middle tier as target | Behavioral economics, any tiered product |
| Decoy Effect | TPM | Lower tier designed to make middle tier look better | Behavioral pricing, any product |
| Price Anchoring (Tversky & Kahneman) | BVM, TPM | First price seen influences fairness perception; ~32% perceived value increase from anchoring | Cognitive psychology, universal |
| Van Westendorp PSM | BVM, TPM | Four-question price sensitivity meter; PMC/PME/OPP/IDP outputs | Pricing research method, any product |
| Gabor-Granger Method | BVM, TPM | Sequential price presentation to find max WTP; produces demand curves | Pricing research method, any product |
| Conjoint Analysis (CBC) | BVM, TAF, TPM | Choice-based analysis of attribute trade-offs; part-worth utilities; WTP derivation | Research method, any multi-attribute product |
| MaxDiff Scaling | BVM | Rank importance of bundle components (must-have / nice-to-have / dead weight) | Research method, any feature set |
| Value-Based Pricing | TAF, TPM | Price based on perceived customer value, not cost or competition | Pricing philosophy, any product |
| Cost-Plus Pricing (Floor Pricing) | TPM | Calculate total cost + target margin = floor price | Unit economics, any product |
| Cross-Subsidy Models | BVM | High-margin components subsidize low-margin ones within a bundle | Any multi-component product with mixed margins |
| Cannibalization Analysis | BVM, TPM | Positive (ARPU uplift, churn reduction) vs. negative (migration from standalone) effects | Any product that replaces existing offerings |
| Perceived Value vs. Actual Value | BVM | Three-dimension model (benefits, sacrifices, overall); 15% satisfaction penalty for undelivered promises | Consumer psychology, any product |
| "Highest-priced item" heuristic | BVM | Consumers evaluate bundles based on most expensive/desired component | Cognitive bias, any bundle |
| Framing Effects (Xia Wei et al.) | BVM | Heterogeneous bundles support higher prices; variety signals value | Behavioral economics, any bundle |
| Reference Price Distortion | BVM | Bundling creates price opacity; consumers lose ability to assess component values | Any complex product/bundle |
| Wansink Bundle Framing | BVM | "3 for $5" vs. "$1.67 each" -- 32% sales boost from bundle framing | Behavioral economics, any product |
| Option Value of Unused Components | BVM | Unused services still contribute perceived value through availability | Any feature-rich product |
| Switching Cost Creation | MF, BVM | Multi-service active use increases exit friction; <3 active services = weak lock-in | Any ecosystem/platform product |
| ICE Model (Image/Communication/Execution) | MF, TAF | Brand assessment for premium permission: image + communication + execution | Brand framework, any industry |
| McKinsey Value Delivery System | TAF | Choose value -> Provide value -> Communicate value | Strategic framework, any product |
| NPS / CSAT / CES | MF, TAF | Standard customer experience metrics | Universal customer metrics |
| CLV / LTV Calculation | MF, TAF | ARPU x Margin x Lifespan | Any subscription/recurring product |
| Go/No-Go Decision Matrix | MF, TAF | Weighted scoring across dimensions; thresholds: >=4.0 Strong Go, 3.0-3.9 Conditional, etc. | Decision framework, any product launch |
| Risk Assessment Matrix | MF | Probability x Impact scoring with mitigation | Universal risk framework |
| SWOT Analysis | MF | Strengths / Weaknesses / Opportunities / Threats | Universal strategic framework |
| Stage-Gate Product Launch | TAF | Stage-gate framework for portfolio decisions | Universal product management |

### (b) Adaptable -- Structure Universal, KPIs/Thresholds Need Replacement

| Framework/Element | Source Doc | What's Universal | What Needs Replacement |
|-------------------|-----------|------------------|----------------------|
| 7 Assessment Dimensions | MF | Structure of 7 evaluation dimensions with criteria IDs and pass thresholds | Telecom-specific criteria content: PMF-1 subscriber base, PMF-6 MTS funnel numbers, BND-5 geographic gym constraints, REG-* all telecom/Russia-specific |
| KPI Reference List (6 categories) | MF, TAF | Category structure: Revenue, Customer, Product Performance, Segment, Bundle Economics, Market Context | Specific KPIs: ARPU, Revenue per GB, VAS revenue share, roaming metrics, SIM registrations |
| Evaluation Logic Flow | MF | Sequential demand->pricing->bundle->competitive->financial->CX->regional->decision flow | "Regional Viability" assumes physical infrastructure; needs generalization to "Market Reach" |
| Strategic Fit Scorecard | MF, TAF | Weighted scoring template with criteria | Criteria names: "Infrastructure readiness for premium QoS", "Regulatory risk profile" |
| Financial Viability Scorecard | MF, TAF | Template structure with metrics vs. benchmarks | Metrics: ARPU, content/VAS licensing ratio, AMPU (telecom-specific margin) |
| Bundle Composition Scorecard | MF | Component-level L/F/K analysis template | Components listed are MTS-specific (KION, World Class, White Rabbit) |
| Competitive Position Scorecard | MF | Side-by-side comparison template | Operators named (MTS, MegaFon, Beeline, Tele2) |
| Customer Impact Scorecard | MF | Projected vs. benchmark gap analysis | Some benchmarks telecom-specific (NPS 15-25 industry avg) |
| Regional Viability Matrix | MF | Geographic affordability and service availability matrix | Russia-specific regions, income data, gym/restaurant locations |
| Conversion Funnel Model | MF, TPM | Funnel narrowing from total base -> eligible -> converters at each tier | MTS-specific numbers (82.2M, 17.5M, etc.) |
| Premium Differentiation Levers (5) | TAF | Five-lever framework for premium differentiation | Levers are telecom-flavored (QoS, content, lifestyle, support, device) |
| Premiumization Assessment Criteria (7) | TAF | Seven-point checklist for premiumization readiness | Criteria mention "network deliver differentiated QoS", "content partner ecosystem" |
| New Product Launch Framework (6 dimensions) | TAF | Six assessment dimensions: market sizing, competitive, revenue, cannibalization, investment, risk | Revenue model uses ARPU formula; investment mentions "spectrum licenses" |
| Competitive Benchmarking (4-step) | TAF | Feature-by-feature -> value comparison -> dynamic tracking -> segment alignment | Steps reference "data allowances, call minutes, SMS" |
| Cost Model (8-step) | TPM | Define scope -> map delivery chain -> identify cost drivers -> allocate -> model -> validate -> iterate | "Shared infrastructure costs", "spectrum licenses" in examples |
| Price Elasticity Ranges | TPM | Elasticity measurement methodology and segment patterns | Specific elasticity values (-0.43, -1.10) are telecom-derived |
| Churn Reduction Benchmarks | BVM, TPM | Churn reduction as primary bundle financial lever | Specific numbers: 59% Disney+ reduction, 6.93 vs 6.15 years broadband |

### (c) Telecom-Only -- Does Not Transfer

| Element | Source Doc | Why Telecom-Only |
|---------|-----------|-----------------|
| ARPU (Average Revenue Per User) | MF, TAF, TPM | Telecom-specific revenue metric. Universal equivalent: Revenue Per Customer (RPC) or Average Revenue Per Account (ARPA) |
| Revenue per GB | MF, TAF | Data consumption metric unique to telecom/ISP |
| AMPU (Average Margin Per User) | MF | Telecom-specific margin metric |
| Roaming analysis (168 days) | MF, BVM | Telecom-specific feature |
| FAS regulatory risk / 3B RUB fine | MF, TPM | Russia-specific telecom regulation |
| SIM registration requirements | TPM | Telecom-specific regulation |
| Network QoS / 5G / network slicing | TAF, TPM | Telecom infrastructure |
| Spectrum licenses | TPM | Telecom-specific cost |
| Double/Triple/Quad/Ecosystem Play evolution | BVM | Telecom-specific historical progression |
| GSMA MCI Methodology | TAF | Telecom industry body methodology |
| TM Forum eTOM Framework | TAF | Telecom-specific process framework |
| Zero-rating (Facebook, WhatsApp) | TAF | Telecom-specific net neutrality practice |
| Prepaid vs. Postpaid dynamics | TAF, TPM | Telecom billing model distinction |
| Device-as-a-Service (DaaS) | TAF | Telecom-specific upselling mechanism |
| MTS-specific ecosystem data (82.2M subs, 17.5M ecosystem, KION, etc.) | MF, TPM | Company-specific data points |
| Russian market price points (543 RUB avg, operator plan tables) | TPM | Country and company-specific data |
| Speed vs. Volume differentiation | TAF | Telecom-specific (GB vs Mbps) |
| ITU affordability target (<2% of income) | TAF | Telecom/UN-specific benchmark |

---

## 2. Benchmarks and Numbers Inventory

### (a) Universal Benchmarks

| Benchmark | Value | Source Claimed | Used In |
|-----------|-------|---------------|---------|
| McKinsey: bundling increases revenues | 10-30% | McKinsey research | BVM |
| HBR: mixed bundling outperforms pure bundling | 25-35% revenue | Harvard Business Review | BVM |
| Bundles perceived as more valuable than sum | 20-25% | Not attributed | BVM |
| Wansink: bundle framing sales boost | 32% | Wansink experiment | BVM |
| Simon-Kucher: neutral/low interest in VAS | ~40% | Simon-Kucher 2024 | MF, BVM |
| Simon-Kucher: swappable benefits conversion | 13% | Simon-Kucher 2024 | MF, BVM |
| Anchoring increases perceived value | ~32% | Tversky & Kahneman referenced | BVM |
| BVR threshold: adequate | >1.5x | Not attributed | MF, TPM |
| BVR threshold: strong | >2.0x | Not attributed | MF |
| GBB distribution: middle tier | ~66% | Research referenced | MF, TPM |
| GBB distribution: top tier | ~14% | Research referenced | MF, TPM |
| GBB distribution: bottom tier | ~20% | Research referenced | MF, TPM |
| Removing middle tier revenue drop | 48-60% | Not attributed | TPM |
| Multi-product customers lower churn | 25-35% | Not attributed | BVM, TPM |
| Disney+ bundle churn reduction | 59% | Disney+ data | BVM |
| Broadband bundle retention | 6.93 vs 6.15 years | Prince & Greenstein | BVM |
| Streaming bundle churn cut | ~10% | Not attributed | BVM |
| Value perception drop post-purchase | 15% | Research referenced | MF, BVM |
| Over-provisioning recoverable costs | 20-30% | Not attributed | BVM |
| Retention vs acquisition cost ratio | 6-7x | Not attributed | TPM |
| Personalized experiences retention boost | 30% | Not attributed | TPM |
| Dead weight threshold | <20% usage = dead weight | Simon-Kucher implied | MF |
| Feature utilization target | >60% of features by >60% of subs | Industry practice | MF |
| NPS: industry average | 15-25 | Not attributed | MF |
| NPS: premium target | >30 | Not attributed | MF |
| CLV: premium vs base target | 2-4x | Not attributed | MF, TAF |
| CAC payback target | <6 months | Not attributed | MF, TAF |
| Time to break-even | 6-18 months | Not attributed | MF, TAF |
| Churn reduction impact target | 25-35% | Not attributed | MF |
| Content/VAS cost ratio target | <30% of ARPU delta | Not attributed | MF |
| Cannibalization rate target | <50% of premium subs | Not attributed | MF |
| Penetration rate target | 10-25% of eligible base | Not attributed | MF, TAF |
| Migration rate target | 2-5% monthly | Not attributed | MF |
| Component activation target (Leaders) | >70% within 30 days | Not attributed | MF |
| Component activation target (Fillers) | >40% within 30 days | Not attributed | MF |
| CSAT target | >80% | Not attributed | MF |
| Price-value ratio target | >1.0 | Not attributed | MF |
| WTP validation rule | Price < 80% of WTP | Not attributed | MF |
| Price as % of disposable income | <3-5% | Not attributed | MF |
| BCG: broadband prices dropped | 60% (2015-2024) | BCG | TAF |
| Bain: NPS improvement from implementation | 30% | Bain case study | TAF |
| Simon-Kucher: price-value ratio decline | -7% in 2024 | Simon-Kucher 2024 | TAF |
| Simon-Kucher: 43% consumers reduced telco spend | 43% | Simon-Kucher 2024 | TAF |
| Unbundling profit decrease | 10.14% | Luo 2023 RAND Journal | TAF |
| Unbundling consumer surplus decrease | 17.18% | Luo 2023 RAND Journal | TAF |
| Freemium conversion rate (apps) | 3-5% (high: 6-8%) | Not attributed | TPM |
| SaaS freemium conversion | 2-5% | Not attributed | TPM |
| Telecom upsell conversion | 5-15% | Not attributed | TPM |
| ARPPU vs ARPU in freemium | 10-20x | Not attributed | TPM |
| Bundle discount effective range | 15-30% | Academic guidance | TPM |
| Simon-Kucher: offer-related criteria in purchase | 50% | Simon-Kucher 2024 | TAF |
| Simon-Kucher: DaaS upselling potential | 6-30% | Simon-Kucher 2024 | TAF |
| Simon-Kucher: 60% positive to free additional benefits | 60% | Simon-Kucher 2024 | TAF |
| Deloitte: telco-mediated subscriptions | 20% -> 25% (2023->2028) | Deloitte 2025 | BVM |
| Deloitte: streaming stacking peak | ~4 services per household | Deloitte 2025 | BVM |

### (b) Telecom-Specific Benchmarks (Do Not Transfer)

| Benchmark | Value | Context |
|-----------|-------|---------|
| ARPU uplift from premium | +12-18% | Telecom industry benchmark |
| Premium ARPU multiple | 1.5-3x base | Telecom industry |
| Prepaid vs postpaid ARPU gap | EUR 25.50 vs EUR 37.70 | GSMA global |
| Russia avg mobile bundle price | ~543 RUB (2023) | Russia market |
| Mobile data CAGR | 6.7% five-year | Russia/global telecom |
| Russia telecom market size | ~$16.4B (2025) | Russia market |
| Mobile VAS market projection | EUR 1,043B by 2029 | Global telecom |
| Triple Play subscriptions | 1.2B (2024) -> 3.5B (2030) | Global telecom |
| Comcast ARPU increase from bundles | 10-15% | US cable/telecom |
| Comcast churn reduction | ~20% | US cable/telecom |
| Verizon churn reduction (quad play) | 15% | US telecom |
| Verizon cross-sell increase | 25% | US telecom |
| FAS fine on MTS | 3B RUB | Russia regulation |
| Price elasticity: mobile subscription | -0.43 | Telecom-specific |
| Price elasticity: firm-specific demand | -0.47 to -1.10 | Telecom-specific |
| UK SVOD via third party | 43% | UK media market |
| 50% US mobile users want super bundling | 50% | US telecom |
| 68% US subscribers pay via indirect channel | 68% | US telecom |
| Avg US subscriber: 5.4 subscriptions | 5.4 | US market |
| STL Partners: loyalty visible in ~40% countries | ~40% | Global telecom |

---

## 3. Case Studies Inventory

### (a) Universal Lessons (Transfer)

| Case | Universal Lesson |
|------|-----------------|
| T-Mobile + Netflix | Simple, universally-valued content inclusion works; wholesale cost pressure over time |
| Reliance Jio ecosystem | Ecosystem integration must be seamless; financial services highest long-term value; aggressive initial pricing unsustainable |
| Rakuten ecosystem | Rewards/loyalty across ecosystem is powerful retention; mobile can be loss leader if ecosystem profits |
| SK Telecom T Universe | Physical retail helps complex bundle adoption; wide partner diversification reduces risk |
| Comcast Xfinity | Traditional bundling yields 10-15% ARPU uplift and ~20% churn reduction; price stability enhances perception |
| Verizon quad play | Cross-selling from bundle entry point is key revenue driver |
| Akimbo (failure) | Core perceived value must be highest-quality component; hidden costs destroy trust; bundling cannot compensate poor quality |
| European convergence failures | Must include defensible, hard-to-replicate components; easy replication -> price wars |
| Airtel Black | Premium positioning requires consistent quality across ALL bundled services |
| Disney+ bundle | Bundled subscribers 59% less likely to churn than standalone |

### (b) Telecom-Specific Cases (Context Only)

| Case | Why Telecom-Specific |
|------|---------------------|
| McKinsey Asian operator pricing | Mobile price-per-minute optimization |
| BCG Bharti Airtel premiumization | Feature phone -> smartphone, prepaid -> postpaid migration |
| Simon-Kucher CEE tariff mapping | 31 telecom providers in 9 countries |
| ADL 5G pricing strategy | 5G network slicing monetization |
| Bango super bundling | Telecom as aggregator of OTT subscriptions |

---

## 4. Research Methods Inventory

### (a) Universal Methods

| Method | Application | Source |
|--------|------------|--------|
| Choice-Based Conjoint (CBC) | Measure WTP for bundle components and configurations | BVM, TAF, TPM |
| Van Westendorp PSM | Establish acceptable price ranges | BVM, TPM |
| Gabor-Granger | Pinpoint exact WTP at specific price points | BVM, TPM |
| MaxDiff Scaling | Rank component importance | BVM |
| BCG Pathways Conjoint | Qualitative + quantitative pricing research | TPM |
| Economic Value Estimation (EVE) | Quantify tangible economic benefit per component | TPM |
| Market Simulator | Forecast market share, revenue, cannibalization for different configs | TPM |
| A/B Testing | In-market validation of price points | TPM |
| Pilot Launch | Limited geography before national rollout | TPM |

### (b) Telecom-Specific Methods

| Method | Why Telecom-Specific |
|--------|---------------------|
| GSMA tariff basket approach | Standard usage baskets (100MB, 500MB, etc.) |
| GSMA MCI scoring | Mobile Connectivity Index methodology |
| TM Forum eTOM process model | Telecom-specific lifecycle management |
| Tarifica pricing data | Telecom plan pricing database |

---

## 5. Summary Statistics

| Category | Universal | Adaptable | Telecom-Only |
|----------|-----------|-----------|--------------|
| Frameworks | 34 | 15 | 8 |
| Benchmarks | 50+ | 0 | 20+ |
| Case Studies | 10 | 0 | 5 |
| Research Methods | 9 | 0 | 4 |

**Key finding:** The vast majority of the methodology is either directly universal or adaptable. The telecom-specific elements are mostly surface-level (metric names, specific numbers, industry jargon) rather than structural. The underlying evaluation logic transfers cleanly to any complex bundled product.
