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

| Module | Functions | Description |
|--------|-----------|-------------|
| pricing | bvr, pvr, wtp, price_position, elasticity, tier_analysis | Price-value, willingness to pay, market positioning |
| bundle | dead_weight, attach_rate, bundle_discount, component_contribution, bundle_efficiency, optimal_bundle | Bundle composition and value analysis |
| financial | clv, unit_economics, payback, margin_structure, revenue_mix, break_even | Unit economics and financial viability |
| customer | adoption_curve, churn_risk, nps_impact, segment_fit, switching_cost, retention_value | Customer behavior and segmentation |
| product | feature_gap, quality_score, lifecycle_stage, innovation_index, platform_dependency, ecosystem_lock | Product maturity and competitive position |
| scoring | go_no_go, risk_matrix, scenario_model, sensitivity, weighted_score, composite_index | Decision scoring and risk assessment |

### Test

```bash
make -C tools/appraise test
```
