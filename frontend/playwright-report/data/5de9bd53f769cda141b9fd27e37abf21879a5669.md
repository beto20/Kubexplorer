# Test info

- Name: visits the app root url
- Location: /Users/beto/Documents_local/1.AmbDesarrollo/1.AMB_Interno/kubexplorer/frontend/e2e/vue.spec.ts:5:1

# Error details

```
Error: browserType.launch: Executable doesn't exist at /Users/beto/Library/Caches/ms-playwright/chromium-1169/chrome-mac/Chromium.app/Contents/MacOS/Chromium
╔═════════════════════════════════════════════════════════════════════════╗
║ Looks like Playwright Test or Playwright was just installed or updated. ║
║ Please run the following command to download new browsers:              ║
║                                                                         ║
║     npx playwright install                                              ║
║                                                                         ║
║ <3 Playwright Team                                                      ║
╚═════════════════════════════════════════════════════════════════════════╝
```

# Test source

```ts
  1 | import { test, expect } from '@playwright/test';
  2 |
  3 | // See here how to get started:
  4 | // https://playwright.dev/docs/intro
> 5 | test('visits the app root url', async ({ page }) => {
    | ^ Error: browserType.launch: Executable doesn't exist at /Users/beto/Library/Caches/ms-playwright/chromium-1169/chrome-mac/Chromium.app/Contents/MacOS/Chromium
  6 |   await page.goto('/');
  7 |   await expect(page.locator('h1')).toHaveText('You did it!');
  8 | })
  9 |
```