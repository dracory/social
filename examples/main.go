package main

import (
	"fmt"
	"net/http"

	"github.com/dracory/social"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shareLinks := social.NewQuick("https://example.com/my-article", "Check out this awesome article!", "https://example.com/cover.jpg")

		// 1. Default Widget (Bootstrap icons)
		defaultWidget := shareLinks.Widget(social.WidgetOptions{
			ShareText:   "Share this:",
			IconLibrary: social.IconLibraryBootstrap,
		})

		// 2. Custom Platforms Widget
		customWidget := shareLinks.Widget(social.WidgetOptions{
			Platforms: []string{
				social.PlatformFacebook,
				social.PlatformTwitter,
				social.PlatformLinkedIn,
				social.PlatformWhatsApp,
				social.PlatformReddit,
				social.PlatformTelegram,
				social.PlatformEmail,
				social.PlatformCopyLink,
			},
			IconLibrary: social.IconLibraryBootstrap,
			ShareText:   "Popular platforms:",
		})

		// 3. Font Awesome 4.7 Icons Widget (matching class names like fa-facebook-official, fa-twitter, etc.)
		faWidget := shareLinks.Widget(social.WidgetOptions{
			Platforms: []string{
				social.PlatformFacebook,
				social.PlatformTwitter,
				social.PlatformLinkedIn,
				social.PlatformPinterest,
				social.PlatformReddit,
			},
			IconLibrary: social.IconLibraryFontAwesome,
			ShareText:   "Font Awesome icons:",
		})

		html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Social Media Share Links Widget Example</title>
    <!-- Bootstrap Icons CSS -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.3/font/bootstrap-icons.min.css">
    <!-- Font Awesome 4.7 CSS -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
            background-color: #f8fafc;
            color: #1e293b;
            margin: 0;
            padding: 40px 20px;
            display: flex;
            justify-content: center;
        }
        .container {
            max-width: 720px;
            width: 100%%;
            background: #ffffff;
            padding: 32px;
            border-radius: 12px;
            box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -2px rgba(0, 0, 0, 0.1);
        }
        h1 {
            font-size: 1.75rem;
            margin-top: 0;
            margin-bottom: 8px;
            color: #0f172a;
        }
        p.subtitle {
            color: #64748b;
            margin-bottom: 32px;
            font-size: 1rem;
        }
        .card {
            background-color: #f1f5f9;
            border-radius: 8px;
            padding: 20px;
            margin-bottom: 24px;
        }
        .card h2 {
            font-size: 1.1rem;
            margin-top: 0;
            margin-bottom: 12px;
            color: #334155;
        }
        /* Widget Styling */
        #social-links {
            display: flex;
            align-items: center;
            flex-wrap: wrap;
            gap: 12px;
        }
        #social-links .share-text {
            font-weight: 600;
            color: #475569;
            font-size: 0.95rem;
        }
        #social-links ul {
            list-style: none;
            padding: 0;
            margin: 0;
            display: flex;
            flex-wrap: wrap;
            gap: 10px;
            align-items: center;
        }
        #social-links li {
            display: inline-block;
            margin: 0;
        }
        .social-button {
            display: inline-flex;
            align-items: center;
            justify-content: center;
            width: 42px;
            height: 42px;
            border-radius: 50%%;
            background-color: #ffffff;
            color: #0f172a;
            text-decoration: none;
            font-size: 1.25rem;
            box-shadow: 0 1px 3px rgba(0,0,0,0.1);
            transition: all 0.2s ease-in-out;
        }
        .social-button:hover {
            transform: translateY(-2px);
            box-shadow: 0 4px 6px rgba(0,0,0,0.15);
            background-color: #2563eb;
            color: #ffffff;
        }
        /* Font Awesome 4 requires .fa base class styling if not provided */
        #social-links span[class^="fa-"], #social-links span[class*=" fa-"] {
            font-family: FontAwesome;
            font-style: normal;
            font-weight: normal;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Social Share Widget Preview</h1>
        <p class="subtitle">Go package <code>github.com/dracory/social</code> demo page.</p>

        <div class="card">
            <h2>1. Default Widget (Bootstrap Icons)</h2>
            %s
        </div>

        <div class="card">
            <h2>2. Custom Selected Platforms</h2>
            %s
        </div>

        <div class="card">
            <h2>3. Font Awesome Icon Library</h2>
            %s
        </div>
    </div>
</body>
</html>
`, defaultWidget, customWidget, faWidget)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(html))
	})

	fmt.Println("Server running at http://localhost:8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
