/* global React */

/* Flint-geometry card — chipped top-right, struck-edge accent rule.
 * All cards carry .fc-flint-card + one of .fc-chip-tr / .fc-chip-tr-lg / .fc-chip-tr-bl.
 */

function CardLight({ eyebrow, title, body, cta }) {
  return (
    <div className="fc-flint-card fc-flint-card--light fc-chip-tr" style={{ border: '0.5px solid #D5D7DA' }}>
      <div className="fc-flint-card__strike" />
      {eyebrow && <div style={{ fontFamily: 'var(--fc-font-body)', fontSize: 10, fontWeight: 500, letterSpacing: '0.2em', textTransform: 'uppercase', color: '#666B72', marginBottom: 10 }}>{eyebrow}</div>}
      <div style={{ fontFamily: 'var(--fc-font-display)', fontSize: 22, fontWeight: 600, color: '#17181A', marginBottom: 10, lineHeight: 1.3 }}>{title}</div>
      <div style={{ fontFamily: 'var(--fc-font-body)', fontSize: 14, color: '#3B3E43', lineHeight: 1.65 }}>{body}</div>
      {cta && <a style={{ display: 'inline-block', marginTop: 18, fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500, letterSpacing: '0.08em', textTransform: 'uppercase', color: '#A85A14', textDecoration: 'none' }}>{cta} →</a>}
    </div>
  );
}

function CapabilityCard({ headline, body, number }) {
  return (
    <div className="fc-flint-card fc-flint-card--dark fc-chip-tr">
      <div className="fc-flint-card__strike" />
      {number && (
        <div style={{ fontFamily: 'var(--fc-font-display)', fontStyle: 'italic', fontSize: 14, color: '#EE9440', marginBottom: 10, letterSpacing: '0.04em' }}>
          № {number}
        </div>
      )}
      <h3 style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontSize: 24, fontWeight: 400, color: '#F5F5F3', lineHeight: 1.3, marginBottom: 12 }}>{headline}</h3>
      <p style={{ margin: 0, fontFamily: 'var(--fc-font-body)', fontSize: 15, color: '#D5D7DA', lineHeight: 1.65 }}>{body}</p>
    </div>
  );
}

function ProjectCard({ project, onClick }) {
  return (
    <a onClick={onClick}
      className="fc-flint-card--dark fc-chip-tr"
      style={{ cursor: 'pointer', display: 'flex', flexDirection: 'column', textDecoration: 'none', padding: 0, position: 'relative', background: '#3B3E43' }}>
      <div className="fc-flint-card__strike" />
      <div style={{ height: 180, background: '#17181A', display: 'flex', alignItems: 'center', justifyContent: 'center', borderBottom: '0.5px solid #3B3E43' }}>
        <img src="../../assets/logo.svg" alt="" style={{ height: 42, width: 'auto', opacity: 0.2 }} />
      </div>
      <div style={{ padding: 24, display: 'flex', flexDirection: 'column', flex: 1 }}>
        <div style={{ display: 'flex', gap: 12, marginBottom: 10, flexWrap: 'wrap' }}>
          <span style={{ fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500, letterSpacing: '0.18em', textTransform: 'uppercase', color: '#EE9440' }}>{project.client}</span>
          <span style={{ fontFamily: 'var(--fc-font-body)', fontSize: 11, color: '#D5D7DA' }}>{project.location}</span>
        </div>
        <h3 style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontSize: 22, fontWeight: 400, color: '#F5F5F3', lineHeight: 1.3, marginBottom: 16 }}>{project.title}</h3>
        {project.results && (
          <div style={{ marginTop: 'auto', paddingTop: 16, borderTop: '0.5px solid rgba(213,215,218, 0.2)', display: 'flex', gap: 24, flexWrap: 'wrap' }}>
            {project.results.map((r, i) => (
              <div key={i}>
                <div style={{ fontFamily: 'var(--fc-font-display)', fontSize: 22, fontWeight: 600, color: '#F5F5F3' }}>
                  {r.value}<span style={{ fontSize: 13, fontWeight: 400, color: '#EE9440', marginLeft: 2 }}>{r.unit}</span>
                </div>
                <div style={{ fontFamily: 'var(--fc-font-body)', fontSize: 10, letterSpacing: '0.18em', textTransform: 'uppercase', color: '#D5D7DA', marginTop: 4 }}>{r.label}</div>
              </div>
            ))}
          </div>
        )}
      </div>
    </a>
  );
}

function FeaturedProjectCard({ project }) {
  return (
    <div className="fc-flint-card--dark fc-chip-tr-bl"
         style={{ padding: 0, display: 'grid', gridTemplateColumns: '3fr 2fr', background: '#3B3E43', position: 'relative' }}>
      <div className="fc-flint-card__strike" style={{ width: 44, transform: 'translate(-20px, 12px) rotate(-45deg)' }} />
      <div style={{ background: '#17181A', display: 'flex', alignItems: 'center', justifyContent: 'center', minHeight: 340 }}>
        <img src="../../assets/logo.svg" alt="" style={{ height: 72, width: 'auto', opacity: 0.18 }} />
      </div>
      <div style={{ padding: 40, display: 'flex', flexDirection: 'column', justifyContent: 'space-between' }}>
        <div>
          <div style={{ display: 'flex', gap: 12, marginBottom: 14, flexWrap: 'wrap' }}>
            <span style={{ fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500, letterSpacing: '0.18em', textTransform: 'uppercase', color: '#EE9440' }}>{project.client}</span>
            <span style={{ fontFamily: 'var(--fc-font-body)', fontSize: 11, color: '#D5D7DA' }}>{project.location}</span>
          </div>
          <h3 style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontSize: 30, fontWeight: 400, color: '#F5F5F3', lineHeight: 1.2, marginBottom: 18 }}>{project.title}</h3>
          <p style={{ margin: 0, fontFamily: 'var(--fc-font-display)', fontStyle: 'italic', fontSize: 16, color: '#D5D7DA', lineHeight: 1.5 }}>
            "{project.quote.text}"
            <span style={{ display: 'block', marginTop: 8, fontStyle: 'normal', fontSize: 10, letterSpacing: '0.18em', textTransform: 'uppercase', color: '#A7ADB5', fontFamily: 'var(--fc-font-body)' }}>— {project.quote.attribution}</span>
          </p>
        </div>
        <div style={{ paddingTop: 22, borderTop: '0.5px solid rgba(213,215,218, 0.2)', display: 'flex', gap: 28, flexWrap: 'wrap', marginTop: 22 }}>
          {project.results.map((r, i) => (
            <div key={i}>
              <div style={{ fontFamily: 'var(--fc-font-display)', fontSize: 28, fontWeight: 600, color: '#F5F5F3' }}>
                {r.value}<span style={{ fontSize: 14, fontWeight: 400, color: '#EE9440', marginLeft: 2 }}>{r.unit}</span>
              </div>
              <div style={{ fontFamily: 'var(--fc-font-body)', fontSize: 10, letterSpacing: '0.18em', textTransform: 'uppercase', color: '#D5D7DA', marginTop: 4 }}>{r.label}</div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

Object.assign(window, { CardLight, CapabilityCard, ProjectCard, FeaturedProjectCard });
