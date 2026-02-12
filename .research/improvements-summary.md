# Improvements Summary: Telecom-Specific -> Universal Methodology

**Date:** 2026-02-12
**Purpose:** Changelog of what was verified, corrected, removed, and improved during generalization.

---

## A. What Was Verified (Kept As-Is or With Source URLs Added)

### Academic Foundations (All Verified)
1. **Adams-Yellen (1976)** -- bundling as price discrimination via reservation prices. Core of our bundle pricing theory.
2. **Schmalensee (1984)** -- variance reduction through bundling. Explains why bundles enable better price extraction.
3. **McAfee-McMillan-Whinston (1989)** -- mixed bundling optimality proof. Justifies our recommendation for mixed (not pure) bundling.
4. **Blut et al. (2024)** -- perceived value meta-analysis (687 articles, 780 samples). Validates our three-dimension value assessment.
5. **Xia Wei et al. (2025)** -- heterogeneous bundles support premium pricing. Validates our approach to diverse bundle composition.
6. **Shaddy & Fishbach (2017)** -- dilution effect in bundles. Validates our dead weight / Killer identification methodology.
7. **Luo (2023)** -- unbundling losses (10.14% profit, 17.18% consumer surplus). Empirical evidence for bundle value.
8. **Prince & Greenstein (2014)** -- bundling reduces churn modestly. Validates churn dimension with measured (not inflated) expectations.
9. **Wansink et al. (1998)** -- bundle framing boosts sales 32%. Validates our framing/anchoring guidance.

### Frameworks (All Verified)
1. **Leaders / Fillers / Killers** -- Simon-Kucher framework, verified in their publications.
2. **Good-Better-Best** -- three-tier pricing, verified via HBR (Rafi Mohammed, 2018) and multiple sources.
3. **Van Westendorp PSM** -- standard method, Peter Van Westendorp (1976), verified.
4. **Gabor-Granger** -- standard method, verified.
5. **Conjoint Analysis (CBC)** -- gold standard for WTP research, verified across academic and consulting sources.
6. **Go/No-Go Decision Matrix** -- standard decision framework, universal.

### Specific Benchmarks (Verified With Sources)
1. **Disney+ bundle: 59% less likely to churn** (Ampere Analysis) -- strong evidence for bundling's churn effect.
2. **Deloitte 2025: streaming peak ~4 services/household; telco-mediated subscriptions 20% -> 25% by 2028** -- verified.
3. **Bain NPS implementation: >30 percentage points improvement in call centers** -- verified with nuance.
4. **Simon-Kucher: 7% price-value ratio decline in telco (2024)** -- verified (telecom-specific example).
5. **GBB: ~66% choose middle tier** -- verified as approximate heuristic.
6. **Freemium conversion: 3-5% self-serve, 5-7% sales-assisted** -- verified (First Page Sage, 2026).
7. **Acquisition vs. retention cost: 5-25x** (HBR) -- verified with corrected range.

---

## B. What Was Corrected

| # | Original Claim | Correction | Reason |
|---|---------------|------------|--------|
| 1 | NPS industry avg: 15-25 for telecom | Telecom NPS avg is ~29-30 (2023-2024). Among lowest industries but higher than claimed. | CustomerGauge 2025 data. Original figure may be outdated (pre-2020). |
| 2 | Retention vs. acquisition: 6-7x more expensive | HBR states 5-25x. The 6-7x is mid-range but not the canonical cite. | Corrected to cite HBR primary source. |
| 3 | Personalized experiences: 30% retention boost | McKinsey documents 10-15% revenue lift and 10-30% marketing ROI improvement. Specific examples: Nike 40%, Spotify 40%. The "30% retention" is not directly sourced. | Replaced with verified McKinsey numbers. |
| 4 | Shaddy & Fishbach: "bundling expensive + cheap lowers value below expensive alone" | The actual mechanism is more nuanced: consumers pay less for items ADDED to bundles and demand more for items REMOVED. The "dilution" interpretation is directionally correct but oversimplified. | Nuanced the description to match actual findings. |
| 5 | Prince & Greenstein: broadband churn reduction | The authors themselves characterize the effect as "modest" (6.93 vs 6.15 years). Original doc presented it without this qualification. | Added qualification. |

---

## C. What Was Removed or Reworded (Unverifiable)

| # | Original Claim | Action | Reason |
|---|---------------|--------|--------|
| 1 | "Anchoring increases perceived value by approximately 32%" (attributed to Tversky/Kahneman) | Removed specific %. Reworded to "Price anchoring significantly increases perceived bundle value." | The 32% figure is not from T&K. It appears in secondary sources without primary attribution. Wansink's 32% is a different experiment (multi-unit framing, not anchoring). |
| 2 | "BCG: bundles perceived as 20-25% more valuable than sum of parts" | Removed BCG attribution. Kept directional claim. | No BCG publication found. Claim appears in secondary aggregation sites only. |
| 3 | "HBR: mixed bundling outperforms pure bundling by 25-35%" | Removed specific % and HBR attribution. Cited McAfee-McMillan-Whinston for theoretical proof instead. | The 25-35% figure cannot be traced to a specific HBR article. |
| 4 | "Simon-Kucher: 13% of disinterested customers convert with swappable benefits" | Removed specific %. Reworded to "Offering swappable benefits converts a portion of initially disinterested customers." | Not found in accessible Simon-Kucher publications. |
| 5 | "ICE Model (Simon-Kucher)" brand framework | Removed Simon-Kucher attribution. Kept framework as "brand permission assessment: Image, Communication, Execution." | No Simon-Kucher source found for this specific framework name. |
| 6 | "McKinsey: bundling increases revenues 10-30%" | Softened to "Industry research suggests 10-30%." McKinsey verified for personalization (10-15%), not bundling specifically. | Direct McKinsey bundling publication not found for this exact claim. |

---

## D. What Was Generalized (Telecom -> Universal)

### Terminology Replacements

| Telecom Term | Universal Term | Applies To |
|-------------|---------------|------------|
| ARPU (Average Revenue Per User) | Revenue Per Customer (RPC) | Any subscription/recurring product |
| AMPU (Average Margin Per User) | Gross Margin Per Customer | Any product |
| Subscriber | Customer | Universal |
| VAS revenue share | Service/add-on revenue share | Any product with modules/add-ons |
| Revenue per GB | *Removed* | Telecom-only, no universal equivalent |
| Regional Viability (Dimension 7) | Market Reach | Covers geographic, demographic, platform, channel constraints |
| Network QoS differentiation | Service quality differentiation | Any product with quality tiers |
| Churn rate (telecom context) | Churn/attrition rate | Any subscription/recurring product |

### Structural Generalizations

| Original Structure | Generalized Structure | What Changed |
|-------------------|----------------------|-------------|
| 7 Assessment Dimensions | 7 Assessment Dimensions | Dimension 7 renamed "Market Reach" from "Regional Viability" |
| 48 criteria (PMF-1 through REG-7) | 48 criteria (PMF-1 through MR-7) | REG-* -> MR-*; telecom-specific criteria content replaced |
| 6 KPI categories | 6 KPI categories | Same categories, telecom KPIs replaced with universal ones |
| Evaluation Logic Flow | Evaluation Logic Flow | "Regional Feasibility" -> "Market Reach Assessment" |
| Go/No-Go Matrix | Go/No-Go Matrix | Unchanged -- already universal |

### Thresholds Marked "Calibrate Per Industry"

| Threshold | Telecom Value | Why Calibration Needed |
|-----------|--------------|----------------------|
| Revenue uplift | +12-18% | SaaS may see 50-100%+; physical goods 10-20% |
| CAC payback | <6 months | SaaS: 12-18 months; consumer apps: 1-3 months |
| Penetration rate | 10-25% | Depends on market maturity, pricing, segment size |
| Affordability threshold | <3-5% of income | Telecom = utility; luxury/SaaS have different norms |
| NPS target | >30 premium | NPS varies: tech ~40-60, retail ~50-70, financial ~20-40 |
| Time to break-even | 6-18 months | Depends on investment, margin structure |
| Content/partner cost ratio | <30% | Depends on partner economics, margin structure |
| Feature utilization | >60% | Reasonable but unverified; may differ by product complexity |

---

## E. Proposed Improvements to Methodology

### 1. Source Attribution Standard
**Problem:** Original methodology cited "McKinsey says X" or "research shows Y" without URLs.
**Improvement:** Every external claim in METHODOLOGY.md now has one of:
- Direct source URL
- "Practitioner guidance" label (for reasonable but unverifiable heuristics)
- "Calibrate per industry" label (for industry-specific thresholds)
- "Calculated" label (for numbers we derived)

### 2. Nuanced Churn Benchmarks
**Problem:** Original cited multiple churn reduction figures (25-35%, 59%, 10%) without context on when each applies.
**Improvement:** Present a range with qualifiers:
- Modest effect (5-15%): general bundling, mature markets (Prince & Greenstein)
- Moderate effect (25-35%): well-designed multi-product bundles (industry consensus)
- Strong effect (50%+): tightly integrated entertainment bundles (Disney+/Ampere Analysis)
- Note: Churn reduction is the primary financial justification for bundling, but magnitude depends on bundle design and market.

### 3. Dead Weight Treatment
**Problem:** Original presented dead weight as purely negative.
**Improvement:** Add nuance from research: dead weight has both negative effects (dilution, cost waste) AND positive effects (option value, premium signaling). The threshold (<40% dead weight) remains, but note that some dead weight may be intentional if it serves signaling or option value purposes.

### 4. Calibration Guidance
**Problem:** All thresholds were set for telecom.
**Improvement:** Added explicit "calibrate per industry" markers with examples of how the same metric behaves differently across industries (telecom, SaaS, retail, financial services).

### 5. Dilution Effect Precision
**Problem:** Original oversimplified Shaddy & Fishbach as "cheap items lower bundle value."
**Improvement:** Correct mechanism: consumers resist altering bundles as gestalt units -- they pay less for additions and demand more for removals. The practical implication (don't include low-value components in premium bundles) remains valid, but the underlying psychology is about gestalt perception, not simple averaging.

### 6. Mixed Bundling as Default Recommendation
**Problem:** Original mentioned mixed bundling but didn't strongly position it.
**Improvement:** McAfee-McMillan-Whinston's proof is the strongest theoretical result in bundling theory. Make mixed bundling (offering bundle AND standalone components) the explicit default recommendation, not just an option.

---

## F. What Remains Telecom-Specific (Not in Universal Methodology)

The following elements were documented in the audit but deliberately excluded from the universal methodology:

1. ARPU, Revenue per GB, AMPU (telecom-specific metrics)
2. GSMA MCI methodology (telecom industry body)
3. TM Forum eTOM framework (telecom process model)
4. FAS regulatory risk (Russia-specific)
5. Double/Triple/Quad Play evolution (telecom history)
6. Speed vs. Volume differentiation (telecom-specific)
7. Prepaid vs. Postpaid dynamics (telecom billing model)
8. Device-as-a-Service (telecom upselling)
9. All MTS-specific numbers (82.2M subs, 17.5M ecosystem, KION, etc.)
10. Russian market data (543 RUB avg, operator comparisons)
11. Zero-rating practices (net neutrality)
12. SIM registration requirements (telecom regulation)
13. 5G/network slicing monetization (telecom infrastructure)

These are preserved in the audit inventory (`.research/audit-inventory.md`) for reference when the methodology is applied to telecom specifically.
