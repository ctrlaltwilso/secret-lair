# GoLang: Secret Lair Alerts

**Go Secret Lair Alerts** is a Go-based automation tool designed to ensure you never miss a MTG Secret Lair drop. By converting official email announcements into a structured notification pipeline, it keeps you informed every step of the way.

- **Instant Announcement Alerts:** Automatically processes incoming email announcements to notify you of new drop reveals.
- **Pre-Drop Reminders:** Opt-in to specific drops to receive a heads-up notification 24 hours before sales go live.
- **Drop Day Tracking:** Delivers scheduled, timely alerts throughout launch day so you can secure your order before stock runs out.

## Status

This is a personal, in-progress learning project — the feature list above is the target, not what's built yet. Here's where things actually stand:

- [x] Connect to Gmail over IMAP and poll a dedicated label for new mail
- [ ] Parse Secret Lair drop details (name, release date, link) out of announcement emails
- [ ] Discord bot integration for alerts
- [ ] RSVP / opt-in tracking per drop
- [ ] 24-hour and drop-day reminder alerts

Right now, running this tool connects to your Gmail account, selects the label you configure, and logs the subject line of the latest message in it. Parsing, alerting, and reminders aren't implemented yet.

## Getting Started

### Prerequisites

- Go (see `go.mod` for the version this is built against)
- A Gmail account
- 2-Step Verification enabled on that account, plus an [App Password](https://myaccount.google.com/apppasswords) — your normal Gmail password won't work over IMAP
- A Gmail label dedicated to Secret Lair emails, with **Show in IMAP** enabled for it (Gmail Settings → Labels)

### Setup

1. Clone this repository
2. Copy `.env.example` to `.env` and fill in your values:
   ```sh
   cp .env.example .env
   ```
3. Run it:
   ```sh
   go run main.go
   ```

## License

MIT — see [LICENSE](./LICENSE).
