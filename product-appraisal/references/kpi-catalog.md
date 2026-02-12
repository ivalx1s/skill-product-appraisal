# Universal KPI Catalog

Quick-reference catalog of KPIs for premium product and bundle assessment. Organized into 6 categories. Every benchmark is either fact-checked with a source URL or marked as practitioner guidance / calibrate-per-industry.

**Usage:** Look up KPIs by category. Use `appraise` CLI for all calculations (see CLI commands below each table). Use benchmark ranges as starting points, then calibrate for your specific industry using the calibration notes at the bottom.

**CLI:** All KPIs with formulas can be computed via `appraise calc <module> <function> --input data.json`. Run `appraise q 'schema()'` for the full API contract.

---

## Category 1: Revenue KPIs

| KPI | Definition | Formula | Benchmark Range | Source / Notes |
|-----|-----------|---------|-----------------|----------------|
| Revenue Per Customer (RPC) | Average revenue generated per customer over a period | `Total Product Revenue / Average Customer Count` | Premium RPC: 1.5-3x base RPC | Calibrate per industry. Replaces telecom-specific ARPU. |
| Revenue Uplift | Delta between premium customer revenue and base/prior revenue | `(Premium RPC - Base RPC) / Base RPC` | Calibrate per industry | Industry research suggests effective bundling can increase revenues by 10-30%. McKinsey documents 10-15% revenue lift from personalization specifically. [McKinsey](https://www.mckinsey.com/capabilities/growth-marketing-and-sales/our-insights/the-value-of-getting-personalization-right-or-wrong-is-multiplying) |
| Blended Revenue Impact | Net effect on total customer base revenue after cannibalization | `Total Revenue (post-launch) - Total Revenue (pre-launch)` | Must be positive after accounting for cannibalization | Universal. If negative, premium tier is destroying value. |
| Gross Margin Per Customer | Revenue minus cost of goods sold per customer | `(Revenue - COGS) / Average Customer Count` | Must cover all licensing, partner, and delivery costs | Universal unit economics metric. |
| Service/Add-on Revenue Share | Portion of revenue from value-added services or modules | `Add-on Revenue / Total Revenue` | Calibrate per industry | SaaS services: ~30-60%; retail add-ons: ~5-15%. Growing trend is positive. |
| Ecosystem Revenue Share | Revenue from adjacent/ecosystem operations vs. total | `Ecosystem Revenue / Total Revenue` | Positive trend | Applicable to platform businesses and ecosystem plays. |
| Revenue Growth Rate | Year-over-year growth in premium segment revenue | `(Revenue_t - Revenue_t-1) / Revenue_t-1` | Should outpace overall company average | Universal. |

> **CLI:** `appraise calc financial revenue_uplift`, `appraise calc financial gross_margin`, `appraise calc financial unit_economics`, `appraise calc customer service_revenue_share`, `appraise calc customer revenue_growth`

---

## Category 2: Customer KPIs

| KPI | Definition | Formula | Benchmark Range | Source / Notes |
|-----|-----------|---------|-----------------|----------------|
| Churn Rate | Customers lost as a proportion of total customers per period | `Customers Lost in Period / Total Customers at Start of Period` | Premium churn < 50% of base churn | Churn reduction from bundling ranges from modest (5-15%) to dramatic (50%+) depending on bundle design. Disney+ bundle subscribers are 59% less likely to churn than standalone. [Ampere Analysis / NextTV](https://www.nexttv.com/news/disney-bundlers-59-less-likely-to-churn-research-company-says-chart) |
| Retention Rate | Complement of churn rate | `1 - Churn Rate` | Higher than base product | Universal complement to churn. |
| NPS (Net Promoter Score) | Promoter-detractor spread from survey | `% Promoters (9-10) - % Detractors (0-6)` | Calibrate per industry | Telecom ~29-30; tech/SaaS ~40-60; retail ~50-70; financial services ~20-40. [CustomerGauge](https://customergauge.com/benchmarks/blog/telecommunications-nps-benchmarks-and-cx-trends). Originated by Bain. |
| CSAT (Customer Satisfaction) | Satisfaction survey score | `Satisfied Responses / Total Responses` | >80% for premium | Universal. Measure via post-interaction surveys. |
| CES (Customer Effort Score) | Ease of interaction score from survey | `Sum of Effort Ratings / Total Responses` | Lower = better; premium should outperform base | Universal. Premium experience should minimize friction. |
| CLV / LTV (Customer Lifetime Value) | Total value of a customer over their entire relationship | `RPC x Gross Margin % x Average Customer Lifespan` | Premium CLV: 2-4x base (practitioner target) | Widely used target for premium tiers in subscription businesses. Exact multiplier depends on margin structure and retention. |
| CAC (Customer Acquisition Cost) | Cost to acquire one new customer | `Total Acquisition Spend / New Customers Acquired` | Calibrate per industry | SaaS: 12-18 month payback; consumer apps: 1-3 months. Acquiring new customers costs 5-25x more than retaining existing ones. [HBR](https://hbr.org/2014/10/the-value-of-keeping-the-right-customers) |
| Churn Reduction Impact | Change in churn rate after premium/bundle launch | `(Churn_before - Churn_after) / Churn_before` | 5-50%+ depending on bundle design | Modest: 5-15% (general bundling, [Prince & Greenstein 2014](https://host.kelley.iu.edu/riharbau/RePEc/iuk/wpaper/bepp2011-05-prince-greenstein.pdf)); moderate: 25-35% (multi-product bundles); strong: 50%+ (tightly integrated bundles, [Ampere/Disney+](https://www.nexttv.com/news/disney-bundlers-59-less-likely-to-churn-research-company-says-chart)). |

> **CLI:** `appraise calc customer churn_rate`, `appraise calc customer retention_rate`, `appraise calc customer nps`, `appraise calc customer csat`, `appraise calc customer churn_reduction`, `appraise calc financial clv`, `appraise calc financial cac_payback`

---

## Category 3: Product Performance KPIs

| KPI | Definition | Formula | Benchmark Range | Source / Notes |
|-----|-----------|---------|-----------------|----------------|
| Penetration Rate | Share of total customer base on premium tier | `Premium Customers / Total Customer Base` | Calibrate per industry | Depends on market maturity, pricing, segment size. SaaS freemium-to-paid: 3-5% self-serve, 5-7% sales-assisted. [First Page Sage](https://firstpagesage.com/seo-blog/saas-freemium-conversion-rates/) |
| Upgrade/Migration Rate | Rate of customers moving to premium | `Customers Upgrading / Eligible Base per Period` | Calibrate per industry | Track monthly. Healthy rate depends on pricing gap and value proposition. |
| Cannibalization Rate | Existing customers migrating vs. net new acquired | `Migrated Existing Customers / Total Premium Customers` | Net cannibalization < 50% of premium base | Universal when new product overlaps existing offerings. High cannibalization without revenue uplift signals pricing failure. |
| Cross-sell Rate | Premium customers purchasing additional services | `Premium Customers Buying Add-ons / Total Premium Customers` | Higher than base segment cross-sell rate | Universal. Measures ecosystem stickiness. |
| Feature Utilization Rate | Share of premium features actively used | `Features Used per Customer / Total Available Features` | >60% (practitioner target) | Widely used product management heuristic. Low utilization signals over-provisioning or poor feature-market fit. |
| Time to Break-even | Months until cumulative revenue exceeds cumulative cost per customer | `Month where Cumulative Revenue >= Cumulative Cost` | Calibrate per industry (practitioner guidance: 6-18 months) | Depends heavily on investment level and margin structure. |
| Component Activation Rate | Share of customers activating each bundle component within 30 days | `Customers Activating Component / Total Bundle Customers` | >70% for Leaders, >40% for Fillers | Based on Leaders/Fillers/Killers framework. [Simon-Kucher](https://www.simon-kucher.com/en/insights/future-proof-telco-marketing-6-strategies-best-practices) |
| Attach Rate | Monthly usage of each bundle component | `Customers Using Component Monthly / Total Bundle Customers` | Track per component; declining attach signals dead weight | Universal for multi-component products. |
| Trial-to-Paid Conversion | Share of trial/freemium users converting to paid | `Paid Conversions / Trial Users` | Self-serve: 3-5%; sales-assisted: 5-7%; top performers: 8-15% | [First Page Sage 2026](https://firstpagesage.com/seo-blog/saas-freemium-conversion-rates/) |

> **CLI:** `appraise calc product penetration_rate`, `appraise calc product migration_rate`, `appraise calc product cannibalization_rate`, `appraise calc product feature_utilization`, `appraise calc product component_activation`, `appraise calc product attach_rate`, `appraise calc product trial_conversion`, `appraise calc financial break_even`

---

## Category 4: Premium Segment KPIs

| KPI | Definition | Formula | Benchmark Range | Source / Notes |
|-----|-----------|---------|-----------------|----------------|
| Price-Value Ratio | Customer-perceived value relative to price | `Survey: Perceived Value Score / Actual Price` | >1.0 (customers perceive more value than price paid) | Universal. Declining price-value ratio signals value erosion. Example: Simon-Kucher found 7% price-value ratio decline in one industry in 2024. [Simon-Kucher](https://www.simon-kucher.com/en/insights/2024-telco-growth-strategies-brand-portfolio-and-pricing-insights) |
| WTP (Willingness to Pay) | Maximum price customers would pay | Measured via Van Westendorp PSM, Gabor-Granger, or Conjoint Analysis | Price should be < 80% of WTP | Standard pricing research methods. [Van Westendorp](https://sawtoothsoftware.com/resources/blog/posts/van-westendorp-pricing-sensitivity-meter), [Gabor-Granger](https://sawtoothsoftware.com/resources/blog/posts/gabor-granger-pricing-method), [Conjoint/BCG](https://www.bcg.com/publications/2014/telecommunications-pricing-pathways-conjoint-new-approach-pricing-mobile) |
| Premium Price Index | Premium product price relative to market average | `Premium Price / Market Average Price` | Track and justify with value ratio | Universal for any premium-positioned product. |
| Bundle Attractiveness Score | Composite value of components vs. bundle price | `Sum of Component Standalone Prices / Bundle Price` (= BVR) | >1.5x (practitioner guidance) | See Bundle Economics below for BVR detail. |
| Switching Cost Index | Perceived cost/effort for customer to leave | Survey-based composite score | Higher than any direct competitor | Universal for subscription and platform products. Higher switching costs improve retention but beware regulatory risk. |
| Share of Wallet | Company's share of customer's total category spend | `Company Revenue from Customer / Customer's Total Category Spend` | Growing quarter-over-quarter | Universal. Measure via survey or transaction data. |
| Premiumization Rate | Share of base customers moving to higher tiers over time | `Customers Moving Up / Total Base Customers per Period` | Positive trend | Universal for tiered product lines. |
| Component Engagement Rate | Usage depth/frequency of bundled services | `Usage Events per Component per Customer per Period` | Track vs. standalone benchmarks | Universal for any bundle. Low engagement = potential dead weight. |
| Dead Weight Ratio | Share of components with very low usage | `Components with <20% Monthly Usage / Total Components` | <40% (practitioner guidance, per Simon-Kucher data) | Some dead weight may be intentional (option value, premium signaling) but excessive dead weight erodes perceived value via dilution effect. [Shaddy & Fishbach 2017](https://www.anderson.ucla.edu/documents/areas/fac/marketing/Seminars/Fall%202017/Shaddy%20%20Fishbach%20-%20How%20Bundling%20Affects%20Valuation%20(job%20market%20paper).pdf) |

> **CLI:** `appraise calc pricing bvr`, `appraise calc pricing price_value_ratio`, `appraise calc pricing premium_price_index`, `appraise calc bundle dead_weight`, `appraise calc bundle classify`

---

## Category 5: Bundle Economics KPIs

| KPI | Definition | Formula | Benchmark Range | Source / Notes |
|-----|-----------|---------|-----------------|----------------|
| Bundle Value Ratio (BVR) | Perceived savings from buying the bundle vs. standalone | `Sum of Standalone Prices / Bundle Price` | >1.5x adequate, >2.0x strong (practitioner guidance) | Reasonable heuristic but not externally validated. Optimal ratio depends on category and competitive context. |
| Cross-subsidy Efficiency | Net margin contribution across high-margin and low-margin components | `High-Margin Component Margin - Low-Margin Component Subsidy Cost` | Positive after all subsidies | Universal for mixed-margin bundles. If negative, bundle pricing is unsustainable. |
| Partner/Licensing Cost Ratio | Third-party costs relative to premium revenue delta | `(Partner + Licensing Costs per Customer) / Premium Revenue Delta` | <30% (practitioner guidance) | Applicable when bundle includes third-party components. Industry-specific threshold. |
| Incremental Revenue | Bundle revenue minus lost standalone revenue from migration | `Bundle Revenue per Customer - Lost Standalone Revenue per Customer` | Positive | Universal when bundle replaces existing products. Negative = cannibalization exceeds uplift. |
| Multi-component Usage Rate | Share of customers using 3+ bundle components | `Customers Using 3+ Components / Total Bundle Customers` | >60% (practitioner target) | Universal for multi-component bundles. Low rate signals poor bundle composition or over-provisioning. |

> **CLI:** `appraise calc pricing bvr`, `appraise calc bundle cross_subsidy`, `appraise calc financial cannibalization`, `appraise calc financial incremental_revenue`, `appraise calc bundle multi_component_usage`

---

## Category 6: Market Context KPIs

| KPI | Definition | Data Source | Benchmark Range | Notes |
|-----|-----------|-------------|-----------------|-------|
| Market Average Product Price | Average price for comparable products in the market | Industry reports, market research | Use as anchor for Premium Price Index | Universal. The baseline against which premium pricing is justified. |
| Market Growth Rate (CAGR) | Growth rate of the relevant market segment | Industry reports, analyst forecasts | Calibrate per industry | Universal. A growing market is more forgiving of premium pricing. |
| Total Addressable Market (TAM) | Total market size for the product category | Market research firms | Calibrate per industry | Universal. Sets the ceiling for revenue potential. |
| Serviceable Addressable Market (SAM) | TAM filtered by company's actual reach | Internal data + market research | SAM < TAM; typically 10-40% of TAM for specific products | Universal. Realistic revenue ceiling. |
| Competitor Premium Penetration | Share of market served by competitor premium offerings | Competitive intelligence | Track and benchmark against own penetration | Universal. If competitors have low premium penetration, market may be underdeveloped or unreceptive. |
| Category Budget Share | Average customer spend on this category as % of total relevant budget | Consumer surveys, industry data | Calibrate per industry | Essential utility categories (telecom, insurance) command 3-5% of income; discretionary categories vary widely. |
| Industry Value Trend | Year-over-year change in price-value perception across the industry | Industry surveys | Positive or stable trend | A declining trend means customers perceive less value for the same price. Example: Simon-Kucher found 7% decline in one industry. [Simon-Kucher 2024](https://www.simon-kucher.com/en/insights/2024-telco-growth-strategies-brand-portfolio-and-pricing-insights) |

---

## Calibration Notes

Benchmarks above are starting points. The same KPI behaves very differently across industries. Use this table to calibrate thresholds for your specific context.

### Revenue Uplift

| Industry | Typical Range | Why |
|----------|--------------|-----|
| SaaS | 50-100%+ from premium tier | High margins, low incremental cost per user |
| Retail / Physical Goods | 10-20% | Physical COGS limits margin expansion |
| Subscriptions (media, services) | 15-40% | Content/service costs scale sub-linearly |
| Financial Services | 20-50% | Premium products (wealth management, cards) carry high margins |

### NPS Targets

| Industry | Typical NPS Range | Source |
|----------|------------------|--------|
| Technology / SaaS | 40-60 | [CustomerGauge](https://customergauge.com/benchmarks/blog/telecommunications-nps-benchmarks-and-cx-trends) |
| Retail | 50-70 | Industry benchmarks |
| Financial Services | 20-40 | Industry benchmarks |
| Telecom | 25-35 (~29-30 avg) | [CustomerGauge](https://customergauge.com/benchmarks/blog/telecommunications-nps-benchmarks-and-cx-trends) |
| Healthcare | 30-50 | Industry benchmarks |

### CAC Payback Period

| Industry | Typical Target | Why |
|----------|---------------|-----|
| SaaS (enterprise) | 12-18 months | Long sales cycles, high CLV |
| SaaS (SMB/self-serve) | 3-6 months | Short cycles, lower CLV |
| Consumer Apps | 1-3 months | High churn, must recover fast |
| Financial Services | 6-12 months | Moderate CLV, regulated acquisition |
| Retail / E-commerce | 1-3 months | Low margins, high volume |

### Penetration Rate (Premium Tier)

| Industry | Typical Range | Why |
|----------|--------------|-----|
| SaaS (freemium-to-paid) | 3-7% | Self-serve: 3-5%, sales-assisted: 5-7%. [First Page Sage](https://firstpagesage.com/seo-blog/saas-freemium-conversion-rates/) |
| Mobile Apps / Gaming | 1-5% | High volume, low conversion. Gaming: 1-2% |
| Established Subscription Products | 10-25% | Mature base with proven value proposition |
| Luxury / Premium Retail | 5-15% | Smaller segment, higher willingness to pay |

### Churn Reduction from Bundling

| Bundle Type | Typical Reduction | Source |
|-------------|------------------|--------|
| General product bundling | 5-15% (modest) | [Prince & Greenstein 2014](https://host.kelley.iu.edu/riharbau/RePEc/iuk/wpaper/bepp2011-05-prince-greenstein.pdf) |
| Well-designed multi-product bundles | 25-35% (moderate) | Industry consensus across multiple sources |
| Tightly integrated entertainment bundles | 50%+ (strong) | Disney+ bundle: 59% reduction. [Ampere Analysis](https://www.nexttv.com/news/disney-bundlers-59-less-likely-to-churn-research-company-says-chart) |

### Service/Add-on Revenue Share

| Industry | Typical Range | Why |
|----------|--------------|-----|
| SaaS | 30-60% | Professional services, integrations, premium support |
| Retail | 5-15% | Warranties, accessories, installation |
| Telecom | 15-25% | Value-added services (content, insurance, cloud) |
| Financial Services | 20-40% | Fees, advisory, premium card features |

### Affordability Threshold (Category Budget Share)

| Category Type | Typical % of Income | Why |
|---------------|-------------------|-----|
| Essential Utilities (telecom, energy) | 3-5% | Non-discretionary, regulated |
| Entertainment / Media | 2-4% | Semi-discretionary |
| Enterprise SaaS | Variable (% of IT budget) | Budget-driven, not income-driven |
| Luxury / Premium Discretionary | No fixed ceiling | Status-driven, less price-sensitive |

### Freemium Conversion Rates

| Model | Typical Range | Source |
|-------|--------------|--------|
| Self-serve freemium | 3-5% | [First Page Sage 2026](https://firstpagesage.com/seo-blog/saas-freemium-conversion-rates/) |
| Sales-assisted freemium | 5-7% | [First Page Sage 2026](https://firstpagesage.com/seo-blog/saas-freemium-conversion-rates/) |
| Top performers | 8-15% | [First Page Sage 2026](https://firstpagesage.com/seo-blog/saas-freemium-conversion-rates/) |
| Mobile gaming | 1-2% | Industry benchmarks |

---

## GBB (Good-Better-Best) Tier Distribution Reference

When using three-tier pricing, research consistently shows the majority of customers choose the middle tier.

| Tier | Approximate Share | Source |
|------|------------------|--------|
| Good (low) | 10-25% | [HBR: Good-Better-Best Approach to Pricing](https://hbr.org/2018/09/the-good-better-best-approach-to-pricing) |
| Better (middle) | ~66% (60-70%) | Multiple sources converge on this range |
| Best (high) | 10-20% | Varies by price gap and perceived premium value |

Removing the middle tier can reduce revenue by approximately 50% (industry sources, partially verified).

---

## Source Index

All fact-checked sources referenced in this catalog:

| Label | URL |
|-------|-----|
| McKinsey Personalization | https://www.mckinsey.com/capabilities/growth-marketing-and-sales/our-insights/the-value-of-getting-personalization-right-or-wrong-is-multiplying |
| HBR: Good-Better-Best Pricing | https://hbr.org/2018/09/the-good-better-best-approach-to-pricing |
| HBR: Value of Keeping Customers | https://hbr.org/2014/10/the-value-of-keeping-the-right-customers |
| CustomerGauge NPS Benchmarks | https://customergauge.com/benchmarks/blog/telecommunications-nps-benchmarks-and-cx-trends |
| Simon-Kucher 2024 Telco Study | https://www.simon-kucher.com/en/insights/2024-telco-growth-strategies-brand-portfolio-and-pricing-insights |
| Simon-Kucher: Leaders/Fillers/Killers | https://www.simon-kucher.com/en/insights/future-proof-telco-marketing-6-strategies-best-practices |
| First Page Sage: Freemium Conversion | https://firstpagesage.com/seo-blog/saas-freemium-conversion-rates/ |
| Ampere Analysis: Disney+ Churn | https://www.nexttv.com/news/disney-bundlers-59-less-likely-to-churn-research-company-says-chart |
| Prince & Greenstein 2014 | https://host.kelley.iu.edu/riharbau/RePEc/iuk/wpaper/bepp2011-05-prince-greenstein.pdf |
| Shaddy & Fishbach 2017 (Dilution) | https://www.anderson.ucla.edu/documents/areas/fac/marketing/Seminars/Fall%202017/Shaddy%20%20Fishbach%20-%20How%20Bundling%20Affects%20Valuation%20(job%20market%20paper).pdf |
| Deloitte TMT Predictions 2025 | https://www.deloitte.com/us/en/insights/industry/technology/technology-media-and-telecom-predictions/2025/tmt-predictions-video-streaming-bundles-bigger-than-ever.html |
| Bain: NPS in Telecom | https://www.bain.com/client-results/dialing-up-customer-experience-in-telecommunications/ |
| Van Westendorp PSM | https://sawtoothsoftware.com/resources/blog/posts/van-westendorp-pricing-sensitivity-meter |
| Gabor-Granger Method | https://sawtoothsoftware.com/resources/blog/posts/gabor-granger-pricing-method |
| BCG Pathways Conjoint | https://www.bcg.com/publications/2014/telecommunications-pricing-pathways-conjoint-new-approach-pricing-mobile |
