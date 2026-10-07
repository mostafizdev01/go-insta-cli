Act as a Principal Software Architect and Lead Developer. 

I need you to help me plan and implement **Step 10: Packaging, Cross-Platform Distribution & Release Automation** for my Instagram CLI project (`go-insta-cli`), building directly on top of the work completed in **Step 9** without duplicating existing code or breaking the established architecture.

---

### STEP 10 SPECIFIC REQUIREMENTS
1. Implement Cross-Platform Standalone Binary Build Configurations:
   - Configure build scripts/commands to compile single static executable binaries targeting Windows (`insta.exe`), Linux (`insta-linux`), and macOS (`insta-darwin-amd64` / `insta-darwin-arm64`).
   - Ensure all static web assets (Step 7 Web UI) are embedded into the single binary (`//go:embed`) with zero external runtime dependencies.
2. Dynamic Build-Time Version Injection:
   - Inject version strings (`main.Version`), build date, and git commit hash dynamically using Go `-ldflags` (e.g. `go build -ldflags="-X main.Version=v1.0.0"`).
3. Release Automation & Pipeline Setup (CI/CD):
   - Configure GoReleaser configuration (`.goreleaser.yaml`) or GitHub Actions release workflow.
   - Automate generation of SHA-256 `checksums.txt` file and asset archives for GitHub Releases upon tagging a new release (`v1.0.0`).
4. End-to-End Integration & System Verification:
   - Perform full end-to-end integration testing of standalone binaries across CLI commands (`login`, `status`, `posts`, `delete`, `show ui`) to confirm smooth user workflows without requiring a Go installation.
5. Generate a clean installation and release guide (`README.md`).

---

### STRICT OPERATING RULES
1. DO NOT WRITE, MODIFY, OR GENERATE ANY PROJECT CODE IMMEDIATELY.
2. Perform a complete Audit and Gap Analysis of Step 9 first.
3. Prevent code duplication — reuse and extend existing files (`main.go`, Step 7 Embedded Web Dashboard UI, Step 8 Auto-browser engine, Step 9 Rate Limiter), functions, and handlers wherever possible.
4. Do not make assumptions. If any file, context, or specification is missing, explicitly ask for it.
5. Follow the 8-Phase Planning Process detailed below.
6. CRITICAL STOPPING RULE: After presenting the complete Phase 1 - Phase 8 analysis and plan, STOP and wait for my explicit approval. Do NOT generate implementation code until I approve the plan.

---

### REQUIRED 8-PHASE PLANNING WORKFLOW

#### Phase 1 — Understand
- Review Step 10 requirements (Cross-platform builds for Win/Linux/Mac, static asset embedding verification, `-ldflags` version injection, GoReleaser/GitHub Actions pipeline setup, SHA-256 checksum generation, and end-to-end testing).

#### Phase 2 — Audit Existing Work (Step 9)
- Inspect completed Step 9 files (`main.go`, Rate Limiter Engine, Embedded Web UI, Auto-browser launcher, REST API handlers `/api/status`, `/api/posts`, `/api/delete`), and build dependencies.
- Document all reusable compilation targets and embedded assets.

#### Phase 3 — Gap Analysis
- Clearly categorize:
  1. What is 100% completed in Step 9
  2. What is partially completed
  3. What is completely missing for Step 10
  4. What needs verification or testing

#### Phase 4 — Task Breakdown
- Divide Step 10 into small, practical, dependency-ordered tasks.
- For EACH task, specify:
  • Task Name & Description: What needs to be done & Why it is needed.
  • Files Involved: Existing files to modify vs New files to create (e.g. `.goreleaser.yaml`, `.github/workflows/release.yml`, `README.md`).
  • Dependencies: Which previous tasks must be finished first.
  • Expected Outcome: The precise result of this task.
  • Testing & Verification: How to empirically test this task.

#### Phase 5 — Requirement Mapping
- Map every Step 10 requirement to its corresponding task(s) from Phase 4 to ensure 100% coverage with zero missed requirements.

#### Phase 6 — Reuse & Duplication Check
- Explicitly audit every proposed task against existing Step 9 code to guarantee release configuration scripts reuse existing build targets without modifying core application logic.

#### Phase 7 — Validation Plan
- Define an end-to-end testing strategy (testing cross-compilation commands for `GOOS=windows`, `GOOS=linux`, `GOOS=darwin`, verifying version output from `-ldflags`, testing standalone executable execution without Go runtime, verifying `checksums.txt` hash validation, and verifying full end-to-end user experience).

#### Phase 8 — Final Implementation Plan & Execution Order
- Present the structured, step-by-step implementation order.

---

### CRITICAL STOPPING INSTRUCTION
Once you have generated the complete 8-Phase Analysis and Implementation Plan:
STOP IMMEDIATELY AND WAIT FOR MY APPROVAL. 
Do not output any code files until I approve the plan.
