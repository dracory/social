package social

import (
	"testing"
)

func TestFormatTextURL_EdgeCases(t *testing.T) {
	if got := formatTextURL("", "https://example.com"); got != "https://example.com" {
		t.Errorf("Expected https://example.com, got %q", got)
	}

	if got := formatTextURL("Title Only", ""); got != "Title Only" {
		t.Errorf("Expected 'Title Only', got %q", got)
	}

	if got := formatTextURL("   ", "   "); got != "" {
		t.Errorf("Expected empty string, got %q", got)
	}
}

func TestWidget_UnknownPlatformAndFontAwesome(t *testing.T) {
	s := NewQuick("https://example.com", "Title", "")

	// Test unknown platform (should be ignored gracefully)
	html := s.Widget(WidgetOptions{
		Platforms:   []string{"unknown_platform", PlatformFacebook},
		IconLibrary: IconLibraryFontAwesome,
	})

	if !testing.Short() && len(html) == 0 {
		t.Error("Widget HTML should not be empty")
	}

	if fontAwesomeFacebook := FontAwesomeFacebook; fontAwesomeFacebook != "" {
		if !containsString(html, fontAwesomeFacebook) {
			t.Errorf("Widget should contain FontAwesome icon %s", fontAwesomeFacebook)
		}
	}
}

func containsString(s, substr string) bool {
	return len(s) >= len(substr) && searchSubstr(s, substr)
}

func searchSubstr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestSortPlatforms_EdgeCases(t *testing.T) {
	// Unknown platform sorting
	sortedAlpha := sortPlatformsAlphabetically([]string{"unknown1", "unknown2", PlatformFacebook})
	if len(sortedAlpha) != 3 {
		t.Errorf("Expected 3 items, got %d", len(sortedAlpha))
	}

	sortedPop := sortPlatformsByPopularity([]string{"unknown1", "unknown2", PlatformFacebook})
	if len(sortedPop) != 3 {
		t.Errorf("Expected 3 items, got %d", len(sortedPop))
	}
}
