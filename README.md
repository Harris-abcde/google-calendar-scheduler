# Google Calendar Event Scheduler

![Go](https://img.shields.io/badge/Go-1.20+-blue)
![Gin](https://img.shields.io/badge/Gin-Gonic-green)
![OAuth2](https://img.shields.io/badge/OAuth2-Google-red)
![Calendar API](https://img.shields.io/badge/Google%20Calendar-API-yellow)

A Proof of Concept application demonstrating Google Calendar integration with OAuth2 authentication and automated event scheduling.

## 🚀 Features

- 🔐 Google OAuth 2.0 authentication flow
- 📅 Automated calendar event creation
- 🎥 Google Meet link generation
- ⏱️ Timezone-aware scheduling (Asia/Kolkata default)
- 🏗️ MVC architecture with clean code structure

## 🛠️ Tech Stack

| Component       | Technology                          |
|-----------------|-------------------------------------|
| Backend         | Go 1.20+                           |
| Web Framework   | Gin                                 |
| Authentication  | Google OAuth2                       |
| Calendar API    | Google Calendar v3                  |
| Configuration   | godotenv                            |



## Authorized Redirect URI:

http://localhost:8080/auth/google/callback


## Create .env:

GOOGLE_CLIENT_ID=your_client_id.apps.googleusercontent.com

GOOGLE_CLIENT_SECRET=your_client_secret

REDIRECT_URL=http://localhost:8080/auth/google/callback

PORT=8080

##Access the application at:

http://localhost:8080/auth/google

## 🌐 API Endpoints
Endpoint	Method	Description

/auth/google	 - GET	Initiates OAuth2 flow

/auth/google/callback	 - GET	Handles OAuth2 callback

