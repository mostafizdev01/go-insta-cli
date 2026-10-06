# Step 2: Master AI Implementation Prompt

Act as a Principal Software Architect and Lead Developer. 

I need you to help me plan and implement **Step 2: Secure Configuration & Local Storage** for my Instagram CLI project (`go-insta-cli`), building directly on top of the work completed in **Step 1** without duplicating existing code or breaking the established architecture.

---

### STEP 2 SPECIFIC REQUIREMENTS
1. Set up local config file storage at `~/.config/insta-cli/config.json` (handle OS-specific user home directory dynamically).
2. Create a `Config` struct containing: `Username`, `SessionToken`, `IsLoggedIn`, and `LastLogin`.
3. Encrypt/Obfuscate and store Instagram session tokens safely on disk with secure file permissions (0600).
4. Implement config helper functions (`GetConfigPath()`, `LoadConfig()`, `SaveConfig()`).
5. Update `login` subcommand to store session data in `config.json`.
6. Add a new `status` subcommand to read `config.json` and display current login status, active username, and session state using ANSI colors.

---

### STRICT OPERATING RULES
1. DO NOT WRITE, MODIFY, OR GENERATE ANY PROJECT CODE IMMEDIATELY.
2. Perform a complete Audit and Gap Analysis of Step 1 first.
3. Prevent code duplication — reuse and extend existing files (`main.go`), functions, and logic wherever possible.
4. Do not make assumptions. If any file, context, or specification is missing, explicitly ask for it.
5. Follow the 8-Phase Planning Process detailed below.
6. CRITICAL STOPPING RULE: After presenting the complete Phase 1 - Phase 8 analysis and plan, STOP and wait for my explicit approval. Do NOT generate implementation code until I approve the plan.

---

### REQUIRED 8-PHASE PLANNING WORKFLOW

#### Phase 1 — Understand
- Review Step 2 requirements (`~/.config/insta-cli/config.json`, encryption, session management, `login` and `status` subcommands).

#### Phase 2 — Audit Existing Work (Step 1)
- Inspect completed Step 1 files (`main.go`, `go.mod`), CLI flags (`--version`, `--help`), ANSI color constants, and existing subcommand skeletons.
- Document all reusable components and helper functions.

#### Phase 3 — Gap Analysis
- Clearly categorize:
  1. What is 100% completed in Step 1
  2. What is partially completed
  3. What is completely missing for Step 2
  4. What needs verification or testing

#### Phase 4 — Task Breakdown
- Divide Step 2 into small, practical, dependency-ordered tasks.
- For EACH task, specify:
  • Task Name & Description: What needs to be done & Why it is needed.
  • Files Involved: Existing files to modify vs New files to create.
  • Dependencies: Which previous tasks must be finished first.
  • Expected Outcome: The precise result of this task.
  • Testing & Verification: How the result can be tested.

#### Phase 5 — Requirement Mapping
- Map every Step 2 requirement to its corresponding task(s) from Phase 4.

#### Phase 6 — Reuse & Duplication Check
- Audit proposed tasks against existing Step 1 code to prevent duplication.

#### Phase 7 — Validation & Test Plan
- Define an end-to-end testing strategy for Step 2.

#### Phase 8 — Final Implementation Plan & Execution Order
- Present the structured implementation order.

---

### CRITICAL STOPPING INSTRUCTION
Once you have generated the complete 8-Phase Analysis and Implementation Plan:
STOP IMMEDIATELY AND WAIT FOR MY APPROVAL.
