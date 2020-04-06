package chrome

const (
	// Exclude this switch will set window.navigator.webdriver to undefined
	// but not work on ChromeDriver 79.0.3945.16 and above
	// see https://chromedriver.chromium.org/downloads
	SWITCH_ENABLE_AUTOMATION = "enable-automation"
)

// Capabilities for ChromeDriver
// See https://chromedriver.chromium.org/capabilities
type Capabilities struct {
	// List of command-line arguments to use when starting Chrome.
	// Arguments with an associated value should be separated by a '=' sign
	// (e.g., ['start-maximized', 'user-data-dir=/tmp/temp_profile']).
	// See here(http://peter.sh/experiments/chromium-command-line-switches/) for a list of Chrome arguments.
	Args []string `json:"args,omitempty"`

	// Path to the Chrome executable to use
	// (on Mac OS X, this should be the actual binary, not just the app.
	// e.g., '/Applications/Google Chrome.app/Contents/MacOS/Google Chrome')
	Binary string `json:"binary,omitempty"`

	// A list of Chrome extensions to install on startup.
	// Each item in the list should be a base-64 encoded packed Chrome extension (.crx)
	Extensions []string `json:"extensions,omitempty"`

	// A dictionary with each entry consisting of the name of the preference and its value.
	// These preferences are only applied to the user profile in use.
	// See the 'Preferences' file in Chrome's user data directory for examples.
	LocalState map[string]string `json:"localState,omitempty"`

	// If false, Chrome will be quit when ChromeDriver is killed, regardless of whether the session is quit.
	// If true, Chrome will only be quit if the session is quit (or closed).
	// Note, if true, and the session is not quit, ChromeDriver cannot clean up the temporary user data
	// directory that the running Chrome instance is using.
	Detach bool `json:"detach,omitempty"`

	// An address of a Chrome debugger server to connect to,
	// in the form of <hostname/ip:port>, e.g. '127.0.0.1:38947'
	DebuggerAddress string `json:"debuggerAddress,omitempty"`

	// List of Chrome command line switches to exclude that ChromeDriver by default passes when starting Chrome.
	// Do not prefix switches with --.
	ExcludeSwitches []string `json:"excludeSwitches,omitempty"`

	// Directory to store Chrome minidumps . (Supported only on Linux.)
	MinidumpPath string `json:"minidumpPath,omitempty"`

	// A dictionary with either a value for "deviceName" or values for "deviceMetrics" and "userAgent"
	// Refer to Mobile Emulation for more information.
	MobileEmulation *MobileEmulation `json:"mobileEmulation,omitempty"`

	// An optional dictionary that specifies performance logging preferences. See below for more information.
	PerfLoggingPrefs map[string]string `json:"perfLoggingPrefs,omitempty"`

	// A list of window types that will appear in the list of window handles.
	// For access to <webview> elements, include "webview" in this list.
	WindowTypes []string `json:"windowTypes,omitempty"`
}

// MobileEmulation enables the Mobile Emulation feature in Chrome DevTools.
// Either a value for "deviceName" or values for "deviceMetrics" and "userAgent"
// https://chromedriver.chromium.org/mobile-emulation
type MobileEmulation struct {
	// DeviceName is known device name from DevTools Emulation panel. e.g. "Nexus 5"
	DeviceName string `json:"deviceName,omitempty"`

	// DeviceMetrics specifies individual device attributes
	DeviceMetrics *DeviceMetrics `json:"deviceMetrics,omitempty"`

	// UserAgent specifies the user agent of the individual device
	UserAgent string `json:"userAgent,omitempty"`
}

// DeviceMetrics specifies individual device attributes.
type DeviceMetrics struct {
	// Width is the width of the device screen.
	Width uint `json:"width"`

	// Height is the height of the device screen.
	Height uint `json:"height"`

	// PixelRatio is the device’s pixel ratio.
	PixelRatio float64 `json:"pixelRatio"`

	// Touch indicates whether to emulate touch events (defaults to true, usually does not need to be set)
	Touch *bool `json:"touch,omitempty"`
}
