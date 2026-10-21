/* global React */
function Nav({ active, onNav }) {
  const links = [['work', 'Work'], ['services', 'Services'], ['process', 'Process'], ['about', 'About']];
  return (
    <nav style={{ background: '#17181A' }}>
      <div style={{
        maxWidth: 1100, margin: '0 auto',
        display: 'flex', justifyContent: 'space-between', alignItems: 'center',
        padding: '12px 32px',
      }}>
        <a onClick={() => onNav('home')} style={{ display: 'flex', alignItems: 'center', gap: 12, cursor: 'pointer', textDecoration: 'none' }}>
          <img src="../../assets/logo.svg" alt="FlintCraft Studio" style={{ height: 32, width: 'auto' }} />
          <div>
            <div style={{ fontFamily: 'var(--fc-font-body)', fontSize: 13, fontWeight: 500, letterSpacing: '0.12em', textTransform: 'uppercase', color: '#F5F5F3' }}>
              FlintCraft Studio
            </div>
            <div style={{ fontFamily: 'var(--fc-font-display)', fontStyle: 'italic', fontSize: 11, letterSpacing: '0.04em', color: '#A7ADB5', marginTop: 2 }}>
              Still shaped by hand.
            </div>
          </div>
        </a>
        <div style={{ display: 'flex', alignItems: 'center', gap: 24 }}>
          {links.map(([k, label]) => (
            <a key={k} onClick={() => onNav(k)} style={{
              fontFamily: 'var(--fc-font-body)', fontSize: 12, letterSpacing: '0.06em',
              textTransform: 'uppercase', textDecoration: 'none', cursor: 'pointer',
              color: active === k ? '#F5F5F3' : '#A7ADB5',
              transition: 'color 0.2s',
            }}>{label}</a>
          ))}
          <a onClick={() => onNav('contact')} style={{
            fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500, letterSpacing: '0.08em',
            textTransform: 'uppercase', color: '#EE9440', cursor: 'pointer',
            border: '0.5px solid #EE9440', borderRadius: 3, padding: '5px 14px', textDecoration: 'none',
          }}>Start a project</a>
        </div>
      </div>
    </nav>
  );
}

function Footer() {
  return (
    <footer style={{ background: '#17181A', paddingTop: 60, paddingBottom: 40 }}>
      <div style={{ maxWidth: 1100, margin: '0 auto', padding: '0 32px' }}>
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', gap: 40, flexWrap: 'wrap' }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 12 }}>
            <img src="../../assets/logo.svg" alt="" style={{ height: 32, width: 'auto' }} />
            <div>
              <div style={{ fontFamily: 'var(--fc-font-body)', fontSize: 13, fontWeight: 500, letterSpacing: '0.12em', textTransform: 'uppercase', color: '#F5F5F3' }}>FlintCraft Studio</div>
              <div style={{ fontFamily: 'var(--fc-font-display)', fontStyle: 'italic', fontSize: 11, letterSpacing: '0.04em', color: '#A7ADB5', marginTop: 2 }}>Still shaped by hand.</div>
            </div>
          </div>
          <div style={{ display: 'flex', gap: 32, flexWrap: 'wrap' }}>
            {['Services', 'Work', 'Process', 'About', 'Start a project'].map(l => (
              <a key={l} style={{ fontFamily: 'var(--fc-font-body)', fontSize: 13, letterSpacing: '0.08em', textTransform: 'uppercase', color: '#A7ADB5', textDecoration: 'none' }}>{l}</a>
            ))}
          </div>
        </div>
        <div style={{ marginTop: 40, paddingTop: 28, borderTop: '0.5px solid #3B3E43', display: 'flex', gap: 32, flexWrap: 'wrap', color: '#A7ADB5', fontSize: 14 }}>
          <span>(406) 871-9875</span>
          <span>hello@flintcraftstudio.com</span>
          <span>Helena, Montana</span>
        </div>
        <div style={{ marginTop: 28, display: 'flex', justifyContent: 'space-between', fontFamily: 'var(--fc-font-body)', fontSize: 11, letterSpacing: '0.08em', textTransform: 'uppercase', color: '#A7ADB5' }}>
          <span>© 2026 FlintCraft Studio. All rights reserved.</span>
          <div style={{ display: 'flex', gap: 18 }}>
            <a style={{ color: '#A7ADB5', textDecoration: 'none' }}>Privacy</a>
            <a style={{ color: '#A7ADB5', textDecoration: 'none' }}>Terms</a>
          </div>
        </div>
      </div>
    </footer>
  );
}

Object.assign(window, { Nav, Footer });
