import { test, expect, Page } from '@playwright/test';
import AxeBuilder from '@axe-core/playwright';

// Portfolio slugs mirror ProjectOrder in internal/view/portfolio.go and
// landing pages mirror IndustryOrder in internal/view/industries.go; add
// new entries here too. Content is fixed in code, so all get screenshots.
const work = [
  'northstar-aluminum',
  'schluters-metal-art',
  'rockabilly-roasting',
  'lo-mo-outfitting',
  'nautilus-group',
  'a-team-asphalt',
  'a-team-gutters',
  'traver-hardwood-floors',
];

const routes = [
  '/',
  '/services',
  '/work',
  ...work.map((s) => `/work/${s}`),
  '/process',
  '/about',
  '/contact',
  '/privacy',
  '/terms',
  '/industries',
  '/chiropractor-website-design',
  '/law-firm-website-design',
  '/demos/chiropractor/',
  '/demos/law/',
];

// Seeded Math.random, so anything random (the law demo's confetti) lands
// the same on every run.
async function seedRandom(page: Page) {
  await page.addInitScript(() => {
    let s = 0x9e3779b9;
    Math.random = () => {
      s = (s + 0x6d2b79f5) | 0;
      let t = Math.imul(s ^ (s >>> 15), 1 | s);
      t = (t + Math.imul(t ^ (t >>> 7), 61 | t)) ^ t;
      return ((t ^ (t >>> 14)) >>> 0) / 4294967296;
    };
  });
}

for (const path of routes) {
  test.describe(path, () => {
    test('loads without errors and passes axe (WCAG 2.1 AA)', async ({ page }) => {
      const errors: string[] = [];
      page.on('pageerror', (e) => errors.push(e.message));
      page.on('console', (m) => m.type() === 'error' && errors.push(m.text()));

      // Check colours at rest, not mid-transition or mid-reveal.
      await page.emulateMedia({ reducedMotion: 'reduce' });
      const res = await page.goto(path);
      expect(res?.status()).toBe(200);
      await expect(page.locator('h1')).toHaveCount(1);

      const axe = await new AxeBuilder({ page })
        .withTags(['wcag2a', 'wcag2aa', 'wcag21a', 'wcag21aa'])
        .analyze();
      const summary = axe.violations.map(
        (v) => `${v.id} (${v.impact}): ${v.help} → ${v.nodes.map((n) => n.target.join(' ')).join(', ')}`,
      );
      expect(summary, 'axe violations').toEqual([]);
      expect(errors, 'console/page errors').toEqual([]);
    });

    test('matches screenshot', async ({ page }) => {
      // Screenshots check layout, not motion.
      await page.emulateMedia({ reducedMotion: 'reduce' });
      await seedRandom(page);
      await page.goto(path);
      await page.evaluate(async () => {
        // Full-page captures include below-the-fold lazy images; load them
        // all first or the two "stable" frames never match.
        const imgs = [...document.images];
        imgs.forEach((img) => (img.loading = 'eager'));
        await Promise.all(imgs.map((img) => img.decode().catch(() => {})));
        await document.fonts.ready;
        // A page height ending in half a pixel is captured 1px taller or
        // shorter from frame to frame. Round it up to a whole pixel.
        const html = document.documentElement;
        html.style.minHeight = Math.ceil(html.getBoundingClientRect().height) + 'px';
      });
      await expect(page).toHaveScreenshot({ fullPage: true });
    });
  });
}
