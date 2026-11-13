import { test, expect, Page } from '@playwright/test';

// What Go tests can't see: that htmx swaps the /contact form in place (no
// full page load) with the field errors or the thank-you, and that its
// hx-on::after-settle handler moves focus to the error summary or the
// thank-you so screen readers announce the result. Postmark is off
// in the test server (playwright.config.ts refuses to start if .env sets
// it), so a valid submission is only logged.

// Marks the window so we can tell an htmx swap from a full page load.
async function markWindow(page: Page) {
  await page.evaluate(() => ((window as any).__noReload = true));
}
async function stillSameWindow(page: Page) {
  return page.evaluate(() => (window as any).__noReload === true);
}

// The inputs carry `required` and type="email", so the browser would block
// these submits before htmx sees them. Turn that off to exercise the
// server's validation.
async function skipBrowserValidation(page: Page) {
  await page.evaluate(() => document.querySelector('form[hx-post="/contact"]')!.setAttribute('novalidate', ''));
}

test('empty submit swaps in the error summary and focuses it', async ({ page }) => {
  await page.goto('/contact');
  await markWindow(page);
  await skipBrowserValidation(page);

  await page.getByRole('button', { name: 'Send it' }).click();

  const summary = page.locator('#contact-errsum');
  await expect(summary).toBeVisible();
  await expect(summary).toBeFocused();
  await expect(summary).toContainText('Name is required.');
  await expect(summary).toContainText('Email is required.');
  await expect(summary).toContainText('Message is required.');
  // The inline error stays under each field too.
  await expect(page.locator('#name ~ p')).toHaveText('Name is required.');
  expect(await stillSameWindow(page), 'form should swap via htmx, not reload').toBe(true);
});

test('bad email is rejected and typed values are kept', async ({ page }) => {
  await page.goto('/contact');
  await skipBrowserValidation(page);
  await page.locator('#name').fill('Jordan');
  await page.locator('#email').fill('not-an-email');
  await page.locator('#message').fill('We need a new site.');
  await page.getByRole('button', { name: 'Send it' }).click();

  const summary = page.locator('#contact-errsum');
  await expect(summary).toContainText('Enter a valid email address.');
  await expect(summary).toBeFocused();
  await expect(summary).not.toContainText('Name is required.');
  await expect(page.locator('#name')).toHaveValue('Jordan');
  await expect(page.locator('#message')).toHaveValue('We need a new site.');
});

test('valid message swaps in the thank-you and focuses it', async ({ page }) => {
  await page.goto('/contact');
  await markWindow(page);

  await page.locator('#name').fill('Jordan');
  await page.locator('#email').fill('jordan@example.com');
  await page.locator('#message').fill('We need a new site.');
  await page.getByRole('button', { name: 'Send it' }).click();

  const sent = page.locator('#contact-sent');
  await expect(sent).toContainText("Got it. We'll look at your current site first");
  await expect(sent).toBeFocused();
  expect(await stillSameWindow(page)).toBe(true);
});
