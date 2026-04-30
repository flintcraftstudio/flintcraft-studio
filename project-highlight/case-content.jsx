/* global React */

/* Shared content + image placeholders for the Rockabilly Roasting Co. case study.
 * No real client photography — we draw warm placeholders in the FlintCraft idiom:
 * Obsidian / Flint / Parchment fields with monospace labels for what should land there. */

const CASE = {
  client: 'Rockabilly Roasting Co.',
  location: 'Kennewick, Washington',
  year: '2025',
  type: 'Marketing site · E-commerce',
  url: 'rockabillyroasting.com',
  tagline: 'Small-batch coffee, roasted in Kennewick. A site that pours like the cup.',

  brief: {
    eyebrow: 'The brief',
    headline: 'A roaster known by the regulars, invisible to everyone else.',
    body: [
      "Rockabilly had outgrown its first site. The cafe was full most mornings; the wholesale list kept growing on word of mouth alone. The website didn't reflect any of it — it loaded slowly, buried the weekly roast, and had no clear way to buy a bag of beans without driving in.",
      "The owners wanted one thing made obvious: small-batch coffee, roasted weekly, poured without fuss. Everything else — the cafe hours, the subscription, the wholesale form — needed to sit underneath that, calmly, in plain English.",
    ],
  },

  quote: {
    text: "FlintCraft Studio is an absolutely must have company working for you in your corner. They are super approachable, helpful and friendly. If you are looking to take your company to the next level I would highly recommend you give FlintCraft a call.",
    attribution: 'Kagen Cox',
    role: 'Owner, Rockabilly Roasting Co.',
  },
};

/* ─── Image placeholder primitive ───────────────────────────────────────────
 * Subtly-striped, monospace-labelled placeholder. Honors light/dark surfaces.
 * Renders inside the 6px card frame the system prescribes.
 */
function ImagePlaceholder({ label, ratio = '16 / 10', dark = false, accent = false, fill = false }) {
  const bg = dark ? '#1c1e24' : '#f0ece5';
  const stripe = dark ? 'rgba(196, 189, 178, 0.045)' : 'rgba(28, 30, 36, 0.035)';
  const text = dark ? '#96897a' : '#6b5f52';
  const accentColor = dark ? '#e8873b' : '#a84812';

  return (
    <div style={{
      aspectRatio: fill ? undefined : ratio,
      width: '100%',
      height: fill ? '100%' : undefined,
      background: `repeating-linear-gradient(135deg, ${bg} 0 14px, ${stripe} 14px 28px)`,
      borderRadius: 6,
      border: dark ? '0.5px solid #474d55' : '0.5px solid #d4cfc9',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
      position: 'relative',
      overflow: 'hidden',
    }}>
      {accent && (
        <div style={{
          position: 'absolute', top: 14, left: 14,
          width: 28, height: 1, background: accentColor,
        }} />
      )}
      <div style={{
        fontFamily: 'SF Mono, ui-monospace, Menlo, monospace',
        fontSize: 11,
        letterSpacing: '0.14em',
        textTransform: 'uppercase',
        color: text,
        textAlign: 'center',
        padding: '0 24px',
        lineHeight: 1.5,
      }}>
        {label}
      </div>
    </div>
  );
}

/* Hero placeholder — full-bleed, with logo watermark + scrim, per system rule
 * "the only gradient permitted is a hero scrim over a full-bleed photograph". */
function HeroBackdrop({ children, scrim = 0.55 }) {
  return (
    <div style={{
      position: 'absolute',
      inset: 0,
      background: `repeating-linear-gradient(135deg, #1c1e24 0 18px, #232730 18px 36px)`,
      overflow: 'hidden',
    }}>
      <img src="logo.svg" alt=""
           style={{ position: 'absolute', right: '-60px', bottom: '-40px', width: 460, height: 'auto', opacity: 0.14 }} />
      <div style={{
        position: 'absolute', inset: 0,
        background: `linear-gradient(180deg, rgba(28,30,36,${scrim - 0.15}) 0%, rgba(28,30,36,${scrim}) 60%, rgba(28,30,36,${scrim + 0.05}) 100%)`,
      }} />
      <div style={{
        position: 'absolute', top: 18, left: 18,
        fontFamily: 'SF Mono, ui-monospace, Menlo, monospace',
        fontSize: 10, letterSpacing: '0.18em', textTransform: 'uppercase',
        color: 'rgba(196, 189, 178, 0.55)',
      }}>
        Hero photograph · cafe interior at open
      </div>
      {children}
    </div>
  );
}

/* Tiny FlintCraft footer used inside each artboard so the page reads as a real site. */
function FooterStrip({ tone = 'dark' }) {
  const dark = tone === 'dark';
  return (
    <footer style={{
      background: dark ? '#1c1e24' : '#faf9f7',
      borderTop: dark ? '0.5px solid #363b42' : '0.5px solid #d4cfc9',
      padding: '36px 64px',
      display: 'flex', justifyContent: 'space-between', alignItems: 'center',
      gap: 24, flexWrap: 'wrap',
    }}>
      <div style={{ display: 'flex', alignItems: 'center', gap: 14 }}>
        <img src="logo.svg" alt="" style={{ height: 26, width: 'auto', opacity: dark ? 0.85 : 1 }} />
        <span style={{
          fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500,
          letterSpacing: '0.18em', textTransform: 'uppercase',
          color: dark ? '#96897a' : '#6b5f52',
        }}>
          FlintCraft Studio · Helena, MT
        </span>
      </div>
      <span style={{
        fontFamily: 'var(--fc-font-display)', fontStyle: 'italic',
        fontSize: 15, color: dark ? '#c4bdb2' : '#4a4038',
      }}>
        Still shaped by hand.
      </span>
    </footer>
  );
}

/* Lightweight nav strip matching the marketing site. */
function NavStrip({ tone = 'dark' }) {
  const dark = tone === 'dark';
  const fg = dark ? '#c4bdb2' : '#6b5f52';
  const accent = dark ? '#e8873b' : '#a84812';
  return (
    <nav style={{
      background: dark ? '#1c1e24' : '#faf9f7',
      borderBottom: dark ? '0.5px solid #363b42' : '0.5px solid #d4cfc9',
      padding: '20px 64px',
      display: 'flex', justifyContent: 'space-between', alignItems: 'center',
      position: 'relative', zIndex: 5,
    }}>
      <div style={{ display: 'flex', alignItems: 'center', gap: 12 }}>
        <img src="logo.svg" alt="" style={{ height: 24, width: 'auto' }} />
        <span style={{
          fontFamily: 'var(--fc-font-display)',
          fontSize: 18, color: dark ? '#f0ece5' : '#1c1e24',
          letterSpacing: '0.01em',
        }}>FlintCraft Studio</span>
      </div>
      <div style={{ display: 'flex', gap: 32, alignItems: 'center' }}>
        {['Services', 'Process', 'Work', 'Contact'].map((l, i) => (
          <span key={l} style={{
            fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500,
            letterSpacing: '0.18em', textTransform: 'uppercase',
            color: l === 'Work' ? accent : fg,
          }}>{l}</span>
        ))}
      </div>
    </nav>
  );
}

/* Eyebrow with the small accent rule above it. */
function CaseEyebrow({ children, dark = false, align = 'left' }) {
  return (
    <div style={{ marginBottom: 14, textAlign: align }}>
      <div style={{
        width: 28, height: 1,
        background: dark ? '#e8873b' : '#a84812',
        marginBottom: 10,
        marginLeft: align === 'center' ? 'auto' : 0,
        marginRight: align === 'center' ? 'auto' : 0,
      }} />
      <span style={{
        fontFamily: 'var(--fc-font-body)',
        fontSize: 11, fontWeight: 500, letterSpacing: '0.2em',
        textTransform: 'uppercase',
        color: dark ? '#96897a' : '#6b5f52',
      }}>{children}</span>
    </div>
  );
}

/* Project metadata block — appears in both variants. */
function MetaBlock({ dark = false }) {
  const labelColor = dark ? '#96897a' : '#6b5f52';
  const valueColor = dark ? '#f0ece5' : '#1c1e24';
  const border = dark ? '#363b42' : '#d4cfc9';
  const items = [
    ['Client', CASE.client],
    ['Location', CASE.location],
    ['Year', CASE.year],
    ['Scope', CASE.type],
    ['Live at', CASE.url],
  ];
  return (
    <dl style={{
      display: 'grid', gridTemplateColumns: 'repeat(5, 1fr)',
      borderTop: `0.5px solid ${border}`,
      borderBottom: `0.5px solid ${border}`,
      margin: 0, padding: 0,
    }}>
      {items.map(([k, v], i) => (
        <div key={k} style={{
          padding: '22px 24px',
          borderRight: i < items.length - 1 ? `0.5px solid ${border}` : 'none',
        }}>
          <dt style={{
            fontFamily: 'var(--fc-font-body)', fontSize: 10, fontWeight: 500,
            letterSpacing: '0.2em', textTransform: 'uppercase',
            color: labelColor, marginBottom: 8,
          }}>{k}</dt>
          <dd style={{
            margin: 0,
            fontFamily: 'var(--fc-font-display)', fontSize: 18,
            color: valueColor, lineHeight: 1.3,
          }}>{v}</dd>
        </div>
      ))}
    </dl>
  );
}

Object.assign(window, {
  CASE, ImagePlaceholder, HeroBackdrop, FooterStrip, NavStrip, CaseEyebrow, MetaBlock,
});
