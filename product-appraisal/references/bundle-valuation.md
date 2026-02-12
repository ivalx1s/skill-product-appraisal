# Bundle Valuation Reference

Methods, frameworks, and benchmarks for evaluating bundled products and multi-component offerings. Universal -- applies to any industry with complex product bundles.

**CLI:** Bundle calculations available via `appraise calc bundle <function>`. Key functions: `classify` (L/F/K), `dead_weight`, `cross_subsidy`, `component_activation`, `multi_component_usage`.

---

## Table of Contents

1. [Bundle Discount Theory](#1-bundle-discount-theory)
2. [Mixed vs. Pure Bundling](#2-mixed-vs-pure-bundling)
3. [Perceived Value vs. Actual Value](#3-perceived-value-vs-actual-value)
4. [Leaders / Fillers / Killers Framework](#4-leaders--fillers--killers-framework)
5. [Dead Weight Analysis](#5-dead-weight-analysis)
6. [Cross-Subsidy Models](#6-cross-subsidy-models)
7. [Cannibalization Analysis](#7-cannibalization-analysis)
8. [Verified Benchmarks](#8-verified-benchmarks)
9. [Sources](#9-sources)

---

## 1. Bundle Discount Theory

### 1.1 Adams-Yellen (1976): Bundling as Price Discrimination

Adams and Yellen formalized bundling as a mechanism for price discrimination. The core insight: every buyer is described by a vector of reservation prices for individual goods. The seller cannot observe these reservation prices directly, but bundling allows sorting customers into groups based on their reservation price combinations and extracting consumer surplus that would be unattainable through component pricing alone.

Three strategies:

| Strategy | Description | When Optimal |
|----------|-------------|-------------|
| **Pure components** | Each product sold separately | Low consumer heterogeneity; high correlation of valuations |
| **Pure bundling** | Products available only as a package | High consumer heterogeneity; low/negative correlation |
| **Mixed bundling** | Both bundle and standalone available | Almost always (see McAfee-McMillan-Whinston below) |

**Source:** Adams, W.J. & Yellen, J.L. (1976). "Commodity Bundling and the Burden of Monopoly." *Quarterly Journal of Economics*, 90(3), 475-498. [URL](https://academic.oup.com/qje/article-abstract/90/3/475/1854397)

### 1.2 Schmalensee (1984): Variance Reduction

Schmalensee provided the statistical foundation for why bundling enables better surplus extraction. The key mechanism: consumer heterogeneity is the primary obstacle to extracting full surplus. Bundling reduces the variance of consumer valuations because the variance of the sum is lower than the sum of variances (assuming imperfect positive correlation).

As the number of components increases, the law of large numbers guarantees that the coefficient of variation decreases, making the consumer population more homogeneous in their bundle valuation. This enables the seller to set a single bundle price that captures surplus from a larger share of the market.

**Key properties:**

- Works even when individual component valuations are positively correlated (though negatively correlated valuations amplify the effect).
- The effect strengthens with more components -- larger bundles are more amenable to efficient pricing.
- Does not require the seller to know individual reservation prices.

**Source:** Schmalensee, R. (1984). "Gaussian Demand and Commodity Bundling." *Journal of Business*, 57(1), S211-S230. [URL](https://www.jstor.org/stable/2352937)

### 1.3 McAfee, McMillan & Whinston (1989): Mixed Bundling Optimality

The strongest theoretical result in bundling theory. McAfee, McMillan, and Whinston proved that mixed bundling (offering both the bundle and standalone components) almost always strictly increases profits compared to pure bundling or pure component selling.

**Why it matters for product appraisal:** Any bundle evaluation that considers only pure bundling (no standalone options) is leaving money on the table. The default recommendation is always mixed bundling -- offering the bundle AND standalone components simultaneously.

**Source:** McAfee, R.P., McMillan, J. & Whinston, M.D. (1989). "Multiproduct Monopoly, Commodity Bundling, and Correlation of Values." *Quarterly Journal of Economics*, 104(2), 371-383. [URL](https://academic.oup.com/qje/article-abstract/104/2/371/1854649)

### 1.4 Pricing Psychology in Bundles

**Price anchoring** (Tversky & Kahneman, 1974): The first price consumers encounter anchors their fairness assessment. Displaying individual component prices before the bundle price creates a powerful anchoring effect that increases perceived bundle value. The anchoring mechanism is well-established in behavioral economics.

**Bundle framing** (Wansink et al., 1998): Presenting items as a bundle ("3 for $5") rather than individually ("$1.67 each") boosted sales by approximately 32%, even at mathematically equivalent prices. The framing itself creates perceived value independent of actual savings.

**Reference price distortion:** When multiple products are bundled, consumers lose the ability to accurately assess individual component values. This "price opacity" works in the bundler's favor for premium bundles, as there is no easy market comparison for the combined offering.

**Perceived heterogeneity** (Xia Wei et al., 2025): Bundles with diverse components support higher prices because variety itself signals value. Heterogeneous bundles (spanning different product categories) can command premiums over homogeneous bundles (multiple similar items).

---

## 2. Mixed vs. Pure Bundling

### 2.1 Comparison

| Dimension | Pure Bundling | Mixed Bundling | Pure Components |
|-----------|--------------|----------------|-----------------|
| **Availability** | Bundle only | Bundle + standalone | Standalone only |
| **Consumer choice** | None (take-it-or-leave-it) | Full (choose bundle or any combination) | Full (choose any combination) |
| **Revenue extraction** | Good when valuations are negatively correlated | Almost always optimal (McAfee et al., 1989) | Good when valuations are highly correlated |
| **Consumer surplus** | Lowest (most surplus extracted) | Moderate (some surplus remains) | Highest (least surplus extracted) |
| **Dead weight risk** | Highest (forced components) | Low (customers can self-select) | None |
| **Cannibalization risk** | Moderate (standalone products discontinued) | Requires careful price management | None |
| **Complexity** | Low (one SKU) | High (must manage bundle + standalone pricing simultaneously) | Low (individual SKUs) |

### 2.2 When Pure Bundling Can Work

Pure bundling is justified in narrow conditions:

- Components have near-zero marginal cost (e.g., digital services, SaaS modules).
- Consumer valuations are strongly negatively correlated.
- Market power is sufficient to prevent competitive unbundling.
- The bundle itself becomes a coherent product identity (not "a collection of things").

### 2.3 When to Default to Mixed Bundling

In virtually all other cases, mixed bundling dominates. The theoretical proof (McAfee et al., 1989) is robust across assumptions. Mixed bundling captures:

- **High-value-for-all customers** buy the bundle at a discount vs. sum of parts.
- **High-value-for-one customers** buy their preferred component at standalone price.
- **Marginal customers** who might reject the bundle can still buy a single component.

### 2.4 Unbundling Losses

Empirical research found that unbundling leads to approximately 10% profit decrease and 17% consumer surplus decrease. Both producer and consumer are worse off when bundles are dismantled.

**Source:** Luo, Y. (2023). "Bundling and nonlinear pricing in telecommunications." *RAND Journal of Economics*, 54(2), 268-298. [URL](https://onlinelibrary.wiley.com/doi/10.1111/1756-2171.12437)
**Note:** These are empirical results from one industry. The direction is consistent with theory, though exact magnitudes will vary.

---

## 3. Perceived Value vs. Actual Value

### 3.1 Three-Dimension Value Model

A comprehensive meta-analysis (Blut et al., 2024; 687 articles, 780 samples, 357,247 customers) established that customer perceived value operates through three dimensions:

| Dimension | Description | Bundle Implication |
|-----------|-------------|-------------------|
| **Benefits** | Functional, emotional, social gains from the product | Each component must deliver tangible benefits to be perceived as valuable |
| **Sacrifices** | Monetary cost, time, effort, risk | Total bundle sacrifice must feel proportionate to total perceived benefits |
| **Overall value** | Net assessment integrating benefits and sacrifices | The overall evaluation -- not a simple sum of individual component ratings |

The most integrative model incorporating all three dimensions performs best at predicting purchase behavior.

**Source:** Blut, M., Chaney, D., Lunardo, R., Mencarelli, R. & Grewal, D. (2024). "Customer Perceived Value: A Comprehensive Meta-analysis." *Journal of Service Research*, 27(4), 501-524. [URL](https://journals.sagepub.com/doi/10.1177/10946705231222295)

### 3.2 The "Highest-Priced Item" Heuristic

When products are bundled, consumers mentally evaluate the package based on the perceived value of the most valued or highest-priced component. Everything else is processed as a bonus. This means:

- The **Leader** component (see section 4) determines the bundle's perceived value ceiling.
- Additional components contribute diminishing marginal perceived value.
- If no single component is clearly the highest-value item, the bundle lacks a cognitive anchor and feels like a "random collection."

**Practical implication:** Always ensure the bundle has one unmistakable Leader that alone justifies a significant portion of the bundle price.

### 3.3 The Dilution Effect

Shaddy and Fishbach (2017) demonstrated that consumers perceive bundles as gestalt units. The actual mechanism:

- Consumers **pay less** for items added to a bundle (additions yield diminished perceived gain).
- Consumers **demand more compensation** for items removed from a bundle (removals trigger amplified perceived loss).

The practical consequence: adding low-value components to a premium bundle can reduce total willingness-to-pay below what consumers would pay for the high-value component alone. This is the primary theoretical justification for eliminating "Killer" components (see section 4).

**Source:** Shaddy, F. & Fishbach, A. (2017). "Seller Beware: How Bundling Affects Valuation." *Journal of Marketing Research*. [URL](https://www.anderson.ucla.edu/documents/areas/fac/marketing/Seminars/Fall%202017/Shaddy%20%20Fishbach%20-%20How%20Bundling%20Affects%20Valuation%20(job%20market%20paper).pdf)

### 3.4 Post-Purchase Value Perception

Research in hospitality found approximately 15% value perception drop when bundle promises are unmet. Bundled services must deliver on promises -- including services that are hard to access or of poor quality is worse than not including them.

**Status:** [Partially Verified -- domain-specific, hospitality research]

### 3.5 Option Value of Unused Components

Not all unused components are pure dead weight. Unused services can still contribute to perceived value through **option value** -- the knowledge they could be used if needed. This is strongest for:

- Aspirational features (premium status, exclusive access).
- Insurance-like features (protection services, emergency coverage).
- Infrequent-but-high-value features (travel benefits, concierge services).

The test is empirical: does the component's removal increase or decrease WTP? If removal decreases WTP despite low usage, the component has genuine option value.

---

## 4. Leaders / Fillers / Killers Framework

A classification framework for bundle components, originating from Simon-Kucher's consulting practice and now broadly adopted in bundle strategy.

### 4.1 Definitions

| Role | Definition | Target Count | Activation Target |
|------|-----------|-------------|-------------------|
| **Leader** | High perceived value. Drives purchase intent. Customers cite this as the primary reason for buying the bundle. | 2-3 per bundle | >70% activation within 30 days [Practitioner Guidance] |
| **Filler** | Adds perceived value at low marginal cost. Enhances the "bonus" perception. Not the purchase driver, but makes the decision easier. | 3-5 per bundle | >40% activation within 30 days [Practitioner Guidance] |
| **Killer** | Reduces WTP, confuses the offer, attracts wrong customers, or triggers the dilution effect. | **0 per bundle** | N/A -- should be removed |

**Source:** Simon-Kucher publications on bundling strategy. [URL](https://www.simon-kucher.com/en/insights/future-proof-telco-marketing-6-strategies-best-practices)

### 4.2 Component Assessment Table

Use this table to classify each component in a bundle:

| Component | Perceived Value (1-5) | Standalone WTP | Marginal Cost | Usage Forecast | Classification | Rationale |
|-----------|----------------------|----------------|---------------|----------------|----------------|-----------|
| *(name)* | | | | | Leader / Filler / Killer | |

**Classification rules:**

| If... | Then... |
|-------|---------|
| High perceived value AND drives purchase intent | **Leader** |
| Moderate perceived value AND low marginal cost | **Filler** |
| Low perceived value AND high marginal cost | **Killer** -- remove |
| Low perceived value AND low marginal cost | **Filler** (if it has option value) or **Killer** (if it triggers dilution) |
| High perceived value BUT high access constraints (geography, credentials, etc.) | **Leader for eligible segment**, **Dead weight for others** -- consider swappable alternatives |
| Removing it increases WTP (Shaddy & Fishbach test) | **Killer** -- remove immediately |

### 4.3 Ideal Bundle Composition

```
IDEAL:     2-3 Leaders + 3-5 Fillers + 0 Killers
CAUTION:   1 Leader + 6+ Fillers (weak value proposition, "random collection" risk)
DANGER:    0 clear Leaders (nothing drives purchase intent)
DANGER:    Any Killers present (dilution effect active)
```

### 4.4 Dead Weight vs. Killer Distinction

Not all low-usage components are Killers. The distinction:

| Component Type | Usage | WTP Impact | Action |
|---------------|-------|------------|--------|
| **Dead weight with option value** | Low usage | Removing decreases WTP | Keep (it contributes perceived value despite low activation) |
| **True dead weight** | Low usage | Removing has no WTP impact | Consider removing to reduce cost; or keep if cost is negligible |
| **Killer** | Low usage | Removing *increases* WTP | Remove immediately |

---

## 5. Dead Weight Analysis

### 5.1 Definition

Dead weight: components included in a bundle that customers rarely or never use. The threshold: components used by **less than 20% of bundle customers** within 3 months of activation.

### 5.2 Analysis Method

**Step 1: Usage measurement.** Track activation and ongoing usage per component. Define "active use" per component type (e.g., logged in at least once per month, completed at least one transaction, etc.).

**Step 2: Classification.** Apply the 20% threshold:

| Usage Rate | Classification |
|-----------|---------------|
| >60% monthly active | Core component (Leader or strong Filler) |
| 20-60% monthly active | Moderate usage (likely Filler) |
| <20% monthly active | Dead weight candidate |

**Step 3: WTP impact test.** For each dead weight candidate, test whether removal increases, decreases, or has no effect on WTP (see section 3.3). This distinguishes option-value dead weight from true dead weight from Killers.

**Step 4: Cost-benefit assessment.** For each dead weight candidate:

| Factor | Question |
|--------|----------|
| Marginal cost | What does this component cost per customer per month? |
| Option value | Does availability contribute to premium signaling or peace of mind? |
| Dilution risk | Does this component trigger the dilution effect (Shaddy & Fishbach)? |
| Swappability | Could this be replaced with a customer-chosen alternative? |
| Removal risk | Would removal generate negative PR or breach expectations? |

**Step 5: Decision.**

| WTP Impact | Cost | Action |
|-----------|------|--------|
| Removal decreases WTP | Any | Keep (genuine option value) |
| Removal neutral | Low | Keep (no harm, low cost) |
| Removal neutral | High | Remove or make optional |
| Removal increases WTP | Any | Remove (Killer) |

### 5.3 Dead Weight Threshold

The aggregate dead weight ratio for a bundle should stay below **40%** -- meaning fewer than 40% of components should be dead weight.

This is derived from Simon-Kucher's finding that approximately 60% of customers respond positively to free additional benefits (inverse: ~40% are neutral to negative toward any given component).

**Source:** Simon-Kucher Global Telecommunications Study, 2024. [URL](https://www.simon-kucher.com/en/insights/2024-telco-growth-strategies-brand-portfolio-and-pricing-insights)

### 5.4 Over-Provisioning Waste

Industry analysis suggests 20-30% of feature/service costs may be wasted on capabilities customers do not value. This represents recoverable cost that could either improve margins or fund higher-value components.

**Status:** [Practitioner Guidance -- broadly cited in TEM and cloud infrastructure contexts, no single primary source]

### 5.5 Swappable Components as Mitigation

When dead weight is unavoidable (diverse customer base with divergent needs), offering swappable components -- letting customers choose N of M options -- converts a portion of initially disinterested customers into engaged users. Research shows that approximately 60% of customers respond positively to free additional benefits in bundles.

**Practical guidance:** Design bundles with a fixed core (Leaders) plus a choice layer (customer selects from a menu of Fillers).

---

## 6. Cross-Subsidy Models

### 6.1 Definition

Cross-subsidization: using profits from high-margin components to offset losses on low-margin or loss-making components within the same bundle. This is a deliberate structural choice, not an accident.

### 6.2 Cross-Subsidy Flows

```
HIGH MARGIN (subsidy sources)           LOW MARGIN (subsidy recipients)
--------------------------------------  --------------------------------------
Core product with low marginal cost     Premium features with high delivery cost
Financial products / payment fees       Content licensing
Insurance / protection services         Physical-world services (gym, dining)
Data / analytics services               Concierge / human-operated services
Platform fees                           Partner-dependent services
```

### 6.3 Assessment Framework

For each component, calculate:

| Component | Revenue Contribution | Direct Cost | Net Margin | Role |
|-----------|---------------------|-------------|------------|------|
| *(name)* | | | Positive = **Source** / Negative = **Recipient** | Leader/Filler/Killer |

**Rules of thumb:**

- Subsidy sources should be components with **low marginal cost** and **high perceived value** (ideal Leaders).
- Subsidy recipients must justify their cost through one or more of: churn reduction, brand differentiation, cross-sell enablement, or customer acquisition.
- If subsidized components consume **more than 50% of the incremental margin** from the premium tier, the cross-subsidy is unsustainable (BND-4 criterion).

### 6.4 Cross-Subsidy Sustainability Test

Stress-test the model under adverse conditions:

| Scenario | Test |
|----------|------|
| Partner cost increase +20% | Does the model remain margin-positive? |
| Customer adoption -30% | Does the model remain margin-positive? |
| Key partner exits | Can the component be replaced without destroying the value proposition? |
| Usage exceeds forecast by 2x | Do variable costs blow up the model? |

If the model fails under any single scenario, the cross-subsidy structure needs redesign.

### 6.5 Risk Indicators

Cross-subsidy models fail when:

- The subsidized service costs more than the margin it helps retain.
- Customers perceive the subsidized service as the primary value, creating adverse selection (attracting cost-seekers rather than loyal customers).
- Partner relationships deteriorate (cost increases, quality decreases, availability restrictions).
- The subsidy source erodes (core product commoditizes, competitors undercut).

---

## 7. Cannibalization Analysis

### 7.1 The Dual Effect

Bundling creates two simultaneous revenue effects:

| Effect | Mechanism | Impact |
|--------|-----------|--------|
| **Positive** | New customers adopt the bundle; existing customers upgrade; churn decreases; cross-selling increases | Revenue growth |
| **Negative** | Customers who previously purchased individual components at full price migrate to the discounted bundle | Revenue erosion |

The net effect depends on the ratio of new/upgraded customers to migrating standalone customers.

### 7.2 Measuring Cannibalization

| Metric | Formula | Purpose |
|--------|---------|---------|
| Migration rate | Standalone customers moving to bundle / Total standalone customers | Measures migration velocity |
| Incremental revenue | Bundle RPC - Lost standalone revenue per migrating customer | Net revenue impact per customer |
| Attach rate | Bundle customers using each component / Total bundle customers | Identifies genuinely valued components vs. ignored ones |
| Component-level churn | Churn rate for customers using 3+ components vs. 1-2 | Validates bundling's retention thesis |
| LTV comparison | Bundle customer CLV / Standalone customer CLV over 12-24 month cohorts | Long-term value validation |

### 7.3 Cannibalization Offset: Churn Reduction

Bundling's primary strategic value is often churn reduction. Even thin or negative per-customer margins can be justified if churn reduction delivers sufficient lifetime value improvement.

**Verified churn benchmarks:**

| Context | Churn Reduction | Qualification | Source |
|---------|----------------|---------------|--------|
| Entertainment bundle (Disney+/Hulu/ESPN+) | 59% less likely to churn | 70% of standalone re-churned in 12 months vs. 29% of bundle | Ampere Analysis. [URL](https://www.nexttv.com/news/disney-bundlers-59-less-likely-to-churn-research-company-says-chart) |
| Fixed broadband bundle | Modest: 6.93 vs. 6.15 years retention | Authors characterize the effect as "modest." Mobile: 6.43 vs. 5.89 years. | Prince & Greenstein (2014). [URL](https://host.kelley.iu.edu/riharbau/RePEc/iuk/wpaper/bepp2011-05-prince-greenstein.pdf) |
| Multi-product bundles (general) | 25-35% lower churn | Conservative industry consensus range. Documented cases span 5-15% (modest) to 50%+ (entertainment). | [Practitioner Guidance] |

### 7.4 Dynamic Effects Over Time

- **Short-term:** Mixed bundling enhances revenues for both primary and secondary products.
- **Medium-term:** Bundle customers develop cross-component usage patterns, increasing switching costs and deepening engagement.
- **Long-term:** If bundle services are underutilized, customers may downgrade or churn when novelty wears off. Ongoing component refresh is essential.

---

## 8. Verified Benchmarks

All benchmarks are tagged by verification status from factcheck review.

### [Verified] -- Confirmed with primary source

| Benchmark | Value | Source |
|-----------|-------|--------|
| Mixed bundling profit superiority | Strictly increases profits vs. pure bundling or component selling | McAfee, McMillan & Whinston (1989). [URL](https://academic.oup.com/qje/article-abstract/104/2/371/1854649) |
| Variance reduction via bundling | Bundling reduces valuation variance, enables better surplus extraction | Schmalensee (1984). [URL](https://www.jstor.org/stable/2352937) |
| Bundle framing sales boost | ~32% from multi-unit framing ("3 for $5") | Wansink et al. (1998). [URL](https://www.scribd.com/document/895589393/Wansink-Study-Simplified-Notes) |
| Dilution effect in bundles | Adding low-value items can reduce total bundle WTP | Shaddy & Fishbach (2017). [URL](https://www.anderson.ucla.edu/documents/areas/fac/marketing/Seminars/Fall%202017/Shaddy%20%20Fishbach%20-%20How%20Bundling%20Affects%20Valuation%20(job%20market%20paper).pdf) |
| Heterogeneous bundles support premium pricing | Diverse component bundles command higher prices | Xia Wei et al. (2025). [URL](https://journals.sagepub.com/doi/10.1177/00472875231222263) |
| Perceived value meta-analysis | Three dimensions: benefits, sacrifices, overall (687 articles, 357K customers) | Blut et al. (2024). [URL](https://journals.sagepub.com/doi/10.1177/10946705231222295) |
| GBB middle tier selection | ~66% choose middle option (range varies; some studies give 60-70%) | HBR (Mohammed, 2018). [URL](https://hbr.org/2018/09/the-good-better-best-approach-to-pricing) |
| Entertainment bundle churn reduction | 59% less likely to churn (Disney+/Hulu/ESPN+ bundle) | Ampere Analysis. [URL](https://www.nexttv.com/news/disney-bundlers-59-less-likely-to-churn-research-company-says-chart) |
| Broadband bundle retention | Modest: 6.93 vs. 6.15 years | Prince & Greenstein (2014). [URL](https://host.kelley.iu.edu/riharbau/RePEc/iuk/wpaper/bepp2011-05-prince-greenstein.pdf) |
| Unbundling losses | ~10% profit decrease, ~17% consumer surplus decrease (single industry) | Luo (2023). [URL](https://onlinelibrary.wiley.com/doi/10.1111/1756-2171.12437) |
| Acquisition vs. retention cost ratio | 5-25x more expensive to acquire than retain | HBR (Gallo, 2014). [URL](https://hbr.org/2014/10/the-value-of-keeping-the-right-customers) |
| Personalization revenue lift | 10-15% revenue lift; 10-30% marketing ROI improvement | McKinsey. [URL](https://www.mckinsey.com/capabilities/growth-marketing-and-sales/our-insights/the-value-of-getting-personalization-right-or-wrong-is-multiplying) |
| Freemium conversion (self-serve) | 3-5% average; 6-8% top performers | First Page Sage (2026). [URL](https://firstpagesage.com/seo-blog/saas-freemium-conversion-rates/) |
| Freemium conversion (sales-assisted) | 5-7% average; 10-15% top performers | First Page Sage (2026). [URL](https://firstpagesage.com/seo-blog/saas-freemium-conversion-rates/) |
| Positive response to free additional benefits | ~60% | Simon-Kucher (2024). [URL](https://www.simon-kucher.com/en/insights/2024-telco-growth-strategies-brand-portfolio-and-pricing-insights) |
| Leaders / Fillers / Killers framework | Classification system for bundle components | Simon-Kucher. [URL](https://www.simon-kucher.com/en/insights/future-proof-telco-marketing-6-strategies-best-practices) |

### [Practitioner Guidance] -- Reasonable but unverifiable

| Benchmark | Value | Notes |
|-----------|-------|-------|
| BVR threshold | >1.5x adequate; >2.0x strong | Widely used in practice; no primary academic source. Optimal ratio depends on category and competitive context. |
| Dead weight threshold | <40% of components | Derived from Simon-Kucher ~60% positive response (inverse). Reasonable heuristic. |
| Multi-product churn reduction | 25-35% lower churn | Conservative industry consensus. Real range spans 5% (modest) to 59% (entertainment). |
| Removing middle tier revenue loss | ~50% | Industry sources suggest 48-60% range. Consistent with GBB middle tier attracting ~66% of customers. |
| Over-provisioning waste | 20-30% of feature costs | Broadly cited in TEM and cloud infrastructure; no single primary source. |
| Post-purchase disappointment | ~15% perceived value drop | Hospitality bundle research. Domain-specific. |
| Premium CLV multiple | 2-4x base CLV | Widely cited; "VIP subscribers 3x LTV of basic" appears in industry sources. |
| Feature utilization target | >60% of features by >60% of customers | Common product management heuristic; no research-backed validation. |
| Partner cost ratio ceiling | <30% of premium revenue delta | Not externally validated. Reasonable structural constraint. |
| Leader activation target | >70% within 30 days | Not externally validated. Reasonable operational target. |
| Filler activation target | >40% within 30 days | Not externally validated. Reasonable operational target. |

---

## 9. Sources

### Academic Papers

- Adams, W.J. & Yellen, J.L. (1976). "Commodity Bundling and the Burden of Monopoly." *Quarterly Journal of Economics*, 90(3), 475-498. [URL](https://academic.oup.com/qje/article-abstract/90/3/475/1854397)
- Schmalensee, R. (1984). "Gaussian Demand and Commodity Bundling." *Journal of Business*, 57(1), S211-S230. [URL](https://www.jstor.org/stable/2352937)
- McAfee, R.P., McMillan, J. & Whinston, M.D. (1989). "Multiproduct Monopoly, Commodity Bundling, and Correlation of Values." *Quarterly Journal of Economics*, 104(2), 371-383. [URL](https://academic.oup.com/qje/article-abstract/104/2/371/1854649)
- Tversky, A. & Kahneman, D. (1974). "Judgment Under Uncertainty: Heuristics and Biases." *Science*, 185(4157), 1124-1131.
- Wansink, B., Kent, R.J. & Hoch, S.J. (1998). "An Anchoring and Adjustment Model of Purchase Quantity Decisions." *Journal of Marketing Research*, 35(1), 71-81. [URL](https://www.scribd.com/document/895589393/Wansink-Study-Simplified-Notes)
- Shaddy, F. & Fishbach, A. (2017). "Seller Beware: How Bundling Affects Valuation." *Journal of Marketing Research*. [URL](https://www.anderson.ucla.edu/documents/areas/fac/marketing/Seminars/Fall%202017/Shaddy%20%20Fishbach%20-%20How%20Bundling%20Affects%20Valuation%20(job%20market%20paper).pdf)
- Luo, Y. (2023). "Bundling and nonlinear pricing in telecommunications." *RAND Journal of Economics*, 54(2), 268-298. [URL](https://onlinelibrary.wiley.com/doi/10.1111/1756-2171.12437)
- Blut, M., Chaney, D., Lunardo, R., Mencarelli, R. & Grewal, D. (2024). "Customer Perceived Value: A Comprehensive Meta-analysis." *Journal of Service Research*, 27(4), 501-524. [URL](https://journals.sagepub.com/doi/10.1177/10946705231222295)
- Xia Wei, Yu, S. & Li, X. (2025). "Price it High if it is Varied: Perceived Heterogeneity and the Effectiveness of Discount Framing Strategies for Travel Packages." *Journal of Travel Research*. [URL](https://journals.sagepub.com/doi/10.1177/00472875231222263)
- Prince, J. & Greenstein, S. (2014). "Does Service Bundling Reduce Churn?" *Journal of Economics & Management Strategy*. [URL](https://host.kelley.iu.edu/riharbau/RePEc/iuk/wpaper/bepp2011-05-prince-greenstein.pdf)

### Industry and Consulting Sources

- Simon-Kucher (2024). "2024 Telco Growth Strategies: Brand, Portfolio, and Pricing Insights." [URL](https://www.simon-kucher.com/en/insights/2024-telco-growth-strategies-brand-portfolio-and-pricing-insights)
- Simon-Kucher. "Future-proof telco marketing: 6 strategies & best practices." [URL](https://www.simon-kucher.com/en/insights/future-proof-telco-marketing-6-strategies-best-practices)
- Mohammed, R. (2018). "The Good-Better-Best Approach to Pricing." *Harvard Business Review*. [URL](https://hbr.org/2018/09/the-good-better-best-approach-to-pricing)
- Gallo, A. (2014). "The Value of Keeping the Right Customers." *Harvard Business Review*. [URL](https://hbr.org/2014/10/the-value-of-keeping-the-right-customers)
- McKinsey. "The value of getting personalization right -- or wrong -- is multiplying." [URL](https://www.mckinsey.com/capabilities/growth-marketing-and-sales/our-insights/the-value-of-getting-personalization-right-or-wrong-is-multiplying)
- Ampere Analysis. "Disney+ bundlers 59% less likely to churn." [URL](https://www.nexttv.com/news/disney-bundlers-59-less-likely-to-churn-research-company-says-chart)
- First Page Sage (2026). "SaaS Freemium Conversion Rates." [URL](https://firstpagesage.com/seo-blog/saas-freemium-conversion-rates/)
