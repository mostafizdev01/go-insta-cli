# Step 3: Requirements & Acceptance Criteria

## Target Scope
Implement the Instagram Authentication Module, including interactive terminal credential prompts, session validation, logout subcommand, and the `RequireAuth()` guard function to protect restricted subcommands.

---

## Requested Requirements
1. **Interactive Login Flow**:
   - Update `insta login` to prompt interactively for Instagram Username and Password / Session Cookie (`sessionid`).
   - Securely capture terminal input.
2. **Session Persistence**:
   - Validate token format and save active session to `~/.config/insta-cli/config.json` with `IsLoggedIn = true`, active username, encrypted session token, and ISO timestamp.
3. **Logout Subcommand (`insta logout`)**:
   - Implement `logout` command to wipe session token, update `IsLoggedIn = false`, and display notification.
4. **Authentication Guard (`RequireAuth`)**:
   - Create helper function `RequireAuth(cfg config.Config) bool` in package `pkg/auth`.
   - If user is not logged in, print red error `"Error: You must be logged in. Run 'insta login' first."` and block execution.
5. **Protect Subcommands**:
   - Integrate `RequireAuth()` into `posts`, `delete`, and `show ui` subcommands.

---

## Verification Criteria
- Running `insta posts` when logged out prints red error and exits with code 1.
- Running `insta login` prompts interactively and saves session to `config.json`.
- Running `insta status` confirms active login.
- Running `insta logout` wipes session token and marks status as Logged Out.
