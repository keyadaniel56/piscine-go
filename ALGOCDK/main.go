package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
    "strings"
)

var (
    users     = map[string]string{} // user email -> password (for demo only)
    sessions  = map[string]string{} // sessionID -> email
    templates *template.Template    // Initialized below in init()
)

// Bot represents a trading bot's data
type Bot struct {
    Name        string
    Image       string
    ROI         string
    Platform    string
    Rating      string
    Creator     string
    Description string
    Price       string
}

// BotData holds the data for different sections
type BotData struct {
    FeaturedBots []Bot
    TrendingBots []Bot
}

// Initialize templates with custom functions
func init() {
    funcMap := template.FuncMap{
        "add": func(a, b int) int {
            return a + b
        },
    }
    var err error
    templates, err = template.New("").Funcs(funcMap).ParseGlob("*.html")
    if err != nil {
        log.Fatalf("Failed to parse templates: %v", err)
    }
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/register", registerHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/dashboard", dashboardHandler)
    http.HandleFunc("/logout", logoutHandler)
    http.HandleFunc("/bot/", botHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    fmt.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func getBotData() BotData {
    return BotData{
        FeaturedBots: []Bot{
            {
                Name:        "Gold Scalper Pro",
                Image:       "/static/Images/bit.jpg",
                ROI:         "+82% ROI",
                Platform:    "Deriv",
                Rating:      "4.8",
                Creator:     "QuantAlgo",
                Description: "Advanced scalping bot for XAU/USD with 85% win rate",
                Price:       "$149",
            },
            {
                Name:        "Bitcoin Trend Rider",
                Image:       "/static/Images/bit.jpg",
                ROI:         "+65% ROI",
                Platform:    "Deriv",
                Rating:      "4.6",
                Creator:     "CryptoBots",
                Description: "Follows BTC/USD trends with AI-powered signals",
                Price:       "$199",
            },
            {
                Name:        "Forex Hedge Master",
                Image:      "/static/Images/bit.jpg",
                ROI:         "+57% ROI",
                Platform:    "Deriv",
                Rating:      "4.9",
                Creator:     "ForexQuant",
                Description: "Advanced hedging strategy for major currency pairs",
                Price:       "$249",
            },
            {
                Name:        "Oil Surfer Pro",
                Image:      "/static/Images/bit.jpg",
                ROI:         "+73% ROI",
                Platform:    "Deriv",
                Rating:      "4.7",
                Creator:     "CommodityAlgo",
                Description: "Trades crude oil with wave pattern recognition",
                Price:       "$179",
            },
        },
        TrendingBots: []Bot{
            {
                Name:        "Index Master",
                Image:       "/static/Images/bit.jpg",
                ROI:         "+68% ROI",
                Platform:    "Deriv",
                Rating:      "4.9",
                Creator:     "IndexTraders",
                Description: "Trades major indices with 80% accuracy",
                Price:       "$299",
            },
        },
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    data := getBotData()
    err := templates.ExecuteTemplate(w, "home.html", data)
    if err != nil {
        http.Error(w, "Template execution failed", http.StatusInternalServerError)
        log.Printf("Template error: %v", err)
    }
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        email := strings.TrimSpace(r.FormValue("email"))
        password := strings.TrimSpace(r.FormValue("password"))

        if email == "" || password == "" {
            http.Error(w, "Email and password required", http.StatusBadRequest)
            return
        }
        if _, exists := users[email]; exists {
            http.Error(w, "User already exists", http.StatusBadRequest)
            return
        }
        users[email] = password
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }
    err := templates.ExecuteTemplate(w, "register.html", nil)
    if err != nil {
        http.Error(w, "Template execution failed", http.StatusInternalServerError)
        log.Printf("Template error: %v", err)
    }
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        email := strings.TrimSpace(r.FormValue("email"))
        password := strings.TrimSpace(r.FormValue("password"))

        if storedPass, ok := users[email]; ok && storedPass == password {
            sessionID := fmt.Sprintf("%d", len(sessions)+1)
            sessions[sessionID] = email
            cookie := &http.Cookie{
                Name:  "session",
                Value: sessionID,
                Path:  "/",
            }
            http.SetCookie(w, cookie)
            http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
            return
        }
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }
    err := templates.ExecuteTemplate(w, "login.html", nil)
    if err != nil {
        http.Error(w, "Template execution failed", http.StatusInternalServerError)
        log.Printf("Template error: %v", err)
    }
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session")
    if err != nil || sessions[cookie.Value] == "" {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }
    email := sessions[cookie.Value]
    data := struct {
        Email string
        Bots  BotData
    }{
        Email: email,
        Bots:  getBotData(),
    }
    err = templates.ExecuteTemplate(w, "dashboard.html", data)
    if err != nil {
        http.Error(w, "Template execution failed", http.StatusInternalServerError)
        log.Printf("Template error: %v", err)
    }
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session")
    if err == nil {
        delete(sessions, cookie.Value)
        cookie.MaxAge = -1
        cookie.Path = "/"
        http.SetCookie(w, cookie)
    }
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func botHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session")
    if err != nil || sessions[cookie.Value] == "" {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }
    email := sessions[cookie.Value]

    // Extract bot ID from URL (e.g., /bot/0)
    idStr := strings.TrimPrefix(r.URL.Path, "/bot/")
    id, err := strconv.Atoi(idStr)
    if err != nil || id < 0 {
        http.Error(w, "Invalid bot ID", http.StatusBadRequest)
        return
    }

    botData := getBotData()
    var selectedBot Bot
    var botType string

    if id < len(botData.FeaturedBots) {
        selectedBot = botData.FeaturedBots[id]
        botType = "Featured"
    } else if id < len(botData.FeaturedBots)+len(botData.TrendingBots) {
        selectedBot = botData.TrendingBots[id-len(botData.FeaturedBots)]
        botType = "Trending"
    } else {
        http.Error(w, "Bot not found", http.StatusNotFound)
        return
    }

    data := struct {
        Email   string
        Bot     Bot
        BotType string
    }{
        Email:   email,
        Bot:     selectedBot,
        BotType: botType,
    }
    err = templates.ExecuteTemplate(w, "bot.html", data)
    if err != nil {
        http.Error(w, "Template execution failed", http.StatusInternalServerError)
        log.Printf("Template error: %v", err)
    }
}