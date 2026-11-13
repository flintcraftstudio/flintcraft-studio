import { defineConfig, devices } from '@playwright/test';

// The suite builds and starts its own server on :8093, so it never
// touches a dev server on :8080. The site has no database.
const PORT = 8093;
const baseURL = `http://localhost:${PORT}`;

export default defineConfig({
  testDir: './tests',
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 1 : 0,
  reporter: [['list'], ['html', { open: 'never' }]],
  // First run writes baselines instead of failing; later runs compare.
  updateSnapshots: 'missing',
  expect: {
    toHaveScreenshot: { maxDiffPixelRatio: 0.01, animations: 'disabled' },
  },
  use: {
    baseURL,
    trace: 'retain-on-failure',
  },
  projects: [
    { name: 'desktop', use: { ...devices['Desktop Chrome'] } },
    { name: 'mobile', use: { ...devices['Pixel 7'] } },
  ],
  webServer: {
    cwd: '..',
    // main.go loads .env, and the contact test submits a valid message. If
    // .env configures Postmark that would send real email, so refuse to run.
    // The pinned templ CLI matches go.mod (see CLAUDE.md).
    command: [
      `! grep -qs '^POSTMARK_SERVER_TOKEN=.' .env || { echo 'e2e: .env sets POSTMARK_SERVER_TOKEN; the contact test would send real email' >&2; exit 1; }`,
      'go run github.com/a-h/templ/cmd/templ@v0.3.1020 generate',
      'mage buildCSS buildChiroCSS buildLawCSS',
      'go build -o e2e/.bin/server ./cmd/server',
      'e2e/.bin/server',
    ].join(' && '),
    url: baseURL,
    reuseExistingServer: !process.env.CI,
    timeout: 180_000,
    env: {
      PORT: String(PORT),
      BASE_URL: baseURL,
    },
  },
});
