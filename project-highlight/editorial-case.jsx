/* global React, CASE, ImagePlaceholder, HeroBackdrop, FooterStrip, NavStrip, CaseEyebrow, MetaBlock */

/* Editorial cut — quiet, type-forward.
 * Dark hero (Cormorant italic, full-bleed photo placeholder, no card),
 * meta strip on Birchwood, brief on Birchwood with one carded screen,
 * pull quote on Flint, two-up screens on Parchment, light CTA. */

function EditorialCase({ flintGeometry = false }) {
  const featureBorder = flintGeometry ? { border: '2px solid #a84812', borderRadius: 0 } : { border: '0.5px solid #d4cfc9', borderRadius: 6 };
  const chipClass = flintGeometry ? 'fc-chip-tr' : '';

  return (
    <div style={{ background: '#faf9f7', fontFamily: 'var(--fc-font-body)', color: '#4a4038' }}>
      <NavStrip tone="dark" />

      {/* ── Hero — full-bleed dark photograph with scrim, eyebrow + Cormorant italic title ── */}
      <section style={{ background: '#1c1e24', position: 'relative', minHeight: 560 }}>
        <HeroBackdrop scrim={0.55} />
        <div style={{
          position: 'relative', zIndex: 2,
          maxWidth: 1100, margin: '0 auto',
          padding: '120px 64px 96px',
        }}>
          <div style={{ marginBottom: 22 }}>
            <div style={{ width: 28, height: 1, background: '#e8873b', marginBottom: 10 }} />
            <span style={{
              fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500,
              letterSpacing: '0.2em', textTransform: 'uppercase', color: '#e8873b',
            }}>Selected work · Case 04</span>
          </div>
          <h1 style={{
            margin: 0, fontFamily: 'var(--fc-font-display)',
            fontSize: 'clamp(46px, 5.6vw, 76px)', fontWeight: 400,
            lineHeight: 1.04, letterSpacing: '-0.02em',
            color: '#faf9f7', maxWidth: 920,
          }}>
            Rockabilly Roasting Co. — <em style={{ fontStyle: 'italic', color: '#e8873b' }}>a site that pours like the cup.</em>
          </h1>
          <p style={{
            margin: '32px 0 0', maxWidth: 560,
            fontFamily: 'var(--fc-font-body)', fontSize: 18,
            lineHeight: 1.7, color: '#c4bdb2',
          }}>
            Specialty coffee, roasted in Helena. The cafe was already full; the website needed to catch up.
          </p>
        </div>
      </section>

      {/* ── Meta strip — Parchment, between hero and brief ── */}
      <section style={{ background: '#f0ece5', borderBottom: '0.5px solid #d4cfc9' }}>
        <div style={{ maxWidth: 1100, margin: '0 auto' }}>
          <MetaBlock />
        </div>
      </section>

      {/* ── The brief — 5/7 split, eyebrow + headline left, body right ── */}
      <section style={{ background: '#faf9f7', padding: '112px 64px' }}>
        <div style={{
          maxWidth: 1100, margin: '0 auto',
          display: 'grid', gridTemplateColumns: '5fr 7fr', gap: 80, alignItems: 'start',
        }}>
          <div>
            <CaseEyebrow>{CASE.brief.eyebrow}</CaseEyebrow>
            <h2 style={{
              margin: 0, fontFamily: 'var(--fc-font-display)',
              fontSize: 38, fontWeight: 400,
              color: '#1c1e24', lineHeight: 1.18, letterSpacing: '-0.01em',
              textWrap: 'pretty',
            }}>{CASE.brief.headline}</h2>
          </div>
          <div>
            {CASE.brief.body.map((p, i) => (
              <p key={i} style={{
                margin: i === 0 ? 0 : '20px 0 0',
                fontFamily: 'var(--fc-font-body)', fontSize: 17,
                lineHeight: 1.75, color: '#4a4038',
              }}>{p}</p>
            ))}
          </div>
        </div>
      </section>

      {/* ── Featured screen — single carded mockup, slightly larger ── */}
      <section style={{ background: '#faf9f7', padding: '0 64px 112px' }}>
        <div style={{ maxWidth: 1100, margin: '0 auto' }}>
          <CaseEyebrow>The home page</CaseEyebrow>
          <div className={chipClass} style={{ ...featureBorder, padding: 0, overflow: 'hidden', position: 'relative' }}>
            <ImagePlaceholder
              label="Home page · hero, menu lockup, subscribe block"
              ratio="16 / 9"
              accent
            />
          </div>
          <p style={{
            marginTop: 18, fontFamily: 'var(--fc-font-body)',
            fontSize: 13, color: '#6b5f52', letterSpacing: '0.02em',
          }}>
            01 — Home. Roast date and origin foregrounded; cafe hours one click away.
          </p>
        </div>
      </section>

      {/* ── Pull quote — Flint surface, single column, italic Cormorant ── */}
      <section style={{ background: '#363b42', padding: '120px 64px', borderTop: '0.5px solid #474d55', borderBottom: '0.5px solid #474d55' }}>
        <div style={{ maxWidth: 880, margin: '0 auto', textAlign: 'left' }}>
          <div style={{ width: 28, height: 1, background: '#e8873b', marginBottom: 28 }} />
          <p style={{
            margin: 0,
            fontFamily: 'var(--fc-font-display)', fontStyle: 'italic',
            fontSize: 'clamp(28px, 3.4vw, 40px)',
            lineHeight: 1.35, letterSpacing: '0.005em',
            color: '#faf9f7', textWrap: 'pretty',
          }}>
            “{CASE.quote.text}”
          </p>
          <div style={{ marginTop: 36, display: 'flex', alignItems: 'center', gap: 14 }}>
            <span style={{ fontSize: 14, color: '#e8873b', fontFamily: 'var(--fc-font-body)' }}>★★★★★</span>
            <span style={{
              fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500,
              letterSpacing: '0.2em', textTransform: 'uppercase', color: '#96897a',
            }}>
              {CASE.quote.attribution} · {CASE.quote.role}
            </span>
          </div>
        </div>
      </section>

      {/* ── Two-up screens — carded, on Birchwood ── */}
      <section style={{ background: '#faf9f7', padding: '112px 64px' }}>
        <div style={{ maxWidth: 1100, margin: '0 auto' }}>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'baseline', marginBottom: 40, gap: 32, flexWrap: 'wrap' }}>
            <CaseEyebrow>Selected screens</CaseEyebrow>
            <span style={{ fontFamily: 'var(--fc-font-body)', fontSize: 13, color: '#6b5f52' }}>
              02 — 03
            </span>
          </div>
          <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: 28 }}>
            <div>
              <div className={chipClass} style={{ border: '0.5px solid #d4cfc9', borderRadius: flintGeometry ? 0 : 6, overflow: 'hidden', background: '#f0ece5' }}>
                <ImagePlaceholder label="Single-origin product page · roast notes, brew guide" ratio="4 / 3" />
              </div>
              <p style={{ marginTop: 14, fontFamily: 'var(--fc-font-body)', fontSize: 13, color: '#6b5f52' }}>
                02 — Bean detail. Origin and roast date sit above the price.
              </p>
            </div>
            <div>
              <div className={chipClass} style={{ border: '0.5px solid #d4cfc9', borderRadius: flintGeometry ? 0 : 6, overflow: 'hidden', background: '#f0ece5' }}>
                <ImagePlaceholder label="Wholesale form · cafes, restaurants, offices" ratio="4 / 3" />
              </div>
              <p style={{ marginTop: 14, fontFamily: 'var(--fc-font-body)', fontSize: 13, color: '#6b5f52' }}>
                03 — Wholesale. Three fields, plain English, no portal login.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* ── Closing CTA on Parchment ── */}
      <section style={{ background: '#f0ece5', padding: '112px 64px', borderTop: '0.5px solid #d4cfc9' }}>
        <div style={{ maxWidth: 720, margin: '0 auto' }}>
          <CaseEyebrow>Next</CaseEyebrow>
          <h2 style={{
            margin: 0, fontFamily: 'var(--fc-font-display)',
            fontSize: 'clamp(32px, 3.4vw, 44px)', fontWeight: 400,
            lineHeight: 1.15, color: '#1c1e24', letterSpacing: '-0.01em',
          }}>
            Have a business that's outgrown its first site?
          </h2>
          <p style={{
            margin: '20px 0 32px', fontFamily: 'var(--fc-font-body)',
            fontSize: 17, color: '#4a4038', lineHeight: 1.7, maxWidth: 540,
          }}>
            The first conversation costs nothing. No pitch, no pressure. Just an honest read on whether it makes sense.
          </p>
          <div style={{ display: 'flex', gap: 20, alignItems: 'center', flexWrap: 'wrap' }}>
            <a className={`fc-flint-btn ${flintGeometry ? 'fc-chip-btn' : ''}`} style={{
              background: '#a84812', color: '#faf9f7',
              padding: '14px 26px', fontSize: 13,
              borderRadius: flintGeometry ? 0 : 3,
            }}>Start a project</a>
            <a style={{
              fontFamily: 'var(--fc-font-body)', fontSize: 13, fontWeight: 500,
              letterSpacing: '0.08em', textTransform: 'uppercase',
              color: '#a84812',
              borderBottom: '1.5px solid #a84812', paddingBottom: 4,
            }}>See more work →</a>
          </div>
        </div>
      </section>

      <FooterStrip tone="dark" />
    </div>
  );
}

window.EditorialCase = EditorialCase;
