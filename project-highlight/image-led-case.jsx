/* global React, CASE, ImagePlaceholder, FooterStrip, NavStrip, CaseEyebrow, MetaBlock */

/* Image-led cut — warmer, mockup-forward.
 * Parchment hero with a carded screen sitting alongside the title (no full-bleed photo),
 * image grid (3-up) early, brief follows imagery, pull quote on Obsidian with image flank,
 * featured device-style mockup, dark CTA. */

function ImageLedCase({ flintGeometry = false }) {
  const chipClass = flintGeometry ? 'fc-chip-tr' : '';
  const featureChip = flintGeometry ? 'fc-chip-tr-bl' : '';

  return (
    <div style={{ background: '#faf9f7', fontFamily: 'var(--fc-font-body)', color: '#4a4038' }}>
      <NavStrip tone="light" />

      {/* ── Hero — Parchment, two-column: type left, real video right ── */}
      <section style={{ background: '#f0ece5', padding: '88px 64px 72px', borderBottom: '0.5px solid #d4cfc9' }}>
        <div style={{
          maxWidth: 1100, margin: '0 auto',
          display: 'grid', gridTemplateColumns: '1fr 1.05fr', gap: 56, alignItems: 'center',
        }}>
          <div>
            <div style={{ marginBottom: 22 }}>
              <div style={{ width: 28, height: 1, background: '#a84812', marginBottom: 10 }} />
              <span style={{
                fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500,
                letterSpacing: '0.2em', textTransform: 'uppercase', color: '#a84812',
              }}>Selected work · Case 04</span>
            </div>
            <h1 style={{
              margin: 0, fontFamily: 'var(--fc-font-display)',
              fontSize: 'clamp(40px, 4.6vw, 60px)', fontWeight: 400,
              lineHeight: 1.06, letterSpacing: '-0.018em',
              color: '#1c1e24', textWrap: 'pretty',
            }}>
              Rockabilly Roasting Co.
            </h1>
            <p style={{
              margin: '20px 0 0',
              fontFamily: 'var(--fc-font-display)', fontStyle: 'italic',
              fontSize: 22, lineHeight: 1.4, color: '#a84812',
            }}>
              A site that pours like the cup.
            </p>
            <p style={{
              margin: '28px 0 0', maxWidth: 460,
              fontFamily: 'var(--fc-font-body)', fontSize: 17,
              lineHeight: 1.7, color: '#4a4038',
            }}>
              Small-batch coffee, roasted in Kennewick. The cafe was already full; the website needed to catch up.
            </p>
          </div>
          <div className={chipClass} style={{
            border: flintGeometry ? '2px solid #a84812' : '0.5px solid #d4cfc9',
            borderRadius: flintGeometry ? 0 : 6,
            overflow: 'hidden', background: '#1c1e24', position: 'relative',
            aspectRatio: '4 / 3',
          }}>
            <video
              src="rockabilly-hero.mp4"
              autoPlay muted loop playsInline
              style={{ width: '100%', height: '100%', objectFit: 'cover', display: 'block' }}
            />
            <div style={{
              position: 'absolute', top: 14, left: 14,
              width: 28, height: 1, background: '#e8873b',
            }} />
            <div style={{
              position: 'absolute', bottom: 14, right: 14,
              fontFamily: 'SF Mono, ui-monospace, Menlo, monospace',
              fontSize: 10, letterSpacing: '0.18em', textTransform: 'uppercase',
              color: 'rgba(250, 249, 247, 0.6)',
              padding: '4px 8px', background: 'rgba(28, 30, 36, 0.5)',
              borderRadius: 2,
            }}>
              Live · home page hero loop
            </div>
          </div>
        </div>
      </section>

      {/* ── Meta strip — Birchwood ── */}
      <section style={{ background: '#faf9f7', borderBottom: '0.5px solid #d4cfc9' }}>
        <div style={{ maxWidth: 1100, margin: '0 auto' }}>
          <MetaBlock />
        </div>
      </section>

      {/* ── Three-up image grid — first thing after the hero, before the brief ── */}
      <section style={{ background: '#faf9f7', padding: '88px 64px 32px' }}>
        <div style={{ maxWidth: 1100, margin: '0 auto' }}>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'baseline', marginBottom: 32, gap: 24, flexWrap: 'wrap' }}>
            <CaseEyebrow>The work, top to bottom</CaseEyebrow>
            <span style={{ fontFamily: 'var(--fc-font-body)', fontSize: 13, color: '#6b5f52' }}>
              01 — 03 of 06
            </span>
          </div>
          <div style={{ display: 'grid', gridTemplateColumns: 'repeat(3, 1fr)', gap: 22 }}>
            {[
              { label: 'Home · hero, roast date, subscribe', tag: '01' },
              { label: 'Bean grid · single-origins, blends', tag: '02' },
              { label: 'Cafe · hours, address, what\'s on bar', tag: '03' },
            ].map((s) => (
              <div key={s.tag}>
                <div className={chipClass} style={{
                  border: '0.5px solid #d4cfc9', borderRadius: flintGeometry ? 0 : 6,
                  overflow: 'hidden', background: '#f0ece5',
                }}>
                  <ImagePlaceholder label={s.label} ratio="4 / 5" />
                </div>
                <div style={{ marginTop: 14, display: 'flex', gap: 10, alignItems: 'baseline' }}>
                  <span style={{
                    fontFamily: 'var(--fc-font-display)', fontStyle: 'italic',
                    fontSize: 13, color: '#a84812',
                  }}>№ {s.tag}</span>
                  <span style={{ fontFamily: 'var(--fc-font-body)', fontSize: 13, color: '#6b5f52' }}>{s.label}</span>
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* ── The brief — narrow column on Birchwood, image bookends ── */}
      <section style={{ background: '#faf9f7', padding: '80px 64px 112px' }}>
        <div style={{ maxWidth: 760, margin: '0 auto' }}>
          <CaseEyebrow>{CASE.brief.eyebrow}</CaseEyebrow>
          <h2 style={{
            margin: 0, fontFamily: 'var(--fc-font-display)',
            fontSize: 'clamp(30px, 3.2vw, 40px)', fontWeight: 400,
            color: '#1c1e24', lineHeight: 1.18, letterSpacing: '-0.01em',
            textWrap: 'pretty',
          }}>{CASE.brief.headline}</h2>
          <div style={{ marginTop: 32 }}>
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

      {/* ── Customer flow highlight — Flint, the subscription journey, mobile-anchored ── */}
      <section style={{ background: '#363b42', padding: '96px 64px' }}>
        <div style={{ maxWidth: 1100, margin: '0 auto' }}>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'baseline', marginBottom: 14, gap: 24, flexWrap: 'wrap' }}>
            <CaseEyebrow dark>The subscription flow</CaseEyebrow>
            <span style={{
              fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500,
              letterSpacing: '0.2em', textTransform: 'uppercase', color: '#96897a',
            }}>04 — Featured</span>
          </div>
          <h2 style={{
            margin: '0 0 14px', fontFamily: 'var(--fc-font-display)',
            fontSize: 'clamp(28px, 3.2vw, 38px)', fontWeight: 400,
            color: '#faf9f7', lineHeight: 1.2, letterSpacing: '-0.01em',
            maxWidth: 720, textWrap: 'pretty',
          }}>
            From <em style={{ fontStyle: 'italic', color: '#e8873b' }}>walked past the cafe</em> to <em style={{ fontStyle: 'italic', color: '#e8873b' }}>standing order</em> in five steps.
          </h2>
          <p style={{
            margin: '0 0 56px', fontFamily: 'var(--fc-font-body)', fontSize: 16,
            color: '#c4bdb2', lineHeight: 1.7, maxWidth: 640,
          }}>
            The new site treats subscribing like the obvious next step, not a destination. Each screen leans on the one before it — no portal login, no account wall before checkout, no surprises at the billing summary.
          </p>

          <div style={{ display: 'grid', gridTemplateColumns: '300px 1fr', gap: 56, alignItems: 'start' }}>
            {/* Anchor — real mobile screenshot inside a phone-ish frame */}
            <div style={{ position: 'sticky', top: 24 }}>
              <div className={chipClass} style={{
                background: '#1c1e24',
                borderRadius: flintGeometry ? 0 : 28,
                padding: 8,
                border: '0.5px solid #474d55',
              }}>
                <img
                  src="rockabilly-mobile.webp"
                  alt="Rockabilly Roasting home page on mobile"
                  style={{
                    display: 'block', width: '100%', height: 'auto',
                    borderRadius: flintGeometry ? 0 : 22,
                  }}
                />
              </div>
              <p style={{
                margin: '14px 4px 0', fontFamily: 'var(--fc-font-body)',
                fontSize: 12, color: '#96897a', lineHeight: 1.5,
              }}>
                Step 01 · Home, on a phone — what most first-time visitors actually see.
              </p>
            </div>

            {/* Steps — connected vertical flow */}
            <ol style={{ listStyle: 'none', margin: 0, padding: 0, position: 'relative' }}>
              {/* connector */}
              <div style={{
                position: 'absolute', left: 11, top: 14, bottom: 14,
                width: 1, background: '#474d55',
              }} />
              {[
                {
                  n: '01', tag: 'Land',
                  h: 'Burn both ends. Roast the beans.',
                  b: 'Looping cafe footage behind the title. Two CTAs only — shop this week\'s roast, visit the cafe.',
                },
                {
                  n: '02', tag: 'Browse',
                  h: 'This week\'s roast, freshly stamped.',
                  b: 'Roast date, origin and tasting notes sit above the price. No carousel, no filters menu.',
                },
                {
                  n: '03', tag: 'Pick',
                  h: 'Bag size, grind, cadence.',
                  b: 'Three picks on one screen. No portal account required to see what it costs.',
                },
                {
                  n: '04', tag: 'Confirm',
                  h: 'A billing summary you can read in one breath.',
                  b: 'Total per shipment, total per month, when the first bag arrives. Nothing more.',
                },
                {
                  n: '05', tag: 'Recur',
                  h: 'A bag a week, until you say otherwise.',
                  b: 'Manage from a one-click email link. No login, no recovered-password loop.',
                },
              ].map((s, i, arr) => (
                <li key={s.n} style={{
                  position: 'relative',
                  paddingLeft: 44,
                  paddingBottom: i === arr.length - 1 ? 0 : 36,
                }}>
                  <span style={{
                    position: 'absolute', left: 0, top: 4,
                    width: 22, height: 22, borderRadius: '50%',
                    background: '#e8873b', color: '#1c1e24',
                    display: 'flex', alignItems: 'center', justifyContent: 'center',
                    fontFamily: 'var(--fc-font-body)', fontSize: 10, fontWeight: 500,
                    letterSpacing: '0.04em',
                  }}>{s.n}</span>
                  <div style={{
                    fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500,
                    letterSpacing: '0.22em', textTransform: 'uppercase',
                    color: '#e8873b', marginBottom: 8,
                  }}>{s.tag}</div>
                  <h3 style={{
                    margin: 0, fontFamily: 'var(--fc-font-display)',
                    fontSize: 22, fontWeight: 400, color: '#faf9f7',
                    lineHeight: 1.3,
                  }}>{s.h}</h3>
                  <p style={{
                    margin: '8px 0 0', fontFamily: 'var(--fc-font-body)',
                    fontSize: 15, color: '#c4bdb2', lineHeight: 1.65,
                    maxWidth: 520,
                  }}>{s.b}</p>
                </li>
              ))}
            </ol>
          </div>
        </div>
      </section>

      {/* ── Pull quote — Obsidian, image flanks, Cormorant italic ── */}
      <section style={{ background: '#1c1e24', padding: '120px 64px' }}>
        <div style={{
          maxWidth: 1100, margin: '0 auto',
          display: 'grid', gridTemplateColumns: '1fr 1.4fr', gap: 64, alignItems: 'center',
        }}>
          <div className={chipClass} style={{
            border: '0.5px solid #474d55', borderRadius: flintGeometry ? 0 : 6,
            overflow: 'hidden', background: '#363b42',
          }}>
            <ImagePlaceholder label={'Portrait · Kagen Cox\nbehind the bar'} ratio="3 / 4" dark accent />
          </div>
          <div>
            <div style={{ width: 28, height: 1, background: '#e8873b', marginBottom: 24 }} />
            <p style={{
              margin: 0,
              fontFamily: 'var(--fc-font-display)', fontStyle: 'italic',
              fontSize: 'clamp(24px, 2.6vw, 32px)',
              lineHeight: 1.4, color: '#faf9f7', textWrap: 'pretty',
            }}>
              “{CASE.quote.text}”
            </p>
            <div style={{ marginTop: 32, display: 'flex', alignItems: 'center', gap: 14, flexWrap: 'wrap' }}>
              <span style={{ fontSize: 14, color: '#e8873b', fontFamily: 'var(--fc-font-body)' }}>★★★★★</span>
              <span style={{
                fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500,
                letterSpacing: '0.2em', textTransform: 'uppercase', color: '#96897a',
              }}>
                {CASE.quote.attribution} · {CASE.quote.role}
              </span>
            </div>
          </div>
        </div>
      </section>

      {/* ── Two more screens, this time on Parchment ── */}
      <section style={{ background: '#f0ece5', padding: '96px 64px', borderTop: '0.5px solid #d4cfc9' }}>
        <div style={{ maxWidth: 1100, margin: '0 auto' }}>
          <CaseEyebrow>More from the build</CaseEyebrow>
          <div style={{ display: 'grid', gridTemplateColumns: '1.2fr 1fr', gap: 22, marginTop: 24 }}>
            <div className={chipClass} style={{
              border: '0.5px solid #d4cfc9', borderRadius: flintGeometry ? 0 : 6,
              overflow: 'hidden', background: '#faf9f7',
            }}>
              <ImagePlaceholder label="05 · Wholesale form — three fields, plain English" ratio="5 / 4" />
            </div>
            <div className={chipClass} style={{
              border: '0.5px solid #d4cfc9', borderRadius: flintGeometry ? 0 : 6,
              overflow: 'hidden', background: '#faf9f7',
            }}>
              <ImagePlaceholder label="06 · Mobile menu — bar, beans, hours" ratio="5 / 4" />
            </div>
          </div>
        </div>
      </section>

      {/* ── Closing CTA on Obsidian ── */}
      <section style={{ background: '#1c1e24', padding: '112px 64px' }}>
        <div style={{ maxWidth: 720, margin: '0 auto' }}>
          <CaseEyebrow dark>Next</CaseEyebrow>
          <h2 style={{
            margin: 0, fontFamily: 'var(--fc-font-display)',
            fontSize: 'clamp(32px, 3.4vw, 44px)', fontWeight: 400,
            lineHeight: 1.15, color: '#faf9f7', letterSpacing: '-0.01em',
          }}>
            Outgrown your first site?
          </h2>
          <p style={{
            margin: '20px 0 32px', fontFamily: 'var(--fc-font-body)',
            fontSize: 17, color: '#c4bdb2', lineHeight: 1.7, maxWidth: 540,
          }}>
            The first conversation costs nothing. No pitch, no pressure. Just an honest read on whether it makes sense.
          </p>
          <div style={{ display: 'flex', gap: 20, alignItems: 'center', flexWrap: 'wrap' }}>
            <a className={`fc-flint-btn ${flintGeometry ? 'fc-chip-btn' : ''}`} style={{
              background: '#e8873b', color: '#1c1e24',
              padding: '14px 26px', fontSize: 13,
              borderRadius: flintGeometry ? 0 : 3,
            }}>Start a project</a>
            <a style={{
              fontFamily: 'var(--fc-font-body)', fontSize: 13, fontWeight: 500,
              letterSpacing: '0.08em', textTransform: 'uppercase',
              color: '#e8873b',
              borderBottom: '1.5px solid #e8873b', paddingBottom: 4,
            }}>See more work →</a>
          </div>
        </div>
      </section>

      <FooterStrip tone="light" />
    </div>
  );
}

window.ImageLedCase = ImageLedCase;
