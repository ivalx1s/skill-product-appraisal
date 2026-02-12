# Plan: Product Appraisal Skill Creation

**Date:** 2026-02-12 23:35
**Epic:** EPIC-260212-m2waop — Create Product Appraisal Skill
**Repo:** skill-product-appraisal (git@github.com:ivalx1s/skill-product-appraisal.git)

---

## Phase Overview

### Phase 1: Research Chain (sequential, 1 agent)
**Agent:** agent-research (opus, foreground)
**Story:** STORY-260212-1wl7c7 (extract-generalize-methodology)

| Wave | Task | ID |
|------|------|----|
| 1 | audit-source-methodology | TASK-260212-wadul3 |
| 2a | generalize-assessment-dimensions | TASK-260212-1detcq |
| 2b | generalize-kpis | TASK-260212-3s5bk1 |
| 3 | factcheck-generalized-frameworks | TASK-260212-m2z8ik |
| 4 | write-improvements-summary | TASK-260212-32zwnl |
| 5 | write-methodology-description | TASK-260212-2m6de7 |

Why 1 agent: needs continuous context (audit informs generalization, generalization informs factcheck, factcheck informs improvements, improvements inform methodology). Splitting would lose context.

### Phase 2: Write Skill Files (parallel, 6 agents)
**Agents:** agent-skill-md, agent-assessment, agent-bundle, agent-pricing, agent-kpis, agent-template
**Story:** STORY-260212-17hlts (write-skill-and-references)

All 6 tasks can run in parallel after Phase 1 completes:

| Agent | Task | ID | Output |
|-------|------|----|--------|
| agent-skill-md | write-skill-md | TASK-260212-rbvs9s | product-appraisal/SKILL.md |
| agent-assessment | write-assessment-criteria-ref | TASK-260212-ty55pf | product-appraisal/references/assessment-criteria.md |
| agent-bundle | write-bundle-valuation-ref | TASK-260212-2l8au7 | product-appraisal/references/bundle-valuation.md |
| agent-pricing | write-pricing-methods-ref | TASK-260212-3gruhi | product-appraisal/references/pricing-methods.md |
| agent-kpis | write-kpi-catalog-ref | TASK-260212-239sj5 | product-appraisal/references/kpi-catalog.md |
| agent-template | write-evaluation-template | TASK-260212-376cy8 | product-appraisal/references/evaluation-template.md |

### Phase 2.5: Example (sequential, after template)
**Agent:** agent-example (opus, foreground)
**Task:** write-example-telecom-case (TASK-260212-29tbxu)
**Output:** product-appraisal/references/example-telecom-appraisal.md
Blocked by template (needs scorecard format).

### Phase 3: Package & Install (coordinator, no agent needed)
**Story:** STORY-260212-3hfcog (package-and-install)

| Task | ID | Who |
|------|----|-----|
| validate-and-package | TASK-260212-10vdw7 | coordinator |
| install-global-symlinks | TASK-260212-6oaiuc | coordinator |
| update-skill-triggers | TASK-260212-2whya6 | coordinator |

---

## Critical Path

```
audit → generalize (2 parallel) → factcheck → improvements → methodology
  → write files (6 parallel) → example → validate → symlinks → triggers
```

## Agent Allocation

- Phase 1: 1 opus agent (sequential research chain)
- Phase 2: up to 6 opus agents (parallel file writing)
- Phase 2.5: 1 opus agent (example case)
- Phase 3: coordinator (me)

## Source Data

All methodology docs in: `/Users/aagrigore1/src/mts-tariff-research/.research/`
- 260212_methodology-framework.md (639 lines)
- 260212_bundle-valuation-methods.md (565 lines)
- 260212_telecom-assessment-frameworks.md (760 lines)
- 260212_telecom-pricing-methodologies.md (700 lines)

Final report for example: `/Users/aagrigore1/src/mts-tariff-research/REPORT.md`

## Output

Skill directory: `/Users/aagrigore1/src/skill-product-appraisal/product-appraisal/`
```
product-appraisal/
├── SKILL.md
├── references/
│   ├── assessment-criteria.md
│   ├── bundle-valuation.md
│   ├── pricing-methods.md
│   ├── kpi-catalog.md
│   ├── evaluation-template.md
│   └── example-telecom-appraisal.md
├── scripts/
└── assets/
```

Research artifacts: `/Users/aagrigore1/src/skill-product-appraisal/.research/`
METHODOLOGY.md: project root
