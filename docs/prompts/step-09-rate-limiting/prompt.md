Act as a Principal Software Architect and Lead Developer. 

I need you to help me plan and implement **Step 9: API Rate Limiting & Account Protection Controls** for my Instagram CLI project (`go-insta-cli`), building directly on top of the work completed in **Step 8** without duplicating existing code or breaking the established architecture.

---

### STEP 9 SPECIFIC REQUIREMENTS
1. Implement Request Throttling & Rate Limiter Engine:
   - Create a central rate limiter/throttler middleware intercepting all Instagram API requests (Post Retrieval, Post Deletion, Authentication).
   - Enforce mandatory safety delays (e.g., 1-2 seconds) between API requests to mimic human behavior and keep the Instagram account safe.
2. Implement Bulk Action Throttling:
   - When deleting multiple posts (bulk delete via CLI or Web UI), execute requests sequentially with a forced safety interval (e.g., 3 seconds per deletion) instead of parallel spamming.
3. Implement Exponential Backoff & Circuit Breaker:
   - If Instagram returns HTTP 429 (Too Many Requests) or a rate limit flag, automatically apply exponential backoff.
   - Circuit Breaker: If rate limit is hit repeatedly, immediately halt execution, print a red/yellow terminal warning (`"⚠ Rate limit detected. Pausing requests for account safety."`), and return user-friendly feedback to Web UI.
4. Expose Rate Limit status info in `insta status` CLI output and `GET /api/status` REST API response.

---

### STRICT OPERATING RULES
1. DO NOT WRITE, MODIFY, OR GENERATE ANY PROJECT CODE IMMEDIATELY.
2. Perform a complete Audit and Gap Analysis of Step 8 first.
3. Prevent code duplication — reuse and extend existing files (`main.go`, Step 4 Post Retrieval logic, Step 5 Post Deletion logic, Step 6 REST API handlers, Step 7 Web UI), functions, and handlers wherever possible.
4. Do not make assumptions. If any file, context, or specification is missing, explicitly ask for it.
5. Follow the 8-Phase Planning Process detailed below.
6. CRITICAL STOPPING RULE: After presenting the complete Phase 1 - Phase 8 analysis and plan, STOP and wait for my explicit approval. Do NOT generate implementation code until I approve the plan.

---

### REQUIRED 8-PHASE PLANNING WORKFLOW

#### Phase 1 — Understand
- Review Step 9 requirements (Rate Limiter Engine, Safety Delays, Sequential Bulk Throttling, Exponential Backoff, Circuit Breaker, and Rate Limit status feedback).

#### Phase 2 — Audit Existing Work (Step 8)
- Inspect completed Step 8 files (`main.go`, Embedded Web Dashboard UI, Auto-browser launcher, REST API handlers `/api/status`, `/api/posts`, `/api/delete`), and API call mechanisms.
- Document all reusable API handlers and HTTP client logic.

#### Phase 3 — Gap Analysis
- Clearly categorize:
  1. What is 100% completed in Step 8
  2. What is partially completed
  3. What is completely missing for Step 9
  4. What needs verification or testing

#### Phase 4 — Task Breakdown
- Divide Step 9 into small, practical, dependency-ordered tasks.
- For EACH task, specify:
  • Task Name & Description: What needs to be done & Why it is needed.
  • Files Involved: Existing files to modify vs New files to create.
  • Dependencies: Which previous tasks must be finished first.
  • Expected Outcome: The precise result of this task.
  • Testing & Verification: How to empirically test this task.

#### Phase 5 — Requirement Mapping
- Map every Step 9 requirement to its corresponding task(s) from Phase 4 to ensure 100% coverage with zero missed requirements.

#### Phase 6 — Reuse & Duplication Check
- Explicitly audit every proposed task against existing Step 8 code to guarantee rate limiting wrappers wrap existing HTTP handlers without duplicating post retrieval or deletion code redundantly.

#### Phase 7 — Validation Plan
- Define an end-to-end testing strategy (testing sequential delay during bulk post deletion, testing rate limit simulation / HTTP 429 response handling, verifying Circuit Breaker warning in terminal and Web UI, and verifying account safety throttle logs).

#### Phase 8 — Final Implementation Plan & Execution Order
- Present the structured, step-by-step implementation order.

---

### CRITICAL STOPPING INSTRUCTION
Once you have generated the complete 8-Phase Analysis and Implementation Plan:
STOP IMMEDIATELY AND WAIT FOR MY APPROVAL. 
Do not output any code files until I approve the plan.
