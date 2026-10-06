# Step 2: Requirements & Acceptance Criteria

## 📌 Target Scope
Implement a secure, OS-aware local JSON storage engine for `go-insta-cli` at `~/.config/insta-cli/config.json`.

---

## 📋 Requested Requirements
1. **Config Path**: Dynamically resolve OS home directory (`~/.config/insta-cli/config.json`). Create parent directories if missing.
2. **Config Struct**:
   - `Username` (string)
   - `SessionToken` (string, obfuscated/encrypted)
   - `IsLoggedIn` (bool)
   - `LastLogin` (string ISO timestamp)
3. **Security**: Obfuscate/encrypt session token string before disk storage; write file with strict `0600` permissions.
4. **Helper Functions**:
   - `GetConfigFilePath()`
   - `LoadConfig()`
   - `SaveConfig(cfg Config)`
5. **Subcommand Updates**:
   - `login`: Update to accept `<username>`, save mock session to `config.json`, and display green success message.
   - `status`: New subcommand reading `config.json` and displaying active user, session status, and login timestamp with ANSI colors.

---

## 🎯 Verification Criteria
- `insta login testuser` writes formatted `config.json` to disk with `0600` permissions.
- `insta status` reads `config.json` and outputs current login state.
