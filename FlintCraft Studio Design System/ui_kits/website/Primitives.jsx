/* global React */
const { useState } = React;

function Eyebrow({ children, dark = false }) {
  return (
    <div style={{ marginBottom: '1rem' }}>
      <div style={{ width: 28, height: 1, background: '#a84812', marginBottom: 8 }} />
      <span style={{
        fontFamily: 'var(--fc-font-body)',
        fontSize: 11, fontWeight: 500, letterSpacing: '0.2em',
        textTransform: 'uppercase',
        color: dark ? '#96897a' : '#6b5f52',
      }}>{children}</span>
    </div>
  );
}

function RuleAccent() {
  return <hr style={{ border: 'none', borderTop: '1px solid #a84812', width: 48, margin: '1.5rem 0' }} />;
}

function PullQuote({ children, dark = false }) {
  return (
    <p style={{
      fontFamily: 'var(--fc-font-display)',
      fontStyle: 'italic',
      fontSize: 18,
      letterSpacing: '0.02em',
      color: dark ? '#e8873b' : '#a84812',
      lineHeight: 1.4,
      margin: 0,
    }}>{children}</p>
  );
}

/* Flint-geometry button — chipped top-right corner. */
function Button({ variant = 'primary', dark = false, children, onClick, size = 'md' }) {
  const pad = size === 'lg' ? '14px 26px' : '11px 22px';
  if (variant === 'primary') {
    return (
      <a onClick={onClick}
        className={`fc-flint-btn fc-chip-btn ${dark ? 'fc-flint-btn--primary-dark' : 'fc-flint-btn--primary-light'}`}
        style={{ padding: pad }}>
        {children}
      </a>
    );
  }
  if (variant === 'outline') {
    return (
      <a onClick={onClick}
        className={`fc-flint-btn fc-chip-btn ${dark ? 'fc-flint-btn--outline-dark' : 'fc-flint-btn--outline-light'}`}
        style={{ padding: pad }}>
        {children}
      </a>
    );
  }
  // ghost — underline, no chip
  return (
    <a onClick={onClick} style={{
      fontFamily: 'var(--fc-font-body)', fontSize: 13, fontWeight: 500,
      letterSpacing: '0.08em', textTransform: 'uppercase',
      color: dark ? '#e8873b' : '#a84812',
      borderBottom: `1.5px solid ${dark ? '#e8873b' : '#a84812'}`,
      padding: '10px 0', textDecoration: 'none', cursor: 'pointer',
      transition: 'opacity 0.2s',
    }}>{children}</a>
  );
}

function GhostLink({ dark = false, children, onClick }) {
  return <Button variant="ghost" dark={dark} onClick={onClick}>{children}</Button>;
}

Object.assign(window, { Eyebrow, RuleAccent, PullQuote, Button, GhostLink });
