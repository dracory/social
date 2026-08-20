# Proposal for Further Improvements to `dracory/social`

## Executive Summary
This proposal outlines recommended future architecture enhancements, new features, and developer experience improvements for the `dracory/social` Go library.

---

## 1. Extension Architecture & Dynamic Custom Platforms

### Current State
Currently, platforms and their share URL generators, display names, brand colors, and icon classes are defined statically across `constants.go`, `functions.go`, `share_links.go`, and `widget.go`.

### Proposed Enhancement
Introduce a dynamic platform registration mechanism allowing consumers to register custom or enterprise platforms at runtime without modifying library source code:

```go
type CustomPlatform struct {
    ID            string
    NiceName      string
    Color         string
    FontAwesome   string
    BootstrapIcon string
    GetURL        func(params ShareLinksParams) string
}

func RegisterPlatform(cp CustomPlatform) error
```

---

## 2. Advanced Analytics & UTM Campaign Tracking

### Proposed Enhancement
Provide built-in support for appending UTM tracking parameters (e.g., `utm_source`, `utm_medium`, `utm_campaign`, `utm_content`) automatically to generated share links.

```go
params := social.ShareLinksParams{
    URL:          "https://example.com/page",
    UTMCampaign:  "spring_sale",
    UTMMedium:    "social_share",
    AutoUTMSource: true, // sets utm_source=<platform_id>
}
```

---

## 3. SVG Icon Support & Templating Flexibility

### Proposed Enhancement
Extend `WidgetOptions` to support embedded inline SVG icons or custom HTML templates:

```go
type WidgetOptions struct {
    IconStyle      string // "bootstrap", "fontawesome", "inline-svg"
    CustomClass    string
    ContainerTag   string // "div", "nav", "section"
}
```

---

## 4. URL Shortener Integration Interface

### Proposed Enhancement
Integrate an optional asynchronous or synchronous URL shortener interface (e.g., Bitly, TinyURL, or custom service) into `ShareLinksParams`.

```go
type URLShortener interface {
    Shorten(ctx context.Context, longURL string) (string, error)
}
```

---

## 5. UI Framework / Component Adaptors

### Proposed Enhancement
Provide lightweight wrapper packages or export clean JSON/struct representations to facilitate integration with modern frontend frameworks (React, Vue, Svelte, Templ, Go html/template):

```go
type ShareButton struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    URL       string `json:"url"`
    IconClass string `json:"icon_class"`
    Color     string `json:"color"`
}

func (s *ShareLinks) ExportButtons(opts WidgetOptions) []ShareButton
```
