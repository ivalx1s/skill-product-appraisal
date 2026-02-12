# TASK-260212-m2z8ik: factcheck-generalized-frameworks

## Description
After generalization, fact-check ALL numbers, metrics, benchmarks, and citations taken from internet sources. This includes:

(a) Academic/theoretical frameworks: Adams-Yellen/Schmalensee bundle theory, Simon-Kucher Leaders/Fillers/Killers (40% dead weight benchmark), Tversky/Kahneman anchoring research, Wansink bundling experiments (32% sales boost)

(b) Industry benchmarks: McKinsey bundle revenue uplift (10-30%), HBR mixed bundling outperformance (25-35%), Simon-Kucher 2024 survey (3/5 respondents, 13% conversion), churn reduction benchmarks (25-35%), NPS targets (15-25 avg, 30+ premium)

(c) Pricing thresholds: Bundle Value Ratio >1.5x, affordability <3-5% disposable income, GBB distribution 20/66/14, tier gap recommendations, 15% satisfaction penalty for bundles

(d) Financial metrics: ARPU uplift 12-18%, CAC recovery <6 months, cannibalization <50%, partner cost ratio <30%, CLV 2-4x base

(e) Any other specific numbers presented as facts/benchmarks in the methodology docs

Method: web search each claim against original sources. Flag hallucinated, misattributed, or outdated data.

SOURCE ATTRIBUTION RULE: Every EXTERNAL number (from third-party research, publications, reports) must have a direct source URL. No 'McKinsey says X' without a link to the actual McKinsey publication. If a source cannot be found — flag for removal or reword as 'industry estimate' without specific attribution.

Exception: numbers that WE calculated or derived (e.g. value ratios, price gaps, projected revenues) don't need external sources — they are our analysis based on input data. Mark these clearly as 'calculated' vs 'sourced'.

## Scope
(define task scope)

## Acceptance Criteria
1. Every external number has a source URL or is flagged for removal. 2. Each claim categorized as: verified (with URL), unverifiable (flagged), or calculated (our analysis). 3. Results stored in .research/factcheck-results.md with clear pass/fail/flag per item. 4. No hallucinated citations remain.
