/* global React */
const { useState: useStateForm } = React;

function TextInput({ label, value, onChange, placeholder, type = 'text', error, optional }) {
  const [focused, setFocused] = useStateForm(false);
  return (
    <div>
      <label style={{ display: 'block', fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500, letterSpacing: '0.06em', textTransform: 'uppercase', color: '#6b5f52', marginBottom: 6 }}>
        {label} {optional && <span style={{ textTransform: 'none', letterSpacing: 0, color: '#b4ada6' }}>(optional)</span>}
      </label>
      <input
        type={type}
        value={value || ''}
        placeholder={placeholder}
        onChange={e => onChange(e.target.value)}
        onFocus={() => setFocused(true)}
        onBlur={() => setFocused(false)}
        style={{
          width: '100%',
          padding: '10px 14px',
          background: '#faf9f7',
          border: error ? '0.5px solid #b44' : `0.5px solid ${focused ? '#a84812' : '#d4cfc9'}`,
          borderRadius: 3,
          fontFamily: 'var(--fc-font-body)',
          fontSize: 16,
          color: '#1c1e24',
          outline: 'none',
          boxShadow: focused ? '0 0 0 2px rgba(168, 72, 18, 0.12)' : 'none',
          boxSizing: 'border-box',
          transition: 'all 0.15s',
        }}
      />
      {error && <p style={{ margin: '5px 0 0', fontSize: 11, letterSpacing: '0.02em', color: '#b44' }}>{error}</p>}
    </div>
  );
}

function TextArea({ label, value, onChange, placeholder, error }) {
  const [focused, setFocused] = useStateForm(false);
  return (
    <div>
      <label style={{ display: 'block', fontFamily: 'var(--fc-font-body)', fontSize: 11, fontWeight: 500, letterSpacing: '0.06em', textTransform: 'uppercase', color: '#6b5f52', marginBottom: 6 }}>{label}</label>
      <textarea
        value={value || ''}
        placeholder={placeholder}
        onChange={e => onChange(e.target.value)}
        onFocus={() => setFocused(true)}
        onBlur={() => setFocused(false)}
        rows={5}
        style={{
          width: '100%',
          padding: '10px 14px',
          background: '#faf9f7',
          border: `0.5px solid ${focused ? '#a84812' : '#d4cfc9'}`,
          borderRadius: 3,
          fontFamily: 'var(--fc-font-body)',
          fontSize: 16,
          lineHeight: 1.6,
          color: '#1c1e24',
          outline: 'none',
          resize: 'vertical',
          boxShadow: focused ? '0 0 0 2px rgba(168, 72, 18, 0.12)' : 'none',
          boxSizing: 'border-box',
        }}
      />
      {error && <p style={{ margin: '5px 0 0', fontSize: 11, color: '#b44' }}>{error}</p>}
    </div>
  );
}

function ContactForm() {
  const [values, setValues] = useStateForm({ name: '', business: '', email: '', phone: '', website: '', message: '' });
  const [submitted, setSubmitted] = useStateForm(false);
  const set = k => v => setValues(p => ({ ...p, [k]: v }));
  if (submitted) {
    return (
      <div style={{ background: '#f0ece5', borderLeft: '2.5px solid #a84812', borderRadius: 0, padding: 28 }}>
        <p style={{ margin: 0, fontFamily: 'var(--fc-font-body)', fontSize: 17, lineHeight: 1.65, color: '#1c1e24' }}>
          Got it. You'll hear from Logan. Not a bot. Not a VA. The person who will actually build your project is reading this right now.
        </p>
      </div>
    );
  }
  return (
    <form style={{ display: 'flex', flexDirection: 'column', gap: 20 }} onSubmit={e => { e.preventDefault(); setSubmitted(true); }}>
      <TextInput label="Name" value={values.name} onChange={set('name')} placeholder="Your name" />
      <TextInput label="Business name" value={values.business} onChange={set('business')} placeholder="Your business" />
      <TextInput label="Email" type="email" value={values.email} onChange={set('email')} placeholder="you@example.com" />
      <TextInput label="Phone" type="tel" value={values.phone} onChange={set('phone')} placeholder="(406) 555-0199" optional />
      <TextInput label="Current website URL" value={values.website} onChange={set('website')} placeholder="yourbusiness.com" optional />
      <TextArea label="What's frustrating you about your current setup?" value={values.message} onChange={set('message')} placeholder="Tell us what's not working and what you'd like to change." />
      <button type="submit" style={{
        alignSelf: 'flex-start',
        background: '#a84812', color: '#faf9f7',
        border: 'none', borderRadius: 3,
        padding: '12px 26px',
        fontFamily: 'var(--fc-font-body)', fontSize: 13, fontWeight: 500, letterSpacing: '0.08em', textTransform: 'uppercase',
        cursor: 'pointer',
      }}>Send it</button>
    </form>
  );
}

Object.assign(window, { TextInput, TextArea, ContactForm });
