Act as a Principal Software Architect and Lead Developer. 

I need you to help me plan and implement **Step 8: Auto Browser Launch & CLI-Web Sync Engine** for my Instagram CLI project (`go-insta-cli`), building directly on top of the work completed in **Step 7** without duplicating existing code or breaking the established architecture.

---

### STEP 8 SPECIFIC REQUIREMENTS
1. Implement Cross-Platform Auto Browser Launcher:
   - When user executes `insta show ui`, automatically launch the OS default web browser opening `http://localhost:8080/`.
   - Support cross-platform browser opening helper (Windows: `cmd /c start`, macOS: `open`, Linux: `xdg-open`).
2. Implement Non-Blocking Fallback: If auto-browser launch fails or user is in a headless environment, gracefully fall back to printing a clickable ANSI colored terminal link: `"🔗 Open Dashboard: http://localhost:8080/"`.
3. Implement Real-Time Data Sync Engine:
   - Ensure changes made via CLI (e.g., `insta delete <id>`) are instantly reflected when Web UI refreshes.
   - Implement periodic polling or event-based auto-refresh in the Web UI frontend (Step 7) so the dashboard automatically updates without manual page reload.
4. Maintain `RequireAuth()` guard to prevent unauthenticated server startup or browser launch.

---

### STRICT OPERATING RULES
1. DO NOT WRITE, MODIFY, OR GENERATE ANY PROJECT CODE IMMEDIATELY.
2. Perform a complete Audit and Gap Analysis of Step 7 first.
3. Prevent code duplication — reuse and extend existing files (`main.go`, Step 7 Embedded Web Dashboard UI, Step 6 REST API endpoints), functions, and handlers wherever possible.
4. Do not make assumptions. If any file, context, or specification is missing, explicitly ask for it.
5. Follow the 8-Phase Planning Process detailed below.
6. CRITICAL STOPPING RULE: After presenting the complete Phase 1 - Phase 8 analysis and plan, STOP and wait for my explicit approval. Do NOT generate implementation code until I approve the plan.

---

### REQUIRED 8-PHASE PLANNING WORKFLOW

#### Phase 1 — Understand
- Review Step 8 requirements (Cross-platform Browser Auto-Launcher, Fallback Terminal Link, Two-Way Data Sync Engine, and Web UI auto-refresh/polling).

#### Phase 2 — Audit Existing Work (Step 7)
- Inspect completed Step 7 files (`main.go`, Embedded Web Dashboard UI, REST API endpoints `/api/status`, `/api/posts`, `/api/delete`), server startup logic, and ANSI color constants.
- Document all reusable helper functions and frontend scripts.

#### Phase 3 — Gap Analysis
- Clearly categorize:
  1. What is 100% completed in Step 7
  2. What is partially completed
  3. What is completely missing for Step 8
  4. What needs verification or testing

#### Phase 4 — Task Breakdown
- Divide Step 8 into small, practical, dependency-ordered tasks.
- For EACH task, specify:
  • Task Name & Description: What needs to be done & Why it is needed.
  • Files Involved: Existing files to modify vs New files to create.
  • Dependencies: Which previous tasks must be finished first.
  • Expected Outcome: The precise result of this task.
  • Testing & Verification: How to empirically test this task.

#### Phase 5 — Requirement Mapping
- Map every Step 8 requirement to its corresponding task(s) from Phase 4 to ensure 100% coverage with zero missed requirements.

#### Phase 6 — Reuse & Duplication Check
- Explicitly audit every proposed task against existing Step 7 code to guarantee no browser launch or sync logic duplicates existing HTTP handlers redundantly.

#### Phase 7 — Validation Plan
- Define an end-to-end testing strategy (testing `insta show ui` auto-opening default browser on Windows/Mac/Linux, verifying fallback link output when browser launch is disabled, testing post deletion via CLI while Web UI is open to verify real-time dashboard auto-refresh).

#### Phase 8 — Final Implementation Plan & Execution Order
- Present the structured, step-by-step implementation order.

---

### CRITICAL STOPPING INSTRUCTION
Once you have generated the complete 8-Phase Analysis and Implementation Plan:
STOP IMMEDIATELY AND WAIT FOR MY APPROVAL. 
Do not output any code files until I approve the plan.
