# Step 10: 8-Phase Architectural Plan & Audit Log

- **Target Step**: Step 10 - Packaging, Cross-Platform Distribution & Release Automation
- **Status**: PLAN_APPROVED_READY_FOR_IMPLEMENTATION
- **Timestamp**: 2026-10-07

---

## Phase 1 — Understand
- Cross-platform builds (`windows`, `linux`, `darwin`).
- Dynamic `-ldflags` version injection.
- GitHub Actions CI/CD automation & `.goreleaser.yaml`.
- SHA-256 `checksums.txt` generation.
- End-to-end binary verification and `README.md`.

---

## Phase 2 — Audit of Existing Work
- `main.go`, `cmd_handlers.go`: Version string and subcommands.
- `pkg/server/`: Embedded Web UI and REST API.
- `pkg/instagram/`: Meta Graph API integration and rate limiter.

---

## Phase 3 — Gap Analysis
- [ ] Task 1: Cross-compilation script (`scripts/build.ps1` / `Makefile`).
- [ ] Task 2: Dynamic `-ldflags` version injection setup in `main.go`.
- [ ] Task 3: `.goreleaser.yaml` and `.github/workflows/release.yml` CI/CD pipeline.
- [ ] Task 4: Complete `README.md` user and installation guide.

---

## Phase 4 — Task Breakdown
1. **Task 1: Cross-Compilation Script**: Create build automation for Windows, Linux, macOS.
2. **Task 2: Dynamic `-ldflags` Injection**: Update `main.go` for build date & commit hash variables.
3. **Task 3: Release Pipeline (CI/CD)**: Configure `.goreleaser.yaml` and `.github/workflows/release.yml`.
4. **Task 4: System Verification & `README.md`**: Perform end-to-end verification and generate documentation.

---

## Phase 5 — Requirement Mapping
- Req 1 & 2 -> Task 1 & Task 2
- Req 3 -> Task 3
- Req 4 & 5 -> Task 4

---

## Phase 6 — Reuse & Duplication Check
- 100% reuse of existing application code, subcommands, and embedded assets.

---

## Phase 7 — Validation Plan
- Test cross-compilation for `GOOS=windows`, `GOOS=linux`, `GOOS=darwin`.
- Verify `insta --version` injected output.
- Verify `README.md` and release workflow.
