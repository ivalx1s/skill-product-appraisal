# Worked Example: Telecom Premium Bundle Appraisal

> **Product:** Operator A -- "Connect" Premium Tariff Lineup (3 tiers)
> **Industry:** Telecom (mobile + ecosystem)
> **Market:** Fictional national operator, ~70M subscribers, mid-tier brand positioning
> **Date:** 2026-01-15 (fictional)
> **Verdict:** Conditional Go (weighted score 3.30 / 5.0)

This is a condensed example showing how the 7-dimension methodology works on a real-world-style telecom evaluation. All operator names, prices, and bundle compositions are fictional or modified. Competitors referenced (MegaFon, T2, Beeline, Yota) use publicly available data for illustrative purposes only.

---

## (a) Product Overview

Operator A plans to launch a 3-tier premium lineup under the "Connect" brand. The lineup targets the 1,500-5,000 RUB/month price band -- a segment with limited competition in the Russian market.

### Tier Structure

| Tier | Name | Price (RUB/mo) | Model | Target Segment |
|------|------|---------------|-------|----------------|
| Entry | Connect Go | 1,800 | Subscription | Digital families, ecosystem consolidation |
| Middle | Connect Plus | 2,800 | Tariff | Business travelers, mobile professionals |
| Top | Connect Max | 4,200 | Tariff | High-income lifestyle, status-seekers |

### Bundle Composition

| Component | Connect Go | Connect Plus | Connect Max |
|-----------|-----------|-------------|-------------|
| Voice minutes | 1,500 | 2,500 | 5,000 |
| Mobile data | Unlimited | Unlimited | Unlimited + priority QoS |
| SMS | 300 | 500 | Unlimited |
| Home broadband (300 Mbps) | Yes | No | Yes (500 Mbps) |
| Video Platform subscription | Yes | Yes | Yes + 200 TV channels |
| Music Service | Yes | Yes | Yes |
| Digital Library | Yes | Yes | Yes |
| Sports Club Chain access | Yes | No | Yes |
| Restaurant Network benefit | No | No | Yes (monthly dining credit) |
| International roaming | No | 150 days/year | 150 days/year |
| Personal manager | No | No | Yes |
| AI voice assistant | Yes | Yes | Yes |
| Cloud storage (256 GB) | Yes | Yes | Yes (512 GB) |
| Security suite | Yes | Yes | Yes |
| Bank integration | Basic | Basic | Premium tier |

**Standalone value sums** (from retail pricing of each component individually):

| Tier | Standalone Sum | Bundle Price | BVR |
|------|---------------|-------------|-----|
| Connect Go | 7,900 | 1,800 | 4.39x |
| Connect Plus | 7,400 | 2,800 | 2.64x |
| Connect Max | 18,700 | 4,200 | 4.45x |

**Usage-adjusted value** (conservative -- not all subscribers use every service):

| Tier | Usage-Adjusted Value | Bundle Price | Adjusted BVR |
|------|---------------------|-------------|--------------|
| Connect Go | 4,200 (53% utilization) | 1,800 | 2.33x |
| Connect Plus | 4,500 (61% utilization) | 2,800 | 1.61x |
| Connect Max | 6,800 (36% utilization) | 4,200 | 1.62x |

Note: even under conservative usage assumptions, all tiers exceed the 1.5x minimum BVR threshold (PRC-1). Connect Plus has the thinnest margin at 1.61x because roaming (its biggest value component at 4,800 RUB/year annualized) is used by only ~18% of subscribers on average.

### Competitive Context

The Russian premium telecom segment is sparse above 1,500 RUB/month. The competitive landscape at launch:

| Operator | Top Product | Price | Key Strengths | Key Gaps |
|----------|-----------|-------|---------------|----------|
| MegaFon | Premium | 3,000 RUB | 6,000 minutes, free broadband, personal manager, 3 family members | No fitness, no dining, weak roaming (3 GB/mo) |
| T2 | Premium | 1,300-1,800 RUB | Free unlimited roaming internet, price freeze, data rollover | No ecosystem depth, per-minute roaming voice |
| Beeline | UP5 (max) | 1,110 RUB | Grocery cashback via X5 Retail | No premium tier, most expensive roaming, no ecosystem |
| Yota | Top plan | ~1,200 RUB | Simple pricing, speed-based tiers | Capped at mid-market, no bundled services |

Operator A would create a new competitive category: the "premium ecosystem tariff" with lifestyle components. No existing Russian operator bundles fitness, dining, and entertainment into a mobile plan above 3,000 RUB.

---

## (b) Dimension-by-Dimension Walkthrough

### Dimension 1: Product-Market Fit -- Score: 4/5, Gate: PASS

The addressable segment is approximately 8-10% of Operator A's 70M subscriber base (PMF-1: pass at >5%). Premium engagement is growing: the operator's existing loyalty tier (a 400 RUB/month subscription) has 2.8M active subscribers, demonstrating demand for value-added services above the base tariff (PMF-2: pass). Willingness-to-pay analysis using Van Westendorp PSM on a 3,000-person survey placed the Optimal Price Point at 2,100 RUB for the entry tier and 3,800 RUB for the top tier -- both above the proposed prices (PMF-3: pass). Service-need alignment is strong: 68% of high-ARPU subscribers expressed interest in 4+ of the bundled components (PMF-4: pass). Conversion funnel modeling projects 4.5% of eligible subscribers at entry tier and 1.2% at the top tier, within the 3-10% / 1-5% benchmarks (PMF-6: pass). Brand permission is the weakest point -- Operator A is perceived as "reliable and innovative" but not yet as a "premium lifestyle brand" (PMF-7: 3/5).

**Key metric:** 8-10% addressable segment = 5.6-7.0M potential customers.

### Dimension 2: Pricing Adequacy -- Score: 3/5, Gate: CONDITIONAL

BVR exceeds 2.0x at all tiers (PRC-1: strong pass). However, the tier gap architecture has a structural problem: Connect Go at 1,800 to Connect Plus at 2,800 is a 56% jump, and Connect Plus at 2,800 to Connect Max at 4,200 is a 50% jump. Both are at the threshold of the >50% red flag (PRC-3: borderline). The GBB anchoring analysis predicts approximately 52% choosing the entry tier, 38% the middle, and 10% the top -- deviation from the ideal ~66% middle tier concentration indicates the entry tier is too attractive relative to the middle (PRC-4: below target). Affordability checks pass in the top 8 metropolitan areas but fail in 40+ smaller regions where 4,200 RUB exceeds 5% of average monthly income (PRC-6: conditional pass, primary markets only). Price floor clearance is positive at all tiers with estimated unit costs of 750 RUB (Go), 920 RUB (Plus), and 1,350 RUB (Max) per month (PRC-7: pass).

**Key metric:** Entry tier attracts 52% of subscribers vs. target 20-25%, indicating pricing architecture imbalance.

### Dimension 3: Bundle Composition -- Score: 3/5, Gate: CONDITIONAL

**Leaders:** Video Platform and home broadband are the clear purchase drivers (BND-1: pass). Sports Club Chain access is a secondary leader for Connect Go and Connect Max. **Fillers:** Music Service, Digital Library, AI assistant, cloud storage, and security suite add perceived value at low marginal cost. **Potential Killers:** Restaurant Network benefit (Connect Max only) is projected to be used by only 12% of subscribers and is geographically constrained to 6 cities (BND-3: dilution risk present). Dead weight ratio is estimated at 28% -- below the 40% red flag but worth monitoring (BND-2: pass). The critical issue is access constraints: Sports Club Chain operates in 14 cities covering roughly 35% of the subscriber base, and Restaurant Network covers only 6 cities. Combined, approximately 32% of the top-tier standalone value is tied to geographically constrained components (BND-5: borderline at 30% threshold).

**Key metric:** 32% of Connect Max value is access-constrained -- just above the 30% red flag.

### Dimension 4: Competitive Positioning -- Score: 4/5, Gate: PASS

Operator A's lineup enters a market segment with limited competition. MegaFon Premium at 3,000 RUB is the most direct competitor, offering more voice minutes (6,000) and free home internet but no fitness or dining components. T2 Premium at 1,300-1,800 RUB undercuts on price and includes free unlimited roaming internet, but lacks any ecosystem depth. Beeline has no premium offering above 1,100 RUB. The Connect lineup has 4 features not replicable within 6 months: the Sports Club Chain partnership (exclusive contract), the Video Platform integration (proprietary), the banking tier (in-house), and the convergent home broadband bundling (own infrastructure) (CMP-1, CMP-5: pass). Time to full competitive imitation is estimated at 14-18 months due to partner exclusivity and infrastructure requirements (CMP-3: pass). Competitive BVR comparison: Connect Go at 4.39x exceeds MegaFon Premium at approximately 1.9x and T2 Premium at approximately 2.5x (CMP-2: pass).

**Key metric:** 4 defensible components; 14-18 months estimated imitation time.

### Dimension 5: Financial Viability -- Score: 3/5, Gate: CONDITIONAL

Unit economics are positive at all tiers: gross margin per subscriber estimated at 1,050 RUB/month (Go), 1,880 RUB/month (Plus), and 2,850 RUB/month (Max) (FIN-2: pass). Cannibalization is a concern: an estimated 60% of Connect Go adopters will migrate from existing mid-tier plans (800-1,100 RUB), yielding a net revenue uplift of only 700-1,000 RUB per migrator vs. the full 1,800 RUB for net-new subscribers. However, net incremental revenue remains positive (FIN-3: pass). Partner and licensing costs for the Sports Club Chain and Video Platform consume approximately 34% of the premium revenue delta -- slightly above the 30% red-flag threshold (FIN-4: borderline). Stress test (costs +20%, growth -30%): Connect Go and Connect Plus remain margin-positive; Connect Max breaks even under stress due to high partner costs for Sports Club and Restaurant Network (FIN-7: conditional pass). Year 1 subscriber projection: 180K total (100K Go, 60K Plus, 20K Max), generating approximately 6.5B RUB annual revenue.

**Key metric:** Partner cost ratio at 34% exceeds the 30% target; stress test is tight for the top tier.

Cannibalization paths and net revenue impact:

| Migration Path | Volume (Y1) | Net Revenue Impact per User |
|---------------|-------------|---------------------------|
| Mid-tier (800-1,100) to Connect Go (1,800) | ~60K | +700-1,000 RUB/mo |
| Existing loyalty sub (400) to Connect Go (1,800) | ~25K | +1,400 RUB/mo |
| Net new (from competitors / unsubscribed) to any tier | ~55K | Full incremental |
| Standalone Video Platform (280) to Connect Go (1,800) | ~15K | +1,520 but -280 lost standalone |
| Premium mid-tier (1,500+) to Connect Plus (2,800) | ~25K | +1,300 RUB/mo |

Total net incremental revenue remains positive. The primary financial lever here is not ARPU uplift alone -- it is churn reduction. Premium bundle subscribers in similar markets show 40-55% lower churn, which translates to 1.8-2.2x CLV improvement even at modest revenue premiums (FIN-5: expected pass).

### Dimension 6: Customer Experience -- Score: 4/5, Gate: PASS

Premium bundle subscribers in comparable markets show 40-55% lower churn than base subscribers (CX-1: expected pass, pending launch data). Feature utilization projections from the operator's existing loyalty program suggest 65% of subscribers will actively use 4+ components within 3 months (CX-4: expected pass). The WTP validation shows proposed prices are at 72-78% of measured WTP across tiers -- comfortably below the 80% ceiling (CX-5: pass). The primary disappointment risk centers on the Sports Club Chain component: users in cities without locations may perceive a "phantom benefit" that erodes satisfaction. Mitigation through regional component swapping is recommended (CX-8: conditional, pending regional adaptation).

**Key metric:** Price-to-WTP ratio at 72-78% provides headroom for satisfaction.

Component-level engagement projections based on the operator's existing loyalty program data:

| Component | Projected Monthly Active Rate | Standalone Benchmark | Gap |
|-----------|------------------------------|---------------------|-----|
| Video Platform | 72% | 68% (industry avg) | +4pp -- Leader effect |
| Music Service | 48% | 55% (standalone apps) | -7pp -- passive inclusion |
| Home broadband | 85% | 90% (always-on) | -5pp -- some users already have ISP |
| Sports Club Chain | 22% | 35% (gym membership avg) | -13pp -- access constraint |
| Digital Library | 18% | 20% (e-book subs) | -2pp -- niche interest |
| Security suite | 88% | 85% (passive service) | +3pp -- bundled activation |

The Sports Club Chain's 22% engagement rate is the clear outlier. In cities with locations, the rate rises to 38% -- strong. In cities without locations, it drops to 0%. This bimodal distribution is the core argument for regional component swapping.

### Dimension 7: Market Reach -- Score: 2/5, Gate: CONDITIONAL

This is the weakest dimension. Connect Go (with home broadband and Sports Club Chain) is fully deliverable in 14 cities covering approximately 40% of the subscriber base. Connect Plus (roaming-centric, no constrained physical components) works nationally. Connect Max (Sports Club + Restaurant Network + broadband) is only fully deliverable in 6 cities covering approximately 22% of the subscriber base (MR-1: fail for Max tier). Affordability drops sharply outside the top 10 regions: Connect Max at 4,200 RUB exceeds 7% of average monthly income in 50+ regions (MR-2: fail outside primary markets). Demand heterogeneity testing shows viable segments in primary (Moscow, St. Petersburg) and secondary (top-10 cities) markets, but tertiary markets lack both affordability and access coverage (MR-6: partial pass).

**Key metric:** Top tier is fully deliverable to only 22% of the subscriber base.

---

## (c) Condensed Scorecard

| Dimension | Weight | Score (1-5) | Weighted | Go/No-Go | Gate | Rationale |
|-----------|--------|-------------|----------|----------|------|-----------|
| Product-Market Fit | 15% | 4 | 0.60 | Go | PASS | Strong demand signal; brand permission is the only soft spot |
| Pricing Adequacy | 15% | 3 | 0.45 | Conditional | COND | BVR is excellent but tier gaps drive wrong distribution (52% entry vs. 20-25% target) |
| Bundle Composition | 10% | 3 | 0.30 | Conditional | COND | Leaders are clear; access constraints at 32% are borderline |
| Competitive Position | 15% | 4 | 0.60 | Go | PASS | 4 defensible features; 14-18 months to imitation; best BVR in market |
| Financial Viability | 25% | 3 | 0.75 | Conditional | COND | Positive unit economics; partner cost ratio (34%) and stress test are tight |
| Customer Experience | 10% | 4 | 0.40 | Go | PASS | Strong WTP headroom; utilization projections favorable |
| Market Reach | 10% | 2 | 0.20 | Conditional | COND | Top tier deliverable to 22% of base; affordability fails outside top 10 |
| **Total** | **100%** | | **3.30** | **Conditional Go** | | |

**Overall Verdict: Conditional Go** -- address pricing architecture, access constraints, and partner costs before national launch.

### Gate Failures Requiring Action

| Dimension | Gate | Status | Required Action |
|-----------|------|--------|----------------|
| Pricing Adequacy | PRC-3: tier gaps >50% | TRIGGERED | Reduce Connect Max price to close the gap |
| Pricing Adequacy | PRC-4: GBB anchoring off-target | TRIGGERED | Rebalance feature differentiation between Go and Plus |
| Bundle Composition | BND-5: access constraints >30% | TRIGGERED | Regional component swapping for Sports Club + Restaurant |
| Financial Viability | FIN-4: partner costs >30% | TRIGGERED | Renegotiate Sports Club Chain terms or restructure cost model |
| Market Reach | MR-1: access coverage <70% for Max | TRIGGERED | Adapt top-tier components for non-primary markets |
| Market Reach | MR-2: affordability fails outside top 10 | TRIGGERED | Price reduction or regional pricing tiers |

### Top 3 Risks

1. **Geographic concentration of lifestyle components** -- P:4 x I:4 = 16 -- Mitigation: regional swapping program, digital alternatives for non-covered cities.
2. **Cannibalization of existing mid-tier plans** -- P:4 x I:3 = 12 -- Mitigation: sunset overlapping plans, position Connect as the clear upgrade path.
3. **Partner cost escalation (Sports Club Chain)** -- P:3 x I:4 = 12 -- Mitigation: volume-based contracts, per-usage pricing model, cap annual increases at CPI+3%.

---

## (d) Key Recommendations

1. **Fix the tier gap architecture.** Reduce Connect Max from 4,200 to 3,600 RUB (28% gap from Plus instead of 50%). This improves GBB dynamics, shifts more subscribers toward the middle tier, and expands affordability to 12+ regions. The BVR at 3,600 remains a strong 5.19x.

2. **Implement regional component swapping for Connect Max.** In cities without Sports Club Chain or Restaurant Network, offer equivalent alternatives: additional digital benefits (extended cloud, family lines, enhanced cashback) or partnerships with regional fitness/dining providers. This raises access coverage from 22% to an estimated 60%+.

3. **Renegotiate partner economics or restructure cost sharing.** The Sports Club Chain partnership at current terms pushes partner costs to 34% of the premium delta. Target: bring total partner costs below 28% either through volume-based rate reductions (commit to 50K+ active users) or by shifting to a per-usage model instead of flat per-subscriber fees.

4. **Lead acquisition with Connect Go (entry tier).** It has the strongest value ratio (4.39x), the lowest price barrier, and the subscription model positions it against ecosystem competitors (Sber, Yandex) rather than pure telecom. Target: 55% of total premium subscribers.

5. **Position Connect Plus as the "traveler's choice."** Roaming is its sole differentiator over Connect Go. All marketing for this tier should target business travelers, digital nomads, and internationally-mobile professionals. Consider adding a roaming add-on for Connect Go to capture occasional travelers who do not need 150 full days.

---

## (e) Lessons Learned

**What this case teaches about applying the methodology:**

**Market Reach is often the hidden killer.** This evaluation scored well on demand, competition, and customer experience -- but nearly failed on Market Reach (Dimension 7). Physical components like fitness club access and restaurant networks create powerful value propositions in primary markets but become dead weight in secondary ones. The methodology's sequential gate structure caught this: without Dimension 7, the product would have launched nationally with components that 78% of top-tier subscribers cannot use.

**Tier gap architecture is not just math -- it is psychology.** The 50% gap between Plus and Max violated GBB principles and skewed projected distribution heavily toward the entry tier (52% vs. target 20-25%). The BVR looked excellent at every tier, but the pricing architecture undermined the anchoring effect that makes Good-Better-Best work. Lesson: BVR alone is insufficient -- you must evaluate tier gaps (PRC-3) and anchoring effectiveness (PRC-4) as a system.

**Access constraints compound across tiers.** The Sports Club Chain alone was borderline (14 cities, 35% coverage). Adding Restaurant Network to the top tier pushed combined access constraints to 32%, just above the 30% red flag. Each new physical component narrows the addressable market. The methodology's BND-5 criterion exists precisely to catch this accumulation effect before launch.

**Partner costs are the most volatile input in the financial model.** The stress test (FIN-7) revealed that the top tier breaks even under pessimistic conditions -- driven entirely by high partner costs for lifestyle components. In a pure-digital bundle, costs scale predictably. Physical partnerships introduce cost volatility that must be modeled explicitly.

**The entry tier is the product.** Despite the top tier getting the most attention in product design, the entry tier (Connect Go) scored highest on value ratio, affordability, and market reach. It is the primary acquisition vehicle, the highest-volume tier, and the gateway to the premium ecosystem. Designing the entry tier to be independently compelling -- not just a stripped-down version of the top tier -- is critical.

---

*This example uses the Product Appraisal methodology (7 dimensions, 48 criteria). All IDs (PMF-1 through MR-7) reference assessment-criteria.md. Prices, features, and projections are fictional. Competitor data (MegaFon, T2, Beeline) uses publicly available information for illustrative context.*
