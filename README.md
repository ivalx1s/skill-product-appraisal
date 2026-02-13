# product-appraisal

AI agent skill for evaluating complex products, bundles, and subscription offerings. MapReduce orchestration pattern with 9 evaluation phases across 7 execution waves.

## Structure

```
product-appraisal/          # Skill (SKILL.md + references)
tools/appraise/             # CLI calculation engine (Go)
```

## Skill

Install as a global skill:

```bash
# Clone
git clone git@github.com:ivalx1s/skill-product-appraisal.git ~/agents/skills/product-appraisal-repo

# Symlink skill directory
ln -s ~/agents/skills/product-appraisal-repo/product-appraisal ~/.claude/skills/product-appraisal
ln -s ~/agents/skills/product-appraisal-repo/product-appraisal ~/.codex/skills/product-appraisal
```

## CLI Tool (`appraise`)

38 calculator functions across 6 modules: pricing, bundle, financial, customer, product, scoring.

### Install

```bash
# Via go install (requires Go 1.25+)
go install github.com/ivalx1s/skill-product-appraisal/tools/appraise@v0.1.0

# Or build from source (installs to ~/.local/bin/)
make -C tools/appraise install
```

### Usage

```bash
# Run a calculation
appraise calc pricing bvr --input data.json

# Query available modules/functions
appraise q "list(modules)"
appraise q "list(functions, module=pricing)"

# Schema introspection
appraise q "schema(pricing, bvr)"

# Scoped grep
appraise grep "dead_weight" --module bundle
```

### Modules

| Module | # | Functions | Description |
|--------|---|-----------|-------------|
| pricing | 6 | bvr, tier_gap, cost_floor, price_value_ratio, premium_price_index, bundle_discount | Price-value ratios, tier analysis, cost floors, premium indexing |
| bundle | 5 | classify, dead_weight, cross_subsidy, component_activation, multi_component_usage | Component classification, dead weight, cross-subsidy analysis |
| financial | 9 | unit_economics, gross_margin, clv, cac_payback, break_even, cannibalization, stress_test, incremental_revenue, revenue_uplift | Unit economics, margins, CLV, payback, stress testing |
| customer | 7 | churn_rate, retention_rate, nps, csat, churn_reduction, revenue_growth, service_revenue_share | Churn, retention, NPS, CSAT, revenue growth |
| product | 8 | penetration_rate, migration_rate, cannibalization_rate, cross_sell_rate, feature_utilization, component_activation_rate, attach_rate, trial_conversion | Adoption rates, cross-sell, feature utilization, conversions |
| scoring | 3 | go_no_go, risk_matrix, dimension_score | Go/no-go decision, risk matrix, dimension scoring |

### Test

```bash
make -C tools/appraise test
```
