# Step 3: Master AI Implementation Prompt

Act as a Principal Software Architect and Lead Developer. 

I need you to help me plan and implement **Step 3: Instagram Authentication Module** for my Instagram CLI project (`go-insta-cli`), building directly on top of the work completed in **Step 2** without duplicating existing code or breaking the established architecture.

---

### STEP 3 SPECIFIC REQUIREMENTS
1. Implement Instagram Authentication Engine: Enhance the `login` subcommand with an interactive terminal prompt for Instagram Username and Password / Session Cookie (`sessionid`).
2. Validate authentication credentials and active session token format.
3. Update `config.json` securely with validated session credentials (`IsLoggedIn = true`, `Username`, encrypted `SessionToken`, and `LastLogin` timestamp).
4. Implement a new `logout` subcommand: Clear session credentials from `config.json` (`IsLoggedIn = false`, wipe `SessionToken`), and display a colored notification.
5. Implement an authentication guard helper function `RequireAuth(cfg Config) bool`:
   - Checks if `cfg.IsLoggedIn` is true and `SessionToken` is present.
   - If false, prints a red error: "Error: You must be logged in. Run 'insta login' first." and blocks command execution.
6. Integrate `RequireAuth()` into protected subcommands (`posts`, `delete`, `show ui`) so unauthenticated users cannot access them.

---

### STRICT OPERATING RULES
1. DO NOT WRITE, MODIFY, OR GENERATE ANY PROJECT CODE IMMEDIATELY.
2. Perform a complete Audit and Gap Analysis of Step 2 first.
3. Prevent code duplication — reuse and extend existing files (`main.go`, `config.json` logic), functions, and structs wherever possible.
4. Do not make assumptions. If any file, context, or specification is missing, explicitly ask for it.
5. Follow the 8-Phase Planning Process detailed below.
6. CRITICAL STOPPING RULE: After presenting the complete Phase 1 - Phase 8 analysis and plan, STOP and wait for my explicit approval. Do NOT generate implementation code until I approve the plan.

---

### REQUIRED 8-PHASE PLANNING WORKFLOW

#### Phase 1 — Understand
- Review Step 3 requirements (Interactive Login, Session Validation, `logout` subcommand, `RequireAuth()` guard, and protecting subcommands).

#### Phase 2 — Audit Existing Work (Step 2)
- Inspect completed Step 2 files (`main.go`, `pkg/config/config.go`, `pkg/cli/colors.go`), existing structs (`Config`), CLI flags (`--version`, `--help`), ANSI color constants, and subcommands (`login`, `status`).
- Document all reusable helper functions (`GetConfigFilePath()`, `LoadConfig()`, `SaveConfig()`).

#### Phase 3 — Gap Analysis
- Clearly categorize:
  1. What is 100% completed in Step 2
  2. What is partially completed
  3. What is completely missing for Step 3
  4. What needs verification or testing

#### Phase 4 — Task Breakdown
- Divide Step 3 into small, practical, dependency-ordered tasks.

#### Phase 5 — Requirement Mapping
- Map every Step 3 requirement to its corresponding task(s) from Phase 4.

#### Phase 6 — Reuse & Duplication Check
- Audit proposed tasks against existing Step 2 code to prevent duplication.

#### Phase 7 — Validation & Test Plan
- Define an end-to-end testing strategy for Step 3.

#### Phase 8 — Final Implementation Plan & Execution Order
- Present the structured implementation order.

---

### CRITICAL STOPPING INSTRUCTION
Once you have generated the complete 8-Phase Analysis and Implementation Plan:
STOP IMMEDIATELY AND WAIT FOR MY APPROVAL.
