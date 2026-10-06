# Step 4: Master AI Implementation Prompt

Act as a Principal Software Architect and Lead Developer. 

I need you to help me plan and implement **Step 4: Post Retrieval Engine** for my Instagram CLI project (`go-insta-cli`), building directly on top of the work completed in **Step 3** without duplicating existing code or breaking the established architecture.

---

### STEP 4 SPECIFIC REQUIREMENTS
1. Implement Post Retrieval Engine: Create the core logic to fetch user's recent Instagram posts (Media ID, Caption, Timestamp, Like Count, Comment Count, and Media Type).
2. Enhance `posts` subcommand: When user executes `insta posts`, fetch posts and render a beautifully formatted, colorized terminal table/list using ANSI colors.
3. Support optional CLI flags for `posts`:
   - `--limit <n>`: Limit number of posts returned (default: 10).
   - `--json`: Output raw JSON formatted string for scripting.
4. Integrate `RequireAuth()` authentication guard from Step 3 to ensure unauthenticated users cannot run `posts`.
5. Implement graceful error handling: Handle API rate limits, network timeouts, invalid session errors, and empty post states with clear colored terminal messages.
6. Maintain mock/test data provider fallback so the feature can be safely tested offline without triggering Instagram rate limits.

---

### STRICT OPERATING RULES
1. DO NOT WRITE, MODIFY, OR GENERATE ANY PROJECT CODE IMMEDIATELY.
2. Perform a complete Audit and Gap Analysis of Step 3 first.
3. Prevent code duplication — reuse and extend existing files (`main.go`, `config.json` logic, `RequireAuth()`), functions, and structs wherever possible.
4. Do not make assumptions. If any file, context, or specification is missing, explicitly ask for it.
5. Follow the 8-Phase Planning Process detailed below.
6. CRITICAL STOPPING RULE: After presenting the complete Phase 1 - Phase 8 analysis and plan, STOP and wait for my explicit approval. Do NOT generate implementation code until I approve the plan.

---

### REQUIRED 8-PHASE PLANNING WORKFLOW

#### Phase 1 — Understand
- Review Step 4 requirements (Post Retrieval Engine, `posts` subcommand formatting, `--limit` and `--json` flags, `RequireAuth()` integration, error handling).

#### Phase 2 — Audit Existing Work (Step 3)
- Inspect completed Step 3 files (`main.go`, `config.json` storage engine), authentication guard (`RequireAuth()`), ANSI color constants, and existing subcommands (`login`, `logout`, `status`).
- Document all reusable helper functions and data structures.

#### Phase 3 — Gap Analysis
- Clearly categorize:
  1. What is 100% completed in Step 3
  2. What is partially completed
  3. What is completely missing for Step 4
  4. What needs verification or testing

#### Phase 4 — Task Breakdown
- Divide Step 4 into small, practical, dependency-ordered tasks.

#### Phase 5 — Requirement Mapping
- Map every Step 4 requirement to its corresponding task(s) from Phase 4.

#### Phase 6 — Reuse & Duplication Check
- Audit proposed tasks against existing Step 3 code to prevent duplication.

#### Phase 7 — Validation & Test Plan
- Define an end-to-end testing strategy for Step 4.

#### Phase 8 — Final Implementation Plan & Execution Order
- Present the structured implementation order.

---

### CRITICAL STOPPING INSTRUCTION
Once you have generated the complete 8-Phase Analysis and Implementation Plan:
STOP IMMEDIATELY AND WAIT FOR MY APPROVAL.
