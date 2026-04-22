/* global React, Eyebrow, Button, CapabilityCard, ProjectCard, FeaturedProjectCard, ContactForm, PullQuote */
const { useState: useSt } = React;

const PROJECTS = [
  { slug: 'nautilus', client: 'Nautilus Group', location: 'Kennewick, WA', tags: ['Brand', 'Marketing site'],
    title: 'A clean-services site that reads like a referral.',
    quote: { text: 'He treated us like family. Patience and understanding with all the changes, and someone willing to listen to your needs.', attribution: 'David Nash' },
    results: [{ value: '3×', unit: '', label: 'Inbound leads' }, { value: '28', unit: 'd', label: 'To launch' }] },
  { slug: 'lo-mo', client: 'Lo Mo Outfitting', location: 'Helena, MT', tags: ['Brand', 'Marketing site'], title: 'A Montana outfitter whose website matches the guide.', results: [{ value: '100', unit: '%', label: 'Custom build' }, { value: '6', unit: 'wk', label: 'Total build' }] },
  { slug: 'schluters', client: "Schluter's Metal Art", location: 'Helena, MT', tags: ['Marketing site'], title: 'Steel sculpture work, presented like steel sculpture.', results: [{ value: '12', unit: '', label: 'Commissions / mo' }] },
  { slug: 'k9', client: 'K9 Elements', location: 'Helena, MT', tags: ['Brand', 'Services'], title: 'Dog training services, written in plain English.', results: [{ value: '2×', unit: '', label: 'Enrollment' }] },
];

function HomeScreen({ onNav }) {
  return (
    <div>
      {/* Hero */}
      <section style={{ background: '#1c1e24', position: 'relative', overflow: 'hidden' }}>
        <div style={{ maxWidth: 1100, margin: '0 auto', padding: '120px 32px', display: 'grid', gridTemplateColumns: 'minmax(0, 1.35fr) minmax(0, 1fr)', gap: 72, alignItems: 'center' }}>
          <div style={{ minWidth: 0 }}>
            <Eyebrow dark>Custom web development · Helena, MT</Eyebrow>
            <h1 style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontSize: 'clamp(44px, 5.4vw, 72px)', fontWeight: 400, lineHeight: 1.0, letterSpacing: '-0.02em', color: '#f0ece5' }}>
              <em style={{ fontStyle: 'italic' }}>Carved for purpose.</em>
            </h1>
            <p style={{ margin: '36px 0 0', fontFamily: 'var(--fc-font-body)', fontSize: 18, lineHeight: 1.7, color: '#c4bdb2', maxWidth: 480 }}>
              One relationship. <em style={{ fontFamily: 'var(--fc-font-display)', fontStyle: 'italic' }}>Everything handled.</em> For businesses that have earned their reputation and are ready for a website that reflects it.
            </p>
            <div style={{ marginTop: 40, display: 'flex', gap: 20, flexWrap: 'wrap' }}>
              <Button variant="primary" dark onClick={() => onNav('contact')}>Start a project</Button>
              <Button variant="ghost" dark onClick={() => onNav('work')}>See our work →</Button>
            </div>
          </div>
          <div style={{ justifySelf: 'end', minWidth: 0 }}>
            <img src="../../assets/logo.svg" alt="" style={{ width: '100%', maxWidth: 320, height: 'auto', opacity: 0.45, display: 'block' }} />
          </div>
        </div>
      </section>

      {/* Stats */}
      <section style={{ background: '#f0ece5', borderTop: '0.5px solid #d4cfc9', borderBottom: '0.5px solid #d4cfc9' }}>
        <div style={{ maxWidth: 1100, margin: '0 auto', display: 'grid', gridTemplateColumns: 'repeat(4, 1fr)' }}>
          {[['100%', 'Deadlines met'], ['2', 'Engineers — you always know who built it'], ['0', 'Templates — every line written for you'], ['Same day', 'Support response']].map(([v, l], i) => (
            <div key={i} style={{ padding: '40px 22px', textAlign: 'center', borderRight: i < 3 ? '0.5px solid #d4cfc9' : 'none' }}>
              <p style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontSize: 26, color: '#1c1e24' }}>{v}</p>
              <p style={{ margin: '6px 0 0', fontFamily: 'var(--fc-font-body)', fontSize: 12, color: '#6b5f52', lineHeight: 1.5 }}>{l}</p>
            </div>
          ))}
        </div>
      </section>

      {/* Who we are */}
      <section style={{ background: '#faf9f7', padding: '100px 32px' }}>
        <div style={{ maxWidth: 1100, margin: '0 auto' }}>
          <div style={{ maxWidth: 640 }}>
            <Eyebrow>Who we are</Eyebrow>
            <h2 style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontSize: 40, fontWeight: 400, color: '#1c1e24', lineHeight: 1.15 }}>
              Small by design. Accountable by nature.
            </h2>
            <div style={{ marginTop: 28 }}>
              <p style={{ fontFamily: 'var(--fc-font-body)', fontSize: 17, color: '#4a4038', lineHeight: 1.7 }}>
                FlintCraft Studio is a small custom development studio in Helena, Montana. We build websites for businesses that have earned their reputation and are ready for a web presence that reflects it.
              </p>
              <p style={{ fontFamily: 'var(--fc-font-body)', fontSize: 17, color: '#4a4038', lineHeight: 1.7 }}>
                Everything we ship is written from scratch. No WordPress. No plugins. No templates. No page builders.
              </p>
            </div>
            <div style={{ marginTop: 26 }}>
              <PullQuote>"As durable as the mountains, and just as unlikely to move to San Francisco."</PullQuote>
              <p style={{ margin: '6px 0 0', fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500, letterSpacing: '0.18em', textTransform: 'uppercase', color: '#6b5f52' }}>
                Helena, MT · Est. 2019
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* ICP Statement */}
      <section style={{ background: '#1c1e24', padding: '100px 32px' }}>
        <div style={{ maxWidth: 760, margin: '0 auto' }}>
          <Eyebrow dark>Who we work with</Eyebrow>
          <h2 style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontSize: 40, fontWeight: 400, color: '#f0ece5', lineHeight: 1.15 }}>
            Done improvising. Ready to build something that lasts.
          </h2>
          <p style={{ marginTop: 28, fontFamily: 'var(--fc-font-body)', fontSize: 17, color: '#c4bdb2', lineHeight: 1.7, maxWidth: 640 }}>
            You've done the hard part. You built something real, earned real customers, and proven the business works. Your website is the one thing you never got right — or the one thing you built when you were still figuring everything out.
          </p>
        </div>
      </section>

      {/* CTA */}
      <section style={{ background: '#363b42', padding: '96px 32px', textAlign: 'center' }}>
        <div style={{ maxWidth: 720, margin: '0 auto' }}>
          <h2 style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontSize: 40, fontWeight: 400, color: '#f0ece5', lineHeight: 1.15 }}>
            Ready to find out if we're the right fit?
          </h2>
          <p style={{ marginTop: 20, fontFamily: 'var(--fc-font-body)', fontSize: 17, color: '#c4bdb2', lineHeight: 1.7 }}>
            The first conversation costs nothing. No pitch, no pressure. Just an honest read on whether it makes sense.
          </p>
          <div style={{ marginTop: 36, display: 'flex', gap: 24, justifyContent: 'center', flexWrap: 'wrap' }}>
            <Button variant="primary" dark onClick={() => onNav('contact')}>Start a project</Button>
            <Button variant="ghost" dark>(406) 871-9875</Button>
          </div>
        </div>
      </section>
    </div>
  );
}

function ServicesScreen() {
  const CAPS = [
    ['A site that reflects where you actually are.', 'Not where you were when you started. Custom designed, custom built, no templates.'],
    ['Copy that does the selling.', "We write it. You verify it. Most owners don't have time to stare at a blank page."],
    ['Built to get found.', 'SEO foundations, Google Business Profile setup and optimization, and copy written for the markets you actually serve.'],
    ['We watch what\'s working and adjust.', 'After launch we track what\'s converting and tailor your content around it.'],
    ['Hosting, maintenance, and updates — handled.', 'Your site lives on our infrastructure. When something needs updating, you call us directly.'],
    ['One person to call. Always.', "You'll know who built your site and you'll have their number. That doesn't change after launch."],
  ];
  return (
    <div>
      <section style={{ background: '#faf9f7', padding: '96px 32px' }}>
        <div style={{ maxWidth: 1100, margin: '0 auto' }}>
          <Eyebrow>Services</Eyebrow>
          <h1 style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontSize: 52, fontWeight: 400, color: '#1c1e24', lineHeight: 1.1, maxWidth: 720 }}>
            One relationship. <em style={{ fontStyle: 'italic', color: '#a84812' }}>Everything handled.</em>
          </h1>
          <p style={{ marginTop: 24, fontFamily: 'var(--fc-font-body)', fontSize: 17, color: '#4a4038', maxWidth: 560, lineHeight: 1.7 }}>
            For businesses that have earned their reputation and are ready for a website that reflects it.
          </p>
        </div>
      </section>
      <section style={{ background: '#1c1e24', padding: '96px 32px' }}>
        <div style={{ maxWidth: 1100, margin: '0 auto' }}>
          <Eyebrow dark>What we do</Eyebrow>
          <div style={{ marginTop: 32, display: 'grid', gridTemplateColumns: '1fr 1fr', gap: 32 }}>
            {CAPS.map(([h, b], i) => <CapabilityCard key={h} headline={h} body={b} number={String(i+1).padStart(2, '0')} />)}
          </div>
        </div>
      </section>
      <section style={{ background: '#faf9f7', padding: '96px 32px' }}>
        <div style={{ maxWidth: 720, margin: '0 auto' }}>
          <Eyebrow>What it costs</Eyebrow>
          <h2 style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontSize: 36, fontWeight: 400, color: '#1c1e24', lineHeight: 1.2 }}>Straightforward pricing. No surprises.</h2>
          <p style={{ marginTop: 24, fontFamily: 'var(--fc-font-body)', fontSize: 17, color: '#4a4038', lineHeight: 1.7 }}>
            Most FlintCraft engagements run between $150 and $350 per month, all in. No setup fees, no surprise invoices, no separate hosting bill. One rate covers design, build, hosting, maintenance, updates, and ongoing visibility work.
          </p>
        </div>
      </section>
    </div>
  );
}

/* ====== EDITORIAL WORK SCREEN ====== */

function FieldReport({ index, project, report, dark }) {
  const bg = dark ? '#1c1e24' : '#faf9f7';
  const ink = dark ? '#f0ece5' : '#1c1e24';
  const muted = dark ? '#e8e3dc' : '#4a4038';
  const bodyInk = dark ? '#faf9f7' : '#1c1e24';
  const sub = dark ? '#96897a' : '#6b5f52';
  const accent = dark ? '#e8873b' : '#a84812';
  return (
    <article style={{ background: bg, padding: '96px 32px', borderBottom: `0.5px solid ${dark ? '#363b42' : '#d4cfc9'}` }}>
      <div style={{ maxWidth: 1100, margin: '0 auto', display: 'grid', gridTemplateColumns: '240px 1fr', gap: 56 }}>
        {/* Marginalia */}
        <aside>
          <div style={{ fontFamily: 'var(--fc-font-display)', fontStyle: 'italic', fontSize: 48, color: accent, lineHeight: 1 }}>№{String(index).padStart(2, '0')}</div>
          <div style={{ marginTop: 28, fontFamily: 'var(--fc-font-body)', fontSize: 10, fontWeight: 500, letterSpacing: '0.2em', textTransform: 'uppercase', color: sub, lineHeight: 2.2 }}>
            <div style={{ color: accent }}>Client</div>
            <div style={{ color: ink, letterSpacing: '0.08em', fontSize: 13, fontWeight: 500, textTransform: 'none', fontFamily: 'var(--fc-font-display)', fontStyle: 'italic' }}>{project.client}</div>
            <div style={{ marginTop: 14, color: accent }}>Location</div>
            <div style={{ color: muted, letterSpacing: '0.06em', fontSize: 12, textTransform: 'none' }}>{project.location}</div>
            <div style={{ marginTop: 14, color: accent }}>Scope</div>
            <div style={{ color: muted, letterSpacing: '0.06em', fontSize: 12, textTransform: 'none' }}>{project.tags.join(', ')}</div>
            <div style={{ marginTop: 14, color: accent }}>Shipped</div>
            <div style={{ color: muted, letterSpacing: '0.06em', fontSize: 12, textTransform: 'none' }}>{report.shipped}</div>
          </div>
        </aside>

        {/* Column */}
        <div>
          <div style={{ fontFamily: 'var(--fc-font-body)', fontSize: 10, fontWeight: 500, letterSpacing: '0.28em', textTransform: 'uppercase', color: accent, marginBottom: 10 }}>
            Field Report
          </div>
          <h2 style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontSize: 44, fontWeight: 400, color: ink, lineHeight: 1.1, letterSpacing: '-0.01em' }}>
            {project.title}
          </h2>

          {/* Image plate */}
          <figure style={{ margin: '44px 0', position: 'relative' }}>
            <div className="fc-chip-tr-lg" style={{ background: dark ? '#363b42' : '#f0ece5', height: 340, display: 'flex', alignItems: 'center', justifyContent: 'center', position: 'relative' }}>
              <img src="../../assets/logo.svg" alt="" style={{ height: 80, width: 'auto', opacity: dark ? 0.16 : 0.22 }} />
              <div style={{ position: 'absolute', top: 0, right: 0, width: 44, height: 2.5, background: accent, transformOrigin: '100% 50%', transform: 'translate(-20px, 13px) rotate(-45deg)' }} />
            </div>
            <figcaption style={{ marginTop: 12, fontFamily: 'var(--fc-font-display)', fontStyle: 'italic', fontSize: 13, color: sub, letterSpacing: '0.02em' }}>
              Plate {index}.1 — {report.caption}
            </figcaption>
          </figure>

          {/* Body — proper editorial column */}
          <div style={{ fontFamily: 'var(--fc-font-body)', fontSize: 17, lineHeight: 1.75, color: bodyInk, maxWidth: 620 }}>
            <p style={{ marginTop: 0, color: bodyInk }}>
              <span style={{ fontFamily: 'var(--fc-font-display)', fontSize: 11, fontWeight: 500, letterSpacing: '0.2em', textTransform: 'uppercase', color: accent, marginRight: 10 }}>The brief —</span>
              {report.brief}
            </p>
            <p style={{ color: bodyInk }}>
              <span style={{ fontFamily: 'var(--fc-font-display)', fontSize: 11, fontWeight: 500, letterSpacing: '0.2em', textTransform: 'uppercase', color: accent, marginRight: 10 }}>What we did —</span>
              {report.work}
            </p>
            <p style={{ color: bodyInk }}>
              <span style={{ fontFamily: 'var(--fc-font-display)', fontSize: 11, fontWeight: 500, letterSpacing: '0.2em', textTransform: 'uppercase', color: accent, marginRight: 10 }}>The outcome —</span>
              {report.outcome}
            </p>
          </div>

          {/* Pull quote, if present */}
          {project.quote && (
            <blockquote style={{ margin: '44px 0 0', paddingLeft: 24, borderLeft: `2.5px solid ${accent}` }}>
              <p style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontStyle: 'italic', fontSize: 22, lineHeight: 1.45, color: ink, letterSpacing: '-0.005em' }}>
                "{project.quote.text}"
              </p>
              <footer style={{ marginTop: 12, fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500, letterSpacing: '0.2em', textTransform: 'uppercase', color: sub }}>
                — {project.quote.attribution}
              </footer>
            </blockquote>
          )}

          {/* Results strip */}
          <div style={{ marginTop: 44, paddingTop: 28, borderTop: `0.5px solid ${dark ? '#363b42' : '#d4cfc9'}`, display: 'flex', gap: 48, flexWrap: 'wrap' }}>
            {project.results.map((r, i) => (
              <div key={i}>
                <div style={{ fontFamily: 'var(--fc-font-display)', fontSize: 36, fontWeight: 400, color: ink, letterSpacing: '-0.02em' }}>
                  {r.value}<span style={{ fontSize: 16, color: accent, marginLeft: 2 }}>{r.unit}</span>
                </div>
                <div style={{ fontFamily: 'var(--fc-font-body)', fontSize: 10, letterSpacing: '0.2em', textTransform: 'uppercase', color: sub, marginTop: 6 }}>{r.label}</div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </article>
  );
}

function WorkScreen({ onNav }) {
  const reports = [
    { shipped: 'March 2026', caption: 'Home hero and service ledger, as shipped.',
      brief: "Nautilus had outgrown a templated Squarespace site. Their phones were ringing on referrals alone — the website was an afterthought, and it showed. They wanted something that matched the quality of the work David and his crew had been putting in for fifteen years.",
      work: 'A custom build from scratch. We interviewed David, shadowed two job walkthroughs, and wrote the copy ourselves before touching design. The information architecture runs three verticals — commercial, residential, specialty — each with its own entry point. No generic service blocks. Every photograph is their crew, their trucks, their job sites.',
      outcome: "The site converts on first read. Inbound lead volume tripled in the first eight weeks, and the inquiries that come through are already qualified — they've read the services page and know which vertical they want." },
    { shipped: 'January 2026', caption: 'Landing page and booking flow.',
      brief: "Lo Mo's hunting and fishing outfits had a reputation built over two decades of Montana backcountry. Their old site looked like it was built on a card table in 2009. Trips were booking by phone, by word of mouth, by sheer force of reputation.",
      work: "Two on-location days in the Bob Marshall Wilderness. We came back with the photography, the stories, and a feel for the voice the site needed. The build itself is tight: five pages, one booking flow, a straightforward rates table. The trip pages read like a guide briefing — practical, specific, honest about what clients can expect.",
      outcome: "Guide-to-guide trip handoffs moved from paper to a simple internal dashboard. Online inquiries are up meaningfully, but the bigger shift is that the site now does the qualifying work before the first call." },
    { shipped: 'November 2025', caption: 'Commission gallery and inquiry form.',
      brief: "Mike Schluter's steel sculpture is commissioned work — ranches, civic plazas, private collections. Photos were sitting in an unsorted Dropbox. There was no website at all, just an Instagram feed and word of mouth.",
      work: "A quiet portfolio site. Full-bleed photography, minimal chrome, a clear path from piece to inquiry. We built a simple commissions ledger so Mike can mark pieces as available, sold, or on commission, and we wrote the inquiry form to pre-qualify scope and timeline.",
      outcome: "Roughly a dozen serious inquiries a month, most of which convert to commissions. The site stays out of the way of the work." },
  ];

  return (
    <div>
      {/* Masthead */}
      <header style={{ background: '#faf9f7', borderBottom: '0.5px solid #d4cfc9', padding: '72px 32px 48px' }}>
        <div style={{ maxWidth: 1100, margin: '0 auto', display: 'grid', gridTemplateColumns: '240px 1fr', gap: 56, alignItems: 'end' }}>
          <div style={{ fontFamily: 'var(--fc-font-body)', fontSize: 10, fontWeight: 500, letterSpacing: '0.28em', textTransform: 'uppercase', color: '#6b5f52', lineHeight: 1.8 }}>
            Vol. 7<br/>Spring 2026<br/><span style={{ color: '#a84812' }}>Helena, Montana</span>
          </div>
          <div>
            <div style={{ fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500, letterSpacing: '0.28em', textTransform: 'uppercase', color: '#a84812', marginBottom: 14 }}>
              Selected Work — A Journal
            </div>
            <h1 style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontSize: 88, fontWeight: 400, color: '#1c1e24', lineHeight: 0.95, letterSpacing: '-0.03em' }}>
              <em style={{ fontStyle: 'italic' }}>Field Reports.</em>
            </h1>
            <p style={{ margin: '24px 0 0', fontFamily: 'var(--fc-font-body)', fontSize: 17, color: '#4a4038', lineHeight: 1.7, maxWidth: 560 }}>
              A running record of recent projects. Real businesses, real problems, real outcomes — written the way we wish case studies were always written.
            </p>
          </div>
        </div>
      </header>

      {/* Ledger (table of contents) */}
      <section style={{ background: '#f0ece5', padding: '48px 32px', borderBottom: '0.5px solid #d4cfc9' }}>
        <div style={{ maxWidth: 1100, margin: '0 auto' }}>
          <div style={{ fontFamily: 'var(--fc-font-body)', fontSize: 10, fontWeight: 500, letterSpacing: '0.28em', textTransform: 'uppercase', color: '#6b5f52', marginBottom: 20 }}>
            In this volume
          </div>
          <div style={{ display: 'grid', gap: 2 }}>
            {PROJECTS.slice(0, 3).map((p, i) => (
              <div key={p.slug} style={{ display: 'grid', gridTemplateColumns: '40px 1fr auto', gap: 20, alignItems: 'baseline', padding: '10px 0', borderTop: i === 0 ? '0.5px solid #d4cfc9' : '0.5px solid #d4cfc9' }}>
                <span style={{ fontFamily: 'var(--fc-font-display)', fontStyle: 'italic', fontSize: 14, color: '#a84812' }}>№{String(i+1).padStart(2, '0')}</span>
                <span style={{ fontFamily: 'var(--fc-font-display)', fontSize: 19, color: '#1c1e24' }}>
                  <em style={{ fontStyle: 'italic', color: '#a84812' }}>{p.client}</em> — {p.title.replace(/\.$/, '')}
                </span>
                <span style={{ fontFamily: 'var(--fc-font-body)', fontSize: 10, letterSpacing: '0.2em', textTransform: 'uppercase', color: '#6b5f52' }}>p. {(i+1)*4}</span>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Field reports, alternating surfaces for editorial rhythm */}
      {reports.map((r, i) => (
        <FieldReport key={i} index={i+1} project={PROJECTS[i]} report={r} dark={i % 2 === 1} />
      ))}

      {/* Colophon / CTA */}
      <section style={{ background: '#faf9f7', padding: '96px 32px', textAlign: 'center' }}>
        <div style={{ maxWidth: 620, margin: '0 auto' }}>
          <div style={{ fontFamily: 'var(--fc-font-body)', fontSize: 10, fontWeight: 500, letterSpacing: '0.28em', textTransform: 'uppercase', color: '#a84812', marginBottom: 14 }}>Colophon</div>
          <p style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontStyle: 'italic', fontSize: 26, color: '#1c1e24', lineHeight: 1.4, letterSpacing: '-0.005em' }}>
            Each report is a faithful account of work done in Montana and points west. If you see your business in one of these, the next one could be yours.
          </p>
          <div style={{ marginTop: 36 }}>
            <Button variant="primary" onClick={() => onNav('contact')}>Start a project</Button>
          </div>
        </div>
      </section>
    </div>
  );
}

function ContactScreen() {
  return (
    <section style={{ background: '#faf9f7', padding: '96px 32px' }}>
      <div style={{ maxWidth: 1100, margin: '0 auto', display: 'grid', gridTemplateColumns: '1fr 1fr', gap: 80 }}>
        <div>
          <Eyebrow>Start a project</Eyebrow>
          <h1 style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontSize: 44, fontWeight: 400, color: '#1c1e24', lineHeight: 1.15 }}>
            <em style={{ fontStyle: 'italic' }}>Let's find out if we're the right fit.</em>
          </h1>
          <p style={{ marginTop: 22, fontFamily: 'var(--fc-font-body)', fontSize: 17, color: '#4a4038', lineHeight: 1.7, maxWidth: 400 }}>
            No pitch, no deck. Tell us what you're working on — we'll take a look at your current setup and come to the first conversation prepared.
          </p>
          <div style={{ marginTop: 36, fontFamily: 'var(--fc-font-body)', fontSize: 15, color: '#6b5f52', lineHeight: 1.9 }}>
            <p style={{ margin: 0 }}>Prefer to call? <a style={{ color: '#a84812' }}>(406) 871-9875</a></p>
            <p style={{ margin: 0 }}>Or email: <a style={{ color: '#a84812' }}>hello@flintcraftstudio.com</a></p>
            <p style={{ margin: 0 }}>We're based in Helena, Montana. We generally respond same business day.</p>
          </div>
        </div>
        <div><ContactForm /></div>
      </div>
    </section>
  );
}

Object.assign(window, { HomeScreen, ServicesScreen, WorkScreen, ContactScreen });
